package ibapi

import (
	"time"
)

type MsgInHistDataItem struct {
	Date string
	Open float64
	High float64
	Low float64
	Close float64
	Volume int64
	WAP float64
	HasGaps string
	BarCount int64 `minVer:"3"`
}

type MsgInHistData struct {
	ReqId int64
	StartDate string
	EndDate string
	Items []MsgInHistDataItem
}

type MsgOutReqHistData struct {
	ReqId int64
	Symbol string
	SecurityType string
	Expiry string
	Strike float64
	Right string
	Multiplier string
	Exchange string
	PrimaryExchange string
	Currency string
	LocalSymbol string
	IncludeExpired string `minVer:"31"`
	EndDateTime time.Time `minVer:"20"`
	BarSizeSetting string `minVer:"20"`
	Duration string
	UseRTH bool
	WhatToShow string
	FormatDate int64 `minVer:"16"`
}

const (
	FormatDateString = 1
	FormatDateSeconds = 2
	WhatToShowTrades = "TRADES"
	WhatToShowMidpoint = "MIDPOINT"
)
