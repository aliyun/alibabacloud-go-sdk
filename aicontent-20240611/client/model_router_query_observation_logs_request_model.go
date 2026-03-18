// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryObservationLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v int64) *ModelRouterQueryObservationLogsRequest
	GetApiKeyId() *int64
	SetClientId(v int64) *ModelRouterQueryObservationLogsRequest
	GetClientId() *int64
	SetEndTime(v string) *ModelRouterQueryObservationLogsRequest
	GetEndTime() *string
	SetGroupBy(v string) *ModelRouterQueryObservationLogsRequest
	GetGroupBy() *string
	SetMaxResults(v int32) *ModelRouterQueryObservationLogsRequest
	GetMaxResults() *int32
	SetModelId(v int64) *ModelRouterQueryObservationLogsRequest
	GetModelId() *int64
	SetNeedTotalCount(v bool) *ModelRouterQueryObservationLogsRequest
	GetNeedTotalCount() *bool
	SetNextToken(v string) *ModelRouterQueryObservationLogsRequest
	GetNextToken() *string
	SetOrderBy(v string) *ModelRouterQueryObservationLogsRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ModelRouterQueryObservationLogsRequest
	GetOrderDirection() *string
	SetPageIndex(v int32) *ModelRouterQueryObservationLogsRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterQueryObservationLogsRequest
	GetPageSize() *int32
	SetStartTime(v string) *ModelRouterQueryObservationLogsRequest
	GetStartTime() *string
	SetTimeRange(v string) *ModelRouterQueryObservationLogsRequest
	GetTimeRange() *string
}

type ModelRouterQueryObservationLogsRequest struct {
	// example:
	//
	// 1
	ApiKeyId *int64 `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	// example:
	//
	// 1
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// example:
	//
	// 2024-01-02T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// resourceId
	GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 1
	ModelId *int64 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// true
	NeedTotalCount *bool `json:"needTotalCount,omitempty" xml:"needTotalCount,omitempty"`
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// resourceId
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// example:
	//
	// DESC
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 24h
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s ModelRouterQueryObservationLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryObservationLogsRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryObservationLogsRequest) GetApiKeyId() *int64 {
	return s.ApiKeyId
}

func (s *ModelRouterQueryObservationLogsRequest) GetClientId() *int64 {
	return s.ClientId
}

func (s *ModelRouterQueryObservationLogsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ModelRouterQueryObservationLogsRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *ModelRouterQueryObservationLogsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryObservationLogsRequest) GetModelId() *int64 {
	return s.ModelId
}

func (s *ModelRouterQueryObservationLogsRequest) GetNeedTotalCount() *bool {
	return s.NeedTotalCount
}

func (s *ModelRouterQueryObservationLogsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryObservationLogsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ModelRouterQueryObservationLogsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ModelRouterQueryObservationLogsRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterQueryObservationLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryObservationLogsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ModelRouterQueryObservationLogsRequest) GetTimeRange() *string {
	return s.TimeRange
}

func (s *ModelRouterQueryObservationLogsRequest) SetApiKeyId(v int64) *ModelRouterQueryObservationLogsRequest {
	s.ApiKeyId = &v
	return s
}

func (s *ModelRouterQueryObservationLogsRequest) SetClientId(v int64) *ModelRouterQueryObservationLogsRequest {
	s.ClientId = &v
	return s
}

func (s *ModelRouterQueryObservationLogsRequest) SetEndTime(v string) *ModelRouterQueryObservationLogsRequest {
	s.EndTime = &v
	return s
}

func (s *ModelRouterQueryObservationLogsRequest) SetGroupBy(v string) *ModelRouterQueryObservationLogsRequest {
	s.GroupBy = &v
	return s
}

func (s *ModelRouterQueryObservationLogsRequest) SetMaxResults(v int32) *ModelRouterQueryObservationLogsRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryObservationLogsRequest) SetModelId(v int64) *ModelRouterQueryObservationLogsRequest {
	s.ModelId = &v
	return s
}

func (s *ModelRouterQueryObservationLogsRequest) SetNeedTotalCount(v bool) *ModelRouterQueryObservationLogsRequest {
	s.NeedTotalCount = &v
	return s
}

func (s *ModelRouterQueryObservationLogsRequest) SetNextToken(v string) *ModelRouterQueryObservationLogsRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryObservationLogsRequest) SetOrderBy(v string) *ModelRouterQueryObservationLogsRequest {
	s.OrderBy = &v
	return s
}

func (s *ModelRouterQueryObservationLogsRequest) SetOrderDirection(v string) *ModelRouterQueryObservationLogsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ModelRouterQueryObservationLogsRequest) SetPageIndex(v int32) *ModelRouterQueryObservationLogsRequest {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterQueryObservationLogsRequest) SetPageSize(v int32) *ModelRouterQueryObservationLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryObservationLogsRequest) SetStartTime(v string) *ModelRouterQueryObservationLogsRequest {
	s.StartTime = &v
	return s
}

func (s *ModelRouterQueryObservationLogsRequest) SetTimeRange(v string) *ModelRouterQueryObservationLogsRequest {
	s.TimeRange = &v
	return s
}

func (s *ModelRouterQueryObservationLogsRequest) Validate() error {
	return dara.Validate(s)
}
