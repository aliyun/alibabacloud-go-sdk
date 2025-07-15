// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpnPbrRouteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateVpnPbrRouteEntryRequest
	GetClientToken() *string
	SetDescription(v string) *CreateVpnPbrRouteEntryRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateVpnPbrRouteEntryRequest
	GetDryRun() *bool
	SetNextHop(v string) *CreateVpnPbrRouteEntryRequest
	GetNextHop() *string
	SetOverlayMode(v string) *CreateVpnPbrRouteEntryRequest
	GetOverlayMode() *string
	SetOwnerAccount(v string) *CreateVpnPbrRouteEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateVpnPbrRouteEntryRequest
	GetOwnerId() *int64
	SetPriority(v int32) *CreateVpnPbrRouteEntryRequest
	GetPriority() *int32
	SetPublishVpc(v bool) *CreateVpnPbrRouteEntryRequest
	GetPublishVpc() *bool
	SetRegionId(v string) *CreateVpnPbrRouteEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateVpnPbrRouteEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateVpnPbrRouteEntryRequest
	GetResourceOwnerId() *int64
	SetRouteDest(v string) *CreateVpnPbrRouteEntryRequest
	GetRouteDest() *string
	SetRouteSource(v string) *CreateVpnPbrRouteEntryRequest
	GetRouteSource() *string
	SetVpnGatewayId(v string) *CreateVpnPbrRouteEntryRequest
	GetVpnGatewayId() *string
	SetWeight(v int32) *CreateVpnPbrRouteEntryRequest
	GetWeight() *int32
}

type CreateVpnPbrRouteEntryRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate a token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID is different for each request.
	//
	// example:
	//
	// d7d24a21-f4ba-4454-9173-b3****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the policy-based route.
	//
	// The description must be 1 to 100 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
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
	// The priority of the policy-based route. Valid values: **1*	- to **100**. Default value: **10**.
	//
	// A smaller value indicates a higher priority.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Specifies whether to advertise the policy-based route to a virtual private cloud (VPC) route table. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	PublishVpc *bool `json:"PublishVpc,omitempty" xml:"PublishVpc,omitempty"`
	// The region ID of the VPN gateway. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
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
	// The VPN gateway ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpn-bp1a3kqjiiq9legfx****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
	// The weight of the policy-based route.
	//
	// If you use the same VPN gateway to establish active/standby IPsec-VPN connections, you can configure route weights to specify which connection is active. A value of 100 specifies the active connection, whereas a value of 0 specifies the standby connection.
	//
	// You can configure health checks to automatically check the connectivity of IPsec-VPN connections. If the active connection is down, the standby connection automatically takes over. For more information, see [CreateVpnConnection](https://help.aliyun.com/document_detail/120391.html).
	//
	// 	- **100**: The IPsec-VPN connection associated with the policy-based route serves as an active connection.
	//
	// 	- **0**: The IPsec-VPN connection associated with the policy-based route serves as a standby connection.
	//
	// >  If you specify active/standby IPsec-VPN connections, the active policy-based route and the standby policy-based route must have the same source and destination CIDR blocks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateVpnPbrRouteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVpnPbrRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *CreateVpnPbrRouteEntryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVpnPbrRouteEntryRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVpnPbrRouteEntryRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateVpnPbrRouteEntryRequest) GetNextHop() *string {
	return s.NextHop
}

func (s *CreateVpnPbrRouteEntryRequest) GetOverlayMode() *string {
	return s.OverlayMode
}

func (s *CreateVpnPbrRouteEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateVpnPbrRouteEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateVpnPbrRouteEntryRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateVpnPbrRouteEntryRequest) GetPublishVpc() *bool {
	return s.PublishVpc
}

func (s *CreateVpnPbrRouteEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVpnPbrRouteEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateVpnPbrRouteEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateVpnPbrRouteEntryRequest) GetRouteDest() *string {
	return s.RouteDest
}

func (s *CreateVpnPbrRouteEntryRequest) GetRouteSource() *string {
	return s.RouteSource
}

func (s *CreateVpnPbrRouteEntryRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *CreateVpnPbrRouteEntryRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateVpnPbrRouteEntryRequest) SetClientToken(v string) *CreateVpnPbrRouteEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVpnPbrRouteEntryRequest) SetDescription(v string) *CreateVpnPbrRouteEntryRequest {
	s.Description = &v
	return s
}

func (s *CreateVpnPbrRouteEntryRequest) SetDryRun(v bool) *CreateVpnPbrRouteEntryRequest {
	s.DryRun = &v
	return s
}

func (s *CreateVpnPbrRouteEntryRequest) SetNextHop(v string) *CreateVpnPbrRouteEntryRequest {
	s.NextHop = &v
	return s
}

func (s *CreateVpnPbrRouteEntryRequest) SetOverlayMode(v string) *CreateVpnPbrRouteEntryRequest {
	s.OverlayMode = &v
	return s
}

func (s *CreateVpnPbrRouteEntryRequest) SetOwnerAccount(v string) *CreateVpnPbrRouteEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateVpnPbrRouteEntryRequest) SetOwnerId(v int64) *CreateVpnPbrRouteEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVpnPbrRouteEntryRequest) SetPriority(v int32) *CreateVpnPbrRouteEntryRequest {
	s.Priority = &v
	return s
}

func (s *CreateVpnPbrRouteEntryRequest) SetPublishVpc(v bool) *CreateVpnPbrRouteEntryRequest {
	s.PublishVpc = &v
	return s
}

func (s *CreateVpnPbrRouteEntryRequest) SetRegionId(v string) *CreateVpnPbrRouteEntryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVpnPbrRouteEntryRequest) SetResourceOwnerAccount(v string) *CreateVpnPbrRouteEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateVpnPbrRouteEntryRequest) SetResourceOwnerId(v int64) *CreateVpnPbrRouteEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateVpnPbrRouteEntryRequest) SetRouteDest(v string) *CreateVpnPbrRouteEntryRequest {
	s.RouteDest = &v
	return s
}

func (s *CreateVpnPbrRouteEntryRequest) SetRouteSource(v string) *CreateVpnPbrRouteEntryRequest {
	s.RouteSource = &v
	return s
}

func (s *CreateVpnPbrRouteEntryRequest) SetVpnGatewayId(v string) *CreateVpnPbrRouteEntryRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *CreateVpnPbrRouteEntryRequest) SetWeight(v int32) *CreateVpnPbrRouteEntryRequest {
	s.Weight = &v
	return s
}

func (s *CreateVpnPbrRouteEntryRequest) Validate() error {
	return dara.Validate(s)
}
