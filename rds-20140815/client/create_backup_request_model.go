// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupMethod(v string) *CreateBackupRequest
	GetBackupMethod() *string
	SetBackupRetentionPeriod(v int64) *CreateBackupRequest
	GetBackupRetentionPeriod() *int64
	SetBackupStrategy(v string) *CreateBackupRequest
	GetBackupStrategy() *string
	SetBackupType(v string) *CreateBackupRequest
	GetBackupType() *string
	SetDBInstanceId(v string) *CreateBackupRequest
	GetDBInstanceId() *string
	SetDBName(v string) *CreateBackupRequest
	GetDBName() *string
	SetResourceOwnerId(v int64) *CreateBackupRequest
	GetResourceOwnerId() *int64
}

type CreateBackupRequest struct {
	// The backup type of the instance. Valid values:
	//
	// 	- **Logical**: logical backup
	//
	// 	- **Physical**: physical backup
	//
	// 	- **Snapshot**: snapshot backup
	//
	// Default value: **Physical**.
	//
	// > 	- You can perform a logical backup only when databases are created on the instance.
	//
	// > 	- When you perform a snapshot backup on an ApsaraDB RDS for MariaDB instance, you must set this parameter to **Physical**.
	//
	// > 	- For more information about the supported backup types, see [Use the data backup feature](https://help.aliyun.com/document_detail/98818.html).
	//
	// > 	- When you perform a snapshot backup on an ApsaraDB RDS for SQL Server instance that uses cloud disks, you must set this parameter to **Snapshot**.
	//
	// example:
	//
	// Physical
	BackupMethod          *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	BackupRetentionPeriod *int64  `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// The backup policy. Valid values:
	//
	// 	- **db**: a database-level backup.
	//
	// 	- **instance**: an instance-level backup.
	//
	// > You can specify this parameter when you perform a logical backup on an ApsaraDB RDS for MySQL instance. You can also specify this parameter when you perform a full physical backup on an ApsaraDB RDS for SQL Server instance.
	//
	// example:
	//
	// db
	BackupStrategy *string `json:"BackupStrategy,omitempty" xml:"BackupStrategy,omitempty"`
	// The backup method. Valid values:
	//
	// 	- **Auto**: full or incremental backup that is automatically selected
	//
	// 	- **FullBackup**: full backup
	//
	// Default value: **Auto**.
	//
	// > 	- You must set this parameter only when the instance runs SQL Server.
	//
	// > 	- This parameter is valid only when you set the **BackupMethod*	- parameter to **Physical**.
	//
	// example:
	//
	// Auto
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The names of the databases whose data you want to back up. Separate the names of the databases with commas (,).
	//
	// > You can specify this parameter when you perform a logical backup on individual databases of an ApsaraDB RDS for MySQL instance. You can also specify this parameter when you perform a full physical backup on individual databases of an ApsaraDB RDS for SQL Server instance.
	//
	// example:
	//
	// rds_mysql
	DBName          *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateBackupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupRequest) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *CreateBackupRequest) GetBackupRetentionPeriod() *int64 {
	return s.BackupRetentionPeriod
}

func (s *CreateBackupRequest) GetBackupStrategy() *string {
	return s.BackupStrategy
}

func (s *CreateBackupRequest) GetBackupType() *string {
	return s.BackupType
}

func (s *CreateBackupRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateBackupRequest) GetDBName() *string {
	return s.DBName
}

func (s *CreateBackupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateBackupRequest) SetBackupMethod(v string) *CreateBackupRequest {
	s.BackupMethod = &v
	return s
}

func (s *CreateBackupRequest) SetBackupRetentionPeriod(v int64) *CreateBackupRequest {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *CreateBackupRequest) SetBackupStrategy(v string) *CreateBackupRequest {
	s.BackupStrategy = &v
	return s
}

func (s *CreateBackupRequest) SetBackupType(v string) *CreateBackupRequest {
	s.BackupType = &v
	return s
}

func (s *CreateBackupRequest) SetDBInstanceId(v string) *CreateBackupRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateBackupRequest) SetDBName(v string) *CreateBackupRequest {
	s.DBName = &v
	return s
}

func (s *CreateBackupRequest) SetResourceOwnerId(v int64) *CreateBackupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateBackupRequest) Validate() error {
	return dara.Validate(s)
}
