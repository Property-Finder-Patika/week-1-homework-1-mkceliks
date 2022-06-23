package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i] // Takes an arg from command line and sets it to a string
		sep = " "
	}

	fmt.Println(s) // printing that string.
}
