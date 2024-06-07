// Copyright (c) 2024 Timo Savola
// SPDX-License-Identifier: GPL-2.0-or-later

package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode/utf8"
)

//go:embed words.txt
var wordsTXT string

func main() {
	words := strings.Split(wordsTXT, "\n")
	if last := len(words) - 1; words[last] == "" {
		words = words[:last]
	}

	var letters map[rune]struct{}
	for len(letters) != 7 {
		letters = getLetterSet(words[rand.Intn(len(words))])
	}

	b := bytes.NewBuffer(nil)
	for c := range letters {
		b.WriteRune(c)
	}
	fmt.Println("Kirjaimet:", strings.ToUpper(b.String()))

	r := bufio.NewReader(os.Stdin)
	var found []string

mainloop:
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}

		s = strings.TrimSpace(s[:len(s)-1])

		for _, f := range found {
			if f == s {
				fmt.Println("Sana on jo löytynyt")
				continue
			}
		}

		if utf8.RuneCountInString(s) < 4 {
			fmt.Println("Sana on liian lyhyt")
			continue
		}

		set := make(map[rune]struct{})
		for _, c := range s {
			if _, ok := letters[c]; !ok {
				fmt.Printf("Kirjain ei kuulu joukkoon: %c\n", c)
				continue mainloop
			}
			set[c] = struct{}{}
		}
		if len(set) >= 7 {
			fmt.Println("Sana sisältää kaikki kirjaimet!")
		}

		found = append([]string{s}, found...)
		fmt.Println("Löytyneet sanat:", strings.Join(found, ", "))
	}
}

func getLetters(word string) []rune {
	runes := make([]rune, len(word))
	for _, c := range word {
		runes = append(runes, c)
	}
	return runes
}

func getLetterSet(word string) map[rune]struct{} {
	set := make(map[rune]struct{}, len(word))
	for _, c := range word {
		set[c] = struct{}{}
	}
	return set
}
