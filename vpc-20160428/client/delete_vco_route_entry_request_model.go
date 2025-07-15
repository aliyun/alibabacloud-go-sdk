// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVcoRouteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteVcoRouteEntryRequest
	GetClientToken() *string
	SetNextHop(v string) *DeleteVcoRouteEntryRequest
	GetNextHop() *string
	SetOverlayMode(v string) *DeleteVcoRouteEntryRequest
	GetOverlayMode() *string
	SetOwnerAccount(v string) *DeleteVcoRouteEntryRequest
	GetOwnerAccount() *string
	SetRegionId(v string) *DeleteVcoRouteEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteVcoRouteEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteVcoRouteEntryRequest
	GetResourceOwnerId() *int64
	SetRouteDest(v string) *DeleteVcoRouteEntryRequest
	GetRouteDest() *string
	SetVpnConnectionId(v string) *DeleteVcoRouteEntryRequest
	GetVpnConnectionId() *string
	SetWeight(v int32) *DeleteVcoRouteEntryRequest
	GetWeight() *int32
}

type DeleteVcoRouteEntryRequest struct {
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
	// The next hop of the destination-based route that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// vco-p0w5112fgnl2ihlmf****
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
	// The destination CIDR block of the destination-based route that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.10.0/24
	RouteDest *string `json:"RouteDest,omitempty" xml:"RouteDest,omitempty"`
	// The ID of the IPsec-VPN attachment.
	//
	// This parameter is required.
	//
	// example:
	//
	// vco-p0w5112fgnl2ihlmf****
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" xml:"VpnConnectionId,omitempty"`
	// The weight of the destination-based route that you want to delete. Valid values:
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

func (s DeleteVcoRouteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVcoRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteVcoRouteEntryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteVcoRouteEntryRequest) GetNextHop() *string {
	return s.NextHop
}

func (s *DeleteVcoRouteEntryRequest) GetOverlayMode() *string {
	return s.OverlayMode
}

func (s *DeleteVcoRouteEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteVcoRouteEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVcoRouteEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteVcoRouteEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteVcoRouteEntryRequest) GetRouteDest() *string {
	return s.RouteDest
}

func (s *DeleteVcoRouteEntryRequest) GetVpnConnectionId() *string {
	return s.VpnConnectionId
}

func (s *DeleteVcoRouteEntryRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *DeleteVcoRouteEntryRequest) SetClientToken(v string) *DeleteVcoRouteEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVcoRouteEntryRequest) SetNextHop(v string) *DeleteVcoRouteEntryRequest {
	s.NextHop = &v
	return s
}

func (s *DeleteVcoRouteEntryRequest) SetOverlayMode(v string) *DeleteVcoRouteEntryRequest {
	s.OverlayMode = &v
	return s
}

func (s *DeleteVcoRouteEntryRequest) SetOwnerAccount(v string) *DeleteVcoRouteEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteVcoRouteEntryRequest) SetRegionId(v string) *DeleteVcoRouteEntryRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVcoRouteEntryRequest) SetResourceOwnerAccount(v string) *DeleteVcoRouteEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVcoRouteEntryRequest) SetResourceOwnerId(v int64) *DeleteVcoRouteEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVcoRouteEntryRequest) SetRouteDest(v string) *DeleteVcoRouteEntryRequest {
	s.RouteDest = &v
	return s
}

func (s *DeleteVcoRouteEntryRequest) SetVpnConnectionId(v string) *DeleteVcoRouteEntryRequest {
	s.VpnConnectionId = &v
	return s
}

func (s *DeleteVcoRouteEntryRequest) SetWeight(v int32) *DeleteVcoRouteEntryRequest {
	s.Weight = &v
	return s
}

func (s *DeleteVcoRouteEntryRequest) Validate() error {
	return dara.Validate(s)
}
