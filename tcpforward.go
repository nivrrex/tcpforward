package main

import (
    "fmt"
    "io"
    "net"
    "os"
    "flag"
)

func main() {
    var local string
    var remote string

    flag.StringVar(&local, "l", "", "local address - address:port")
    flag.StringVar(&remote, "r", "", "remote address - address:port")
    flag.Parse()

    l, err := net.Listen("tcp", local)
    if err != nil {
        fmt.Println(err, err.Error())
        os.Exit(0)
    }

    for {
        s_conn, err := l.Accept()
        if err != nil {
            continue
        }

        d_tcpAddr, _ := net.ResolveTCPAddr("tcp4", remote)
        d_conn, err := net.DialTCP("tcp", nil, d_tcpAddr)
        if err != nil {
            fmt.Println(err)
            s_conn.Write([]byte(fmt.Sprintf("can't connect ",remote)))
            s_conn.Close()
            continue
        }
        go io.Copy(s_conn, d_conn)
        go io.Copy(d_conn, s_conn)
    }
}
