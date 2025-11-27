// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebuildDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMigrationId(v int32) *RebuildDBInstanceResponseBody
	GetMigrationId() *int32
	SetRequestId(v string) *RebuildDBInstanceResponseBody
	GetRequestId() *string
	SetTaskId(v int32) *RebuildDBInstanceResponseBody
	GetTaskId() *int32
}

type RebuildDBInstanceResponseBody struct {
	// The serial number of the task in the rebuild task queue. When the serial number becomes 0, the system starts to rebuild the secondary instance.
	//
	// example:
	//
	// 329****
	MigrationId *int32 `json:"MigrationId,omitempty" xml:"MigrationId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 355DA57C-8CC4-40AB-B3F8-B684BA32EB9E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 20867****
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RebuildDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebuildDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RebuildDBInstanceResponseBody) GetMigrationId() *int32 {
	return s.MigrationId
}

func (s *RebuildDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebuildDBInstanceResponseBody) GetTaskId() *int32 {
	return s.TaskId
}

func (s *RebuildDBInstanceResponseBody) SetMigrationId(v int32) *RebuildDBInstanceResponseBody {
	s.MigrationId = &v
	return s
}

func (s *RebuildDBInstanceResponseBody) SetRequestId(v string) *RebuildDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebuildDBInstanceResponseBody) SetTaskId(v int32) *RebuildDBInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *RebuildDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
