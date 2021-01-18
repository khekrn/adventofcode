package io

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

// ReadAll - reads all the file contents and returns string
func ReadAll(fileName string) string {

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal("file not found " + fileName)
	}

	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("error while reading file " + err.Error())
	}

	return string(content)
}

// ScanFile provides the scanner object which helps to read the given file
// line by line
func ScanFile(fileName string) *bufio.Scanner {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal("file not found " + fileName)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return scanner
}
