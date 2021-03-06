package polish

import (
	"net"
)

type Connection interface {
	Handshake(conn net.Conn) error
	Polish(input []byte) ([]byte, error)
	Unpolish(input []byte) ([]byte, error)
	GetChunkSize() int
}

type Server interface {
	NewConnection(net.Conn) Connection
}
