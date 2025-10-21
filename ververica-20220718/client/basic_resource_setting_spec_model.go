// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBasicResourceSettingSpec interface {
	dara.Model
	String() string
	GoString() string
	SetCpu(v float64) *BasicResourceSettingSpec
	GetCpu() *float64
	SetMemory(v string) *BasicResourceSettingSpec
	GetMemory() *string
}

type BasicResourceSettingSpec struct {
	// example:
	//
	// 2.0
	Cpu *float64 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// example:
	//
	// 4Gi
	Memory *string `json:"memory,omitempty" xml:"memory,omitempty"`
}

func (s BasicResourceSettingSpec) String() string {
	return dara.Prettify(s)
}

func (s BasicResourceSettingSpec) GoString() string {
	return s.String()
}

func (s *BasicResourceSettingSpec) GetCpu() *float64 {
	return s.Cpu
}

func (s *BasicResourceSettingSpec) GetMemory() *string {
	return s.Memory
}

func (s *BasicResourceSettingSpec) SetCpu(v float64) *BasicResourceSettingSpec {
	s.Cpu = &v
	return s
}

func (s *BasicResourceSettingSpec) SetMemory(v string) *BasicResourceSettingSpec {
	s.Memory = &v
	return s
}

func (s *BasicResourceSettingSpec) Validate() error {
	return dara.Validate(s)
}
