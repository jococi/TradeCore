package main

import (
	"TradeCore/lib/signal"
)

func main() {
	signalCH := signal.InitSignal()
	signal.HandleSignal(signalCH)
}
