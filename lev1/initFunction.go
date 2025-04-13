package main

import "fmt"

/**

1. init function use to initalize task that exicute before the main fnction exicute you can use in any packege
2. init function not take argument and nither return anything
3. use for initialize state to exicute main function effeceantly like initalize variable ...
4. can have multiple init() int same packege it will exicute in order thy are declear 
5. init function use for initializing packege level state 


**/
func init() {
	fmt.Println("this is init function1")
}
func init() {
	fmt.Println("this is init function2")
}
func init() {
	fmt.Println("this is init function3")
}
func main() {

	fmt.Println("this is main function")

}
