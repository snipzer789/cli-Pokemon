package main

import (
	"fmt"
	"encoding/json"
	"os"
	"log"
	"math"
)

type Pokemon struct {
    NAME    string
	GENDER 	string
	SPECIES string

	MOVE1	string
	MOVE1_POWER	float64
	MOVE1_PP	float64
	MOVE1_DAMAGEING	string
	MOVE2	string
	MOVE2_POWER	float64
	MOVE2_PP	float64
	MOVE2_DAMAGEING	string
	MOVE3	string
	MOVE3_POWER	float64
	MOVE3_PP	float64
	MOVE3_DAMAGEING string
	MOVE4	string
	MOVE4_POWER float64
	MOVE4_PP 	float64
	MOVE4_DAMAGEING string

	LEVEL 	float64
	IVHP 	float64
    IVATK	float64
    IVDefence	float64
    IVSpeed	float64
    IVSP_ATK	float64
    IVSP_DEF	float64

    EVHP	float64
    EVATK	float64
    EVDEF	float64
    EVSPEED	float64
    EVSP_ATK	float64
    EVSP_DEF	float64
}

func roundFloat(val float64, precision uint) float64 {
    ratio := math.Pow(10, float64(precision))
    return math.Round(val*ratio) / ratio
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
	fmt.Print("\033[H\033[2J") 
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
	Data, err := os.ReadFile("/workspaces/cli-Pokemon/main/data/test/1.json")
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
	var onehund float64 = 0.01
	var onefourth float64 = 0.25
	var health float64 = (onehund * (2 * 20 + Pokemon1.IVHP + (onefourth * Pokemon1.EVHP)) * Pokemon1.LEVEL) + Pokemon1.LEVEL + 10
	
	var health_bar string = "███████████████"
	fmt.Print("\033[H\033[2J") 
	fmt.Println("Name:", Pokemon1.NAME)
	fmt.Println("Level:", Pokemon1.LEVEL)
	fmt.Println("Health:", roundFloat(health, 0))
	fmt.Println(health_bar)
	fmt.Println("Move 1:", Pokemon1.MOVE1, "| PP:", Pokemon1.MOVE1_PP)
	fmt.Println("Move 2:", Pokemon1.MOVE2, "| PP:", Pokemon1.MOVE2_PP)
	fmt.Println("Move 3:", Pokemon1.MOVE3, "| PP:", Pokemon1.MOVE3_PP)
	fmt.Println("Move 4:", Pokemon1.MOVE4, "| PP:", Pokemon1.MOVE4_PP)
	

}

func Continue(){
	fmt.Print("\033[H\033[2J") 
	fmt.Println("Contiune")
	fmt.Println("------------------")
	fmt.Println("Why are you here")
	fmt.Println("------------------")
	fmt.Scanln()
	fmt.Println("------------------")
}

func main(){
	fmt.Print("\033[H\033[2J") 
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
