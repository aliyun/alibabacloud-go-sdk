// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserUsageDataExportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *CreateUserUsageDataExportTaskResponseBody
	GetEndTime() *string
	SetRequestId(v string) *CreateUserUsageDataExportTaskResponseBody
	GetRequestId() *string
	SetStartTime(v string) *CreateUserUsageDataExportTaskResponseBody
	GetStartTime() *string
	SetTaskId(v string) *CreateUserUsageDataExportTaskResponseBody
	GetTaskId() *string
}

type CreateUserUsageDataExportTaskResponseBody struct {
	// The end of the time range during which data was queried.
	//
	// example:
	//
	// 2015-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ED61C6C3-8241-4187-AAA7-5157AE175CEC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 129456
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateUserUsageDataExportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserUsageDataExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserUsageDataExportTaskResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateUserUsageDataExportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserUsageDataExportTaskResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateUserUsageDataExportTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateUserUsageDataExportTaskResponseBody) SetEndTime(v string) *CreateUserUsageDataExportTaskResponseBody {
	s.EndTime = &v
	return s
}

func (s *CreateUserUsageDataExportTaskResponseBody) SetRequestId(v string) *CreateUserUsageDataExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserUsageDataExportTaskResponseBody) SetStartTime(v string) *CreateUserUsageDataExportTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *CreateUserUsageDataExportTaskResponseBody) SetTaskId(v string) *CreateUserUsageDataExportTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateUserUsageDataExportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
