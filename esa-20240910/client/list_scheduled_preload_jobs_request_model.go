// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduledPreloadJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ListScheduledPreloadJobsRequest
	GetEndTime() *int64
	SetPageNumber(v int32) *ListScheduledPreloadJobsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListScheduledPreloadJobsRequest
	GetPageSize() *int32
	SetSiteId(v int64) *ListScheduledPreloadJobsRequest
	GetSiteId() *int64
	SetStartTime(v int64) *ListScheduledPreloadJobsRequest
	GetStartTime() *int64
}

type ListScheduledPreloadJobsRequest struct {
	// The end time for querying the task creation time range. Specify the timestamp in milliseconds.
	//
	// example:
	//
	// 1717225277000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number.
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
	// The website ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ListScheduledPreloadJobs
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The beginning of the task creation time range to query. Specify the timestamp in milliseconds.
	//
	// example:
	//
	// 1717225276000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListScheduledPreloadJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledPreloadJobsRequest) GoString() string {
	return s.String()
}

func (s *ListScheduledPreloadJobsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListScheduledPreloadJobsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListScheduledPreloadJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScheduledPreloadJobsRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListScheduledPreloadJobsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListScheduledPreloadJobsRequest) SetEndTime(v int64) *ListScheduledPreloadJobsRequest {
	s.EndTime = &v
	return s
}

func (s *ListScheduledPreloadJobsRequest) SetPageNumber(v int32) *ListScheduledPreloadJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListScheduledPreloadJobsRequest) SetPageSize(v int32) *ListScheduledPreloadJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListScheduledPreloadJobsRequest) SetSiteId(v int64) *ListScheduledPreloadJobsRequest {
	s.SiteId = &v
	return s
}

func (s *ListScheduledPreloadJobsRequest) SetStartTime(v int64) *ListScheduledPreloadJobsRequest {
	s.StartTime = &v
	return s
}

func (s *ListScheduledPreloadJobsRequest) Validate() error {
	return dara.Validate(s)
}
