package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isValidWord(word string, availableLetters string) bool {
	// Checks if the given word can be formed using the available letters.
	wordLetters := strings.Split(word, "")
	for _, letter := range wordLetters {
		if strings.Count(word, letter) > strings.Count(availableLetters, letter) {
			return false
		}
	}
	return true
}

func findPossibleWords(availableLetters string, dictionaryFile string) []string {
	// Finds all possible words from the dictionary that can be formed using the available letters.
	var possibleWords []string

	file, err := os.Open(dictionaryFile)
	if err != nil {
		fmt.Println("Error opening dictionary file:", err)
		return possibleWords
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if len(word) > 0 && isValidWord(word, availableLetters) && len(word) == len(availableLetters) {
			possibleWords = append(possibleWords, word)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading dictionary file:", err)
	}

	return possibleWords
}

func main() {
	for {
		fmt.Print("Enter the available letters (or 'exit' to quit): ")
		var availableLetters string
		fmt.Scanln(&availableLetters)
		availableLetters = strings.ToLower(availableLetters)

		if availableLetters == "exit" {
			break
		}

		possibleWords := findPossibleWords(availableLetters, "russian.txt")

		if len(possibleWords) > 0 {
			fmt.Println("Possible words:")
			for _, word := range possibleWords {
				fmt.Println(word)
			}
		} else {
			fmt.Println("No possible words matching the given letters.")
		}
	}
}
