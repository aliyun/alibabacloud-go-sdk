// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdgeMachineModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModels(v []*DescribeEdgeMachineModelsResponseBodyModels) *DescribeEdgeMachineModelsResponseBody
	GetModels() []*DescribeEdgeMachineModelsResponseBodyModels
}

type DescribeEdgeMachineModelsResponseBody struct {
	// The cloud-native box models.
	Models []*DescribeEdgeMachineModelsResponseBodyModels `json:"models,omitempty" xml:"models,omitempty" type:"Repeated"`
}

func (s DescribeEdgeMachineModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeMachineModelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEdgeMachineModelsResponseBody) GetModels() []*DescribeEdgeMachineModelsResponseBodyModels {
	return s.Models
}

func (s *DescribeEdgeMachineModelsResponseBody) SetModels(v []*DescribeEdgeMachineModelsResponseBodyModels) *DescribeEdgeMachineModelsResponseBody {
	s.Models = v
	return s
}

func (s *DescribeEdgeMachineModelsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEdgeMachineModelsResponseBodyModels struct {
	// The number of vCores.
	//
	// example:
	//
	// 6
	Cpu *int32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The CPU architecture.
	//
	// example:
	//
	// x86_64/arm64
	CpuArch *string `json:"cpu_arch,omitempty" xml:"cpu_arch,omitempty"`
	// The time when the cloud-native box was created.
	//
	// example:
	//
	// 2021-07-07T20:44:00+08:00
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// The description of the cloud-native box.
	//
	// example:
	//
	// B010
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Indicates whether the cloud-native box model manages the Docker runtime.
	//
	// example:
	//
	// 0/1
	ManageRuntime *int32 `json:"manage_runtime,omitempty" xml:"manage_runtime,omitempty"`
	// The memory. Unit: GB.
	//
	// example:
	//
	// 8
	Memory *int32 `json:"memory,omitempty" xml:"memory,omitempty"`
	// The model of the cloud-native box.
	//
	// example:
	//
	// ACK-V-B010
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// The ID of the cloud-native box.
	//
	// example:
	//
	// c34cc753-8908-4739-bd10-ebd922a4****
	ModelId *string `json:"model_id,omitempty" xml:"model_id,omitempty"`
}

func (s DescribeEdgeMachineModelsResponseBodyModels) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeMachineModelsResponseBodyModels) GoString() string {
	return s.String()
}

func (s *DescribeEdgeMachineModelsResponseBodyModels) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeEdgeMachineModelsResponseBodyModels) GetCpuArch() *string {
	return s.CpuArch
}

func (s *DescribeEdgeMachineModelsResponseBodyModels) GetCreated() *string {
	return s.Created
}

func (s *DescribeEdgeMachineModelsResponseBodyModels) GetDescription() *string {
	return s.Description
}

func (s *DescribeEdgeMachineModelsResponseBodyModels) GetManageRuntime() *int32 {
	return s.ManageRuntime
}

func (s *DescribeEdgeMachineModelsResponseBodyModels) GetMemory() *int32 {
	return s.Memory
}

func (s *DescribeEdgeMachineModelsResponseBodyModels) GetModel() *string {
	return s.Model
}

func (s *DescribeEdgeMachineModelsResponseBodyModels) GetModelId() *string {
	return s.ModelId
}

func (s *DescribeEdgeMachineModelsResponseBodyModels) SetCpu(v int32) *DescribeEdgeMachineModelsResponseBodyModels {
	s.Cpu = &v
	return s
}

func (s *DescribeEdgeMachineModelsResponseBodyModels) SetCpuArch(v string) *DescribeEdgeMachineModelsResponseBodyModels {
	s.CpuArch = &v
	return s
}

func (s *DescribeEdgeMachineModelsResponseBodyModels) SetCreated(v string) *DescribeEdgeMachineModelsResponseBodyModels {
	s.Created = &v
	return s
}

func (s *DescribeEdgeMachineModelsResponseBodyModels) SetDescription(v string) *DescribeEdgeMachineModelsResponseBodyModels {
	s.Description = &v
	return s
}

func (s *DescribeEdgeMachineModelsResponseBodyModels) SetManageRuntime(v int32) *DescribeEdgeMachineModelsResponseBodyModels {
	s.ManageRuntime = &v
	return s
}

func (s *DescribeEdgeMachineModelsResponseBodyModels) SetMemory(v int32) *DescribeEdgeMachineModelsResponseBodyModels {
	s.Memory = &v
	return s
}

func (s *DescribeEdgeMachineModelsResponseBodyModels) SetModel(v string) *DescribeEdgeMachineModelsResponseBodyModels {
	s.Model = &v
	return s
}

func (s *DescribeEdgeMachineModelsResponseBodyModels) SetModelId(v string) *DescribeEdgeMachineModelsResponseBodyModels {
	s.ModelId = &v
	return s
}

func (s *DescribeEdgeMachineModelsResponseBodyModels) Validate() error {
	return dara.Validate(s)
}
