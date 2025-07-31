// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceTDEInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceTDEInfoRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeDBInstanceTDEInfoRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstanceTDEInfoRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBInstanceTDEInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceTDEInfoRequest
	GetResourceOwnerId() *int64
}

type DescribeDBInstanceTDEInfoRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bpxxxxxxxx
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBInstanceTDEInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceTDEInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTDEInfoRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceTDEInfoRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstanceTDEInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstanceTDEInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstanceTDEInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceTDEInfoRequest) SetDBInstanceId(v string) *DescribeDBInstanceTDEInfoRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceTDEInfoRequest) SetOwnerAccount(v string) *DescribeDBInstanceTDEInfoRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceTDEInfoRequest) SetOwnerId(v int64) *DescribeDBInstanceTDEInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceTDEInfoRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceTDEInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceTDEInfoRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceTDEInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceTDEInfoRequest) Validate() error {
	return dara.Validate(s)
}
