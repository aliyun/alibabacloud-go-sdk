// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpnGatewayAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPropagate(v bool) *ModifyVpnGatewayAttributeRequest
	GetAutoPropagate() *bool
	SetClientToken(v string) *ModifyVpnGatewayAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyVpnGatewayAttributeRequest
	GetDescription() *string
	SetName(v string) *ModifyVpnGatewayAttributeRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyVpnGatewayAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyVpnGatewayAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyVpnGatewayAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyVpnGatewayAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyVpnGatewayAttributeRequest
	GetResourceOwnerId() *int64
	SetVpnGatewayId(v string) *ModifyVpnGatewayAttributeRequest
	GetVpnGatewayId() *string
}

type ModifyVpnGatewayAttributeRequest struct {
	// Specifies whether to automatically advertise BGP routes to the virtual private cloud (VPC). Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AutoPropagate *bool `json:"AutoPropagate,omitempty" xml:"AutoPropagate,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-4266****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The new description of the VPN connection.
	//
	// The description must be 1 to 100 characters in length.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The new name of the VPN gateway.
	//
	// The name must be 2 to 100 characters in length and cannot start with `http://` or `https://`. It must start with a letter and can contain letters, digits, underscores (_), hyphens (-), and periods (.). Other characters are not supported.
	//
	// example:
	//
	// myvpn
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the VPN gateway is created. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VPN gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpn-bp1q8bgx4xnkm2ogj****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s ModifyVpnGatewayAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnGatewayAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpnGatewayAttributeRequest) GetAutoPropagate() *bool {
	return s.AutoPropagate
}

func (s *ModifyVpnGatewayAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyVpnGatewayAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyVpnGatewayAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyVpnGatewayAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyVpnGatewayAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyVpnGatewayAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyVpnGatewayAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyVpnGatewayAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyVpnGatewayAttributeRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *ModifyVpnGatewayAttributeRequest) SetAutoPropagate(v bool) *ModifyVpnGatewayAttributeRequest {
	s.AutoPropagate = &v
	return s
}

func (s *ModifyVpnGatewayAttributeRequest) SetClientToken(v string) *ModifyVpnGatewayAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyVpnGatewayAttributeRequest) SetDescription(v string) *ModifyVpnGatewayAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyVpnGatewayAttributeRequest) SetName(v string) *ModifyVpnGatewayAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyVpnGatewayAttributeRequest) SetOwnerAccount(v string) *ModifyVpnGatewayAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyVpnGatewayAttributeRequest) SetOwnerId(v int64) *ModifyVpnGatewayAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyVpnGatewayAttributeRequest) SetRegionId(v string) *ModifyVpnGatewayAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyVpnGatewayAttributeRequest) SetResourceOwnerAccount(v string) *ModifyVpnGatewayAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyVpnGatewayAttributeRequest) SetResourceOwnerId(v int64) *ModifyVpnGatewayAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyVpnGatewayAttributeRequest) SetVpnGatewayId(v string) *ModifyVpnGatewayAttributeRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *ModifyVpnGatewayAttributeRequest) Validate() error {
	return dara.Validate(s)
}
