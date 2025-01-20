package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	readNumber := func(prompt string) int {

		fmt.Println(prompt)
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		num, _ := strconv.Atoi(input)

		return num
	}

	a := readNumber("Enter the first number")

	b := readNumber("Enter the second number")

	c := a + b

	fmt.Printf("%d\n", c)

}
