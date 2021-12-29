package data

import (
	"encoding/xml"
	"io"
)

type Project struct {
	XMLName        xml.Name `xml:"project"`
	Text           string   `xml:",chardata"`
	Xmlns          string   `xml:"xmlns,attr"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	ModelVersion   string   `xml:"modelVersion"`
	Parent         struct {
		Text       string `xml:",chardata"`
		GroupId    string `xml:"groupId"`
		ArtifactId string `xml:"artifactId"`
		Version    string `xml:"version"`
	} `xml:"parent"`
	GroupId    string `xml:"groupId"`
	ArtifactId string `xml:"artifactId"`
	Version    string `xml:"version"`
	Name       string `xml:"name"`
	Packaging  string `xml:"packaging"`
	//Properties []interface{} `xml:"properties"`
	Properties Props `xml:"properties"`
	Modules    struct {
		Text   string   `xml:",chardata"`
		Module []string `xml:"module"`
	} `xml:"modules"`
	DependencyManagement struct {
		Text         string `xml:",chardata"`
		Dependencies struct {
			Text       string `xml:",chardata"`
			Dependency []struct {
				Text       string `xml:",chardata"`
				GroupId    string `xml:"groupId"`
				ArtifactId string `xml:"artifactId"`
				Version    string `xml:"version"`
				Type       string `xml:"type"`
				Scope      string `xml:"scope"`
			} `xml:"dependency"`
		} `xml:"dependencies"`
	} `xml:"dependencyManagement"`
	Build struct {
		Text    string `xml:",chardata"`
		Plugins struct {
			Text   string `xml:",chardata"`
			Plugin []struct {
				Text       string `xml:",chardata"`
				GroupId    string `xml:"groupId"`
				ArtifactId string `xml:"artifactId"`
				Version    string `xml:"version"`
				Executions struct {
					Text      string `xml:",chardata"`
					Execution []struct {
						Text  string `xml:",chardata"`
						Phase string `xml:"phase"`
						Goals struct {
							Text string `xml:",chardata"`
							Goal string `xml:"goal"`
						} `xml:"goals"`
						ID string `xml:"id"`
					} `xml:"execution"`
				} `xml:"executions"`
				Configuration struct {
					Text                string `xml:",chardata"`
					ShortRevisionLength string `xml:"shortRevisionLength"`
					DoCheck             string `xml:"doCheck"`
					DoUpdate            string `xml:"doUpdate"`
				} `xml:"configuration"`
			} `xml:"plugin"`
		} `xml:"plugins"`
		PluginManagement struct {
			Text    string `xml:",chardata"`
			Plugins struct {
				Text   string `xml:",chardata"`
				Plugin []struct {
					Text          string `xml:",chardata"`
					GroupId       string `xml:"groupId"`
					ArtifactId    string `xml:"artifactId"`
					Version       string `xml:"version"`
					Configuration struct {
						Text                      string `xml:",chardata"`
						Source                    string `xml:"source"`
						Target                    string `xml:"target"`
						UseIncrementalCompilation string `xml:"useIncrementalCompilation"`
						Archive                   struct {
							Text     string `xml:",chardata"`
							Manifest struct {
								Text                            string `xml:",chardata"`
								AddDefaultImplementationEntries string `xml:"addDefaultImplementationEntries"`
								AddDefaultSpecificationEntries  string `xml:"addDefaultSpecificationEntries"`
							} `xml:"manifest"`
							ManifestEntries struct {
								Text        string `xml:",chardata"`
								BuildNumber string `xml:"Build-Number"`
								BuildDate   string `xml:"Build-Date"`
								SCMRevision string `xml:"SCM-Revision"`
								SCMBranch   string `xml:"SCM-Branch"`
							} `xml:"manifestEntries"`
						} `xml:"archive"`
						ID                    string `xml:"id"`
						Hostname              string `xml:"hostname"`
						Skip                  string `xml:"skip"`
						AutoVersionSubmodules string `xml:"autoVersionSubmodules"`
						FlowInitContext       struct {
							Text             string `xml:",chardata"`
							VersionTagPrefix string `xml:"versionTagPrefix"`
						} `xml:"flowInitContext"`
						LocalOnly string `xml:"localOnly"`
						Schemas   string `xml:"schemas"`
						Locations struct {
							Text     string   `xml:",chardata"`
							Location []string `xml:"location"`
						} `xml:"locations"`
						CleanDisabled            string `xml:"cleanDisabled"`
						BaselineVersion          string `xml:"baselineVersion"`
						Additionalparam          string `xml:"additionalparam"`
						LifecycleMappingMetadata struct {
							Text             string `xml:",chardata"`
							PluginExecutions struct {
								Text            string `xml:",chardata"`
								PluginExecution struct {
									Text                  string `xml:",chardata"`
									PluginExecutionFilter struct {
										Text         string `xml:",chardata"`
										GroupId      string `xml:"groupId"`
										ArtifactId   string `xml:"artifactId"`
										VersionRange string `xml:"versionRange"`
										Goals        struct {
											Text string `xml:",chardata"`
											Goal string `xml:"goal"`
										} `xml:"goals"`
									} `xml:"pluginExecutionFilter"`
									Action struct {
										Text   string `xml:",chardata"`
										Ignore string `xml:"ignore"`
									} `xml:"action"`
								} `xml:"pluginExecution"`
							} `xml:"pluginExecutions"`
						} `xml:"lifecycleMappingMetadata"`
						Includes struct {
							Text    string   `xml:",chardata"`
							Include []string `xml:"include"`
						} `xml:"includes"`
					} `xml:"configuration"`
					Executions struct {
						Text      string `xml:",chardata"`
						Execution struct {
							Text  string `xml:",chardata"`
							ID    string `xml:"id"`
							Goals struct {
								Text string   `xml:",chardata"`
								Goal []string `xml:"goal"`
							} `xml:"goals"`
						} `xml:"execution"`
					} `xml:"executions"`
					Inherited    string `xml:"inherited"`
					Dependencies struct {
						Text       string `xml:",chardata"`
						Dependency []struct {
							Text       string `xml:",chardata"`
							GroupId    string `xml:"groupId"`
							ArtifactId string `xml:"artifactId"`
							Version    string `xml:"version"`
						} `xml:"dependency"`
					} `xml:"dependencies"`
				} `xml:"plugin"`
			} `xml:"plugins"`
		} `xml:"pluginManagement"`
	} `xml:"build"`
	Profiles struct {
		Text    string `xml:",chardata"`
		Profile struct {
			Text  string `xml:",chardata"`
			ID    string `xml:"id"`
			Build struct {
				Text             string `xml:",chardata"`
				PluginManagement struct {
					Text    string `xml:",chardata"`
					Plugins struct {
						Text   string `xml:",chardata"`
						Plugin []struct {
							Text       string `xml:",chardata"`
							GroupId    string `xml:"groupId"`
							ArtifactId string `xml:"artifactId"`
							Version    string `xml:"version"`
							Executions struct {
								Text      string `xml:",chardata"`
								Execution struct {
									Text  string `xml:",chardata"`
									ID    string `xml:"id"`
									Phase string `xml:"phase"`
									Goals struct {
										Text string `xml:",chardata"`
										Goal string `xml:"goal"`
									} `xml:"goals"`
									Configuration struct {
										Text           string `xml:",chardata"`
										PropertiesUser struct {
											Text     string `xml:",chardata"`
											Hostname string `xml:"hostname"`
											Port     string `xml:"port"`
											Protocol string `xml:"protocol"`
										} `xml:"propertiesUser"`
									} `xml:"configuration"`
								} `xml:"execution"`
							} `xml:"executions"`
							Configuration struct {
								Text        string `xml:",chardata"`
								Source      string `xml:"source"`
								CheckResult struct {
									Text       string `xml:",chardata"`
									Throughput struct {
										Text               string `xml:",chardata"`
										Threshold          string `xml:"threshold"`
										ToleranceDirection string `xml:"toleranceDirection"`
									} `xml:"throughput"`
								} `xml:"checkResult"`
							} `xml:"configuration"`
						} `xml:"plugin"`
					} `xml:"plugins"`
				} `xml:"pluginManagement"`
			} `xml:"build"`
		} `xml:"profile"`
	} `xml:"profiles"`
	URL string `xml:"url"`
	Scm struct {
		Text                string `xml:",chardata"`
		Connection          string `xml:"connection"`
		DeveloperConnection string `xml:"developerConnection"`
		URL                 string `xml:"url"`
	} `xml:"scm"`
}

type Props struct {
	Elems map[string]Entry
}

type Entry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

type MarshalEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

type properties struct {
	Properties []MarshalEntry `xml:"properties"`
}

func (p *Props) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	e := Entry{}
	p.Elems = map[string]Entry{}
	for err = d.Decode(&e); err == nil; err = d.Decode(&e) {
		p.Elems[e.XMLName.Local] = e
	}
	if err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (p *Props) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	//type entry struct {
	//XMLName xml.Name
	//		Value   string `xml:",chardata"`
	//}

	var props []MarshalEntry
	for _, v := range p.Elems {
		v.XMLName.Space = ""
		e := MarshalEntry{
			XMLName: v.XMLName,
			Value:   v.Value,
		}

		props = append(props, e)
	}

	toMarshall := properties{
		Properties: props,
	}

	return e.Encode(toMarshall)
}
