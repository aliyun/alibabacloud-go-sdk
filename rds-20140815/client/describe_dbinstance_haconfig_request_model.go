// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceHAConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceHAConfigRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeDBInstanceHAConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstanceHAConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBInstanceHAConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceHAConfigRequest
	GetResourceOwnerId() *int64
}

type DescribeDBInstanceHAConfigRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxx
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBInstanceHAConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceHAConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceHAConfigRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceHAConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstanceHAConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstanceHAConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstanceHAConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceHAConfigRequest) SetDBInstanceId(v string) *DescribeDBInstanceHAConfigRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceHAConfigRequest) SetOwnerAccount(v string) *DescribeDBInstanceHAConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceHAConfigRequest) SetOwnerId(v int64) *DescribeDBInstanceHAConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceHAConfigRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceHAConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceHAConfigRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceHAConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceHAConfigRequest) Validate() error {
	return dara.Validate(s)
}
