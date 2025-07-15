// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpnPbrRouteEntryAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyVpnPbrRouteEntryAttributeRequest
	GetClientToken() *string
	SetNewPriority(v int32) *ModifyVpnPbrRouteEntryAttributeRequest
	GetNewPriority() *int32
	SetNewWeight(v int32) *ModifyVpnPbrRouteEntryAttributeRequest
	GetNewWeight() *int32
	SetNextHop(v string) *ModifyVpnPbrRouteEntryAttributeRequest
	GetNextHop() *string
	SetOwnerAccount(v string) *ModifyVpnPbrRouteEntryAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyVpnPbrRouteEntryAttributeRequest
	GetOwnerId() *int64
	SetPriority(v int32) *ModifyVpnPbrRouteEntryAttributeRequest
	GetPriority() *int32
	SetRegionId(v string) *ModifyVpnPbrRouteEntryAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyVpnPbrRouteEntryAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyVpnPbrRouteEntryAttributeRequest
	GetResourceOwnerId() *int64
	SetRouteDest(v string) *ModifyVpnPbrRouteEntryAttributeRequest
	GetRouteDest() *string
	SetRouteSource(v string) *ModifyVpnPbrRouteEntryAttributeRequest
	GetRouteSource() *string
	SetVpnGatewayId(v string) *ModifyVpnPbrRouteEntryAttributeRequest
	GetVpnGatewayId() *string
	SetWeight(v int32) *ModifyVpnPbrRouteEntryAttributeRequest
	GetWeight() *int32
}

type ModifyVpnPbrRouteEntryAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// d7d24a21-f4ba-4454-9173-b3****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The new priority of the policy-based route. Valid values: **1*	- to **100**.
	//
	// A smaller value indicates a higher priority.
	//
	// If you do not specify this parameter, the priority of the policy-based route is not modified.
	//
	// >  You must specify at least one of **NewPriority*	- and **NewWeight**.
	//
	// example:
	//
	// 10
	NewPriority *int32 `json:"NewPriority,omitempty" xml:"NewPriority,omitempty"`
	// The new weight of the policy-based route. Valid values:
	//
	// 	- **100**: The IPsec-VPN connection associated with the policy-based route serves as the active connection.
	//
	// 	- **0**: The IPsec-VPN connection associated with the policy-based route serves as the standby connection.
	//
	// If you do not specify this parameter, the weight of the policy-based route is not modified.
	//
	// >  You must specify at least one of **NewPriority*	- and **NewWeight**.
	//
	// example:
	//
	// 0
	NewWeight *int32 `json:"NewWeight,omitempty" xml:"NewWeight,omitempty"`
	// The next hop of the policy-based route.
	//
	// This parameter is required.
	//
	// example:
	//
	// vco-bp15oes1py4i66rmd****
	NextHop      *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The original priority of the policy-based route. Valid values: **1*	- to **100**.
	//
	// A smaller value indicates a higher priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The region ID of the VPN gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-heyuan
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The destination CIDR block of the policy-based route.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.0.0/24
	RouteDest *string `json:"RouteDest,omitempty" xml:"RouteDest,omitempty"`
	// The source CIDR block of the policy-based route.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.1.0/24
	RouteSource *string `json:"RouteSource,omitempty" xml:"RouteSource,omitempty"`
	// The ID of the VPN gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpn-bp1a3kqjiiq9legfx****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
	// The original weight of the policy-based route. Valid values:
	//
	// 	- **100**: The IPsec-VPN connection associated with the policy-based route serves as an active connection.
	//
	// 	- **0**: The IPsec-VPN connection associated with the policy-based route serves as a standby connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ModifyVpnPbrRouteEntryAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnPbrRouteEntryAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) GetNewPriority() *int32 {
	return s.NewPriority
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) GetNewWeight() *int32 {
	return s.NewWeight
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) GetNextHop() *string {
	return s.NextHop
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) GetRouteDest() *string {
	return s.RouteDest
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) GetRouteSource() *string {
	return s.RouteSource
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) SetClientToken(v string) *ModifyVpnPbrRouteEntryAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) SetNewPriority(v int32) *ModifyVpnPbrRouteEntryAttributeRequest {
	s.NewPriority = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) SetNewWeight(v int32) *ModifyVpnPbrRouteEntryAttributeRequest {
	s.NewWeight = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) SetNextHop(v string) *ModifyVpnPbrRouteEntryAttributeRequest {
	s.NextHop = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) SetOwnerAccount(v string) *ModifyVpnPbrRouteEntryAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) SetOwnerId(v int64) *ModifyVpnPbrRouteEntryAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) SetPriority(v int32) *ModifyVpnPbrRouteEntryAttributeRequest {
	s.Priority = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) SetRegionId(v string) *ModifyVpnPbrRouteEntryAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) SetResourceOwnerAccount(v string) *ModifyVpnPbrRouteEntryAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) SetResourceOwnerId(v int64) *ModifyVpnPbrRouteEntryAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) SetRouteDest(v string) *ModifyVpnPbrRouteEntryAttributeRequest {
	s.RouteDest = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) SetRouteSource(v string) *ModifyVpnPbrRouteEntryAttributeRequest {
	s.RouteSource = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) SetVpnGatewayId(v string) *ModifyVpnPbrRouteEntryAttributeRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) SetWeight(v int32) *ModifyVpnPbrRouteEntryAttributeRequest {
	s.Weight = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryAttributeRequest) Validate() error {
	return dara.Validate(s)
}
