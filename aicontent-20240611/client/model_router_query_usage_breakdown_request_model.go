// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryUsageBreakdownRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ModelRouterQueryUsageBreakdownRequest
	GetEndTime() *int64
	SetGranularity(v string) *ModelRouterQueryUsageBreakdownRequest
	GetGranularity() *string
	SetMaxResults(v int32) *ModelRouterQueryUsageBreakdownRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryUsageBreakdownRequest
	GetNextToken() *string
	SetPage(v int32) *ModelRouterQueryUsageBreakdownRequest
	GetPage() *int32
	SetPageSize(v int32) *ModelRouterQueryUsageBreakdownRequest
	GetPageSize() *int32
	SetStartTime(v int64) *ModelRouterQueryUsageBreakdownRequest
	GetStartTime() *int64
}

type ModelRouterQueryUsageBreakdownRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1700086400
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hourly
	Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// xxxx-xxx-xxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1700000000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ModelRouterQueryUsageBreakdownRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryUsageBreakdownRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryUsageBreakdownRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ModelRouterQueryUsageBreakdownRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *ModelRouterQueryUsageBreakdownRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryUsageBreakdownRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryUsageBreakdownRequest) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterQueryUsageBreakdownRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryUsageBreakdownRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ModelRouterQueryUsageBreakdownRequest) SetEndTime(v int64) *ModelRouterQueryUsageBreakdownRequest {
	s.EndTime = &v
	return s
}

func (s *ModelRouterQueryUsageBreakdownRequest) SetGranularity(v string) *ModelRouterQueryUsageBreakdownRequest {
	s.Granularity = &v
	return s
}

func (s *ModelRouterQueryUsageBreakdownRequest) SetMaxResults(v int32) *ModelRouterQueryUsageBreakdownRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryUsageBreakdownRequest) SetNextToken(v string) *ModelRouterQueryUsageBreakdownRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryUsageBreakdownRequest) SetPage(v int32) *ModelRouterQueryUsageBreakdownRequest {
	s.Page = &v
	return s
}

func (s *ModelRouterQueryUsageBreakdownRequest) SetPageSize(v int32) *ModelRouterQueryUsageBreakdownRequest {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryUsageBreakdownRequest) SetStartTime(v int64) *ModelRouterQueryUsageBreakdownRequest {
	s.StartTime = &v
	return s
}

func (s *ModelRouterQueryUsageBreakdownRequest) Validate() error {
	return dara.Validate(s)
}
