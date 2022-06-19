package quadratic_equation

import (
	"errors"
	"math"
)

const (
	Epsilon = 10e-5
)

var (
	ErrNotQuadratic = errors.New("a == 0")
	ErrInvalidCoefficients = errors.New("invalid coefficients")
)

func Solve(a, b, c float64) ([]float64, error) {
	var roots []float64
    var discriminant float64

    if a < Epsilon {
    	return roots, ErrNotQuadratic
    }

    discriminant = (b * b) - (4 * a * c)
    switch {
    case discriminant > Epsilon:
        roots = append(roots, -b + math.Sqrt(discriminant) / (2 * a))
        roots = append(roots, -b - math.Sqrt(discriminant) / (2 * a))
    case discriminant < Epsilon && discriminant > -Epsilon:
        roots = append(roots, -b / (2 * a))
        roots = append(roots, -b / (2 * a))
    case discriminant < 0:
        return roots, nil
    default:
    	return roots, ErrInvalidCoefficients
    }
    return roots, nil
}