package main

import (
	"ascii-art/functions"
	"fmt"
	"os"
	"strings"
)

func main() {
	var text string
	var style string
	if len(os.Args) != 3 {
		if len(os.Args) == 2 {
			text = os.Args[1]
			style = "standard"
		} else {
			fmt.Println("Usage: go run . [STRING] [BANNER]")
			return
		}
		
	} else {
		text = os.Args[1]
		if !functions.ValidateASCII(text) {
			fmt.Println("Error: Non-ASCII character detected")
			return
		}
		if strings.ToLower(os.Args[2]) == "shadow" || strings.ToLower(os.Args[2]) == "standard" || strings.ToLower(os.Args[2]) == "thinkertoy" {
			style = strings.ToLower(os.Args[2])
		} else {
			fmt.Println("Invalid style. Please use either shadow, standard, or thinkertoy")
			return
		}
	}
	fileLines := functions.Read(style)
	asciiRep := functions.AsciiRep(fileLines)
	functions.PrintStr(text, asciiRep)
}
