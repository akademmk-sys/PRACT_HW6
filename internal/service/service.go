package service

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(str string) (string, error) {
	if str == "" {
		err := errors.New("empty message")
		return "", err
	}
	isNotMorse := func(r rune) bool {
		return r != '.' && r != '-' && r != ' '
	}
	if strings.ContainsFunc(str, isNotMorse) {
		tempStr := strings.ReplaceAll(str, " ", "")
		for _, elem := range tempStr {
			upperElem := unicode.ToUpper(elem)
			if _, ok := morse.DefaultMorse[upperElem]; !ok {
				err := fmt.Errorf("wrong string format:%s", str)
				return "", err
			}
		}
		msg := morse.ToMorse(str)
		return msg, nil
	}
	msg := morse.ToText(str)
	return msg, nil
}
