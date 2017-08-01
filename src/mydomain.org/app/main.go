package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	programName    string
	programVersion string
	gitCommit      string
	buildTimestamp string
	printVersion   = flag.Bool("version", false, "displays version information")
)

type server struct{}

func main() {
	flag.Parse()

	versionString := fmt.Sprintf("%s version:%s commit:%s timestamp:%s", programName, programVersion, gitCommit, buildTimestamp)
	if *printVersion {
		fmt.Println(versionString)
		os.Exit(0)
	}

	log.Info(versionString)

	log.Info("Important code goes here")
}
