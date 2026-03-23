// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTempDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v int64) *CreateTempDBInstanceRequest
	GetBackupId() *int64
	SetDBInstanceId(v string) *CreateTempDBInstanceRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *CreateTempDBInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTempDBInstanceRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *CreateTempDBInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateTempDBInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTempDBInstanceRequest
	GetResourceOwnerId() *int64
	SetRestoreTime(v string) *CreateTempDBInstanceRequest
	GetRestoreTime() *string
}

type CreateTempDBInstanceRequest struct {
	BackupId *int64 `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// This parameter is required.
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RestoreTime          *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
}

func (s CreateTempDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTempDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateTempDBInstanceRequest) GetBackupId() *int64 {
	return s.BackupId
}

func (s *CreateTempDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateTempDBInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTempDBInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTempDBInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateTempDBInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTempDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTempDBInstanceRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *CreateTempDBInstanceRequest) SetBackupId(v int64) *CreateTempDBInstanceRequest {
	s.BackupId = &v
	return s
}

func (s *CreateTempDBInstanceRequest) SetDBInstanceId(v string) *CreateTempDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateTempDBInstanceRequest) SetOwnerAccount(v string) *CreateTempDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTempDBInstanceRequest) SetOwnerId(v int64) *CreateTempDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTempDBInstanceRequest) SetResourceGroupId(v string) *CreateTempDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTempDBInstanceRequest) SetResourceOwnerAccount(v string) *CreateTempDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTempDBInstanceRequest) SetResourceOwnerId(v int64) *CreateTempDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTempDBInstanceRequest) SetRestoreTime(v string) *CreateTempDBInstanceRequest {
	s.RestoreTime = &v
	return s
}

func (s *CreateTempDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
