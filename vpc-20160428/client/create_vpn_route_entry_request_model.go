// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpnRouteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateVpnRouteEntryRequest
	GetClientToken() *string
	SetDescription(v string) *CreateVpnRouteEntryRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateVpnRouteEntryRequest
	GetDryRun() *bool
	SetNextHop(v string) *CreateVpnRouteEntryRequest
	GetNextHop() *string
	SetOverlayMode(v string) *CreateVpnRouteEntryRequest
	GetOverlayMode() *string
	SetOwnerAccount(v string) *CreateVpnRouteEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateVpnRouteEntryRequest
	GetOwnerId() *int64
	SetPublishVpc(v bool) *CreateVpnRouteEntryRequest
	GetPublishVpc() *bool
	SetRegionId(v string) *CreateVpnRouteEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateVpnRouteEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateVpnRouteEntryRequest
	GetResourceOwnerId() *int64
	SetRouteDest(v string) *CreateVpnRouteEntryRequest
	GetRouteDest() *string
	SetVpnGatewayId(v string) *CreateVpnRouteEntryRequest
	GetVpnGatewayId() *string
	SetWeight(v int32) *CreateVpnRouteEntryRequest
	GetWeight() *int32
}

type CreateVpnRouteEntryRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// d7d24a21-f4ba-4454-9173-b3828dae****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the destination route.
	//
	// The description must be **1*	- to **100*	- characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// mytest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without creating the route. The system checks the required parameters, request syntax, and business restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code `DryRunOperation` is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. If the check passes, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The next hop of the destination route.
	//
	// This parameter is required.
	//
	// example:
	//
	// vco-bp15oes1py4i66rmd****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// The tunneling protocol. Set the value to **Ipsec*	- (IPsec tunneling).
	//
	// example:
	//
	// Ipsec
	OverlayMode  *string `json:"OverlayMode,omitempty" xml:"OverlayMode,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to publish the destination route to the VPC route table. Valid values:
	//
	// - **true**: Publishes the destination route to the VPC route table. The system publishes the route only to the VPC system route table, not to VPC custom route tables.
	//
	//   If you want the VPC custom route table to contain this route, manually add the route. For more information, see [CreateRouteEntry](https://help.aliyun.com/document_detail/448722.html).
	//
	// - **false**: Does not publish the destination route to the VPC route table.
	//
	//   You must manually add a destination route whose next hop points to the VPN gateway instance in both the VPC system route table and custom route tables. Otherwise, the VPC cannot access resources in the destination CIDR block through the IPsec-VPN connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	PublishVpc *bool `json:"PublishVpc,omitempty" xml:"PublishVpc,omitempty"`
	// The region ID of the VPN gateway instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The destination CIDR block of the destination route.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.0.0/24
	RouteDest *string `json:"RouteDest,omitempty" xml:"RouteDest,omitempty"`
	// The instance ID of the VPN gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpn-bp1a3kqjiiq9legfx****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
	// The weight of the destination route.
	//
	// If you use the same VPN gateway instance to establish active/standby IPsec-VPN connections, you can specify the active and standby links by configuring the weight of the destination routing. A destination route with a weight of 100 is the active link by default, and a destination route with a weight of 0 is the standby link by default.
	//
	// You can configure health checks for the IPsec-VPN connection to automatically detect the connectivity of the link. If the active link is unavailable, the system automatically switches traffic to the standby link, which ensures high availability of the cloud connection. For more information, see [CreateVpnConnection](https://help.aliyun.com/document_detail/120391.html).
	//
	// - **100**: The IPsec-VPN connection associated with the destination route serves as the active link.
	//
	// - **0**: The IPsec-VPN connection associated with the destination route serves as the standby link.
	//
	// > - When you specify active and standby links, the active and standby destination routes must have the same destination CIDR block, different next hops, and different weights.
	//
	// > - For VPN gateway instances that support dual-tunnel pattern IPsec-VPN connections, you do not need to configure this parameter. A dual-tunnel pattern IPsec-VPN connection contains two tunnels that automatically form active/standby links. You do not need to specify active/standby links by using the parameter settings. If you configure this parameter, the parameter settings do not take effect.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateVpnRouteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVpnRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *CreateVpnRouteEntryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVpnRouteEntryRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVpnRouteEntryRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateVpnRouteEntryRequest) GetNextHop() *string {
	return s.NextHop
}

func (s *CreateVpnRouteEntryRequest) GetOverlayMode() *string {
	return s.OverlayMode
}

func (s *CreateVpnRouteEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateVpnRouteEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateVpnRouteEntryRequest) GetPublishVpc() *bool {
	return s.PublishVpc
}

func (s *CreateVpnRouteEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVpnRouteEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateVpnRouteEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateVpnRouteEntryRequest) GetRouteDest() *string {
	return s.RouteDest
}

func (s *CreateVpnRouteEntryRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *CreateVpnRouteEntryRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateVpnRouteEntryRequest) SetClientToken(v string) *CreateVpnRouteEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVpnRouteEntryRequest) SetDescription(v string) *CreateVpnRouteEntryRequest {
	s.Description = &v
	return s
}

func (s *CreateVpnRouteEntryRequest) SetDryRun(v bool) *CreateVpnRouteEntryRequest {
	s.DryRun = &v
	return s
}

func (s *CreateVpnRouteEntryRequest) SetNextHop(v string) *CreateVpnRouteEntryRequest {
	s.NextHop = &v
	return s
}

func (s *CreateVpnRouteEntryRequest) SetOverlayMode(v string) *CreateVpnRouteEntryRequest {
	s.OverlayMode = &v
	return s
}

func (s *CreateVpnRouteEntryRequest) SetOwnerAccount(v string) *CreateVpnRouteEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateVpnRouteEntryRequest) SetOwnerId(v int64) *CreateVpnRouteEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVpnRouteEntryRequest) SetPublishVpc(v bool) *CreateVpnRouteEntryRequest {
	s.PublishVpc = &v
	return s
}

func (s *CreateVpnRouteEntryRequest) SetRegionId(v string) *CreateVpnRouteEntryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVpnRouteEntryRequest) SetResourceOwnerAccount(v string) *CreateVpnRouteEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateVpnRouteEntryRequest) SetResourceOwnerId(v int64) *CreateVpnRouteEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateVpnRouteEntryRequest) SetRouteDest(v string) *CreateVpnRouteEntryRequest {
	s.RouteDest = &v
	return s
}

func (s *CreateVpnRouteEntryRequest) SetVpnGatewayId(v string) *CreateVpnRouteEntryRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *CreateVpnRouteEntryRequest) SetWeight(v int32) *CreateVpnRouteEntryRequest {
	s.Weight = &v
	return s
}

func (s *CreateVpnRouteEntryRequest) Validate() error {
	return dara.Validate(s)
}
