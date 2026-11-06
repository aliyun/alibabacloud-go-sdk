// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppInstanceTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceTypeModels(v []*ListAppInstanceTypeResponseBodyAppInstanceTypeModels) *ListAppInstanceTypeResponseBody
	GetAppInstanceTypeModels() []*ListAppInstanceTypeResponseBodyAppInstanceTypeModels
	SetPageNumber(v int32) *ListAppInstanceTypeResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAppInstanceTypeResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAppInstanceTypeResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAppInstanceTypeResponseBody
	GetTotalCount() *int32
}

type ListAppInstanceTypeResponseBody struct {
	// The list of session instance types on the current page.
	AppInstanceTypeModels []*ListAppInstanceTypeResponseBodyAppInstanceTypeModels `json:"AppInstanceTypeModels,omitempty" xml:"AppInstanceTypeModels,omitempty" type:"Repeated"`
	// The page number used in this query, which corresponds to the request parameters and the total number of entries.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page used in this query.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique identifier of the request, which can be used to troubleshoot request issues.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of instance types that meet all filtering conditions, not the number of entries on the current page.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAppInstanceTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppInstanceTypeResponseBody) GetAppInstanceTypeModels() []*ListAppInstanceTypeResponseBodyAppInstanceTypeModels {
	return s.AppInstanceTypeModels
}

func (s *ListAppInstanceTypeResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAppInstanceTypeResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppInstanceTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppInstanceTypeResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAppInstanceTypeResponseBody) SetAppInstanceTypeModels(v []*ListAppInstanceTypeResponseBodyAppInstanceTypeModels) *ListAppInstanceTypeResponseBody {
	s.AppInstanceTypeModels = v
	return s
}

