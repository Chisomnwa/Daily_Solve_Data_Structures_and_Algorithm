package main

import (
	"fmt"
	gosolutions "piscine/go-solutions"
)

func main() {
	// fmt.Println(gosolutions.IsPalindrome("A man, a plan, a canal: Panama"))
	// fmt.Println(gosolutions.IsPalindrome(" "))
	// fmt.Println(gosolutions.IsPalindrome("race a car"))
	// fmt.Println(gosolutions.RomanToInt("IV"))
	// fmt.Println(gosolutions.DefangIPaddr("1.1.1.1"))
	// fmt.Println(gosolutions.IsPalindromeII("aba"))
	// fmt.Println(gosolutions.IsPalindromeII("abca"))
	// fmt.Println(gosolutions.IsPalindromeII("abc"))
	fmt.Println(gosolutions.Interpret("G()(al)"))
	fmt.Println(gosolutions.Interpret("G()()()()(al)"))
	fmt.Println(gosolutions.Interpret("(al)G(al)()()G"))
}
