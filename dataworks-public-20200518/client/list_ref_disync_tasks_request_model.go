// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRefDISyncTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasourceName(v string) *ListRefDISyncTasksRequest
	GetDatasourceName() *string
	SetPageNumber(v int64) *ListRefDISyncTasksRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListRefDISyncTasksRequest
	GetPageSize() *int64
	SetProjectId(v int64) *ListRefDISyncTasksRequest
	GetProjectId() *int64
	SetRefType(v string) *ListRefDISyncTasksRequest
	GetRefType() *string
	SetTaskType(v string) *ListRefDISyncTasksRequest
	GetTaskType() *string
}

type ListRefDISyncTasksRequest struct {
	// The name of the data source. You can call the [ListDataSources](https://help.aliyun.com/document_detail/211431.html) operation to query the name of the data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql_datasource
	DatasourceName *string `json:"DatasourceName,omitempty" xml:"DatasourceName,omitempty"`
	// The page number. Valid values: 1 to 100.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 100
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The condition used to filter synchronization tasks. Valid values:
	//
	// 	- from: queries the synchronization tasks that use the data source as the source.
	//
	// 	- to: queries the synchronization tasks that use the data source as the destination.
	//
	// This parameter is required.
	//
	// example:
	//
	// from
	RefType *string `json:"RefType,omitempty" xml:"RefType,omitempty"`
	// The type of the synchronization task that you want to query. Valid values:
	//
	// 	- DI_OFFLINE: batch synchronization task
	//
	// 	- DI_REALTIME: real-time synchronization task
	//
	// You can call the ListRefDISyncTasks operation to query only one type of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// DI_OFFLINE
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListRefDISyncTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRefDISyncTasksRequest) GoString() string {
	return s.String()
}

func (s *ListRefDISyncTasksRequest) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *ListRefDISyncTasksRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListRefDISyncTasksRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRefDISyncTasksRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListRefDISyncTasksRequest) GetRefType() *string {
	return s.RefType
}

func (s *ListRefDISyncTasksRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *ListRefDISyncTasksRequest) SetDatasourceName(v string) *ListRefDISyncTasksRequest {
	s.DatasourceName = &v
	return s
}

func (s *ListRefDISyncTasksRequest) SetPageNumber(v int64) *ListRefDISyncTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRefDISyncTasksRequest) SetPageSize(v int64) *ListRefDISyncTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListRefDISyncTasksRequest) SetProjectId(v int64) *ListRefDISyncTasksRequest {
	s.ProjectId = &v
	return s
}

func (s *ListRefDISyncTasksRequest) SetRefType(v string) *ListRefDISyncTasksRequest {
	s.RefType = &v
	return s
}

func (s *ListRefDISyncTasksRequest) SetTaskType(v string) *ListRefDISyncTasksRequest {
	s.TaskType = &v
	return s
}

func (s *ListRefDISyncTasksRequest) Validate() error {
	return dara.Validate(s)
}
