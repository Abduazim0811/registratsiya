package main

import (
	sup "Registratsiya/internal/Users/Signup"
	"fmt"
	sin "Registratsiya/internal/Users/Signin"
	// pr "Registratsiya/internal/Product"
	"os"
)

func main() {
	fmt.Println("------------------------")
	fmt.Println("Qaysi xizmatni tanlaysiz")
	
	fmt.Println("[1] Users")
	fmt.Println("[2] Admin")
	fmt.Println("[3] Exit")
	fmt.Println("------------------------")
	var son int
	var son2 int
	fmt.Scanln(&son)
	if son == 1 {
		fmt.Println("[1] Sign up")
		fmt.Println("[2] Sign in")
		fmt.Scanln(&son2)
		if son2==1{
			sup.SIgn_Up()
		}else if son2==2{
			sin.Sign_in()
		}	
	} else if son == 2 {
		
		// pr.PRoduct()
	} else {
		os.Exit(0)
	}
}
