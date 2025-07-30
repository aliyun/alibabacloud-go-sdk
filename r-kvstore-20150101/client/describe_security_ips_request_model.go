// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeSecurityIpsRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeSecurityIpsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSecurityIpsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeSecurityIpsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSecurityIpsRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeSecurityIpsRequest
	GetSecurityToken() *string
}

type DescribeSecurityIpsRequest struct {
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

func (s DescribeSecurityIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSecurityIpsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSecurityIpsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSecurityIpsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSecurityIpsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSecurityIpsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeSecurityIpsRequest) SetInstanceId(v string) *DescribeSecurityIpsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSecurityIpsRequest) SetOwnerAccount(v string) *DescribeSecurityIpsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSecurityIpsRequest) SetOwnerId(v int64) *DescribeSecurityIpsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSecurityIpsRequest) SetResourceOwnerAccount(v string) *DescribeSecurityIpsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSecurityIpsRequest) SetResourceOwnerId(v int64) *DescribeSecurityIpsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSecurityIpsRequest) SetSecurityToken(v string) *DescribeSecurityIpsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeSecurityIpsRequest) Validate() error {
	return dara.Validate(s)
}
