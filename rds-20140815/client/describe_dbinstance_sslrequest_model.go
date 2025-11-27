// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceSSLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceSSLRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeDBInstanceSSLRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstanceSSLRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBInstanceSSLRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceSSLRequest
	GetResourceOwnerId() *int64
}

type DescribeDBInstanceSSLRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp162dfr55g47****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBInstanceSSLRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSSLRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSSLRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceSSLRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstanceSSLRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstanceSSLRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstanceSSLRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceSSLRequest) SetDBInstanceId(v string) *DescribeDBInstanceSSLRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceSSLRequest) SetOwnerAccount(v string) *DescribeDBInstanceSSLRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceSSLRequest) SetOwnerId(v int64) *DescribeDBInstanceSSLRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceSSLRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceSSLRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceSSLRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceSSLRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceSSLRequest) Validate() error {
	return dara.Validate(s)
}
