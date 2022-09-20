package main

import "fmt"

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	fmt.Print(lcm(lcm(a, b), c))
}
