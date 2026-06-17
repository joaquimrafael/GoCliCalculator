package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("invalid operation %s", op)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("-----------------------------------------------")
		fmt.Println("*Go CLI Calculator*")
		fmt.Println("By Joaquim Prieto 2026")

		fmt.Println("Type: value1 value2 operator ->")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			}
			break
		}
		fields := strings.Fields(scanner.Text())

		if len(fields) == 1 && fields[0] == "close" {
			break
		}

		if len(fields) != 3 {
			fmt.Println("Expected 3 values: value1 value2 operator")
			continue
		}

		v1, err1 := strconv.ParseFloat(fields[0], 64)
		v2, err2 := strconv.ParseFloat(fields[1], 64)
		if err1 != nil || err2 != nil {
			fmt.Println("values must be numbers")
			continue
		}

		op := fields[2]

		result, err := calculate(v1, v2, op)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Printf("Result: %.2f\n", result)

		fmt.Println("-----------------------------------------------")
	}
}
