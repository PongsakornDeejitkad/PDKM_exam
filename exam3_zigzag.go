package main

import "fmt"
import "time"

func convert(s string, numRows int) string {
  numberOfRows,numberOfCols :=numRows,len(s)
  matrix := make([][]string,numberOfRows)

	// create matrix
  for i := range matrix{
	matrix[i] = make([]string,numberOfCols)
	for j := range matrix[i]{
		matrix[i][j] = " "
	}
  }
	row, col := 0, 0
	moveDown := true
	for _, char := range s {
		matrix[row][col] = string(char)

		if numRows > 1 {
			if row == 0 {
				moveDown = true
			} 
			if row == numRows-1 {
				moveDown = false
			} 
			if moveDown{
				row++
			}else {
				row--
			}
		} 
		col++
	}

	result :=convertToString(matrix)

  return result
}

func convertToString(matrix [][]string) string {
	var result string

	for _, row := range matrix {
		result =""
		for _, char := range row {
			result += char
		}
		fmt.Print("\n")
		fmt.Print(result)
	}

	return result
}

func main() {
	var number int
	var integer string

	fmt.Print("Enter a number of row and sentence : ")
	fmt.Scan(&number)
	fmt.Scan(&integer)

    startTime := time.Now()

	convert(integer,number)

    endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	fmt.Printf("\nElapsed Time: %v\n", elapsedTime)
    
}