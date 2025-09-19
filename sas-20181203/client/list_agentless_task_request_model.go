// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentlessTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListAgentlessTaskRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *ListAgentlessTaskRequest
	GetEndTime() *int64
	SetInternetIp(v string) *ListAgentlessTaskRequest
	GetInternetIp() *string
	SetIntranetIp(v string) *ListAgentlessTaskRequest
	GetIntranetIp() *string
	SetLang(v string) *ListAgentlessTaskRequest
	GetLang() *string
	SetMachineName(v string) *ListAgentlessTaskRequest
	GetMachineName() *string
	SetPageSize(v int32) *ListAgentlessTaskRequest
	GetPageSize() *int32
	SetRootTask(v bool) *ListAgentlessTaskRequest
	GetRootTask() *bool
	SetRootTaskId(v string) *ListAgentlessTaskRequest
	GetRootTaskId() *string
	SetStartTime(v int64) *ListAgentlessTaskRequest
	GetStartTime() *int64
	SetStatus(v int32) *ListAgentlessTaskRequest
	GetStatus() *int32
	SetTargetName(v string) *ListAgentlessTaskRequest
	GetTargetName() *string
	SetTargetType(v int32) *ListAgentlessTaskRequest
	GetTargetType() *int32
	SetTaskId(v string) *ListAgentlessTaskRequest
	GetTaskId() *string
	SetUuid(v string) *ListAgentlessTaskRequest
	GetUuid() *string
}

type ListAgentlessTaskRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end timestamp of the task.
	//
	// example:
	//
	// 1635575219000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The public IP address of the asset that you want to query.
	//
	// example:
	//
	// 1.1.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the asset that you want to query.
	//
	// example:
	//
	// 172.26.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The language type. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
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
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to query main tasks. Valid values:
	//
	// 	- **true**: queries main tasks.
	//
	// 	- **false**: queries subtasks.
	//
	// example:
	//
	// false
	RootTask *bool `json:"RootTask,omitempty" xml:"RootTask,omitempty"`
	// The ID of the main task.
	//
	// example:
	//
	// 12c27343861610c5db3f7a2573b4****
	RootTaskId *string `json:"RootTaskId,omitempty" xml:"RootTaskId,omitempty"`
	// The start timestamp of the task.
	//
	// example:
	//
	// 1651290987000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the detection task.
	//
	// 	- **1**: The detection task is in progress.
	//
	// 	- **2**: The detection task is complete.
	//
	// 	- **3**: The detection task fails.
	//
	// 	- **4**: The detection task times out.
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the asset that you want to detect.
	//
	// example:
	//
	// source-test-obj-0****
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The type of the asset that you want to detect. Valid values:
	//
	// 	- **1**: snapshot
	//
	// 	- **2**: image
	//
	// example:
	//
	// 1
	TargetType *int32 `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The ID of the main task. If you want to query subtasks of a main task, you must specify this parameter.
	//
	// example:
	//
	// d7b2acf8d362742123e4a84e1bf8****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// e4af3620-6895-4e2f-a641-a9d8fb53****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListAgentlessTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessTaskRequest) GoString() string {
	return s.String()
}

func (s *ListAgentlessTaskRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAgentlessTaskRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListAgentlessTaskRequest) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListAgentlessTaskRequest) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ListAgentlessTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *ListAgentlessTaskRequest) GetMachineName() *string {
	return s.MachineName
}

func (s *ListAgentlessTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAgentlessTaskRequest) GetRootTask() *bool {
	return s.RootTask
}

func (s *ListAgentlessTaskRequest) GetRootTaskId() *string {
	return s.RootTaskId
}

func (s *ListAgentlessTaskRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListAgentlessTaskRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListAgentlessTaskRequest) GetTargetName() *string {
	return s.TargetName
}

func (s *ListAgentlessTaskRequest) GetTargetType() *int32 {
	return s.TargetType
}

func (s *ListAgentlessTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListAgentlessTaskRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ListAgentlessTaskRequest) SetCurrentPage(v int32) *ListAgentlessTaskRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAgentlessTaskRequest) SetEndTime(v int64) *ListAgentlessTaskRequest {
	s.EndTime = &v
	return s
}

func (s *ListAgentlessTaskRequest) SetInternetIp(v string) *ListAgentlessTaskRequest {
	s.InternetIp = &v
	return s
}

func (s *ListAgentlessTaskRequest) SetIntranetIp(v string) *ListAgentlessTaskRequest {
	s.IntranetIp = &v
	return s
}

func (s *ListAgentlessTaskRequest) SetLang(v string) *ListAgentlessTaskRequest {
	s.Lang = &v
	return s
}

func (s *ListAgentlessTaskRequest) SetMachineName(v string) *ListAgentlessTaskRequest {
	s.MachineName = &v
	return s
}

func (s *ListAgentlessTaskRequest) SetPageSize(v int32) *ListAgentlessTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListAgentlessTaskRequest) SetRootTask(v bool) *ListAgentlessTaskRequest {
	s.RootTask = &v
	return s
}

func (s *ListAgentlessTaskRequest) SetRootTaskId(v string) *ListAgentlessTaskRequest {
	s.RootTaskId = &v
	return s
}

func (s *ListAgentlessTaskRequest) SetStartTime(v int64) *ListAgentlessTaskRequest {
	s.StartTime = &v
	return s
}

func (s *ListAgentlessTaskRequest) SetStatus(v int32) *ListAgentlessTaskRequest {
	s.Status = &v
	return s
}

func (s *ListAgentlessTaskRequest) SetTargetName(v string) *ListAgentlessTaskRequest {
	s.TargetName = &v
	return s
}

func (s *ListAgentlessTaskRequest) SetTargetType(v int32) *ListAgentlessTaskRequest {
	s.TargetType = &v
	return s
}

func (s *ListAgentlessTaskRequest) SetTaskId(v string) *ListAgentlessTaskRequest {
	s.TaskId = &v
	return s
}

func (s *ListAgentlessTaskRequest) SetUuid(v string) *ListAgentlessTaskRequest {
	s.Uuid = &v
	return s
}

func (s *ListAgentlessTaskRequest) Validate() error {
	return dara.Validate(s)
}
