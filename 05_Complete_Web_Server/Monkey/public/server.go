package main

import (
	"Monkey/evaluator"
	"Monkey/lexer"
	"Monkey/object"
	"Monkey/parser"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", serveHTML)
	r.HandleFunc("/run", runREPL).Methods("POST")

	port := getenv("PORT", "8080")
	fmt.Printf("Server is listening on port %s...\n", port)
	http.ListenAndServe(":"+port, r)
}

func getenv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func runREPL(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Code string `json:"code"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	env := object.NewEnvironment()
	evaluated, errors := evaluate(req.Code, env)
	if len(errors) != 0 {
		printParserErrors(w, errors)
		return
	}

	if evaluated != nil && evaluated.Inspect() != "null" {
		io.WriteString(w, evaluated.Inspect())
	} else {
		io.WriteString(w, "null")
	}
}

func evaluate(input string, env *object.Environment) (object.Object, []string) {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		return nil, p.Errors()
	}

	evaluated := evaluator.Eval(program, env)
	return evaluated, nil
}

const MONKEY_FACE = `
           __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some trouble here!\n\n")
	io.WriteString(out, "👾👾👾👾👾👾👾👾👾👾👾\n\n")
	io.WriteString(out, "Parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
