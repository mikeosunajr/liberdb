package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Println("expected 'foo' or 'bar' subcommands")
	os.Exit(1)
}

func main() {

	if len(os.Args) < 2 {
		usage()
	}

	switch os.Args[1] {
	case "node":
		nodeCmd()
	default:
		usage()
	}
}
