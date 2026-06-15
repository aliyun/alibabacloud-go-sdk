// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateRouteEntryRequest
	GetClientToken() *string
	SetDestinationCidrBlock(v string) *CreateRouteEntryRequest
	GetDestinationCidrBlock() *string
	SetNextHopId(v string) *CreateRouteEntryRequest
	GetNextHopId() *string
	SetNextHopList(v []*CreateRouteEntryRequestNextHopList) *CreateRouteEntryRequest
	GetNextHopList() []*CreateRouteEntryRequestNextHopList
	SetNextHopType(v string) *CreateRouteEntryRequest
	GetNextHopType() *string
	SetOwnerAccount(v string) *CreateRouteEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateRouteEntryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateRouteEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateRouteEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateRouteEntryRequest
	GetResourceOwnerId() *int64
	SetRouteTableId(v string) *CreateRouteEntryRequest
	GetRouteTableId() *string
}

type CreateRouteEntryRequest struct {
	// A unique, case-sensitive identifier to ensure request idempotency. The token must be 1 to 64 ASCII characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The destination CIDR block. The value must be a valid CIDR block. Set the value to `0.0.0.0/0` to create a default route.
	//
	// This parameter is required.
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The ID of the next hop.
	//
	// **Note*	- Specify either `NextHopId` or `NextHopList`, but not both.
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// The list of next hops for an ECMP route.
	//
	// **Note*	- Specify either `NextHopId` or `NextHopList`, but not both.
	NextHopList []*CreateRouteEntryRequestNextHopList `json:"NextHopList,omitempty" xml:"NextHopList,omitempty" type:"Repeated"`
	// The type of the next hop. Valid values:
	//
	// - **Instance**: an ECS instance.
	//
	// - **HaVip**: a high-availability virtual IP address.
	//
	// - **RouterInterface**: a router interface.
	//
	// - **NetworkInterface**: an elastic network interface.
	//
	// - **VpnGateway**: a VPN gateway.
	//
	// - **Ipv6Gateway**: an IPv6 gateway.
	//
	// - **NatGateway**: a NAT gateway.
	//
	// **Note*	- This parameter is required if you specify `NextHopId`.
	NextHopType  *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](~~docid:36063~~) operation to get the latest list of regions.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the route table.
	//
	// This parameter is required.
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s CreateRouteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *CreateRouteEntryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateRouteEntryRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *CreateRouteEntryRequest) GetNextHopId() *string {
	return s.NextHopId
}

func (s *CreateRouteEntryRequest) GetNextHopList() []*CreateRouteEntryRequestNextHopList {
	return s.NextHopList
}

func (s *CreateRouteEntryRequest) GetNextHopType() *string {
	return s.NextHopType
}

func (s *CreateRouteEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateRouteEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateRouteEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRouteEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateRouteEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateRouteEntryRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *CreateRouteEntryRequest) SetClientToken(v string) *CreateRouteEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRouteEntryRequest) SetDestinationCidrBlock(v string) *CreateRouteEntryRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *CreateRouteEntryRequest) SetNextHopId(v string) *CreateRouteEntryRequest {
	s.NextHopId = &v
	return s
}

func (s *CreateRouteEntryRequest) SetNextHopList(v []*CreateRouteEntryRequestNextHopList) *CreateRouteEntryRequest {
	s.NextHopList = v
	return s
}

func (s *CreateRouteEntryRequest) SetNextHopType(v string) *CreateRouteEntryRequest {
	s.NextHopType = &v
	return s
}

func (s *CreateRouteEntryRequest) SetOwnerAccount(v string) *CreateRouteEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateRouteEntryRequest) SetOwnerId(v int64) *CreateRouteEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRouteEntryRequest) SetRegionId(v string) *CreateRouteEntryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRouteEntryRequest) SetResourceOwnerAccount(v string) *CreateRouteEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateRouteEntryRequest) SetResourceOwnerId(v int64) *CreateRouteEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateRouteEntryRequest) SetRouteTableId(v string) *CreateRouteEntryRequest {
	s.RouteTableId = &v
	return s
}

func (s *CreateRouteEntryRequest) Validate() error {
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

type CreateRouteEntryRequestNextHopList struct {
	// The ID of the next hop in the ECMP route.
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// The type of the next hop in the ECMP route. Valid values:
	//
	// - **Instance**: an ECS instance.
	//
	// - **HaVip**: a high-availability virtual IP address.
	//
	// - **RouterInterface**: a router interface.
	//
	// - **NetworkInterface**: an elastic network interface.
	//
	// - **VpnGateway**: a VPN gateway.
	//
	// - **Ipv6Gateway**: an IPv6 gateway.
	//
	// - **NatGateway**: a NAT gateway.
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
}

func (s CreateRouteEntryRequestNextHopList) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteEntryRequestNextHopList) GoString() string {
	return s.String()
}

func (s *CreateRouteEntryRequestNextHopList) GetNextHopId() *string {
	return s.NextHopId
}

func (s *CreateRouteEntryRequestNextHopList) GetNextHopType() *string {
	return s.NextHopType
}

func (s *CreateRouteEntryRequestNextHopList) SetNextHopId(v string) *CreateRouteEntryRequestNextHopList {
	s.NextHopId = &v
	return s
}

func (s *CreateRouteEntryRequestNextHopList) SetNextHopType(v string) *CreateRouteEntryRequestNextHopList {
	s.NextHopType = &v
	return s
}

func (s *CreateRouteEntryRequestNextHopList) Validate() error {
	return dara.Validate(s)
}
