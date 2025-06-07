//go:build darwin

package cmd

import (
	"github.com/fsnotify/fsevents"
	"log"
)

func test() {
	stream := fsevents.EventStream{}
	err := stream.Start()
	if err != nil {
		log.Fatal(err)
	}
	stream.Stop()
}
