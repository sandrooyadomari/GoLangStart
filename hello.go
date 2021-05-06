package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	introduction()

	for {
		menu()
		command := readCommand()

		switch command {
		case 1:
			monitoring()
		case 2:
			fmt.Println("Exibir")
		case 0:
			fmt.Println("Quit")
			os.Exit(0)
		default:
			fmt.Println("ta errado neh")
			os.Exit(-1)
		}
	}
}

func introduction() {
	nome := "Sandro"
	versao := 1
	fmt.Println("Hello", nome)
	fmt.Println("Version: ", versao)
}

func readCommand() int {
	var command int
	fmt.Scan(&command)

	return command
}

func menu() {
	fmt.Println("1 - Monitor")
	fmt.Println("2 - Show monitor")
	fmt.Println("0 - Quit")
}

func monitoring() {
	fmt.Println("Initializing")

	webSites := []string{"https://google.com", "https://www.alura.com.br", "https://www.caelum.com.br"}

	for _, webSite := range webSites {
		testUrl(webSite)
	}
}

func testUrl(site string) {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println(site, "is up!")
	} else {
		fmt.Println("It's Down!")
	}
}
