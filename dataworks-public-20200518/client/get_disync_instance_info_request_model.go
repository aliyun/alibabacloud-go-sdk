// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDISyncInstanceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v int64) *GetDISyncInstanceInfoRequest
	GetFileId() *int64
	SetProjectId(v int64) *GetDISyncInstanceInfoRequest
	GetProjectId() *int64
	SetTaskType(v string) *GetDISyncInstanceInfoRequest
	GetTaskType() *string
}

type GetDISyncInstanceInfoRequest struct {
	// 	- If you set the TaskType parameter to DI_REALTIME, set the FileId parameter to the ID of the real-time synchronization task that you want to query.
	//
	// 	- If you set the TaskType parameter to DI_SOLUTION, set the FileId parameter to the ID of the data synchronization solution that you want to query.
	//
	// You can call the [ListFiles](https://help.aliyun.com/document_detail/173942.html) operation to query the ID of the real-time synchronization task or data synchronization solution.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
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

func (s GetDISyncInstanceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDISyncInstanceInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDISyncInstanceInfoRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *GetDISyncInstanceInfoRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDISyncInstanceInfoRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *GetDISyncInstanceInfoRequest) SetFileId(v int64) *GetDISyncInstanceInfoRequest {
	s.FileId = &v
	return s
}

func (s *GetDISyncInstanceInfoRequest) SetProjectId(v int64) *GetDISyncInstanceInfoRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDISyncInstanceInfoRequest) SetTaskType(v string) *GetDISyncInstanceInfoRequest {
	s.TaskType = &v
	return s
}

func (s *GetDISyncInstanceInfoRequest) Validate() error {
	return dara.Validate(s)
}
