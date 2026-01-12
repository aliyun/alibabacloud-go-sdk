// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceLimit interface {
	dara.Model
	String() string
	GoString() string
	SetCPU(v string) *ResourceLimit
	GetCPU() *string
	SetGPU(v string) *ResourceLimit
	GetGPU() *string
	SetMemory(v string) *ResourceLimit
	GetMemory() *string
}

type ResourceLimit struct {
	CPU    *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	GPU    *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ResourceLimit) String() string {
	return dara.Prettify(s)
}

func (s ResourceLimit) GoString() string {
	return s.String()
}

func (s *ResourceLimit) GetCPU() *string {
	return s.CPU
}

func (s *ResourceLimit) GetGPU() *string {
	return s.GPU
}

func (s *ResourceLimit) GetMemory() *string {
	return s.Memory
}

func (s *ResourceLimit) SetCPU(v string) *ResourceLimit {
	s.CPU = &v
	return s
}

func (s *ResourceLimit) SetGPU(v string) *ResourceLimit {
	s.GPU = &v
	return s
}

func (s *ResourceLimit) SetMemory(v string) *ResourceLimit {
	s.Memory = &v
	return s
}

func (s *ResourceLimit) Validate() error {
	return dara.Validate(s)
}
