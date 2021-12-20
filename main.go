package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorWhite  = "\033[37m"
	colorOrange = "\033[40m"
)

// THIS IS TO LOWERCASE THE COLOR WANTED
func ToLower(s string) string {
	lowstr := []rune(s)
	for i, char := range lowstr {
		if char >= 65 && char <= 90 {
			lowstr[i] = lowstr[i] + 32
		}
	}
	return string(lowstr)
}

// THIS IS TO EXTRACT THE POSITION OF THE LETTER THAT WANTS TO BE COLORED
func TrimAtoi(s string) int {
	myRunes := []rune(s)
	num := 0
	for _, b := range myRunes {
		if b >= 48 && b <= 57 {
			num = num*10 + int(b-'0')
		}
	}
	return num
}

// THIS PUT CONVERTS A STRING INTO ASCII-BUBBLES WITH COLORS
func main() {
	var emptyString string
	var inputString []string
	if len(os.Args) == 3 || len(os.Args) == 5 {
		inputString = strings.Split(os.Args[1], "\\n")
		// this takes the argument that we are printing and seperates them into a []string via \n
		// this will then therefore automatically will print each []string on a new line.
	} else {
		fmt.Println("Usage: go run . [STRING] [OPTION]")
		fmt.Println("EX: go run . something --color=<color>")
		os.Exit(0)
	}
	for _, ele := range inputString {
		if ele == "" || ele == "\\n" {
			os.Exit(0)
		}
	}
	Content, _ := os.ReadFile("standard.txt")
	asciiSlice2 := make([][]string, 95)
	// this stores the ascii-bubbles in order of the
	// there are 95 ascii characters and this lets us index the dimension holding each bubble
	for i := 0; i < len(asciiSlice2); i++ {
		asciiSlice2[i] = make([]string, 9)
	}
	// this makes the asciiSlice2[i] have a length of 8
	var bubbleCount int
	count := 0
	for i := 1; i < len(Content); i++ {
		if Content[i] == '\n' && bubbleCount <= 94 {
			asciiSlice2[bubbleCount][count] = emptyString
			// so bubbleCount is the index and count is the row
			// so asciiSlice2[1][0] is the 1st row of the exclamation mark
			emptyString = ""
			count++
			// we want count to get to 8 as that is the number of rows (height of the 8)
		}
		if count == 9 {
			count = 0
			bubbleCount++
			// i++
			// once we have the 8 rows of the bubble text, we want to move onto the next index of the
			// asciiSlice2, hence bubbleCount++
			// We also have i++
		} else {
			if Content[i] != '\n' && Content[i] != '\r' {
				emptyString += string(Content[i])
				// as count != 8 and Contet[i]!= '\n', it will append the contents of each row.
				// Once it reaches the '\n' at the end of the row, the first if statement is activated.
			}
		}
	}

	var colorFlag []string
	if strings.HasPrefix(os.Args[2], "--color=") {
		colorFlag = strings.Split(os.Args[2], "--color=")
	} else {
		fmt.Println("Usage: go run . [STRING] [OPTION]")
		fmt.Println("EX: go run . something --color=<color>")
		os.Exit(0)
	}

	colorFlag[1] = ToLower(colorFlag[1])
	Paint := colorWhite
	if colorFlag[1] == "red" {
		Paint = colorRed
	}
	if colorFlag[1] == "green" {
		Paint = colorGreen
	}
	if colorFlag[1] == "yellow" {
		Paint = colorYellow
	}
	if colorFlag[1] == "blue" {
		Paint = colorBlue
	}
	if colorFlag[1] == "purple" {
		Paint = colorPurple
	}
	if colorFlag[1] == "orange" {
		Paint = colorOrange
	}

	colorCount := 0
	var tempOutput [][]string
	for j, str := range inputString {
		for _, aRune := range str {
			tempOutput = append(tempOutput, asciiSlice2[aRune-rune(32)])
			// due to the loop it will append the bubble eqivalent of the every letter inside inputString
		}
		for i := range tempOutput[0] {
			for _, char := range tempOutput {
				if len(os.Args) == 3 {
					fmt.Print(string(Paint), (char[i]))
				}
				if len(os.Args) == 5 {
					min := TrimAtoi(os.Args[3])
					max := TrimAtoi(os.Args[4])
					if max > min {
						if colorCount >= min-1 && colorCount <= max-1 {
							fmt.Print(string(Paint), (char[i]))
						} else {
							fmt.Print(string(colorWhite), (char[i]))
						}
						colorCount++
						if colorCount == len(inputString[j]) {
							colorCount = 0
						}
					}
					if min > max || min == max {
						if colorCount == min-1 {
							fmt.Print(string(Paint), (char[i]))
						} else {
							fmt.Print(string(colorWhite), (char[i]))
						}
						colorCount++
						if colorCount == len(inputString[j]) {
							colorCount = 0
						}
					}
				}
			}
			fmt.Println()
		}
		tempOutput = nil
		// once the word has been printed, we want to reset tempOutput to nil, ready to be filled
		// by the next string element in inputString.
		// to run this file:
		// to print all letters in colour: go run main.go [string] --color=<color>
		// to print a set of letters in colour: go run main.go [string] --color=<color> start==int end==int
		// to print one letter either: go run main.go [string] --color=<color> start==int end==0
		// or go run main.go [string] --color=<color> start==int end==int --> where start and end integers are equal
	}
}
