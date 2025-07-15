// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationCidrBlock(v string) *DeleteRouteEntryRequest
	GetDestinationCidrBlock() *string
	SetDryRun(v bool) *DeleteRouteEntryRequest
	GetDryRun() *bool
	SetNextHopId(v string) *DeleteRouteEntryRequest
	GetNextHopId() *string
	SetNextHopList(v []*DeleteRouteEntryRequestNextHopList) *DeleteRouteEntryRequest
	GetNextHopList() []*DeleteRouteEntryRequestNextHopList
	SetOwnerAccount(v string) *DeleteRouteEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteRouteEntryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteRouteEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteRouteEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteRouteEntryRequest
	GetResourceOwnerId() *int64
	SetRouteEntryId(v string) *DeleteRouteEntryRequest
	GetRouteEntryId() *string
	SetRouteTableId(v string) *DeleteRouteEntryRequest
	GetRouteTableId() *string
}

type DeleteRouteEntryRequest struct {
	// The destination CIDR block of the route. Only IPv4 CIDR blocks, IPv6 CIDR blocks, and prefix lists are supported.
	//
	// example:
	//
	// 47.100.XX.XX/16
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the next hop.
	//
	// 	- To delete a route other than an equal-cost multi-path (ECMP) route, set the **NextHopId*	- parameter and ignore the **NextHopList*	- parameter.
	//
	// 	- To delete an ECMP route, set the **NextHopList*	- parameter and ignore the **NextHopId*	- parameter.
	//
	// example:
	//
	// ri-2zeo3xzyf38r4urzd****
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// The list of the next hop of the ECMP route.
	NextHopList  []*DeleteRouteEntryRequestNextHopList `json:"NextHopList,omitempty" xml:"NextHopList,omitempty" type:"Repeated"`
	OwnerAccount *string                               `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the route table.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the route that you want to delete.
	//
	// example:
	//
	// rte-bp1mnnr2al0naomnpv****
	RouteEntryId *string `json:"RouteEntryId,omitempty" xml:"RouteEntryId,omitempty"`
	// The ID of the route table to which the route belongs.
	//
	// example:
	//
	// vtb-2ze3jgygk9bmsj23s****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DeleteRouteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteRouteEntryRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DeleteRouteEntryRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteRouteEntryRequest) GetNextHopId() *string {
	return s.NextHopId
}

func (s *DeleteRouteEntryRequest) GetNextHopList() []*DeleteRouteEntryRequestNextHopList {
	return s.NextHopList
}

func (s *DeleteRouteEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteRouteEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteRouteEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRouteEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteRouteEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteRouteEntryRequest) GetRouteEntryId() *string {
	return s.RouteEntryId
}

func (s *DeleteRouteEntryRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DeleteRouteEntryRequest) SetDestinationCidrBlock(v string) *DeleteRouteEntryRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DeleteRouteEntryRequest) SetDryRun(v bool) *DeleteRouteEntryRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteRouteEntryRequest) SetNextHopId(v string) *DeleteRouteEntryRequest {
	s.NextHopId = &v
	return s
}

func (s *DeleteRouteEntryRequest) SetNextHopList(v []*DeleteRouteEntryRequestNextHopList) *DeleteRouteEntryRequest {
	s.NextHopList = v
	return s
}

func (s *DeleteRouteEntryRequest) SetOwnerAccount(v string) *DeleteRouteEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteRouteEntryRequest) SetOwnerId(v int64) *DeleteRouteEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteRouteEntryRequest) SetRegionId(v string) *DeleteRouteEntryRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRouteEntryRequest) SetResourceOwnerAccount(v string) *DeleteRouteEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteRouteEntryRequest) SetResourceOwnerId(v int64) *DeleteRouteEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteRouteEntryRequest) SetRouteEntryId(v string) *DeleteRouteEntryRequest {
	s.RouteEntryId = &v
	return s
}

func (s *DeleteRouteEntryRequest) SetRouteTableId(v string) *DeleteRouteEntryRequest {
	s.RouteTableId = &v
	return s
}

func (s *DeleteRouteEntryRequest) Validate() error {
	return dara.Validate(s)
}

type DeleteRouteEntryRequestNextHopList struct {
	// The ID of the next hop that is configured for ECMP routing. You can specify information about at most 16 next hops.
	//
	// example:
	//
	// ri-2zeo3xzyf38r43cd****
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// The type of the next hop that is configured for ECMP routing. Set the value to **RouterInterface**. You can specify information about at most 16 next hops.
	//
	// example:
	//
	// RouterInterface
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
}

func (s DeleteRouteEntryRequestNextHopList) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteEntryRequestNextHopList) GoString() string {
	return s.String()
}

func (s *DeleteRouteEntryRequestNextHopList) GetNextHopId() *string {
	return s.NextHopId
}

func (s *DeleteRouteEntryRequestNextHopList) GetNextHopType() *string {
	return s.NextHopType
}

func (s *DeleteRouteEntryRequestNextHopList) SetNextHopId(v string) *DeleteRouteEntryRequestNextHopList {
	s.NextHopId = &v
	return s
}

func (s *DeleteRouteEntryRequestNextHopList) SetNextHopType(v string) *DeleteRouteEntryRequestNextHopList {
	s.NextHopType = &v
	return s
}

func (s *DeleteRouteEntryRequestNextHopList) Validate() error {
	return dara.Validate(s)
}
