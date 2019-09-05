package lexer

import "github.com/palash25/chimp/token"

// Lexer represents the lexer used to tokenize our input programs
type Lexer struct {
	// TODO: Add filename, col and line no and read file using io.Reader
	input        string
	position     int
	readPosition int
	// TODO: change to rune later to accomadate unicode characters
	ch byte
}

// New creates a Lexer instance
func New(s string) *Lexer {
	l := &Lexer{input: s}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() *token.Token {
	var tok *token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = &token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = &token.Token{Type: token.ASSIGN, Literal: string(l.ch)}
		}
	case '+':
		tok = &token.Token{Type: token.PLUS, Literal: string(l.ch)}
	case '-':
		tok = &token.Token{Type: token.MINUS, Literal: string(l.ch)}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = &token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = &token.Token{Type: token.BANG, Literal: string(l.ch)}
		}
	case '/':
		tok = &token.Token{Type: token.SLASH, Literal: string(l.ch)}
	case '*':
		tok = &token.Token{Type: token.ASTERISK, Literal: string(l.ch)}
	case '<':
		tok = &token.Token{Type: token.LT, Literal: string(l.ch)}
	case '>':
		tok = &token.Token{Type: token.GT, Literal: string(l.ch)}
	case ',':
		tok = &token.Token{Type: token.COMMA, Literal: string(l.ch)}
	case ';':
		tok = &token.Token{Type: token.SEMICOLON, Literal: string(l.ch)}
	case '(':
		tok = &token.Token{Type: token.LPAREN, Literal: string(l.ch)}
	case ')':
		tok = &token.Token{Type: token.RPAREN, Literal: string(l.ch)}
	case '{':
		tok = &token.Token{Type: token.LBRACE, Literal: string(l.ch)}
	case '}':
		tok = &token.Token{Type: token.RBRACE, Literal: string(l.ch)}
	case 0:
		tok = &token.Token{Type: token.EOF, Literal: ""}
	default:
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			tokenType := token.LookupIdent(literal)
			return &token.Token{Type: tokenType, Literal: literal}
		} else if isDigit(l.ch) {
			literal := l.readNumber()
			return &token.Token{Type: token.INT, Literal: literal}
		} else {
			tok = &token.Token{Type: token.ILLEGAL, Literal: "ILLEGAL"}
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	// allow underscore in identifiers
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readIdentifier() string {
	pos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
