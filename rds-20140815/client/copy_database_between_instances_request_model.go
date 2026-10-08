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
	// The backup set ID of the source instance. To copy a database from a backup set, call DescribeBackups to query the backup set ID.
	//
	// >You must specify either **BackupId*	- or **RestoreTime**.
	//
	// example:
	//
	// 259321****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The source instance ID. You can call DescribeDBInstances to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp172446ys9cf****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The list of database names to be copied. Format: `{"Source database name":"Destination database name"}`. Separate multiple databases with commas (,). Examples:
	//
	// - Copy a single database: `{"zhttest":"zhttest"}`
	//
	// - Copy multiple databases: `{"zhttest01":"zhttest01","zhttest02":"zhttest02"}`
	//
	// > The database name on the target instance can be different from that on the source instance. However, make sure that the target instance does not contain a database with the same name before copying.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"zhttest":"zhttest"}
	DbNames         *string `json:"DbNames,omitempty" xml:"DbNames,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The point in time to which you want to copy the database. You can specify any point in time within the backup retention period. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// >You must specify either **BackupId*	- or **RestoreTime**.
	//
	// example:
	//
	// 2025-06-08T17:41:14Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// Specifies whether to copy users and permissions. Valid values:
	//
	// 	- **YES**: Users and permissions are copied. If the target instance contains a user with the same name, the permissions of the user on the source instance are merged with those of the user on the target instance.
	//
	// 	- **NO*	- (default): Users and permissions are not copied.
	//
	// example:
	//
	// NO
	SyncUserPrivilege *string `json:"SyncUserPrivilege,omitempty" xml:"SyncUserPrivilege,omitempty"`
	// The target instance ID. You can invoke DescribeDBInstances to query the instance ID.
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
