package io

import (
	"fmt"
	"os"
)

// This method is not going to be used frequently in favour of
// `go:embed input.txt`
func ReadInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("error: cannot able to read part1.txt - ", err.Error())
		os.Exit(1)
	}
	return string(data)
}
