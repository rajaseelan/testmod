package testmod

import (
	"fmt"
	"errors"
)

// Hi Returns friendly greeting in a language lang
func Hi(name, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hi. %s!", name), nil
	case "pt":
		return fmt.Sprintg("Oi, %s!", name), nil
	default:
		return "", errors.New("Unknown language")
	}
}