func (s *ListAppInstanceTypeResponseBody) SetPageNumber(v int32) *ListAppInstanceTypeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAppInstanceTypeResponseBody) SetPageSize(v int32) *ListAppInstanceTypeResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAppInstanceTypeResponseBody) SetRequestId(v string) *ListAppInstanceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppInstanceTypeResponseBody) SetTotalCount(v int32) *ListAppInstanceTypeResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAppInstanceTypeResponseBody) Validate() error {
	if s.AppInstanceTypeModels != nil {
		for _, item := range s.AppInstanceTypeModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppInstanceTypeResponseBodyAppInstanceTypeModels struct {
	// The session instance type code, which can be used as the session instance type parameter for subsequent operations.
	//
	// example:
	//
	// appstreaming.vgpu.4c8g
	AppInstanceType *string `json:"AppInstanceType,omitempty" xml:"AppInstanceType,omitempty"`
	// The instance type family code to which the session instance type belongs.
	//
	// example:
	//
	// ecs.g6
	AppInstanceTypeFamily *string `json:"AppInstanceTypeFamily,omitempty" xml:"AppInstanceTypeFamily,omitempty"`
	// The number of CPU cores included in the instance type, represented by a numeric string. Unit: cores.
	//
	// example:
	//
	// 4
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The data cloud disk capacity of the instance type. Unit: GB.
	//
	// example:
	//
	// 100
	DataDiskSize *int64 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	// Not enabled yet. The example value is only used to demonstrate the format and does not represent the actual number of GPUs returned by this operation.
	//
	// example:
	//
	// 1
	Gpu *string `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The GPU memory capacity of the instance type. Unit: MB.
	//
	// example:
	//
	// 4096
	GpuMemory *int64 `json:"GpuMemory,omitempty" xml:"GpuMemory,omitempty"`
	// The display value or display range of the GPU memory. Display it as the returned text. Do not directly use it as a numeric value for capacity conversion.
	//
	// example:
	//
	// 1~2
	GpuSpec *string `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty"`
	// The identifier of the instance type record. When integrating with pricing or creation operations, use `AppInstanceType`. Do not use this field to replace the instance type code.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The display name of the instance type. The Chinese name is used only when `Language` is set to `zh-CN`. For other values or when omitted, the instance type code is used.
	//
	// example:
	//
	// appstreaming.vgpu.4c8g
	InstanceTypeName *string `json:"InstanceTypeName,omitempty" xml:"InstanceTypeName,omitempty"`
	// The inventory information of the instance type. This parameter is returned only when the inventory is queried and the relevant information is obtained. If this parameter is not returned, it does not mean that the inventory is 0.
	InventoryInfo *ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfo `json:"InventoryInfo,omitempty" xml:"InventoryInfo,omitempty" type:"Struct"`
	// The memory capacity of the instance type. Unit: MB.
	//
	// example:
	//
	// 8192
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The type of the product to which the instance type belongs. The current document displays the following values:
	//
	// - `CloudApp`: Wuying Cloud Application.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The system cloud disk capacity of the instance type. Unit: GB.
	//
	// example:
	//
	// 40
	SystemDiskSize *int64 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s ListAppInstanceTypeResponseBodyAppInstanceTypeModels) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceTypeResponseBodyAppInstanceTypeModels) GoString() string {
	return s.String()
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) GetAppInstanceType() *string {
	return s.AppInstanceType
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) GetAppInstanceTypeFamily() *string {
	return s.AppInstanceTypeFamily
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) GetCpu() *string {
	return s.Cpu
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) GetDataDiskSize() *int64 {
	return s.DataDiskSize
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) GetGpu() *string {
	return s.Gpu
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) GetGpuMemory() *int64 {
	return s.GpuMemory
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) GetGpuSpec() *string {
	return s.GpuSpec
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) GetId() *int64 {
	return s.Id
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) GetInstanceTypeName() *string {
	return s.InstanceTypeName
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) GetInventoryInfo() *ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfo {
	return s.InventoryInfo
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) GetMemory() *int64 {
	return s.Memory
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) GetProductType() *string {
	return s.ProductType
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) GetSystemDiskSize() *int64 {
	return s.SystemDiskSize
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) SetAppInstanceType(v string) *ListAppInstanceTypeResponseBodyAppInstanceTypeModels {
	s.AppInstanceType = &v
	return s
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) SetAppInstanceTypeFamily(v string) *ListAppInstanceTypeResponseBodyAppInstanceTypeModels {
	s.AppInstanceTypeFamily = &v
	return s
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) SetCpu(v string) *ListAppInstanceTypeResponseBodyAppInstanceTypeModels {
	s.Cpu = &v
	return s
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) SetDataDiskSize(v int64) *ListAppInstanceTypeResponseBodyAppInstanceTypeModels {
	s.DataDiskSize = &v
	return s
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) SetGpu(v string) *ListAppInstanceTypeResponseBodyAppInstanceTypeModels {
	s.Gpu = &v
	return s
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) SetGpuMemory(v int64) *ListAppInstanceTypeResponseBodyAppInstanceTypeModels {
	s.GpuMemory = &v
	return s
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) SetGpuSpec(v string) *ListAppInstanceTypeResponseBodyAppInstanceTypeModels {
	s.GpuSpec = &v
	return s
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) SetId(v int64) *ListAppInstanceTypeResponseBodyAppInstanceTypeModels {
	s.Id = &v
	return s
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) SetInstanceTypeName(v string) *ListAppInstanceTypeResponseBodyAppInstanceTypeModels {
	s.InstanceTypeName = &v
	return s
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) SetInventoryInfo(v *ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfo) *ListAppInstanceTypeResponseBodyAppInstanceTypeModels {
	s.InventoryInfo = v
	return s
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) SetMemory(v int64) *ListAppInstanceTypeResponseBodyAppInstanceTypeModels {
	s.Memory = &v
	return s
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) SetProductType(v string) *ListAppInstanceTypeResponseBodyAppInstanceTypeModels {
	s.ProductType = &v
	return s
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) SetSystemDiskSize(v int64) *ListAppInstanceTypeResponseBodyAppInstanceTypeModels {
	s.SystemDiskSize = &v
	return s
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModels) Validate() error {
	if s.InventoryInfo != nil {
		if err := s.InventoryInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfo struct {
	// The business region ID corresponding to the inventory query.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance type inventory status obtained in the current query.
	//
	// - `Available`: The total inventory of the queried zones is greater than 0.
	//
	// - `SoldOut`: The total inventory of the queried zones is not greater than 0.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of inventory information returned by zone.
	ZoneInventoryInfo []*ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfoZoneInventoryInfo `json:"ZoneInventoryInfo,omitempty" xml:"ZoneInventoryInfo,omitempty" type:"Repeated"`
}

func (s ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfo) GoString() string {
	return s.String()
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfo) GetStatus() *string {
	return s.Status
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfo) GetZoneInventoryInfo() []*ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfoZoneInventoryInfo {
	return s.ZoneInventoryInfo
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfo) SetRegionId(v string) *ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfo {
	s.RegionId = &v
	return s
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfo) SetStatus(v string) *ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfo {
	s.Status = &v
	return s
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfo) SetZoneInventoryInfo(v []*ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfoZoneInventoryInfo) *ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfo {
	s.ZoneInventoryInfo = v
	return s
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfo) Validate() error {
	if s.ZoneInventoryInfo != nil {
		for _, item := range s.ZoneInventoryInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfoZoneInventoryInfo struct {
	// The available quantity in the zone obtained in this query. Unit: instances. This value is an inventory snapshot and does not indicate that resources have been reserved for the current account.
	//
	// example:
	//
	// 10
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The zone ID, which can be queried by calling the [ListZones](https://api.aliyun.com/document/appstream-center/2021-09-01/ListZones) operation.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfoZoneInventoryInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfoZoneInventoryInfo) GoString() string {
	return s.String()
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfoZoneInventoryInfo) GetAmount() *string {
	return s.Amount
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfoZoneInventoryInfo) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfoZoneInventoryInfo) SetAmount(v string) *ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfoZoneInventoryInfo {
	s.Amount = &v
	return s
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfoZoneInventoryInfo) SetZoneId(v string) *ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfoZoneInventoryInfo {
	s.ZoneId = &v
	return s
}

func (s *ListAppInstanceTypeResponseBodyAppInstanceTypeModelsInventoryInfoZoneInventoryInfo) Validate() error {
	return dara.Validate(s)
}
