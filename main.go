package main

import(
    "os"
    "fmt"
    "log"
    "flag"
    "bufio"
    "strings"
    "net/http"
    "math/rand"    
)

type handler func(w http.ResponseWriter, r *http.Request)

func loadQuotes() []string {
    var quotes []string
    file, err := os.Open("quotes.txt")
    
    if err != nil {
        log.Fatalf("Cannot open quotes.txt: %v\n",err)
    }
    reader := bufio.NewReader(file)
    quote, err := reader.ReadString('\n')
    for err == nil {
        quotes = append(quotes, strings.Trim(quote, "\n "))
        quote,err = reader.ReadString('\n')
    }
    
    return quotes
}

func makeHandler(quotes []string) handler {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Access-Control-Allow-Origin", "*")
        fmt.Fprintf(w, "%s", quotes[rand.Intn(len(quotes))])
    }
}

func main() {
    addr := flag.String("addr", ":8080", "The address on which to listen")
    flag.Parse()
    
    quotes := loadQuotes()
    
    fmt.Printf("Loaded %d quotes\n", len(quotes))
    
    http.HandleFunc("/", makeHandler(quotes))
    http.ListenAndServe(*addr, nil)
}