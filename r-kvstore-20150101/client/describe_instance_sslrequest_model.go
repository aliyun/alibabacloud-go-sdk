// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSSLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeInstanceSSLRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeInstanceSSLRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstanceSSLRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeInstanceSSLRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceSSLRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeInstanceSSLRequest
	GetSecurityToken() *string
}

type DescribeInstanceSSLRequest struct {
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

func (s DescribeInstanceSSLRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSSLRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSSLRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceSSLRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstanceSSLRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceSSLRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceSSLRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceSSLRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeInstanceSSLRequest) SetInstanceId(v string) *DescribeInstanceSSLRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceSSLRequest) SetOwnerAccount(v string) *DescribeInstanceSSLRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceSSLRequest) SetOwnerId(v int64) *DescribeInstanceSSLRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceSSLRequest) SetResourceOwnerAccount(v string) *DescribeInstanceSSLRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceSSLRequest) SetResourceOwnerId(v int64) *DescribeInstanceSSLRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceSSLRequest) SetSecurityToken(v string) *DescribeInstanceSSLRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInstanceSSLRequest) Validate() error {
	return dara.Validate(s)
}
