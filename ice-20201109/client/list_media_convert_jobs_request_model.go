// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaConvertJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndOfCreateTime(v string) *ListMediaConvertJobsRequest
	GetEndOfCreateTime() *string
	SetJobId(v string) *ListMediaConvertJobsRequest
	GetJobId() *string
	SetNextPageToken(v string) *ListMediaConvertJobsRequest
	GetNextPageToken() *string
	SetOrderBy(v string) *ListMediaConvertJobsRequest
	GetOrderBy() *string
	SetPageSize(v int32) *ListMediaConvertJobsRequest
	GetPageSize() *int32
	SetStartOfCreateTime(v string) *ListMediaConvertJobsRequest
	GetStartOfCreateTime() *string
	SetStatus(v string) *ListMediaConvertJobsRequest
	GetStatus() *string
}

type ListMediaConvertJobsRequest struct {
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-07-15T00:00:00Z
	EndOfCreateTime *string `json:"EndOfCreateTime,omitempty" xml:"EndOfCreateTime,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 88c6ca184c0e47098a5b665e2****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// ab4802364a2e49208c99efab82dfa8e8
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The sorting order. Valid values: CreateTimeDesc: sorts by create time in descending order. CreateTimeAsc: sorts by create time in ascending order.
	//
	// example:
	//
	// CreateTimeDesc
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The number of entries per page. Valid values: 0 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-07-01T00:00:00Z
	StartOfCreateTime *string `json:"StartOfCreateTime,omitempty" xml:"StartOfCreateTime,omitempty"`
	// The task status.
	//
	// 	- Inited: submitted
	//
	// 	- Running
	//
	// 	- Complete
	//
	// 	- Error
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListMediaConvertJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMediaConvertJobsRequest) GoString() string {
	return s.String()
}

func (s *ListMediaConvertJobsRequest) GetEndOfCreateTime() *string {
	return s.EndOfCreateTime
}

func (s *ListMediaConvertJobsRequest) GetJobId() *string {
	return s.JobId
}

func (s *ListMediaConvertJobsRequest) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListMediaConvertJobsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListMediaConvertJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMediaConvertJobsRequest) GetStartOfCreateTime() *string {
	return s.StartOfCreateTime
}

func (s *ListMediaConvertJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListMediaConvertJobsRequest) SetEndOfCreateTime(v string) *ListMediaConvertJobsRequest {
	s.EndOfCreateTime = &v
	return s
}

func (s *ListMediaConvertJobsRequest) SetJobId(v string) *ListMediaConvertJobsRequest {
	s.JobId = &v
	return s
}

func (s *ListMediaConvertJobsRequest) SetNextPageToken(v string) *ListMediaConvertJobsRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListMediaConvertJobsRequest) SetOrderBy(v string) *ListMediaConvertJobsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListMediaConvertJobsRequest) SetPageSize(v int32) *ListMediaConvertJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMediaConvertJobsRequest) SetStartOfCreateTime(v string) *ListMediaConvertJobsRequest {
	s.StartOfCreateTime = &v
	return s
}

func (s *ListMediaConvertJobsRequest) SetStatus(v string) *ListMediaConvertJobsRequest {
	s.Status = &v
	return s
}

func (s *ListMediaConvertJobsRequest) Validate() error {
	return dara.Validate(s)
}
