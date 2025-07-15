// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomerGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerGatewayId(v string) *DescribeCustomerGatewayRequest
	GetCustomerGatewayId() *string
	SetOwnerAccount(v string) *DescribeCustomerGatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCustomerGatewayRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeCustomerGatewayRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeCustomerGatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCustomerGatewayRequest
	GetResourceOwnerId() *int64
}

type DescribeCustomerGatewayRequest struct {
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

func (s DescribeCustomerGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerGatewayRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomerGatewayRequest) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *DescribeCustomerGatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCustomerGatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCustomerGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCustomerGatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCustomerGatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCustomerGatewayRequest) SetCustomerGatewayId(v string) *DescribeCustomerGatewayRequest {
	s.CustomerGatewayId = &v
	return s
}

func (s *DescribeCustomerGatewayRequest) SetOwnerAccount(v string) *DescribeCustomerGatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCustomerGatewayRequest) SetOwnerId(v int64) *DescribeCustomerGatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCustomerGatewayRequest) SetRegionId(v string) *DescribeCustomerGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCustomerGatewayRequest) SetResourceOwnerAccount(v string) *DescribeCustomerGatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCustomerGatewayRequest) SetResourceOwnerId(v int64) *DescribeCustomerGatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCustomerGatewayRequest) Validate() error {
	return dara.Validate(s)
}
