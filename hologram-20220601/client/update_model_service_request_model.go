// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCpu(v int64) *UpdateModelServiceRequest
	GetCpu() *int64
	SetGpu(v int64) *UpdateModelServiceRequest
	GetGpu() *int64
	SetMemory(v int64) *UpdateModelServiceRequest
	GetMemory() *int64
	SetModelServiceName(v string) *UpdateModelServiceRequest
	GetModelServiceName() *string
	SetModelType(v string) *UpdateModelServiceRequest
	GetModelType() *string
	SetServiceCount(v int64) *UpdateModelServiceRequest
	GetServiceCount() *int64
}

type UpdateModelServiceRequest struct {
	// example:
	//
	// 32
	Cpu *int64 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// example:
	//
	// 1
	Gpu *int64 `json:"gpu,omitempty" xml:"gpu,omitempty"`
	// example:
	//
	// 60
	Memory *int64 `json:"memory,omitempty" xml:"memory,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// model-qwen
	ModelServiceName *string `json:"modelServiceName,omitempty" xml:"modelServiceName,omitempty"`
	// example:
	//
	// Qwen/Qwen2.5-VL-32B-Instruct
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// example:
	//
	// 2
	ServiceCount *int64 `json:"serviceCount,omitempty" xml:"serviceCount,omitempty"`
}

func (s UpdateModelServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelServiceRequest) GetCpu() *int64 {
	return s.Cpu
}

func (s *UpdateModelServiceRequest) GetGpu() *int64 {
	return s.Gpu
}

func (s *UpdateModelServiceRequest) GetMemory() *int64 {
	return s.Memory
}

func (s *UpdateModelServiceRequest) GetModelServiceName() *string {
	return s.ModelServiceName
}

func (s *UpdateModelServiceRequest) GetModelType() *string {
	return s.ModelType
}

func (s *UpdateModelServiceRequest) GetServiceCount() *int64 {
	return s.ServiceCount
}

func (s *UpdateModelServiceRequest) SetCpu(v int64) *UpdateModelServiceRequest {
	s.Cpu = &v
	return s
}

func (s *UpdateModelServiceRequest) SetGpu(v int64) *UpdateModelServiceRequest {
	s.Gpu = &v
	return s
}

func (s *UpdateModelServiceRequest) SetMemory(v int64) *UpdateModelServiceRequest {
	s.Memory = &v
	return s
}

func (s *UpdateModelServiceRequest) SetModelServiceName(v string) *UpdateModelServiceRequest {
	s.ModelServiceName = &v
	return s
}

func (s *UpdateModelServiceRequest) SetModelType(v string) *UpdateModelServiceRequest {
	s.ModelType = &v
	return s
}

func (s *UpdateModelServiceRequest) SetServiceCount(v int64) *UpdateModelServiceRequest {
	s.ServiceCount = &v
	return s
}

func (s *UpdateModelServiceRequest) Validate() error {
	return dara.Validate(s)
}
