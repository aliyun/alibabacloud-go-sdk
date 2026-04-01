// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLCollectorRetentionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeSQLCollectorRetentionRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeSQLCollectorRetentionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSQLCollectorRetentionRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DescribeSQLCollectorRetentionRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeSQLCollectorRetentionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSQLCollectorRetentionRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeSQLCollectorRetentionRequest
	GetSecurityToken() *string
}

type DescribeSQLCollectorRetentionRequest struct {
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
	// The resource group ID. You can call the DescribeDBInstanceAttribute operation to query the resource group ID.
	//
	// example:
	//
	// rg-acfmyxxxx
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeSQLCollectorRetentionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLCollectorRetentionRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLCollectorRetentionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSQLCollectorRetentionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSQLCollectorRetentionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSQLCollectorRetentionRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSQLCollectorRetentionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSQLCollectorRetentionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSQLCollectorRetentionRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeSQLCollectorRetentionRequest) SetDBInstanceId(v string) *DescribeSQLCollectorRetentionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLCollectorRetentionRequest) SetOwnerAccount(v string) *DescribeSQLCollectorRetentionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSQLCollectorRetentionRequest) SetOwnerId(v int64) *DescribeSQLCollectorRetentionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSQLCollectorRetentionRequest) SetResourceGroupId(v string) *DescribeSQLCollectorRetentionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSQLCollectorRetentionRequest) SetResourceOwnerAccount(v string) *DescribeSQLCollectorRetentionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSQLCollectorRetentionRequest) SetResourceOwnerId(v int64) *DescribeSQLCollectorRetentionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSQLCollectorRetentionRequest) SetSecurityToken(v string) *DescribeSQLCollectorRetentionRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeSQLCollectorRetentionRequest) Validate() error {
	return dara.Validate(s)
}
