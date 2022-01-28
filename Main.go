package main

import (
	"fmt"
	"oppgave/rivercrossing/state"
)

func main() {
	fmt.Println(state.ViewState())
	state.PutInKylling()
	fmt.Println(state.ViewState())
}
