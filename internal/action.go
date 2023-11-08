package internal

func Actions(runeAscii []rune, j int) []int {
	mas := []int{}
	a := 0

	for _, i := range runeAscii {
		a = (int(i)-32)*9 + j
		mas = append(mas, a)

	}
	return mas
}

func Print(textSplit []string, words string) string {
	runeWords := []rune(words)

	result := ""
	for j := 1; j <= 8; j++ {
		a := Actions(runeWords, j)
		for j := 0; j < len(a); j++ {
			result += textSplit[a[j]]
		}
		result += "\n"

	}
	return result
}
