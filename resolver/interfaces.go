package resolver

import (
	"github.com/ninech/nine-dhcp2/dhcp/v4"
)

type Offerer interface {
	OfferV4ByMAC(clientInfo *v4.ClientInfoV4, xid, mac string) error
	OfferV4ByID(clientInfo *v4.ClientInfoV4, xid, duid, iaid string) error
}

type Acknowledger interface {
	AcknowledgeV4ByMAC(clientInfo *v4.ClientInfoV4, xid, mac, ip string) error
	AcknowledgeV4ByID(clientInfo *v4.ClientInfoV4, xid, duid, iaid, ip string) error
}

type Decliner interface {
	DeclineV4ByMAC(xid, mac, ip string) error
	DeclineV4ByID(xid, duid, iaid, ip string) error
}

type Releaser interface {
	ReleaseV4ByMAC(xid, mac, ip string) error
	ReleaseV4ByID(xid, duid, iaid, ip string) error
}

type Resolver interface {
	Offerer
	Acknowledger
	Releaser
	Decliner
}