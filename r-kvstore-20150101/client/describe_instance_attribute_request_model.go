// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeInstanceAttributeRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeInstanceAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstanceAttributeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeInstanceAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceAttributeRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeInstanceAttributeRequest
	GetSecurityToken() *string
}

type DescribeInstanceAttributeRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstanceAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceAttributeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeInstanceAttributeRequest) SetInstanceId(v string) *DescribeInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceAttributeRequest) SetOwnerAccount(v string) *DescribeInstanceAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceAttributeRequest) SetOwnerId(v int64) *DescribeInstanceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceAttributeRequest) SetResourceOwnerAccount(v string) *DescribeInstanceAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceAttributeRequest) SetResourceOwnerId(v int64) *DescribeInstanceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceAttributeRequest) SetSecurityToken(v string) *DescribeInstanceAttributeRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
