// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusScanTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListVirusScanTaskRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *ListVirusScanTaskRequest
	GetEndTime() *int64
	SetInternetIp(v string) *ListVirusScanTaskRequest
	GetInternetIp() *string
	SetIntranetIp(v string) *ListVirusScanTaskRequest
	GetIntranetIp() *string
	SetLang(v string) *ListVirusScanTaskRequest
	GetLang() *string
	SetMachineName(v string) *ListVirusScanTaskRequest
	GetMachineName() *string
	SetPageSize(v int32) *ListVirusScanTaskRequest
	GetPageSize() *int32
	SetRootTask(v bool) *ListVirusScanTaskRequest
	GetRootTask() *bool
	SetRootTaskId(v string) *ListVirusScanTaskRequest
	GetRootTaskId() *string
	SetScanType(v string) *ListVirusScanTaskRequest
	GetScanType() *string
	SetStartTime(v int64) *ListVirusScanTaskRequest
	GetStartTime() *int64
	SetStatus(v int32) *ListVirusScanTaskRequest
	GetStatus() *int32
	SetStatusList(v []*int32) *ListVirusScanTaskRequest
	GetStatusList() []*int32
	SetTaskId(v string) *ListVirusScanTaskRequest
	GetTaskId() *string
}

type ListVirusScanTaskRequest struct {
	// The page number of the results to return. Default value: **1**, which indicates that results start from page 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The timestamp of the task end time to query, in milliseconds.
	//
	// example:
	//
	// 1680919232999
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The public IP address.
	//
	// example:
	//
	// 120.27.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address.
	//
	// example:
	//
	// 172.26.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The language type of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// oracle-win-001****
	MachineName *string `json:"MachineName,omitempty" xml:"MachineName,omitempty"`
	// The number of tasks per page in a paged query. Default value: **20**, which indicates that each page contains 20 tasks.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether the task is the root task of the virus scan.
	//
	// example:
	//
	// true
	RootTask *bool `json:"RootTask,omitempty" xml:"RootTask,omitempty"`
	// The root task ID.
	//
	// > Call [GetVirusScanLatestTaskStatistic](~~GetVirusScanLatestTaskStatistic~~) to obtain this parameter.
	//
	// example:
	//
	// 89f5d7813bd59dd237580a8664b3xxxx
	RootTaskId *string `json:"RootTaskId,omitempty" xml:"RootTaskId,omitempty"`
	// The scan type of the virus scan task to query. Valid values:
	//
	// - **system**: automatic system scan
	//
	// - **user**: custom user scan.
	//
	// example:
	//
	// user
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// The timestamp of the task start time to query, in milliseconds.
	//
	// example:
	//
	// 1680919232000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The execution status of the virus scan task. Valid values:
	//
	// - **1**: Scanning.
	//
	// - **2**: Completed.
	//
	// - **3**: Failed.
	//
	// - **4**: Timed out.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of statuses used to filter tasks by multiple statuses.
	StatusList []*int32 `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	// The ID of the virus scan task to query.
	//
	// > Call [ListVirusScanTask](~~ListVirusScanTask~~) to obtain this parameter.
	//
	// example:
	//
	// 1471d8ebb96795b41ede090b9758****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListVirusScanTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanTaskRequest) GoString() string {
	return s.String()
}

func (s *ListVirusScanTaskRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListVirusScanTaskRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListVirusScanTaskRequest) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListVirusScanTaskRequest) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ListVirusScanTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *ListVirusScanTaskRequest) GetMachineName() *string {
	return s.MachineName
}

func (s *ListVirusScanTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVirusScanTaskRequest) GetRootTask() *bool {
	return s.RootTask
}

func (s *ListVirusScanTaskRequest) GetRootTaskId() *string {
	return s.RootTaskId
}

func (s *ListVirusScanTaskRequest) GetScanType() *string {
	return s.ScanType
}

func (s *ListVirusScanTaskRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListVirusScanTaskRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListVirusScanTaskRequest) GetStatusList() []*int32 {
	return s.StatusList
}

func (s *ListVirusScanTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListVirusScanTaskRequest) SetCurrentPage(v int32) *ListVirusScanTaskRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListVirusScanTaskRequest) SetEndTime(v int64) *ListVirusScanTaskRequest {
	s.EndTime = &v
	return s
}

func (s *ListVirusScanTaskRequest) SetInternetIp(v string) *ListVirusScanTaskRequest {
	s.InternetIp = &v
	return s
}

func (s *ListVirusScanTaskRequest) SetIntranetIp(v string) *ListVirusScanTaskRequest {
	s.IntranetIp = &v
	return s
}

func (s *ListVirusScanTaskRequest) SetLang(v string) *ListVirusScanTaskRequest {
	s.Lang = &v
	return s
}

func (s *ListVirusScanTaskRequest) SetMachineName(v string) *ListVirusScanTaskRequest {
	s.MachineName = &v
	return s
}

func (s *ListVirusScanTaskRequest) SetPageSize(v int32) *ListVirusScanTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListVirusScanTaskRequest) SetRootTask(v bool) *ListVirusScanTaskRequest {
	s.RootTask = &v
	return s
}

func (s *ListVirusScanTaskRequest) SetRootTaskId(v string) *ListVirusScanTaskRequest {
	s.RootTaskId = &v
	return s
}

func (s *ListVirusScanTaskRequest) SetScanType(v string) *ListVirusScanTaskRequest {
	s.ScanType = &v
	return s
}

func (s *ListVirusScanTaskRequest) SetStartTime(v int64) *ListVirusScanTaskRequest {
	s.StartTime = &v
	return s
}

func (s *ListVirusScanTaskRequest) SetStatus(v int32) *ListVirusScanTaskRequest {
	s.Status = &v
	return s
}

func (s *ListVirusScanTaskRequest) SetStatusList(v []*int32) *ListVirusScanTaskRequest {
	s.StatusList = v
	return s
}

func (s *ListVirusScanTaskRequest) SetTaskId(v string) *ListVirusScanTaskRequest {
	s.TaskId = &v
	return s
}

func (s *ListVirusScanTaskRequest) Validate() error {
	return dara.Validate(s)
}
