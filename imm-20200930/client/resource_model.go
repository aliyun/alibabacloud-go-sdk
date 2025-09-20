// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResource interface {
	dara.Model
	String() string
	GoString() string
	SetCPU(v int64) *Resource
	GetCPU() *int64
	SetECSInstance(v string) *Resource
	GetECSInstance() *string
	SetGPUModel(v string) *Resource
	GetGPUModel() *string
	SetGPUNum(v int64) *Resource
	GetGPUNum() *int64
	SetName(v string) *Resource
	GetName() *string
	SetRAM(v int64) *Resource
	GetRAM() *int64
}

type Resource struct {
	// example:
	//
	// 2
	CPU *int64 `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// example:
	//
	// ecs.gn5i-c2g1.large
	ECSInstance *string `json:"ECSInstance,omitempty" xml:"ECSInstance,omitempty"`
	// example:
	//
	// string	NVIDIA_P4
	GPUModel *string `json:"GPUModel,omitempty" xml:"GPUModel,omitempty"`
	// example:
	//
	// 1
	GPUNum *int64 `json:"GPUNum,omitempty" xml:"GPUNum,omitempty"`
	// example:
	//
	// string	ecs.gn5i-c2g1.large-2vCPU-8GB-1*NVIDIA_P4
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 8
	RAM *int64 `json:"RAM,omitempty" xml:"RAM,omitempty"`
}

func (s Resource) String() string {
	return dara.Prettify(s)
}

func (s Resource) GoString() string {
	return s.String()
}

func (s *Resource) GetCPU() *int64 {
	return s.CPU
}

func (s *Resource) GetECSInstance() *string {
	return s.ECSInstance
}

func (s *Resource) GetGPUModel() *string {
	return s.GPUModel
}

func (s *Resource) GetGPUNum() *int64 {
	return s.GPUNum
}

func (s *Resource) GetName() *string {
	return s.Name
}

func (s *Resource) GetRAM() *int64 {
	return s.RAM
}

func (s *Resource) SetCPU(v int64) *Resource {
	s.CPU = &v
	return s
}

func (s *Resource) SetECSInstance(v string) *Resource {
	s.ECSInstance = &v
	return s
}

func (s *Resource) SetGPUModel(v string) *Resource {
	s.GPUModel = &v
	return s
}

func (s *Resource) SetGPUNum(v int64) *Resource {
	s.GPUNum = &v
	return s
}

func (s *Resource) SetName(v string) *Resource {
	s.Name = &v
	return s
}

func (s *Resource) SetRAM(v int64) *Resource {
	s.RAM = &v
	return s
}

func (s *Resource) Validate() error {
	return dara.Validate(s)
}
