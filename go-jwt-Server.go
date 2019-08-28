package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var mySignKey = []byte("mysupersecretphase")

func homPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super secrite ")
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (i interface{}, e error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error ")
				}
				return mySignKey, nil
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

func handelRequet() {
	r := mux.NewRouter()
	r.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":1000", r))

}

func main() {
	handelRequet()
}
