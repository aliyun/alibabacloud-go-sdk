// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *DeleteRouteEntriesRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DeleteRouteEntriesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteRouteEntriesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteRouteEntriesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteRouteEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteRouteEntriesRequest
	GetResourceOwnerId() *int64
	SetRouteEntries(v []*DeleteRouteEntriesRequestRouteEntries) *DeleteRouteEntriesRequest
	GetRouteEntries() []*DeleteRouteEntriesRequestRouteEntries
}

type DeleteRouteEntriesRequest struct {
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the route table.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The information about the routes that you want to delete.
	RouteEntries []*DeleteRouteEntriesRequestRouteEntries `json:"RouteEntries,omitempty" xml:"RouteEntries,omitempty" type:"Repeated"`
}

func (s DeleteRouteEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *DeleteRouteEntriesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteRouteEntriesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteRouteEntriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteRouteEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRouteEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteRouteEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteRouteEntriesRequest) GetRouteEntries() []*DeleteRouteEntriesRequestRouteEntries {
	return s.RouteEntries
}

func (s *DeleteRouteEntriesRequest) SetDryRun(v bool) *DeleteRouteEntriesRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteRouteEntriesRequest) SetOwnerAccount(v string) *DeleteRouteEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteRouteEntriesRequest) SetOwnerId(v int64) *DeleteRouteEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteRouteEntriesRequest) SetRegionId(v string) *DeleteRouteEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRouteEntriesRequest) SetResourceOwnerAccount(v string) *DeleteRouteEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteRouteEntriesRequest) SetResourceOwnerId(v int64) *DeleteRouteEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteRouteEntriesRequest) SetRouteEntries(v []*DeleteRouteEntriesRequestRouteEntries) *DeleteRouteEntriesRequest {
	s.RouteEntries = v
	return s
}

func (s *DeleteRouteEntriesRequest) Validate() error {
	return dara.Validate(s)
}

type DeleteRouteEntriesRequestRouteEntries struct {
	// The destination CIDR block of the route that you want to delete. IPv4 and IPv6 CIDR blocks are supported. You can specify up to 50 destination CIDR blocks.
	//
	// >  If **RouteEntryId*	- is not specified, **DstCidrBlock*	- and **NextHop*	- are required.
	//
	// example:
	//
	// 47.100.XX.XX/24
	DstCidrBlock *string `json:"DstCidrBlock,omitempty" xml:"DstCidrBlock,omitempty"`
	// The ID of the next hop that you want to delete. You can specify up to 50 next hop IDs.
	//
	// >  If **RouteEntryId*	- is not specified, **DstCidrBlock*	- and **NextHop*	- are required.
	//
	// example:
	//
	// i-j6c2fp57q8rr4jlu****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// The ID of the route that you want to delete. You can specify up to 50 route IDs.
	//
	// >  If **RouteEntryId*	- is not specified, **DstCidrBlock*	- and **NextHop*	- are required.
	//
	// example:
	//
	// rte-bp1mnnr2al0naomnpv****
	RouteEntryId *string `json:"RouteEntryId,omitempty" xml:"RouteEntryId,omitempty"`
	// The ID of the route table to which the routes to be deleted belongs. You can specify up to 50 route table IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-2ze3jgygk9bmsj23s****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DeleteRouteEntriesRequestRouteEntries) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteEntriesRequestRouteEntries) GoString() string {
	return s.String()
}

func (s *DeleteRouteEntriesRequestRouteEntries) GetDstCidrBlock() *string {
	return s.DstCidrBlock
}

func (s *DeleteRouteEntriesRequestRouteEntries) GetNextHop() *string {
	return s.NextHop
}

func (s *DeleteRouteEntriesRequestRouteEntries) GetRouteEntryId() *string {
	return s.RouteEntryId
}

func (s *DeleteRouteEntriesRequestRouteEntries) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DeleteRouteEntriesRequestRouteEntries) SetDstCidrBlock(v string) *DeleteRouteEntriesRequestRouteEntries {
	s.DstCidrBlock = &v
	return s
}

func (s *DeleteRouteEntriesRequestRouteEntries) SetNextHop(v string) *DeleteRouteEntriesRequestRouteEntries {
	s.NextHop = &v
	return s
}

func (s *DeleteRouteEntriesRequestRouteEntries) SetRouteEntryId(v string) *DeleteRouteEntriesRequestRouteEntries {
	s.RouteEntryId = &v
	return s
}

func (s *DeleteRouteEntriesRequestRouteEntries) SetRouteTableId(v string) *DeleteRouteEntriesRequestRouteEntries {
	s.RouteTableId = &v
	return s
}

func (s *DeleteRouteEntriesRequestRouteEntries) Validate() error {
	return dara.Validate(s)
}
