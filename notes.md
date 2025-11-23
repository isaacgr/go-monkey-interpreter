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


# References 

- "Writing an Interpreter in Go" by Thorsten Ball
