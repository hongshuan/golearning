package main

import (
    "fmt"
    "strings"
)

type CodeLines struct {
    lines []string
}

func NewCodeLines() *CodeLines {
    return &CodeLines{make([]string, 0)}
}

func (c *CodeLines) Push(s string) {
    c.lines = append(c.lines, s)
}

func (c *CodeLines) ToString() string {
    return strings.Join(c.lines, "\n")
}

func main() {
    strfmt := fmt.Sprintf
    lines := NewCodeLines()
    lines.Push("<Request>")
    lines.Push("<Item>")
    lines.Push(strfmt("<Weight>%d</Weight>", 23))
    lines.Push("</Item>")
    lines.Push("</Request>")
    fmt.Println(lines.ToString())
}
