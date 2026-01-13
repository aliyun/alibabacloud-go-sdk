// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutorEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExecutorEventList(v []*ListExecutorEventsResponseBodyExecutorEventList) *ListExecutorEventsResponseBody
	GetExecutorEventList() []*ListExecutorEventsResponseBodyExecutorEventList
	SetPageNumber(v int32) *ListExecutorEventsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListExecutorEventsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListExecutorEventsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListExecutorEventsResponseBody
	GetTotalCount() *int32
}

type ListExecutorEventsResponseBody struct {
	// The list of the running event.
	ExecutorEventList []*ListExecutorEventsResponseBodyExecutorEventList `json:"ExecutorEventList,omitempty" xml:"ExecutorEventList,omitempty" type:"Repeated"`
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
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 40
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListExecutorEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExecutorEventsResponseBody) GetExecutorEventList() []*ListExecutorEventsResponseBodyExecutorEventList {
	return s.ExecutorEventList
}

func (s *ListExecutorEventsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListExecutorEventsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExecutorEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExecutorEventsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListExecutorEventsResponseBody) SetExecutorEventList(v []*ListExecutorEventsResponseBodyExecutorEventList) *ListExecutorEventsResponseBody {
	s.ExecutorEventList = v
	return s
}

func (s *ListExecutorEventsResponseBody) SetPageNumber(v int32) *ListExecutorEventsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListExecutorEventsResponseBody) SetPageSize(v int32) *ListExecutorEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListExecutorEventsResponseBody) SetRequestId(v string) *ListExecutorEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExecutorEventsResponseBody) SetTotalCount(v int32) *ListExecutorEventsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListExecutorEventsResponseBody) Validate() error {
	if s.ExecutorEventList != nil {
		for _, item := range s.ExecutorEventList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListExecutorEventsResponseBodyExecutorEventList struct {
	// The content of the running event.
	//
	// example:
	//
	// Executor created successfully
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the executor. The format is JobId-TaskName-ArrayIndex.
	//
	// example:
	//
	// job-xxxx-Task0-1
	ExecutorId *string `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// job-xxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The level of the running event. Valid values:
	//
	// 	- Normal
	//
	// 	- Warning
	//
	// 	- Error
	//
	// example:
	//
	// Normal
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The event of the running event.
	//
	// example:
	//
	// 2024-02-20 10:04:13
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s ListExecutorEventsResponseBodyExecutorEventList) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorEventsResponseBodyExecutorEventList) GoString() string {
	return s.String()
}

func (s *ListExecutorEventsResponseBodyExecutorEventList) GetContent() *string {
	return s.Content
}

func (s *ListExecutorEventsResponseBodyExecutorEventList) GetExecutorId() *string {
	return s.ExecutorId
}

func (s *ListExecutorEventsResponseBodyExecutorEventList) GetJobId() *string {
	return s.JobId
}

func (s *ListExecutorEventsResponseBodyExecutorEventList) GetLevel() *string {
	return s.Level
}

func (s *ListExecutorEventsResponseBodyExecutorEventList) GetTime() *string {
	return s.Time
}

func (s *ListExecutorEventsResponseBodyExecutorEventList) SetContent(v string) *ListExecutorEventsResponseBodyExecutorEventList {
	s.Content = &v
	return s
}

func (s *ListExecutorEventsResponseBodyExecutorEventList) SetExecutorId(v string) *ListExecutorEventsResponseBodyExecutorEventList {
	s.ExecutorId = &v
	return s
}

func (s *ListExecutorEventsResponseBodyExecutorEventList) SetJobId(v string) *ListExecutorEventsResponseBodyExecutorEventList {
	s.JobId = &v
	return s
}

func (s *ListExecutorEventsResponseBodyExecutorEventList) SetLevel(v string) *ListExecutorEventsResponseBodyExecutorEventList {
	s.Level = &v
	return s
}

func (s *ListExecutorEventsResponseBodyExecutorEventList) SetTime(v string) *ListExecutorEventsResponseBodyExecutorEventList {
	s.Time = &v
	return s
}

func (s *ListExecutorEventsResponseBodyExecutorEventList) Validate() error {
	return dara.Validate(s)
}
