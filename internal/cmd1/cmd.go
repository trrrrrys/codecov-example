package cmd1

import "log"

func Command1(n int) error {
	if n%2 == 0 {
		log.Println("cmd1.Func1: n is even")
	} else if n%3 == 0 {
		log.Println("saan by nabeatsu")
	} else {
		log.Println("cmd1.Func1: n is odd")
	}
	return nil
}
