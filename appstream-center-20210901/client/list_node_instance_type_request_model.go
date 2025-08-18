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
	// The ID of the region where the resource resides. For information about the supported regions, see [Limits](https://help.aliyun.com/document_detail/426036.html).
	//
	// Valid values:
	//
	// 	- cn-shanghai: China (Shanghai)
	//
	// 	- cn-hangzhou: China (Hangzhou)
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string  `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	Cpu         *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu         *float32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	GpuMemory   *int32   `json:"GpuMemory,omitempty" xml:"GpuMemory,omitempty"`
	// The language that you want to use.
	//
	// Valid values:
	//
	// 	- en-US: English (US)
	//
	// 	- zh-CN: Simplified Chinese
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Memory   *int32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The resource type that you want to query. If you do not configure this parameter, all resource types are returned.
	//
	// example:
	//
	// appstreaming.vgpu.4c8g.2g
	NodeInstanceType       *string `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	NodeInstanceTypeFamily *string `json:"NodeInstanceTypeFamily,omitempty" xml:"NodeInstanceTypeFamily,omitempty"`
	OrderBy                *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The operating system that is supported.
	//
	// Valid value:
	//
	// 	- Windows: the Windows operating system
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SortType    *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
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
