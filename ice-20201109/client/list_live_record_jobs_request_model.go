// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveRecordJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListLiveRecordJobsRequest
	GetEndTime() *string
	SetKeyword(v string) *ListLiveRecordJobsRequest
	GetKeyword() *string
	SetPageNo(v int64) *ListLiveRecordJobsRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListLiveRecordJobsRequest
	GetPageSize() *int64
	SetSortBy(v string) *ListLiveRecordJobsRequest
	GetSortBy() *string
	SetStartTime(v string) *ListLiveRecordJobsRequest
	GetStartTime() *string
	SetStatus(v string) *ListLiveRecordJobsRequest
	GetStatus() *string
}

type ListLiveRecordJobsRequest struct {
	// The end of the time range to query. The maximum time range between EndTime and StartTime cannot exceed 30 days. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-07-11T08:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The search keyword. You can use the job ID or name as the keyword to search for jobs.
	//
	// example:
	//
	// ab0e3e76-1e9d-11ed-ba64-0c42a1b73d66
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sorting order. By default, the query results are sorted by creation time in descending order.
	//
	// Valid values:
	//
	// 	- asc: sorts the query results in ascending order.
	//
	// 	- desc: sorts the query results in descending order.
	//
	// example:
	//
	// desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-07-15T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the job. By default, the state is not filtered.
	//
	// Valid values:
	//
	// 	- paused: The job is paused.
	//
	// 	- initial: The job is not started.
	//
	// 	- started: The job is in progress.
	//
	// example:
	//
	// started
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLiveRecordJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRecordJobsRequest) GoString() string {
	return s.String()
}

func (s *ListLiveRecordJobsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListLiveRecordJobsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListLiveRecordJobsRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListLiveRecordJobsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListLiveRecordJobsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLiveRecordJobsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListLiveRecordJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListLiveRecordJobsRequest) SetEndTime(v string) *ListLiveRecordJobsRequest {
	s.EndTime = &v
	return s
}

func (s *ListLiveRecordJobsRequest) SetKeyword(v string) *ListLiveRecordJobsRequest {
	s.Keyword = &v
	return s
}

func (s *ListLiveRecordJobsRequest) SetPageNo(v int64) *ListLiveRecordJobsRequest {
	s.PageNo = &v
	return s
}

func (s *ListLiveRecordJobsRequest) SetPageSize(v int64) *ListLiveRecordJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListLiveRecordJobsRequest) SetSortBy(v string) *ListLiveRecordJobsRequest {
	s.SortBy = &v
	return s
}

func (s *ListLiveRecordJobsRequest) SetStartTime(v string) *ListLiveRecordJobsRequest {
	s.StartTime = &v
	return s
}

func (s *ListLiveRecordJobsRequest) SetStatus(v string) *ListLiveRecordJobsRequest {
	s.Status = &v
	return s
}

func (s *ListLiveRecordJobsRequest) Validate() error {
	return dara.Validate(s)
}
