// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppliedScope(v string) *DescribeDesktopTypesRequest
	GetAppliedScope() *string
	SetCpuCount(v int32) *DescribeDesktopTypesRequest
	GetCpuCount() *int32
	SetDesktopGroupIdForModify(v string) *DescribeDesktopTypesRequest
	GetDesktopGroupIdForModify() *string
	SetDesktopIdForModify(v string) *DescribeDesktopTypesRequest
	GetDesktopIdForModify() *string
	SetDesktopTypeId(v string) *DescribeDesktopTypesRequest
	GetDesktopTypeId() *string
	SetDesktopTypeIdList(v []*string) *DescribeDesktopTypesRequest
	GetDesktopTypeIdList() []*string
	SetGpuCount(v float32) *DescribeDesktopTypesRequest
	GetGpuCount() *float32
	SetGpuDriverType(v string) *DescribeDesktopTypesRequest
	GetGpuDriverType() *string
	SetGpuMemory(v int32) *DescribeDesktopTypesRequest
	GetGpuMemory() *int32
	SetInstanceTypeFamily(v string) *DescribeDesktopTypesRequest
	GetInstanceTypeFamily() *string
	SetMemorySize(v int32) *DescribeDesktopTypesRequest
	GetMemorySize() *int32
	SetOrderBy(v string) *DescribeDesktopTypesRequest
	GetOrderBy() *string
	SetOrderType(v string) *DescribeDesktopTypesRequest
	GetOrderType() *string
	SetRegionId(v string) *DescribeDesktopTypesRequest
	GetRegionId() *string
	SetScope(v string) *DescribeDesktopTypesRequest
	GetScope() *string
	SetScopeSet(v []*string) *DescribeDesktopTypesRequest
	GetScopeSet() []*string
	SetSortType(v string) *DescribeDesktopTypesRequest
	GetSortType() *string
	SetSupportMinSessionCount(v int32) *DescribeDesktopTypesRequest
	GetSupportMinSessionCount() *int32
	SetZoneId(v string) *DescribeDesktopTypesRequest
	GetZoneId() *string
}

type DescribeDesktopTypesRequest struct {
	// Applicable Scope of specifications. Default value: `Public`
	//
	// example:
	//
	// Public
	AppliedScope *string `json:"AppliedScope,omitempty" xml:"AppliedScope,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 2
	CpuCount *int32 `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	// The ID of the cloud computer share you want to modify. If this parameter is provided, the response will include compatibility information for the specified specification.
	//
	// example:
	//
	// dg-abcdefg****
	DesktopGroupIdForModify *string `json:"DesktopGroupIdForModify,omitempty" xml:"DesktopGroupIdForModify,omitempty"`
	// The ID of the cloud computer when you change instance types of cloud computers. If you specify this parameter, the information about whether the instance type is compatible with the cloud computer is included in the response.
	//
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopIdForModify *string `json:"DesktopIdForModify,omitempty" xml:"DesktopIdForModify,omitempty"`
	// The specification ID.
	//
	// >  If both `InstanceTypeFamily` and `DesktopTypeId` are empty, all cloud computer specifications will be queried.
	//
	// Valid values:
	//
	// 	- eds.enterprise_office.4c8g
	//
	// 	- eds.hf.4c8g
	//
	// 	- ecd.basic.large
	//
	// 	- ecd.advanced.large
	//
	// 	- eds.enterprise_office.8c16g
	//
	// 	- ecd.basic.small
	//
	// 	- ecd.graphics.2xlarge
	//
	// 	- eds.hf.8c16g
	//
	// 	- eds.hf.12c24g
	//
	// 	- eds.general.8c16g
	//
	// 	- eds.general.16c32g
	//
	// 	- ecd.advanced.xlarge
	//
	// 	- eds.graphics.16c1t4
	//
	// 	- ecd.graphics.xlarge
	//
	// 	- ecd.performance.2xlarge
	//
	// 	- eds.general.8c32g
	//
	// 	- eds.general.2c2g
	//
	// 	- eds.general.2c4g
	//
	// 	- eds.graphics.24c1t4
	//
	// 	- eds.general.4c8g
	//
	// 	- eds.enterprise_office.2c4g
	//
	// 	- eds.general.4c16g
	//
	// 	- eds.general.2c8g
	//
	// example:
	//
	// ecd.graphics.xlarge
	DesktopTypeId *string `json:"DesktopTypeId,omitempty" xml:"DesktopTypeId,omitempty"`
	// The specification IDs.
	DesktopTypeIdList []*string `json:"DesktopTypeIdList,omitempty" xml:"DesktopTypeIdList,omitempty" type:"Repeated"`
	// The number of GPUs.
	//
	// example:
	//
	// 1
	GpuCount *float32 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	// The GPU driver type.
	//
	// Valid values:
	//
	// 	- T4
	//
	// 	- A10
	//
	// 	- G28
	//
	// 	- G39
	//
	// example:
	//
	// A10
	GpuDriverType *string `json:"GpuDriverType,omitempty" xml:"GpuDriverType,omitempty"`
	// The GPU memory size. Unit: MB.
	//
	// example:
	//
	// 2048
	GpuMemory *int32 `json:"GpuMemory,omitempty" xml:"GpuMemory,omitempty"`
	// The name of the specification family.
	//
	// >  If both `InstanceTypeFamily` and `DesktopTypeId` are empty, all specification families will be queried.
	//
	// Valid values:
	//
	// 	- ecd.advanced
	//
	// 	- eds.graphics
	//
	// 	- ecd.basic
	//
	// 	- eds.enterprise_office
	//
	// 	- eds.hf
	//
	// 	- ecd.graphics
	//
	// 	- eds.general
	//
	// 	- ecd.performance
	//
	// example:
	//
	// ecd.graphics
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	// The memory size. Unit: MiB.
	//
	// example:
	//
	// 4
	MemorySize *int32 `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	// The sorting field. If this parameter is not provided, results are sorted by creation time in descending order.
	//
	// Valid values:
	//
	// 	- Memory: sorts by memory size.
	//
	// 	- Cpu: sorts by the number of vCPUs.
	//
	// example:
	//
	// Memory
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The order type.
	//
	// example:
	//
	// DOWNGRADE
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The sales mode of the specification.
	//
	// Valid values:
	//
	// 	- MonthPackage: the monthly subscription mode.
	//
	// 	- FastBuy: the quick purchase mode.
	//
	// example:
	//
	// FastBuy
	Scope    *string   `json:"Scope,omitempty" xml:"Scope,omitempty"`
	ScopeSet []*string `json:"ScopeSet,omitempty" xml:"ScopeSet,omitempty" type:"Repeated"`
	// The sorting order.
	//
	// Valid values:
	//
	// 	- ASC (default): the ascending order.
	//
	// 	- DESC: the descending order.
	//
	// example:
	//
	// ASC
	SortType *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
	// The number of sessions supported by the specification.
	//
	// example:
	//
	// 2
	SupportMinSessionCount *int32 `json:"SupportMinSessionCount,omitempty" xml:"SupportMinSessionCount,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// null
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDesktopTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopTypesRequest) GetAppliedScope() *string {
	return s.AppliedScope
}

