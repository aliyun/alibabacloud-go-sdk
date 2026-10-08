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
	// The backup type. Valid values:
	//
	// 	- **Logical**: logical backup. Only MySQL instances with local disks support this type.
	//
	// 	- **Physical**: physical backup. MySQL instances with local disks, SQL Server instances, and PostgreSQL instances support this type.
	//
	// 	- **Snapshot**: snapshot backup. MySQL instances with cloud disks, SQL Server instances, PostgreSQL instances, and MariaDB instances support this type.
	//
	// Default value: **Physical**.
	//
	// > 	- When you use logical backup, the database must contain data (the data cannot be empty).
	//
	// > 	- MariaDB instances support only snapshot backup. However, set this parameter to **Physical**.
	//
	// example:
	//
	// Physical
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// - **SQL Server**: When the BackupStrategy parameter is set to db, the BackupMethod parameter is set to Physical, and the BackupType parameter is set to FullBackup, you can specify the retention period of the backup set. Valid values: 7 to 730 days, or -1 (long-term retention (LTR)).
	//
	// - **MySQL**: You can specify the retention period of the backup set. Valid values: 7 to 730 days, or -1 (long-term retention (LTR)).
	//
	// example:
	//
	// 7
	BackupRetentionPeriod *int64 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// The backup strategy. Valid values:
	//
	// 	- **db**: single-database backup
	//
	// 	- **instance**: instance backup
	//
	// > This parameter takes effect only when the following conditions are met:
	//
	// > - MySQL: The **BackupMethod*	- parameter is set to **Logical**.
	//
	// > - SQL Server: The **BackupType*	- parameter is set to **FullBackup**.
	//
	// example:
	//
	// db
	BackupStrategy *string `json:"BackupStrategy,omitempty" xml:"BackupStrategy,omitempty"`
	// The backup method for SQL Server instances. Valid values:
	//
	// 	- **Auto*	- (default): automatically selects full backup or incremental backup.
	//
	// 	- **FullBackup**: full backup.
	//
	// > This parameter takes effect only when the **BackupMethod*	- parameter is set to **Physical**.
	//
	// example:
	//
	// Auto
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The instance ID. You can call DescribeDBInstances to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The list of databases. Separate multiple databases with commas (,).
	//
	// > This parameter takes effect only when the **BackupStrategy*	- parameter is set to **db**.
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
