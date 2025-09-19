// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOnceTaskResultInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetOnceTaskResultInfoRequest
	GetTaskId() *string
	SetTaskName(v string) *GetOnceTaskResultInfoRequest
	GetTaskName() *string
	SetTaskType(v string) *GetOnceTaskResultInfoRequest
	GetTaskType() *string
}

type GetOnceTaskResultInfoRequest struct {
	// The ID of the scan task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9dfa3a7eb9547781632785b49003****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the task. Valid values:
	//
	// 	- **CLIENT_PROBLEM_CHECK**: a task of the Security Center agent
	//
	// 	- **CLIENT_DEV_OPS**: an O\\&M task of Cloud Assistant
	//
	// 	- **ASSET_SECURITY_CHECK**: a task of asset information collection
	//
	// This parameter is required.
	//
	// example:
	//
	// ASSETS_COLLECTION
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the task. Valid values:
	//
	// 	- **CLIENT_PROBLEM_CHECK**: a task of the Security Center agent
	//
	// 	- **CLIENT_DEV_OPS**: an O\\&M task of Cloud Assistant
	//
	// 	- **ASSET_SECURITY_CHECK**: a task of asset information collection
	//
	// This parameter is required.
	//
	// example:
	//
	// ASSETS_COLLECTION
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetOnceTaskResultInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOnceTaskResultInfoRequest) GoString() string {
	return s.String()
}

func (s *GetOnceTaskResultInfoRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetOnceTaskResultInfoRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *GetOnceTaskResultInfoRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *GetOnceTaskResultInfoRequest) SetTaskId(v string) *GetOnceTaskResultInfoRequest {
	s.TaskId = &v
	return s
}

func (s *GetOnceTaskResultInfoRequest) SetTaskName(v string) *GetOnceTaskResultInfoRequest {
	s.TaskName = &v
	return s
}

func (s *GetOnceTaskResultInfoRequest) SetTaskType(v string) *GetOnceTaskResultInfoRequest {
	s.TaskType = &v
	return s
}

func (s *GetOnceTaskResultInfoRequest) Validate() error {
	return dara.Validate(s)
}
