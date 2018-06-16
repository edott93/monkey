package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/edott93/monkey/lexer"
	"github.com/edott93/monkey/token"
)

// PROMPT is handles the input to the REPL by giving an indicator of where input is going.
const PROMPT = ">> "

// Start initiates the REPL and currently lexes the input into tokens
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
