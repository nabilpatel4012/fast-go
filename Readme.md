# Let's Learn Go
## Variables

We can declare variable by typing 
```go
var <"name"><"data-type"> 
```
Or we can declare and initialize at the same without specifying the data-type 
```go
var <"name"> = <"value"> 
```
Also there is a short hand method as well which we can use for declaring and initializing a variable short hand by using ":=" only 
```go
<"name"> := <"value">
```

We can initialize multiple variables as well by using all of the methods above
```go
var <"name1">, <"name2"> <"optional data-type"> = <"value1">,<"value2">
```
OR 
```go
<"name1">, <"name2"> := <"value1">, <"value2">
```
## Data Types

The Data-types in golang are simple not that difficult if you are coming from a different language
I am not talking about "javascript" XD
```go
uint, uint8, uint16, uint32, uint64

int, int8, int16, int32, int64

float32, float64

rune

bool

string
```
The default value for these data-types is as follows
```go
[Default Values of data-types]

uint, uint8, uint16, uint32, uint64 = 0

int, int8, int16, int32, int64 = 0

float32, float64 = 0

bool = false

string = ""
```
**NOTE:**
Constants cannot be declared without any value, that is we have to initialize it with a value
```go
const myConst int               <-(This will give an error)

// Correct way of writing
const myNewConst int = 100      <- Will execute without any problem
```
## Functions
Functions plays very important role in every programming language. 
They help us avoid repetative code. 
Functions can be considered as lego blocks where we can join the blocks to make something useful and meaningful out of it.

Let's take a look at the syntax of Function
```go
func <"Function name">(<"input argument">) <"return-type"> {
    <"return statement">
}
```
**NOTE:** 
The "input argument", "return-type" and "return statement" are optional and they depend on the requirement and what we want

Below are some examples
```go
func printMe() {
  fmt.Println("I am printed!!!")
}

func myFunctionWithInput(value string) {
  fmt.Println(value)
}

func myFunctionWithReturnType(value int) int {
  var square = value*value
  return square
}

func myFunctionWithMultipleReturnType(value int, name string) (int, string) {
  var square = value*value
  myMessage := "Hello Nabil"
  return square,myMessage
}
```
