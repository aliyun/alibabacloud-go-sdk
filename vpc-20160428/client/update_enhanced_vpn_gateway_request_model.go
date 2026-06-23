// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnhancedVpnGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPropagate(v bool) *UpdateEnhancedVpnGatewayRequest
	GetAutoPropagate() *bool
	SetClientToken(v string) *UpdateEnhancedVpnGatewayRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateEnhancedVpnGatewayRequest
	GetDescription() *string
	SetName(v string) *UpdateEnhancedVpnGatewayRequest
	GetName() *string
	SetOwnerAccount(v string) *UpdateEnhancedVpnGatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateEnhancedVpnGatewayRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateEnhancedVpnGatewayRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpdateEnhancedVpnGatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateEnhancedVpnGatewayRequest
	GetResourceOwnerId() *int64
	SetVpnGatewayId(v string) *UpdateEnhancedVpnGatewayRequest
	GetVpnGatewayId() *string
}

type UpdateEnhancedVpnGatewayRequest struct {
	// Specifies whether to enable automatic route propagation for the enhanced VPN gateway. Valid values:
	//
	// - **true**<br>
	//
	//   The enhanced VPN gateway automatically learns system routes from the system route table of the VPC and propagates routes from the on-premises data center to the system route table of the VPC.<br>
	//
	// - **false**<br>
	//
	//   Automatic route propagation is disabled. Before you disable this feature, make sure that BGP dynamic routing is disabled for all IPsec-VPN connections of the enhanced VPN gateway.<br>
	//
	// example:
	//
	// true
	AutoPropagate *bool `json:"AutoPropagate,omitempty" xml:"AutoPropagate,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate a token from your client to make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify a `ClientToken`, the system automatically uses the `RequestId` of the API request as the `ClientToken`. Each API request has a different `RequestId`.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-4266****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The new description of the enhanced VPN gateway.
	//
	// The description must be 1 to 100 characters in length.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The new name of the enhanced VPN gateway.
	//
	// The name must be 2 to 100 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// myvpn
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the enhanced VPN gateway is located. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to obtain the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the enhanced VPN gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpn-bp1q8bgx4xnkm2ogj****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s UpdateEnhancedVpnGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnhancedVpnGatewayRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnhancedVpnGatewayRequest) GetAutoPropagate() *bool {
	return s.AutoPropagate
}

func (s *UpdateEnhancedVpnGatewayRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateEnhancedVpnGatewayRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateEnhancedVpnGatewayRequest) GetName() *string {
	return s.Name
}

func (s *UpdateEnhancedVpnGatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateEnhancedVpnGatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateEnhancedVpnGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEnhancedVpnGatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateEnhancedVpnGatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateEnhancedVpnGatewayRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *UpdateEnhancedVpnGatewayRequest) SetAutoPropagate(v bool) *UpdateEnhancedVpnGatewayRequest {
	s.AutoPropagate = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayRequest) SetClientToken(v string) *UpdateEnhancedVpnGatewayRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayRequest) SetDescription(v string) *UpdateEnhancedVpnGatewayRequest {
	s.Description = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayRequest) SetName(v string) *UpdateEnhancedVpnGatewayRequest {
	s.Name = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayRequest) SetOwnerAccount(v string) *UpdateEnhancedVpnGatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayRequest) SetOwnerId(v int64) *UpdateEnhancedVpnGatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayRequest) SetRegionId(v string) *UpdateEnhancedVpnGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayRequest) SetResourceOwnerAccount(v string) *UpdateEnhancedVpnGatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayRequest) SetResourceOwnerId(v int64) *UpdateEnhancedVpnGatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayRequest) SetVpnGatewayId(v string) *UpdateEnhancedVpnGatewayRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayRequest) Validate() error {
	return dara.Validate(s)
}
