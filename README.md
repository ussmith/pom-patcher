# pom-patcher

This is simple WIP used to version dependencies.  It started due to the log4j problem, and I thought it was fitting to write something to fix Java in Golang.

First, it identifies the dependency you've requested, determines whether it's an inline version or a property.  Once that is known, it updates the value and rewrites 
the pom.xml back out.

Usage of ./pom-patcher:
  -dep string
    	dependency name
  -ver string
    	version
  -start
      directory to start the crawl
