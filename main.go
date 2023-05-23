package main

import (
	"main/spring"
)

func main() {
	sp := spring.NewSpring()
	sp.HandleInputs()
	sp.FindSolution()
	sp.Animate(600, 600, 2, "spring.gif")
}
