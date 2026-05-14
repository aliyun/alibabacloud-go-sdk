// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *CreateModelServiceRequest
	GetApiKey() *string
	SetCpu(v int64) *CreateModelServiceRequest
	GetCpu() *int64
	SetGpu(v int64) *CreateModelServiceRequest
	GetGpu() *int64
	SetGpuMemory(v int64) *CreateModelServiceRequest
	GetGpuMemory() *int64
	SetMemory(v int64) *CreateModelServiceRequest
	GetMemory() *int64
	SetModelParams(v string) *CreateModelServiceRequest
	GetModelParams() *string
	SetModelServiceName(v string) *CreateModelServiceRequest
	GetModelServiceName() *string
	SetModelType(v string) *CreateModelServiceRequest
	GetModelType() *string
	SetProvider(v string) *CreateModelServiceRequest
	GetProvider() *string
	SetServiceCount(v int64) *CreateModelServiceRequest
	GetServiceCount() *int64
	SetTaskType(v string) *CreateModelServiceRequest
	GetTaskType() *string
}

type CreateModelServiceRequest struct {
	// example:
	//
	// api-key-xxx
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// example:
	//
	// 16
	Cpu *int64 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// example:
	//
	// 1
	Gpu *int64 `json:"gpu,omitempty" xml:"gpu,omitempty"`
	// example:
	//
	// 64
	GpuMemory *int64 `json:"gpuMemory,omitempty" xml:"gpuMemory,omitempty"`
	// example:
	//
	// 64
	Memory *int64 `json:"memory,omitempty" xml:"memory,omitempty"`
	// example:
	//
	// {"timeout":600,"max_retries":10,"max_retry_delay":8,"initial_retry_delay":0.5}
	ModelParams *string `json:"modelParams,omitempty" xml:"modelParams,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my_model
	ModelServiceName *string `json:"modelServiceName,omitempty" xml:"modelServiceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qwen3.5-plus
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// example:
	//
	// bailian
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// example:
	//
	// 2
	ServiceCount *int64 `json:"serviceCount,omitempty" xml:"serviceCount,omitempty"`
	// example:
	//
	// embedding
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s CreateModelServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateModelServiceRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateModelServiceRequest) GetCpu() *int64 {
	return s.Cpu
}

func (s *CreateModelServiceRequest) GetGpu() *int64 {
	return s.Gpu
}

func (s *CreateModelServiceRequest) GetGpuMemory() *int64 {
	return s.GpuMemory
}

func (s *CreateModelServiceRequest) GetMemory() *int64 {
	return s.Memory
}

func (s *CreateModelServiceRequest) GetModelParams() *string {
	return s.ModelParams
}

func (s *CreateModelServiceRequest) GetModelServiceName() *string {
	return s.ModelServiceName
}

func (s *CreateModelServiceRequest) GetModelType() *string {
	return s.ModelType
}

func (s *CreateModelServiceRequest) GetProvider() *string {
	return s.Provider
}

func (s *CreateModelServiceRequest) GetServiceCount() *int64 {
	return s.ServiceCount
}

func (s *CreateModelServiceRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateModelServiceRequest) SetApiKey(v string) *CreateModelServiceRequest {
	s.ApiKey = &v
	return s
}

func (s *CreateModelServiceRequest) SetCpu(v int64) *CreateModelServiceRequest {
	s.Cpu = &v
	return s
}

func (s *CreateModelServiceRequest) SetGpu(v int64) *CreateModelServiceRequest {
	s.Gpu = &v
	return s
}

func (s *CreateModelServiceRequest) SetGpuMemory(v int64) *CreateModelServiceRequest {
	s.GpuMemory = &v
	return s
}

func (s *CreateModelServiceRequest) SetMemory(v int64) *CreateModelServiceRequest {
	s.Memory = &v
	return s
}

func (s *CreateModelServiceRequest) SetModelParams(v string) *CreateModelServiceRequest {
	s.ModelParams = &v
	return s
}

func (s *CreateModelServiceRequest) SetModelServiceName(v string) *CreateModelServiceRequest {
	s.ModelServiceName = &v
	return s
}

func (s *CreateModelServiceRequest) SetModelType(v string) *CreateModelServiceRequest {
	s.ModelType = &v
	return s
}

func (s *CreateModelServiceRequest) SetProvider(v string) *CreateModelServiceRequest {
	s.Provider = &v
	return s
}

func (s *CreateModelServiceRequest) SetServiceCount(v int64) *CreateModelServiceRequest {
	s.ServiceCount = &v
	return s
}

func (s *CreateModelServiceRequest) SetTaskType(v string) *CreateModelServiceRequest {
	s.TaskType = &v
	return s
}

func (s *CreateModelServiceRequest) Validate() error {
	return dara.Validate(s)
}
