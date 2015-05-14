// MoldUDP transport protocol
//
// This implementation is based on the following specifications provided by
// NASDAQ OMX:
//
//  MoldUDP
//  Version 1.02a
//  October 19, 2006
//
// and
//
//  MoldUDP for NASDAQ OMX Nordic
//  Version 1.0.1
//  February 10, 2014

package main

type MoldUDPHeader struct {
	Session        [10]byte
	SequenceNumber uint32
	MessageCount   uint16
}

type MoldUDPMessageBlock struct {
	MessageLength uint16
}
