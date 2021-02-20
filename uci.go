package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var tell func(text ...string)

func uci(frGUI chan string, myTell func(text ...string)) {
	tell = myTell
	tell("info string hello world from uci")
	fromEngine, toEngine := engine()
	quit := false
	cmd := ""
	bestMove := ""

	for quit == false {
		select {
		case cmd = <-frGUI:
		case bestMove = <-fromEngine:
			handleBestMove(bestMove)
			continue
		}
		switch cmd {
		case "uci":
		case "stop":
			handleStop(toEngine)
		case "quit", "q":
			quit = true
			continue
		}
	}
	fmt.Println("Hello uci")
}

func handleBestMove(bestMove string) {
	tell(bestMove)
}

func handleStop(toEngine chan string) {
	toEngine <- "stop"
}

func mainTell(text ...string) {
	//toGUI := ""
	line := ""
	for _, t := range text {
		line += t
	}
	fmt.Println()
}

func input() chan string {
	line := make(chan string)
	go func() {
		var reader *bufio.Reader
		reader = bufio.NewReader(os.Stdin)
		for {
			text, err := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			if err != io.EOF && len(text) > 0 {
				line <- text
			}
		}

	}()
	return line
}
