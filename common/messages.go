package common

import (
	"net"
)

// Message types
const (
	_ = iota
	MSGJOINREQ
	MSGJOINRSP
	MSGFRAME
)

// sent before each message to help identify message type
type MessageType struct {
	Type uint8 //256 kinds should be enough for all messages, huh?
}

// sent from client to master, representing request to join
type JoinReq struct {
	Identity int
	MACAddr  net.HardwareAddr
}

// sent from master back to client, indicating assigned IP address and Mask
type JoinRsp struct {
	Address net.IP
	Mask    net.IPMask
	Success bool
}

// represent a MAC frame
type Frame []byte
