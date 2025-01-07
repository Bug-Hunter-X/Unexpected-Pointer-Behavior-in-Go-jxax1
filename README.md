# Unexpected Pointer Behavior in Go

This repository demonstrates a common pitfall when working with pointers in Go.  Many programmers coming from languages with different memory management models may encounter unexpected behavior. The example shows how changing a variable through a pointer can have seemingly unrelated effects.

## Bug
The `bug.go` file contains the problematic code. It shows how assigning a new value to a pointer does not modify the original variable but rather updates the value pointed to by the pointer.

## Solution
The `bugSolution.go` file offers a corrected understanding and usage of pointers, demonstrating how to properly use pointers to manipulate values effectively without unexpected side effects.