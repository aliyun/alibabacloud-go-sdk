// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEipGatewayInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeEipGatewayInfoRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeEipGatewayInfoRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeEipGatewayInfoRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeEipGatewayInfoRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeEipGatewayInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeEipGatewayInfoRequest
	GetResourceOwnerId() *int64
}

type DescribeEipGatewayInfoRequest struct {
	// The ID of the secondary ENI that is associated with the EIP.
	//
	// This parameter is required.
	//
	// example:
	//
	// eni-bp1d66qjxb3qoin3****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region to which the EIP that you want to query belongs. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeEipGatewayInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipGatewayInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeEipGatewayInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeEipGatewayInfoRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeEipGatewayInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeEipGatewayInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEipGatewayInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeEipGatewayInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeEipGatewayInfoRequest) SetInstanceId(v string) *DescribeEipGatewayInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeEipGatewayInfoRequest) SetOwnerAccount(v string) *DescribeEipGatewayInfoRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeEipGatewayInfoRequest) SetOwnerId(v int64) *DescribeEipGatewayInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEipGatewayInfoRequest) SetRegionId(v string) *DescribeEipGatewayInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEipGatewayInfoRequest) SetResourceOwnerAccount(v string) *DescribeEipGatewayInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeEipGatewayInfoRequest) SetResourceOwnerId(v int64) *DescribeEipGatewayInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeEipGatewayInfoRequest) Validate() error {
	return dara.Validate(s)
}
