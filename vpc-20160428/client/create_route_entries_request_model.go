// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *CreateRouteEntriesRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *CreateRouteEntriesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateRouteEntriesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateRouteEntriesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateRouteEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateRouteEntriesRequest
	GetResourceOwnerId() *int64
	SetRouteEntries(v []*CreateRouteEntriesRequestRouteEntries) *CreateRouteEntriesRequest
	GetRouteEntries() []*CreateRouteEntriesRequestRouteEntries
}

type CreateRouteEntriesRequest struct {
	// Specifies whether to perform a dry run. Valid values:
	//
	// **true**: performs a dry run without creating routes. The system checks the AccessKey pair, the authorization of the Resource Access Management (RAM) user, and the required parameters. If the check fails, the corresponding error is returned. If the check passes, the `DryRunOperation` error code is returned.
	//
	// **false*	- (default): sends a Normal request. If the check passes, a 2xx HTTP status code is returned and the routes are created.
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the route table.
	//
	// You can call [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of route entry information.
	//
	// This parameter is required.
	RouteEntries []*CreateRouteEntriesRequestRouteEntries `json:"RouteEntries,omitempty" xml:"RouteEntries,omitempty" type:"Repeated"`
}

func (s CreateRouteEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *CreateRouteEntriesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateRouteEntriesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateRouteEntriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateRouteEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRouteEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateRouteEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateRouteEntriesRequest) GetRouteEntries() []*CreateRouteEntriesRequestRouteEntries {
	return s.RouteEntries
}

func (s *CreateRouteEntriesRequest) SetDryRun(v bool) *CreateRouteEntriesRequest {
	s.DryRun = &v
	return s
}

func (s *CreateRouteEntriesRequest) SetOwnerAccount(v string) *CreateRouteEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateRouteEntriesRequest) SetOwnerId(v int64) *CreateRouteEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRouteEntriesRequest) SetRegionId(v string) *CreateRouteEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRouteEntriesRequest) SetResourceOwnerAccount(v string) *CreateRouteEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateRouteEntriesRequest) SetResourceOwnerId(v int64) *CreateRouteEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateRouteEntriesRequest) SetRouteEntries(v []*CreateRouteEntriesRequestRouteEntries) *CreateRouteEntriesRequest {
	s.RouteEntries = v
	return s
}

func (s *CreateRouteEntriesRequest) Validate() error {
	if s.RouteEntries != nil {
		for _, item := range s.RouteEntries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateRouteEntriesRequestRouteEntries struct {
	// The description of the custom route entry. You can specify a maximum of 50 descriptions.
	//
	// The description must be 1 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination CIDR block of the custom route entry. Both IPv4 and IPv6 destination CIDR blocks are supported. You can specify a maximum of 50 destination CIDR blocks. The following requirements must be met:
	//
	//
	//
	// - The destination CIDR block cannot point to or be contained by 100.64.0.0/10.
	//
	//
	//
	// - The destination CIDR blocks of different route entries in the same route table must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.0/24
	DstCidrBlock *string `json:"DstCidrBlock,omitempty" xml:"DstCidrBlock,omitempty"`
	// The version of the IP protocol. You can specify a maximum of 50 IP protocol versions. Valid values:
	//
	// - **4**: IPv4.
	//
	// - **6**: IPv6.
	//
	// example:
	//
	// 4
	IpVersion *int32 `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The name of the custom route entry to add. You can specify a maximum of 50 names.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the next hop instance for the custom route entry. You can specify a maximum of 50 instance IDs.
	//
	// > If NextHopType is set to ECR, you can call [DescribeExpressConnectRouterAssociation](https://help.aliyun.com/document_detail/2712069.html) to obtain the AssociationId as the next hop ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-j6c2fp57q8rr4jlu****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// The type of the next hop for the custom route entry. You can specify a maximum of 50 next hop types. Valid values:
	//
	// - **Instance*	- (default): ECS instance. Forwards traffic to an ECS instance.
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
	// - **Attachment**: transit router. Forwards traffic to a transit router.
	//
	// - **VpcPeer**: VPC peering connection.
	//
	// - **Ipv4Gateway**: IPv4 gateway.
	//
	// - **GatewayEndpoint**: gateway endpoint.
	//
	// - **CenBasic**: CEN that does not support transit routers.
	//
	// - **Ecr**: Express Connect Router (ECR).
	//
	// - **GatewayLoadBalancerEndpoint**: Gateway Load Balancer endpoint (GWLBe).
	//
	// - **RouteTargetGroup**: routing target group.
	//
	// This parameter is required.
	//
	// example:
	//
	// RouterInterface
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The ID of the route table to which you want to add the custom route entry. You can specify a maximum of 50 route table IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-bp145q7glnuzd****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s CreateRouteEntriesRequestRouteEntries) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteEntriesRequestRouteEntries) GoString() string {
	return s.String()
}

func (s *CreateRouteEntriesRequestRouteEntries) GetDescription() *string {
	return s.Description
}

func (s *CreateRouteEntriesRequestRouteEntries) GetDstCidrBlock() *string {
	return s.DstCidrBlock
}

func (s *CreateRouteEntriesRequestRouteEntries) GetIpVersion() *int32 {
	return s.IpVersion
}

func (s *CreateRouteEntriesRequestRouteEntries) GetName() *string {
	return s.Name
}

func (s *CreateRouteEntriesRequestRouteEntries) GetNextHop() *string {
	return s.NextHop
}

func (s *CreateRouteEntriesRequestRouteEntries) GetNextHopType() *string {
	return s.NextHopType
}

func (s *CreateRouteEntriesRequestRouteEntries) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *CreateRouteEntriesRequestRouteEntries) SetDescription(v string) *CreateRouteEntriesRequestRouteEntries {
	s.Description = &v
	return s
}

func (s *CreateRouteEntriesRequestRouteEntries) SetDstCidrBlock(v string) *CreateRouteEntriesRequestRouteEntries {
	s.DstCidrBlock = &v
	return s
}

func (s *CreateRouteEntriesRequestRouteEntries) SetIpVersion(v int32) *CreateRouteEntriesRequestRouteEntries {
	s.IpVersion = &v
	return s
}

func (s *CreateRouteEntriesRequestRouteEntries) SetName(v string) *CreateRouteEntriesRequestRouteEntries {
	s.Name = &v
	return s
}

func (s *CreateRouteEntriesRequestRouteEntries) SetNextHop(v string) *CreateRouteEntriesRequestRouteEntries {
	s.NextHop = &v
	return s
}

func (s *CreateRouteEntriesRequestRouteEntries) SetNextHopType(v string) *CreateRouteEntriesRequestRouteEntries {
	s.NextHopType = &v
	return s
}

func (s *CreateRouteEntriesRequestRouteEntries) SetRouteTableId(v string) *CreateRouteEntriesRequestRouteEntries {
	s.RouteTableId = &v
	return s
}

func (s *CreateRouteEntriesRequestRouteEntries) Validate() error {
	return dara.Validate(s)
}
