// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpnPbrRouteEntryPriorityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyVpnPbrRouteEntryPriorityRequest
	GetClientToken() *string
	SetNewPriority(v int32) *ModifyVpnPbrRouteEntryPriorityRequest
	GetNewPriority() *int32
	SetNextHop(v string) *ModifyVpnPbrRouteEntryPriorityRequest
	GetNextHop() *string
	SetOwnerAccount(v string) *ModifyVpnPbrRouteEntryPriorityRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyVpnPbrRouteEntryPriorityRequest
	GetOwnerId() *int64
	SetPriority(v int32) *ModifyVpnPbrRouteEntryPriorityRequest
	GetPriority() *int32
	SetRegionId(v string) *ModifyVpnPbrRouteEntryPriorityRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyVpnPbrRouteEntryPriorityRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyVpnPbrRouteEntryPriorityRequest
	GetResourceOwnerId() *int64
	SetRouteDest(v string) *ModifyVpnPbrRouteEntryPriorityRequest
	GetRouteDest() *string
	SetRouteSource(v string) *ModifyVpnPbrRouteEntryPriorityRequest
	GetRouteSource() *string
	SetVpnGatewayId(v string) *ModifyVpnPbrRouteEntryPriorityRequest
	GetVpnGatewayId() *string
	SetWeight(v int32) *ModifyVpnPbrRouteEntryPriorityRequest
	GetWeight() *int32
}

type ModifyVpnPbrRouteEntryPriorityRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// 10
	NewPriority *int32 `json:"NewPriority,omitempty" xml:"NewPriority,omitempty"`
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
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the region where the VPN gateway is created.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-henyuan
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
	// This parameter is required.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ModifyVpnPbrRouteEntryPriorityRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnPbrRouteEntryPriorityRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) GetNewPriority() *int32 {
	return s.NewPriority
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) GetNextHop() *string {
	return s.NextHop
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) GetRouteDest() *string {
	return s.RouteDest
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) GetRouteSource() *string {
	return s.RouteSource
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) SetClientToken(v string) *ModifyVpnPbrRouteEntryPriorityRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) SetNewPriority(v int32) *ModifyVpnPbrRouteEntryPriorityRequest {
	s.NewPriority = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) SetNextHop(v string) *ModifyVpnPbrRouteEntryPriorityRequest {
	s.NextHop = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) SetOwnerAccount(v string) *ModifyVpnPbrRouteEntryPriorityRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) SetOwnerId(v int64) *ModifyVpnPbrRouteEntryPriorityRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) SetPriority(v int32) *ModifyVpnPbrRouteEntryPriorityRequest {
	s.Priority = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) SetRegionId(v string) *ModifyVpnPbrRouteEntryPriorityRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) SetResourceOwnerAccount(v string) *ModifyVpnPbrRouteEntryPriorityRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) SetResourceOwnerId(v int64) *ModifyVpnPbrRouteEntryPriorityRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) SetRouteDest(v string) *ModifyVpnPbrRouteEntryPriorityRequest {
	s.RouteDest = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) SetRouteSource(v string) *ModifyVpnPbrRouteEntryPriorityRequest {
	s.RouteSource = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) SetVpnGatewayId(v string) *ModifyVpnPbrRouteEntryPriorityRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) SetWeight(v int32) *ModifyVpnPbrRouteEntryPriorityRequest {
	s.Weight = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryPriorityRequest) Validate() error {
	return dara.Validate(s)
}
