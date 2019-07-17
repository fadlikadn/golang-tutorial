package main

func main() {
	//cards := []string{"Ace of Diamonds", newCard()}
	//cards := deck{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of Spades")

	//for i, card := range cards {
	//	fmt.Println(i, card)
	//}

	cards := newDeck()

	cards.print()
}

func newCard() string {
	return "Five of Diamons"
}
