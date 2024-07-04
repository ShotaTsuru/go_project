package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	words := []string{"Go", "语言", "程序设计语言", "Go", "语言", "微笑み", "プログラミング", "タコのあしは8本", "タコのあしは8本でイカの足は10本"}
	for _, word := range words {
		switch size := utf8.RuneCountInString(word); size {
		case 1, 2, 3, 4:
			fmt.Printf("%s の文字数は %d です。\n短い単語だ", word, utf8.RuneCountInString(word))
		case 5, 6, 7, 8:
			fmt.Printf("%s の文字数は %d です。\n中くらいの単語だ", word, utf8.RuneCountInString(word))
		default:
			fmt.Printf("%s の文字数は %d です。\n長い単語だ", word, utf8.RuneCountInString(word))
			if n := len(word); size < n {
				fmt.Printf("%d バイトのデータを使っている\n", n)
			} else {
				fmt.Println()
			}
		}
	}
}
