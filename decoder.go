package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Decoder(inputStr string) (string, error) {
	findNumbers := regexp.MustCompile(`\[(\d+) (.*?)\]`)
	matches := findNumbers.FindAllStringSubmatch(inputStr, -1)
	for _, match := range matches {
		number, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Println("Error conversion :", err)
			os.Exit(1)
		}
		symbol := match[2]
		if symbol == "" {
			return "", fmt.Errorf("\nArt decoder was not successful due to empty string.\n")
		}
		numSmbols := strings.Repeat(symbol, number)
		inputStr = strings.Replace(inputStr, match[0], numSmbols, -1)
	}
	if strings.Contains(inputStr, "[") || strings.Contains(inputStr, "]") {
		return "", fmt.Errorf("\nArt decoder was not successful due to invalid format.\n")
	}
	fmt.Println("\n==========RESULT==========\n", inputStr)
	return inputStr, nil
}

func Encoder(input string) (string, error) {
	var result []string
	var currentChar rune
	var count int
	for i, char := range input {
		if i == 0 || char == currentChar {
			count++
		} else {
			countStr := strconv.Itoa(count)
			charStr := string(currentChar)
			if countStr == "1" {
				result = append(result, string(charStr))
			} else {
				group := "[" + countStr + " " + string(charStr) + "]"
				result = append(result, group)

				count = 1 //reset count for next symbol
			}
		}
		currentChar = char
	}
	if count > 0 { //add fix for last symbol, required if programm runs in single-line mode
		if count == 1 {
			result = append(result, string(currentChar))
		} else {
			group := "[" + strconv.Itoa(count) + " " + string(currentChar) + "]"
			result = append(result, group)
		}
	}
	finalResult := strings.Join(result, "")
	fmt.Println("\n==========RESULT==========\n", finalResult)
	return finalResult, nil
}

func main() {
	GetData()
}

func GetData() {
	argslimit := len(os.Args)
	if argslimit > 3 {
		fmt.Println("\nTo run the program, use flags. For more help, use -h flag.\n")
		os.Exit(0)
	}
	helpFlag := flag.Bool("h", false, "Help")
	decodeFlag := flag.Bool("d", false, "Decode art")
	encodeFlag := flag.Bool("e", false, "Encode art")
	multiline := flag.Bool("multi", false, "Multi-Line mode")
	flag.Parse()
	if *helpFlag {
		fmt.Println("\nTo use DECODER, use falg -d. For example: go run . -d \"[1 #][2 -_]-[3 #]\"")
		fmt.Println("\nTo use ENCODER, use flag -e. For example: go run . -e AAABBC")
		fmt.Println("\nFor MULTI-LINE, add flag -multi and press enter. After that you should type in your input. Once you have finished press 'Ctrl+d'." +
			"\nFor example:  go run . -d -multi" +
			"\n[1 _]/[1 ?][1 _]" +
			"\n[1 _]/[1 ?][1 _]")
		os.Exit(0)
	}

	scanner := bufio.NewScanner(os.Stdin)
	var inputStr string
	if *multiline {
		for scanner.Scan() {
			inputStr += scanner.Text() + "\n"
		}
	} else {
		if len(os.Args) > 2 {
			inputStr = os.Args[2]
		} else {
			fmt.Println("\nNo input provided. For help, type in go run . -h.")
			os.Exit(1)
		}
	}

	if *decodeFlag {
		_, err := Decoder(inputStr)
		if err != nil {
			fmt.Println("Error decoding:", err)
		}
	}
	// if used -e flag, will call func DecodeArt
	if *encodeFlag {
		_, err := Encoder(inputStr)
		if err != nil {
			fmt.Println("Error encoding:", err)
		}
	}
}
