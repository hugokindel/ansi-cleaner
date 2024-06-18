package main

import (
	"github.com/hugokindel/ansi-cleaner/cmd"
	"log"
	_ "net/http/pprof"
)

func main() {
	executeCmd()
}

func executeCmd() {
	if err := cmd.Cmd.Execute(); err != nil {
		log.Fatalf("Whoops. There was a fatal error: %s", err)
	}
}
