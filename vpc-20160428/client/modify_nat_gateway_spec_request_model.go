// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatGatewaySpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyNatGatewaySpecRequest
	GetAutoPay() *bool
	SetClientToken(v string) *ModifyNatGatewaySpecRequest
	GetClientToken() *string
	SetNatGatewayId(v string) *ModifyNatGatewaySpecRequest
	GetNatGatewayId() *string
	SetOwnerAccount(v string) *ModifyNatGatewaySpecRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyNatGatewaySpecRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyNatGatewaySpecRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyNatGatewaySpecRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyNatGatewaySpecRequest
	GetResourceOwnerId() *int64
	SetSpec(v string) *ModifyNatGatewaySpecRequest
	GetSpec() *string
}

type ModifyNatGatewaySpecRequest struct {
	// Specifies whether to automatically complete the payment.
	//
	// 	- **true**: enables automatic payment. Payments are automatically completed.
	//
	// 	- **false*	- (default): disables automatic payment. If you select this option, you must go to the Order Center to complete the payment after an order is generated.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the Internet NAT gateway that you want to upgrade.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-bp1uewa15k4iy5770****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the Internet NAT gateway is deployed.
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
	// The size of the Internet NAT gateway. Valid values:
	//
	// 	- **Small**: small
	//
	// 	- **Middle**: medium
	//
	// 	- **Large**: large
	//
	// This parameter is required.
	//
	// example:
	//
	// Middle
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s ModifyNatGatewaySpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatGatewaySpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyNatGatewaySpecRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyNatGatewaySpecRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyNatGatewaySpecRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *ModifyNatGatewaySpecRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyNatGatewaySpecRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyNatGatewaySpecRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyNatGatewaySpecRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyNatGatewaySpecRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyNatGatewaySpecRequest) GetSpec() *string {
	return s.Spec
}

func (s *ModifyNatGatewaySpecRequest) SetAutoPay(v bool) *ModifyNatGatewaySpecRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyNatGatewaySpecRequest) SetClientToken(v string) *ModifyNatGatewaySpecRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyNatGatewaySpecRequest) SetNatGatewayId(v string) *ModifyNatGatewaySpecRequest {
	s.NatGatewayId = &v
	return s
}

func (s *ModifyNatGatewaySpecRequest) SetOwnerAccount(v string) *ModifyNatGatewaySpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyNatGatewaySpecRequest) SetOwnerId(v int64) *ModifyNatGatewaySpecRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyNatGatewaySpecRequest) SetRegionId(v string) *ModifyNatGatewaySpecRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyNatGatewaySpecRequest) SetResourceOwnerAccount(v string) *ModifyNatGatewaySpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyNatGatewaySpecRequest) SetResourceOwnerId(v int64) *ModifyNatGatewaySpecRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyNatGatewaySpecRequest) SetSpec(v string) *ModifyNatGatewaySpecRequest {
	s.Spec = &v
	return s
}

func (s *ModifyNatGatewaySpecRequest) Validate() error {
	return dara.Validate(s)
}
