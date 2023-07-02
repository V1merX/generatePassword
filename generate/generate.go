package generate

import (
	"fmt"
	"log"
	"math/rand"
)

type password struct {
	length   int
	password string
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_"

func Start() {
	pass := new(password)
	pass.scanLengthPass()
	pass.generatePassword()
	pass.print()
}

func (p *password) print() {
	fmt.Println(p.password)
}

func (p *password) scanLengthPass() {
	fmt.Print("Введите длину пароля: ")

	_, err := fmt.Scan(&p.length)
	if err != nil {
		log.Fatal(err)
	}
}

func (p *password) generatePassword() {
	b := make([]byte, p.length)

	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	p.password = string(b)
}
