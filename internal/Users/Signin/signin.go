package Signin

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"log"
)

type SignIn struct {
	Email    string
	Password string
}

func SIGNIN(s_in SignIn) {
	// arr:=make([]string, 0)
	file, err := os.Open("/home/abduazim/Projects/Golang/registratsiya/malumotlar.txt")
	if err != nil {
		fmt.Println("Faylni ochishda xatolik:", err)
		return
	}
	defer file.Close()
	fmt.Printf("Email:\t")
	fmt.Scanln(&s_in.Email)
	fmt.Printf("Password:\t")
	fmt.Scanln(&s_in.Password)
	lampochka:=false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, s_in.Email) {
			if strings.Contains(line, s_in.Password) {
				fmt.Println("Siz muvaffaqiyatli kirdingiz 🥳🥳🥳")
				lampochka=true
				
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scannerda xatolik:", err)
		return
	}

	if lampochka{
		file3, err := os.Open("/home/abduazim/Projects/Golang/registratsiya/warehouse.txt")
		if err != nil {
			log.Fatalf("Faylni ochishda xatolik: %v", err)
		}
		defer file.Close()


		scanner2 := bufio.NewScanner(file3)

		for scanner2.Scan() {
			line:=scanner2.Text()
			fmt.Println(line)
		}
	}

}

func Sign_in() {
	var s_in SignIn
	SIGNIN(s_in)
}
