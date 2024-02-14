package main

import (
    "fmt"
    "time"
    "regexp"
    "strings"
)

func isPalindrome(sentence string)string{
    preprocessed := strings.ToLower(regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(sentence,""))
    left,right := 0 , len(preprocessed)-1
    wrongCharacter:=0

    for left < right{
        if preprocessed[left] != preprocessed[right]{
            wrongCharacter++
        }
        left++
        right--
    }

    
    if wrongCharacter == 0 {
        return "It s already a palindrome"
    } else if wrongCharacter == 1 {
        return "It can be a palindrome by changing only 1 character"
    } else {
        return "It cannot be a palindrome"
    }

}

func main() {

    var input string

	fmt.Print("Enter a sentence :")
	fmt.Scan(&input)

	startTime := time.Now()
	checkPalindrome := isPalindrome(input)

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

    fmt.Printf("%s\n", checkPalindrome)
	fmt.Printf("Elapsed Time: %v\n", elapsedTime)
}
