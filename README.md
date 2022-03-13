# FlowScript

![FlowScript Logo](icon/FlowScriptIcon.jpeg)

A embeded script language interpreter written by Go.

## Usage

Use:

```go
import "https://github.com/zhangzheheng12345/flowscript"
```

and add:

```
require github.com/zhangzheheng12345/flowscript v0.1.1
replace github.com/zhangzheheng12345/flowscript v0.1.1 => github.com/zhangzheheng12345/FlowScript v0.1.1
```

in go.mod to import it to your Go project.

( *I don't know why there must be a replacement. Who could give me an answer?* )

Use:

```go
parser.RunCode( <your codes(type string)> )
```

to run FlowScript code.