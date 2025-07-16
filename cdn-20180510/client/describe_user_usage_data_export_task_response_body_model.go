// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserUsageDataExportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUserUsageDataExportTaskResponseBody
	GetRequestId() *string
	SetUsageDataPerPage(v *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage) *DescribeUserUsageDataExportTaskResponseBody
	GetUsageDataPerPage() *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage
}

type DescribeUserUsageDataExportTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A91BE91F-0B34-4CBF-8E0F-A2977E15AA52
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The usage details returned per page.
	UsageDataPerPage *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage `json:"UsageDataPerPage,omitempty" xml:"UsageDataPerPage,omitempty" type:"Struct"`
}

func (s DescribeUserUsageDataExportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserUsageDataExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDataExportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserUsageDataExportTaskResponseBody) GetUsageDataPerPage() *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage {
	return s.UsageDataPerPage
}

func (s *DescribeUserUsageDataExportTaskResponseBody) SetRequestId(v string) *DescribeUserUsageDataExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBody) SetUsageDataPerPage(v *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage) *DescribeUserUsageDataExportTaskResponseBody {
	s.UsageDataPerPage = v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage struct {
	// The information about the tasks.
	Data *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage) GetData() *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageData {
	return s.Data
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage) SetData(v *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageData) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage {
	s.Data = v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage) SetPageNumber(v int32) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage) SetPageSize(v int32) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage {
	s.PageSize = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage) SetTotalCount(v int32) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage {
	s.TotalCount = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage) Validate() error {
	return dara.Validate(s)
}

type DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageData struct {
	DataItem []*DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem `json:"DataItem,omitempty" xml:"DataItem,omitempty" type:"Repeated"`
}

func (s DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageData) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageData) GetDataItem() []*DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	return s.DataItem
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageData) SetDataItem(v []*DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageData {
	s.DataItem = v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageData) Validate() error {
	return dara.Validate(s)
}

type DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem struct {
	// The time when the task was created.
	//
	// example:
	//
	// 2019-12-31T08:43:21Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The download URL.
	//
	// example:
	//
	// https://cdn-polaris.xxxx
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The state of the task.
	//
	// 	- created: The task is being created.
	//
	// 	- success: The task is successful.
	//
	// 	- failed: The task failed.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The configurations of the task.
	TaskConfig *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig `json:"TaskConfig,omitempty" xml:"TaskConfig,omitempty" type:"Struct"`
	// The ID of the task.
	//
	// example:
	//
	// A91BE91F-0B34-4CBF-8E0F-A2977
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// Refresh
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The time when the task was last modified.
	//
	// example:
	//
	// 2019-12-31T08:45:02Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GetStatus() *string {
	return s.Status
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GetTaskConfig() *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig {
	return s.TaskConfig
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetCreateTime(v string) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.CreateTime = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetDownloadUrl(v string) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetStatus(v string) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.Status = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetTaskConfig(v *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.TaskConfig = v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetTaskId(v string) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.TaskId = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetTaskName(v string) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.TaskName = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetUpdateTime(v string) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.UpdateTime = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) Validate() error {
	return dara.Validate(s)
}

type DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig struct {
	// The end of the time range that was queried.
	//
	// example:
	//
	// 2019-12-30T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2019-12-29T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) SetEndTime(v string) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig {
	s.EndTime = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) SetStartTime(v string) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig {
	s.StartTime = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) Validate() error {
	return dara.Validate(s)
}
