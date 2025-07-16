// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserUsageDetailDataExportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUserUsageDetailDataExportTaskResponseBody
	GetRequestId() *string
	SetUsageDataPerPage(v *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) *DescribeUserUsageDetailDataExportTaskResponseBody
	GetUsageDataPerPage() *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage
}

type DescribeUserUsageDetailDataExportTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A91BE91F-0B34-4CBF-8E0F-A2977E15AA52
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The usage details returned per page.
	UsageDataPerPage *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage `json:"UsageDataPerPage,omitempty" xml:"UsageDataPerPage,omitempty" type:"Struct"`
}

func (s DescribeUserUsageDetailDataExportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserUsageDetailDataExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBody) GetUsageDataPerPage() *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage {
	return s.UsageDataPerPage
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBody) SetRequestId(v string) *DescribeUserUsageDetailDataExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBody) SetUsageDataPerPage(v *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) *DescribeUserUsageDetailDataExportTaskResponseBody {
	s.UsageDataPerPage = v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage struct {
	// The information about the task.
	Data *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) GetData() *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData {
	return s.Data
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) SetData(v *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage {
	s.Data = v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) SetPageNumber(v int32) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) SetPageSize(v int32) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage {
	s.PageSize = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) SetTotalCount(v int32) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage {
	s.TotalCount = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) Validate() error {
	return dara.Validate(s)
}

type DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData struct {
	DataItem []*DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem `json:"DataItem,omitempty" xml:"DataItem,omitempty" type:"Repeated"`
}

func (s DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData) GetDataItem() []*DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	return s.DataItem
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData) SetDataItem(v []*DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData {
	s.DataItem = v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData) Validate() error {
	return dara.Validate(s)
}

type DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem struct {
	// The time when the task was created.
	//
	// example:
	//
	// 2018-10-09T06:33:38Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The download URL.
	//
	// example:
	//
	// https://test.oss-cn-beijing.aliyuncs.com/billing_data/xxx
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The status of the task.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The configurations of the task.
	TaskConfig *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig `json:"TaskConfig,omitempty" xml:"TaskConfig,omitempty" type:"Struct"`
	// The ID of the task.
	//
	// example:
	//
	// 11
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
	// 2018-10-09T06:35:46Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GetStatus() *string {
	return s.Status
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GetTaskConfig() *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig {
	return s.TaskConfig
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetCreateTime(v string) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.CreateTime = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetDownloadUrl(v string) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetStatus(v string) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.Status = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetTaskConfig(v *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.TaskConfig = v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetTaskId(v string) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.TaskId = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetTaskName(v string) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.TaskName = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetUpdateTime(v string) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.UpdateTime = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) Validate() error {
	return dara.Validate(s)
}

type DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig struct {
	// The end of the time range during which data was queried.
	//
	// example:
	//
	// 2018-08-31T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2018-07-31T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) SetEndTime(v string) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig {
	s.EndTime = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) SetStartTime(v string) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig {
	s.StartTime = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) Validate() error {
	return dara.Validate(s)
}
