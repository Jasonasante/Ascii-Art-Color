package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Color string

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

func colorise(color Color, s string) {
	fmt.Print(string(color), s, string(ColorReset))
}

func main() {
	var emptyString string
	var inputString []string
	if len(os.Args) == 4 {
		inputString = strings.Split(os.Args[1], "\\n")
		// this takes the argument that we are printing and seperates them into a []string via \n
		// this will then therefore automatically will print each []string on a new line.
	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		os.Exit(0)
	}
	// fmt.Println(inputString)
	Content, _ := os.ReadFile(os.Args[2] + ".txt")
	// fmt.Println(Content)
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
	// for _,strarr := range asciiSlice2{
	// 	for _,str := range strarr{
	// 		fmt.Println(str)
	// 	}
	// }

	colorStr := "Color"
	colorFlag := flag.Bool("color", false, "to colorise string")
	flag.Parse()
	for i := range os.Args[3] {
		if os.Args[3][i] == 61 {
			for j := i + 1; j < len(os.Args[3]); j++ {
				colorStr += string(os.Args[3][j])
			}
		}
	}

	var tempOutput [][]string
	// why is it that when we used make, it did not print the first index?
	for _, str := range inputString {
		for _, aRune := range str {
			tempOutput = append(tempOutput, asciiSlice2[aRune-rune(32)])
			// due to the loop it will append the bubble eqivalent of the every letter inside inputString
		}
		for i := range tempOutput[0] {
			// why does it have to be 0???
			// with tempout[4] the bro disappears but hello and there are printed.
			// tempOutput[0] is the first slice of the 2D array ( which ssshould be the slice of
			// bubble letters that make up inputString)
			// i is each line inside the []string.
			// so for the range of slice of tempout (which is the bublble version of inputString)
			for _, char := range tempOutput {
				if *colorFlag {
					colorise(colorStr, char[i])
				}
				fmt.Print(char[i])
				// this prints each line of each bubble letter
			}
			fmt.Println()
		}
		tempOutput = nil
		// once the word has been printed, we want to reset tempOutput to nil, ready to be filled
		// by the next string element in inputString.
	}
}
