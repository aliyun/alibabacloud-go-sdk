// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrawlerRunsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListCrawlerRunsResponseBodyPagingInfo) *ListCrawlerRunsResponseBody
	GetPagingInfo() *ListCrawlerRunsResponseBodyPagingInfo
	SetRequestId(v string) *ListCrawlerRunsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCrawlerRunsResponseBody
	GetSuccess() *bool
}

type ListCrawlerRunsResponseBody struct {
	// The pagination information.
	PagingInfo *ListCrawlerRunsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID. Used for locating logs and troubleshooting issues.
	//
	// example:
	//
	// 9252F32F-D855-549E-8898-61CF5A733050
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCrawlerRunsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCrawlerRunsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCrawlerRunsResponseBody) GetPagingInfo() *ListCrawlerRunsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListCrawlerRunsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCrawlerRunsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCrawlerRunsResponseBody) SetPagingInfo(v *ListCrawlerRunsResponseBodyPagingInfo) *ListCrawlerRunsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListCrawlerRunsResponseBody) SetRequestId(v string) *ListCrawlerRunsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCrawlerRunsResponseBody) SetSuccess(v bool) *ListCrawlerRunsResponseBody {
	s.Success = &v
	return s
}

func (s *ListCrawlerRunsResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCrawlerRunsResponseBodyPagingInfo struct {
	// The list of metadata crawler run records.
	CrawlerRuns []*ListCrawlerRunsResponseBodyPagingInfoCrawlerRuns `json:"CrawlerRuns,omitempty" xml:"CrawlerRuns,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records that meet the query conditions.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCrawlerRunsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCrawlerRunsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListCrawlerRunsResponseBodyPagingInfo) GetCrawlerRuns() []*ListCrawlerRunsResponseBodyPagingInfoCrawlerRuns {
	return s.CrawlerRuns
}

func (s *ListCrawlerRunsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCrawlerRunsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCrawlerRunsResponseBodyPagingInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCrawlerRunsResponseBodyPagingInfo) SetCrawlerRuns(v []*ListCrawlerRunsResponseBodyPagingInfoCrawlerRuns) *ListCrawlerRunsResponseBodyPagingInfo {
	s.CrawlerRuns = v
	return s
}

func (s *ListCrawlerRunsResponseBodyPagingInfo) SetPageNumber(v int32) *ListCrawlerRunsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListCrawlerRunsResponseBodyPagingInfo) SetPageSize(v int32) *ListCrawlerRunsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListCrawlerRunsResponseBodyPagingInfo) SetTotalCount(v int64) *ListCrawlerRunsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListCrawlerRunsResponseBodyPagingInfo) Validate() error {
	if s.CrawlerRuns != nil {
		for _, item := range s.CrawlerRuns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCrawlerRunsResponseBodyPagingInfoCrawlerRuns struct {
	// The run duration, in seconds.
	//
	// example:
	//
	// 60
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The end time, in millisecond-level UNIX timestamp.
	//
	// example:
	//
	// 1710239065403
	FinishedTime *int64 `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The start time, in millisecond-level UNIX timestamp.
	//
	// example:
	//
	// 1710239005403
	StartedTime *int64 `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	// The run status. Valid values: WAITING, RUNNING, SUCCESS, ERROR, SHUTDOWN.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The DataWorks task instance ID associated with this run.
	//
	// example:
	//
	// 1234
	TaskInstanceId *int64 `json:"TaskInstanceId,omitempty" xml:"TaskInstanceId,omitempty"`
	// The total number of tables involved in this run.
	//
	// example:
	//
	// 42
	TotalTableCount *int64 `json:"TotalTableCount,omitempty" xml:"TotalTableCount,omitempty"`
}

func (s ListCrawlerRunsResponseBodyPagingInfoCrawlerRuns) String() string {
	return dara.Prettify(s)
}

func (s ListCrawlerRunsResponseBodyPagingInfoCrawlerRuns) GoString() string {
	return s.String()
}

func (s *ListCrawlerRunsResponseBodyPagingInfoCrawlerRuns) GetDuration() *float64 {
	return s.Duration
}

func (s *ListCrawlerRunsResponseBodyPagingInfoCrawlerRuns) GetFinishedTime() *int64 {
	return s.FinishedTime
}

func (s *ListCrawlerRunsResponseBodyPagingInfoCrawlerRuns) GetStartedTime() *int64 {
	return s.StartedTime
}

func (s *ListCrawlerRunsResponseBodyPagingInfoCrawlerRuns) GetStatus() *string {
	return s.Status
}

func (s *ListCrawlerRunsResponseBodyPagingInfoCrawlerRuns) GetTaskInstanceId() *int64 {
	return s.TaskInstanceId
}

func (s *ListCrawlerRunsResponseBodyPagingInfoCrawlerRuns) GetTotalTableCount() *int64 {
	return s.TotalTableCount
}

func (s *ListCrawlerRunsResponseBodyPagingInfoCrawlerRuns) SetDuration(v float64) *ListCrawlerRunsResponseBodyPagingInfoCrawlerRuns {
	s.Duration = &v
	return s
}

func (s *ListCrawlerRunsResponseBodyPagingInfoCrawlerRuns) SetFinishedTime(v int64) *ListCrawlerRunsResponseBodyPagingInfoCrawlerRuns {
	s.FinishedTime = &v
	return s
}

func (s *ListCrawlerRunsResponseBodyPagingInfoCrawlerRuns) SetStartedTime(v int64) *ListCrawlerRunsResponseBodyPagingInfoCrawlerRuns {
	s.StartedTime = &v
	return s
}

func (s *ListCrawlerRunsResponseBodyPagingInfoCrawlerRuns) SetStatus(v string) *ListCrawlerRunsResponseBodyPagingInfoCrawlerRuns {
	s.Status = &v
	return s
}

func (s *ListCrawlerRunsResponseBodyPagingInfoCrawlerRuns) SetTaskInstanceId(v int64) *ListCrawlerRunsResponseBodyPagingInfoCrawlerRuns {
	s.TaskInstanceId = &v
	return s
}

func (s *ListCrawlerRunsResponseBodyPagingInfoCrawlerRuns) SetTotalTableCount(v int64) *ListCrawlerRunsResponseBodyPagingInfoCrawlerRuns {
	s.TotalTableCount = &v
	return s
}

func (s *ListCrawlerRunsResponseBodyPagingInfoCrawlerRuns) Validate() error {
	return dara.Validate(s)
}
