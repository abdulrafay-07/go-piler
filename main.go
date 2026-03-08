package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
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
	UNKNOWN = "UNKNOWN"

	SEMICOLON = "SEMICOLON"
	COLON = "COLON"
	EQUALS = "EQUALS"
	QUICK_ASSIGNMENT = "QUICK_ASSIGNMENT"

	INT = "INT"
	FLOAT = "FLOAT"
)

var keywords = map[string]TokenType{
	"const": CONST,
	"var": VAR,
	"int":    TYPE,
	"float":  TYPE,
	"bool":   TYPE,
	"string": TYPE,
	"rune":   TYPE,
	"unknown": UNKNOWN,
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
	fmt.Println("word", word)
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

	return getTokenFromRegex(word)
}

func isWordBreaker(ch rune) bool {
	return slices.Contains(wordBreakers, ch)
}

func isWhitespace(ch rune) bool {
	return unicode.IsSpace(ch)
}

func getTokenFromRegex(word string) Token {
	intRegex := regexp.MustCompile(`^[0-9]+$`)
	floatRegex := regexp.MustCompile(`^[0-9]+\.[0-9]+$`)
	identifierRegex := regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

	switch {
		case intRegex.MatchString(word):
			return Token{
				name: INT,
				value: word,
			}
		case floatRegex.MatchString(word):
			return Token{
				name: FLOAT,
				value: word,
			}
		case identifierRegex.MatchString(word):
			return Token{
				name: ID,
				value: word,
			}
		default:
			return Token{
				name: UNKNOWN,
				value: word,
			}
	}
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

		if isWhitespace(ch) {
			if builder.Len() > 0 {
				tokens = append(tokens, LookupWords(builder.String()))
				builder.Reset()
			}
			continue
		}

		if isWordBreaker(ch) {
			if builder.Len() > 0 {
				tokens = append(tokens, LookupWords(builder.String()))
				builder.Reset()
			}
		}

		_, err = builder.WriteRune(ch)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	if builder.Len() > 0 {
		tokens = append(tokens, LookupWords(builder.String()))
		builder.Reset()
	}

	fmt.Println("Tokens", tokens)
}
