package main

type Seconds struct {
	MsgType byte
	Second  [5]byte
}

type Milliseconds struct {
	MsgType     byte
	Millisecond [3]byte
}

type SystemEvent struct {
	MsgType   byte
	EventCode byte
}

type MarketSegmentState struct {
	MsgType         byte
	MarketSegmentID [3]byte
	EventCode       byte
}

type OrderBookDirectory struct {
	MsgType          byte
	OrderBook        [6]byte
	Symbol           [16]byte
	ISIN             [12]byte
	FinancialProduct [3]byte
	TradingCurrency  [3]byte
	MIC              [4]byte
	MarketSegmentID  [3]byte
	NoteCodes        [8]byte
	RoundLotSize     [9]byte
}

type OrderBookTradingAction struct {
	MsgType      byte
	OrderBook    [6]byte
	TradingState byte
	Reserved     byte
	Reason       [4]byte
}

type AddOrder struct {
	MsgType              byte
	OrderReferenceNumber [9]byte
	BuySellIndicator     byte
	Quantity             [9]byte
	OrderBook            [6]byte
	Price                [10]byte
}

type AddOrderMPID struct {
	MsgType              byte
	OrderReferenceNumber [9]byte
	BuySellIndicator     byte
	Quantity             [9]byte
	OrderBook            [6]byte
	Price                [10]byte
	Attribution          [4]byte
}

type OrderExecuted struct {
	MsgType                   byte
	OrderReferenceNumber      [9]byte
	ExecutedQuantity          [9]byte
	MatchNumber               [9]byte
	OwnerParticipantID        [4]byte
	CounterpartyParticipantID [4]byte
}

type OrderExecutedWithPrice struct {
	MsgType                   byte
	OrderReferenceNumber      [9]byte
	ExecutedQuantity          [9]byte
	MatchNumber               [9]byte
	Printable                 byte
	TradePrice                [10]byte
	OwnerParticipantID        [4]byte
	CounterpartyParticipantID [4]byte
}

type OrderCancel struct {
	MsgType              byte
	OrderReferenceNumber [9]byte
	CanceledQuantity     [9]byte
}

type OrderDelete struct {
	MsgType              byte
	OrderReferenceNumber [9]byte
}

type Trade struct {
	MsgType                   byte
	OrderReferenceNumber      [9]byte
	TradeType                 byte
	Quantity                  [9]byte
	OrderBook                 [6]byte
	MatchNumber               [9]byte
	TradePrice                [10]byte
	OwnerParticipantID        [4]byte
	CounterpartyParticipantID [4]byte
}

type CrossTrade struct {
	MsgType        byte
	Quantity       [9]byte
	OrderBook      [6]byte
	CrossPrice     [10]byte
	MatchNumber    [9]byte
	CrossType      byte
	NumberOfTrades [10]byte
}

type BrokenTrade struct {
	MsgType     byte
	MatchNumber [9]byte
}

type NOII struct {
	MsgType            byte
	PairedQuantity     [9]byte
	ImbalanceQuantity  [9]byte
	ImbalanceDirection byte
	OrderBook          [6]byte
	EquilibriumPrice   [10]byte
	CrossType          byte
	BestBidPrice       [10]byte
	BestBidQuantity    [9]byte
	BestAskPrice       [10]byte
	BestAskQuantity    [9]byte
}

func ItchUatoi(buf []byte, len int) uint64 {
	var ret uint64 = 0
	var idx int = 0
	for idx < len {
		if buf[idx] != ' ' {
			break
		}
		idx++
	}
	for idx < len {
		ch := buf[idx]
		ret = (ret * 10) + uint64(ch-byte('0'))
		idx += 1
	}
	return ret
}
