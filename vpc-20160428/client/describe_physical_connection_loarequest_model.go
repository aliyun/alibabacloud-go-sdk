// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhysicalConnectionLOARequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribePhysicalConnectionLOARequest
	GetClientToken() *string
	SetInstanceId(v string) *DescribePhysicalConnectionLOARequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribePhysicalConnectionLOARequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribePhysicalConnectionLOARequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribePhysicalConnectionLOARequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribePhysicalConnectionLOARequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePhysicalConnectionLOARequest
	GetResourceOwnerId() *int64
}

type DescribePhysicalConnectionLOARequest struct {
	// A client token that ensures the idempotence of the request.
	//
	// Your client generates this value, which must be unique for each request and have a maximum length of 64 ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. Each API request has a unique **RequestId**.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the physical connection instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1ca4wca27ex****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the physical connection is located. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to get a list of available region IDs.
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

func (s DescribePhysicalConnectionLOARequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePhysicalConnectionLOARequest) GoString() string {
	return s.String()
}

func (s *DescribePhysicalConnectionLOARequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribePhysicalConnectionLOARequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePhysicalConnectionLOARequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribePhysicalConnectionLOARequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePhysicalConnectionLOARequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePhysicalConnectionLOARequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePhysicalConnectionLOARequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePhysicalConnectionLOARequest) SetClientToken(v string) *DescribePhysicalConnectionLOARequest {
	s.ClientToken = &v
	return s
}

func (s *DescribePhysicalConnectionLOARequest) SetInstanceId(v string) *DescribePhysicalConnectionLOARequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePhysicalConnectionLOARequest) SetOwnerAccount(v string) *DescribePhysicalConnectionLOARequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePhysicalConnectionLOARequest) SetOwnerId(v int64) *DescribePhysicalConnectionLOARequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhysicalConnectionLOARequest) SetRegionId(v string) *DescribePhysicalConnectionLOARequest {
	s.RegionId = &v
	return s
}

func (s *DescribePhysicalConnectionLOARequest) SetResourceOwnerAccount(v string) *DescribePhysicalConnectionLOARequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhysicalConnectionLOARequest) SetResourceOwnerId(v int64) *DescribePhysicalConnectionLOARequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePhysicalConnectionLOARequest) Validate() error {
	return dara.Validate(s)
}
