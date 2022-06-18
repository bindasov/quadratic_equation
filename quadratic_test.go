package quadratic_equation

import (
    "testing"
    "github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
    tests := []struct {
        name    string
        handler func(*testing.T)
    }{
        {
            name: "no roots",
            handler: func(t *testing.T) {
                coefficients := [3]int{1, 0, 1}
                result := Solve(coefficients)
                require.Equal(t, []int{}, result)
            },
        },
        {
            name: "wrong coefficients",
            handler: func(t *testing.T) {
                coefficients := [3]int{1, 1, 1}
                result := Solve(coefficients)
                require.Equal(t, []int{1}, result)
            },
        },
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            tc.handler(t)
        })
    }
}
