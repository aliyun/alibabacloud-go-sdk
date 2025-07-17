// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSyntheticTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *GetSyntheticTaskListResponseBodyPageInfo) *GetSyntheticTaskListResponseBody
	GetPageInfo() *GetSyntheticTaskListResponseBodyPageInfo
	SetRequestId(v string) *GetSyntheticTaskListResponseBody
	GetRequestId() *string
}

type GetSyntheticTaskListResponseBody struct {
	// The query results.
	PageInfo *GetSyntheticTaskListResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A5EC8221-08F2-4C95-9AF1-49FD998C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSyntheticTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskListResponseBody) GetPageInfo() *GetSyntheticTaskListResponseBodyPageInfo {
	return s.PageInfo
}

func (s *GetSyntheticTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSyntheticTaskListResponseBody) SetPageInfo(v *GetSyntheticTaskListResponseBodyPageInfo) *GetSyntheticTaskListResponseBody {
	s.PageInfo = v
	return s
}

func (s *GetSyntheticTaskListResponseBody) SetRequestId(v string) *GetSyntheticTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSyntheticTaskListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSyntheticTaskListResponseBodyPageInfo struct {
	// Indicates whether the current page is followed by a page.
	//
	// example:
	//
	// false
	HasNextPage *string `json:"HasNextPage,omitempty" xml:"HasNextPage,omitempty"`
	// Indicates whether a previous page exists.
	//
	// example:
	//
	// true
	HasPreviousPage *bool `json:"HasPreviousPage,omitempty" xml:"HasPreviousPage,omitempty"`
	// Indicates whether the page is the first page.
	//
	// example:
	//
	// true
	IsFirstPage *bool `json:"IsFirstPage,omitempty" xml:"IsFirstPage,omitempty"`
	// Indicates whether the page is the last page.
	//
	// example:
	//
	// true
	IsLastPage *bool `json:"IsLastPage,omitempty" xml:"IsLastPage,omitempty"`
	// The task information.
	List []*GetSyntheticTaskListResponseBodyPageInfoList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The first page on the navigation bar.
	//
	// example:
	//
	// 1
	NavigateFirstPage *string `json:"NavigateFirstPage,omitempty" xml:"NavigateFirstPage,omitempty"`
	// The last page on the navigation bar.
	//
	// example:
	//
	// 3
	NavigateLastPage *string `json:"NavigateLastPage,omitempty" xml:"NavigateLastPage,omitempty"`
	// All navigation page numbers.
	//
	// example:
	//
	// 1,2,3
	NavigatePageNums *string `json:"NavigatePageNums,omitempty" xml:"NavigatePageNums,omitempty"`
	// The next page.
	//
	// example:
	//
	// 3
	NextPage *string `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 10
	Pages *string `json:"Pages,omitempty" xml:"Pages,omitempty"`
	// The previous page.
	//
	// example:
	//
	// 1
	Prepage *string `json:"Prepage,omitempty" xml:"Prepage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 15
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSyntheticTaskListResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) GetHasNextPage() *string {
	return s.HasNextPage
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) GetHasPreviousPage() *bool {
	return s.HasPreviousPage
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) GetIsFirstPage() *bool {
	return s.IsFirstPage
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) GetIsLastPage() *bool {
	return s.IsLastPage
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) GetList() []*GetSyntheticTaskListResponseBodyPageInfoList {
	return s.List
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) GetNavigateFirstPage() *string {
	return s.NavigateFirstPage
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) GetNavigateLastPage() *string {
	return s.NavigateLastPage
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) GetNavigatePageNums() *string {
	return s.NavigatePageNums
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) GetNextPage() *string {
	return s.NextPage
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) GetPages() *string {
	return s.Pages
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) GetPrepage() *string {
	return s.Prepage
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) GetSize() *int64 {
	return s.Size
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) GetTotal() *int64 {
	return s.Total
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) SetHasNextPage(v string) *GetSyntheticTaskListResponseBodyPageInfo {
	s.HasNextPage = &v
	return s
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) SetHasPreviousPage(v bool) *GetSyntheticTaskListResponseBodyPageInfo {
	s.HasPreviousPage = &v
	return s
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) SetIsFirstPage(v bool) *GetSyntheticTaskListResponseBodyPageInfo {
	s.IsFirstPage = &v
	return s
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) SetIsLastPage(v bool) *GetSyntheticTaskListResponseBodyPageInfo {
	s.IsLastPage = &v
	return s
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) SetList(v []*GetSyntheticTaskListResponseBodyPageInfoList) *GetSyntheticTaskListResponseBodyPageInfo {
	s.List = v
	return s
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) SetNavigateFirstPage(v string) *GetSyntheticTaskListResponseBodyPageInfo {
	s.NavigateFirstPage = &v
	return s
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) SetNavigateLastPage(v string) *GetSyntheticTaskListResponseBodyPageInfo {
	s.NavigateLastPage = &v
	return s
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) SetNavigatePageNums(v string) *GetSyntheticTaskListResponseBodyPageInfo {
	s.NavigatePageNums = &v
	return s
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) SetNextPage(v string) *GetSyntheticTaskListResponseBodyPageInfo {
	s.NextPage = &v
	return s
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) SetPages(v string) *GetSyntheticTaskListResponseBodyPageInfo {
	s.Pages = &v
	return s
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) SetPrepage(v string) *GetSyntheticTaskListResponseBodyPageInfo {
	s.Prepage = &v
	return s
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) SetSize(v int64) *GetSyntheticTaskListResponseBodyPageInfo {
	s.Size = &v
	return s
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) SetTotal(v int64) *GetSyntheticTaskListResponseBodyPageInfo {
	s.Total = &v
	return s
}

