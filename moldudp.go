package main

type MoldUDPHeader struct {
	Session        [10]byte
	SequenceNumber uint32
	MessageCount   uint16
}

type MoldUDPMessageBlock struct {
	MessageLength uint16
}
