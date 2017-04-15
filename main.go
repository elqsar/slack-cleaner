package main

import (
	"flag"
	"fmt"
	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	token = flag.String("token", "", "Slack token")
)

func init() {
	log.SetOutput(os.Stderr)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	flag.Parse()

	api := slack.New(*token)

	// ignore pagination
	files, _, err := api.GetFiles(slack.NewGetFilesParameters())
	if err != nil {
		log.Error("Error getting files info: ", err)
	}

	for _, f := range files {
		fmt.Printf("File: ID:%s Name:%s\n", f.ID, f.Name)
		if err := api.DeleteFile(f.ID); err != nil {
			log.Errorf("Error deleting file: %s %s", f.ID, err)
		}
	}
}
