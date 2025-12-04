// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpdRouteEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationCidrBlock(v string) *ListVpdRouteEntriesRequest
	GetDestinationCidrBlock() *string
	SetEnablePage(v bool) *ListVpdRouteEntriesRequest
	GetEnablePage() *bool
	SetIgnoreDetailedRouteEntry(v bool) *ListVpdRouteEntriesRequest
	GetIgnoreDetailedRouteEntry() *bool
	SetNextHopId(v string) *ListVpdRouteEntriesRequest
	GetNextHopId() *string
	SetNextHopType(v string) *ListVpdRouteEntriesRequest
	GetNextHopType() *string
	SetPageNumber(v int32) *ListVpdRouteEntriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListVpdRouteEntriesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListVpdRouteEntriesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListVpdRouteEntriesRequest
	GetResourceGroupId() *string
	SetRouteType(v string) *ListVpdRouteEntriesRequest
	GetRouteType() *string
	SetStatus(v string) *ListVpdRouteEntriesRequest
	GetStatus() *string
	SetVpdId(v string) *ListVpdRouteEntriesRequest
	GetVpdId() *string
	SetVpdRouteEntryId(v string) *ListVpdRouteEntriesRequest
	GetVpdRouteEntryId() *string
}

type ListVpdRouteEntriesRequest struct {
	// Destination CIDR block
	//
	// example:
	//
	// 0.0.0.0/0
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Specifies whether to enable paged query. Optional values:
	//
	// 	- **true**: Enable pagination query
	//
	// 	- **false**: Pagination query is disabled
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
	// rg-acfm4mlwqjalz7a
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
	// Lingjun CIDR block instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-fuliephf
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// Lingjun CIDR block route entry instance ID
	//
	// example:
	//
	// vpd-rte-4r1zbhoh
	VpdRouteEntryId *string `json:"VpdRouteEntryId,omitempty" xml:"VpdRouteEntryId,omitempty"`
}

func (s ListVpdRouteEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVpdRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *ListVpdRouteEntriesRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *ListVpdRouteEntriesRequest) GetEnablePage() *bool {
	return s.EnablePage
}

func (s *ListVpdRouteEntriesRequest) GetIgnoreDetailedRouteEntry() *bool {
	return s.IgnoreDetailedRouteEntry
}

func (s *ListVpdRouteEntriesRequest) GetNextHopId() *string {
	return s.NextHopId
}

func (s *ListVpdRouteEntriesRequest) GetNextHopType() *string {
	return s.NextHopType
}

func (s *ListVpdRouteEntriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVpdRouteEntriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVpdRouteEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpdRouteEntriesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVpdRouteEntriesRequest) GetRouteType() *string {
	return s.RouteType
}

func (s *ListVpdRouteEntriesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListVpdRouteEntriesRequest) GetVpdId() *string {
	return s.VpdId
}

func (s *ListVpdRouteEntriesRequest) GetVpdRouteEntryId() *string {
	return s.VpdRouteEntryId
}

func (s *ListVpdRouteEntriesRequest) SetDestinationCidrBlock(v string) *ListVpdRouteEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ListVpdRouteEntriesRequest) SetEnablePage(v bool) *ListVpdRouteEntriesRequest {
	s.EnablePage = &v
	return s
}

func (s *ListVpdRouteEntriesRequest) SetIgnoreDetailedRouteEntry(v bool) *ListVpdRouteEntriesRequest {
	s.IgnoreDetailedRouteEntry = &v
	return s
}

func (s *ListVpdRouteEntriesRequest) SetNextHopId(v string) *ListVpdRouteEntriesRequest {
	s.NextHopId = &v
	return s
}

func (s *ListVpdRouteEntriesRequest) SetNextHopType(v string) *ListVpdRouteEntriesRequest {
	s.NextHopType = &v
	return s
}

func (s *ListVpdRouteEntriesRequest) SetPageNumber(v int32) *ListVpdRouteEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVpdRouteEntriesRequest) SetPageSize(v int32) *ListVpdRouteEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListVpdRouteEntriesRequest) SetRegionId(v string) *ListVpdRouteEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpdRouteEntriesRequest) SetResourceGroupId(v string) *ListVpdRouteEntriesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpdRouteEntriesRequest) SetRouteType(v string) *ListVpdRouteEntriesRequest {
	s.RouteType = &v
	return s
}

func (s *ListVpdRouteEntriesRequest) SetStatus(v string) *ListVpdRouteEntriesRequest {
	s.Status = &v
	return s
}

func (s *ListVpdRouteEntriesRequest) SetVpdId(v string) *ListVpdRouteEntriesRequest {
	s.VpdId = &v
	return s
}

func (s *ListVpdRouteEntriesRequest) SetVpdRouteEntryId(v string) *ListVpdRouteEntriesRequest {
	s.VpdRouteEntryId = &v
	return s
}

func (s *ListVpdRouteEntriesRequest) Validate() error {
	return dara.Validate(s)
}
