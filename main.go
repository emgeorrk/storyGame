package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func reverseString(input string) string {
	runes := []rune(input)

	length := len(runes)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func extractNumberFromEnd(s string) int {
	i := len(s) - 1
	numberStr := ""

	for i >= 0 {
		if unicode.IsDigit(rune(s[i])) {
			numberStr += string(s[i])
		} else {
			break
		}
		i--
	}

	numberStr = reverseString(numberStr)

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		return -1
	}

	return number
}

func main() {
	path := "/home/egormerk/GolandProjects/story-game/content/"
	file := "1.txt"
	for {
		filePath := path + file
		content, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatal(err)
		}
		lines := strings.Split(string(content), "\n")

		nextFile := -1
		options := make(map[int]int)
		for _, line := range lines {
			nextFile = extractNumberFromEnd(line)
			if nextFile == -1 {
				fmt.Println(line)
			} else {
				digit, err := strconv.Atoi(string(line[0]))
				if err != nil {
					fmt.Println("Ошибка при преобразовании:", err)
					return
				}

				options[digit] = nextFile

				j := 0
				for j < len(line) {
					if j+1 < len(line) && line[j] == ' ' && unicode.IsDigit(rune(line[j+1])) {
						break
					}
					j++
				}
				fmt.Println(line[0:j])
			}
		}

		if nextFile == -1 {
			break
		}

		fmt.Print("Выберите действие: ")
		var input int
		_, err = fmt.Scan(&input)
		if err != nil {
			fmt.Println("Ошибка ввода")
			return
		}
		file = fmt.Sprintf("%d.txt", options[input])
	}
}
