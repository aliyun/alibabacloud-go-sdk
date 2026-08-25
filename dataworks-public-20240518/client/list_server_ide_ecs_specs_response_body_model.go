// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServerIdeEcsSpecsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEcsSpecs(v []*ListServerIdeEcsSpecsResponseBodyEcsSpecs) *ListServerIdeEcsSpecsResponseBody
	GetEcsSpecs() []*ListServerIdeEcsSpecsResponseBodyEcsSpecs
	SetMaxResults(v int32) *ListServerIdeEcsSpecsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListServerIdeEcsSpecsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListServerIdeEcsSpecsResponseBody
	GetRequestId() *string
}

type ListServerIdeEcsSpecsResponseBody struct {
	// The list of available ECS instance types for personal development environments.
	EcsSpecs []*ListServerIdeEcsSpecsResponseBodyEcsSpecs `json:"EcsSpecs,omitempty" xml:"EcsSpecs,omitempty" type:"Repeated"`
	// The maximum number of records returned in this response.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page. An empty value indicates that no more results are available.
	//
	// example:
	//
	// CAESG****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListServerIdeEcsSpecsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServerIdeEcsSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServerIdeEcsSpecsResponseBody) GetEcsSpecs() []*ListServerIdeEcsSpecsResponseBodyEcsSpecs {
	return s.EcsSpecs
}

func (s *ListServerIdeEcsSpecsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServerIdeEcsSpecsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServerIdeEcsSpecsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServerIdeEcsSpecsResponseBody) SetEcsSpecs(v []*ListServerIdeEcsSpecsResponseBodyEcsSpecs) *ListServerIdeEcsSpecsResponseBody {
	s.EcsSpecs = v
	return s
}

func (s *ListServerIdeEcsSpecsResponseBody) SetMaxResults(v int32) *ListServerIdeEcsSpecsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServerIdeEcsSpecsResponseBody) SetNextToken(v string) *ListServerIdeEcsSpecsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServerIdeEcsSpecsResponseBody) SetRequestId(v string) *ListServerIdeEcsSpecsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServerIdeEcsSpecsResponseBody) Validate() error {
	if s.EcsSpecs != nil {
		for _, item := range s.EcsSpecs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServerIdeEcsSpecsResponseBodyEcsSpecs struct {
	// The accelerator type. Valid values:
	//
	// - CPU: uses only CPU.
	//
	// - GPU: uses GPU acceleration.
	//
	// example:
	//
	// CPU
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// The number of CPU cores.
	//
	// example:
	//
	// 4
	Cpu *int64 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The number of compute units (CUs) consumed by this instance type.
	//
	// example:
	//
	// 10
	Cu *float32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The number of GPU cards.
	//
	// example:
	//
	// 1
	Gpu *int64 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The GPU memory size.
	//
	// example:
	//
	// 16
	GpuMemorySize *float32 `json:"GpuMemorySize,omitempty" xml:"GpuMemorySize,omitempty"`
	// The GPU model.
	//
	// example:
	//
	// V100
	GpuType *string `json:"GpuType,omitempty" xml:"GpuType,omitempty"`
	// The ECS instance type.
	//
	// example:
	//
	// ecs.g6.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Indicates whether the instance type is available.
	IsAvailable *bool `json:"IsAvailable,omitempty" xml:"IsAvailable,omitempty"`
	// The memory size, in GB.
	//
	// example:
	//
	// 16
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListServerIdeEcsSpecsResponseBodyEcsSpecs) String() string {
	return dara.Prettify(s)
}

func (s ListServerIdeEcsSpecsResponseBodyEcsSpecs) GoString() string {
	return s.String()
}

func (s *ListServerIdeEcsSpecsResponseBodyEcsSpecs) GetAcceleratorType() *string {
	return s.AcceleratorType
}

func (s *ListServerIdeEcsSpecsResponseBodyEcsSpecs) GetCpu() *int64 {
	return s.Cpu
}

func (s *ListServerIdeEcsSpecsResponseBodyEcsSpecs) GetCu() *float32 {
	return s.Cu
}

func (s *ListServerIdeEcsSpecsResponseBodyEcsSpecs) GetGpu() *int64 {
	return s.Gpu
}

func (s *ListServerIdeEcsSpecsResponseBodyEcsSpecs) GetGpuMemorySize() *float32 {
	return s.GpuMemorySize
}

func (s *ListServerIdeEcsSpecsResponseBodyEcsSpecs) GetGpuType() *string {
	return s.GpuType
}

func (s *ListServerIdeEcsSpecsResponseBodyEcsSpecs) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListServerIdeEcsSpecsResponseBodyEcsSpecs) GetIsAvailable() *bool {
	return s.IsAvailable
}

func (s *ListServerIdeEcsSpecsResponseBodyEcsSpecs) GetMemory() *float32 {
	return s.Memory
}

func (s *ListServerIdeEcsSpecsResponseBodyEcsSpecs) SetAcceleratorType(v string) *ListServerIdeEcsSpecsResponseBodyEcsSpecs {
	s.AcceleratorType = &v
	return s
}

func (s *ListServerIdeEcsSpecsResponseBodyEcsSpecs) SetCpu(v int64) *ListServerIdeEcsSpecsResponseBodyEcsSpecs {
	s.Cpu = &v
	return s
}

func (s *ListServerIdeEcsSpecsResponseBodyEcsSpecs) SetCu(v float32) *ListServerIdeEcsSpecsResponseBodyEcsSpecs {
	s.Cu = &v
	return s
}

func (s *ListServerIdeEcsSpecsResponseBodyEcsSpecs) SetGpu(v int64) *ListServerIdeEcsSpecsResponseBodyEcsSpecs {
	s.Gpu = &v
	return s
}

func (s *ListServerIdeEcsSpecsResponseBodyEcsSpecs) SetGpuMemorySize(v float32) *ListServerIdeEcsSpecsResponseBodyEcsSpecs {
	s.GpuMemorySize = &v
	return s
}

func (s *ListServerIdeEcsSpecsResponseBodyEcsSpecs) SetGpuType(v string) *ListServerIdeEcsSpecsResponseBodyEcsSpecs {
	s.GpuType = &v
	return s
}

func (s *ListServerIdeEcsSpecsResponseBodyEcsSpecs) SetInstanceType(v string) *ListServerIdeEcsSpecsResponseBodyEcsSpecs {
	s.InstanceType = &v
	return s
}

func (s *ListServerIdeEcsSpecsResponseBodyEcsSpecs) SetIsAvailable(v bool) *ListServerIdeEcsSpecsResponseBodyEcsSpecs {
	s.IsAvailable = &v
	return s
}

func (s *ListServerIdeEcsSpecsResponseBodyEcsSpecs) SetMemory(v float32) *ListServerIdeEcsSpecsResponseBodyEcsSpecs {
	s.Memory = &v
	return s
}

func (s *ListServerIdeEcsSpecsResponseBodyEcsSpecs) Validate() error {
	return dara.Validate(s)
}
