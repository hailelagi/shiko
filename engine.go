package main

import "fmt"

func engine() (fromEngine, toEngine chan string) {
	fmt.Println("Hello from engine")
	fromEngine = make(chan string)
	toEngine = make(chan string)

	go func() {
		for cmd := range toEngine {
			switch cmd {
			case "stop":
			case "quit":

			}
		}
	}()
	return
}
