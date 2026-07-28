// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRouteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyRouteEntryRequest
	GetDescription() *string
	SetDestinationCidrBlock(v string) *ModifyRouteEntryRequest
	GetDestinationCidrBlock() *string
	SetDryRun(v bool) *ModifyRouteEntryRequest
	GetDryRun() *bool
	SetNewNextHopId(v string) *ModifyRouteEntryRequest
	GetNewNextHopId() *string
	SetNewNextHopType(v string) *ModifyRouteEntryRequest
	GetNewNextHopType() *string
	SetOwnerAccount(v string) *ModifyRouteEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyRouteEntryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyRouteEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyRouteEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyRouteEntryRequest
	GetResourceOwnerId() *int64
	SetRouteEntryId(v string) *ModifyRouteEntryRequest
	GetRouteEntryId() *string
	SetRouteEntryName(v string) *ModifyRouteEntryRequest
	GetRouteEntryName() *string
	SetRouteTableId(v string) *ModifyRouteEntryRequest
	GetRouteTableId() *string
}

type ModifyRouteEntryRequest struct {
	// The description of the route entry.
	//
	// The description must be 1 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// EntryDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IPv4 CIDR block of the route entry. IPv4 and IPv6 CIDR blocks are supported.
	//
	// > If the **RouteEntryId*	- parameter is not specified, the **DestinationCidrBlock*	- and **RouteTableId*	- parameters are required.
	//
	// > To change the IPv4 CIDR block of a route to a **prefix list**, you must specify the **RouteEntryId*	- parameter. The **DestinationCidrBlock*	- parameter does not support prefix list CIDR blocks or prefix list instance IDs.
	//
	// example:
	//
	// 192.168.0.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// **true**: performs a dry run without modifying the route. The system checks the AccessKey pair, the authorization of the Resource Access Management (RAM) user, and the required parameters. If the check fails, the corresponding error is returned. If the check succeeds, the `DryRunOperation` error code is returned.
	//
	// **false*	- (default): sends a Normal request. If the check succeeds, a 2xx HTTP status code is returned and the route is modified.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The new next hop instance ID of the route entry.
	//
	// example:
	//
	// eni-bp17y37ytsenqyim****
	NewNextHopId *string `json:"NewNextHopId,omitempty" xml:"NewNextHopId,omitempty"`
	// The new next hop type of the route entry. Valid values:
	//
	// - **Instance**: ECS instance.
	//
	// - **HaVip**: high-availability virtual IP address.
	//
	// - **RouterInterface**: vRouter interface.
	//
	// - **NetworkInterface**: elastic network interfaces (ENIs).
	//
	// - **VpnGateway**: VPN gateway.
	//
	// - **IPv6Gateway**: IPv6 gateway.
	//
	// - **NatGateway**: NAT gateway.
	//
	// - **Attachment**: forward router.
	//
	// - **VpcPeer**: VPC peering connection.
	//
	// - **Ipv4Gateway**: IPv4 gateway.
	//
	// - **GatewayEndpoint**: gateway endpoint.
	//
	// - **Ecr**: Express Connect Router (ECR).
	//
	// - **GatewayLoadBalancerEndpoint**: Gateway Load Balancer endpoint (GWLBe).
	//
	// - **RouteTargetGroup**: route target group.
	//
	// example:
	//
	// NetworkInterface
	NewNextHopType *string `json:"NewNextHopType,omitempty" xml:"NewNextHopType,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the route entry.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the custom route entry.
	//
	// example:
	//
	// rte-acfvgfsghfdd****
	RouteEntryId *string `json:"RouteEntryId,omitempty" xml:"RouteEntryId,omitempty"`
	// The name of the route entry.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// EntryName
	RouteEntryName *string `json:"RouteEntryName,omitempty" xml:"RouteEntryName,omitempty"`
	// The route table ID.
	//
	// example:
	//
	// vtb-bp1nk7zk65du3pni8z9td
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s ModifyRouteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *ModifyRouteEntryRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyRouteEntryRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *ModifyRouteEntryRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyRouteEntryRequest) GetNewNextHopId() *string {
	return s.NewNextHopId
}

func (s *ModifyRouteEntryRequest) GetNewNextHopType() *string {
	return s.NewNextHopType
}

func (s *ModifyRouteEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyRouteEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyRouteEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRouteEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyRouteEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyRouteEntryRequest) GetRouteEntryId() *string {
	return s.RouteEntryId
}

func (s *ModifyRouteEntryRequest) GetRouteEntryName() *string {
	return s.RouteEntryName
}

func (s *ModifyRouteEntryRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *ModifyRouteEntryRequest) SetDescription(v string) *ModifyRouteEntryRequest {
	s.Description = &v
	return s
}

func (s *ModifyRouteEntryRequest) SetDestinationCidrBlock(v string) *ModifyRouteEntryRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ModifyRouteEntryRequest) SetDryRun(v bool) *ModifyRouteEntryRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyRouteEntryRequest) SetNewNextHopId(v string) *ModifyRouteEntryRequest {
	s.NewNextHopId = &v
	return s
}

func (s *ModifyRouteEntryRequest) SetNewNextHopType(v string) *ModifyRouteEntryRequest {
	s.NewNextHopType = &v
	return s
}

func (s *ModifyRouteEntryRequest) SetOwnerAccount(v string) *ModifyRouteEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyRouteEntryRequest) SetOwnerId(v int64) *ModifyRouteEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyRouteEntryRequest) SetRegionId(v string) *ModifyRouteEntryRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRouteEntryRequest) SetResourceOwnerAccount(v string) *ModifyRouteEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyRouteEntryRequest) SetResourceOwnerId(v int64) *ModifyRouteEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyRouteEntryRequest) SetRouteEntryId(v string) *ModifyRouteEntryRequest {
	s.RouteEntryId = &v
	return s
}

func (s *ModifyRouteEntryRequest) SetRouteEntryName(v string) *ModifyRouteEntryRequest {
	s.RouteEntryName = &v
	return s
}

func (s *ModifyRouteEntryRequest) SetRouteTableId(v string) *ModifyRouteEntryRequest {
	s.RouteTableId = &v
	return s
}

func (s *ModifyRouteEntryRequest) Validate() error {
	return dara.Validate(s)
}
