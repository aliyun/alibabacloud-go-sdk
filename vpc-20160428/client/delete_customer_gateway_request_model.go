// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomerGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteCustomerGatewayRequest
	GetClientToken() *string
	SetCustomerGatewayId(v string) *DeleteCustomerGatewayRequest
	GetCustomerGatewayId() *string
	SetOwnerAccount(v string) *DeleteCustomerGatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteCustomerGatewayRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteCustomerGatewayRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteCustomerGatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteCustomerGatewayRequest
	GetResourceOwnerId() *int64
}

type DeleteCustomerGatewayRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the customer gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// cgw-bp1pvpl9r9adju6l5****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the customer gateway. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
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

func (s DeleteCustomerGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomerGatewayRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomerGatewayRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteCustomerGatewayRequest) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *DeleteCustomerGatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteCustomerGatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCustomerGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCustomerGatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteCustomerGatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteCustomerGatewayRequest) SetClientToken(v string) *DeleteCustomerGatewayRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCustomerGatewayRequest) SetCustomerGatewayId(v string) *DeleteCustomerGatewayRequest {
	s.CustomerGatewayId = &v
	return s
}

func (s *DeleteCustomerGatewayRequest) SetOwnerAccount(v string) *DeleteCustomerGatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCustomerGatewayRequest) SetOwnerId(v int64) *DeleteCustomerGatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCustomerGatewayRequest) SetRegionId(v string) *DeleteCustomerGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCustomerGatewayRequest) SetResourceOwnerAccount(v string) *DeleteCustomerGatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCustomerGatewayRequest) SetResourceOwnerId(v int64) *DeleteCustomerGatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteCustomerGatewayRequest) Validate() error {
	return dara.Validate(s)
}
