package main

import "fmt"

func mypanic(){
	var x,y=30,0
	var c = x/y
	fmt.Println(c)
}

func main(){
	fmt.Println("test panic")
	defer func(){
		fmt.Println("last func")
	}()
	defer func(){
		if err := recover();err!=nil{
			fmt.Println("err:",err)
		}
	}()

	mypanic()
	fmt.Println("end")
}