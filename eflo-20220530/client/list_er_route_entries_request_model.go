// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListErRouteEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationCidrBlock(v string) *ListErRouteEntriesRequest
	GetDestinationCidrBlock() *string
	SetEnablePage(v bool) *ListErRouteEntriesRequest
	GetEnablePage() *bool
	SetErId(v string) *ListErRouteEntriesRequest
	GetErId() *string
	SetIgnoreDetailedRouteEntry(v bool) *ListErRouteEntriesRequest
	GetIgnoreDetailedRouteEntry() *bool
	SetInstanceId(v string) *ListErRouteEntriesRequest
	GetInstanceId() *string
	SetNextHopId(v string) *ListErRouteEntriesRequest
	GetNextHopId() *string
	SetNextHopType(v string) *ListErRouteEntriesRequest
	GetNextHopType() *string
	SetPageNumber(v int32) *ListErRouteEntriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListErRouteEntriesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListErRouteEntriesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListErRouteEntriesRequest
	GetResourceGroupId() *string
	SetRouteType(v string) *ListErRouteEntriesRequest
	GetRouteType() *string
	SetStatus(v string) *ListErRouteEntriesRequest
	GetStatus() *string
}

type ListErRouteEntriesRequest struct {
	// Destination CIDR Block
	//
	// example:
	//
	// 0.0.0.0/0
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Specifies whether to enable pagination query.
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// Lingjun HUB ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Filter 32 detailed CIDR blocks. Default value: true
	//
	// example:
	//
	// true
	IgnoreDetailedRouteEntry *bool `json:"IgnoreDetailedRouteEntry,omitempty" xml:"IgnoreDetailedRouteEntry,omitempty"`
	// Network Instance ID
	//
	// example:
	//
	// vcc-cn-209300qha01
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Next Hop Instance
	//
	// example:
	//
	// vcc-cn-209300qha01
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// Next Hop Instance Type
	//
	// example:
	//
	// VCC
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The page number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-acfmyuzlx2iihcy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Route type
	//
	// example:
	//
	// VCC
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	// The status of the enterprise-level snapshot policy.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListErRouteEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListErRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *ListErRouteEntriesRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *ListErRouteEntriesRequest) GetEnablePage() *bool {
	return s.EnablePage
}

func (s *ListErRouteEntriesRequest) GetErId() *string {
	return s.ErId
}

func (s *ListErRouteEntriesRequest) GetIgnoreDetailedRouteEntry() *bool {
	return s.IgnoreDetailedRouteEntry
}

func (s *ListErRouteEntriesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListErRouteEntriesRequest) GetNextHopId() *string {
	return s.NextHopId
}

func (s *ListErRouteEntriesRequest) GetNextHopType() *string {
	return s.NextHopType
}

func (s *ListErRouteEntriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListErRouteEntriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListErRouteEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListErRouteEntriesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListErRouteEntriesRequest) GetRouteType() *string {
	return s.RouteType
}

func (s *ListErRouteEntriesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListErRouteEntriesRequest) SetDestinationCidrBlock(v string) *ListErRouteEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ListErRouteEntriesRequest) SetEnablePage(v bool) *ListErRouteEntriesRequest {
	s.EnablePage = &v
	return s
}

func (s *ListErRouteEntriesRequest) SetErId(v string) *ListErRouteEntriesRequest {
	s.ErId = &v
	return s
}

func (s *ListErRouteEntriesRequest) SetIgnoreDetailedRouteEntry(v bool) *ListErRouteEntriesRequest {
	s.IgnoreDetailedRouteEntry = &v
	return s
}

func (s *ListErRouteEntriesRequest) SetInstanceId(v string) *ListErRouteEntriesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListErRouteEntriesRequest) SetNextHopId(v string) *ListErRouteEntriesRequest {
	s.NextHopId = &v
	return s
}

func (s *ListErRouteEntriesRequest) SetNextHopType(v string) *ListErRouteEntriesRequest {
	s.NextHopType = &v
	return s
}

func (s *ListErRouteEntriesRequest) SetPageNumber(v int32) *ListErRouteEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListErRouteEntriesRequest) SetPageSize(v int32) *ListErRouteEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListErRouteEntriesRequest) SetRegionId(v string) *ListErRouteEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *ListErRouteEntriesRequest) SetResourceGroupId(v string) *ListErRouteEntriesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListErRouteEntriesRequest) SetRouteType(v string) *ListErRouteEntriesRequest {
	s.RouteType = &v
	return s
}

func (s *ListErRouteEntriesRequest) SetStatus(v string) *ListErRouteEntriesRequest {
	s.Status = &v
	return s
}

func (s *ListErRouteEntriesRequest) Validate() error {
	return dara.Validate(s)
}
