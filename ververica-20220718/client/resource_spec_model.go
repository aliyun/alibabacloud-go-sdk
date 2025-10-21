// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceSpec interface {
	dara.Model
	String() string
	GoString() string
	SetCpu(v float64) *ResourceSpec
	GetCpu() *float64
	SetMemory(v string) *ResourceSpec
	GetMemory() *string
}

type ResourceSpec struct {
	// example:
	//
	// 1.0
	Cpu *float64 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// example:
	//
	// 4Gi
	Memory *string `json:"memory,omitempty" xml:"memory,omitempty"`
}

func (s ResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s ResourceSpec) GoString() string {
	return s.String()
}

func (s *ResourceSpec) GetCpu() *float64 {
	return s.Cpu
}

func (s *ResourceSpec) GetMemory() *string {
	return s.Memory
}

func (s *ResourceSpec) SetCpu(v float64) *ResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ResourceSpec) SetMemory(v string) *ResourceSpec {
	s.Memory = &v
	return s
}

func (s *ResourceSpec) Validate() error {
	return dara.Validate(s)
}
