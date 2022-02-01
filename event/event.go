// lage array for items(person, kylling, rev, korn)? bruke conditional statements?
// https://www.golangprograms.com/go-language/arrays.html
// metoder som kan brukes: start() GetInput() CheckInput()
package event

import "fmt"

// arrays med alle items som kan/må lastes på båten
// få input fra bruker om hvilke items som skal flyttes på
// båtstate, true=vest false=øst
var items [3]string = [3]string{"høne", "rev", "korn"}
var itemToPut string

//Test for å velge hvilken slice "item" man vil laste på båten
func Test(items [3]string) string {
	return "[|rev|korn|  vvv\\_høne_pers_/ _________________/øøø]"
}

// viser initiell status
func Start() {
	fmt.Printf("[%v|%v|%v|  vvv\\__/ _________________/øøø]"+"\n", items[0], items[1], items[2])
}

func PuttIBåt() {
	fmt.Println("tester")
}

func GetInputInit() {
	fmt.Printf("\n"+"Skriv hva du vil putte i båten, du kan velge mellom %v, %v og %v"+"\n", items[0], items[1], items[2])

	fmt.Scanln(&itemToPut)

}

func CheckUserInput() {

	if itemToPut != items[0] && itemToPut != items[1] && itemToPut != items[2] {
		fmt.Printf("Du må velge mellom %v, %v, og %v"+"\n", items[0], items[1], items[2])
	}
}

func FlyttH() {
	fmt.Println(fmt.Printf("[|%v|%v|  vvv\\________________\\___//øøø %v]", items[1], items[2], items[0]))

}

func PutH() {
	fmt.Println(fmt.Printf("[|%v|%v|  vvv\\_%v_/ _________________/øøø]", items[1], items[2], items[0]))
}

func IsValid() {

}

func Final() {
	fmt.Printf("[vvv\\________________\\___//øøø %v|%v|%v|]", items[1], items[2], items[0])
}
