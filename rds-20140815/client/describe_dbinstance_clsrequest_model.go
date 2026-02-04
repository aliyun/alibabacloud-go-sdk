// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceCLSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceCLSRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeDBInstanceCLSRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstanceCLSRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBInstanceCLSRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceCLSRequest
	GetResourceOwnerId() *int64
}

type DescribeDBInstanceCLSRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rm-t4n8t18o******6d5
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBInstanceCLSRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceCLSRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceCLSRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceCLSRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstanceCLSRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstanceCLSRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstanceCLSRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceCLSRequest) SetDBInstanceId(v string) *DescribeDBInstanceCLSRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceCLSRequest) SetOwnerAccount(v string) *DescribeDBInstanceCLSRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceCLSRequest) SetOwnerId(v int64) *DescribeDBInstanceCLSRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceCLSRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceCLSRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceCLSRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceCLSRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceCLSRequest) Validate() error {
	return dara.Validate(s)
}
