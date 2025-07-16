// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserUsageDataExportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *CreateUserUsageDataExportTaskRequest
	GetEndTime() *string
	SetLanguage(v string) *CreateUserUsageDataExportTaskRequest
	GetLanguage() *string
	SetStartTime(v string) *CreateUserUsageDataExportTaskRequest
	GetStartTime() *string
	SetTaskName(v string) *CreateUserUsageDataExportTaskRequest
	GetTaskName() *string
}

type CreateUserUsageDataExportTaskRequest struct {
	// The end of the time range to query. The end time must be later than the start time.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2015-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language in which you want to export the file. Default value: zh-cn. Valid values:
	//
	// 	- **zh-cn**: Chinese
	//
	// 	- **en-us**: English
	//
	// example:
	//
	// zh-cn
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The start of the time range to query. The data is collected every 5 minutes.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// Refresh
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateUserUsageDataExportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserUsageDataExportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateUserUsageDataExportTaskRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateUserUsageDataExportTaskRequest) GetLanguage() *string {
	return s.Language
}

func (s *CreateUserUsageDataExportTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateUserUsageDataExportTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateUserUsageDataExportTaskRequest) SetEndTime(v string) *CreateUserUsageDataExportTaskRequest {
	s.EndTime = &v
	return s
}

func (s *CreateUserUsageDataExportTaskRequest) SetLanguage(v string) *CreateUserUsageDataExportTaskRequest {
	s.Language = &v
	return s
}

func (s *CreateUserUsageDataExportTaskRequest) SetStartTime(v string) *CreateUserUsageDataExportTaskRequest {
	s.StartTime = &v
	return s
}

func (s *CreateUserUsageDataExportTaskRequest) SetTaskName(v string) *CreateUserUsageDataExportTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateUserUsageDataExportTaskRequest) Validate() error {
	return dara.Validate(s)
}
