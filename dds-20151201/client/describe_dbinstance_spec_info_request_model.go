// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceSpecInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceClass(v string) *DescribeDBInstanceSpecInfoRequest
	GetInstanceClass() *string
	SetOwnerAccount(v string) *DescribeDBInstanceSpecInfoRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstanceSpecInfoRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBInstanceSpecInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceSpecInfoRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeDBInstanceSpecInfoRequest
	GetSecurityToken() *string
}

type DescribeDBInstanceSpecInfoRequest struct {
	// The instance type. You can query this parameter by calling the [DescribeDBInstanceAttribute](https://next.api.aliyun.com/api/Dds/2015-12-01/DescribeDBInstanceAttribute) operation.
	//
	// For instance types of different instance categories, see the following topics:
	//
	// - [Specifications of standalone instances](https://help.aliyun.com/document_detail/311407.html)
	//
	// - [Specifications of replica set instances](https://help.aliyun.com/document_detail/311410.html)
	//
	// - [Specifications of sharded cluster instances](https://help.aliyun.com/document_detail/311414.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// mdb.shard.4x.large.d
	InstanceClass        *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDBInstanceSpecInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSpecInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSpecInfoRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeDBInstanceSpecInfoRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstanceSpecInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstanceSpecInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstanceSpecInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceSpecInfoRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDBInstanceSpecInfoRequest) SetInstanceClass(v string) *DescribeDBInstanceSpecInfoRequest {
	s.InstanceClass = &v
	return s
}

func (s *DescribeDBInstanceSpecInfoRequest) SetOwnerAccount(v string) *DescribeDBInstanceSpecInfoRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceSpecInfoRequest) SetOwnerId(v int64) *DescribeDBInstanceSpecInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceSpecInfoRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceSpecInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceSpecInfoRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceSpecInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceSpecInfoRequest) SetSecurityToken(v string) *DescribeDBInstanceSpecInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDBInstanceSpecInfoRequest) Validate() error {
	return dara.Validate(s)
}
