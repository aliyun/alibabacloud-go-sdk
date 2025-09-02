// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDISyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v int64) *DeleteDISyncTaskRequest
	GetFileId() *int64
	SetProjectId(v int64) *DeleteDISyncTaskRequest
	GetProjectId() *int64
	SetTaskType(v string) *DeleteDISyncTaskRequest
	GetTaskType() *string
}

type DeleteDISyncTaskRequest struct {
	// The ID of the real-time synchronization task. You can call the [ListFiles](https://help.aliyun.com/document_detail/173942.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the workspace ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The type of the task. Set the value to DI_REALTIME, which indicates a real-time synchronization task.
	//
	// This parameter is required.
	//
	// example:
	//
	// DI_REALTIME
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DeleteDISyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDISyncTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteDISyncTaskRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *DeleteDISyncTaskRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteDISyncTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DeleteDISyncTaskRequest) SetFileId(v int64) *DeleteDISyncTaskRequest {
	s.FileId = &v
	return s
}

func (s *DeleteDISyncTaskRequest) SetProjectId(v int64) *DeleteDISyncTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteDISyncTaskRequest) SetTaskType(v string) *DeleteDISyncTaskRequest {
	s.TaskType = &v
	return s
}

func (s *DeleteDISyncTaskRequest) Validate() error {
	return dara.Validate(s)
}
