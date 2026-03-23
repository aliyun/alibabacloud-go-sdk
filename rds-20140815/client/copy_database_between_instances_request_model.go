// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyDatabaseBetweenInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *CopyDatabaseBetweenInstancesRequest
	GetBackupId() *string
	SetDBInstanceId(v string) *CopyDatabaseBetweenInstancesRequest
	GetDBInstanceId() *string
	SetDbNames(v string) *CopyDatabaseBetweenInstancesRequest
	GetDbNames() *string
	SetResourceOwnerId(v int64) *CopyDatabaseBetweenInstancesRequest
	GetResourceOwnerId() *int64
	SetRestoreTime(v string) *CopyDatabaseBetweenInstancesRequest
	GetRestoreTime() *string
	SetSyncUserPrivilege(v string) *CopyDatabaseBetweenInstancesRequest
	GetSyncUserPrivilege() *string
	SetTargetDBInstanceId(v string) *CopyDatabaseBetweenInstancesRequest
	GetTargetDBInstanceId() *string
}

type CopyDatabaseBetweenInstancesRequest struct {
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	DbNames           *string `json:"DbNames,omitempty" xml:"DbNames,omitempty"`
	ResourceOwnerId   *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RestoreTime       *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	SyncUserPrivilege *string `json:"SyncUserPrivilege,omitempty" xml:"SyncUserPrivilege,omitempty"`
	// This parameter is required.
	TargetDBInstanceId *string `json:"TargetDBInstanceId,omitempty" xml:"TargetDBInstanceId,omitempty"`
}

func (s CopyDatabaseBetweenInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyDatabaseBetweenInstancesRequest) GoString() string {
	return s.String()
}

func (s *CopyDatabaseBetweenInstancesRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *CopyDatabaseBetweenInstancesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CopyDatabaseBetweenInstancesRequest) GetDbNames() *string {
	return s.DbNames
}

func (s *CopyDatabaseBetweenInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CopyDatabaseBetweenInstancesRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *CopyDatabaseBetweenInstancesRequest) GetSyncUserPrivilege() *string {
	return s.SyncUserPrivilege
}

func (s *CopyDatabaseBetweenInstancesRequest) GetTargetDBInstanceId() *string {
	return s.TargetDBInstanceId
}

func (s *CopyDatabaseBetweenInstancesRequest) SetBackupId(v string) *CopyDatabaseBetweenInstancesRequest {
	s.BackupId = &v
	return s
}

func (s *CopyDatabaseBetweenInstancesRequest) SetDBInstanceId(v string) *CopyDatabaseBetweenInstancesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CopyDatabaseBetweenInstancesRequest) SetDbNames(v string) *CopyDatabaseBetweenInstancesRequest {
	s.DbNames = &v
	return s
}

func (s *CopyDatabaseBetweenInstancesRequest) SetResourceOwnerId(v int64) *CopyDatabaseBetweenInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CopyDatabaseBetweenInstancesRequest) SetRestoreTime(v string) *CopyDatabaseBetweenInstancesRequest {
	s.RestoreTime = &v
	return s
}

func (s *CopyDatabaseBetweenInstancesRequest) SetSyncUserPrivilege(v string) *CopyDatabaseBetweenInstancesRequest {
	s.SyncUserPrivilege = &v
	return s
}

func (s *CopyDatabaseBetweenInstancesRequest) SetTargetDBInstanceId(v string) *CopyDatabaseBetweenInstancesRequest {
	s.TargetDBInstanceId = &v
	return s
}

func (s *CopyDatabaseBetweenInstancesRequest) Validate() error {
	return dara.Validate(s)
}
