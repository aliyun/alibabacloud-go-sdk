// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeResourceUsageRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeResourceUsageRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeResourceUsageRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DescribeResourceUsageRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeResourceUsageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeResourceUsageRequest
	GetResourceOwnerId() *int64
}

type DescribeResourceUsageRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeResourceUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeResourceUsageRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeResourceUsageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeResourceUsageRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeResourceUsageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeResourceUsageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeResourceUsageRequest) SetDBInstanceId(v string) *DescribeResourceUsageRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeResourceUsageRequest) SetOwnerAccount(v string) *DescribeResourceUsageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeResourceUsageRequest) SetOwnerId(v int64) *DescribeResourceUsageRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeResourceUsageRequest) SetResourceGroupId(v string) *DescribeResourceUsageRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeResourceUsageRequest) SetResourceOwnerAccount(v string) *DescribeResourceUsageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeResourceUsageRequest) SetResourceOwnerId(v int64) *DescribeResourceUsageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeResourceUsageRequest) Validate() error {
	return dara.Validate(s)
}
