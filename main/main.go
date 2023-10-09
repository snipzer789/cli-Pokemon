package main

import "fmt";



func Main_Menu()(string){
	var Continue string = "(unavailable)"
	var Option string
	fmt.Println("Cli-Pokemon")
	fmt.Println("-------------------------")
	fmt.Println("1) Start")
	fmt.Println("2) Continue", Continue)
	fmt.Println("3) About")
	fmt.Println("4) Exit")
	fmt.Scanln(&Option)
	if(Continue == "(unavailable)" && Option == "2"){
		Option = "5"
	}
	return Option;
}
func About(){
	fmt.Println("About")
	fmt.Println("------------------")
	fmt.Println("Bored in class nah?")
	fmt.Scanln()
}

func Continue(){
	fmt.Println("Contiune")
	fmt.Println("------------------")
	fmt.Println("Why are you here")
}

func main(){
	option := Main_Menu()
	switch {
	case option == "1":
		fmt.Println("Start")
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
