// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodUserUsageDetailDataExportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVodUserUsageDetailDataExportTaskResponseBody
	GetRequestId() *string
	SetUsageDataPerPage(v *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) *DescribeVodUserUsageDetailDataExportTaskResponseBody
	GetUsageDataPerPage() *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage
}

type DescribeVodUserUsageDetailDataExportTaskResponseBody struct {
	RequestId        *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UsageDataPerPage *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage `json:"UsageDataPerPage,omitempty" xml:"UsageDataPerPage,omitempty" type:"Struct"`
}

func (s DescribeVodUserUsageDetailDataExportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserUsageDetailDataExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBody) GetUsageDataPerPage() *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage {
	return s.UsageDataPerPage
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBody) SetRequestId(v string) *DescribeVodUserUsageDetailDataExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBody) SetUsageDataPerPage(v *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) *DescribeVodUserUsageDetailDataExportTaskResponseBody {
	s.UsageDataPerPage = v
	return s
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage struct {
	Data       *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	PageNumber *int32                                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) GoString() string {
	return s.String()
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) GetData() *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData {
	return s.Data
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) SetData(v *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData) *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage {
	s.Data = v
	return s
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) SetPageNumber(v int32) *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage {
	s.PageNumber = &v
	return s
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) SetPageSize(v int32) *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage {
	s.PageSize = &v
	return s
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) SetTotalCount(v int32) *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage {
	s.TotalCount = &v
	return s
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) Validate() error {
	return dara.Validate(s)
}

type DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData struct {
	DataItem []*DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem `json:"DataItem,omitempty" xml:"DataItem,omitempty" type:"Repeated"`
}

func (s DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData) GoString() string {
	return s.String()
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData) GetDataItem() []*DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	return s.DataItem
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData) SetDataItem(v []*DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData {
	s.DataItem = v
	return s
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData) Validate() error {
	return dara.Validate(s)
}

type DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem struct {
	CreateTime  *string                                                                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DownloadUrl *string                                                                                     `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	Status      *string                                                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskConfig  *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig `json:"TaskConfig,omitempty" xml:"TaskConfig,omitempty" type:"Struct"`
	TaskId      *string                                                                                     `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName    *string                                                                                     `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	UpdateTime  *string                                                                                     `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GoString() string {
	return s.String()
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GetStatus() *string {
	return s.Status
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GetTaskConfig() *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig {
	return s.TaskConfig
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetCreateTime(v string) *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.CreateTime = &v
	return s
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetDownloadUrl(v string) *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetStatus(v string) *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.Status = &v
	return s
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetTaskConfig(v *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.TaskConfig = v
	return s
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetTaskId(v string) *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.TaskId = &v
	return s
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetTaskName(v string) *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.TaskName = &v
	return s
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetUpdateTime(v string) *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.UpdateTime = &v
	return s
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) Validate() error {
	return dara.Validate(s)
}

type DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) GoString() string {
	return s.String()
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) SetEndTime(v string) *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig {
	s.EndTime = &v
	return s
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) SetStartTime(v string) *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig {
	s.StartTime = &v
	return s
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) Validate() error {
	return dara.Validate(s)
}
