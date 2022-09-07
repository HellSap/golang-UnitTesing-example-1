package main

import "fmt"

func main() {
	fmt.Println(SimpleIntrest(20,4, 2))

}

// func Add(x, y int) int {
// 	return x+y
// }

// func SI(x, y, z float64) float64 {
// 	return (x*y*z)/100
// }


func SimpleIntrest(p , r , t int)int{
    if p <=0 || r <=0 || t <=0 {
        return 0
    }
    return (p*r*t)/100
}