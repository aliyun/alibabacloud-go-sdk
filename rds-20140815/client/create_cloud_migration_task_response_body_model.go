// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudMigrationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreateCloudMigrationTaskResponseBody
	GetDBInstanceName() *string
	SetRequestId(v string) *CreateCloudMigrationTaskResponseBody
	GetRequestId() *string
	SetTaskId(v int64) *CreateCloudMigrationTaskResponseBody
	GetTaskId() *int64
	SetTaskName(v string) *CreateCloudMigrationTaskResponseBody
	GetTaskName() *string
}

type CreateCloudMigrationTaskResponseBody struct {
	// The name of the instance.
	//
	// example:
	//
	// pgm-bp102g323jd4****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8B993DA9-5272-5414-94E3-4CA8BA0146C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 440437220
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// 362c6c7a-4d20-4eac-898c-1495ceab374c
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateCloudMigrationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudMigrationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudMigrationTaskResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateCloudMigrationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCloudMigrationTaskResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CreateCloudMigrationTaskResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateCloudMigrationTaskResponseBody) SetDBInstanceName(v string) *CreateCloudMigrationTaskResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *CreateCloudMigrationTaskResponseBody) SetRequestId(v string) *CreateCloudMigrationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCloudMigrationTaskResponseBody) SetTaskId(v int64) *CreateCloudMigrationTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateCloudMigrationTaskResponseBody) SetTaskName(v string) *CreateCloudMigrationTaskResponseBody {
	s.TaskName = &v
	return s
}

func (s *CreateCloudMigrationTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
