package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

	hand, remainingDeck := deal(cards, 5)
	fmt.Println("========")
	hand.print()
	fmt.Println("========")
	remainingDeck.print()
}
