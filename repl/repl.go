package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/palash25/chimp/lexer"
	"github.com/palash25/chimp/token"
)

const PROMPT = ">> "

func Start(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)

	for {
		fmt.Println(PROMPT)
		inputString := scanner.Scan()

		if !inputString {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
