// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMigrateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupMode(v string) *CreateMigrateTaskResponseBody
	GetBackupMode() *string
	SetDBInstanceId(v string) *CreateMigrateTaskResponseBody
	GetDBInstanceId() *string
	SetDBName(v string) *CreateMigrateTaskResponseBody
	GetDBName() *string
	SetMigrateTaskId(v string) *CreateMigrateTaskResponseBody
	GetMigrateTaskId() *string
	SetRequestId(v string) *CreateMigrateTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateMigrateTaskResponseBody
	GetTaskId() *string
}

type CreateMigrateTaskResponseBody struct {
	// The type of the migration task. Valid values:
	//
	// 	- **FULL**: The migration task migrates full backup files.
	//
	// 	- **UPDF**: The migration task migrates incremental or log backup files.
	//
	// example:
	//
	// FULL
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk******
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// test02
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The ID of the migration task.
	//
	// example:
	//
	// 564******
	MigrateTaskId *string `json:"MigrateTaskId,omitempty" xml:"MigrateTaskId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 866F5EB8-4650-4061-87F0-379F6F968BCE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 545****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateMigrateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMigrateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMigrateTaskResponseBody) GetBackupMode() *string {
	return s.BackupMode
}

func (s *CreateMigrateTaskResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateMigrateTaskResponseBody) GetDBName() *string {
	return s.DBName
}

func (s *CreateMigrateTaskResponseBody) GetMigrateTaskId() *string {
	return s.MigrateTaskId
}

func (s *CreateMigrateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMigrateTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateMigrateTaskResponseBody) SetBackupMode(v string) *CreateMigrateTaskResponseBody {
	s.BackupMode = &v
	return s
}

func (s *CreateMigrateTaskResponseBody) SetDBInstanceId(v string) *CreateMigrateTaskResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateMigrateTaskResponseBody) SetDBName(v string) *CreateMigrateTaskResponseBody {
	s.DBName = &v
	return s
}

func (s *CreateMigrateTaskResponseBody) SetMigrateTaskId(v string) *CreateMigrateTaskResponseBody {
	s.MigrateTaskId = &v
	return s
}

func (s *CreateMigrateTaskResponseBody) SetRequestId(v string) *CreateMigrateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMigrateTaskResponseBody) SetTaskId(v string) *CreateMigrateTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateMigrateTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
