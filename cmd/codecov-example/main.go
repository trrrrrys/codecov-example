package main

import (
	"github.com/trrrrrys/codecov-example/internal/cmd1"
	"github.com/trrrrrys/codecov-example/internal/cmd2"
)

func main() {
	_ = cmd1.Command1(1)
	_ = cmd1.Command1(2)
	_ = cmd2.Command2(1)
	_ = cmd2.Command2(2)
}
