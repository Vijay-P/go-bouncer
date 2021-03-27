package event

import "github.com/wagoodman/go-partybus"

const (
	ModuleScanStarted partybus.EventType = "bouncer-module-scan-started"
	ModuleScanResult  partybus.EventType = "bouncer-module-scan-result"
)
