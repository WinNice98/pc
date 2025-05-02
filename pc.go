package pc

import (
	"fmt"

	"github.com/WinNice98/evm"
)

type PC struct {
	*evm.EVM
	GPU string
	SSD uint64
}

func NewPC(name, cpu, gpu string, ram, ssd uint64) *PC {
	return &PC{
		EVM: evm.NewEVM(name, cpu, ram),
		GPU: gpu,
		SSD: ssd,
	}
}

func (p *PC) ShowSpecs() string {
	return fmt.Sprintf("%s | GPU: %s | SSD: %d GB", p.EVM.ShowSpecs(), p.GPU, p.SSD)
}

