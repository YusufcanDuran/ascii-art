package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var asciiArt []string
	var group []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		group = append(group, line)
		if len(group) == 9 {
			asciiArt = append(asciiArt, strings.Join(group, "\n"))
			group = []string{}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return asciiArt
}

func main() {	
    if len(os.Args) != 2 {
		fmt.Println("Usage: go run . \"myMessage\"")
		return
	} else if os.Args[1] == "" {
	   return
	} else if os.Args[1] == "\\n" {
        fmt.Println()
        return
    } 

	asciiArt := readFile("ascii.txt")
    input := os.Args[1]

	parts := strings.Split(input, "\\n") // Girdiyi \n karakterinden ayır
	for _, part := range parts {
		if part != "" {
			PrintAsciiArt(part, asciiArt)
		} else {
			fmt.Println()
		}
	}
}

func PrintAsciiArt(input string, ascii []string) {
	for line := 0; line < 8; line++ {
		for _, char := range input {
			asciiCode := int(char)
			if asciiCode >= 32 && asciiCode <= 126 {
				asciiIndex := asciiCode - 32
				if asciiIndex < 0 || asciiIndex >= len(ascii) {
					fmt.Printf("No ASCII art representation found for character '%c'\n", char)
					continue
				}
				lines := strings.Split(ascii[asciiIndex], "\n")
				if line < len(lines) {
					fmt.Print(lines[line])
				}
			} else {
				fmt.Printf("Invalid character: %c\n", char)
			}
		}
		fmt.Println()
	}
}
