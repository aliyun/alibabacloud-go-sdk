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
	// The destination CIDR block of the route. IPv4 CIDR blocks, IPv6 CIDR blocks, prefix list CIDR blocks, and prefix list instance IDs are supported. This parameter is mutually exclusive with the RouteEntryId parameter.
	//
	// > If the **RouteEntryId*	- parameter is not specified, the **DestinationCidrBlock*	- and **RouteTableId*	- parameters are required. Configure the **NextHopId*	- or **NextHopList*	- parameter as needed.
	//
	// example:
	//
	// 47.100.XX.XX/16
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// **true**: performs a dry run without deleting the route. The system checks the AccessKey pair, the authorization of the Resource Access Management (RAM) user, and the required parameters. If the check fails, the corresponding error is returned. If the check succeeds, the error code `DryRunOperation` is returned.
	//
	// **false*	- (default): sends a normal request. After the check succeeds, a 2xx HTTP status code is returned and the route is deleted.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the next hop instance.
	//
	// - To delete a non-ECMP route, specify **NextHopId**. Do not specify **NextHopList**.
	//
	// - To delete an ECMP route, specify **NextHopList**. Do not specify **NextHopId**.
	//
	// example:
	//
	// ri-2zeo3xzyf38r4urzd****
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// The information about the next hop instances of the ECMP route. A maximum of 16 next hop instances are supported.
	NextHopList  []*DeleteRouteEntryRequestNextHopList `json:"NextHopList,omitempty" xml:"NextHopList,omitempty" type:"Repeated"`
	OwnerAccount *string                               `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the route table resides.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the route that you want to delete. This parameter is mutually exclusive with the DestinationCidrBlock parameter.
	//
	// > If the **DestinationCidrBlock*	- parameter is not specified, the **RouteEntryId*	- parameter is required.
	//
	// example:
	//
	// rte-bp1mnnr2al0naomnpv****
	RouteEntryId *string `json:"RouteEntryId,omitempty" xml:"RouteEntryId,omitempty"`
	// The ID of the route table that contains the route.
	//
	// > If the **RouteEntryId*	- parameter is not specified, the **DestinationCidrBlock*	- and **RouteTableId*	- parameters are required. Configure the **NextHopId*	- or **NextHopList*	- parameter as needed.
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
	if s.NextHopList != nil {
		for _, item := range s.NextHopList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteRouteEntryRequestNextHopList struct {
	// The ID of the next hop instance of the ECMP route. A maximum of 16 next hop instances are supported.
	//
	// example:
	//
	// ri-2zeo3xzyf38r43cd****
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// The type of the next hop of the ECMP route. Set the value to **RouterInterface*	- (router interface). A maximum of 16 next hop instances are supported.
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
