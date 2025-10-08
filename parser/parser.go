package parser

import (
	"github.com/isaacgr/go-monkey-interpreter/ast"
	"github.com/isaacgr/go-monkey-interpreter/lexer"
	"github.com/isaacgr/go-monkey-interpreter/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// read two tokens, so curToken and peekToken are both set

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
