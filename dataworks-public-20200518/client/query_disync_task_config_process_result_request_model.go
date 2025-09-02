// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDISyncTaskConfigProcessResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncProcessId(v int64) *QueryDISyncTaskConfigProcessResultRequest
	GetAsyncProcessId() *int64
	SetProjectId(v int64) *QueryDISyncTaskConfigProcessResultRequest
	GetProjectId() *int64
	SetTaskType(v string) *QueryDISyncTaskConfigProcessResultRequest
	GetTaskType() *string
}

type QueryDISyncTaskConfigProcessResultRequest struct {
	// The asynchronous thread ID. You can call the [GenerateDISyncTaskConfigForCreating](https://help.aliyun.com/document_detail/383463.html) or [GenerateDISyncTaskConfigForUpdating](https://help.aliyun.com/document_detail/383464.html) operation to obtain the ID.
	//
	// 	- The GenerateDISyncTaskConfigForCreating operation is used to generate the ID of the asynchronous thread that is used to create a real-time synchronization task in Data Integration.
	//
	// 	- The GenerateDISyncTaskConfigForUpdating operation is used to generate the ID of the asynchronous thread that is used to update a real-time synchronization task in Data Integration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	AsyncProcessId *int64 `json:"AsyncProcessId,omitempty" xml:"AsyncProcessId,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The type of the object that you want to create or update in Data Integration in asynchronous mode. Valid values:
	//
	// 	- DI_REALTIME: real-time synchronization task
	//
	// 	- DI_SOLUTION: synchronization solution DataWorks allows you to create or update real-time synchronization tasks and synchronization solutions in Data Integration only in asynchronous mode.
	//
	// Valid values:
	//
	// 	- DI_OFFLINE
	//
	// 	- DI_REALTIME
	//
	// 	- DI_SOLUTION
	//
	// This parameter is required.
	//
	// example:
	//
	// DI_REALTIME
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s QueryDISyncTaskConfigProcessResultRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDISyncTaskConfigProcessResultRequest) GoString() string {
	return s.String()
}

func (s *QueryDISyncTaskConfigProcessResultRequest) GetAsyncProcessId() *int64 {
	return s.AsyncProcessId
}

func (s *QueryDISyncTaskConfigProcessResultRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *QueryDISyncTaskConfigProcessResultRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *QueryDISyncTaskConfigProcessResultRequest) SetAsyncProcessId(v int64) *QueryDISyncTaskConfigProcessResultRequest {
	s.AsyncProcessId = &v
	return s
}

func (s *QueryDISyncTaskConfigProcessResultRequest) SetProjectId(v int64) *QueryDISyncTaskConfigProcessResultRequest {
	s.ProjectId = &v
	return s
}

func (s *QueryDISyncTaskConfigProcessResultRequest) SetTaskType(v string) *QueryDISyncTaskConfigProcessResultRequest {
	s.TaskType = &v
	return s
}

func (s *QueryDISyncTaskConfigProcessResultRequest) Validate() error {
	return dara.Validate(s)
}
