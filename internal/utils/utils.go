package internal

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type EventRequest struct {
	EventName string `json:"eventName"`
}

// LoadEnv accepts a variable number of keys and returns the corresponding value from the env file
func LoadEnv(keys ...string) ([]string, error) {
	// err := godotenv.Load("../../cmd/app/.env")
	err := godotenv.Load("cmd/app/.env")
	if err != nil {
		return nil, err
	}

	var values []string

	for _, key := range keys {
		value := os.Getenv(key)
		if value == "" {
			return nil, fmt.Errorf("%s does not exist in the env file", key)
		}
		values = append(values, value)
	}

	return values, nil
}

// ReverseString reverses a given imput
func ReverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	reversed := string(runes)
	return reversed
}

// Logger creates||opens a log txt file and sets log outputs to be the created||opened file. returns *os.File or an error
func Logger(logFile string) (*os.File, error) {
	// creates a txt file for basic logging if it does not exist
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("error: could not open log file: %s", err)
	}

	return file, nil
}

// isFileEmpty checks if a file is empty by checking the size of the file
func IsFileEmpty(filename string) (bool, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return false, err
	}

	if fileInfo.Size() == 0 {
		return true, nil
	}

	return false, nil
}
