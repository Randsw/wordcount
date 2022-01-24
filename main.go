package main

import (
    "os"
    "fmt"
	"strings"
)

func main() {
	// the first argument is always program name
	argString := os.Args[1]
	words := strings.Fields(argString)
	fmt.Println(len(words))
}