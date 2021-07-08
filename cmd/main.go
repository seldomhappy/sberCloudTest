package main

import (
	"fmt"
	shortString "sberCloudTest/pkg"
)

func main() {
	str := "aaabbccccdeeeff"
	res := shortString.Get(str)
	fmt.Printf("Результат сокращения строки %s: %s\n\n", str, res)

	stringList := []string{
		"aaabbccccdeeeff",
		"jjjsllkkkkklllllll",
		"kkkkkffffffjjjjrrrr",
		"nnnnttttlttttbbbbbb",
		"zzzzzzzzgggggeeeeeeeellll[[[[[",
		"tttttmmmmmggggggg",
		"ooooooooooeeeeeekkkkknnnnnccccc",
		"fffffiiiiiinnnnrrrrrrrrmmmmmmmmm",
		"tttttttggggggnnnnnnnneeeeeee",
	}

	m := make(map[string]string)
	shortString.GetParallelStringsProcessing(stringList, m)

	fmt.Println("Результат параллельной обработки строк:")
	for k, v := range m {
		fmt.Printf("%s -> %s\n", k, v)
	}
}
