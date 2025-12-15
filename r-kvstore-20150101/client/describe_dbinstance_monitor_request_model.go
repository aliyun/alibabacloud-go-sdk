// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceMonitorRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeDBInstanceMonitorRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstanceMonitorRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBInstanceMonitorRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceMonitorRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeDBInstanceMonitorRequest
	GetSecurityToken() *string
}

type DescribeDBInstanceMonitorRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1xp9esa45nll****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDBInstanceMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceMonitorRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceMonitorRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceMonitorRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstanceMonitorRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstanceMonitorRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstanceMonitorRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceMonitorRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDBInstanceMonitorRequest) SetDBInstanceId(v string) *DescribeDBInstanceMonitorRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceMonitorRequest) SetOwnerAccount(v string) *DescribeDBInstanceMonitorRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceMonitorRequest) SetOwnerId(v int64) *DescribeDBInstanceMonitorRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceMonitorRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceMonitorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceMonitorRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceMonitorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceMonitorRequest) SetSecurityToken(v string) *DescribeDBInstanceMonitorRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDBInstanceMonitorRequest) Validate() error {
	return dara.Validate(s)
}
