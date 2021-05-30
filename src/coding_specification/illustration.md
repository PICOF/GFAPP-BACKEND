# illustration
**核心语言：go**

**使用框架：gin**

**数据库：MySQL（可使用gorm）**

## go语言编写规范

### 环境变量配置：

#### （1）**GOROOT**:

Go 语言安装根目录的路径，也就是 GO语言的**安装路径**

#### （2）**GOPATH**：

若干工作区目录的路径，也就是我们**自己定义的工作空间**

- 放置内容：源码文件（source file）；归档文件(.a archive file)；可执行文件（.exe executable file)

#### （3）**GOBIN**：

GO程序生成的**可执行文件**的路径



###  工作区：


（1）src：存放源码文件(或项目)

（2）pkg： 存放归档文件

（3）bin：存放可执行文件

### 二.命名规范：

1）命名时尽可能借名知意，建议采用英文单词或缩写，但**不能用汉语拼音和其缩写**

2）go语言和C语言一样，是大小写敏感的语言

tip： a）当命名（包括常量、变量、类型、函数名、结构字段等）以大写

​          字母开头，那么此对象可以被外部包的代码使用（客户端需先导出

​          这个包）（像面向对象语言中的public）

​          b）命名以小写开头，则对包外不可见，但是其在整个包内可见并可

​          用（像面向对象语言中的private）

#### 1.包命名：package

小写，简短，不与标准库重复

```go
package demo

package main
```

#### 2.文件命名：

小写，有意义，用**下划线**分隔单词

```go
my_test.go
```

#### 3.结构体命名

1）采用驼峰命名法，首字母根据访问

tip：驼峰命名法,就是多个单词拼接起来时，单词衔接处首字母要大写

​        同时，根据你想让它对外是否可见确认其首字母是否大写

​        例如

```go
printEmployeePaychecks();//内部函数
DataBaseUser struct{};//外部可调用的结构体
```

2）struct申明和初始化格式采用多行，例如

```go
//多行申明
type User struct{
    UserName string
    Email string
}
//多行初始化
u:= User{
    UserName: "Polaris6G",
    Email: "860797149@qq.com"
}
```

#### 4.接口命名：

与结构体命名规则相同

只有单个函数的接口命名可以以“er”为后缀，如Reader，Writer

```go
type Reader interface{
    Read(p []byte)(n int, err error)
}
```

#### 5.变量命名

1）与结构体命名规则相同

2）遇到特有名词时，需要遵循以下规则：

​      a）如果变量为私有，且特有名词为首个单词，则使用小写,如apiClient

​      b）其他情况都应当使用该名词原有的写法，如APIClient、UserID

​      c）错误示例：UrlArray,应该写成urlArray或URLArray

3）若变量类型为bool类型，则名称应为以Has，is，Can或Allow开头

```go
var isExist bool
var hasConflict bool
var canManage bool
var allowGitbook bool
```

6.常量命名

均需使用大写字母，并用下划线分词

```go
const APP_VER = "1.0"
```

### 三、注释

1）尽量使用//单行注释，少使用块注释/* */

2）在命名新变量，函数，结构或接口时写明注释，表明用途，方便其他开发人员理解和纠错

3）对于一些关键位置的代码逻辑，或者局部较为复杂的逻辑，应有相应的逻辑说明

### 三.代码风格

#### 1.缩进和折行

1）注意产出代码的格式，格式化以后再git push给远端仓库

tip： a）使用goland开发，则可使用快捷键：ctrl+alt+L 进行格式化

​          c）使用vscode开发，则可使用快捷键：ctrl+shift+F 进行格式化

#### 2.语句结尾

go语言句末可以不加";"，所以规定就不加了。

当然，如果一行有多个语句，那必须使用";"分隔

#### 3.括号和空格

这里使用格式化就可以，然而注意一个问题：

```go
//正确的方式
if a > 0 {

}
//错误的方式
if a>0//a,>和0之间应该空格
{//左大括号不可以换行，会语法报错

}
```

#### 4. import规范

如果你的包引入了三种类型的包，标准库包，程序内部包，第三方包，采用如下方式进行组织

```go
import{
    "fmt"
    "strings"
    "net/http"//标准库包
    
    "myproject/models"
    "myproject/controller"
    "myproject/utils"//程序内部包
    
    "github.com/gin-gonic/gin"
    "github.com/go-sql-driver/mysql"//第三方包
}
```

在项目中不要使用相对路径引入包：

```go
//这是不好的导入
import (
    "../net"
)
```

#### 5.错误处理：

1）错误处理的原则是不能丢弃任何与返回err的调用，不要使用_丢弃，必须全部处理。接收到错误，要么返回err，要么用log记录下来

2）今早return，一旦有错误发生，马上返回

3）尽量不要使用panic，除非你知道自己在做什么

4）采用独立的错误流进行处理

例如：

```go
//错误写法
if err != nil{
    //err handling
}else{
    //normal code
}
//正确写法
if err !=nil{
    //error handling
    return//or continue,etc.
}
//normal code
```

#### 6.测试

单元测试文件名命名规范为example_test.go

测试用例的函数名称必须以Test开头，例如：TestExample

每个重要的函数都要首先编写测试用例，测试用例和正规代码一起提交方便进行回归测试

