package cmd1_test

import (
	"testing"

	"github.com/trrrrrys/codecov-example/internal/cmd1"
)

func TestCommand1(t *testing.T) {
	tests := []struct {
		name string
		in   int
	}{
		{
			"case odd",
			1,
		},
		{
			"case even",
			2,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := cmd1.Command1(tt.in)
			if err != nil {
				t.Fatalf("error: %+v", err)
			}
		})
	}
}
