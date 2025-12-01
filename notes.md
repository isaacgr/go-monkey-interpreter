---
title: Writing an Interpreter in Go
created_at: 2025-11-22 18:51:35
updated_at: 2025-11-22 18:51:35
description:  Notes from "Writing an Interpreter in Go" by Thorsten Ball
tags: golang, go, interpreter, compiler, development
---

# Introduction

Interpreters are magical.

They are programs that take something as their input and produce something
useful.

Seemingly random characters are fed into the interpreter and suddenly become
**meaningful**. It makes sense out of nonsense and the computer now understands
and acts upon this language we feed it.

Understanding how interpreters work means understanding how lexers and parsers
work as well.

What can be said to be common among interpreters is that they take source code
and evaluate it without producing some immedate result that can later be
executed. This is in contrast to compilers, which take source code and produce
output in another language that the underlying system can understand.

Some interpreters are small and do not even bother with a parsing step,
they just interpret the input right away. See `Brainfuck` interpreters.

Others are elaborate and highly optimized, using advanced parsing and evaluation
tecnniques. Some of them dont just evaluate input, but compile it into bytecode
and then evaluate that. Even more advanced are `JIT` interpreters, which compile
the input just-in-time to native machine code that then gets executed.

In between those categories are interpreters that parse the source code, build
an abstract syntax tree (AST) and then evalutate this tree. This type of
interpreter is called a `tree-walking` interpreter. This is what we will be 
building. It will include our own lexer, parser, tree representation and
evaluator.

## The monkey programming language

Every interpreter is built to interpret a specific languate. This is how you
'implement' a language. Without a compiler or an interpreter a language is
nothing more than an idea or specification.

We will be building our own languate called **Monkey**.

Monkey will have the following features:

- C-like syntax
- variable bindings
- integers and booleans
- arithmetic expressions
- built in functions
- first class and higher-order functions
- closures
- a string data structure
- an array data structure
- a hash data structure

Ex.

```c
let age = 1;
let name = "Monkey";
let result = 10 * (20/2);

let myArray = [1,2,3,4,5];
let hashtable = {"name": "Isaac"};

myArray[1]; // 1
hashtable["name"]; // "Isaac"


// let functions can also be used to bind functions to names

let add = fn(a, b) { return a + b; };

// Implicit return is also supported

let add = fn(a, b) { a + b; };

add(1, 2);


// Monkey also supports higher order functions. These are functions that take
// other functions as arguments

// First class functions mean that its just another value like integers or
// strings

let twice = fn(f, x){
    return f(f(x));
};

let addTwo = fn(x){
    return x + 2;
};

twice(addTwo, 2); // 6
```

The interpreter will implement all of these.

It will tokenize and parse monkey code in a **REPL**, building up an internal
representation of the code in an AST and then evaluate that tree.

The major parts will include:

- the lexer
- the parser
- the AST
- the internal object system
- the evaluator

# Chapter 1 - Lexing

## 1.1 Lexical Analysis

- we need to turn source code into a more accessible form in order to work
with it

- we're going to change the representation of our source code 2 times before
we evaluate it

SOURCE CODE ---> TOKENS ---> ABSTRACT SYNTAX TREE

- the first part transformation from source code to tokens is called *lexical analysis*
    - it is performed by a lexer, also called a tokenizer or a scanner

- tokens are small data structures that are fed into the parser which does
the second part of the transformation and turns the tokens into an
*abstract syntax tree*

Ex. The below would be input to a parser

```c
let x = 5 + 5;
```

And the lexer would output 

```
[
    LET,
    IDENTIFIER("x"),
    EQUAL_SIGN,
    INTEGER(5),
    PLUS_SIGN,
    INTEGER(5),
    SEMICOLON
]
```

- these tokens would have the original source code implementation attached
- what constitues a token varies between different lexer implementations
- some implementations care about whitespace, however monkey does not
    - languages like python this would matter
- a production ready lexer might also attach the line number and column and 
file name of a token, to be used when tracking errors

## 1.1 Defining our tokens

- the first thing we must do is to define the tokens our lexer is going to
output

- the subset of the monkey language we're going to lex looks like this

```c
let five = 5;
let ten = 10;

let add = fn(x, y){
    x+y;
};

let result = add(five, ten);
```

- this contains several types:

    - numbers
    - variable names (identifiers)
    - keywords
    - special characters

What fields does our token data structure need?

We definitely need a "type" attribute, so we can distinguish between integers
and right brackets. It also needs a field to hold the literal value of the
token, so we can reuse it later.

We will define the possible `TokenTypes` as constants

> See token/token.go

## 1.3 The Lexer

The lexer will take source code as input and output the tokens that represent
the source code. It will go through its input and output the next token it
recognizes. It doesnt need to buffer or save tokens, since only one method,
`NextToken()` will be called which will output the next token.

- we'll initialize the lexer with our source code and repeatedly call `NextToken()`
on it to go through the source code.
    - we are making life easy here and using a `String` to represent our
    source code
    - as a note, it would be better to attach line numbers and columns for debugging
    in production code, so it would be better to initialize the lexer with an
    `io.Reader` and the filename.

> See lexer/lexer_test.go and lexer/lexer.go

# References 

- "Writing an Interpreter in Go" by Thorsten Ball
