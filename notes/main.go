package main

import "fmt"

//go:noinline
func FuncTypeParam001[T any](t T) {
	fmt.Println(t)
}

//go:noinline
func main() {
	FuncTypeParam001(1) // ir.CallExpr
}

// ir.CurFunc == main ir.Func

// (dlv) p ir.CurFunc.Body
// cmd/compile/internal/ir.Nodes len: 1, cap: 1, [
//         *cmd/compile/internal/ir.CallExpr { // A CallExpr is a function call X(Args).
//                 miniExpr: (*"cmd/compile/internal/ir.miniExpr")(0xc000380750),
//                 origNode: (*"cmd/compile/internal/ir.origNode")(0xc000380788),
//                 X: cmd/compile/internal/ir.Node(*cmd/compile/internal/ir.InstExpr) ...,
//                 Args: cmd/compile/internal/ir.Nodes len: 1, cap: 1, [
//                         ...,
//                 ],
//                 KeepAlive: []*cmd/compile/internal/ir.Name len: 0, cap: 0, nil,
//                 IsDDD: false,
//                 NoInline: false,},
// ]
