// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMigrateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupMode(v string) *CreateMigrateTaskRequest
	GetBackupMode() *string
	SetCheckDBMode(v string) *CreateMigrateTaskRequest
	GetCheckDBMode() *string
	SetDBInstanceId(v string) *CreateMigrateTaskRequest
	GetDBInstanceId() *string
	SetDBName(v string) *CreateMigrateTaskRequest
	GetDBName() *string
	SetIsOnlineDB(v string) *CreateMigrateTaskRequest
	GetIsOnlineDB() *string
	SetMigrateTaskId(v string) *CreateMigrateTaskRequest
	GetMigrateTaskId() *string
	SetOSSUrls(v string) *CreateMigrateTaskRequest
	GetOSSUrls() *string
	SetOssObjectPositions(v string) *CreateMigrateTaskRequest
	GetOssObjectPositions() *string
	SetOwnerId(v int64) *CreateMigrateTaskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateMigrateTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateMigrateTaskRequest
	GetResourceOwnerId() *int64
}

type CreateMigrateTaskRequest struct {
	// The type of the cloud migration task. Valid values:
	//
	// 	- **FULL**: performs a restore operation by using a full backup file. This value is applicable to first-time migrations or full data recovery scenarios.
	//
	// 	- **UPDF**: restores incremental data by using an incremental backup file or log file. This value is applicable to incremental synchronization scenarios where a full backup already exists.
	//
	// This parameter is required.
	//
	// example:
	//
	// FULL
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The consistency check method after the database is brought online. This parameter takes effect only when IsOnlineDB is set to True. Valid values:
	//
	// - **SyncExecuteDBCheck**: performs a synchronous database check. This value is applicable to scenarios that require high data consistency.
	//
	// - **AsyncExecuteDBCheck**: performs an asynchronous database check. This value provides higher performance but may delay the detection of potential issues.
	//
	// Default value: **AsyncExecuteDBCheck*	- (compatible with SQL Server 2008 R2).
	//
	// example:
	//
	// AsyncExecuteDBCheck
	CheckDBMode *string `json:"CheckDBMode,omitempty" xml:"CheckDBMode,omitempty"`
	// The instance ID. You can call DescribeDBInstances to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the destination database.
	//
	// This parameter is required.
	//
	// example:
	//
	// testDB
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// Specifies whether to bring the restored database online so that users can access it. Valid values:
	//
	// 	- **True**: Brings the database online.
	//
	// 	- **False**: Does not bring the database online.
	//
	// > 	- For SQL Server 2008 R2, this value is always True.
	//
	// > 	- When **IsOnlineDB*	- is set to **True**, **BackupMode*	- must be set to **FULL**.
	//
	// > 	- When **IsOnlineDB*	- is set to **False**, **BackupMode*	- must be set to **UPDF**.
	//
	// This parameter is required.
	//
	// example:
	//
	// True
	IsOnlineDB *string `json:"IsOnlineDB,omitempty" xml:"IsOnlineDB,omitempty"`
	// The migration task ID. Valid values:
	//
	// - When **BackupMode*	- is set to **FULL**, leave this parameter empty (compatible with SQL Server 2008 R2).
	//
	// - When **BackupMode*	- is set to **UPDF**, set this parameter to the ID of the corresponding FULL task. You can call DescribeMigrateTasks to query the task ID.
	//
	// example:
	//
	// None
	MigrateTaskId *string `json:"MigrateTaskId,omitempty" xml:"MigrateTaskId,omitempty"`
	// The shared URL of the backup file on OSS (URL-encoded). If multiple URLs exist, separate them with vertical bars (|) before encoding, and then pass the encoded value.
	//
	// > This parameter is required for SQL Server 2008 R2.
	//
	// example:
	//
	// check_cdn_oss.sh www.******.mobi
	OSSUrls *string `json:"OSSUrls,omitempty" xml:"OSSUrls,omitempty"`
	// The OSS file information, which consists of the following three parts separated by colons (:):
	//
	// - **OSS endpoint**: oss-ap-southeast-1.aliyuncs.com.
	//
	// - **OSS bucket name**: rdsmssqlsingapore.
	//
	// - **Backup file name on OSS**: autotest_2008R2_TestMigration_FULL.bak.
	//
	// > This parameter is required for SQL Server versions later than SQL Server 2008 R2.
	//
	// example:
	//
	// oss-ap-southeast-1.aliyuncs.com:rdsmssqlsingapore:autotest_2008R2_TestMigration_FULL.bak
	OssObjectPositions   *string `json:"OssObjectPositions,omitempty" xml:"OssObjectPositions,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateMigrateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMigrateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateMigrateTaskRequest) GetBackupMode() *string {
	return s.BackupMode
}

func (s *CreateMigrateTaskRequest) GetCheckDBMode() *string {
	return s.CheckDBMode
}

func (s *CreateMigrateTaskRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateMigrateTaskRequest) GetDBName() *string {
	return s.DBName
}

func (s *CreateMigrateTaskRequest) GetIsOnlineDB() *string {
	return s.IsOnlineDB
}

func (s *CreateMigrateTaskRequest) GetMigrateTaskId() *string {
	return s.MigrateTaskId
}

func (s *CreateMigrateTaskRequest) GetOSSUrls() *string {
	return s.OSSUrls
}

func (s *CreateMigrateTaskRequest) GetOssObjectPositions() *string {
	return s.OssObjectPositions
}

func (s *CreateMigrateTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateMigrateTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateMigrateTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateMigrateTaskRequest) SetBackupMode(v string) *CreateMigrateTaskRequest {
	s.BackupMode = &v
	return s
}

func (s *CreateMigrateTaskRequest) SetCheckDBMode(v string) *CreateMigrateTaskRequest {
	s.CheckDBMode = &v
	return s
}

func (s *CreateMigrateTaskRequest) SetDBInstanceId(v string) *CreateMigrateTaskRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateMigrateTaskRequest) SetDBName(v string) *CreateMigrateTaskRequest {
	s.DBName = &v
	return s
}

func (s *CreateMigrateTaskRequest) SetIsOnlineDB(v string) *CreateMigrateTaskRequest {
	s.IsOnlineDB = &v
	return s
}

func (s *CreateMigrateTaskRequest) SetMigrateTaskId(v string) *CreateMigrateTaskRequest {
	s.MigrateTaskId = &v
	return s
}

func (s *CreateMigrateTaskRequest) SetOSSUrls(v string) *CreateMigrateTaskRequest {
	s.OSSUrls = &v
	return s
}

func (s *CreateMigrateTaskRequest) SetOssObjectPositions(v string) *CreateMigrateTaskRequest {
	s.OssObjectPositions = &v
	return s
}

func (s *CreateMigrateTaskRequest) SetOwnerId(v int64) *CreateMigrateTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMigrateTaskRequest) SetResourceOwnerAccount(v string) *CreateMigrateTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateMigrateTaskRequest) SetResourceOwnerId(v int64) *CreateMigrateTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateMigrateTaskRequest) Validate() error {
	return dara.Validate(s)
}
