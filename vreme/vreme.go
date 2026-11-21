package vreme

import "fmt"

type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string) {
	if _, obstaja := postaje[oznaka]; obstaja {
		fmt.Printf("Postaja %s že obstaja.\n", oznaka)
		return
	}

	postaje[oznaka] = Postaja{ime: ime, kraj: kraj}
}

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

func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")

	for lb, postaja := range postaje {
		fmt.Printf("%s - %s (%s): %v\n", lb, postaja.ime, postaja.kraj, postaja.meritve)
	}
}

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
