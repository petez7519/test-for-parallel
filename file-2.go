package main

import "fmt"

// Add adds two integers and returns the result.
func Add(a, b int) int {
    return a + b
}

func main() {
    fmt.Println("Addition:", Add(2, 3))
}
