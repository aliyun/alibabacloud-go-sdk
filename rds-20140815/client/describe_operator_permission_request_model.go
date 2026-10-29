// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOperatorPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeOperatorPermissionRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeOperatorPermissionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeOperatorPermissionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeOperatorPermissionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeOperatorPermissionRequest
	GetResourceOwnerId() *int64
}

type DescribeOperatorPermissionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rm-bp1*******
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeOperatorPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperatorPermissionRequest) GoString() string {
	return s.String()
}

func (s *DescribeOperatorPermissionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeOperatorPermissionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeOperatorPermissionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeOperatorPermissionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeOperatorPermissionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeOperatorPermissionRequest) SetDBInstanceId(v string) *DescribeOperatorPermissionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeOperatorPermissionRequest) SetOwnerAccount(v string) *DescribeOperatorPermissionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeOperatorPermissionRequest) SetOwnerId(v int64) *DescribeOperatorPermissionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeOperatorPermissionRequest) SetResourceOwnerAccount(v string) *DescribeOperatorPermissionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeOperatorPermissionRequest) SetResourceOwnerId(v int64) *DescribeOperatorPermissionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeOperatorPermissionRequest) Validate() error {
	return dara.Validate(s)
}
