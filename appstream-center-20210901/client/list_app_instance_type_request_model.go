// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppInstanceTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCenterImageId(v string) *ListAppInstanceTypeRequest
	GetAppCenterImageId() *string
	SetAppInstanceType(v string) *ListAppInstanceTypeRequest
	GetAppInstanceType() *string
	SetAppInstanceTypeFamily(v string) *ListAppInstanceTypeRequest
	GetAppInstanceTypeFamily() *string
	SetAppInstanceTypeList(v []*string) *ListAppInstanceTypeRequest
	GetAppInstanceTypeList() []*string
	SetBizRegionId(v string) *ListAppInstanceTypeRequest
	GetBizRegionId() *string
	SetCategory(v string) *ListAppInstanceTypeRequest
	GetCategory() *string
	SetChargeType(v string) *ListAppInstanceTypeRequest
	GetChargeType() *string
	SetCpu(v string) *ListAppInstanceTypeRequest
	GetCpu() *string
	SetGpuMemory(v int32) *ListAppInstanceTypeRequest
	GetGpuMemory() *int32
	SetLanguage(v string) *ListAppInstanceTypeRequest
	GetLanguage() *string
	SetMemory(v int32) *ListAppInstanceTypeRequest
	GetMemory() *int32
	SetOsType(v string) *ListAppInstanceTypeRequest
	GetOsType() *string
	SetPageNumber(v int32) *ListAppInstanceTypeRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAppInstanceTypeRequest
	GetPageSize() *int32
	SetProductType(v string) *ListAppInstanceTypeRequest
	GetProductType() *string
	SetQueryInventory(v bool) *ListAppInstanceTypeRequest
	GetQueryInventory() *bool
}

