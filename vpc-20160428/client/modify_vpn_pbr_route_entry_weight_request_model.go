// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpnPbrRouteEntryWeightRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyVpnPbrRouteEntryWeightRequest
	GetClientToken() *string
	SetNewWeight(v int32) *ModifyVpnPbrRouteEntryWeightRequest
	GetNewWeight() *int32
	SetNextHop(v string) *ModifyVpnPbrRouteEntryWeightRequest
	GetNextHop() *string
	SetOverlayMode(v string) *ModifyVpnPbrRouteEntryWeightRequest
	GetOverlayMode() *string
	SetOwnerAccount(v string) *ModifyVpnPbrRouteEntryWeightRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyVpnPbrRouteEntryWeightRequest
	GetOwnerId() *int64
	SetPriority(v int32) *ModifyVpnPbrRouteEntryWeightRequest
	GetPriority() *int32
	SetRegionId(v string) *ModifyVpnPbrRouteEntryWeightRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyVpnPbrRouteEntryWeightRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyVpnPbrRouteEntryWeightRequest
	GetResourceOwnerId() *int64
	SetRouteDest(v string) *ModifyVpnPbrRouteEntryWeightRequest
	GetRouteDest() *string
	SetRouteSource(v string) *ModifyVpnPbrRouteEntryWeightRequest
	GetRouteSource() *string
	SetVpnGatewayId(v string) *ModifyVpnPbrRouteEntryWeightRequest
	GetVpnGatewayId() *string
	SetWeight(v int32) *ModifyVpnPbrRouteEntryWeightRequest
	GetWeight() *int32
}

type ModifyVpnPbrRouteEntryWeightRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// d7d24a21-f4ba-4454-9173-b3828dae492b
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The new weight of the policy-based route. Valid values:
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
	NewWeight *int32 `json:"NewWeight,omitempty" xml:"NewWeight,omitempty"`
	// The next hop of the policy-based route.
	//
	// This parameter is required.
	//
	// example:
	//
	// vco-bp15oes1py4i66rmd****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// The tunneling protocol. The value is set to **Ipsec**, which indicates the IPsec tunneling protocol.
	//
	// example:
	//
	// Ipsec
	OverlayMode  *string `json:"OverlayMode,omitempty" xml:"OverlayMode,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The priority of the policy-based route.
	//
	// 	- If the route was not assigned a priority, this parameter is optional.
	//
	// 	- If the route was assigned a priority, this parameter is optional.
	//
	//         If you specify this parameter, set the value to the priority that was assigned to the policy-based route. Otherwise, the operation fails.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the region where the VPN gateway is created. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	// 0
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ModifyVpnPbrRouteEntryWeightRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnPbrRouteEntryWeightRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) GetNewWeight() *int32 {
	return s.NewWeight
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) GetNextHop() *string {
	return s.NextHop
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) GetOverlayMode() *string {
	return s.OverlayMode
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) GetRouteDest() *string {
	return s.RouteDest
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) GetRouteSource() *string {
	return s.RouteSource
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) SetClientToken(v string) *ModifyVpnPbrRouteEntryWeightRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) SetNewWeight(v int32) *ModifyVpnPbrRouteEntryWeightRequest {
	s.NewWeight = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) SetNextHop(v string) *ModifyVpnPbrRouteEntryWeightRequest {
	s.NextHop = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) SetOverlayMode(v string) *ModifyVpnPbrRouteEntryWeightRequest {
	s.OverlayMode = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) SetOwnerAccount(v string) *ModifyVpnPbrRouteEntryWeightRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) SetOwnerId(v int64) *ModifyVpnPbrRouteEntryWeightRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) SetPriority(v int32) *ModifyVpnPbrRouteEntryWeightRequest {
	s.Priority = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) SetRegionId(v string) *ModifyVpnPbrRouteEntryWeightRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) SetResourceOwnerAccount(v string) *ModifyVpnPbrRouteEntryWeightRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) SetResourceOwnerId(v int64) *ModifyVpnPbrRouteEntryWeightRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) SetRouteDest(v string) *ModifyVpnPbrRouteEntryWeightRequest {
	s.RouteDest = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) SetRouteSource(v string) *ModifyVpnPbrRouteEntryWeightRequest {
	s.RouteSource = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) SetVpnGatewayId(v string) *ModifyVpnPbrRouteEntryWeightRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) SetWeight(v int32) *ModifyVpnPbrRouteEntryWeightRequest {
	s.Weight = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryWeightRequest) Validate() error {
	return dara.Validate(s)
}
