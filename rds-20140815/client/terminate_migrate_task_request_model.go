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
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
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
