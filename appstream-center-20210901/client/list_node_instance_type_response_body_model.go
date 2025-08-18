// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeInstanceTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodeInstanceTypeModels(v []*ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) *ListNodeInstanceTypeResponseBody
	GetNodeInstanceTypeModels() []*ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels
	SetPageNumber(v int32) *ListNodeInstanceTypeResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListNodeInstanceTypeResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListNodeInstanceTypeResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListNodeInstanceTypeResponseBody
	GetTotalCount() *int32
}

type ListNodeInstanceTypeResponseBody struct {
	// The resource types.
	NodeInstanceTypeModels []*ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels `json:"NodeInstanceTypeModels,omitempty" xml:"NodeInstanceTypeModels,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodeInstanceTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodeInstanceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeInstanceTypeResponseBody) GetNodeInstanceTypeModels() []*ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	return s.NodeInstanceTypeModels
}

func (s *ListNodeInstanceTypeResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNodeInstanceTypeResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNodeInstanceTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNodeInstanceTypeResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNodeInstanceTypeResponseBody) SetNodeInstanceTypeModels(v []*ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) *ListNodeInstanceTypeResponseBody {
	s.NodeInstanceTypeModels = v
	return s
}

func (s *ListNodeInstanceTypeResponseBody) SetPageNumber(v int32) *ListNodeInstanceTypeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBody) SetPageSize(v int32) *ListNodeInstanceTypeResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBody) SetRequestId(v string) *ListNodeInstanceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBody) SetTotalCount(v int32) *ListNodeInstanceTypeResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels struct {
	// The number of vCPUs.
	//
	// example:
	//
	// 4
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The number of GPUs.
	//
	// example:
	//
	// 2
	Gpu *string `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The GPU size. Unit: MB.
	//
	// example:
	//
	// 8192
	GpuMemory *int64 `json:"GpuMemory,omitempty" xml:"GpuMemory,omitempty"`
	// The maximum number of sessions to which a resource can connect at the same time. If a resource connects to a large number of sessions at the same time, user experience can be compromised. The value range varies based on the resource type. The following items describe the value ranges of different resource types:
	//
	// 	- appstreaming.general.4c8g: 1 to 2
	//
	// 	- appstreaming.general.8c16g: 1 to 4
	//
	// 	- appstreaming.vgpu.8c16g.4g: 1 to 4
	//
	// 	- appstreaming.vgpu.8c31g.16g: 1 to 4
	//
	// 	- appstreaming.vgpu.14c93g.12g: 1 to 6
	//
	// example:
	//
	// 4
	MaxCapacity *int32 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	// The memory size. Unit: MB.
	//
	// example:
	//
	// 8192
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The ID of the resource type.
	//
	// example:
	//
	// appstreaming.vgpu.4c8g.2g
	NodeInstanceType *string `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	// The resource type family.
	//
	// Valid values:
	//
	// 	- appstreaming.general: WUYING - General
	//
	// 	- appstreaming.vgpu: WUYING - Graphics
	//
	// example:
	//
	// appstreaming.vgpu
	NodeInstanceTypeFamily *string `json:"NodeInstanceTypeFamily,omitempty" xml:"NodeInstanceTypeFamily,omitempty"`
	// The name of the resource type.
	//
	// example:
	//
	// WUYING - General - 4 vCPUs 8 GB Memory
	NodeTypeName *string `json:"NodeTypeName,omitempty" xml:"NodeTypeName,omitempty"`
}

func (s ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) String() string {
	return dara.Prettify(s)
}

func (s ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) GoString() string {
	return s.String()
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) GetCpu() *string {
	return s.Cpu
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) GetGpu() *string {
	return s.Gpu
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) GetGpuMemory() *int64 {
	return s.GpuMemory
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) GetMaxCapacity() *int32 {
	return s.MaxCapacity
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) GetMemory() *int64 {
	return s.Memory
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) GetNodeInstanceType() *string {
	return s.NodeInstanceType
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) GetNodeInstanceTypeFamily() *string {
	return s.NodeInstanceTypeFamily
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) GetNodeTypeName() *string {
	return s.NodeTypeName
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetCpu(v string) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.Cpu = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetGpu(v string) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.Gpu = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetGpuMemory(v int64) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.GpuMemory = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetMaxCapacity(v int32) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.MaxCapacity = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetMemory(v int64) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.Memory = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetNodeInstanceType(v string) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.NodeInstanceType = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetNodeInstanceTypeFamily(v string) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.NodeInstanceTypeFamily = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetNodeTypeName(v string) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.NodeTypeName = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) Validate() error {
	return dara.Validate(s)
}
