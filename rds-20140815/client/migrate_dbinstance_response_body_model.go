// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMigrationId(v int32) *MigrateDBInstanceResponseBody
	GetMigrationId() *int32
	SetRequestId(v string) *MigrateDBInstanceResponseBody
	GetRequestId() *string
	SetTaskId(v int32) *MigrateDBInstanceResponseBody
	GetTaskId() *int32
}

type MigrateDBInstanceResponseBody struct {
	// The serial number of the task in the migration task queue. When the serial number becomes 0, the system starts the migration.
	//
	// example:
	//
	// 224****
	MigrationId *int32 `json:"MigrationId,omitempty" xml:"MigrationId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 498AE8CA-8C81-4A01-AF37-2B902014ED30
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 10824****
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s MigrateDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MigrateDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateDBInstanceResponseBody) GetMigrationId() *int32 {
	return s.MigrationId
}

func (s *MigrateDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MigrateDBInstanceResponseBody) GetTaskId() *int32 {
	return s.TaskId
}

func (s *MigrateDBInstanceResponseBody) SetMigrationId(v int32) *MigrateDBInstanceResponseBody {
	s.MigrationId = &v
	return s
}

func (s *MigrateDBInstanceResponseBody) SetRequestId(v string) *MigrateDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *MigrateDBInstanceResponseBody) SetTaskId(v int32) *MigrateDBInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *MigrateDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
