// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateDISyncInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v int64) *TerminateDISyncInstanceRequest
	GetFileId() *int64
	SetProjectId(v int64) *TerminateDISyncInstanceRequest
	GetProjectId() *int64
	SetTaskType(v string) *TerminateDISyncInstanceRequest
	GetTaskType() *string
}

type TerminateDISyncInstanceRequest struct {
	// The ID of the real-time synchronization task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the workspace ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The value DI_REALTIME indicates that the task is a real-time synchronization task.
	//
	// This parameter is required.
	//
	// example:
	//
	// DI_REALTIME
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s TerminateDISyncInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s TerminateDISyncInstanceRequest) GoString() string {
	return s.String()
}

func (s *TerminateDISyncInstanceRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *TerminateDISyncInstanceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *TerminateDISyncInstanceRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *TerminateDISyncInstanceRequest) SetFileId(v int64) *TerminateDISyncInstanceRequest {
	s.FileId = &v
	return s
}

func (s *TerminateDISyncInstanceRequest) SetProjectId(v int64) *TerminateDISyncInstanceRequest {
	s.ProjectId = &v
	return s
}

func (s *TerminateDISyncInstanceRequest) SetTaskType(v string) *TerminateDISyncInstanceRequest {
	s.TaskType = &v
	return s
}

func (s *TerminateDISyncInstanceRequest) Validate() error {
	return dara.Validate(s)
}
