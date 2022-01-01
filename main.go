package main

import(
  "flag"
  "fmt"
)

func main(){
  var addr string
  var port int64
  var buff int64
  
  addr := flag.StringVar(&addr,"--address","192.168.88.251","Server listen address/host.")
  port := flag.Int(&port,"--port",5566,"Server listen port number.")
  buff := flag.Int(&buff,"--buffor",1024,"Server buffor size.")
  
  flag.Parse()
  
  fmt.Println("Addr: ", addr)
  fmt.Println("Port: ", port)
  fmt.Println("Buff: ", buff)
}
