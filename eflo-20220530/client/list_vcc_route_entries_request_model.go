// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVccRouteEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationCidrBlock(v string) *ListVccRouteEntriesRequest
	GetDestinationCidrBlock() *string
	SetEnablePage(v bool) *ListVccRouteEntriesRequest
	GetEnablePage() *bool
	SetIgnoreDetailedRouteEntry(v bool) *ListVccRouteEntriesRequest
	GetIgnoreDetailedRouteEntry() *bool
	SetNextHopId(v string) *ListVccRouteEntriesRequest
	GetNextHopId() *string
	SetNextHopType(v string) *ListVccRouteEntriesRequest
	GetNextHopType() *string
	SetPageNumber(v int32) *ListVccRouteEntriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListVccRouteEntriesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListVccRouteEntriesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListVccRouteEntriesRequest
	GetResourceGroupId() *string
	SetRouteType(v string) *ListVccRouteEntriesRequest
	GetRouteType() *string
	SetStatus(v string) *ListVccRouteEntriesRequest
	GetStatus() *string
	SetVccId(v string) *ListVccRouteEntriesRequest
	GetVccId() *string
	SetVpdRouteEntryId(v string) *ListVccRouteEntriesRequest
	GetVpdRouteEntryId() *string
}

type ListVccRouteEntriesRequest struct {
	// Destination CIDR block
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
	// Filter 32 detailed CIDR blocks. Default value: true
	//
	// example:
	//
	// true
	IgnoreDetailedRouteEntry *bool `json:"IgnoreDetailedRouteEntry,omitempty" xml:"IgnoreDetailedRouteEntry,omitempty"`
	// Next Hop Instance
	//
	// example:
	//
	// vcc-cn-jaj34d75h01
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
	// The number of entries to return on each page. Default value: 20.
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
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Route type
	//
	// example:
	//
	// BGP
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	// The status of the enterprise-level snapshot policy.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the Lingjun connection instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
	// Lingjun CIDR block route entry instance ID
	//
	// example:
	//
	// vpd-rte-toekyqel
	VpdRouteEntryId *string `json:"VpdRouteEntryId,omitempty" xml:"VpdRouteEntryId,omitempty"`
}

func (s ListVccRouteEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVccRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *ListVccRouteEntriesRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *ListVccRouteEntriesRequest) GetEnablePage() *bool {
	return s.EnablePage
}

func (s *ListVccRouteEntriesRequest) GetIgnoreDetailedRouteEntry() *bool {
	return s.IgnoreDetailedRouteEntry
}

func (s *ListVccRouteEntriesRequest) GetNextHopId() *string {
	return s.NextHopId
}

func (s *ListVccRouteEntriesRequest) GetNextHopType() *string {
	return s.NextHopType
}

func (s *ListVccRouteEntriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVccRouteEntriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVccRouteEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVccRouteEntriesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVccRouteEntriesRequest) GetRouteType() *string {
	return s.RouteType
}

func (s *ListVccRouteEntriesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListVccRouteEntriesRequest) GetVccId() *string {
	return s.VccId
}

func (s *ListVccRouteEntriesRequest) GetVpdRouteEntryId() *string {
	return s.VpdRouteEntryId
}

func (s *ListVccRouteEntriesRequest) SetDestinationCidrBlock(v string) *ListVccRouteEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ListVccRouteEntriesRequest) SetEnablePage(v bool) *ListVccRouteEntriesRequest {
	s.EnablePage = &v
	return s
}

func (s *ListVccRouteEntriesRequest) SetIgnoreDetailedRouteEntry(v bool) *ListVccRouteEntriesRequest {
	s.IgnoreDetailedRouteEntry = &v
	return s
}

func (s *ListVccRouteEntriesRequest) SetNextHopId(v string) *ListVccRouteEntriesRequest {
	s.NextHopId = &v
	return s
}

func (s *ListVccRouteEntriesRequest) SetNextHopType(v string) *ListVccRouteEntriesRequest {
	s.NextHopType = &v
	return s
}

func (s *ListVccRouteEntriesRequest) SetPageNumber(v int32) *ListVccRouteEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVccRouteEntriesRequest) SetPageSize(v int32) *ListVccRouteEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListVccRouteEntriesRequest) SetRegionId(v string) *ListVccRouteEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *ListVccRouteEntriesRequest) SetResourceGroupId(v string) *ListVccRouteEntriesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVccRouteEntriesRequest) SetRouteType(v string) *ListVccRouteEntriesRequest {
	s.RouteType = &v
	return s
}

func (s *ListVccRouteEntriesRequest) SetStatus(v string) *ListVccRouteEntriesRequest {
	s.Status = &v
	return s
}

func (s *ListVccRouteEntriesRequest) SetVccId(v string) *ListVccRouteEntriesRequest {
	s.VccId = &v
	return s
}

func (s *ListVccRouteEntriesRequest) SetVpdRouteEntryId(v string) *ListVccRouteEntriesRequest {
	s.VpdRouteEntryId = &v
	return s
}

func (s *ListVccRouteEntriesRequest) Validate() error {
	return dara.Validate(s)
}
