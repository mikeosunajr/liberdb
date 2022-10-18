package main

import (
	"flag"
	"fmt"
	"os"
)

var fooCmd *flag.FlagSet
var fooEnable *bool
var fooName *string

func init() {
	fooCmd = flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable = fooCmd.Bool("enable", false, "enable")
	fooName = fooCmd.String("name", "", "name")
}

func nodeCmd() {
	fooCmd.Parse(os.Args[2:])

	fmt.Println("subcommand 'node'")
	fmt.Println("  enable:", *fooEnable)
	fmt.Println("  name:", *fooName)
	fmt.Println("  tail:", fooCmd.Args())
}
