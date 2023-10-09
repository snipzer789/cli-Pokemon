package main

import (
	"fmt"
	"encoding/json"
	"os"
	"log"
)

type Pokemon struct {
    NAME    string
    TYPE   	string
    HEALTH 	int
	LEVEL 	int
}

func Main_Menu()(string){
	var Continue string = "(unavailable)"
	var Option string
	fmt.Println("Cli-Pokemon")
	fmt.Println("------------------")
	fmt.Println("1) Start")
	fmt.Println("2) Continue", Continue)
	fmt.Println("3) About")
	fmt.Println("4) Exit")
	fmt.Scanln(&Option)
	fmt.Println("------------------")
	if(Continue == "(unavailable)" && Option == "2"){
		Option = "5"
	}
	return Option;
}
func About(){
	fmt.Println("About")
	fmt.Println("------------------")
	fmt.Println("Bored in class nah?")
	fmt.Println("------------------")
	fmt.Scanln()
	fmt.Println("------------------")
}

func Battle_Main(){
	var option string
	fmt.Println("Battle")
	fmt.Println("------------------")
	fmt.Println("1) Test")
	fmt.Println("------------------")
	fmt.Scanln(&option)
	fmt.Println("------------------")
	if(option == "1"){
		Battle_test()
	}
}

func Battle_test(){
	var Pokemon1 Pokemon

	Data, err := os.ReadFile("/workspaces/cli-Pokemon/main/data/1.json")
    if err != nil {
        log.Fatal(err)
    }

    // from json format
    err = json.Unmarshal(Data, &Pokemon1)
 
    if err != nil {
        //if error is not nil
        //print error
        log.Fatal(err)
    }
	fmt.Println(Pokemon1.NAME)
	fmt.Println(Pokemon1.TYPE)
	fmt.Println(Pokemon1.HEALTH)

}

func Continue(){
	fmt.Println("Contiune")
	fmt.Println("------------------")
	fmt.Println("Why are you here")
	fmt.Println("------------------")
	fmt.Scanln()
	fmt.Println("------------------")
}

func main(){
	fmt.Println("------------------")
	option := Main_Menu()
	switch {
	case option == "1":
		Battle_Main()
	case option == "2":
		Continue()
	case option == "3":
		About()
	case option == "4":
		return
	case option == "5":
		Continue()
	}
}
