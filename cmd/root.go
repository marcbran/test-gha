package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsevents"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "test-gha",
	Short: "Test application for GitHub Actions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World!")

		stream := fsevents.EventStream{}
		err := stream.Start()
		if err != nil {
			log.Fatal(err)
		}
		stream.Stop()

		// Create new watcher.
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		defer watcher.Close()

		// Start listening for events.
		go func() {
			for {
				select {
				case event, ok := <-watcher.Events:
					if !ok {
						return
					}
					log.Println("event:", event)
					if event.Has(fsnotify.Write) {
						log.Println("modified file:", event.Name)
					}
				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}
					log.Println("error:", err)
				}
			}
		}()

		// Add a path.
		err = watcher.Add("/tmp")
		if err != nil {
			log.Fatal(err)
		}

		// Block main goroutine forever.
		<-make(chan struct{})
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
