// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceTDERequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceTDERequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeDBInstanceTDERequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstanceTDERequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBInstanceTDERequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceTDERequest
	GetResourceOwnerId() *int64
}

type DescribeDBInstanceTDERequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBInstanceTDERequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceTDERequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTDERequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceTDERequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstanceTDERequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstanceTDERequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstanceTDERequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceTDERequest) SetDBInstanceId(v string) *DescribeDBInstanceTDERequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceTDERequest) SetOwnerAccount(v string) *DescribeDBInstanceTDERequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceTDERequest) SetOwnerId(v int64) *DescribeDBInstanceTDERequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceTDERequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceTDERequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceTDERequest) SetResourceOwnerId(v int64) *DescribeDBInstanceTDERequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceTDERequest) Validate() error {
	return dara.Validate(s)
}
