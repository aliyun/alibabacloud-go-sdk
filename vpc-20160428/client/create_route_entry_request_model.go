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
	SetDescription(v string) *CreateRouteEntryRequest
	GetDescription() *string
	SetDestinationCidrBlock(v string) *CreateRouteEntryRequest
	GetDestinationCidrBlock() *string
	SetDryRun(v bool) *CreateRouteEntryRequest
	GetDryRun() *bool
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
	SetRouteEntryName(v string) *CreateRouteEntryRequest
	GetRouteEntryName() *string
	SetRouteTableId(v string) *CreateRouteEntryRequest
	GetRouteTableId() *string
}

type CreateRouteEntryRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. The ClientToken value can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- for each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the custom route entry.
	//
	// The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination CIDR block of the custom route entry. Both IPv4 and IPv6 CIDR blocks are supported. Make sure that the destination CIDR block meets the following requirements:
	//
	// 	- The destination CIDR block is not 100.64.0.0/10 or a subset of 100.64.0.0/10.
	//
	// 	- The destination CIDR block of the custom route entry is different from the destination CIDR blocks of other route entries in the same route table.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the required parameters, request syntax, and limits. If the request fails, an error message is returned. If the request passes the validation, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): sends the request. If the request passes the check, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the next hop for the custom route.
	//
	// >  [](#-nexthoptype--ecr-describeexpressconnectrouterassociation--associationid--id)If you set the NextHopType parameter to ECR, call the [DescribeExpressConnectRouterAssociation](https://help.aliyun.com/document_detail/2712069.html) operation to access the AssociationId and use it as the next hop ID.
	//
	// example:
	//
	// i-j6c2fp57q8rr4jlu****
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// The next hop list.
	NextHopList []*CreateRouteEntryRequestNextHopList `json:"NextHopList,omitempty" xml:"NextHopList,omitempty" type:"Repeated"`
	// The type of next hop of the custom route entry. Valid values:
	//
	// 	- **Instance**: an Elastic Compute Service (ECS) instance. This is the default value.
	//
	// 	- **HaVip**: a high-availability virtual IP address (HaVip).
	//
	// 	- **RouterInterface**: a router interface.
	//
	// 	- **NetworkInterface**: an elastic network interface (ENI).
	//
	// 	- **VpnGateway**: a VPN gateway.
	//
	// 	- **IPv6Gateway**: an IPv6 gateway.
	//
	// 	- **NatGateway**: a NAT gateway.
	//
	// 	- **Attachment**: a transit router.
	//
	// 	- **VpcPeer**: a VPC peering connection.
	//
	// 	- **Ipv4Gateway**: an IPv4 gateway.
	//
	// 	- **GatewayEndpoint**: a gateway endpoint.
	//
	// 	- **Ecr**: an Express Connect Router (ECR).
	//
	// example:
	//
	// RouterInterface
	NextHopType  *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// The name of the custom route entry that you want to add.
	//
	// The name must be 1 to 128 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	RouteEntryName *string `json:"RouteEntryName,omitempty" xml:"RouteEntryName,omitempty"`
	// The ID of the route table to which you want to add a custom route entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-bp145q7glnuzd****
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

func (s *CreateRouteEntryRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRouteEntryRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *CreateRouteEntryRequest) GetDryRun() *bool {
	return s.DryRun
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

func (s *CreateRouteEntryRequest) GetRouteEntryName() *string {
	return s.RouteEntryName
}

func (s *CreateRouteEntryRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *CreateRouteEntryRequest) SetClientToken(v string) *CreateRouteEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRouteEntryRequest) SetDescription(v string) *CreateRouteEntryRequest {
	s.Description = &v
	return s
}

func (s *CreateRouteEntryRequest) SetDestinationCidrBlock(v string) *CreateRouteEntryRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *CreateRouteEntryRequest) SetDryRun(v bool) *CreateRouteEntryRequest {
	s.DryRun = &v
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

func (s *CreateRouteEntryRequest) SetRouteEntryName(v string) *CreateRouteEntryRequest {
	s.RouteEntryName = &v
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
	// The ID of the next hop of the ECMP route.
	//
	// example:
	//
	// ri-2zeo3xzyf3cd8r4****
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// The type of next hop of the ECMP route entry. Set the value to **RouterInterface**.
	//
	// example:
	//
	// RouterInterface
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The weight of the next hop of the ECMP route entry.
	//
	// example:
	//
	// 10
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
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

func (s *CreateRouteEntryRequestNextHopList) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateRouteEntryRequestNextHopList) SetNextHopId(v string) *CreateRouteEntryRequestNextHopList {
	s.NextHopId = &v
	return s
}

func (s *CreateRouteEntryRequestNextHopList) SetNextHopType(v string) *CreateRouteEntryRequestNextHopList {
	s.NextHopType = &v
	return s
}

func (s *CreateRouteEntryRequestNextHopList) SetWeight(v int32) *CreateRouteEntryRequestNextHopList {
	s.Weight = &v
	return s
}

func (s *CreateRouteEntryRequestNextHopList) Validate() error {
	return dara.Validate(s)
}
