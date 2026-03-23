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
	MigrationId *int32  `json:"MigrationId,omitempty" xml:"MigrationId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId      *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
