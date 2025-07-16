// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUsageDetailDataExportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *CreateUsageDetailDataExportTaskResponseBody
	GetEndTime() *string
	SetRequestId(v string) *CreateUsageDetailDataExportTaskResponseBody
	GetRequestId() *string
	SetStartTime(v string) *CreateUsageDetailDataExportTaskResponseBody
	GetStartTime() *string
	SetTaskId(v string) *CreateUsageDetailDataExportTaskResponseBody
	GetTaskId() *string
}

type CreateUsageDetailDataExportTaskResponseBody struct {
	// The end of the time range for which the data was queried.
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
	// The beginning of the time range for which the data was queried.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 123456
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateUsageDetailDataExportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUsageDetailDataExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUsageDetailDataExportTaskResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateUsageDetailDataExportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUsageDetailDataExportTaskResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateUsageDetailDataExportTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateUsageDetailDataExportTaskResponseBody) SetEndTime(v string) *CreateUsageDetailDataExportTaskResponseBody {
	s.EndTime = &v
	return s
}

func (s *CreateUsageDetailDataExportTaskResponseBody) SetRequestId(v string) *CreateUsageDetailDataExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUsageDetailDataExportTaskResponseBody) SetStartTime(v string) *CreateUsageDetailDataExportTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *CreateUsageDetailDataExportTaskResponseBody) SetTaskId(v string) *CreateUsageDetailDataExportTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateUsageDetailDataExportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
