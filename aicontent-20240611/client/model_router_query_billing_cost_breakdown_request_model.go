// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryBillingCostBreakdownRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ModelRouterQueryBillingCostBreakdownRequest
	GetEndTime() *int64
	SetGranularity(v string) *ModelRouterQueryBillingCostBreakdownRequest
	GetGranularity() *string
	SetMaxResults(v int32) *ModelRouterQueryBillingCostBreakdownRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryBillingCostBreakdownRequest
	GetNextToken() *string
	SetPage(v int32) *ModelRouterQueryBillingCostBreakdownRequest
	GetPage() *int32
	SetPageSize(v int32) *ModelRouterQueryBillingCostBreakdownRequest
	GetPageSize() *int32
	SetStartTime(v int64) *ModelRouterQueryBillingCostBreakdownRequest
	GetStartTime() *int64
}

type ModelRouterQueryBillingCostBreakdownRequest struct {
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

func (s ModelRouterQueryBillingCostBreakdownRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryBillingCostBreakdownRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryBillingCostBreakdownRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ModelRouterQueryBillingCostBreakdownRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *ModelRouterQueryBillingCostBreakdownRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryBillingCostBreakdownRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryBillingCostBreakdownRequest) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterQueryBillingCostBreakdownRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryBillingCostBreakdownRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ModelRouterQueryBillingCostBreakdownRequest) SetEndTime(v int64) *ModelRouterQueryBillingCostBreakdownRequest {
	s.EndTime = &v
	return s
}

func (s *ModelRouterQueryBillingCostBreakdownRequest) SetGranularity(v string) *ModelRouterQueryBillingCostBreakdownRequest {
	s.Granularity = &v
	return s
}

func (s *ModelRouterQueryBillingCostBreakdownRequest) SetMaxResults(v int32) *ModelRouterQueryBillingCostBreakdownRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryBillingCostBreakdownRequest) SetNextToken(v string) *ModelRouterQueryBillingCostBreakdownRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryBillingCostBreakdownRequest) SetPage(v int32) *ModelRouterQueryBillingCostBreakdownRequest {
	s.Page = &v
	return s
}

func (s *ModelRouterQueryBillingCostBreakdownRequest) SetPageSize(v int32) *ModelRouterQueryBillingCostBreakdownRequest {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryBillingCostBreakdownRequest) SetStartTime(v int64) *ModelRouterQueryBillingCostBreakdownRequest {
	s.StartTime = &v
	return s
}

func (s *ModelRouterQueryBillingCostBreakdownRequest) Validate() error {
	return dara.Validate(s)
}
