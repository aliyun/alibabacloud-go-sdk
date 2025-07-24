// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCostInstanceType interface {
	dara.Model
	String() string
	GoString() string
	SetCpu(v int32) *CostInstanceType
	GetCpu() *int32
	SetDataDisks(v []*DataDisk) *CostInstanceType
	GetDataDisks() []*DataDisk
	SetInstanceType(v string) *CostInstanceType
	GetInstanceType() *string
	SetMemory(v int32) *CostInstanceType
	GetMemory() *int32
	SetSystemDisk(v *SystemDisk) *CostInstanceType
	GetSystemDisk() *SystemDisk
}

type CostInstanceType struct {
	// CPU核数。
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// 数据盘列表。
	DataDisks []*DataDisk `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// 实例类型列表。
	//
	// example:
	//
	// ["ecs.g6.4xlarge"]
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// 内存大小。
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// 系统盘信息。
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty"`
}

func (s CostInstanceType) String() string {
	return dara.Prettify(s)
}

func (s CostInstanceType) GoString() string {
	return s.String()
}

func (s *CostInstanceType) GetCpu() *int32 {
	return s.Cpu
}

func (s *CostInstanceType) GetDataDisks() []*DataDisk {
	return s.DataDisks
}

func (s *CostInstanceType) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CostInstanceType) GetMemory() *int32 {
	return s.Memory
}

func (s *CostInstanceType) GetSystemDisk() *SystemDisk {
	return s.SystemDisk
}

func (s *CostInstanceType) SetCpu(v int32) *CostInstanceType {
	s.Cpu = &v
	return s
}

func (s *CostInstanceType) SetDataDisks(v []*DataDisk) *CostInstanceType {
	s.DataDisks = v
	return s
}

func (s *CostInstanceType) SetInstanceType(v string) *CostInstanceType {
	s.InstanceType = &v
	return s
}

func (s *CostInstanceType) SetMemory(v int32) *CostInstanceType {
	s.Memory = &v
	return s
}

func (s *CostInstanceType) SetSystemDisk(v *SystemDisk) *CostInstanceType {
	s.SystemDisk = v
	return s
}

func (s *CostInstanceType) Validate() error {
	return dara.Validate(s)
}
