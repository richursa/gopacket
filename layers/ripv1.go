package layers

import (
	"encoding/binary"
	"net"

	"github.com/google/gopacket"
)

type RIPv1 struct {
	BaseLayer
	Command    uint8
	Version    uint8
	RIPAddress []RIPAddress
}
type RIPAddress struct {
	AddressFamily uint16
	IPAddress     net.IP
	Metric        uint32
}

func (r *RIPv1) LayerType() gopacket.LayerType { return LayerTypeRIPv1 }
func decodeRIPv1(data []byte, p gopacket.PacketBuilder) error {
	return nil
}
func (m *RIPv1) SerializeTo(b gopacket.SerializeBuffer, opts gopacket.SerializeOptions) error {
	bytes, err := b.PrependBytes(2 + len(m.RIPAddress)*10)
	if err != nil {
		return err
	}
	encoded := uint16(0)
	encoded |= uint16(m.Command) << 8
	encoded |= uint16(m.Version)
	binary.BigEndian.PutUint16(bytes[:2], encoded)
	for _, ripAddr := range m.RIPAddress {
		bytes = bytes[2:]
		binary.BigEndian.PutUint16(bytes[:2], ripAddr.AddressFamily)
		copy(bytes[2:], ripAddr.IPAddress)
		binary.BigEndian.PutUint32(bytes[6:], ripAddr.Metric)
	}
	return nil
}
