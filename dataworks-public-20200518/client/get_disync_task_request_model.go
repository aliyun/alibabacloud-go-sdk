// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDISyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v int64) *GetDISyncTaskRequest
	GetFileId() *int64
	SetProjectId(v int64) *GetDISyncTaskRequest
	GetProjectId() *int64
	SetTaskType(v string) *GetDISyncTaskRequest
	GetTaskType() *string
}

type GetDISyncTaskRequest struct {
	// 	- If you set TaskType to DI_REALTIME, set this parameter to the ID of the real-time synchronization task that you want to deploy.
	//
	// 	- If you set TaskType to DI_SOLUTION, set this parameter to the ID of the data synchronization solution that you want to deploy.
	//
	// You can call the [ListFiles](https://help.aliyun.com/document_detail/173942.html) operation to query the ID of the real-time synchronization task or data synchronization solution.
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
	// The type of the object that you want to query. Valid values:
	//
	// 	- DI_REALTIME: real-time synchronization task
	//
	// 	- DI_SOLUTION: data synchronization solution
	//
	// This parameter is required.
	//
	// example:
	//
	// DI_REALTIME
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetDISyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDISyncTaskRequest) GoString() string {
	return s.String()
}

func (s *GetDISyncTaskRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *GetDISyncTaskRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDISyncTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *GetDISyncTaskRequest) SetFileId(v int64) *GetDISyncTaskRequest {
	s.FileId = &v
	return s
}

func (s *GetDISyncTaskRequest) SetProjectId(v int64) *GetDISyncTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDISyncTaskRequest) SetTaskType(v string) *GetDISyncTaskRequest {
	s.TaskType = &v
	return s
}

func (s *GetDISyncTaskRequest) Validate() error {
	return dara.Validate(s)
}
