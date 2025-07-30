// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceNetInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDBInstanceNetInfoRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeDBInstanceNetInfoRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstanceNetInfoRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBInstanceNetInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceNetInfoRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeDBInstanceNetInfoRequest
	GetSecurityToken() *string
}

type DescribeDBInstanceNetInfoRequest struct {
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

func (s DescribeDBInstanceNetInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDBInstanceNetInfoRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstanceNetInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstanceNetInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstanceNetInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceNetInfoRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDBInstanceNetInfoRequest) SetInstanceId(v string) *DescribeDBInstanceNetInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) SetOwnerAccount(v string) *DescribeDBInstanceNetInfoRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) SetOwnerId(v int64) *DescribeDBInstanceNetInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceNetInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceNetInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) SetSecurityToken(v string) *DescribeDBInstanceNetInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) Validate() error {
	return dara.Validate(s)
}
