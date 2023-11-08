package internal

import (
	"errors"
	"io/ioutil"
	"strings"
)

var (
	BadRequest = errors.New("Bad Request")
)

func Args1(banner, words string) (string, error) {
	for _, r := range words {
		if r > 127 {
			BadRequest := errors.New("400: Bad Request")
			return "", BadRequest
		}
	}
	var textSplit []string
	data := banner
	if data == "standard" {
		textSplit1, err := CheckData("data/standard.txt")
		textSplit = textSplit1
		if err != nil {
			return "", BadRequest
		}
	} else if data == "shadow" {
		textSplit1, err := CheckData("data/shadow.txt")
		textSplit = textSplit1
		if err != nil {
			return "", BadRequest
		}
	} else if data == "thinkertoy" {
		textSplit1, err := CheckData("data/thinkertoy.txt")
		textSplit = textSplit1
		if err != nil {
			return "", BadRequest
		}

	} else {
		return "", BadRequest
	}

	words = strings.ReplaceAll(words, "\r\n", "\n")
	splitWords := strings.Split(words, "\n")
	count := 0
	countProbel := 0
	result := ""
	for _, i := range splitWords {
		if i == "" {
			countProbel++
		}
	}
	if countProbel == len(splitWords) {
		for i := 1; i < countProbel; i++ {
			result += "\n"
		}
		return result, nil
	}
	for _, i := range splitWords {
		if len(i) > 0 {
			count++
		}

		if count == 0 {
			result += "\n"
			continue
		}
		count = 0
		result += Print(textSplit, i)
	}
	return result, nil
}

func CheckData(a string) ([]string, error) {
	text, err := ioutil.ReadFile(a)
	text1 := strings.ReplaceAll(string(text), "\r", "")
	textSplit := strings.Split(text1, "\n")

	return textSplit, err
}
