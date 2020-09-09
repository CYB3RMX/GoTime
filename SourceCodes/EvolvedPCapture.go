/* Same as SimplePCapture.go but this one does packet decoding */
package main

import (
    "fmt"
    "log"
    "time"
    "github.com/google/gopacket"
    "github.com/google/gopacket/pcap"
    "github.com/google/gopacket/layers"

)

// Setting variables
var (
    netDevice string = "wlan0"
    snapshotLen int32 = 1024
    promisc bool = false
    err error
    timeout time.Duration = 30 * time.Second
    handle *pcap.Handle
)

func main(){
    // Accessing network device
    handle, err := pcap.OpenLive(netDevice, snapshotLen, promisc, timeout)

    // error handling
    if err != nil {
        log.Fatalln(err)
    }
    defer handle.Close()

    // handling packets
    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
    fmt.Println("[*] Packet capturing starts...")
    for packet := range packetSource.Packets() {
        packetInfo(packet)
    }
}

// parsing packet informations
func packetInfo(packet gopacket.Packet){
    // parsing IP layer
    ipLayer := packet.Layer(layers.LayerTypeIPv4)
    if ipLayer != nil {
        // getting layer variables
        ips, _ := ipLayer.(*layers.IPv4)

        // printing
        fmt.Printf("\nSrcAddr: %s  => DstAddr: %s\n", ips.SrcIP, ips.DstIP)
    }

    // parsing TCP layer
    tcpLayer := packet.Layer(layers.LayerTypeTCP)
    if tcpLayer != nil {
        // getting layer variables
        tcp, _ := tcpLayer.(*layers.TCP)

        // printing
        fmt.Printf("SrcPort: %d  => DstPort: %d\n", tcp.SrcPort, tcp.DstPort)
    }
}
