package main

import (
	"fmt"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func getTCPLayer() *layers.TCP {
	return &layers.TCP{
		SrcPort: layers.TCPPort(12345), // Source port
		DstPort: layers.TCPPort(80),    // Destination port
		Seq:     123456,                // Sequence number
		Ack:     0,                     // Acknowledgment number
		SYN:     true,                  // Flags: SYN
		Window:  14600,                 // Window size
	}
}
func getIPLayer() *layers.IPv4 {
	return &layers.IPv4{
		Version:  4,
		IHL:      5,
		TTL:      64,
		Protocol: layers.IPProtocolTCP,
		SrcIP:    net.IP{192, 168, 1, 1}, // Source IP address
		DstIP:    net.IP{192, 168, 1, 2}, // Destination IP address
	}

}
func getEthernetLayer() *layers.Ethernet {
	return &layers.Ethernet{
		SrcMAC:       net.HardwareAddr{0x00, 0x11, 0x22, 0x33, 0x44, 0x55}, // Source MAC address
		DstMAC:       net.HardwareAddr{0x66, 0x77, 0x88, 0x99, 0xAA, 0xBB}, // Destination MAC address
		EthernetType: layers.EthernetTypeIPv4,                              // EtherType for IPv4
	}
}

func getHTTPPacket() *gopacket.SerializeBuffer {
	// Create a new packet builder
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{}
	// Create an Ethernet layer
	ethernetLayer := getEthernetLayer()

	// Create an IP layer
	ipLayer := getIPLayer()
	// Create a TCP layer
	tcpLayer := getTCPLayer()

	// Create an HTTP payload
	httpPayload := gopacket.Payload([]byte("GET / HTTP/1.1\r\nHost: example.com\r\n\r\n"))

	gopacket.SerializeLayers(buf, opts,
		ethernetLayer,
		ipLayer,
		tcpLayer,
		httpPayload)
	// Access the raw packet bytes
	return &buf
}

func main() {
	packet := getHTTPPacket()
	fmt.Printf("%s\n", *packet)
}
