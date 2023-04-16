# 泛型

命令

```bash

go clean --cache && rm geneticTest
go build -v -x -work

/home/rufeer/Codes/github/go1.18.7/pkg/tool/linux_amd64/compile -o $WORK/b001/_pkg_.a -trimpath "$WORK/b001=>" -p main -complete -buildid hUU3IW8BpJVVPKs8FTYG/hUU3IW8BpJVVPKs8FTYG -goversion go1.18.7 -c=4 -D _/home/rufeer/Codes/goprojects/geneticTest -importcfg $WORK/b001/importcfg -pack ./main.go
/home/rufeer/Codes/github/go1.18.7/pkg/tool/linux_amd64/link -o $WORK/b001/exe/a.out -importcfg $WORK/b001/importcfg.link -buildmode=exe -buildid=EAvmAhDUtqJZRpjyYSDv/hUU3IW8BpJVVPKs8FTYG/BpUJJ06Wivvm3JV_0iXD/EAvmAhDUtqJZRpjyYSDv -extld=gcc $WORK/b001/_pkg_.a

# _/home/rufeer/Codes/goprojects/geneticTest
>>> InstInfo for FuncTypeParam001[go.shape.int_0]
  Typeparam go.shape.int_0
>>> Done Instinfo

=== Creating dictionary .dict.FuncTypeParam001[int]
 * int
Main dictionary in main at generic function call: FuncTypeParam001 - (<node FUNCINST>)(1)

=== Finalizing dictionary .dict.FuncTypeParam001[int]
=== Finalized dictionary .dict.FuncTypeParam001[int]

```

符号

```bash
  1443: 000000000047e080   126 FUNC    GLOBAL DEFAULT    1 main.FuncTypeParam001[go.shape.int_0]
  2032: 00000000004b2ae0     8 OBJECT  GLOBAL DEFAULT    2 main..dict.FuncTypeParam001[int]

```

编译器源码分析

```go

infoPrintMode = true

buildInstantiations
	instantiateMethods               // instantiates all the methods (and associated dictionaries) of all fully-instantiated generic types that have been added to typecheck.instTypeList. It continues until no more types are added to typecheck.instTypeList.
	  getInstantiation
			getInstInfo
		getDictionarySym
	n := len(typecheck.Target.Decls) // functions/methods
	scanForGenCalls                  // Scan all currentdecls for call to generic functions/methods.
		visit
			getInstNameNode              // returns the name node for the method or function being instantiated, and a bool which is true if a method is being instantiated.
			getInstantiation             // gets the instantiantion and dictionary of the function or method nameNode with the type arguments shapes. If the instantiated function is not already cached, then it calls genericSubst to create the new instantiation.
				checkFetchBody             // checks if a generic body can be fetched, but hasn't been loaded yet. If so, it imports the body.
				typecheck.MakeFuncInstSym
			getDictOrSubdict
			transformCall
		instantiateMethods
	finalizeSyms

typecheck.Target.Decls  // []cmd/compile/internal/ir.Node   *cmd/compile/internal/ir.Func  Op: ODCLFUNC Type: 
scanForGenCalls // 为所有非泛型函数的泛型调用 进行实例化对应的泛型函数
getInstantiation // 将实例化表达式InstExpr转化为ir.Func节点
-------------------------------------------------------------------------------------------------------------------------------------------------------------------

// genInst has the information for creating needed instantiations and modifying functions to use instantiations.
type genInst struct

// buildInstantiations scans functions for generic function calls and methods, and
// creates the required instantiations. It also creates instantiated methods for all
// fully-instantiated generic types that have been encountered already or new ones
// that are encountered during the instantiation process. It scans all declarations
// in typecheck.Target.Decls first, before scanning any new instantiations created.
func (g *genInst) buildInstantiations()

0  0x0000000000bf376f in cmd/compile/internal/noder.(*genInst).buildInstantiations  at /home/rufeer/Codes/github/go1.18.7/src/cmd/compile/internal/noder/stencil.go:53
1  0x0000000000bc6b51 in cmd/compile/internal/noder.BuildInstantiations             at /home/rufeer/Codes/github/go1.18.7/src/cmd/compile/internal/noder/stencil.go:44
2  0x0000000000bc6b51 in cmd/compile/internal/noder.(*irgen).generate               at /home/rufeer/Codes/github/go1.18.7/src/cmd/compile/internal/noder/irgen.go:331
3  0x0000000000bc672d in cmd/compile/internal/noder.check2                          at /home/rufeer/Codes/github/go1.18.7/src/cmd/compile/internal/noder/irgen.go:92
4  0x0000000000bca835 in cmd/compile/internal/noder.LoadPackage                     at /home/rufeer/Codes/github/go1.18.7/src/cmd/compile/internal/noder/noder.go:90
5  0x0000000000c3b933 in cmd/compile/internal/gc.Main                               at /home/rufeer/Codes/github/go1.18.7/src/cmd/compile/internal/gc/main.go:191

```