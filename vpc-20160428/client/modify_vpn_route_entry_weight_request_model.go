// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpnRouteEntryWeightRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyVpnRouteEntryWeightRequest
	GetClientToken() *string
	SetNewWeight(v int32) *ModifyVpnRouteEntryWeightRequest
	GetNewWeight() *int32
	SetNextHop(v string) *ModifyVpnRouteEntryWeightRequest
	GetNextHop() *string
	SetOverlayMode(v string) *ModifyVpnRouteEntryWeightRequest
	GetOverlayMode() *string
	SetOwnerAccount(v string) *ModifyVpnRouteEntryWeightRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyVpnRouteEntryWeightRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyVpnRouteEntryWeightRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyVpnRouteEntryWeightRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyVpnRouteEntryWeightRequest
	GetResourceOwnerId() *int64
	SetRouteDest(v string) *ModifyVpnRouteEntryWeightRequest
	GetRouteDest() *string
	SetVpnGatewayId(v string) *ModifyVpnRouteEntryWeightRequest
	GetVpnGatewayId() *string
	SetWeight(v int32) *ModifyVpnRouteEntryWeightRequest
	GetWeight() *int32
}

type ModifyVpnRouteEntryWeightRequest struct {
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
	// The new weight of the destination-based route. Valid values:
	//
	// 	- **0**: a low priority
	//
	// 	- **100**: a high priority
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	NewWeight *int32 `json:"NewWeight,omitempty" xml:"NewWeight,omitempty"`
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
	// The original weight of the destination-based route. Valid values:
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

func (s ModifyVpnRouteEntryWeightRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnRouteEntryWeightRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpnRouteEntryWeightRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyVpnRouteEntryWeightRequest) GetNewWeight() *int32 {
	return s.NewWeight
}

func (s *ModifyVpnRouteEntryWeightRequest) GetNextHop() *string {
	return s.NextHop
}

func (s *ModifyVpnRouteEntryWeightRequest) GetOverlayMode() *string {
	return s.OverlayMode
}

func (s *ModifyVpnRouteEntryWeightRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyVpnRouteEntryWeightRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyVpnRouteEntryWeightRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyVpnRouteEntryWeightRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyVpnRouteEntryWeightRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyVpnRouteEntryWeightRequest) GetRouteDest() *string {
	return s.RouteDest
}

func (s *ModifyVpnRouteEntryWeightRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *ModifyVpnRouteEntryWeightRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *ModifyVpnRouteEntryWeightRequest) SetClientToken(v string) *ModifyVpnRouteEntryWeightRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyVpnRouteEntryWeightRequest) SetNewWeight(v int32) *ModifyVpnRouteEntryWeightRequest {
	s.NewWeight = &v
	return s
}

func (s *ModifyVpnRouteEntryWeightRequest) SetNextHop(v string) *ModifyVpnRouteEntryWeightRequest {
	s.NextHop = &v
	return s
}

func (s *ModifyVpnRouteEntryWeightRequest) SetOverlayMode(v string) *ModifyVpnRouteEntryWeightRequest {
	s.OverlayMode = &v
	return s
}

func (s *ModifyVpnRouteEntryWeightRequest) SetOwnerAccount(v string) *ModifyVpnRouteEntryWeightRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyVpnRouteEntryWeightRequest) SetOwnerId(v int64) *ModifyVpnRouteEntryWeightRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyVpnRouteEntryWeightRequest) SetRegionId(v string) *ModifyVpnRouteEntryWeightRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyVpnRouteEntryWeightRequest) SetResourceOwnerAccount(v string) *ModifyVpnRouteEntryWeightRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyVpnRouteEntryWeightRequest) SetResourceOwnerId(v int64) *ModifyVpnRouteEntryWeightRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyVpnRouteEntryWeightRequest) SetRouteDest(v string) *ModifyVpnRouteEntryWeightRequest {
	s.RouteDest = &v
	return s
}

func (s *ModifyVpnRouteEntryWeightRequest) SetVpnGatewayId(v string) *ModifyVpnRouteEntryWeightRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *ModifyVpnRouteEntryWeightRequest) SetWeight(v int32) *ModifyVpnRouteEntryWeightRequest {
	s.Weight = &v
	return s
}

func (s *ModifyVpnRouteEntryWeightRequest) Validate() error {
	return dara.Validate(s)
}
