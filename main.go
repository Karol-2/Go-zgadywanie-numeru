package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var players map[string]int
	players = wczytaj()
	//fmt.Println(players)
	gra(players)
}

func gra(players map[string]int) {
	fmt.Println("============================================")
	min := 1
	max := 100
	szukana := rand.Intn(max-min) + min
	podejscia := 1
	var odp string

	fmt.Println("Teraz będziesz zgadywać liczbę, którą wylosowałem")
	fmt.Println("Po wpisaniu \"koniec\" nastąpi zakończenie zabawy")

	for {
		//fmt.Println("szukamy liczby", szukana)
		fmt.Print("Podaj liczbę: ")
		fmt.Scan(&odp)
		if odp == "koniec" {
			fmt.Println("Żegnaj :( ")
			zapisz(players)
			leaderboard(players)
			os.Exit(1)
		}

		odp, err := strconv.Atoi(odp)
		if err != nil {
			fmt.Println("Złe dane wejściowe")
			break
		}

		if odp < szukana {
			fmt.Println("### Za mała ###")
			podejscia++
		} else if odp > szukana {
			fmt.Println("### Za duża ###")
			podejscia++
		} else {
			fmt.Println("Zgadł_ś, Gratulacje!")
			fmt.Println("Ilość podejść: ", podejscia)

			var imie string
			fmt.Print("Podaj swoje imie:")
			fmt.Scan(&imie)
			aktualizacjaWyniku(imie, podejscia, players)
			rekordGlobalny(podejscia, players)

			var powtorka string
			fmt.Println("Czy gramy jeszcze raz? [tak/nie]")
			fmt.Scan(&powtorka)
			if powtorka == "tak" {
				gra(players)
			} else if powtorka == "nie" {
				zapisz(players)
				leaderboard(players)
				os.Exit(1)
			} else {
				fmt.Println("Nie zrozumiałem, więc wychodzisz")
				zapisz(players)
				leaderboard(players)
				os.Exit(1)
			}

		}
	}
}
func aktualizacjaWyniku(imie string, wynik int, players map[string]int) {
	if _, ok := players[imie]; ok {
		if wynik < players[imie] {
			fmt.Println("Pobito rekord osobisty!")
			players[imie] = wynik
		}
	} else {
		players[imie] = wynik
	}
}
func rekordGlobalny(wynik int, players map[string]int) {

	for player := range players {
		if players[player] > wynik {
			fmt.Println("Pobito rekord globalny!")
			break
		}
	}

}

func leaderboard(players map[string]int) {

	fmt.Println("Tak prezentuje się tablica wyników")
	fmt.Println("-----------")
	for player, score := range players {
		fmt.Println(player, score)
	}
	fmt.Println("-----------")
}

func zapisz(players map[string]int) {
	currentTime := time.Now()
	f, err := os.OpenFile("scores.txt", os.O_CREATE, 0644)

	if err != nil {
		fmt.Println(err)
	}

	for player, score := range players {

		str := player + "," + strconv.Itoa(score) + "," + currentTime.Format("02-01-2006") + "\n"
		_, err2 := f.WriteString(str)

		if err2 != nil {
			fmt.Println(err2)
		}

	}
}

func wczytaj() map[string]int {
	file, err := os.OpenFile("scores.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return make(map[string]int)
	}

	scores := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			fmt.Printf("Błędny format linii: %s", line)
		}
		score := 0
		_, err := fmt.Sscanf(parts[1], "%d", &score)
		if err != nil {
			fmt.Printf("Błędny format wyniku: %s", parts[1])
		}
		scores[parts[0]] = score
	}

	if err := scanner.Err(); err != nil {
		fmt.Print(err)
	}

	return scores
}
