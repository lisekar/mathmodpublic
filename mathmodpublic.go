package mathmodpublic

import "fmt"

func Greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
	// return fmt.Sprintf("Greeting sent to %s", name)
}

func Add(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
