package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/MatoMatto/vremenskepostaje/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "vremenskepostaje",
		Usage: "Beleži meritve vremenskih postaj",
		Flags: []cli.Flag{
			&cli.FloatFlag{
				Name:  "minTemp",
				Usage: "Najnižja dovoljena temperatura v °C",
				Value: -50,
			},
			&cli.FloatFlag{
				Name:  "maxTemp",
				Usage: "Najvišja dovoljena temperatura v °C",
				Value: 50,
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Usage: "Najmanjše število meritev za izračun statistike",
				Value: 3,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp := cmd.Float("minTemp")
			maxTemp := cmd.Float("maxTemp")
			stMeritev := cmd.Int("stMeritev")
			run(minTemp, maxTemp, stMeritev)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(minTemp, maxTemp float64, stMeritev int) {
	postaje := make(map[string]vreme.Postaja)

	vreme.DodajPostajo(postaje, "KP01", "Koper Luka", "Koper")
	vreme.DodajPostajo(postaje, "KR01", "Kredarica", "Triglav")
	vreme.DodajPostajo(postaje, "LJ01", "Ljubljana Bežigrad", "Ljubljana")
	vreme.DodajPostajo(postaje, "MB01", "Maribor Tabor", "Maribor")
	vreme.DodajPostajo(postaje, "CE01", "Celje", "Celje")

	for _, m := range []float64{24.5, 27, 26.1} {
		vreme.DodajMeritev(postaje, "KP01", m, minTemp, maxTemp)
	}
	for _, m := range []float64{2, -1.5} {
		vreme.DodajMeritev(postaje, "KR01", m, minTemp, maxTemp)
	}
	for _, m := range []float64{18.2, 21.5, 19.8} {
		vreme.DodajMeritev(postaje, "LJ01", m, minTemp, maxTemp)
	}
	for _, m := range []float64{15.2, 16.5, 17.8} {
		vreme.DodajMeritev(postaje, "MB01", m, minTemp, maxTemp)
	}
	vreme.DodajMeritev(postaje, "CE01", 5.0, minTemp, maxTemp)

	vreme.IzpisMeritev(postaje)

	vreme.DodajMeritev(postaje, "KP01", 28.4, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "KR01", 3.5, minTemp, maxTemp)

	fmt.Println("------- Bi mogl dobit napako: -------")
	vreme.DodajMeritev(postaje, "AB01", 3.5, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "LJ01", 200.0, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "MB01", -100.0, minTemp, maxTemp)
	vreme.DodajPostajo(postaje, "LJ01", "Ljubljana", "Ljubljana")
	fmt.Println("-------")

	vreme.IzpisMeritev(postaje)

	fmt.Println("Poročilo:")
	vreme.IzpisPorocila(postaje, stMeritev)
}
