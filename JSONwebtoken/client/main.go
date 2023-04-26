package main

import (
	"fmt"
	"time"

	jwt "github.com/elliotforbes/go-jwt-tutorial/client"
)

var mySigningKey = []byte("mysecretscript")

func GenerateJWT() (string, error) {

	token := jwt.New(jwt.SingningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "hello"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}
	return "", err
}

func main() {
	fmt.Println("My simple clinet")

	tokenString, err := GenerateJWT()

}
