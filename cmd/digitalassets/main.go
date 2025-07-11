// cmd/digitalassets/main.go
package main

import (
"flag"
"log"
"os"

"digitalassets/internal/digitalassets"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := digitalassets.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
