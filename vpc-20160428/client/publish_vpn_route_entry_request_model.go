// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishVpnRouteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *PublishVpnRouteEntryRequest
	GetClientToken() *string
	SetNextHop(v string) *PublishVpnRouteEntryRequest
	GetNextHop() *string
	SetOwnerAccount(v string) *PublishVpnRouteEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *PublishVpnRouteEntryRequest
	GetOwnerId() *int64
	SetPublishVpc(v bool) *PublishVpnRouteEntryRequest
	GetPublishVpc() *bool
	SetRegionId(v string) *PublishVpnRouteEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *PublishVpnRouteEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *PublishVpnRouteEntryRequest
	GetResourceOwnerId() *int64
	SetRouteDest(v string) *PublishVpnRouteEntryRequest
	GetRouteDest() *string
	SetRouteType(v string) *PublishVpnRouteEntryRequest
	GetRouteType() *string
	SetVpnGatewayId(v string) *PublishVpnRouteEntryRequest
	GetVpnGatewayId() *string
}

type PublishVpnRouteEntryRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// d7d24a21-f4ba-4454-9173-b382****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The next hop of the VPN gateway route.
	//
	// This parameter is required.
	//
	// example:
	//
	// vco-bp15oes1py4i66rmd****
	NextHop      *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to advertise the VPN gateway route to the VPC route table. Valid values:
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
	// The ID of the region where the VPN gateway is created.
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
	// The destination CIDR block of the VPN gateway route.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.0.0/24
	RouteDest *string `json:"RouteDest,omitempty" xml:"RouteDest,omitempty"`
	// The type of the VPN gateway route. Valid values:
	//
	// 	- **pbr**: policy-based route
	//
	// 	- **dbr**: destination-based route
	//
	// This parameter is required.
	//
	// example:
	//
	// pbr
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	// The ID of the VPN gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpn-bp1a3kqjiiq9legfx****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s PublishVpnRouteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishVpnRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *PublishVpnRouteEntryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *PublishVpnRouteEntryRequest) GetNextHop() *string {
	return s.NextHop
}

func (s *PublishVpnRouteEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *PublishVpnRouteEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PublishVpnRouteEntryRequest) GetPublishVpc() *bool {
	return s.PublishVpc
}

func (s *PublishVpnRouteEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PublishVpnRouteEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *PublishVpnRouteEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *PublishVpnRouteEntryRequest) GetRouteDest() *string {
	return s.RouteDest
}

func (s *PublishVpnRouteEntryRequest) GetRouteType() *string {
	return s.RouteType
}

func (s *PublishVpnRouteEntryRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *PublishVpnRouteEntryRequest) SetClientToken(v string) *PublishVpnRouteEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *PublishVpnRouteEntryRequest) SetNextHop(v string) *PublishVpnRouteEntryRequest {
	s.NextHop = &v
	return s
}

func (s *PublishVpnRouteEntryRequest) SetOwnerAccount(v string) *PublishVpnRouteEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *PublishVpnRouteEntryRequest) SetOwnerId(v int64) *PublishVpnRouteEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *PublishVpnRouteEntryRequest) SetPublishVpc(v bool) *PublishVpnRouteEntryRequest {
	s.PublishVpc = &v
	return s
}

func (s *PublishVpnRouteEntryRequest) SetRegionId(v string) *PublishVpnRouteEntryRequest {
	s.RegionId = &v
	return s
}

func (s *PublishVpnRouteEntryRequest) SetResourceOwnerAccount(v string) *PublishVpnRouteEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PublishVpnRouteEntryRequest) SetResourceOwnerId(v int64) *PublishVpnRouteEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PublishVpnRouteEntryRequest) SetRouteDest(v string) *PublishVpnRouteEntryRequest {
	s.RouteDest = &v
	return s
}

func (s *PublishVpnRouteEntryRequest) SetRouteType(v string) *PublishVpnRouteEntryRequest {
	s.RouteType = &v
	return s
}

func (s *PublishVpnRouteEntryRequest) SetVpnGatewayId(v string) *PublishVpnRouteEntryRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *PublishVpnRouteEntryRequest) Validate() error {
	return dara.Validate(s)
}
