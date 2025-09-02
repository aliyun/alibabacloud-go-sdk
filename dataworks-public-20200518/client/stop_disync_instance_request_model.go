// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDISyncInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v int64) *StopDISyncInstanceRequest
	GetFileId() *int64
	SetProjectId(v int64) *StopDISyncInstanceRequest
	GetProjectId() *int64
	SetTaskType(v string) *StopDISyncInstanceRequest
	GetTaskType() *string
}

type StopDISyncInstanceRequest struct {
	// The ID of the synchronization task. You can call the [ListFiles](https://help.aliyun.com/document_detail/173942.html) operation to query the ID.
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
	// The type of the synchronization task that you want to stop. Set the value to DI_REALTIME.
	//
	// DI_REALTIME indicates a real-time synchronization task.
	//
	// This parameter is required.
	//
	// example:
	//
	// DI_REALTIME
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s StopDISyncInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StopDISyncInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopDISyncInstanceRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *StopDISyncInstanceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *StopDISyncInstanceRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *StopDISyncInstanceRequest) SetFileId(v int64) *StopDISyncInstanceRequest {
	s.FileId = &v
	return s
}

func (s *StopDISyncInstanceRequest) SetProjectId(v int64) *StopDISyncInstanceRequest {
	s.ProjectId = &v
	return s
}

func (s *StopDISyncInstanceRequest) SetTaskType(v string) *StopDISyncInstanceRequest {
	s.TaskType = &v
	return s
}

func (s *StopDISyncInstanceRequest) Validate() error {
	return dara.Validate(s)
}
