package main

import (
	"fmt"
	"personal/mal/go/src/readline"
)

const PROMPT = "user> "

func READ() (string, error) {
	return readline.Readline(PROMPT)
}

func EVAL(str string) string {
	return str
}

func PRINT(str string) {
	fmt.Println(str)
}

func rep() {

	for {
		in, err := READ()
		if err != nil {
			break
		}
		eval := EVAL(in)
		PRINT(eval)
	}
}

func main() {
	rep()
}
