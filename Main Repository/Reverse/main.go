package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	reverseFlag := flag.String("reverse", "", "File name to reverse")
	flag.Parse()

	// Check if the --reverse flag is provided and in the correct format
	if *reverseFlag == "" || !strings.HasPrefix(*reverseFlag, "--reverse=") {
		printUsage()
		return
	}

	// Extract the file name from the provided argument
	fileName := strings.TrimPrefix(*reverseFlag, "--reverse=")

	// Read the content of the file
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	// Reverse the content
	reversedContent := reverseString(string(content))

	// Print the reversed content
	fmt.Println(reversedContent)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func printUsage() {
	fmt.Println("Usage: go run . --reverse=<fileName>")
}
