package game

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func InitGame() {
	fmt.Println("O jogo é simples, tente acertar, o número é um inteiro de 1 a 100")

	x := rand.Int64N(101)
	scannerInt := bufio.NewScanner((os.Stdin))
	tries := [15]int64{}

	for i := range tries {
		fmt.Println("Qual é o número?")
		scannerInt.Scan()
		try := scannerInt.Text()
		try = strings.TrimSpace(try)
		tryInt, err := strconv.ParseInt(try, 10, 64)

		if err != nil {
			fmt.Println("Hm, seu número precisa ser um número inteiro\n" +
				"Tente novamente")
			return
		}

		switch {
		case tryInt < x:
			fmt.Println("eita, voce errou jogador, o número que precisa acertar é um valor maior", tryInt)
		case tryInt > x:
			fmt.Println("opa, voce errou jogador, o número que precisa acertar é um valor menor", tryInt)
		case tryInt == x:
			fmt.Printf("uau, simplesmente temos um vencedor, o número era: %d\n"+
				"voce acertou em %d tentativas\n"+
				"obrigado por jogar, um verdadeiro genio(a)",
				x, i+1)
			return
		}
		tries[i] = tryInt

	}
	fmt.Printf("vish, parece que voce não acertou o número, que pena, o número era: %d\n"+
		"voce teve 15 tentativas.\n"+
		"boa sorte na proxima vez!",
		x)
}
