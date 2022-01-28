package state

var state = "[kylling rev korn hs ---\\ \\__/ _________________/---]"

func ViewState() string {
	// Sjekke data som er lagret ("kylling til venstre", "rev til venstre")
	return state
}
func PutInKylling() {
	state = "[rev korn hs ---\\ \\_kylling_/ _________________/---]"
}
func PutInRev() {
	state = "[rev korn hs ---\\ \\_Rev_/ _________________/---]"
}
func PutInKorn() {
	state = "[rev korn hs ---\\ \\_Korn_/ _________________/---]"
}
