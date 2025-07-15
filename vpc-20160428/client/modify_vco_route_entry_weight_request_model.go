// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVcoRouteEntryWeightRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyVcoRouteEntryWeightRequest
	GetClientToken() *string
	SetNewWeight(v int32) *ModifyVcoRouteEntryWeightRequest
	GetNewWeight() *int32
	SetNextHop(v string) *ModifyVcoRouteEntryWeightRequest
	GetNextHop() *string
	SetOverlayMode(v string) *ModifyVcoRouteEntryWeightRequest
	GetOverlayMode() *string
	SetOwnerAccount(v string) *ModifyVcoRouteEntryWeightRequest
	GetOwnerAccount() *string
	SetRegionId(v string) *ModifyVcoRouteEntryWeightRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyVcoRouteEntryWeightRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyVcoRouteEntryWeightRequest
	GetResourceOwnerId() *int64
	SetRouteDest(v string) *ModifyVcoRouteEntryWeightRequest
	GetRouteDest() *string
	SetVpnConnectionId(v string) *ModifyVcoRouteEntryWeightRequest
	GetVpnConnectionId() *string
	SetWeight(v int32) *ModifyVcoRouteEntryWeightRequest
	GetWeight() *int32
}

type ModifyVcoRouteEntryWeightRequest struct {
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
	// The new weight of the destination-based route that you want to modify. Valid values:
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
	NewWeight *int32 `json:"NewWeight,omitempty" xml:"NewWeight,omitempty"`
	// The next hop of the destination-based route that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// vco-p0w2jpkhi2eeop6q6****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// The tunneling protocol. Set the value to **Ipsec**, which specifies the IPsec tunneling protocol.
	//
	// example:
	//
	// Ipsec
	OverlayMode  *string `json:"OverlayMode,omitempty" xml:"OverlayMode,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	// The region ID of the IPsec-VPN connection.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The destination CIDR block of the destination-based route that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.10.0/24
	RouteDest *string `json:"RouteDest,omitempty" xml:"RouteDest,omitempty"`
	// The ID of the IPsec-VPN connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// vco-p0w2jpkhi2eeop6q6****
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" xml:"VpnConnectionId,omitempty"`
	// The current weight of the destination-based route that you want to modify. Valid values:
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
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ModifyVcoRouteEntryWeightRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVcoRouteEntryWeightRequest) GoString() string {
	return s.String()
}

func (s *ModifyVcoRouteEntryWeightRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyVcoRouteEntryWeightRequest) GetNewWeight() *int32 {
	return s.NewWeight
}

func (s *ModifyVcoRouteEntryWeightRequest) GetNextHop() *string {
	return s.NextHop
}

func (s *ModifyVcoRouteEntryWeightRequest) GetOverlayMode() *string {
	return s.OverlayMode
}

func (s *ModifyVcoRouteEntryWeightRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyVcoRouteEntryWeightRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyVcoRouteEntryWeightRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyVcoRouteEntryWeightRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyVcoRouteEntryWeightRequest) GetRouteDest() *string {
	return s.RouteDest
}

func (s *ModifyVcoRouteEntryWeightRequest) GetVpnConnectionId() *string {
	return s.VpnConnectionId
}

func (s *ModifyVcoRouteEntryWeightRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *ModifyVcoRouteEntryWeightRequest) SetClientToken(v string) *ModifyVcoRouteEntryWeightRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyVcoRouteEntryWeightRequest) SetNewWeight(v int32) *ModifyVcoRouteEntryWeightRequest {
	s.NewWeight = &v
	return s
}

func (s *ModifyVcoRouteEntryWeightRequest) SetNextHop(v string) *ModifyVcoRouteEntryWeightRequest {
	s.NextHop = &v
	return s
}

func (s *ModifyVcoRouteEntryWeightRequest) SetOverlayMode(v string) *ModifyVcoRouteEntryWeightRequest {
	s.OverlayMode = &v
	return s
}

func (s *ModifyVcoRouteEntryWeightRequest) SetOwnerAccount(v string) *ModifyVcoRouteEntryWeightRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyVcoRouteEntryWeightRequest) SetRegionId(v string) *ModifyVcoRouteEntryWeightRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyVcoRouteEntryWeightRequest) SetResourceOwnerAccount(v string) *ModifyVcoRouteEntryWeightRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyVcoRouteEntryWeightRequest) SetResourceOwnerId(v int64) *ModifyVcoRouteEntryWeightRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyVcoRouteEntryWeightRequest) SetRouteDest(v string) *ModifyVcoRouteEntryWeightRequest {
	s.RouteDest = &v
	return s
}

func (s *ModifyVcoRouteEntryWeightRequest) SetVpnConnectionId(v string) *ModifyVcoRouteEntryWeightRequest {
	s.VpnConnectionId = &v
	return s
}

func (s *ModifyVcoRouteEntryWeightRequest) SetWeight(v int32) *ModifyVcoRouteEntryWeightRequest {
	s.Weight = &v
	return s
}

func (s *ModifyVcoRouteEntryWeightRequest) Validate() error {
	return dara.Validate(s)
}
