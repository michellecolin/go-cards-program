package main

func main() {
	cardsTwo := newDeckFromfile("card-test")
	cardsTwo.print()
	cardsTwo.shuffle()
	cardsTwo.print()

	cards := newDeck()
	cards.print()
}
