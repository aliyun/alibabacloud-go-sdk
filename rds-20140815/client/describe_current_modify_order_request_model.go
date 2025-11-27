// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCurrentModifyOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeCurrentModifyOrderRequest
	GetClientToken() *string
	SetDbInstanceId(v string) *DescribeCurrentModifyOrderRequest
	GetDbInstanceId() *string
	SetOwnerId(v int64) *DescribeCurrentModifyOrderRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeCurrentModifyOrderRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeCurrentModifyOrderRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeCurrentModifyOrderRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCurrentModifyOrderRequest
	GetResourceOwnerId() *int64
}

type DescribeCurrentModifyOrderRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the generated token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp1u775467ggm7j9j
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeCurrentModifyOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCurrentModifyOrderRequest) GoString() string {
	return s.String()
}

func (s *DescribeCurrentModifyOrderRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeCurrentModifyOrderRequest) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *DescribeCurrentModifyOrderRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCurrentModifyOrderRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCurrentModifyOrderRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCurrentModifyOrderRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCurrentModifyOrderRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCurrentModifyOrderRequest) SetClientToken(v string) *DescribeCurrentModifyOrderRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeCurrentModifyOrderRequest) SetDbInstanceId(v string) *DescribeCurrentModifyOrderRequest {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeCurrentModifyOrderRequest) SetOwnerId(v int64) *DescribeCurrentModifyOrderRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCurrentModifyOrderRequest) SetRegionId(v string) *DescribeCurrentModifyOrderRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCurrentModifyOrderRequest) SetResourceGroupId(v string) *DescribeCurrentModifyOrderRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCurrentModifyOrderRequest) SetResourceOwnerAccount(v string) *DescribeCurrentModifyOrderRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCurrentModifyOrderRequest) SetResourceOwnerId(v int64) *DescribeCurrentModifyOrderRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCurrentModifyOrderRequest) Validate() error {
	return dara.Validate(s)
}
