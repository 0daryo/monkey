package parser

import (
	"fmt"

	"github.com/0daryo/monkey/lexer"
	"github.com/0daryo/monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Errors() []string {
	return p.Errors()
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token tobe %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
