package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestMainFunction(t *testing.T) {
	testCases := []struct {
		input string
	}{
		{
			input: "Hello\n",
		},
		{
			input: "hello",
		},
		{
			input: "HeLlO",
		},
		{
			input: "HelloThere",
		},
		{
			input: "1Hello2There",
		},
		{
			input: "{HelloThere}",
		},
		{
			input: "HelloThere",
		},
		{
			input: "HelloThere",
		},
	}

    // Create or clear output.txt
	err := os.WriteFile("output.txt", []byte{}, 0644)
	if err != nil {
		t.Fatalf("Error creating or clearing output.txt: %v", err)
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			// Backup original os.Args
			oldArgs := os.Args
			// Restore os.Args after the test finishes
			defer func() { os.Args = oldArgs }()

			// Run the main function with piping to cat -e and write output to output.txt
			cmd := exec.Command("sh", "-c", fmt.Sprintf("go run main.go %s | cat -e | tee -a output.txt", tc.input)) 
			err := cmd.Run()
			if err != nil {
				t.Fatalf("Error running main function: %v", err)
          fmt.Println(err)
			}
            
			// Open output.txt for reading
			outputFile, err := os.Open("output.txt")
			if err != nil {
				t.Fatalf("Error opening output.txt: %v", err)
			}
			defer outputFile.Close()

			// Open is_output_correct.txt for reading
			correctOutputFile, err := os.Open("is_output_correct.txt")
			if err != nil {
				t.Fatalf("Error opening is_output_correct.txt: %v", err)
			}
			defer correctOutputFile.Close()

			// Compare the content of output.txt with is_output_correct.txt
			outputScanner := bufio.NewScanner(outputFile)
			correctOutputScanner := bufio.NewScanner(correctOutputFile)

			for outputScanner.Scan() && correctOutputScanner.Scan() {
				if outputScanner.Text() != correctOutputScanner.Text() {
					t.Errorf("MainFunction(%q) produced different output from expected", tc.input)
					return
				}
			}

			// Check if one file is longer than the other
			if outputScanner.Scan() || correctOutputScanner.Scan() {
				t.Errorf("MainFunction(%q) produced different output from expected", tc.input)
				return
			}
		})
	}
}
