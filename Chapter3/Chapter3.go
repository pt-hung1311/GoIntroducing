package main
import (
	"fmt"
)

func main() {
	var input float64
	fmt.Printf("Your input: ")
	fmt.Scanf("%f", &input)
	output := (input - 32)*5/9
	fmt.Printf("Your output: %f (Celsius)\n", output)
	output1 := input * 0.3048
	fmt.Printf("Your output: %f (meters)\n", output1)
}