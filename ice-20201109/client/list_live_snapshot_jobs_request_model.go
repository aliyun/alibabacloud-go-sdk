// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveSnapshotJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListLiveSnapshotJobsRequest
	GetEndTime() *string
	SetPageNo(v int32) *ListLiveSnapshotJobsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListLiveSnapshotJobsRequest
	GetPageSize() *int32
	SetSearchKeyWord(v string) *ListLiveSnapshotJobsRequest
	GetSearchKeyWord() *string
	SetSortBy(v string) *ListLiveSnapshotJobsRequest
	GetSortBy() *string
	SetStartTime(v string) *ListLiveSnapshotJobsRequest
	GetStartTime() *string
	SetStatus(v string) *ListLiveSnapshotJobsRequest
	GetStatus() *string
}

type ListLiveSnapshotJobsRequest struct {
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// 	- By default, EndTime is seven days later than StartTime.
	//
	// 	- The time range specified by the StartTime and EndTime parameters cannot exceed 30 days.
	//
	// example:
	//
	// 2022-02-02T23:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number. Valid values: [1,n). Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search keyword. You can use the job ID or name as the keyword to search for jobs. If you search for jobs by name, fuzzy match is supported.
	//
	// 	- It cannot exceed 128 characters in length.
	//
	// example:
	//
	// ****a046-263c-3560-978a-fb287782****
	SearchKeyWord *string `json:"SearchKeyWord,omitempty" xml:"SearchKeyWord,omitempty"`
	// The sorting order. By default, the query results are sorted by creation time in descending order.
	//
	// Valid values:
	//
	// 	- asc: sorts the query results by creation time in ascending order.
	//
	// 	- desc: sorts the query results by creation time in descending order.
	//
	// example:
	//
	// desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// 	- The default value is seven days ago.
	//
	// 	- The time range specified by the StartTime and EndTime parameters cannot exceed 30 days.
	//
	// example:
	//
	// 2022-02-02T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The job state filter. By default, all jobs are queried.
	//
	// Valid values:
	//
	// 	- init: The job is not started.
	//
	// 	- paused: The job is paused.
	//
	// 	- started: The job is in progress.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLiveSnapshotJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLiveSnapshotJobsRequest) GoString() string {
	return s.String()
}

func (s *ListLiveSnapshotJobsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListLiveSnapshotJobsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListLiveSnapshotJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLiveSnapshotJobsRequest) GetSearchKeyWord() *string {
	return s.SearchKeyWord
}

func (s *ListLiveSnapshotJobsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLiveSnapshotJobsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListLiveSnapshotJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListLiveSnapshotJobsRequest) SetEndTime(v string) *ListLiveSnapshotJobsRequest {
	s.EndTime = &v
	return s
}

func (s *ListLiveSnapshotJobsRequest) SetPageNo(v int32) *ListLiveSnapshotJobsRequest {
	s.PageNo = &v
	return s
}

func (s *ListLiveSnapshotJobsRequest) SetPageSize(v int32) *ListLiveSnapshotJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListLiveSnapshotJobsRequest) SetSearchKeyWord(v string) *ListLiveSnapshotJobsRequest {
	s.SearchKeyWord = &v
	return s
}

func (s *ListLiveSnapshotJobsRequest) SetSortBy(v string) *ListLiveSnapshotJobsRequest {
	s.SortBy = &v
	return s
}

func (s *ListLiveSnapshotJobsRequest) SetStartTime(v string) *ListLiveSnapshotJobsRequest {
	s.StartTime = &v
	return s
}

func (s *ListLiveSnapshotJobsRequest) SetStatus(v string) *ListLiveSnapshotJobsRequest {
	s.Status = &v
	return s
}

func (s *ListLiveSnapshotJobsRequest) Validate() error {
	return dara.Validate(s)
}
