// Package greetings to learn module
package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"unicode"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	return fmt.Sprintf(randFormat(), name), nil
}

func Hellos(names ...string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {

		msg, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[capitalizeFirst(name)] = msg
	}
	return messages, nil
}

func randFormat() string {
	formats := []string{
		"Hi, %s. Welcome!",
		"Great to see you, %s!",
		"Hail, %s!  Well met!",
		"Welcome to %s!",
	}
	idx := rand.Intn(len(formats))
	return formats[idx]
}

func capitalizeFirst(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s) // handle Unicode properly
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
