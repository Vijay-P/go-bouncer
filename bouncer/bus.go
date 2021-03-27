package bouncer

import (
	"github.com/wagoodman/go-bouncer/internal/bus"
	"github.com/wagoodman/go-partybus"
)

func SetBus(b *partybus.Bus) {
	bus.SetPublisher(b)
}
