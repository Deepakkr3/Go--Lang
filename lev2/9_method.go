package main

import "fmt"

/*

1. methodd shuld associate with any type not mendotry to any struct anly
2. there are thow type of reciver value reciver and pointer reciver
3.


*/

type Ractangle struct {
	length int
	weadth int
}

//method with value recier
func (r Ractangle) area() int {
	return r.length * r.weadth
}

//method with pointer reciver
//you can modify ractange type instance
func (r *Ractangle) AreaPointer(factor int) int {
	r.length=r.length+factor
	return r.length * r.weadth
}
func main() {
	ract := Ractangle{
		length: 10,
		weadth: 20,
	}
	fmt.Println(" value ract", ract.area())
	fmt.Println("length is before ",ract.length)
	fmt.Println(" pointer ract", ract.AreaPointer(2))
	fmt.Println("length is after ",ract.length)
}
