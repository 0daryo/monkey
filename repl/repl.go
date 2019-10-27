package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/0daryo/monkey/lexer"
	"github.com/0daryo/monkey/token"
)

const PROMPT = "heyyy "

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
		// if last 'tok = l.NextToken()' doesnt exist, it would be an inf loop
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
