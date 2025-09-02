// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployDISyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v int64) *DeployDISyncTaskRequest
	GetFileId() *int64
	SetProjectId(v int64) *DeployDISyncTaskRequest
	GetProjectId() *int64
	SetTaskType(v string) *DeployDISyncTaskRequest
	GetTaskType() *string
}

type DeployDISyncTaskRequest struct {
	// 	- If you set the TaskType parameter to DI_REALTIME, set the FileId parameter to the ID of the real-time synchronization task that you want to deploy.
	//
	// 	- If you set the TaskType parameter to DI_SOLUTION, set the FileId parameter to the ID of the synchronization task that you want to deploy.
	//
	// You can call the [ListFiles](https://help.aliyun.com/document_detail/173942.html) operation to query the ID of the real-time synchronization task or the synchronization task created in Data Integration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The type of the object that you want to deploy. Valid values:
	//
	// 	- DI_REALTIME: real-time synchronization node
	//
	// 	- DI_SOLUTION: data synchronization solution
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The type of the object that you want to deploy. Valid values:
	//
	// 	- DI_REALTIME: real-time synchronization task
	//
	// 	- DI_SOLUTION: synchronization task created in Data Integration
	//
	// This parameter is required.
	//
	// example:
	//
	// DI_REALTIME
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DeployDISyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeployDISyncTaskRequest) GoString() string {
	return s.String()
}

func (s *DeployDISyncTaskRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *DeployDISyncTaskRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeployDISyncTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DeployDISyncTaskRequest) SetFileId(v int64) *DeployDISyncTaskRequest {
	s.FileId = &v
	return s
}

func (s *DeployDISyncTaskRequest) SetProjectId(v int64) *DeployDISyncTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *DeployDISyncTaskRequest) SetTaskType(v string) *DeployDISyncTaskRequest {
	s.TaskType = &v
	return s
}

func (s *DeployDISyncTaskRequest) Validate() error {
	return dara.Validate(s)
}
