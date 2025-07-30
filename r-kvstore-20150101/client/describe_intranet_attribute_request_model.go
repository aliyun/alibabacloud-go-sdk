// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIntranetAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeIntranetAttributeRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeIntranetAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeIntranetAttributeRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DescribeIntranetAttributeRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeIntranetAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeIntranetAttributeRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeIntranetAttributeRequest
	GetSecurityToken() *string
}

type DescribeIntranetAttributeRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeIntranetAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIntranetAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeIntranetAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeIntranetAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeIntranetAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeIntranetAttributeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeIntranetAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeIntranetAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeIntranetAttributeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeIntranetAttributeRequest) SetInstanceId(v string) *DescribeIntranetAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeIntranetAttributeRequest) SetOwnerAccount(v string) *DescribeIntranetAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeIntranetAttributeRequest) SetOwnerId(v int64) *DescribeIntranetAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeIntranetAttributeRequest) SetResourceGroupId(v string) *DescribeIntranetAttributeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeIntranetAttributeRequest) SetResourceOwnerAccount(v string) *DescribeIntranetAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeIntranetAttributeRequest) SetResourceOwnerId(v int64) *DescribeIntranetAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeIntranetAttributeRequest) SetSecurityToken(v string) *DescribeIntranetAttributeRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeIntranetAttributeRequest) Validate() error {
	return dara.Validate(s)
}
