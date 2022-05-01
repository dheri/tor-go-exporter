package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/reiver/go-telnet"
)

func main() {
	fmt.Println("Connecting to telnet.")

	conn, _ := telnet.DialTo("127.0.0.1:42457")

	conn.Write([]byte("authenticate \" \""))
	conn.Write([]byte("\n"))
	fmt.Println("Connected")

	b := make([]byte, 6)
	_, err := conn.Read(b)
	fmt.Println(string(b))

	if err != nil {
		// handle error
		fmt.Println("err.")
		panic(err)
	}

	conn.Write([]byte("setevents  circ stream orconn addrmap status_general status_client guard"))
	conn.Write([]byte("\n"))

	fmt.Println(ReaderTelnet(conn, "19a3a590701e391bd3f4bd4ea6c7e3b6964da5c763dbd201739358312ffe760a"))

}

// stop function if 'expect' is encountered
func ReaderTelnet(telnet_conn *telnet.Conn, expect string) (out string) {
	var buffer [1]byte
	recvData := buffer[:]
	var n int
	var telnet_err error

	tcp_conn, tcp_err := net.Dial("tcp", "localhost:9000")

	// l := logstash.New("127.0.0.1", 9000, 25)
	// _, logstash_err := l.Connect()

	if tcp_err != nil {
		panic(tcp_err)
	} else {
		fmt.Println("logstash connection made sucessfully")
	}

	for {
		n, telnet_err = telnet_conn.Read(recvData)
		// fmt.Println("Bytes: ", n, "Data: ", recvData, string(recvData))
		if n <= 0 || telnet_err != nil || strings.Contains(out, expect) {
			break
		} else {
			out += string(recvData)
		}
		if recvData[0] == 10 {
			fmt.Print(out)

			_, tcp_err := tcp_conn.Write([]byte(out))

			// logstash_err = l.Writeln(out)
			if tcp_err != nil {
				panic(tcp_err)
			}

			out = ""
		}
	}

	return out
}
