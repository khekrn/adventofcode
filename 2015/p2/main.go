package main

/*

--- Day 2: I Was Told There Would Be No Math ---
The elves are running low on wrapping paper, and so they need to submit an order for more. They have a list of the dimensions (length l, width w, and height h) of each present, and only want to order exactly as much as they need.

Fortunately, every present is a box (a perfect right rectangular prism), which makes calculating the required wrapping paper for each gift a little easier: find the surface area of the box, which is 2*l*w + 2*w*h + 2*h*l. The elves also need a little extra paper for each present: the area of the smallest side.

For example:

A present with dimensions 2x3x4 requires 2*6 + 2*12 + 2*8 = 52 square feet of wrapping paper plus 6 square feet of slack, for a total of 58 square feet.
A present with dimensions 1x1x10 requires 2*1 + 2*10 + 2*10 = 42 square feet of wrapping paper plus 1 square foot of slack, for a total of 43 square feet.
All numbers in the elves' list are in feet. How many total square feet of wrapping paper should they order?

*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	res := surfaceArea()
	fmt.Println(res)
	res = ribbon()
	fmt.Println(res)
}

func ribbon() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error while opening file - ", err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	res := 0
	for scanner.Scan() {
		word := scanner.Text()
		wordSplit := strings.Split(word, "x")
		l, _ := strconv.Atoi(wordSplit[0])
		w, _ := strconv.Atoi(wordSplit[1])
		h, _ := strconv.Atoi(wordSplit[2])

		lw := (2 * l) + (2 * w)
		wh := (2 * h) + (2 * w)
		hl := (2 * l) + (2 * h)
		wrap := min(lw, wh, hl)
		bow := l * w * h
		res += (bow + wrap)
	}

	return res
}

func surfaceArea() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error while opening file - ", err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	res := 0
	for scanner.Scan() {
		word := scanner.Text()
		wordSplit := strings.Split(word, "x")
		l, _ := strconv.Atoi(wordSplit[0])
		w, _ := strconv.Atoi(wordSplit[1])
		h, _ := strconv.Atoi(wordSplit[2])

		lw := l * w
		wh := w * h
		hl := h * l

		m := min(lw, wh, hl)

		temp := (2 * lw) + (2 * wh) + (2 * hl) + m
		res += temp

	}
	return res
}

func min(x, y, z int) int {
	if x <= y && x <= z {
		return x
	} else if y <= x && y <= z {
		return y
	}
	return z
}
