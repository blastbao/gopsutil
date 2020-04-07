package main

import (
	"fmt"
	"github.com/blastbao/gopsutil/cpu"
	"github.com/blastbao/gopsutil/disk"
	"github.com/blastbao/gopsutil/load"
	"github.com/blastbao/gopsutil/mem"
	"github.com/blastbao/gopsutil/net"
	"github.com/blastbao/gopsutil/process"
)

func main() {
	fmt.Println("CPU统计:")
	c, _ := cpu.Info()
	fmt.Println(c)
	fmt.Println("内存统计:")
	m, _ := mem.VirtualMemory()
	fmt.Println(m)
	fmt.Println("磁盘用量和IO统计:")
	dp, _ := disk.Partitions(true)
	du, _ := disk.Usage("/")
	di, _ := disk.IOCounters()
	fmt.Println(du)
	fmt.Println(dp)
	fmt.Println(di)
	fmt.Println("网络IO统计:")
	ni, _ := net.IOCounters(true)
	fmt.Println(ni)
	fmt.Println("协议统计:")
	nt, _ := net.ProtoCounters(nil)
	fmt.Println(nt)
	fmt.Println("链接状态统计:")
	nc, _ := net.Connections("all")
	fmt.Println(nc)
	fmt.Println("进程统计:")
	pi, _ := process.Pids()
	fmt.Println(pi)
	p, _ := process.NewProcess(614)
	pm, _ := p.MemoryPercent()
	pn, _ := p.Username()
	fmt.Println(pm)
	fmt.Println(pn)
	fmt.Println("负载统计:")
	pl, _ := load.Avg()
	fmt.Println(pl)
}