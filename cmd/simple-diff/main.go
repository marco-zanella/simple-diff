package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/marco-zanella/simple-diff/diff"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 || len(args) > 2 {
		fmt.Println("Usage: diff-tool <file1> <file2>")
		os.Exit(1)
	}

	path1 := ""
	path2 := ""

	if len(args) == 1 {
		// If only one file is provided, assume the first file is empty
		path2 = args[0]
	} else {
		path1 = args[0]
		path2 = args[1]
	}

	result, err := diff.DiffFiles(path1, path2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	for _, line := range result.Removed {
		fmt.Printf("< %s\n", line)
	}

	for _, line := range result.Added {
		fmt.Printf("> %s\n", line)
	}
}
