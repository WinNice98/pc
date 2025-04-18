package pc

import (
	"evm"
	"fmt"
)

type PC struct {
	evm.EVM
	GPU      string
	VRAM     uint64
	Ethernet bool
}

func (p *PC) Connect() {
	if p.Ethernet {
		fmt.Println("Интернет уже подключён")
	} else {
		p.Ethernet = true
		fmt.Println("Интернет успешно подключён")
	}
}

func showbool(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func (p *PC) ShowSpecs() {
	fmt.Printf("%s\n%s\n%d\n%s\n%d\nEthernet = %s\n",
		p.Name, p.CPU, p.RAM, p.GPU, p.VRAM, showbool(p.Ethernet))
}

