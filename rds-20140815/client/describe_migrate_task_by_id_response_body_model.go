// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMigrateTaskByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupMode(v string) *DescribeMigrateTaskByIdResponseBody
	GetBackupMode() *string
	SetCreateTime(v string) *DescribeMigrateTaskByIdResponseBody
	GetCreateTime() *string
	SetDBInstanceName(v string) *DescribeMigrateTaskByIdResponseBody
	GetDBInstanceName() *string
	SetDBName(v string) *DescribeMigrateTaskByIdResponseBody
	GetDBName() *string
	SetDescription(v string) *DescribeMigrateTaskByIdResponseBody
	GetDescription() *string
	SetEndTime(v string) *DescribeMigrateTaskByIdResponseBody
	GetEndTime() *string
	SetIsDBReplaced(v string) *DescribeMigrateTaskByIdResponseBody
	GetIsDBReplaced() *string
	SetMigrateTaskId(v string) *DescribeMigrateTaskByIdResponseBody
	GetMigrateTaskId() *string
	SetRequestId(v string) *DescribeMigrateTaskByIdResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeMigrateTaskByIdResponseBody
	GetStatus() *string
}

type DescribeMigrateTaskByIdResponseBody struct {
	// The type of the migration task. Valid values:
	//
	// 	- **FULL**: The migration task migrates full backup files that can be used to restore the full data of the instance.
	//
	// 	- **UPDF**: The migration task migrates incremental or log backup files that can be used to restore the incremental data of the instance.
	//
	// example:
	//
	// FULL
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The time when the migration task was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-05-30T12:11:04Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// mytestdb
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The description of the migration task.
	//
	// example:
	//
	// Success to DBCC checkdb asynchronously
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the migration task was completed. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-05-30T15:15:05Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Indicates whether the imported data overwrites the existing data. Valid values:
	//
	// 	- **False**: The imported data does not overwrite the existing data.
	//
	// 	- **True**: The imported data overwrites the existing data.
	//
	// example:
	//
	// False
	IsDBReplaced *string `json:"IsDBReplaced,omitempty" xml:"IsDBReplaced,omitempty"`
	// The ID of the migration task.
	//
	// example:
	//
	// 235943
	MigrateTaskId *string `json:"MigrateTaskId,omitempty" xml:"MigrateTaskId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6ED3635A-01F9-47BD-B9C8-CB3FD70A336E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the migration task. Valid values:
	//
	// 	- **NoStart**: The task has not started.
	//
	// 	- **Running**:The task is in progress.
	//
	// 	- **Success**: The task is successful.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Waiting**: The task is waiting for an incremental backup file to be imported.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMigrateTaskByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrateTaskByIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMigrateTaskByIdResponseBody) GetBackupMode() *string {
	return s.BackupMode
}

func (s *DescribeMigrateTaskByIdResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeMigrateTaskByIdResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeMigrateTaskByIdResponseBody) GetDBName() *string {
	return s.DBName
}

func (s *DescribeMigrateTaskByIdResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeMigrateTaskByIdResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeMigrateTaskByIdResponseBody) GetIsDBReplaced() *string {
	return s.IsDBReplaced
}

func (s *DescribeMigrateTaskByIdResponseBody) GetMigrateTaskId() *string {
	return s.MigrateTaskId
}

func (s *DescribeMigrateTaskByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMigrateTaskByIdResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeMigrateTaskByIdResponseBody) SetBackupMode(v string) *DescribeMigrateTaskByIdResponseBody {
	s.BackupMode = &v
	return s
}

func (s *DescribeMigrateTaskByIdResponseBody) SetCreateTime(v string) *DescribeMigrateTaskByIdResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeMigrateTaskByIdResponseBody) SetDBInstanceName(v string) *DescribeMigrateTaskByIdResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeMigrateTaskByIdResponseBody) SetDBName(v string) *DescribeMigrateTaskByIdResponseBody {
	s.DBName = &v
	return s
}

func (s *DescribeMigrateTaskByIdResponseBody) SetDescription(v string) *DescribeMigrateTaskByIdResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeMigrateTaskByIdResponseBody) SetEndTime(v string) *DescribeMigrateTaskByIdResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeMigrateTaskByIdResponseBody) SetIsDBReplaced(v string) *DescribeMigrateTaskByIdResponseBody {
	s.IsDBReplaced = &v
	return s
}

func (s *DescribeMigrateTaskByIdResponseBody) SetMigrateTaskId(v string) *DescribeMigrateTaskByIdResponseBody {
	s.MigrateTaskId = &v
	return s
}

func (s *DescribeMigrateTaskByIdResponseBody) SetRequestId(v string) *DescribeMigrateTaskByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMigrateTaskByIdResponseBody) SetStatus(v string) *DescribeMigrateTaskByIdResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeMigrateTaskByIdResponseBody) Validate() error {
	return dara.Validate(s)
}
