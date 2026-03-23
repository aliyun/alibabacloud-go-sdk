// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMigrationId(v int32) *StartDBInstanceResponseBody
	GetMigrationId() *int32
	SetRequestId(v string) *StartDBInstanceResponseBody
	GetRequestId() *string
	SetTaskId(v int32) *StartDBInstanceResponseBody
	GetTaskId() *int32
}

type StartDBInstanceResponseBody struct {
	MigrationId *int32  `json:"MigrationId,omitempty" xml:"MigrationId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId      *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartDBInstanceResponseBody) GetMigrationId() *int32 {
	return s.MigrationId
}

func (s *StartDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartDBInstanceResponseBody) GetTaskId() *int32 {
	return s.TaskId
}

func (s *StartDBInstanceResponseBody) SetMigrationId(v int32) *StartDBInstanceResponseBody {
	s.MigrationId = &v
	return s
}

func (s *StartDBInstanceResponseBody) SetRequestId(v string) *StartDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDBInstanceResponseBody) SetTaskId(v int32) *StartDBInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *StartDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
