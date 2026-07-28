// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVcoRouteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateVcoRouteEntryRequest
	GetClientToken() *string
	SetDescription(v string) *CreateVcoRouteEntryRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateVcoRouteEntryRequest
	GetDryRun() *bool
	SetNextHop(v string) *CreateVcoRouteEntryRequest
	GetNextHop() *string
	SetOverlayMode(v string) *CreateVcoRouteEntryRequest
	GetOverlayMode() *string
	SetOwnerAccount(v string) *CreateVcoRouteEntryRequest
	GetOwnerAccount() *string
	SetRegionId(v string) *CreateVcoRouteEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateVcoRouteEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateVcoRouteEntryRequest
	GetResourceOwnerId() *int64
	SetRouteDest(v string) *CreateVcoRouteEntryRequest
	GetRouteDest() *string
	SetVpnConnectionId(v string) *CreateVcoRouteEntryRequest
	GetVpnConnectionId() *string
	SetWeight(v int32) *CreateVcoRouteEntryRequest
	GetWeight() *int32
}

type CreateVcoRouteEntryRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-4266****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the destination route entry.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run. The system checks the required parameters, request syntax, and business restrictions. If the check fails, the corresponding error is returned. If the check succeeds, the error code `DryRunOperation` is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. If the check succeeds, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The next hop of the destination route entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// vco-p0w2jpkhi2eeop6q6****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// The tunneling protocol. Set the value to **Ipsec*	- (default), which specifies the IPsec tunneling protocol.
	//
	// example:
	//
	// Ipsec
	OverlayMode  *string `json:"OverlayMode,omitempty" xml:"OverlayMode,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	// The region ID of the IPsec-VPN connection.
	//
	// You can call [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The destination CIDR block of the destination route entry.
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
	// The weight of the destination route entry. Valid values:
	//
	// - **0**: low priority.
	//
	// - **100**: high priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateVcoRouteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVcoRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *CreateVcoRouteEntryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVcoRouteEntryRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVcoRouteEntryRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateVcoRouteEntryRequest) GetNextHop() *string {
	return s.NextHop
}

func (s *CreateVcoRouteEntryRequest) GetOverlayMode() *string {
	return s.OverlayMode
}

func (s *CreateVcoRouteEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateVcoRouteEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVcoRouteEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateVcoRouteEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateVcoRouteEntryRequest) GetRouteDest() *string {
	return s.RouteDest
}

func (s *CreateVcoRouteEntryRequest) GetVpnConnectionId() *string {
	return s.VpnConnectionId
}

func (s *CreateVcoRouteEntryRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateVcoRouteEntryRequest) SetClientToken(v string) *CreateVcoRouteEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVcoRouteEntryRequest) SetDescription(v string) *CreateVcoRouteEntryRequest {
	s.Description = &v
	return s
}

func (s *CreateVcoRouteEntryRequest) SetDryRun(v bool) *CreateVcoRouteEntryRequest {
	s.DryRun = &v
	return s
}

func (s *CreateVcoRouteEntryRequest) SetNextHop(v string) *CreateVcoRouteEntryRequest {
	s.NextHop = &v
	return s
}

func (s *CreateVcoRouteEntryRequest) SetOverlayMode(v string) *CreateVcoRouteEntryRequest {
	s.OverlayMode = &v
	return s
}

func (s *CreateVcoRouteEntryRequest) SetOwnerAccount(v string) *CreateVcoRouteEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateVcoRouteEntryRequest) SetRegionId(v string) *CreateVcoRouteEntryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVcoRouteEntryRequest) SetResourceOwnerAccount(v string) *CreateVcoRouteEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateVcoRouteEntryRequest) SetResourceOwnerId(v int64) *CreateVcoRouteEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateVcoRouteEntryRequest) SetRouteDest(v string) *CreateVcoRouteEntryRequest {
	s.RouteDest = &v
	return s
}

func (s *CreateVcoRouteEntryRequest) SetVpnConnectionId(v string) *CreateVcoRouteEntryRequest {
	s.VpnConnectionId = &v
	return s
}

func (s *CreateVcoRouteEntryRequest) SetWeight(v int32) *CreateVcoRouteEntryRequest {
	s.Weight = &v
	return s
}

func (s *CreateVcoRouteEntryRequest) Validate() error {
	return dara.Validate(s)
}
