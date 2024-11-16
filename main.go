package main

import "fmt"

type GameState struct {
	Coins     int
	Life      int
	Blood     int
	Character string
}

func main() {
	players := map[string]GameState{
		"Luan": GameState{
			Coins: 100,
			Life: 3,
			Blood: 100,
			Character: "Soldier",
		},
	}

	character := players["Luan"].Character
	switch character {
	case "Soldier":
		fmt.Println("ITS A SOLDIER")
	default:
		fmt.Println("Good Question!")
	}

	character = players["Luan"]
	if character.Blood <= 0 {
		character.Life -= 1
	}

	fmt.Println(character)
}
