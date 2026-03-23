// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudMigrationPrecheckTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreateCloudMigrationPrecheckTaskResponseBody
	GetDBInstanceName() *string
	SetRequestId(v string) *CreateCloudMigrationPrecheckTaskResponseBody
	GetRequestId() *string
	SetTaskId(v int64) *CreateCloudMigrationPrecheckTaskResponseBody
	GetTaskId() *int64
	SetTaskName(v string) *CreateCloudMigrationPrecheckTaskResponseBody
	GetTaskName() *string
}

type CreateCloudMigrationPrecheckTaskResponseBody struct {
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId         *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName       *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateCloudMigrationPrecheckTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudMigrationPrecheckTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudMigrationPrecheckTaskResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateCloudMigrationPrecheckTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCloudMigrationPrecheckTaskResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CreateCloudMigrationPrecheckTaskResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateCloudMigrationPrecheckTaskResponseBody) SetDBInstanceName(v string) *CreateCloudMigrationPrecheckTaskResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *CreateCloudMigrationPrecheckTaskResponseBody) SetRequestId(v string) *CreateCloudMigrationPrecheckTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCloudMigrationPrecheckTaskResponseBody) SetTaskId(v int64) *CreateCloudMigrationPrecheckTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateCloudMigrationPrecheckTaskResponseBody) SetTaskName(v string) *CreateCloudMigrationPrecheckTaskResponseBody {
	s.TaskName = &v
	return s
}

func (s *CreateCloudMigrationPrecheckTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
