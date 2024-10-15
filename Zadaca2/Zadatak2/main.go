package main

import "fmt"

type Adresa struct {
	grad  string
	ulica string
}
type Osoba struct {
	ime     string
	prezime string
	godine  int
	Adresa  Adresa
}

func main() {
	osoba1 := Osoba{
		ime:     "Karla",
		prezime: "Fišić",
		godine:  22,
		Adresa: Adresa{
			grad:  "Uskoplje",
			ulica: "kralja Tomislava",
		},
	}
	osoba1.Opis()

}

func (o Osoba) Opis() {
	fmt.Printf("\n%s %s, %d godine, živi u %s, %s", o.ime, o.prezime, o.godine, o.Adresa.grad, o.Adresa.ulica)
}
