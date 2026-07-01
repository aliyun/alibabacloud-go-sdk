// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTranscodeJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndOfCreateTime(v string) *ListTranscodeJobsRequest
	GetEndOfCreateTime() *string
	SetNextPageToken(v string) *ListTranscodeJobsRequest
	GetNextPageToken() *string
	SetOrderBy(v string) *ListTranscodeJobsRequest
	GetOrderBy() *string
	SetPageSize(v int32) *ListTranscodeJobsRequest
	GetPageSize() *int32
	SetParentJobId(v string) *ListTranscodeJobsRequest
	GetParentJobId() *string
	SetStartOfCreateTime(v string) *ListTranscodeJobsRequest
	GetStartOfCreateTime() *string
	SetStatus(v string) *ListTranscodeJobsRequest
	GetStatus() *string
}

type ListTranscodeJobsRequest struct {
	// The end of the time range to filter jobs by their creation time. The time must be in UTC and formatted as `yyyy-MM-ddTHH:mm:ssZ`.
	//
	// example:
	//
	// 2022-07-15T00:00:00Z
	EndOfCreateTime *string `json:"EndOfCreateTime,omitempty" xml:"EndOfCreateTime,omitempty"`
	// The token for the next page of results. Not required for the first page.
	//
	// example:
	//
	// ab4802364a2e49208c99efab82df****
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The sort order. Valid values:
	//
	// - `CreateTimeDesc`: Sorts by creation time in descending order.
	//
	// - `CreateTimeAsc`: Sorts by creation time in ascending order.
	//
	// example:
	//
	// CreateTimeDesc
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The number of entries per page. Valid values: 1-100. Default: 20.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Filters by job ID.
	//
	// example:
	//
	// 7b38a5d86f1e47838927b6e7ccb1****
	ParentJobId *string `json:"ParentJobId,omitempty" xml:"ParentJobId,omitempty"`
	// The start of the time range to filter jobs by their creation time. The time must be in UTC and formatted as `yyyy-MM-ddTHH:mm:ssZ`.
	//
	// example:
	//
	// 2022-07-01T00:00:00Z
	StartOfCreateTime *string `json:"StartOfCreateTime,omitempty" xml:"StartOfCreateTime,omitempty"`
	// The job status. Valid values:
	//
	// - `Init`: Submitted.
	//
	// - `Success`: Succeeded.
	//
	// - `Fail`: Failed.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListTranscodeJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsRequest) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsRequest) GetEndOfCreateTime() *string {
	return s.EndOfCreateTime
}

func (s *ListTranscodeJobsRequest) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListTranscodeJobsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListTranscodeJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTranscodeJobsRequest) GetParentJobId() *string {
	return s.ParentJobId
}

func (s *ListTranscodeJobsRequest) GetStartOfCreateTime() *string {
	return s.StartOfCreateTime
}

func (s *ListTranscodeJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTranscodeJobsRequest) SetEndOfCreateTime(v string) *ListTranscodeJobsRequest {
	s.EndOfCreateTime = &v
	return s
}

func (s *ListTranscodeJobsRequest) SetNextPageToken(v string) *ListTranscodeJobsRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListTranscodeJobsRequest) SetOrderBy(v string) *ListTranscodeJobsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListTranscodeJobsRequest) SetPageSize(v int32) *ListTranscodeJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTranscodeJobsRequest) SetParentJobId(v string) *ListTranscodeJobsRequest {
	s.ParentJobId = &v
	return s
}

func (s *ListTranscodeJobsRequest) SetStartOfCreateTime(v string) *ListTranscodeJobsRequest {
	s.StartOfCreateTime = &v
	return s
}

func (s *ListTranscodeJobsRequest) SetStatus(v string) *ListTranscodeJobsRequest {
	s.Status = &v
	return s
}

func (s *ListTranscodeJobsRequest) Validate() error {
	return dara.Validate(s)
}
