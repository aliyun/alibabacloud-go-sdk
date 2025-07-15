// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcGatewayEndpointAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointId(v string) *GetVpcGatewayEndpointAttributeRequest
	GetEndpointId() *string
	SetOwnerAccount(v string) *GetVpcGatewayEndpointAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetVpcGatewayEndpointAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetVpcGatewayEndpointAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetVpcGatewayEndpointAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetVpcGatewayEndpointAttributeRequest
	GetResourceOwnerId() *int64
}

type GetVpcGatewayEndpointAttributeRequest struct {
	// The ID of the gateway endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpce-bp1w1dmdqjpwul0v3****
	EndpointId   *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the gateway endpoint.
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
}

func (s GetVpcGatewayEndpointAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVpcGatewayEndpointAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetVpcGatewayEndpointAttributeRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *GetVpcGatewayEndpointAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetVpcGatewayEndpointAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetVpcGatewayEndpointAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVpcGatewayEndpointAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetVpcGatewayEndpointAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetVpcGatewayEndpointAttributeRequest) SetEndpointId(v string) *GetVpcGatewayEndpointAttributeRequest {
	s.EndpointId = &v
	return s
}

func (s *GetVpcGatewayEndpointAttributeRequest) SetOwnerAccount(v string) *GetVpcGatewayEndpointAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetVpcGatewayEndpointAttributeRequest) SetOwnerId(v int64) *GetVpcGatewayEndpointAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *GetVpcGatewayEndpointAttributeRequest) SetRegionId(v string) *GetVpcGatewayEndpointAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *GetVpcGatewayEndpointAttributeRequest) SetResourceOwnerAccount(v string) *GetVpcGatewayEndpointAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetVpcGatewayEndpointAttributeRequest) SetResourceOwnerId(v int64) *GetVpcGatewayEndpointAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetVpcGatewayEndpointAttributeRequest) Validate() error {
	return dara.Validate(s)
}
