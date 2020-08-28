package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int64
	fmt.Println("Please, write a,b,c")
	_, err := fmt.Scanf("%d", &a)
	if err != nil {
		panic(err.Error())
	}
	_, err = fmt.Scanf("%d", &b)
	if err != nil {
		panic(err.Error())
	}
	_, err = fmt.Scanf("%d", &c)
	if err != nil {
		panic(err.Error())
	}
	x1, x2 := CountD(a, b, c)
	fmt.Printf("x1 = %.2f\n", x1)
	fmt.Printf("x2 = %.2f", x2)

}
func CountD(a int64, b int64, c int64) (x1 float64, x2 float64) {
	p := math.Pow(float64(b), 2)
	D := p + float64(4*a*c)
	sqrtD := math.Sqrt(D)
	fmt.Println(-float64(b))
	x1 = (-float64(b) + sqrtD) / 2 * float64(a)
	x2 = (-float64(b) - sqrtD) / 2 * float64(a)
	return x1, x2
}
