package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const delayMonitoramento = 5 //secunds
const qtdMonitoramentos = 3 

func main() {

	exibeIntroducao()

	for {
		exibeMenu()
		comandoMenu := lerComando()

		switch comandoMenu {
			case 1:
				iniciarMonitoramento()
			case 2:
				imprimirLogs()
			case 0:
				fmt.Println("Saindo...")
				os.Exit(0)
			default:
				fmt.Println("Não conheço este comando!")
				os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Carlos" 
	versao := 1.0

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este sistema está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monotoramento;")
	fmt.Println("2 - Exibir Logs;")
	fmt.Println("0 - Sair do programa.")
}

func lerComando() int {
	var comandoMenu int
	fmt.Scan(&comandoMenu)

	fmt.Println("Opção escolhido foi:", comandoMenu)
	return comandoMenu
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := lerArquivoSites()

	for i := 0; i < qtdMonitoramentos; i++ {
		fmt.Println("--------------------------------------------------------")
		for _, site := range sites {
			testarConexao(site)
		}
		fmt.Println("--------------------------------------------------------")
		time.Sleep(delayMonitoramento * time.Second)
	}	
}

func testarConexao(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro ao testar conexão:", err)
		return
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, "foi carregado com sucesso!")
		registrarLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registrarLog(site, false)
	}
}


func lerArquivoSites() []string {
	var sites []string
	
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro na leitura do arquivo:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	return sites
}

func registrarLog (site string, status bool) {
	arquivo, err := os.OpenFile("logs.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro ao ler o arquivo:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") +" - "+ site + "- Online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimirLogs() {
	arquivo, err := os.ReadFile("logs.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro ao ler o arquivo:", err)
	}

	fmt.Println(string(arquivo))
}