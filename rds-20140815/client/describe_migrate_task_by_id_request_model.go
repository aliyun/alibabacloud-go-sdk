// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMigrateTaskByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeMigrateTaskByIdRequest
	GetDBInstanceId() *string
	SetMigrateTaskId(v string) *DescribeMigrateTaskByIdRequest
	GetMigrateTaskId() *string
	SetOwnerId(v int64) *DescribeMigrateTaskByIdRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DescribeMigrateTaskByIdRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeMigrateTaskByIdRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeMigrateTaskByIdRequest
	GetResourceOwnerId() *int64
}

type DescribeMigrateTaskByIdRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp11e1tzgxxxx4ox
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The migration task ID. You can call the DescribeMigrateTasks operation to query the migration task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 235943
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

func (s DescribeMigrateTaskByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrateTaskByIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeMigrateTaskByIdRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeMigrateTaskByIdRequest) GetMigrateTaskId() *string {
	return s.MigrateTaskId
}

func (s *DescribeMigrateTaskByIdRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeMigrateTaskByIdRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeMigrateTaskByIdRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeMigrateTaskByIdRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeMigrateTaskByIdRequest) SetDBInstanceId(v string) *DescribeMigrateTaskByIdRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeMigrateTaskByIdRequest) SetMigrateTaskId(v string) *DescribeMigrateTaskByIdRequest {
	s.MigrateTaskId = &v
	return s
}

func (s *DescribeMigrateTaskByIdRequest) SetOwnerId(v int64) *DescribeMigrateTaskByIdRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMigrateTaskByIdRequest) SetResourceGroupId(v string) *DescribeMigrateTaskByIdRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeMigrateTaskByIdRequest) SetResourceOwnerAccount(v string) *DescribeMigrateTaskByIdRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMigrateTaskByIdRequest) SetResourceOwnerId(v int64) *DescribeMigrateTaskByIdRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeMigrateTaskByIdRequest) Validate() error {
	return dara.Validate(s)
}
