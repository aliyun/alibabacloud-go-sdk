// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomerGatewayAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthKey(v string) *ModifyCustomerGatewayAttributeRequest
	GetAuthKey() *string
	SetClientToken(v string) *ModifyCustomerGatewayAttributeRequest
	GetClientToken() *string
	SetCustomerGatewayId(v string) *ModifyCustomerGatewayAttributeRequest
	GetCustomerGatewayId() *string
	SetDescription(v string) *ModifyCustomerGatewayAttributeRequest
	GetDescription() *string
	SetName(v string) *ModifyCustomerGatewayAttributeRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyCustomerGatewayAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyCustomerGatewayAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyCustomerGatewayAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyCustomerGatewayAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyCustomerGatewayAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyCustomerGatewayAttributeRequest struct {
	// The authentication key of the BGP routing protocol for the gateway device in the data center.
	//
	// The key must be 1 to 64 characters in length. It can contain only ASCII characters and cannot contain spaces or question marks (?).
	//
	// example:
	//
	// AuthKey****
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the customer gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// cgw-bp1pvpl9r9adju6l5****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	// The description of the customer gateway.
	//
	// The description must be 1 to 100 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the customer gateway.
	//
	// The name must be 1 to 100 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// nametest
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the customer gateway is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyCustomerGatewayAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomerGatewayAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyCustomerGatewayAttributeRequest) GetAuthKey() *string {
	return s.AuthKey
}

func (s *ModifyCustomerGatewayAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyCustomerGatewayAttributeRequest) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *ModifyCustomerGatewayAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyCustomerGatewayAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyCustomerGatewayAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyCustomerGatewayAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCustomerGatewayAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCustomerGatewayAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyCustomerGatewayAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyCustomerGatewayAttributeRequest) SetAuthKey(v string) *ModifyCustomerGatewayAttributeRequest {
	s.AuthKey = &v
	return s
}

func (s *ModifyCustomerGatewayAttributeRequest) SetClientToken(v string) *ModifyCustomerGatewayAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyCustomerGatewayAttributeRequest) SetCustomerGatewayId(v string) *ModifyCustomerGatewayAttributeRequest {
	s.CustomerGatewayId = &v
	return s
}

func (s *ModifyCustomerGatewayAttributeRequest) SetDescription(v string) *ModifyCustomerGatewayAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyCustomerGatewayAttributeRequest) SetName(v string) *ModifyCustomerGatewayAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyCustomerGatewayAttributeRequest) SetOwnerAccount(v string) *ModifyCustomerGatewayAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyCustomerGatewayAttributeRequest) SetOwnerId(v int64) *ModifyCustomerGatewayAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCustomerGatewayAttributeRequest) SetRegionId(v string) *ModifyCustomerGatewayAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCustomerGatewayAttributeRequest) SetResourceOwnerAccount(v string) *ModifyCustomerGatewayAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyCustomerGatewayAttributeRequest) SetResourceOwnerId(v int64) *ModifyCustomerGatewayAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyCustomerGatewayAttributeRequest) Validate() error {
	return dara.Validate(s)
}
