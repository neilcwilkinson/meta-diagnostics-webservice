package main

import (
     "fmt"
     "log"
     "net/http"
     "io/ioutil"
     "os"
     "time"
)

var hostname = ""
var filename = "commandfile.cmd"

func Me(w http.ResponseWriter, r *http.Request) {         
     if r.Method == "POST" {
          body, _ := ioutil.ReadAll(r.Body)

          //timestamp;hostname;svc_description;return code;plugin output
          text := fmt.Sprintf("%d;%s;%s;%d;%s\n", time.Now().UnixNano(), hostname, "Meta Trader Diagnostics Check", 0, string(body)) 
          fmt.Printf(text)
          appendToFile(text)
     }
} 

func appendToFile(text string) {
     f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
     if err != nil {
          fmt.Printf("Failed to open file:", err)
     } else {
          defer f.Close()
          if _, err = f.WriteString(text); err != nil {
               fmt.Printf("Failed to write to file:", err, "\n")
          }
     }
}
func main() {
     hostname, _ = os.Hostname()
     fmt.Println(hostname)          

     http.HandleFunc("/MetaDiagnostics/", Me)
     fmt.Println("Started")
     //http.ListenAndServe(":80", nil)
     if err := http.ListenAndServe(":8080", nil); err != nil {
          log.Fatal("ListenAndServe: ", err)
     }
}