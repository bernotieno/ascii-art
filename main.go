package main

import (
	"flag"
	"fmt"
	"log"

	"ascii/ascii"
)

// formats is a map that maps output format flags to their corresponding filenames.
var formats = map[string]string{
	"s":  "standard.txt",
	"t":  "thinkertoy.txt",
	"sh": "shadow.txt",
}

func main() {
	formatPtr := flag.String("f", "s", "Output format: s (standard), t (thinkertoy), sh (shadow)")
	flag.Parse()

	filename, ok := formats[*formatPtr]
	if !ok {
		fmt.Println("Invalid format specified.")
		return
	}

	asciMap, err := ascii.ReadASCIIMapFromFile(filename)
	if err != nil {
		log.Fatalf("Error reading ASCII map: %v", err)
	}

	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Please provide only one input string.")
		return
	}

	input := args[0]

	ascii.PrintArt(input, asciMap)
}
