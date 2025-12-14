package lexer

import "github.com/isaacgr/go-monkey-interpreter/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char
}

// NOTE: If you wanted to attach line numbers and file names then this
// could be expanded to accept io.Reader and the filename as opposed to string
func NewLexer(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

// Give us the next character and advance our position
//
// NOTE: Only supports ascii and not the full unicode range, would require
// changing ch from a byte to a rune which would then change how the next chars
// are read
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for null, which signifies we have nothing or EOF
	} else {
		l.ch = l.input[l.readPosition] // Otherwise set it to the next character
	}
	l.position = l.readPosition // position we last read
	l.readPosition += 1         // next character
}

// We look at the current character and return a token depending on which
// character it is
// Before returning the token we advance our pointer into the input so the
// next call to NextToken will already have an updated l.ch field
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	// Whitespace has no meaning, so we need this to prevent failing when
	// parsing the input string if we come across it
	l.skipWhitespace()
	switch l.ch {
	case '=':
		// Note that here and below, we save l.ch to a local variable before
		// calling readChar again
		// This way we can safely advance the lexer so it leaves NextToken
		// with l.position and l.readPosition in the correct state
		// If we were to support more than these two character tokens in the
		// language, we should probably abstract that away in a new method
		// that both peeks and advances 
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{
				Type:    token.EQ,
				Literal: literal,
			}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{
				Type:    token.NOT_EQ,
				Literal: literal,
			}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	// We need to recognize if the current character is a letter, and if it is
	// it needs to read the rest of the identifier/keyword until it encounters
	// a non letter character
	// We then need to determine if it is an identifier or a keyword so that we
	// can use the correct token.TokenType
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			// We return early here since readIdentifier already advances the
			// position, so we dont want to do it again outside of the switch
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	initialPosition := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[initialPosition:l.position]
}

// NOTE: the lack of support for floats, hex, octal notation etc.
func (l *Lexer) readNumber() string {
	initialPosition := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[initialPosition:l.position]
}

// NOTE: This is where more special characters could be cased
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// similar to readChar except it doesnt increment the cursor
// we simply want to see what the next character is
// some languages may even require that you peek further ahead than
// just one position
func (l *Lexer) peekChar() byte {
	if l.readPosition > len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
