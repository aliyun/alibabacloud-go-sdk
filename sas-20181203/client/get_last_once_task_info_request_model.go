// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLastOnceTaskInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSource(v string) *GetLastOnceTaskInfoRequest
	GetSource() *string
	SetTaskName(v string) *GetLastOnceTaskInfoRequest
	GetTaskName() *string
	SetTaskType(v string) *GetLastOnceTaskInfoRequest
	GetTaskType() *string
}

type GetLastOnceTaskInfoRequest struct {
	// The source of the task.
	//
	// example:
	//
	// console_batch
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The name of the task. Valid values:
	//
	// 	- **CLIENT_PROBLEM_CHECK**: client diagnosis task
	//
	// 	- **CLIENT_DEV_OPS**: O\\&M task of Cloud Assistant
	//
	// 	- **ASSETS_COLLECTION**: asset collection task
	//
	// This parameter is required.
	//
	// example:
	//
	// ASSETS_COLLECTION
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the task. Valid values:
	//
	// 	- **CLIENT_PROBLEM_CHECK**: client diagnosis task
	//
	// 	- **CLIENT_DEV_OPS**: O\\&M task of Cloud Assistant
	//
	// 	- **ASSETS_COLLECTION**: asset collection task
	//
	// This parameter is required.
	//
	// example:
	//
	// ASSETS_COLLECTION
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetLastOnceTaskInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLastOnceTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *GetLastOnceTaskInfoRequest) GetSource() *string {
	return s.Source
}

func (s *GetLastOnceTaskInfoRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *GetLastOnceTaskInfoRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *GetLastOnceTaskInfoRequest) SetSource(v string) *GetLastOnceTaskInfoRequest {
	s.Source = &v
	return s
}

func (s *GetLastOnceTaskInfoRequest) SetTaskName(v string) *GetLastOnceTaskInfoRequest {
	s.TaskName = &v
	return s
}

func (s *GetLastOnceTaskInfoRequest) SetTaskType(v string) *GetLastOnceTaskInfoRequest {
	s.TaskType = &v
	return s
}

func (s *GetLastOnceTaskInfoRequest) Validate() error {
	return dara.Validate(s)
}
