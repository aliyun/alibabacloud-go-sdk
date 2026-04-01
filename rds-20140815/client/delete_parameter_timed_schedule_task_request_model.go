// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteParameterTimedScheduleTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DeleteParameterTimedScheduleTaskRequest
	GetDBInstanceName() *string
	SetTaskId(v int64) *DeleteParameterTimedScheduleTaskRequest
	GetTaskId() *int64
}

type DeleteParameterTimedScheduleTaskRequest struct {
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// 41698
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteParameterTimedScheduleTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteParameterTimedScheduleTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteParameterTimedScheduleTaskRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeleteParameterTimedScheduleTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DeleteParameterTimedScheduleTaskRequest) SetDBInstanceName(v string) *DeleteParameterTimedScheduleTaskRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DeleteParameterTimedScheduleTaskRequest) SetTaskId(v int64) *DeleteParameterTimedScheduleTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DeleteParameterTimedScheduleTaskRequest) Validate() error {
	return dara.Validate(s)
}
