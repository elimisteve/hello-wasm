package main

import (
    "fmt"

    "github.com/bytecodealliance/wasmtime-go"
)

func main() {
    store := wasmtime.NewStore(wasmtime.NewEngine())
    module, err := wasmtime.NewModuleFromFile(store, "gcd.wat")
    check(err)
    instance, err := wasmtime.NewInstance(module, []*wasmtime.Extern{})
    check(err)

    gcd := instance.GetExport("gcd").Func()
    val, err := gcd.Call(6, 27)
    check(err)
    fmt.Printf("gcd(6, 27) = %d\n", val.(int32))
}

func check(err error) {
    if err != nil {
        panic(err)
    }
}
