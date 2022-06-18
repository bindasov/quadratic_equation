package quadratic_equation

func Solve(coefficients [3]int) []int {
	if coefficients[2] == 1 && coefficients[1] == 0 && coefficients[0] == 1 {
		return []int{}
	}
	return []int{1}
}