package main

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

func main() {

	cpuInfo, err := cpu.Info()
	if err != nil{
		log.Fatalf("Error fetching CPU info: %v", err)
	}
	fmt.Println("CPU Info: ")
	for _, ci := range cpuInfo {
		fmt.Printf("Model: %s, Cores: %d, Speed: %.2f MHz\n", ci.ModelName, ci.Cores, ci.Mhz)
	}

	memInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Fatalf("Error fetching memory info: %v", err)
	}

	fmt.Printf("Total Memory: %.2f GB, Used: %.2f GB, Free: %.2f GB\n",
	 float64(memInfo.Total) / (1024 * 1024 * 1024),
	 float64(memInfo.Used) / (1024 * 1024 * 1024),
	 float64(memInfo.Free) / (1024 * 1024 * 1024),
	)

	netInfo, err := net.IOCounters(true)
	if err != nil {
		log.Fatalf("Error fetching network info: %v", err)
	}
	fmt.Println("Network Info: ")
	for _, ni := range netInfo {
		fmt.Printf("Interface: %s, Bytes Sent: %d, Bytes Received: %d\n",ni.Name, ni.BytesSent, ni.BytesRecv)
	}


	hostInfo, err := host.Info()
	if err != nil {
		log.Fatalf("Error fetching host info: %v", err)
	}
	fmt.Printf("Os: %s, Patform: %s, Patform Version: %s, Uptime: %d seconsds\n",
	hostInfo.OS,
	hostInfo.Platform,
	hostInfo.PlatformFamily,
	hostInfo.Uptime)
}
