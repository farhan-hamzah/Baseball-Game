package main
import (
    "fmt"
    "strconv"
)

func main() {
    ops := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
    fmt.Println(calPoints(ops)) // Output: 27
}

func calPoints(operations []string) int {
    var stack []int

    for _, op := range operations {
        switch op {
        case "+":
            n := len(stack)
            stack = append(stack, stack[n-1]+stack[n-2])
        case "D":
            stack = append(stack, 2*stack[len(stack)-1])
        case "C":
            stack = stack[:len(stack)-1]
        default:
            // konversi string ke integer
            num, _ := strconv.Atoi(op)
            stack = append(stack, num)
        }
    }

    // hitung total skor
    total := 0
    for _, score := range stack {
        total += score
    }

    return total
}
