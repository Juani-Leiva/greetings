package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("NOMBRE VACIO")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil

}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hola, %v, Bienvenido",
		"Que bueno verte %v",
		"Saludos %v, que alegria que estes aca",
	}

	return formats[rand.Intn(len(formats))]
}