func (s *DescribeDesktopTypesRequest) GetCpuCount() *int32 {
	return s.CpuCount
}

func (s *DescribeDesktopTypesRequest) GetDesktopGroupIdForModify() *string {
	return s.DesktopGroupIdForModify
}

func (s *DescribeDesktopTypesRequest) GetDesktopIdForModify() *string {
	return s.DesktopIdForModify
}

func (s *DescribeDesktopTypesRequest) GetDesktopTypeId() *string {
	return s.DesktopTypeId
}

func (s *DescribeDesktopTypesRequest) GetDesktopTypeIdList() []*string {
	return s.DesktopTypeIdList
}

func (s *DescribeDesktopTypesRequest) GetGpuCount() *float32 {
	return s.GpuCount
}

func (s *DescribeDesktopTypesRequest) GetGpuDriverType() *string {
	return s.GpuDriverType
}

func (s *DescribeDesktopTypesRequest) GetGpuMemory() *int32 {
	return s.GpuMemory
}

func (s *DescribeDesktopTypesRequest) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *DescribeDesktopTypesRequest) GetMemorySize() *int32 {
	return s.MemorySize
}

func (s *DescribeDesktopTypesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeDesktopTypesRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeDesktopTypesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDesktopTypesRequest) GetScope() *string {
	return s.Scope
}

func (s *DescribeDesktopTypesRequest) GetScopeSet() []*string {
	return s.ScopeSet
}

func (s *DescribeDesktopTypesRequest) GetSortType() *string {
	return s.SortType
}

func (s *DescribeDesktopTypesRequest) GetSupportMinSessionCount() *int32 {
	return s.SupportMinSessionCount
}

func (s *DescribeDesktopTypesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDesktopTypesRequest) SetAppliedScope(v string) *DescribeDesktopTypesRequest {
	s.AppliedScope = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetCpuCount(v int32) *DescribeDesktopTypesRequest {
	s.CpuCount = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetDesktopGroupIdForModify(v string) *DescribeDesktopTypesRequest {
	s.DesktopGroupIdForModify = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetDesktopIdForModify(v string) *DescribeDesktopTypesRequest {
	s.DesktopIdForModify = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetDesktopTypeId(v string) *DescribeDesktopTypesRequest {
	s.DesktopTypeId = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetDesktopTypeIdList(v []*string) *DescribeDesktopTypesRequest {
	s.DesktopTypeIdList = v
	return s
}

func (s *DescribeDesktopTypesRequest) SetGpuCount(v float32) *DescribeDesktopTypesRequest {
	s.GpuCount = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetGpuDriverType(v string) *DescribeDesktopTypesRequest {
	s.GpuDriverType = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetGpuMemory(v int32) *DescribeDesktopTypesRequest {
	s.GpuMemory = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetInstanceTypeFamily(v string) *DescribeDesktopTypesRequest {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetMemorySize(v int32) *DescribeDesktopTypesRequest {
	s.MemorySize = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetOrderBy(v string) *DescribeDesktopTypesRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetOrderType(v string) *DescribeDesktopTypesRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetRegionId(v string) *DescribeDesktopTypesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetScope(v string) *DescribeDesktopTypesRequest {
	s.Scope = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetScopeSet(v []*string) *DescribeDesktopTypesRequest {
	s.ScopeSet = v
	return s
}

func (s *DescribeDesktopTypesRequest) SetSortType(v string) *DescribeDesktopTypesRequest {
	s.SortType = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetSupportMinSessionCount(v int32) *DescribeDesktopTypesRequest {
	s.SupportMinSessionCount = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetZoneId(v string) *DescribeDesktopTypesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeDesktopTypesRequest) Validate() error {
	return dara.Validate(s)
}
