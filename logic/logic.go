package logic

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func stringToASCII(s string) []int {
	asciiValues := make([]int, len(s))

	for i, char := range s {
		asciiValues[i] = int(char)
	}

	return asciiValues
}

func splitArray(arr []int, splitNum int) [][]int {
	var result [][]int
	currentSection := []int{}

	for _, num := range arr {
		if num == splitNum {
			result = append(result, currentSection)
			currentSection = []int{}
		} else {
			currentSection = append(currentSection, num)
		}
	}

	// Append the last section
	result = append(result, currentSection)

	return result
}

func LogicAscii(text string, style string) string{

	// Access the arg
	inputString := text

	// Replace the escape sequence "\n" with an actual newline character
	inputString = strings.ReplaceAll(inputString, "\\n", "\n")

	inputString = strings.Trim(inputString, "")

	filename := style + ".txt"

	// Open the file
	file, err := os.Open("./" + filename)

	if err != nil {
		fmt.Println("Error:", err)
		return "habibi madri"
	}
	defer file.Close()

	// turn the string to an array of ascii representation
	arrayOfCharacters := stringToASCII(inputString)

	// split the array based on the 10
	splittedArrayBasedOn10 := splitArray(arrayOfCharacters, 10)

	// Create a new scanner for the file
	scanner := bufio.NewScanner(file)

	// Slice to store lines
	var linesFromFile []string

	// Iterate over each line and add it to the array
	for scanner.Scan() {
		lineFromFile := scanner.Text()
		linesFromFile = append(linesFromFile, lineFromFile)
	}

	bigAssString := ""

	// for every array in the array [[65 108 105] [104 101 108 108 111]] 0 & 1
	for j := 0; j < len(splittedArrayBasedOn10); j++ {

		// A 9 time for loop since every character has 8 lines and a new line
		for i := 1; i <= 9; i++ {

			// for every character in the array, (this will repeat 9 times due to outer loop)
			for k, asciiRep := range splittedArrayBasedOn10[j] {

				// read the asciiRep to get the position of the pointer for the lines array
				positionOfpointer := (asciiRep-32)*9 + i

				// print out every line without a new line
				bigAssString += linesFromFile[positionOfpointer]

				// if we reach the end of an array, print a new line
				if k == len(splittedArrayBasedOn10[j])-1 {
					bigAssString += "\n"
				}
			}
		}
		// if a specific array in the bigger array is empty, output a new line
		if len(splittedArrayBasedOn10[j]) == 0 {
			bigAssString += "\n"
		}

	}


	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return bigAssString

}
