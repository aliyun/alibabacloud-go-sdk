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
	// The ID of the backup set of the source instance. When you copy a database based on a backup set, you can call the DescribeBackups operation to query the backup set ID.
	//
	// >When you call this operation, you must specify either **BackupId*	- or **RestoreTime**.
	//
	// example:
	//
	// 259321****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The ID of the source instance. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp172446ys9cf****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The names of the databases to be copied. Specify the names in the following format: `{"Source database name":"Target database name"}`. Separate multiple databases with commas (,). Examples:
	//
	// - Copy a single database: `{"zhttest":"zhttest"}`
	//
	// - Copy multiple databases: `{"zhttest01":"zhttest01","zhttest02":"zhttest02"}`
	//
	// > The name of the database in the target instance can be different from that in the source instance. However, before copying, make sure that the target instance does not contain a database with the same name.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"zhttest":"zhttest"}
	DbNames         *string `json:"DbNames,omitempty" xml:"DbNames,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The point in time to which you want to restore the database. You can select any point in time within the backup retention period. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time must be in UTC.
	//
	// >When you call this operation, you must specify either **BackupId*	- or **RestoreTime**.
	//
	// example:
	//
	// 2025-06-08T17:41:14Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// Specifies whether to enable the replication of users and permissions. Valid values:
	//
	// 	- **YES**: Enables the replication of users and permissions. If the target instance contains a user with the same name as a user in the source instance, the permissions of the source user are added to the target user.
	//
	// 	- **NO**: Disables the replication of users and permissions. This is the default value.
	//
	// example:
	//
	// NO
	SyncUserPrivilege *string `json:"SyncUserPrivilege,omitempty" xml:"SyncUserPrivilege,omitempty"`
	// The ID of the target instance. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp1m71wvzfiq7****
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
