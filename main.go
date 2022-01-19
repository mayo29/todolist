package main

import (
	"fmt"
	"math"
)

func main() {

}

func multiply(num1 int, num2 int) int {
	return num1 * num2
}

func subtract(num1 int, num2 int) int {
	return num1 + num2
}

func divide(num1 float64, num2 float64) float64 {
	return num1 / num2
}

func powerOf(num1 float64, num2 float64) float64 {
	return math.Pow(num1, num2)
}

func sayHello() {
	fmt.Println("Hello")
}
