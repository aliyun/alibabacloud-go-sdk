// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceProxyConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceProxyConfigurationRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DescribeDBInstanceProxyConfigurationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBInstanceProxyConfigurationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceProxyConfigurationRequest
	GetResourceOwnerId() *int64
}

type DescribeDBInstanceProxyConfigurationRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBInstanceProxyConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceProxyConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceProxyConfigurationRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceProxyConfigurationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstanceProxyConfigurationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstanceProxyConfigurationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceProxyConfigurationRequest) SetDBInstanceId(v string) *DescribeDBInstanceProxyConfigurationRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceProxyConfigurationRequest) SetOwnerId(v int64) *DescribeDBInstanceProxyConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceProxyConfigurationRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceProxyConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceProxyConfigurationRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceProxyConfigurationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceProxyConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
