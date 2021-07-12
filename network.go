package chord

import (
	"crypto/sha1"
	"net"
	"net/http"
	"net/rpc"
)

func createNewNode(address string) (*Node, error) {
	h := sha1.New()
	h.Write([]byte(address))

	node := &Node{
		id:             h.Sum(nil),
		address:        address,
		predecessorId:  nil,
		predecessorRPC: nil,
		fingerTable:    nil,
	}

	rpc.Register(node)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", address)
	if err != nil {
		return nil, LISTEN_ERROR
	}

	go http.Serve(l, nil)

	client, err := rpc.DialHTTP("tpc", address)
	if err != nil {
		return nil, DIALING_ERROR
	}

	node.self = client

	return node, nil
}