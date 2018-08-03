package yukpiz

import (
	"fmt"
)

func UniqInts(values []int) []int {
	m := make(map[int]bool)
	uniq := []int{}
	for _, value := range values {
		if !m[value] {
			m[value] = true
			uniq = append(uniq, value)
		}
	}
	return uniq
}

func FizzBuzz(n int) {
	for i := 1; i <= n; i += 1 {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
