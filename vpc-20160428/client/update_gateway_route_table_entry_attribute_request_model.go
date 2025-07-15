// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteTableEntryAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateGatewayRouteTableEntryAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateGatewayRouteTableEntryAttributeRequest
	GetDescription() *string
	SetDestinationCidrBlock(v string) *UpdateGatewayRouteTableEntryAttributeRequest
	GetDestinationCidrBlock() *string
	SetDryRun(v bool) *UpdateGatewayRouteTableEntryAttributeRequest
	GetDryRun() *bool
	SetGatewayRouteTableId(v string) *UpdateGatewayRouteTableEntryAttributeRequest
	GetGatewayRouteTableId() *string
	SetIPv4GatewayRouteTableId(v string) *UpdateGatewayRouteTableEntryAttributeRequest
	GetIPv4GatewayRouteTableId() *string
	SetName(v string) *UpdateGatewayRouteTableEntryAttributeRequest
	GetName() *string
	SetNextHopId(v string) *UpdateGatewayRouteTableEntryAttributeRequest
	GetNextHopId() *string
	SetNextHopType(v string) *UpdateGatewayRouteTableEntryAttributeRequest
	GetNextHopType() *string
	SetOwnerAccount(v string) *UpdateGatewayRouteTableEntryAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateGatewayRouteTableEntryAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateGatewayRouteTableEntryAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpdateGatewayRouteTableEntryAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateGatewayRouteTableEntryAttributeRequest
	GetResourceOwnerId() *int64
}

type UpdateGatewayRouteTableEntryAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the gateway route table.
	//
	// The description must be 2 to 256 characters in length. The description must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// new
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination CIDR block of the route entry in the gateway route table.
	//
	// This parameter is required.
	//
	// example:
	//
	// 47.100.XX.XX/16
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Specifies whether to precheck only this request. Valid values:
	//
	// 	- **true**: prechecks the request without modifying the gateway route table. The system checks the required parameters, request format, and service limits. If the request fails to pass the precheck, an error code is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: sends the request. This is the default value. If the request passes the precheck, a 2xx HTTP status code is returned and the gateway route table is modified.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the gateway route table that you want to modify.
	//
	// example:
	//
	// vtb-5ts0ohchwkp3dydt2****
	GatewayRouteTableId *string `json:"GatewayRouteTableId,omitempty" xml:"GatewayRouteTableId,omitempty"`
	// The ID of the gateway route table that you want to modify.
	//
	// example:
	//
	// vtb-5ts0ohchwkp3dydt2****
	IPv4GatewayRouteTableId *string `json:"IPv4GatewayRouteTableId,omitempty" xml:"IPv4GatewayRouteTableId,omitempty"`
	// The name of the gateway route table.
	//
	// The name must be 2 to 128 characters in length and can contain letter, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The new next hop ID of the route entry.
	//
	// 	- If you set **NextHopType*	- to **Instance**, specify an ECS instance ID for **NextHopId**.
	//
	// 	- If you set **NextHopType*	- to **NetworkInterface**, specify an ENI ID for **NextHopId**.
	//
	// 	- If you set **NextHopType*	- to **Local**, leave **NextHopId*	- empty. This indicates a local next hop.
	//
	// >  If the value of NextHopType is **Instance*	- or **NetworkInterface**, and you want to modify the next hop, you must set **NextHopType*	- to **Local*	- first. Then, set **NextHopType*	- to **Instance*	- or **NetworkInterface*	- and specify **NextHopId*	- based on your requirements. If the next hop type of a route entry is Instance or NetworkInterface, you cannot directly specify a different ENI ID or ECS instance ID for the NextHopId parameter.
	//
	// example:
	//
	// i-bp18xq9yguxoxe7m****
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// The new next hop type of the route. Valid values:
	//
	// 	- **Instance**: Elastic Compute Service (ECS) instance
	//
	// 	- **NetworkInterface**: elastic network interface (ENI)
	//
	// 	- **Local**: local next hop
	//
	// This parameter is required.
	//
	// example:
	//
	// EcsInstance
	NextHopType  *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the gateway route table that you want to modify belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-southeast-6
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateGatewayRouteTableEntryAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteTableEntryAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) GetGatewayRouteTableId() *string {
	return s.GatewayRouteTableId
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) GetIPv4GatewayRouteTableId() *string {
	return s.IPv4GatewayRouteTableId
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) GetName() *string {
	return s.Name
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) GetNextHopId() *string {
	return s.NextHopId
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) GetNextHopType() *string {
	return s.NextHopType
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) SetClientToken(v string) *UpdateGatewayRouteTableEntryAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) SetDescription(v string) *UpdateGatewayRouteTableEntryAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) SetDestinationCidrBlock(v string) *UpdateGatewayRouteTableEntryAttributeRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) SetDryRun(v bool) *UpdateGatewayRouteTableEntryAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) SetGatewayRouteTableId(v string) *UpdateGatewayRouteTableEntryAttributeRequest {
	s.GatewayRouteTableId = &v
	return s
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) SetIPv4GatewayRouteTableId(v string) *UpdateGatewayRouteTableEntryAttributeRequest {
	s.IPv4GatewayRouteTableId = &v
	return s
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) SetName(v string) *UpdateGatewayRouteTableEntryAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) SetNextHopId(v string) *UpdateGatewayRouteTableEntryAttributeRequest {
	s.NextHopId = &v
	return s
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) SetNextHopType(v string) *UpdateGatewayRouteTableEntryAttributeRequest {
	s.NextHopType = &v
	return s
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) SetOwnerAccount(v string) *UpdateGatewayRouteTableEntryAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) SetOwnerId(v int64) *UpdateGatewayRouteTableEntryAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) SetRegionId(v string) *UpdateGatewayRouteTableEntryAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) SetResourceOwnerAccount(v string) *UpdateGatewayRouteTableEntryAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) SetResourceOwnerId(v int64) *UpdateGatewayRouteTableEntryAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateGatewayRouteTableEntryAttributeRequest) Validate() error {
	return dara.Validate(s)
}
