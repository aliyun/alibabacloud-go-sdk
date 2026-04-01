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
	// The type of the migration task. Valid values:
	//
	// 	- **FULL**: The migration task migrates full backup files.
	//
	// 	- **UPDF**: The migration task migrates incremental or log backup files.
	//
	// This parameter is required.
	//
	// example:
	//
	// FULL
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The consistency check method for the database. Valid values:
	//
	// 	- **SyncExecuteDBCheck**: synchronous database check
	//
	// 	- **AsyncExecuteDBCheck**: asynchronous database check
	//
	// Default value: **AsyncExecuteDBCheck*	- (compatible with SQL Server 2008 R2)
	//
	// >  This parameter is valid when **IsOnlineDB*	- is set to **True**.
	//
	// example:
	//
	// AsyncExecuteDBCheck
	CheckDBMode *string `json:"CheckDBMode,omitempty" xml:"CheckDBMode,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk******
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the destination database.
	//
	// This parameter is required.
	//
	// example:
	//
	// testDB
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// Specifies whether to make the restored database data available for user access. Valid values:
	//
	// 	- **True**
	//
	// 	- **False**
	//
	// >  Set the value to **True*	- for instances that run SQL Server 2008 R2.
	//
	// This parameter is required.
	//
	// example:
	//
	// True
	IsOnlineDB *string `json:"IsOnlineDB,omitempty" xml:"IsOnlineDB,omitempty"`
	// The migration task ID.
	//
	// 	- If you set **BackupMode*	- to **FULL**, the value of this parameter is empty. The full backup mode is compatible with instance that runs SQL Server 2008 R2.
	//
	// 	- If you set **BackupMode*	- to **UPDF**, the value of this parameter is the ID of the required full migration task.
	//
	// > 	- If you set **IsOnlineDB*	- to **True**, the value of **BackupMode*	- must be **FULL**.
	//
	// > 	- If you set **IsOnlineDB*	- to **False**, the value of **BackupMode*	- must be **UPDF**.
	//
	// example:
	//
	// None
	MigrateTaskId *string `json:"MigrateTaskId,omitempty" xml:"MigrateTaskId,omitempty"`
	// The shared URL of the backup file in the OSS bucket. The URL must be encoded.
	//
	// If you specify multiple URLs, separate them with vertical bars (|) and then encode them.
	//
	// >  This parameter is required for instances that run SQL Server 2008 R2.
	//
	// example:
	//
	// check_cdn_oss.sh www.xxxxxx.mobi
	OSSUrls *string `json:"OSSUrls,omitempty" xml:"OSSUrls,omitempty"`
	// The information about the backup file in the OSS bucket. The values consist of three parts that are separated by colons (:):
	//
	// 	- OSS endpoint: oss-ap-southeast-1.aliyuncs.com.
	//
	// 	- Name of the OSS bucket: rdsmssqlsingapore.
	//
	// 	- Key of the backup file in the OSS bucket: autotest_2008R2_TestMigration_FULL.bak.
	//
	// > 	- This parameter is optional for instances that run SQL Server 2008 R2.
	//
	// > 	- This parameter is required for instances that run a major engine version later than SQL Server 2008 R2.
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
