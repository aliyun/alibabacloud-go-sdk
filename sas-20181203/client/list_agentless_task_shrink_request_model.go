// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentlessTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListAgentlessTaskShrinkRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *ListAgentlessTaskShrinkRequest
	GetEndTime() *int64
	SetInternetIp(v string) *ListAgentlessTaskShrinkRequest
	GetInternetIp() *string
	SetIntranetIp(v string) *ListAgentlessTaskShrinkRequest
	GetIntranetIp() *string
	SetLang(v string) *ListAgentlessTaskShrinkRequest
	GetLang() *string
	SetMachineName(v string) *ListAgentlessTaskShrinkRequest
	GetMachineName() *string
	SetPageSize(v int32) *ListAgentlessTaskShrinkRequest
	GetPageSize() *int32
	SetRootTask(v bool) *ListAgentlessTaskShrinkRequest
	GetRootTask() *bool
	SetRootTaskId(v string) *ListAgentlessTaskShrinkRequest
	GetRootTaskId() *string
	SetStartTime(v int64) *ListAgentlessTaskShrinkRequest
	GetStartTime() *int64
	SetStatus(v int32) *ListAgentlessTaskShrinkRequest
	GetStatus() *int32
	SetTargetName(v string) *ListAgentlessTaskShrinkRequest
	GetTargetName() *string
	SetTargetType(v int32) *ListAgentlessTaskShrinkRequest
	GetTargetType() *int32
	SetTaskId(v string) *ListAgentlessTaskShrinkRequest
	GetTaskId() *string
	SetTaskIdListShrink(v string) *ListAgentlessTaskShrinkRequest
	GetTaskIdListShrink() *string
	SetUuid(v string) *ListAgentlessTaskShrinkRequest
	GetUuid() *string
}

type ListAgentlessTaskShrinkRequest struct {
	// The page number of the current page in a paging query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The timestamp of the end time.
	//
	// example:
	//
	// 1635575219000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The public IP address of the asset to query.
	//
	// example:
	//
	// 1.1.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the asset to query.
	//
	// example:
	//
	// 172.26.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The language type. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// oracle-win-001****
	MachineName *string `json:"MachineName,omitempty" xml:"MachineName,omitempty"`
	// The maximum number of entries to return per page in a paging query.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to query the root task list. Valid values:
	//
	// - **true**: Queries root tasks.
	//
	// - **false**: Queries subtasks.
	//
	// example:
	//
	// false
	RootTask *bool `json:"RootTask,omitempty" xml:"RootTask,omitempty"`
	// The ID of the root task.
	//
	// example:
	//
	// 12c27343861610c5db3f7a2573b4****
	RootTaskId *string `json:"RootTaskId,omitempty" xml:"RootTaskId,omitempty"`
	// The timestamp of the start time.
	//
	// example:
	//
	// 1651290987000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the detection task. Valid values:
	//
	// - **1**: running
	//
	// - **2**: completed
	//
	// - **3**: failed
	//
	// - **4**: timed out
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the scan target.
	//
	// example:
	//
	// source-test-obj-0****
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The object type of the scan. Valid values:
	//
	// - **1**: snapshot
	//
	// - **2**: image
	//
	// example:
	//
	// 1
	TargetType *int32 `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The ID of the root task. Specify this parameter when you query the list of subtasks under a root task.
	//
	// example:
	//
	// d7b2acf8d362742123e4a84e1bf8****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The list of task IDs to return. You can specify up to 100 IDs. You must specify RootTask and cannot specify this parameter together with TaskId. If RootTask is set to true, root tasks are queried. If RootTask is set to false, subtasks are queried, and cross-root task queries are allowed. If RootTaskId is specified, the intersection is returned.
	TaskIdListShrink *string `json:"TaskIdList,omitempty" xml:"TaskIdList,omitempty"`
	// The UUID of the server to query.
	//
	// example:
	//
	// e4af3620-6895-4e2f-a641-a9d8fb53****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListAgentlessTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAgentlessTaskShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAgentlessTaskShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListAgentlessTaskShrinkRequest) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListAgentlessTaskShrinkRequest) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ListAgentlessTaskShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *ListAgentlessTaskShrinkRequest) GetMachineName() *string {
	return s.MachineName
}

func (s *ListAgentlessTaskShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAgentlessTaskShrinkRequest) GetRootTask() *bool {
	return s.RootTask
}

func (s *ListAgentlessTaskShrinkRequest) GetRootTaskId() *string {
	return s.RootTaskId
}

func (s *ListAgentlessTaskShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListAgentlessTaskShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListAgentlessTaskShrinkRequest) GetTargetName() *string {
	return s.TargetName
}

func (s *ListAgentlessTaskShrinkRequest) GetTargetType() *int32 {
	return s.TargetType
}

func (s *ListAgentlessTaskShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListAgentlessTaskShrinkRequest) GetTaskIdListShrink() *string {
	return s.TaskIdListShrink
}

func (s *ListAgentlessTaskShrinkRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ListAgentlessTaskShrinkRequest) SetCurrentPage(v int32) *ListAgentlessTaskShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAgentlessTaskShrinkRequest) SetEndTime(v int64) *ListAgentlessTaskShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListAgentlessTaskShrinkRequest) SetInternetIp(v string) *ListAgentlessTaskShrinkRequest {
	s.InternetIp = &v
	return s
}

func (s *ListAgentlessTaskShrinkRequest) SetIntranetIp(v string) *ListAgentlessTaskShrinkRequest {
	s.IntranetIp = &v
	return s
}

func (s *ListAgentlessTaskShrinkRequest) SetLang(v string) *ListAgentlessTaskShrinkRequest {
	s.Lang = &v
	return s
}

func (s *ListAgentlessTaskShrinkRequest) SetMachineName(v string) *ListAgentlessTaskShrinkRequest {
	s.MachineName = &v
	return s
}

func (s *ListAgentlessTaskShrinkRequest) SetPageSize(v int32) *ListAgentlessTaskShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListAgentlessTaskShrinkRequest) SetRootTask(v bool) *ListAgentlessTaskShrinkRequest {
	s.RootTask = &v
	return s
}

func (s *ListAgentlessTaskShrinkRequest) SetRootTaskId(v string) *ListAgentlessTaskShrinkRequest {
	s.RootTaskId = &v
	return s
}

func (s *ListAgentlessTaskShrinkRequest) SetStartTime(v int64) *ListAgentlessTaskShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListAgentlessTaskShrinkRequest) SetStatus(v int32) *ListAgentlessTaskShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListAgentlessTaskShrinkRequest) SetTargetName(v string) *ListAgentlessTaskShrinkRequest {
	s.TargetName = &v
	return s
}

func (s *ListAgentlessTaskShrinkRequest) SetTargetType(v int32) *ListAgentlessTaskShrinkRequest {
	s.TargetType = &v
	return s
}

func (s *ListAgentlessTaskShrinkRequest) SetTaskId(v string) *ListAgentlessTaskShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *ListAgentlessTaskShrinkRequest) SetTaskIdListShrink(v string) *ListAgentlessTaskShrinkRequest {
	s.TaskIdListShrink = &v
	return s
}

func (s *ListAgentlessTaskShrinkRequest) SetUuid(v string) *ListAgentlessTaskShrinkRequest {
	s.Uuid = &v
	return s
}

func (s *ListAgentlessTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
