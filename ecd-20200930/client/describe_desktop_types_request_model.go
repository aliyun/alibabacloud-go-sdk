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
	SetBusinessChannel(v string) *DescribeDesktopTypesRequest
	GetBusinessChannel() *string
	SetCpuCount(v int32) *DescribeDesktopTypesRequest
	GetCpuCount() *int32
	SetDesktopGroupIdForModify(v string) *DescribeDesktopTypesRequest
	GetDesktopGroupIdForModify() *string
	SetDesktopIdForModify(v string) *DescribeDesktopTypesRequest
	GetDesktopIdForModify() *string
	SetDesktopScenario(v string) *DescribeDesktopTypesRequest
	GetDesktopScenario() *string
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
	SetOfficeSiteId(v string) *DescribeDesktopTypesRequest
	GetOfficeSiteId() *string
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
	// The scope of the instance types to query. Default value: `Public`.
	//
	// example:
	//
	// Public
	AppliedScope    *string `json:"AppliedScope,omitempty" xml:"AppliedScope,omitempty"`
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 2
	CpuCount *int32 `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	// The ID of the desktop group to reconfigure. If you specify this parameter, the response returns only the instance types that are compatible with the specified group.
	//
	// example:
	//
	// dg-abcdefg****
	DesktopGroupIdForModify *string `json:"DesktopGroupIdForModify,omitempty" xml:"DesktopGroupIdForModify,omitempty"`
	// The ID of the WUYING Workspace to reconfigure. If you specify this parameter, the response returns only the instance types that are compatible with the specified workspace.
	//
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopIdForModify *string `json:"DesktopIdForModify,omitempty" xml:"DesktopIdForModify,omitempty"`
	DesktopScenario    *string `json:"DesktopScenario,omitempty" xml:"DesktopScenario,omitempty"`
	// The ID of the instance type.
	//
	// > If you omit both the `InstanceTypeFamily` and `DesktopTypeId` parameters, the operation returns all available WUYING Workspace instance types.
	//
	// example:
	//
	// ecd.graphics.xlarge
	DesktopTypeId *string `json:"DesktopTypeId,omitempty" xml:"DesktopTypeId,omitempty"`
	// An array of instance type IDs.
	DesktopTypeIdList []*string `json:"DesktopTypeIdList,omitempty" xml:"DesktopTypeIdList,omitempty" type:"Repeated"`
	// The number of vGPUs.
	//
	// example:
	//
	// 1
	GpuCount *float32 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	// The GPU driver type.
	//
	// example:
	//
	// A10
	GpuDriverType *string `json:"GpuDriverType,omitempty" xml:"GpuDriverType,omitempty"`
	GpuMemory     *int32  `json:"GpuMemory,omitempty" xml:"GpuMemory,omitempty"`
	// The instance type family.
	//
	// > If you omit both the `InstanceTypeFamily` and `DesktopTypeId` parameters, the operation returns all available WUYING Workspace instance types.
	//
	// example:
	//
	// ecd.graphics
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	// The memory size, in MiB.
	//
	// example:
	//
	// 4096
	MemorySize   *int32  `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The property by which to sort the results. If you omit this parameter, the results are sorted by creation time in descending order.
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
	// The ID of the region. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the regions that Elastic Desktop Service supports.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The billing method of the instance types.
	//
	// example:
	//
	// FastBuy
	Scope    *string   `json:"Scope,omitempty" xml:"Scope,omitempty"`
	ScopeSet []*string `json:"ScopeSet,omitempty" xml:"ScopeSet,omitempty" type:"Repeated"`
	// The sort order.
	//
	// example:
	//
	// ASC
	SortType *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
	// Filters for instance types that support at least the specified number of concurrent sessions. This parameter applies only to multi-session instance types.
	//
	// example:
	//
	// 2
	SupportMinSessionCount *int32 `json:"SupportMinSessionCount,omitempty" xml:"SupportMinSessionCount,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// 无
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

func (s *DescribeDesktopTypesRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
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

func (s *DescribeDesktopTypesRequest) GetDesktopScenario() *string {
	return s.DesktopScenario
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

func (s *DescribeDesktopTypesRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
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

func (s *DescribeDesktopTypesRequest) SetBusinessChannel(v string) *DescribeDesktopTypesRequest {
	s.BusinessChannel = &v
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

func (s *DescribeDesktopTypesRequest) SetDesktopScenario(v string) *DescribeDesktopTypesRequest {
	s.DesktopScenario = &v
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

func (s *DescribeDesktopTypesRequest) SetOfficeSiteId(v string) *DescribeDesktopTypesRequest {
	s.OfficeSiteId = &v
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
