package main

import "fmt"

func main(){
	
	// variable declaration
    var name ="rakib"

	//if wat to use constant value/fixed value
	const PLACE ="dhaka" 
    var age =26
    var gpa  =3.11

	
   

	fmt.Printf("%v is a student\n",name)
	fmt.Printf("%v is %v years old\n",name, age)
	fmt.Printf("%v has got %v/4 in ssc\n",name, gpa)
	fmt.Printf("%v is a student\n",name)
	fmt.Printf("%v is from %v\n",name, PLACE)
}