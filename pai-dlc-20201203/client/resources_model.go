// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResources interface {
	dara.Model
	String() string
	GoString() string
	SetCPU(v string) *Resources
	GetCPU() *string
	SetGPU(v string) *Resources
	GetGPU() *string
	SetMemory(v string) *Resources
	GetMemory() *string
}

type Resources struct {
	// example:
	//
	// 10
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// example:
	//
	// 8
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// example:
	//
	// 1024（单位GB）
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s Resources) String() string {
	return dara.Prettify(s)
}

func (s Resources) GoString() string {
	return s.String()
}

func (s *Resources) GetCPU() *string {
	return s.CPU
}

func (s *Resources) GetGPU() *string {
	return s.GPU
}

func (s *Resources) GetMemory() *string {
	return s.Memory
}

func (s *Resources) SetCPU(v string) *Resources {
	s.CPU = &v
	return s
}

func (s *Resources) SetGPU(v string) *Resources {
	s.GPU = &v
	return s
}

func (s *Resources) SetMemory(v string) *Resources {
	s.Memory = &v
	return s
}

func (s *Resources) Validate() error {
	return dara.Validate(s)
}
