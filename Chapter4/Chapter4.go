package main
import (
	"fmt"
)

func main() {
	for i := 1; i < 100; i++ {
		if i % 3 == 0 {
			fmt.Printf("%d, ", i)
		}
	}
	fmt.Println()

	for i := 1; i < 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Printf("FizzBuzz, ")
		} else if i % 3 == 0 {
			fmt.Printf("Fizz, ")
		} else if i % 5 == 0 {
			fmt.Printf("Buzz, ")
		}
	}
	fmt.Println()
}