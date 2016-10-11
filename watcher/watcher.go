package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/frezadev/photobooth/library/helper"
	"github.com/howeyc/fsnotify"
)

const (
	NOK          = "NOK"
	OK           = "OK"
	TIME_DEFAULT = "60000"
)

func main() {
	config := helper.ReadConfig()

	path := config["watcher_folder"]

	if path == NOK {
		fmt.Println(NOK + " Please provide the Path in conf/app.json")
	} else {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}

		done := make(chan bool)

		fmt.Printf("=========================> I watch you: %v\n", path)
		// Process events
		go func() {
			for {
				select {
				case ev := <-watcher.Event:
					strEV := strings.Split(""+ev.String(), ":")
					action := strings.Trim(strEV[len(strEV)-1], " ")
					if action == "CREATE" {
						processData(ev.Name)
					}
					//					else {
					//						log.Println("event:", ev)
					//					}
				case err := <-watcher.Error:
					log.Printf("error: %v\n", err)
				}
			}
		}()

		err = watcher.Watch(path)
		if err != nil {
			log.Printf("error: %v\n", err)
		}

		<-done

		/* ... do stuff ... */
		watcher.Close()
	}
}

func processData(filePath string) (success bool) {
	fmt.Println("Proccess file: " + filePath)

	if !success {
		log.Printf(NOK+" for file: %v\n", filePath)
	} else {
		log.Printf(OK+" SUCCESS for file: %v\n", filePath)
	}
	return success
}
