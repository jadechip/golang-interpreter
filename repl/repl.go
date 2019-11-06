package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/skogman/interpreter/lexer"
	"github.com/skogman/interpreter/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		// Scan advances the Scanner to the next token
		// Returns false when the scan stops
		scanned := scanner.Scan()
		// TODO: Error handling perhaps?
		// if scanner.Err() != nil {
		// 	fmt.Printf("error: %s\n", scanner.Err())
		// }
		if !scanned {
			return
		}

		line := scanner.Text()
		// Init new lexer with return value from scanner
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
