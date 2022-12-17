package spinwords

import (
	"strings"
)

/*
Write a function that takes in a string of one or more words,
and returns the same string, but with all five or more letter words reversed
(Just like the name of this Kata).
Strings passed in will consist of only letters and spaces.
Spaces will be included only when more than one word is present.

Examples:

spinWords( "Hey fellow warriors" ) => returns "Hey wollef sroirraw"
spinWords( "This is a test") => returns "This is a test"
spinWords( "This is another test" )=> returns "This is rehtona test"
*/
func SpinWords(str string) string {
	arr := strings.Split(str, " ")

	for i, word := range arr {
		if len(word) >= 5 {
			r := []rune(word)
			for n, j := 0, len(r)-1; n < len(r)/2; n, j = n+1, j-1 {
				r[n], r[j] = r[j], r[n]
			}
			arr[i] = string(r)
		}
	}

	return strings.Join(arr, " ")
}
