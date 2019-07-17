package main

func main() {
	//cards := []string{"Ace of Diamonds", newCard()}
	//cards := deck{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of Spades")

	//for i, card := range cards {
	//	fmt.Println(i, card)
	//}

	// Original
	//cards := newDeck()
	//
	//hand, remainingCards := deal(cards, 5)
	//
	//cards.print()
	//fmt.Println("")
	//hand.print()
	//fmt.Println("")
	//remainingCards.print()
	//fmt.Println("")

	cards := newDeck()
	cards.saveToFile("my_cards")
}

func newCard() string {
	return "Five of Diamons"
}
