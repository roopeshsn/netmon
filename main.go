package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/opennetworktools/netmon/internal"
)

func main() {
	c := make(chan internal.CPacket)
	m := internal.NewHostMap()
	mc := make(chan internal.CHost)

	var wg sync.WaitGroup

	args := os.Args
	if len(args) == 1 {
		fmt.Println("No command entered!")
	} else if args[1] == "show" || args[1] == "sh" {
		if args[2] == "interface" || args[2] == "int" {
			if len(args) >= 5 {
				if args[3] == "describe" || args[3] == "des" {
					internal.FindInterfaceDescribe(args[4])
				}
			}
		} else if args[2] == "interfaces" || args[2] == "ints" {
			if len(args) >= 4 {
				if args[3] == "describe" || args[3] == "des" {
					internal.FindAllInterfacesDescribe()
				}
				return
			}
			internal.FindAllInterfaces()
		} else if args[2] == "ip" {
			internal.GetLocalIP()
		}
	} else if args[1] == "watch" {
		if args[2] == "packets" {
			if len(args) >= 5 {
				if args[4] == "html" {
					wg.Add(2)
					go internal.WatchInterface(args[3], c)
					go internal.ResolveHostsInformation(args[3], c, m, mc, true)
					// go internal.StartServer(c, m)
					wg.Wait()
				} else if args[4] == "web" {
					wg.Add(3)
					go internal.WatchInterface(args[3], c)
					go internal.ResolveHostsInformation(args[3], c, m, mc, true)
					go internal.SendEvents(c, mc)
					wg.Wait()
				}
			} else {
				wg.Add(2)
				go internal.WatchInterface(args[3], c)
				go internal.PrintPacket(c)
				wg.Wait()
			}
		} else if args[2] == "hosts" {
			wg.Add(2)
			go internal.WatchInterface(args[3], c)
			go internal.ResolveHostsInformation(args[3], c, m, mc, false)
			wg.Wait()
		}
	} else {
		fmt.Printf("Command %v not found!\n", args[2])
	}
}