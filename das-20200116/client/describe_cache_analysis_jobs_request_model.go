// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCacheAnalysisJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeCacheAnalysisJobsRequest
	GetEndTime() *string
	SetInstanceId(v string) *DescribeCacheAnalysisJobsRequest
	GetInstanceId() *string
	SetPageNo(v string) *DescribeCacheAnalysisJobsRequest
	GetPageNo() *string
	SetPageSize(v string) *DescribeCacheAnalysisJobsRequest
	GetPageSize() *string
	SetStartTime(v string) *DescribeCacheAnalysisJobsRequest
	GetStartTime() *string
}

type DescribeCacheAnalysisJobsRequest struct {
	// The end of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1596177993001
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp18ff4a195d****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. The value must be an integer that is greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1596177993000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCacheAnalysisJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisJobsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCacheAnalysisJobsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCacheAnalysisJobsRequest) GetPageNo() *string {
	return s.PageNo
}

func (s *DescribeCacheAnalysisJobsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeCacheAnalysisJobsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCacheAnalysisJobsRequest) SetEndTime(v string) *DescribeCacheAnalysisJobsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCacheAnalysisJobsRequest) SetInstanceId(v string) *DescribeCacheAnalysisJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCacheAnalysisJobsRequest) SetPageNo(v string) *DescribeCacheAnalysisJobsRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeCacheAnalysisJobsRequest) SetPageSize(v string) *DescribeCacheAnalysisJobsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCacheAnalysisJobsRequest) SetStartTime(v string) *DescribeCacheAnalysisJobsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCacheAnalysisJobsRequest) Validate() error {
	return dara.Validate(s)
}
