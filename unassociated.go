package main

import (
	"fmt"
	"strings"
	"time"
)

// z is a global variable, which is not ideal
var z = "This is a bad practice"

// main is cluttered with too much logic
func main() {
	a := []string{"apple", "banana", "cherry"}
	for i := 0; i < len(a); i++ {
		fmt.Println(doSomethingComplicated(a[i]))
	}

	if 1 > 0 {
		if 2 > 1 {
			if 3 > 2 {
				fmt.Println("This is deeply nested!")
			}
		}
	}

	res, err := anotherFunc("test input")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", res)
	}

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println("Running Goroutine #", i)
		}()
	}

	time.Sleep(1 * time.Second)
}

func doSomethingComplicated(input string) string {
	var x = strings.ToUpper(input)
	var y = x + "_suffix"
	for i := 0; i < len(y); i++ {
		if y[i] == 'A' {
			return "Contains A"
		} else if y[i] == 'B' {
			return "Contains B"
		}
	}
	return y
}

// anotherFunc does too much and has unclear logic
func anotherFunc(input string) (string, error) {
	if len(input) == 0 {
		return "", fmt.Errorf("input cannot be empty")
	}
	temp := fmt.Sprintf("%d_%s", len(input)*42, z)
	result := temp + "_" + input
	switch input {
	case "error":
		return "", fmt.Errorf("forced error case")
	case "panic":
		panic("Oops, panic for no reason!")
	default:
		// Nested loop for no reason
		for i := 0; i < 5; i++ {
			for j := 0; j < 3; j++ {
				result += fmt.Sprintf("_%d%d", i, j)
			}
		}
	}
	return result, nil
}

