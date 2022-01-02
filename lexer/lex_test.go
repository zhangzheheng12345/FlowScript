package lexer

import (
    "testing"
    "fmt"
    "strings"
)

func TestLex(t *testing.T) {
    const t1 = 
        "add 1 var1 > var var2 \n"
    var res1 = []Token{
        Token{ADD, ""},
        Token{NUM, "1"},
        Token{SYMBOL, "var1"},
        Token{SEND, ""},
        Token{VAR, ""},
        Token{SYMBOL, "var2"},
        Token{STOP,""},
    }
    result := Lex(strings.Split(t1,""))
    for key, value := range result {
        if res1[key].Type() != value.Type() ||
           res1[key].Value() != value.Value() {
            fmt.Println("expecting: ",res1[key])
            fmt.Println("but got: ",value)
             fmt.Println("token: ", key)
        }
    }
}