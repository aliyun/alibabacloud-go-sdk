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
	// The ID of the backup set based on which you want to restore databases of the source instance. When you replicate databases by backup set, you can call the DescribeBackups operation to obtain the ID of the backup set.
	//
	// >  You must specify one of the **BackupId*	- and **RestoreTime*	- parameters.
	//
	// example:
	//
	// 106523874****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The source instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The names of the databases that you want to copy. Format: `Source database name 1,Source database name 2`.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"test1":"newtest1","test2":"newtest2"}
	DbNames         *string `json:"DbNames,omitempty" xml:"DbNames,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The point in time when the system replicates databases. You can select a point in time within the backup retention period. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// > You must specify one of the **BackupId*	- and **RestoreTime*	- parameters.
	//
	// example:
	//
	// 2011-06-11T16:00:00Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// Specifies whether to copy users and permissions.
	//
	// 	- **YES**: copies users and permissions. If the destination instance has a user whose name is the same as a user in the source instance, the permissions of the user in the source instance will also be granted to the user in the destination instance after you copy user permissions.
	//
	// 	- **NO**: does not copy users and permissions.
	//
	// Default value: **NO**.
	//
	// example:
	//
	// NO
	SyncUserPrivilege *string `json:"SyncUserPrivilege,omitempty" xml:"SyncUserPrivilege,omitempty"`
	// The destination instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-ut5ajk3xxxxxxx
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
