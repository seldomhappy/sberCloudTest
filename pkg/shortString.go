package pkg

import (
	"strconv"
	"strings"
	"sync"
	"time"
)

func Get(str string) string {
	var sb strings.Builder
	stringOfRuneSlice := []rune(str)
	count := 0
	time.Sleep(2 * time.Second)
	l := len(stringOfRuneSlice)
	for i := 0; ; i++ {
		count++
		if i+1 == l {
			sb.WriteString(strconv.Itoa(count))
			sb.WriteString(string(stringOfRuneSlice[i]))
			break
		}
		if stringOfRuneSlice[i] != stringOfRuneSlice[i+1] {
			sb.WriteString(strconv.Itoa(count))
			sb.WriteString(string(stringOfRuneSlice[i]))
			count = 0
		}
	}
	return sb.String()
}

func GetParallelStringsProcessing(stringList []string, data map[string]string) {
	ch := make(chan string)
	var wg sync.WaitGroup
	var result sync.Map

	wg.Add(len(stringList))
	for i := 0; i <= 16; i++ {
		go func() {
			for val := range ch {
				result.Store(val, Get(val))
				wg.Done()
			}
		}()
	}

	for idx := range stringList {
		ch <- stringList[idx]
	}
	close(ch)
	wg.Wait()


	result.Range(func(in, out interface{},) bool {
		data[in.(string)] = out.(string)
		return true
	})
}
