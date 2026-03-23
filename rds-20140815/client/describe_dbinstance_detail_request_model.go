// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeDBInstanceDetailRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DescribeDBInstanceDetailRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DescribeDBInstanceDetailRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DescribeDBInstanceDetailRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeDBInstanceDetailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceDetailRequest
	GetResourceOwnerId() *int64
}

type DescribeDBInstanceDetailRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBInstanceDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDetailRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeDBInstanceDetailRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstanceDetailRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstanceDetailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstanceDetailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceDetailRequest) SetClientToken(v string) *DescribeDBInstanceDetailRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeDBInstanceDetailRequest) SetDBInstanceId(v string) *DescribeDBInstanceDetailRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceDetailRequest) SetOwnerId(v int64) *DescribeDBInstanceDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceDetailRequest) SetResourceGroupId(v string) *DescribeDBInstanceDetailRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstanceDetailRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceDetailRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceDetailRequest) Validate() error {
	return dara.Validate(s)
}
