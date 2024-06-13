# Building_an_Interpreter

# Interpreter Components

# Lexer (lexical analyzer)
    It is a program that tokenizes the input text and classifies these tokens into predefined categories.

    The following is the output from the lexer. 
    >> let a = fn(x, y) { return x * y;};
    {Type:LET Literal:let}
    {Type:IDENT Literal:a}
    {Type:= Literal:=}
    {Type:FUNCTION Literal:fn}
    {Type:( Literal:(}
    {Type:IDENT Literal:x}
    {Type:, Literal:,}
    {Type:IDENT Literal:y}
    {Type:) Literal:)}
    {Type:{ Literal:{}
    {Type:RETURN Literal:return}
    {Type:IDENT Literal:x}
    {Type:* Literal:*}
    {Type:IDENT Literal:y}
    {Type:; Literal:;}
    {Type:} Literal:}}
    {Type:; Literal:;}

# Parser (Syntax analyzer)
    The parser traverses through this data to create the AST(Abstract Syntax Tree).

    The following is the output from the parser.
    >> let x = 1 + 2 * 3 / 4 + add(y + 6, z * p);

    let x = ((1 + ((2 * 3) / 4)) + add((y + 6), (z * p)));

# Evaluator (Semantic Analyzer)
    It ensures that the declarations and statements of a program are semantically correct.

# REPL Loop
    
# Identifiers can also hold the value of an expression 
# Token types and functions that parse the token will be associated via a Pratt parser.

# Statements
    Statements do not produce values.

# Expressions
    Expressions produce values.

# Example
    - let x = 10; do not produce a value but 
    - answer(3, 6) produces a value.

# Statement Types

# LET statement
    let <identifier> = <expression>

# RETURN statement
    return <expression>

# Expression Statement

# Expression Types

# Identifiers 
    It can be a variable name, function name, or constant.

# IntegerLiterals 
    e.g. 5, 10 etc    

# Prefix Operators
    e.g. -5, !foobar

    <prefix operator><expression>;

# Infix Operators
    e.g. x / y, x == y

    <expression> <infix operator> <expression>

# Boolean Literals
    e.g. true, false

# Grouped Expressions (No separate AST node type required)
    e.g. (5 + 5) * 2;

# If-Else Expression
    if (<condition>) <consequence> else <alternative>

    Consequence and Alternative are block statements which are a series of statements (just like program in Monkey) enclosed by the opening { and closing } .

# functions literals
    fn <parameters> <block statements>
    <parameters> = (<parameter one>, <parameter two>, <parameter three>, ...)

# Call Expression
    function call
    e.g. add(4, 7);
    <expression>(<comma separated expresssions>)