package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

var mySigningKey = []byte("mysecretscript")

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Secret Message cannot be shared")
}

// authorization code
func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error in JWT")
				}
				return mySigningKey, nil

			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}

		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func handleRequest() {

	http.Handle("/", isAuthorized(homePage))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {

	handleRequest()
}
