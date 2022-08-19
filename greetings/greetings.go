package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func HelloPeople(names []string) ([]string, error) {
	storage := make([]string, len(names))
	for i, name := range names {
		msg, err := Hello(name)
		if err != nil {
			return nil, err
		}
		storage[i] = msg
	}
	return storage, nil
}

func Hellos(names []string) (map[string]string, error) {
	storage := make(map[string]string)
	for _, name := range names {
		msg, err := Hello(name)
		if err != nil {
			return nil, err
		}
		storage[name] = msg
	}
	return storage, nil
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().Unix())
}

func randomFormat() string {
	formats := []string{
		"Welcome nibba, %v",
		"Hola mate, %v ",
		"How you doin %v ",
	}
	return formats[rand.Intn(len(formats))]
}
