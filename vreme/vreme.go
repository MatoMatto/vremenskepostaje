// Package vreme simulira vremensko postajo, ki generira naključne vremenske podatke.
//
// Primer uporabe:
//
//	postaje := make(map[string]vreme.Postaja)
//	vreme.DodajPostajo(postaje, "LJ01", "Ljubljana Bežigrad", "Ljubljana")
//	vreme.DodajMeritev(postaje, "LJ01", 18.2, -50, 50)
//	vreme.IzpisMeritev(postaje)
//	vreme.IzpisPorocila(postaje, 3)

package vreme

import "fmt"

// Postaja – vremenska postaja z imenom, krajem in izmerjenimi temperaturami v °C.
// Postajo ustvarimo s funkcijo `DodajPostajo`
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// DodajPostajo v slovar postaj doda novo postajo z dano oznako, imenom in krajem.
// Če postaja z dano oznako že obstaja, izpiše obvestilo.
func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string) {
	if _, obstaja := postaje[oznaka]; obstaja {
		fmt.Printf("Postaja %s že obstaja.\n", oznaka)
		return
	}

	postaje[oznaka] = Postaja{ime: ime, kraj: kraj}
}

// DodajMeritev postaji z dano oznako doda meritev temperature, če velja minTemp <= temperatura <= maxTemp. Če oznaka postaje ne obstaja ali meritev ni veljavne, izpiše obvestilo.
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura, minTemp, maxTemp float64) {
	postaja, obstaja := postaje[oznaka]

	if !obstaja {
		fmt.Printf("Postaja %s ne obstaja.\n", oznaka)
		return
	}

	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Printf("Temperatura %.1f °C je izven obsega.\n", temperatura)
		return
	}

	postaja.meritve = append(postaja.meritve, temperatura)
	postaje[oznaka] = postaja
}

// statistika vrne najnižjo, najvišjo in povprečno temperaturo postaje z dano oznako.
// Če postaja ne obstaja ali ima manj kot stMeritev meritev, vrne ok == false.
func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	postaja, obstaja := postaje[oznaka]

	if !obstaja || len(postaja.meritve) < stMeritev || len(postaja.meritve) == 0 {
		return 0, 0, 0, false
	}

	najnizja, najvisja = postaja.meritve[0], postaja.meritve[0]

	vsota := 0.0

	for _, m := range postaja.meritve {
		if m < najnizja {
			najnizja = m
		}
		if m > najvisja {
			najvisja = m
		}

		vsota += m
	}

	povprecna = vsota / float64(len(postaja.meritve))

	return najnizja, najvisja, povprecna, true
}

// IzpisMeritev izpiše oznake, imena in kraje vseh postaj ter njihove meritve
func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")

	for lb, postaja := range postaje {
		fmt.Printf("%s - %s (%s): %v\n", lb, postaja.ime, postaja.kraj, postaja.meritve)
	}
}

// IzpisPorocila za vsako postajo izpiše ime ter njihovo najnižjo, najvišjo in povprečno temperaturo z opisom.
// Če ima postaja manj kot stMeritev meritev, izpiše "premalo meritev".
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	for lb, postaja := range postaje {
		fmt.Printf("%s: ", postaja.ime)
		mn, mx, avg, ok := statistika(postaje, lb, stMeritev)

		if !ok {
			fmt.Printf("premalo meritev\n")
			continue
		}

		var opis string
		switch {
		case avg >= 25.0:
			opis = "Vroče"
		case avg >= 15.0:
			opis = "Toplo"
		case avg >= 5.0:
			opis = "Hladno"
		default:
			opis = "Mrzlo"
		}

		fmt.Printf("min %.1f °C, max %.1f °C, povprečje %.1f °C -> %s\n", mn, mx, avg, opis)
	}
}
