# Let's Learn Go
## Variables

We can declare variable by typing 
```
<"var"> <"name"><"data-type"> 
```
Or we can declare and initialize at the same without specifying the data-type 
```
<"var"> <"name"> = <"value"> 
```
Also there is a short hand method as well which we can use for declaring and initializing a variable short hand by using ":=" only 
```
<"name"> := <"value">
```

We can initialize multiple variables as well by using all of the methods above
```
<"var"> <"name1">, <"name2"> <"optional data-type"> = <"value1">,<"value2">
```
OR 
```
<"name1">, <"name2"> := <"value1">, <"value2">
```
## Data Types

The Data-types in golang are simple not that difficult if you are coming from a different language
I am not talking about "javascript" XD
```
uint, uint8, uint16, uint32, uint64

int, int8, int16, int32, int64

float32, float64

rune

bool

string
```
The default value for these data-types is as follows
```
[Default Values of data-types]

uint, uint8, uint16, uint32, uint64 = 0

int, int8, int16, int32, int64 = 0

float32, float64 = 0

bool = false

string = ""
```
