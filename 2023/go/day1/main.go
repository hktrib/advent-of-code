package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

var lineLen int = 4096

func main() {

	file, err := os.Open("../../input/day34.txt")
	if err != nil {
		log.Fatalf("Error opening input file: %v", err)
	}
	defer file.Close()

	// Setting scanner.Buffer
	scanner := bufio.NewScanner(file)
	buffer := make([]byte, lineLen)
	scanner.Buffer(buffer, lineLen)

	calibrationsSum := 0

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		seenFirst := false

		var snum1 string
		var snum2 string

		for _, r := range scanner.Text() {
			if unicode.IsNumber(r) {
				if !seenFirst {
					snum1 = string(r)
					seenFirst = true
				} else {
					snum2 = string(r)
				}
			}

		}

		var snum string

		if snum1 == "" {
			// No calibrations in this line
			continue
		} else if snum2 == "" {
			// No second calibration value -> duplicating calibration val 1
			snum = snum1 + snum1
		} else {
			snum = snum1 + snum2
		}

		calibrationVal, err := strconv.Atoi(snum)
		if err != nil {
			log.Fatalf("Unable to convert snum %v to int: %v", snum, err)
		}
		calibrationsSum += calibrationVal
	}

	fmt.Printf("Total Calibrations: %d\n", calibrationsSum)
}
