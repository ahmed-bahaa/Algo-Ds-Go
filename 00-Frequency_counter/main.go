package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(validAnagram("hello", "ollhe"))

	fmt.Println(" \"\" \"\" : ", validAnagram("", ""))
	fmt.Println(" \"aaz\" \"zza\" : ", validAnagram("aaz", "zza"))
	fmt.Println(" \"anagram\" \"nagaram\" : ", validAnagram("anagram", "nagaram"))
	fmt.Println(" \"rat\" \"car\" : ", validAnagram("rat", "car"))
	fmt.Println(" \"awesome\" \"awesome\" : ", validAnagram("awesome", "awesome"))
	fmt.Println(" \"amanaplanacanalpanama\" \"acanalmanplanpamana\" : ", validAnagram("amanaplanacanalpanama", "acanalmanplanpamana"))
	fmt.Println(" \"qwerty\" \"qeywrt\" : ", validAnagram("qwerty", "qeywrt"))
	fmt.Println(" \"texttwisttime\" \"timetwisttext\" : ", validAnagram("texttwisttime", "timetwisttext"))

}

func validAnagram(str1 string, str2 string) bool {
	wordMap := make(map[string]int)
	wordMap2 := make(map[string]int)
	if len(str1) != len(str2) || len(str2) == 0 || len(str1) == 0 {
		return false
	}

	for _, c1 := range str1 {
		if validChar(c1) {
			if val, exist := wordMap[strings.ToLower(string(c1))]; exist {
				wordMap[strings.ToLower(string(c1))] = (val + 1)
			} else {
				wordMap[strings.ToLower(string(c1))] = 1
			}
		}
	}
	for _, c2 := range str2 {
		if validChar(c2) {
			if val, exist := wordMap2[strings.ToLower(string(c2))]; exist {
				wordMap2[strings.ToLower(string(c2))] = (val + 1)
			} else {
				wordMap2[strings.ToLower(string(c2))] = 1
			}
		}
	}
	for i := range wordMap {
		if wordMap[i] != wordMap2[i] {
			return false
		}
	}
	return true
}

func validChar(ch rune) bool {
	if ch >= 97 && ch <= 122 {
		return true
	}
	if ch >= 065 && ch <= 90 {
		return true
	}
	return false
}
