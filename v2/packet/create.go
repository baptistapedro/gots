/*
MIT License

Copyright 2016 Comcast Cable Communications Management, LLC

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package packet

import (
	"github.com/Comcast/gots/v2"
)

var (
	required = []func(*Packet){setSyncByte}
)

var (
	// TestPatPacket is a minimal PAT packet for testing. It contains a single program stream with no payload.
	TestPatPacket = Packet{
		0x47, 0x40, 0x00, 0x10, 0x00, 0x00, 0xb0, 0x0d, 0x00, 0x01, 0xcb, 0x00,
		0x00, 0x00, 0x01, 0xe0, 0x64, 0x68, 0xd6, 0x84, 0x2e, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	TestPmtPacket = Packet{
		0x47, 0x40, 0x64, 0x10, 0x00, 0x02, 0xb0, 0x2d, 0x00, 0x01, 0xcb, 0x00,
		0x00, 0xe0, 0x65, 0xf0, 0x06, 0x05, 0x04, 0x43, 0x55, 0x45, 0x49, 0x1b,
		0xe0, 0x65, 0xf0, 0x05, 0x0e, 0x03, 0x00, 0x04, 0xb0, 0x0f, 0xe0, 0x66,
		0xf0, 0x06, 0x0a, 0x04, 0x65, 0x6e, 0x67, 0x00, 0x86, 0xe0, 0x6e, 0xf0,
		0x00, 0x7f, 0xc9, 0xad, 0x32, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
)

// Create creates a new packet with the provided information. This is primarily used for testing. If you are creating packets for use other than testing, you should reconsider your usage of this package.

// Example usage
// pkt, _ = SetCC(
//        Create(pid,
//              WithHasPayloadFlag,
//              WithContinuousAF,
//              WithPUSI),
//        cc)
func Create(pid int, options ...func(*Packet)) *Packet {
	var pkt Packet
	setPid(&pkt, pid)
	for _, option := range options {
		option(&pkt)
	}
	for _, option := range required {
		option(&pkt)
	}
	return &pkt
}

// CreateTestPacket creates a test packet with the given PID, continuity counter, payload unit start indicator and payload flag
// This is a convenience function for often used packet creatio options functions
func CreateTestPacket(pid int, cc uint8, pusi, hasPay bool) *Packet {
	var pkt *Packet
	if hasPay && pusi {
		pkt = SetCC(
			Create(pid,
				WithHasPayloadFlag,
				WithContinuousAF,
				WithPUSI),
			cc)
	} else if hasPay {
		pkt = SetCC(
			Create(pid,
				WithHasPayloadFlag,
				WithContinuousAF),
			cc)

	} else {
		pkt = SetCC(Create(pid, WithContinuousAF), cc)
	}
	return pkt
}

// CreateDCPacket creates a new packet with a discontinuous adapataion field and the given PID and CC
func CreateDCPacket(pid int, cc uint8) *Packet {
	pkt := SetCC(Create(pid, WithDiscontinuousAF, WithHasPayloadFlag), cc)
	return pkt
}

// CreatePacketWithPayload creates a new packet with the given PID, CC and payload
func CreatePacketWithPayload(pid int, cc uint8, pay []byte) *Packet {
	pkt := SetCC(
		Create(
			pid,
			WithHasPayloadFlag,
			WithContinuousAF,
			func(pkt *Packet) {
				SetPayload(pkt, pay)
			},
		),
		cc,
	)
	return pkt
}

func setPid(pkt *Packet, pid int) {
	pkt[1] = byte(pid >> 8 & 0x1f)
	pkt[2] = byte(pid & 0xff)
}

// WithHasPayloadFlag is an option function for creating a packet with a payload flag
func WithHasPayloadFlag(pkt *Packet) {
	pkt[3] |= 0x10
}

// WithHasAdaptationFieldFlag is an option function for creating a packet with an adaptation field
func WithHasAdaptationFieldFlag(pkt *Packet) {
	pkt[3] |= 0x20
}

// WithAFPrivateDataFlag is an option function for creating a packet with a adaptation field private data flag
func WithAFPrivateDataFlag(pkt *Packet) {
	pkt[5] |= 0x02
}

// WithPUSI is an option function for creating a packet with the payload unit start indicator flag set
func WithPUSI(pkt *Packet) {
	pkt[1] |= 0x40
}

// WithContinuousAF is an option function for creating a packet with a continuous adaptation field
func WithContinuousAF(pkt *Packet) {
	pkt[5] |= 0x7f
}

// WithDisconinuousAF is an option function for creating a packet with a discontinuous adaptation field
func WithDiscontinuousAF(pkt *Packet) {
	pkt[5] |= 0x80
}

// WithPES is an option function for creating a packet with a PES header
func WithPES(pkt *Packet, pts uint64) {
	size := PacketSize - 4
	pay := make([]byte, size, size)
	//packet_start_code_prefix
	pay[0] = 0x00
	pay[1] = 0x00
	pay[2] = 0x01
	// stream id 184 is STREAM_ID_ALL_VIDEO_STREAMS
	pay[3] = byte(184)
	// Packet Length
	pay[4] = 0x0
	//pay[5] = ??length
	// Data alignment indicator
	pay[6] = 0x40
	// PTS DTS indicator 2 is PTS_DTS_INDICATOR_ONLY_PTS
	pay[7] = 0x80
	// Header len
	pay[8] = 14
	gots.InsertPTS(pay[9:14], pts)
	SetPayload(pkt, pay)
	WithHasPayloadFlag(pkt)
}

// SetPayload sets the payload of a given packet
func SetPayload(pkt *Packet, pay []byte) int {
	start := payloadStart(pkt)
	i := start
	j := 0
	for i < PacketSize && j < len(pay) {
		pkt[i] = pay[j]
		i++
		j++
	}
	return i - start
}

// Required parts of a packet
func setSyncByte(pkt *Packet) {
	pkt[0] = SyncByte
}
