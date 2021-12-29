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

	"github.com/ussmith/pom-patcher/data"
)

var (
	dep = flag.String("dep", "", "dependency name")
	ver = flag.String("ver", "", "version")
)

func main() {

	flag.Parse()

	if *dep == "" || *ver == "" {
		log.Fatal("The specific dependency and version is required")
	}

	log.Infof("dependency: %s version: %s", *dep, *ver)

	xmlFile, err := os.Open("./pom.xml")

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

	var oldVersion data.Entry
	var propName string
	isProperty := false
	for _, x := range Pom.DependencyManagement.Dependencies.Dependency {
		if x.ArtifactId == *dep {
			log.Infof("Found it, version = %s", x.Version)
			if strings.HasPrefix(x.Version, "${") {
				log.Info("It's a property")
				propName = x.Version[2 : len(x.Version)-1]
				isProperty = true
				oldVersion, err = getProperty(&Pom, propName)

				if err != nil {
					log.WithError(err).Error("missing property")
				}
			} else {
				oldVersion.Value = x.Version
			}
			if oldVersion.Value != x.Version {
				//swap them and re-write the pom
				if isProperty {
					log.Infof("Changing property %s from %s to %s", propName, Pom.Properties.Elems[propName], *ver)
					e, ok := Pom.Properties.Elems[propName]

					if !ok {
						log.Error("Failed to pull the property")
					}

					e.Value = *ver

					Pom.Properties.Elems[propName] = e
				} else {
					log.Infof("Changing inline version from %s to %s", x.Version, *ver)
					x.Version = *ver
				}

				//err = writeFile(Pom, xmlFile)
				err = writeFile(Pom)

				if err != nil {
					log.WithError(err).Error("Failed to write out the file")
				}
			}
		}
	}
}

func getProperty(pom *data.Project, property string) (data.Entry, error) {
	log.Infof("Getting property %s", property)

	props := pom.Properties

	val, ok := props.Elems[property]

	if !ok {
		return data.Entry{}, fmt.Errorf("Unable to find property %x", property)
	}

	return val, nil
}

//func writeFile(pom data.Project, xmlFile *os.File) error {
func writeFile(pom data.Project) error {
	log.Info("Writing the file here")

	b, err := xml.MarshalIndent(&pom, "  ", "   ")

	if err != nil {
		log.WithError(err).Error("Failed to marshall")
	}

	b = bytes.Replace(b, []byte("&#xA;"), []byte(""), -1)
	err = os.Truncate("./pom.xml", 0)

	if err != nil {
		log.WithError(err).Error("Failed to truncate the file")
		return err
	}

	err = ioutil.WriteFile("pom.xml", b, 0644)

	//pomFile, err := os.OpenFile("pom.xml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

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
