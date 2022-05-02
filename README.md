# ![FlowScript!](icon/flowscript.png)

A embeded script language interpreter written by Go. The language (FlowScript) is functional.

## Usage

Use:

```go
import (
    parser "https://github.com/zhangzheheng12345/flowscript"
)
```

and add:

```go.mod
require github.com/zhangzheheng12345/flowscript v0.2.0
replace github.com/zhangzheheng12345/flowscript v0.2.0 => github.com/zhangzheheng12345/FlowScript v0.2.0
```

in go.mod to import it to your Go project.

( *I don't know why there must be a replacement on my machine. Who could give me an answer?* )

Use:

```go
parser.RunBlock( <your codes(type string)> )
```

to run FlowScript code in a independent scope which will not influence the outside context.

Use:

```go
parser.RunModule( <your codes(type string)> )
```

to run FlowScript code as a module. All the variables declared outside functions & and blocks will be wrapped in a structure and add to global scope.

Use:

```go
parser.AddGoFunc( <function name(type string)>, <function(type func([]interface{}) interface{})>, <number of args(int, -1 -> limitless args)>)
```

to add native Go functions to global scope.

## Examples

### Hello World

<!--- highlight FlowScript as Ruby --->
```ruby
echoln "Hello World!" # output: Hello World!
```

### Fibonacci sequence

```ruby
def fibonacci a b begin
    add a b > echoln -
    add a b > smlr - 10000 > if - begin
        add a b > fibonacci b -
    end
end
echoln "fibonacci begins ..."
fibonacci 1 1
```

### OOP Mock (yes, in a functional language)

```ruby
def Student name li begin # This function will be used as a constructor
    echoln "constructor called!"
    def sum li begin
        len li > equal - 0 > if - begin
            expr 0
        end else begin
            len li > slice li 1 - > sum - > index li 0 > add - -
        end
    end
    sum li > len li > div - - > var avg -
    struct begin
        def getName begin expr name end # Returns the name of the student
        def getScore begin expr li end # Returns a list which contains the student's score
        def addScore score begin
            app li score > Student name - # Add a score to the list and returns a new object (FP)
        end
        def avgScore begin expr avg end
    end
end
list 95 96 97 > Student "Zhang" - > var A -
echo "name: "; A.getName > echoln -
echo "score: " ; A.getScore > echoln -
echo "avg: "; A.avgScore > echoln -
echoln "Now the score is updated."
A.addScore 100 > var B -
echo "score: "; B.getScore > echoln -
echo "avg: "; B.avgScore > echoln -
```