type ListAppInstanceTypeRequest struct {
	// The application image ID. The current session instance type query does not filter based on this parameter. Omit this parameter. Do not infer image compatibility based on whether this parameter is specified.
	//
	// example:
	//
	// img-8z4nztpaqvay4****
	AppCenterImageId *string `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	// The session instance type code. Exact match is performed based on the complete code. If this parameter and `AppInstanceTypeList` are specified at the same time, the results are filtered by intersection. You can perform a conditional query without this condition first, and then select the returned codes.
	//
	// example:
	//
	// appstreaming.vgpu.4c8g
	AppInstanceType *string `json:"AppInstanceType,omitempty" xml:"AppInstanceType,omitempty"`
	// The instance type family code. Exact match is performed based on the complete code. Use the instance type family value from the query results. Do not extract it from the instance type name.
	//
	// example:
	//
	// ecs.g6
	AppInstanceTypeFamily *string `json:"AppInstanceTypeFamily,omitempty" xml:"AppInstanceTypeFamily,omitempty"`
	// The list of session instance type codes. A maximum of 100 items are supported. This parameter takes effect by intersection with other filtering conditions. If this parameter and `AppInstanceType` are specified at the same time, the results must also match the single code.
	//
	// if can be null:
	// true
	AppInstanceTypeList []*string `json:"AppInstanceTypeList,omitempty" xml:"AppInstanceTypeList,omitempty" type:"Repeated"`
	// The business region ID. If this parameter is omitted, the region of the current request is used. The available instance types may vary by region and account. You can call the [ListRegions](https://api.aliyun.com/document/appstream-center/2021-09-01/ListRegions) operation to obtain the business region ID.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The instance type category filtering condition. If this parameter is omitted, no category filtering condition is added. Valid values:
	//
	// - `Cpu`: CPU instance type.
	//
	// - `Gpu`: GPU instance type.
	//
	// example:
	//
	// Gpu
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The billing method used to filter available instance types. If this parameter is omitted, no billing method condition is added. Valid values:
	//
	// - `PrePaid`: subscription.
	//
	// - `PostPaid`: pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The number of CPU cores for exact filtering. Use a decimal numeric string. Unit: cores. If this parameter is omitted, the number of CPU cores is not limited.
	//
	// example:
	//
	// 4
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The GPU memory capacity for exact filtering. Unit: MB. If this parameter is omitted, the GPU memory capacity is not limited.
	//
	// example:
	//
	// 4096
	GpuMemory *int32 `json:"GpuMemory,omitempty" xml:"GpuMemory,omitempty"`
	// The language of the instance type name.
	//
	// - `zh-CN`: Use the Chinese name of the instance type.
	//
	// - Other values or omitted: Use the instance type code as the name.
	//
	// This parameter only affects the display of `InstanceTypeName` and does not change the instance type code.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The memory capacity for exact filtering. Unit: MB. If this parameter is omitted, the memory capacity is not limited.
	//
	// example:
	//
	// 8192
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The operating system type. In the Wuying Cloud Application (CloudApp) scenario, set the value to:
	//
	// - `Windows`: Windows operating system.
	//
	// If this parameter is omitted, no operating system filtering condition is added.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The page number. Explicitly specify a page number starting from 1. If this parameter is omitted, the default value 0 is used to query the first batch of results. For subsequent pages, request in the order of 1, 2, and 3.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of instance types returned per page. Unit: entries. Default value: 200. Explicitly set a positive integer that meets your business requirements. The default value here does not represent the maximum value allowed by the service.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type. Set the value to the following fixed value when you call this operation:
	//
	// - `CloudApp`: Wuying Cloud Application.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// Specifies whether to query the instance type inventory at the same time. Default value: `false`.
	//
	// - `true`: Queries the inventory and returns InventoryInfo when the inventory information is available.
	//
	// - `false`: Queries only the instance types and does not query the inventory.
	//
	// example:
	//
	// false
	QueryInventory *bool `json:"QueryInventory,omitempty" xml:"QueryInventory,omitempty"`
}

func (s ListAppInstanceTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceTypeRequest) GoString() string {
	return s.String()
}

func (s *ListAppInstanceTypeRequest) GetAppCenterImageId() *string {
	return s.AppCenterImageId
}

func (s *ListAppInstanceTypeRequest) GetAppInstanceType() *string {
	return s.AppInstanceType
}

func (s *ListAppInstanceTypeRequest) GetAppInstanceTypeFamily() *string {
	return s.AppInstanceTypeFamily
}

func (s *ListAppInstanceTypeRequest) GetAppInstanceTypeList() []*string {
	return s.AppInstanceTypeList
}

func (s *ListAppInstanceTypeRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *ListAppInstanceTypeRequest) GetCategory() *string {
	return s.Category
}

func (s *ListAppInstanceTypeRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListAppInstanceTypeRequest) GetCpu() *string {
	return s.Cpu
}

func (s *ListAppInstanceTypeRequest) GetGpuMemory() *int32 {
	return s.GpuMemory
}

func (s *ListAppInstanceTypeRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListAppInstanceTypeRequest) GetMemory() *int32 {
	return s.Memory
}

func (s *ListAppInstanceTypeRequest) GetOsType() *string {
	return s.OsType
}

func (s *ListAppInstanceTypeRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAppInstanceTypeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppInstanceTypeRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListAppInstanceTypeRequest) GetQueryInventory() *bool {
	return s.QueryInventory
}

func (s *ListAppInstanceTypeRequest) SetAppCenterImageId(v string) *ListAppInstanceTypeRequest {
	s.AppCenterImageId = &v
	return s
}

func (s *ListAppInstanceTypeRequest) SetAppInstanceType(v string) *ListAppInstanceTypeRequest {
	s.AppInstanceType = &v
	return s
}

func (s *ListAppInstanceTypeRequest) SetAppInstanceTypeFamily(v string) *ListAppInstanceTypeRequest {
	s.AppInstanceTypeFamily = &v
	return s
}

func (s *ListAppInstanceTypeRequest) SetAppInstanceTypeList(v []*string) *ListAppInstanceTypeRequest {
	s.AppInstanceTypeList = v
	return s
}

func (s *ListAppInstanceTypeRequest) SetBizRegionId(v string) *ListAppInstanceTypeRequest {
	s.BizRegionId = &v
	return s
}

func (s *ListAppInstanceTypeRequest) SetCategory(v string) *ListAppInstanceTypeRequest {
	s.Category = &v
	return s
}

func (s *ListAppInstanceTypeRequest) SetChargeType(v string) *ListAppInstanceTypeRequest {
	s.ChargeType = &v
	return s
}

func (s *ListAppInstanceTypeRequest) SetCpu(v string) *ListAppInstanceTypeRequest {
	s.Cpu = &v
	return s
}

func (s *ListAppInstanceTypeRequest) SetGpuMemory(v int32) *ListAppInstanceTypeRequest {
	s.GpuMemory = &v
	return s
}

func (s *ListAppInstanceTypeRequest) SetLanguage(v string) *ListAppInstanceTypeRequest {
	s.Language = &v
	return s
}

func (s *ListAppInstanceTypeRequest) SetMemory(v int32) *ListAppInstanceTypeRequest {
	s.Memory = &v
	return s
}

func (s *ListAppInstanceTypeRequest) SetOsType(v string) *ListAppInstanceTypeRequest {
	s.OsType = &v
	return s
}

func (s *ListAppInstanceTypeRequest) SetPageNumber(v int32) *ListAppInstanceTypeRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppInstanceTypeRequest) SetPageSize(v int32) *ListAppInstanceTypeRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppInstanceTypeRequest) SetProductType(v string) *ListAppInstanceTypeRequest {
	s.ProductType = &v
	return s
}

func (s *ListAppInstanceTypeRequest) SetQueryInventory(v bool) *ListAppInstanceTypeRequest {
	s.QueryInventory = &v
	return s
}

func (s *ListAppInstanceTypeRequest) Validate() error {
	return dara.Validate(s)
}
