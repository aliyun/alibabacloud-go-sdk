// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpnPbrRouteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteVpnPbrRouteEntryRequest
	GetClientToken() *string
	SetNextHop(v string) *DeleteVpnPbrRouteEntryRequest
	GetNextHop() *string
	SetOverlayMode(v string) *DeleteVpnPbrRouteEntryRequest
	GetOverlayMode() *string
	SetOwnerAccount(v string) *DeleteVpnPbrRouteEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteVpnPbrRouteEntryRequest
	GetOwnerId() *int64
	SetPriority(v int32) *DeleteVpnPbrRouteEntryRequest
	GetPriority() *int32
	SetRegionId(v string) *DeleteVpnPbrRouteEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteVpnPbrRouteEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteVpnPbrRouteEntryRequest
	GetResourceOwnerId() *int64
	SetRouteDest(v string) *DeleteVpnPbrRouteEntryRequest
	GetRouteDest() *string
	SetRouteSource(v string) *DeleteVpnPbrRouteEntryRequest
	GetRouteSource() *string
	SetVpnGatewayId(v string) *DeleteVpnPbrRouteEntryRequest
	GetVpnGatewayId() *string
	SetWeight(v int32) *DeleteVpnPbrRouteEntryRequest
	GetWeight() *int32
}

type DeleteVpnPbrRouteEntryRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-4266****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The next hop of the policy-based route.
	//
	// This parameter is required.
	//
	// example:
	//
	// vco-bp15oes1py4i66rmd****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// The tunneling protocol. Set the value to **Ipsec**.
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
	//         If you specify this parameter, set the value to the priority configured for the policy-based route. Otherwise, the operation cannot be performed.
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
	// The weight of the policy-based route. Valid values:
	//
	// You can call [DescribeVpnPbrRouteEntries](https://help.aliyun.com/document_detail/2526959.html) to query weights of policy-based routes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DeleteVpnPbrRouteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpnPbrRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpnPbrRouteEntryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteVpnPbrRouteEntryRequest) GetNextHop() *string {
	return s.NextHop
}

func (s *DeleteVpnPbrRouteEntryRequest) GetOverlayMode() *string {
	return s.OverlayMode
}

func (s *DeleteVpnPbrRouteEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteVpnPbrRouteEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteVpnPbrRouteEntryRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *DeleteVpnPbrRouteEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVpnPbrRouteEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteVpnPbrRouteEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteVpnPbrRouteEntryRequest) GetRouteDest() *string {
	return s.RouteDest
}

func (s *DeleteVpnPbrRouteEntryRequest) GetRouteSource() *string {
	return s.RouteSource
}

func (s *DeleteVpnPbrRouteEntryRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *DeleteVpnPbrRouteEntryRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *DeleteVpnPbrRouteEntryRequest) SetClientToken(v string) *DeleteVpnPbrRouteEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVpnPbrRouteEntryRequest) SetNextHop(v string) *DeleteVpnPbrRouteEntryRequest {
	s.NextHop = &v
	return s
}

func (s *DeleteVpnPbrRouteEntryRequest) SetOverlayMode(v string) *DeleteVpnPbrRouteEntryRequest {
	s.OverlayMode = &v
	return s
}

func (s *DeleteVpnPbrRouteEntryRequest) SetOwnerAccount(v string) *DeleteVpnPbrRouteEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteVpnPbrRouteEntryRequest) SetOwnerId(v int64) *DeleteVpnPbrRouteEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVpnPbrRouteEntryRequest) SetPriority(v int32) *DeleteVpnPbrRouteEntryRequest {
	s.Priority = &v
	return s
}

func (s *DeleteVpnPbrRouteEntryRequest) SetRegionId(v string) *DeleteVpnPbrRouteEntryRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVpnPbrRouteEntryRequest) SetResourceOwnerAccount(v string) *DeleteVpnPbrRouteEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVpnPbrRouteEntryRequest) SetResourceOwnerId(v int64) *DeleteVpnPbrRouteEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVpnPbrRouteEntryRequest) SetRouteDest(v string) *DeleteVpnPbrRouteEntryRequest {
	s.RouteDest = &v
	return s
}

func (s *DeleteVpnPbrRouteEntryRequest) SetRouteSource(v string) *DeleteVpnPbrRouteEntryRequest {
	s.RouteSource = &v
	return s
}

func (s *DeleteVpnPbrRouteEntryRequest) SetVpnGatewayId(v string) *DeleteVpnPbrRouteEntryRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *DeleteVpnPbrRouteEntryRequest) SetWeight(v int32) *DeleteVpnPbrRouteEntryRequest {
	s.Weight = &v
	return s
}

func (s *DeleteVpnPbrRouteEntryRequest) Validate() error {
	return dara.Validate(s)
}
