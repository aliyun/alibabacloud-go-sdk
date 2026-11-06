// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExportTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeExportTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeExportTasksRequest
	GetPageSize() *int32
	SetTaskId(v string) *DescribeExportTasksRequest
	GetTaskId() *string
	SetTaskType(v string) *DescribeExportTasksRequest
	GetTaskType() *string
}

type DescribeExportTasksRequest struct {
	// The page number for paging. Pages start from 1. If this parameter is not specified or is set to a value less than 1, the first page is returned. The query results are sorted in descending order of task creation time, with the most recently created tasks listed first. When polling a single task by TaskId, explicitly set this parameter to 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of tasks to return per page. Default value: 200. If this parameter is not specified, the default value is used. If a value less than 1 is specified, 10 is used. When polling a single task by TaskId, set this parameter to 1 to avoid returning redundant tasks.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Filters tasks by export task ID. The value is the TaskId returned by CreateExportTask, which is a 32-character UUID string without hyphens. This parameter requires exact match and does not support fuzzy queries. If this parameter is not specified, all tasks that meet other conditions under the current account are returned.
	//
	// example:
	//
	// 8119078c29e948dba2e0a9917995332a
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Filters tasks by export task type. The value is case-sensitive and must be one of the following valid values:
	//
	// - WUYING_SERVER_LIST: Wuying host list.
	//
	// - USER_USAGE_REPORT: User usage report.
	//
	// - BROWSER_BOOKMARK_LIST: Cloud browser bookmark list.
	//
	// - BROWSER_RESTRICTED_URL_LIST: Cloud browser URL whitelist.
	//
	// If this parameter is not specified, tasks are not filtered by type, and all tasks that meet other conditions under the current account are returned. If an unrecognized value is specified, the server does not return a parameter error but cancels type filtering. Therefore, you must specify one of the preceding exact enum values. When querying a single task by TaskId, omit this parameter.
	//
	// example:
	//
	// WUYING_SERVER_LIST
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeExportTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExportTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeExportTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeExportTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeExportTasksRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeExportTasksRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeExportTasksRequest) SetPageNumber(v int32) *DescribeExportTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeExportTasksRequest) SetPageSize(v int32) *DescribeExportTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeExportTasksRequest) SetTaskId(v string) *DescribeExportTasksRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeExportTasksRequest) SetTaskType(v string) *DescribeExportTasksRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeExportTasksRequest) Validate() error {
	return dara.Validate(s)
}
