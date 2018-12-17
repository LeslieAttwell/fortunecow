/*
*************************************************
Synopsis: Simple Fortune cookie api
 Date: 16 December 2018
 Usage: 
	- $URL/fortune
	- $URL/cow
	- $URL/tux
	- $URL/tuxtips

To Do:
	* create common interface for exec function
	* add figlet support
*************************************************
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"os/exec"
)

/* 
***********
*  Struct *
***********
*/
type Fortune struct {
    Cookie string `json:"cookie"`
}

type Fortunes []Fortune 

/* 
*****************************
*  Fortune Cookie  - Cowsay *
*****************************
*/
func returnCowsay(w http.ResponseWriter, r *http.Request) {
	cmd := "/usr/games/fortune|/usr/games/cowsay -b"
	command, err := exec.Command("bash","-c",cmd).Output() 
    if err != nil {
        log.Fatal(err)
    }
	
	fortunes := Fortunes{
		Fortune{Cookie: string(command)},
	}
	fmt.Println("Endpoint Hit: returnFortuneCookie")
	json.NewEncoder(w).Encode(fortunes)
}

/* 
**************************
*  Fortune Cookie  - Tux *
**************************
*/
func returnTuxsay(w http.ResponseWriter, r *http.Request) {
	cmd := "/usr/games/fortune|/usr/games/cowsay -f tux"
	command, err := exec.Command("bash","-c",cmd).Output() 
    if err != nil {
        log.Fatal(err)
    }
	
	fortunes := Fortunes{
		Fortune{Cookie: string(command)},
	}
	fmt.Println("Endpoint Hit: returnFortuneCookie")
	json.NewEncoder(w).Encode(fortunes)
}
/* 
****************************
*  Fortune Cookie          *
****************************
*/
func returnFortune(w http.ResponseWriter, r *http.Request) {
	cmd := "/usr/games/fortune"
	command, err := exec.Command("bash","-c",cmd).Output() 
    if err != nil {
        log.Fatal(err)
    }
	
	fortunes := Fortunes{
		Fortune{Cookie: string(command)},
		//Fortune{Cookie: string(fortuneCommand)},
	}
	fmt.Println("Endpoint Hit: returnFortuneCookie")
	json.NewEncoder(w).Encode(fortunes)
}

/* 
****************************
*  Ubuntu Tips Server Tips *
****************************
*/
func returnTuxTips(w http.ResponseWriter, r *http.Request) {
	cmd := "/usr/games/fortune ubuntu-server-tips| /usr/games/cowsay -f tux"
	command, err := exec.Command("bash","-c",cmd).Output() 
    if err != nil {
        log.Fatal(err)
    }
	
	fortunes := Fortunes{
		Fortune{Cookie: string(command)},
		//Fortune{Cookie: string(fortuneCommand)},
	}
	fmt.Println("Endpoint Hit: returnFortuneCookie")
	json.NewEncoder(w).Encode(fortunes)
}

/* 
*********************************
*  Home Page  (used for testing)*
*********************************
*/
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage Endpoint Hit")
}

/* 
*******************************
*  Request Handler  and Serve *
*******************************
*/
func handleRequests() {
	http.HandleFunc("/",homePage) // Point "homePage" to default Fortune function"
	http.HandleFunc("/fortune", returnFortune)
	http.HandleFunc("/cow", returnCowsay)
	http.HandleFunc("/tux", returnTuxsay)
	http.HandleFunc("/tuxtips", returnTuxTips)
	log.Fatal(http.ListenAndServe(":8081",nil))
}

/* 
**********************
*  Main App function *
**********************
*/
func main() {
	handleRequests()
}
