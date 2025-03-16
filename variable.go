package main

import "fmt"

func main(){
	
	// variable declaration
    // var name ="rakib"
	// var place ="dhaka"
    // var age =26
    // var gpa  =3.11

	//if want dynamic variable declaration
    name :="rakib"
	place :="dhaka"
    age :=26
    gpa  :=3.11
   

	fmt.Println(name,"is a student")
	fmt.Println(name,"is" ,age, "years old")
	fmt.Println(name,"has got", gpa,"/4 in ssc")
	fmt.Println(name,"is a student")
	fmt.Println(name,"is from",place)
}