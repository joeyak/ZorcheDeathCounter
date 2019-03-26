package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	death := 0
	fmt.Println("Type enter to increment")
	fmt.Println("Type any non number to show output")
	fmt.Println("Type any number to append the amount to the counter")
	fmt.Println("Type :n to set the count to n")
	fmt.Println("If :n has n as a non numeric, it will act like a letter")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var set bool
		s := scanner.Text()
		if strings.HasPrefix(s, ":") {
			s = s[1:]
			set = true
		}

		if s == "" {
			death++
		} else {
			i, err := strconv.Atoi(s)
			if err == nil {
				if set {
					death = i
				} else {
					death += i
				}
			}
		}

		fmt.Printf("Death Count: %d\n", death)
	}

	if scanner.Err() != nil {
		fmt.Printf("Error %v\n", scanner.Err())
	}
}
