package main

func main() {
	//cards := newDeck()

	//hand, remainingDeck := deal(cards, 5)

	// fmt.Println(hand.toString())
	// fmt.Println(remainingDeck.toString())

	//hand.saveToFile("card-test")

	cardsTwo := newDeckFromfile("card-test")
	cardsTwo.print()
}
