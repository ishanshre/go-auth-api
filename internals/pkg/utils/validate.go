package utils

import (
	"fmt"
	"net/mail"
	"regexp"
)

func CheckUsername(word string) error {
	// checks the word that is alpha-numeric or not
	condition := regexp.MustCompile(`^[a-zA-Z0-9]{7,}`).MatchString(word)
	if !condition {
		return fmt.Errorf("error: username must be alphanumeric and more than 6 characters")
	}
	return nil
}

func CheckPassword(word string) error {
	if len(word) <= 7 {
		return fmt.Errorf("password length must be more than 7 character")
	}
	condition := regexp.MustCompile(`^[a-zA-Z0-9]{7,}`).MatchString(word)
	if !condition {
		return fmt.Errorf("error: password must have at least one uppercase, lowercase and number")
	}
	return nil
}

func CheckEmail(word string) error {
	_, err := mail.ParseAddress(word)
	return err
}
