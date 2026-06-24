// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeInstanceTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizRegionId(v string) *ListNodeInstanceTypeRequest
	GetBizRegionId() *string
	SetCpu(v float32) *ListNodeInstanceTypeRequest
	GetCpu() *float32
	SetGpu(v float32) *ListNodeInstanceTypeRequest
	GetGpu() *float32
	SetGpuMemory(v int32) *ListNodeInstanceTypeRequest
	GetGpuMemory() *int32
	SetInstanceTypeForModify(v string) *ListNodeInstanceTypeRequest
	GetInstanceTypeForModify() *string
	SetLanguage(v string) *ListNodeInstanceTypeRequest
	GetLanguage() *string
	SetMemory(v int32) *ListNodeInstanceTypeRequest
	GetMemory() *int32
	SetNodeInstanceType(v string) *ListNodeInstanceTypeRequest
	GetNodeInstanceType() *string
	SetNodeInstanceTypeFamily(v string) *ListNodeInstanceTypeRequest
	GetNodeInstanceTypeFamily() *string
	SetOrderBy(v string) *ListNodeInstanceTypeRequest
	GetOrderBy() *string
	SetOrderType(v string) *ListNodeInstanceTypeRequest
	GetOrderType() *string
	SetOsType(v string) *ListNodeInstanceTypeRequest
	GetOsType() *string
	SetPageNumber(v int32) *ListNodeInstanceTypeRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListNodeInstanceTypeRequest
	GetPageSize() *int32
	SetProductType(v string) *ListNodeInstanceTypeRequest
	GetProductType() *string
	SetSortType(v string) *ListNodeInstanceTypeRequest
	GetSortType() *string
}

type ListNodeInstanceTypeRequest struct {
	// The region ID of the resource. For more information about supported regions, see [Limits](https://help.aliyun.com/document_detail/426036.html).
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The number of CPU cores.
	//
	// example:
	//
	// 2
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The number of GPUs.
	//
	// example:
	//
	// 1
	Gpu *float32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The GPU memory size. This parameter is meaningful only for GPU-accelerated cloud desktops. Unit: MB.
	//
	// example:
	//
	// 2048
	GpuMemory             *int32  `json:"GpuMemory,omitempty" xml:"GpuMemory,omitempty"`
	InstanceTypeForModify *string `json:"InstanceTypeForModify,omitempty" xml:"InstanceTypeForModify,omitempty"`
	// The language type.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The memory size. Unit: MB.
	//
	// example:
	//
	// 10240
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The resource specification type to query. If you leave this parameter empty, all specification types are returned.
	//
	// example:
	//
	// appstreaming.vgpu.4c8g.2g
	NodeInstanceType *string `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	// The instance family.
	//
	// example:
	//
	// appstreaming.vgpu
	NodeInstanceTypeFamily *string `json:"NodeInstanceTypeFamily,omitempty" xml:"NodeInstanceTypeFamily,omitempty"`
	// CPU/Memory.
	//
	// example:
	//
	// CPU
	OrderBy   *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The supported operating system type.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The page number of the query results to display.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of query results per page. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// DESC/ASC.
	//
	// example:
	//
	// ASC
	SortType *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
}

func (s ListNodeInstanceTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodeInstanceTypeRequest) GoString() string {
	return s.String()
}

func (s *ListNodeInstanceTypeRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *ListNodeInstanceTypeRequest) GetCpu() *float32 {
	return s.Cpu
}

func (s *ListNodeInstanceTypeRequest) GetGpu() *float32 {
	return s.Gpu
}

func (s *ListNodeInstanceTypeRequest) GetGpuMemory() *int32 {
	return s.GpuMemory
}

func (s *ListNodeInstanceTypeRequest) GetInstanceTypeForModify() *string {
	return s.InstanceTypeForModify
}

func (s *ListNodeInstanceTypeRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListNodeInstanceTypeRequest) GetMemory() *int32 {
	return s.Memory
}

func (s *ListNodeInstanceTypeRequest) GetNodeInstanceType() *string {
	return s.NodeInstanceType
}

func (s *ListNodeInstanceTypeRequest) GetNodeInstanceTypeFamily() *string {
	return s.NodeInstanceTypeFamily
}

func (s *ListNodeInstanceTypeRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListNodeInstanceTypeRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListNodeInstanceTypeRequest) GetOsType() *string {
	return s.OsType
}

func (s *ListNodeInstanceTypeRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNodeInstanceTypeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNodeInstanceTypeRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListNodeInstanceTypeRequest) GetSortType() *string {
	return s.SortType
}

func (s *ListNodeInstanceTypeRequest) SetBizRegionId(v string) *ListNodeInstanceTypeRequest {
	s.BizRegionId = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetCpu(v float32) *ListNodeInstanceTypeRequest {
	s.Cpu = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetGpu(v float32) *ListNodeInstanceTypeRequest {
	s.Gpu = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetGpuMemory(v int32) *ListNodeInstanceTypeRequest {
	s.GpuMemory = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetInstanceTypeForModify(v string) *ListNodeInstanceTypeRequest {
	s.InstanceTypeForModify = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetLanguage(v string) *ListNodeInstanceTypeRequest {
	s.Language = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetMemory(v int32) *ListNodeInstanceTypeRequest {
	s.Memory = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetNodeInstanceType(v string) *ListNodeInstanceTypeRequest {
	s.NodeInstanceType = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetNodeInstanceTypeFamily(v string) *ListNodeInstanceTypeRequest {
	s.NodeInstanceTypeFamily = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetOrderBy(v string) *ListNodeInstanceTypeRequest {
	s.OrderBy = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetOrderType(v string) *ListNodeInstanceTypeRequest {
	s.OrderType = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetOsType(v string) *ListNodeInstanceTypeRequest {
	s.OsType = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetPageNumber(v int32) *ListNodeInstanceTypeRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetPageSize(v int32) *ListNodeInstanceTypeRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetProductType(v string) *ListNodeInstanceTypeRequest {
	s.ProductType = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetSortType(v string) *ListNodeInstanceTypeRequest {
	s.SortType = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) Validate() error {
	return dara.Validate(s)
}
