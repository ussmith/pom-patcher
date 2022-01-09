package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/ussmith/crawler"
	"github.com/ussmith/pom-patcher/data"
)

var (
	dep   = flag.String("dep", "", "dependency name")
	ver   = flag.String("ver", "", "version")
	start = flag.String("start", "", "start directory - the program will crawl from here")
)

func main() {

	flag.Parse()

	if *dep == "" || *ver == "" || *start == "" {
		log.Fatal("The specific dependency, version and start directory is required")
	}

	log.Infof("dependency: %s version: %s", *dep, *ver)

	files := crawler.Find(*start, "pom.xml", crawler.Exact)
	log.Debugf("Found %v files", len(files))

	for _, f := range files {
		log.Infof("%s", f)
		patchFile(f, *dep, *ver)
	}
}

func patchFile(file, dependency, version string) {
	xmlFile, err := os.Open(file)

	if err != nil {
		log.WithError(err).Fatal("Couldn't open file")
	}

	//log.Info("Successfully Opened ./pom.xml")
	defer xmlFile.Close()

	bytes, err := ioutil.ReadAll(xmlFile)

	if err != nil {
		log.WithError(err).Fatal("Couldn't read file")
	}

	//log.Info("Successfully Read ./pom.xml")
	var Pom data.Project
	err = xml.Unmarshal(bytes, &Pom)

	if err != nil {
		log.WithError(err).Fatal("Failed to unmarshall")
	}
	//log.Info("Successfully Unmarshalled ./pom.xml")

	found := false
	var oldVersion data.Entry
	var propName string
	isProperty := false
	for i, x := range Pom.DependencyManagement.Dependencies.Dependency {
		if x.ArtifactId == *dep {
			log.Debugf("Found it, version = %s", x.Version)
			found = true
			if strings.HasPrefix(x.Version, "${") {
				log.Debugf("It's a property")
				propName = x.Version[2 : len(x.Version)-1]
				isProperty = true
				log.Debugf("Getting property %s", propName)
				oldVersion, err = getProperty(&Pom, propName)

				if err != nil {
					log.WithError(err).Error("missing property")
				}
			} else {
				oldVersion.Value = x.Version
			}

			//if oldVersion.Value != x.Version {
			if oldVersion.Value != *ver {
				//swap them and re-write the pom
				if isProperty {
					log.Debugf("Changing property %s from %s to %s", propName, Pom.Properties.Elems[propName], *ver)
					e, ok := Pom.Properties.Elems[propName]

					if !ok {
						log.Error("Failed to pull the property")
					}

					e.Value = *ver

					Pom.Properties.Elems[propName] = e
				} else {
					log.Debugf("Changing inline version from %s to %s", x.Version, *ver)
					x.Version = *ver
					log.Debugf("Version now %s", x.Version)
					Pom.DependencyManagement.Dependencies.Dependency[i] = x
				}

				err = writeFile(file, Pom)

				if err != nil {
					log.WithError(err).Error("Failed to write out the file")
				}
			}
		}
	}

	if !found {
		log.Debugf("Dependency %s was not found", *dep)
	}
}

func getProperty(pom *data.Project, property string) (data.Entry, error) {
	log.Debugf("Getting property %s", property)

	props := pom.Properties

	val, ok := props.Elems[property]

	if !ok {
		return data.Entry{}, fmt.Errorf("Unable to find property %x", property)
	}

	return val, nil
}

//func writeFile(pom data.Project, xmlFile *os.File) error {
func writeFile(file string, pom data.Project) error {
	log.Debugf("Writing the file here")

	b, err := xml.MarshalIndent(&pom, " ", " ")

	if err != nil {
		log.WithError(err).Error("Failed to marshall")
	}

	b = bytes.Replace(b, []byte("&#xA;"), []byte(""), -1)
	err = os.Truncate(file, 0)

	if err != nil {
		log.WithError(err).Error("Failed to truncate the file")
		return err
	}

	err = ioutil.WriteFile(file, b, 0644)

	if err != nil {
		log.WithError(err).Error("Failed to open the file to replace it")
	}
	//defer pomFile.Close()

	//log.Infof("Opened %s", pomFile.Name())

	//_, err = pomFile.Write(b)

	//if err != nil {
	//log.WithError(err).Error("Failed to write the file")
	//return err
	//}

	return err
}
