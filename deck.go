package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cardSuit := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	card := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "King", "Queen"}
	cards := deck{}

	for _, suit := range cardSuit {
		for _, value := range card {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) toByte() []byte {
	return []byte(d.toString())
}

func (d deck) writeDeckToFile(filename string) error {
	return ioutil.WriteFile(filename+".txt", d.toByte(), 0666)
}

func readDeckFromFile(filename string) deck {
	filebyte, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := strings.Split(string(filebyte), ",")
	return deck(s)
}

func (d deck) shuffle() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