func (s *GetSyntheticTaskListResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type GetSyntheticTaskListResponseBodyPageInfoList struct {
	// The time when the task was created.
	//
	// example:
	//
	// 1634005438000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The number of detection points.
	//
	// example:
	//
	// 2
	MonitorNumber *int64 `json:"MonitorNumber,omitempty" xml:"MonitorNumber,omitempty"`
	// The ID of the synthetic monitoring task.
	//
	// example:
	//
	// 2118709
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task name.
	//
	// example:
	//
	// net-test
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **0**: The task is stopped.
	//
	// 	- **1**: The task is started.
	//
	// 	- **9**: The task is ended.
	//
	// example:
	//
	// 0
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The type of the task. Valid values:
	//
	// 1.  3: web page performance - IE
	//
	// 2.  34: web page performance - Chrome
	//
	// 3.  0: network quality
	//
	// 4.  40: file download
	//
	// 5.  7: API performance
	//
	// example:
	//
	// 0
	TaskType *int64 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The name of the task type.
	TaskTypeName *string `json:"TaskTypeName,omitempty" xml:"TaskTypeName,omitempty"`
	// The URL for synthetic monitoring.
	//
	// example:
	//
	// www.example.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The availability. Only the data of the last day is counted. If no data is available for the last day, an empty value is returned.
	//
	// example:
	//
	// 0.80
	Usable *float32 `json:"Usable,omitempty" xml:"Usable,omitempty"`
}

func (s GetSyntheticTaskListResponseBodyPageInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskListResponseBodyPageInfoList) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskListResponseBodyPageInfoList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetSyntheticTaskListResponseBodyPageInfoList) GetMonitorNumber() *int64 {
	return s.MonitorNumber
}

func (s *GetSyntheticTaskListResponseBodyPageInfoList) GetTaskId() *string {
	return s.TaskId
}

func (s *GetSyntheticTaskListResponseBodyPageInfoList) GetTaskName() *string {
	return s.TaskName
}

func (s *GetSyntheticTaskListResponseBodyPageInfoList) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetSyntheticTaskListResponseBodyPageInfoList) GetTaskType() *int64 {
	return s.TaskType
}

func (s *GetSyntheticTaskListResponseBodyPageInfoList) GetTaskTypeName() *string {
	return s.TaskTypeName
}

func (s *GetSyntheticTaskListResponseBodyPageInfoList) GetUrl() *string {
	return s.Url
}

func (s *GetSyntheticTaskListResponseBodyPageInfoList) GetUsable() *float32 {
	return s.Usable
}

func (s *GetSyntheticTaskListResponseBodyPageInfoList) SetCreateTime(v string) *GetSyntheticTaskListResponseBodyPageInfoList {
	s.CreateTime = &v
	return s
}

func (s *GetSyntheticTaskListResponseBodyPageInfoList) SetMonitorNumber(v int64) *GetSyntheticTaskListResponseBodyPageInfoList {
	s.MonitorNumber = &v
	return s
}

func (s *GetSyntheticTaskListResponseBodyPageInfoList) SetTaskId(v string) *GetSyntheticTaskListResponseBodyPageInfoList {
	s.TaskId = &v
	return s
}

func (s *GetSyntheticTaskListResponseBodyPageInfoList) SetTaskName(v string) *GetSyntheticTaskListResponseBodyPageInfoList {
	s.TaskName = &v
	return s
}

func (s *GetSyntheticTaskListResponseBodyPageInfoList) SetTaskStatus(v string) *GetSyntheticTaskListResponseBodyPageInfoList {
	s.TaskStatus = &v
	return s
}

func (s *GetSyntheticTaskListResponseBodyPageInfoList) SetTaskType(v int64) *GetSyntheticTaskListResponseBodyPageInfoList {
	s.TaskType = &v
	return s
}

func (s *GetSyntheticTaskListResponseBodyPageInfoList) SetTaskTypeName(v string) *GetSyntheticTaskListResponseBodyPageInfoList {
	s.TaskTypeName = &v
	return s
}

func (s *GetSyntheticTaskListResponseBodyPageInfoList) SetUrl(v string) *GetSyntheticTaskListResponseBodyPageInfoList {
	s.Url = &v
	return s
}

func (s *GetSyntheticTaskListResponseBodyPageInfoList) SetUsable(v float32) *GetSyntheticTaskListResponseBodyPageInfoList {
	s.Usable = &v
	return s
}

func (s *GetSyntheticTaskListResponseBodyPageInfoList) Validate() error {
	return dara.Validate(s)
}
