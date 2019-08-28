package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var mySigningKey = []byte("mysupersecretphase")

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Fprintf(w, validToken)
}
func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "Elliot Forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Somthing happen")
		return "", err
	}
	return tokenString, nil
}

func handelRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8000", r))

}

func main() {
	handelRequest()
}
