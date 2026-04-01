// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateMigrateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *TerminateMigrateTaskRequest
	GetDBInstanceId() *string
	SetMigrateTaskId(v string) *TerminateMigrateTaskRequest
	GetMigrateTaskId() *string
	SetOwnerId(v int64) *TerminateMigrateTaskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *TerminateMigrateTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TerminateMigrateTaskRequest
	GetResourceOwnerId() *int64
}

type TerminateMigrateTaskRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp159vfbu******
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The migration task ID. You can call the DescribeMigrateTasks operation to query the migration task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 56254****
	MigrateTaskId        *string `json:"MigrateTaskId,omitempty" xml:"MigrateTaskId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s TerminateMigrateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s TerminateMigrateTaskRequest) GoString() string {
	return s.String()
}

func (s *TerminateMigrateTaskRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *TerminateMigrateTaskRequest) GetMigrateTaskId() *string {
	return s.MigrateTaskId
}

func (s *TerminateMigrateTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TerminateMigrateTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TerminateMigrateTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TerminateMigrateTaskRequest) SetDBInstanceId(v string) *TerminateMigrateTaskRequest {
	s.DBInstanceId = &v
	return s
}

func (s *TerminateMigrateTaskRequest) SetMigrateTaskId(v string) *TerminateMigrateTaskRequest {
	s.MigrateTaskId = &v
	return s
}

func (s *TerminateMigrateTaskRequest) SetOwnerId(v int64) *TerminateMigrateTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *TerminateMigrateTaskRequest) SetResourceOwnerAccount(v string) *TerminateMigrateTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TerminateMigrateTaskRequest) SetResourceOwnerId(v int64) *TerminateMigrateTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TerminateMigrateTaskRequest) Validate() error {
	return dara.Validate(s)
}
