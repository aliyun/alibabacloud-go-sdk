// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoleTagStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeRoleTagStatusRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeRoleTagStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeRoleTagStatusRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeRoleTagStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRoleTagStatusRequest
	GetResourceOwnerId() *int64
}

type DescribeRoleTagStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dds-wz9760547de1****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRoleTagStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoleTagStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeRoleTagStatusRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeRoleTagStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeRoleTagStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRoleTagStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRoleTagStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRoleTagStatusRequest) SetDBInstanceId(v string) *DescribeRoleTagStatusRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeRoleTagStatusRequest) SetOwnerAccount(v string) *DescribeRoleTagStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRoleTagStatusRequest) SetOwnerId(v int64) *DescribeRoleTagStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRoleTagStatusRequest) SetResourceOwnerAccount(v string) *DescribeRoleTagStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRoleTagStatusRequest) SetResourceOwnerId(v int64) *DescribeRoleTagStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRoleTagStatusRequest) Validate() error {
	return dara.Validate(s)
}
