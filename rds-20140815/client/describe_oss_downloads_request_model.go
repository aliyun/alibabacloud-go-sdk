// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssDownloadsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeOssDownloadsRequest
	GetDBInstanceId() *string
	SetMigrateTaskId(v string) *DescribeOssDownloadsRequest
	GetMigrateTaskId() *string
	SetOwnerId(v int64) *DescribeOssDownloadsRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DescribeOssDownloadsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeOssDownloadsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeOssDownloadsRequest
	GetResourceOwnerId() *int64
}

type DescribeOssDownloadsRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The migration task ID. You can call the DescribeMigrateTasks operation to query the migration task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5625458541
	MigrateTaskId *string `json:"MigrateTaskId,omitempty" xml:"MigrateTaskId,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeOssDownloadsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssDownloadsRequest) GoString() string {
	return s.String()
}

func (s *DescribeOssDownloadsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeOssDownloadsRequest) GetMigrateTaskId() *string {
	return s.MigrateTaskId
}

func (s *DescribeOssDownloadsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeOssDownloadsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeOssDownloadsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeOssDownloadsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeOssDownloadsRequest) SetDBInstanceId(v string) *DescribeOssDownloadsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeOssDownloadsRequest) SetMigrateTaskId(v string) *DescribeOssDownloadsRequest {
	s.MigrateTaskId = &v
	return s
}

func (s *DescribeOssDownloadsRequest) SetOwnerId(v int64) *DescribeOssDownloadsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeOssDownloadsRequest) SetResourceGroupId(v string) *DescribeOssDownloadsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeOssDownloadsRequest) SetResourceOwnerAccount(v string) *DescribeOssDownloadsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeOssDownloadsRequest) SetResourceOwnerId(v int64) *DescribeOssDownloadsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeOssDownloadsRequest) Validate() error {
	return dara.Validate(s)
}
