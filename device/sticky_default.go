//go:build !linux

package device

import (
	"github.com/dragonbreath2000/wireguard-go/conn"
	"github.com/dragonbreath2000/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
