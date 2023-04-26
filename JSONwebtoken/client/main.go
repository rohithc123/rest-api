package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// var mySigningKey = os.Get("MY_JWT_TOKEN")
// don't do like this,since it is being uploaded in a public repo
var mySigningKey = []byte("mysecretscript")

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	req.Header.Set("Token", validToken)

	res, err := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, string(body))
}

func GenerateJWT() (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "hello"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func handleRequest() {
	http.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":9091", nil))
}

func main() {
	fmt.Println("My simple client")

	handleRequest()
}
