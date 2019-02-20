package main

import (
	"flag"
	"fmt"
	"os"
)

func bob() {
	// Parse for flags
	flag.Parse()
	args := flag.Args()
	fmt.Printf("%v\n", args)

	// Check arguments
	if len(args) < 1 {
		fmt.Println("First argument must be the url")
		os.Exit(1)
	}

	//
}
