// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpnRouteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteVpnRouteEntryRequest
	GetClientToken() *string
	SetNextHop(v string) *DeleteVpnRouteEntryRequest
	GetNextHop() *string
	SetOverlayMode(v string) *DeleteVpnRouteEntryRequest
	GetOverlayMode() *string
	SetOwnerAccount(v string) *DeleteVpnRouteEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteVpnRouteEntryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteVpnRouteEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteVpnRouteEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteVpnRouteEntryRequest
	GetResourceOwnerId() *int64
	SetRouteDest(v string) *DeleteVpnRouteEntryRequest
	GetRouteDest() *string
	SetVpnGatewayId(v string) *DeleteVpnRouteEntryRequest
	GetVpnGatewayId() *string
	SetWeight(v int32) *DeleteVpnRouteEntryRequest
	GetWeight() *int32
}

type DeleteVpnRouteEntryRequest struct {
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
	// The next hop of the destination-based route.
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
	// The destination CIDR block of the destination-based route.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.0.0/24
	RouteDest *string `json:"RouteDest,omitempty" xml:"RouteDest,omitempty"`
	// The ID of the VPN gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpn-bp1a3kqjiiq9legfx****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
	// The weight of the destination-based route. Valid values:
	//
	// 	- **0**: a low priority
	//
	// 	- **100**: a high priority
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DeleteVpnRouteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpnRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpnRouteEntryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteVpnRouteEntryRequest) GetNextHop() *string {
	return s.NextHop
}

func (s *DeleteVpnRouteEntryRequest) GetOverlayMode() *string {
	return s.OverlayMode
}

func (s *DeleteVpnRouteEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteVpnRouteEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteVpnRouteEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVpnRouteEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteVpnRouteEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteVpnRouteEntryRequest) GetRouteDest() *string {
	return s.RouteDest
}

func (s *DeleteVpnRouteEntryRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *DeleteVpnRouteEntryRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *DeleteVpnRouteEntryRequest) SetClientToken(v string) *DeleteVpnRouteEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVpnRouteEntryRequest) SetNextHop(v string) *DeleteVpnRouteEntryRequest {
	s.NextHop = &v
	return s
}

func (s *DeleteVpnRouteEntryRequest) SetOverlayMode(v string) *DeleteVpnRouteEntryRequest {
	s.OverlayMode = &v
	return s
}

func (s *DeleteVpnRouteEntryRequest) SetOwnerAccount(v string) *DeleteVpnRouteEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteVpnRouteEntryRequest) SetOwnerId(v int64) *DeleteVpnRouteEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVpnRouteEntryRequest) SetRegionId(v string) *DeleteVpnRouteEntryRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVpnRouteEntryRequest) SetResourceOwnerAccount(v string) *DeleteVpnRouteEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVpnRouteEntryRequest) SetResourceOwnerId(v int64) *DeleteVpnRouteEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVpnRouteEntryRequest) SetRouteDest(v string) *DeleteVpnRouteEntryRequest {
	s.RouteDest = &v
	return s
}

func (s *DeleteVpnRouteEntryRequest) SetVpnGatewayId(v string) *DeleteVpnRouteEntryRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *DeleteVpnRouteEntryRequest) SetWeight(v int32) *DeleteVpnRouteEntryRequest {
	s.Weight = &v
	return s
}

func (s *DeleteVpnRouteEntryRequest) Validate() error {
	return dara.Validate(s)
}
