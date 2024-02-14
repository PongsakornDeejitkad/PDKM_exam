package main

import "fmt"
import "time"


func flipBits(number int) int {
	numBits := 0
	for temp := number; temp > 0; temp /= 2 {
		numBits++
	}
	flippedNumber := number ^ ((1 << numBits)-1)

	return flippedNumber
}

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

    startTime := time.Now()

	binaryNumber := flipBits(number)

    endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

    fmt.Printf("Result: %d\n", binaryNumber)
	fmt.Printf("Elapsed Time: %v\n", elapsedTime)
    
}