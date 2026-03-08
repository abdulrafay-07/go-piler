package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"
	"unicode"
)

type TokenType string

const (
	ID = "ID"
	CONST = "CONST"
	VAR = "VAR"
	TYPE = "TYPE"

	SEMICOLON = "SEMICOLON"
	COLON = "COLON"
	EQUALS = "EQUALS"
	QUICK_ASSIGNMENT = "QUICK_ASSIGNMENT"
)

var keywords = map[string]TokenType{
	"const": CONST,
	"var": VAR,
	"int":    TYPE,
	"float":  TYPE,
	"bool":   TYPE,
	"string": TYPE,
	"rune":   TYPE,
}

var symbols = map[string]TokenType{
	";": SEMICOLON,
	":": COLON,
	"=": EQUALS,
	":=": QUICK_ASSIGNMENT,
}

var wordBreakers = []rune{
	';',
	':',
	'=',
	' ',
}

type Token struct {
	name TokenType
	value string
}

func LookupWords(word string) Token {
	// Check in keywords
	if tok, ok := keywords[word]; ok {
		return Token{
			name: tok,
			value: word,
		}
	}

	// Check in symbols
	if tok, ok := symbols[word]; ok {
		return Token{
			name: tok,
			value: word,
		}
	}

	return Token{
		name: ID,
		value: word,
	}
}

func isWordBreaker(ch rune) bool {
	return slices.Contains(wordBreakers, ch)
}

func isWhitespace(ch rune) bool {
	return unicode.IsSpace(ch)
}

func main() {
	var tokens = []Token{}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	builder := strings.Builder{}
	reader := bufio.NewReader(file)
	for {
		ch, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		if isWordBreaker(ch) {
			if builder.Len() > 0 {
				tokens = append(tokens, LookupWords(builder.String()))
				builder.Reset()
			}

			if isWhitespace(ch) {
				continue
			}
		}

		_, err = builder.WriteRune(ch)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	fmt.Println("Tokens", tokens)
}
