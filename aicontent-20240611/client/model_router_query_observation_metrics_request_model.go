// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryObservationMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v int64) *ModelRouterQueryObservationMetricsRequest
	GetApiKeyId() *int64
	SetClientId(v int64) *ModelRouterQueryObservationMetricsRequest
	GetClientId() *int64
	SetEndTime(v string) *ModelRouterQueryObservationMetricsRequest
	GetEndTime() *string
	SetGroupBy(v string) *ModelRouterQueryObservationMetricsRequest
	GetGroupBy() *string
	SetMaxResults(v int32) *ModelRouterQueryObservationMetricsRequest
	GetMaxResults() *int32
	SetMemberUserIds(v string) *ModelRouterQueryObservationMetricsRequest
	GetMemberUserIds() *string
	SetModelId(v int64) *ModelRouterQueryObservationMetricsRequest
	GetModelId() *int64
	SetNeedTotalCount(v bool) *ModelRouterQueryObservationMetricsRequest
	GetNeedTotalCount() *bool
	SetNextToken(v string) *ModelRouterQueryObservationMetricsRequest
	GetNextToken() *string
	SetOrderBy(v string) *ModelRouterQueryObservationMetricsRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ModelRouterQueryObservationMetricsRequest
	GetOrderDirection() *string
	SetPageIndex(v int32) *ModelRouterQueryObservationMetricsRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterQueryObservationMetricsRequest
	GetPageSize() *int32
	SetStartTime(v string) *ModelRouterQueryObservationMetricsRequest
	GetStartTime() *string
	SetTimeRange(v string) *ModelRouterQueryObservationMetricsRequest
	GetTimeRange() *string
}

type ModelRouterQueryObservationMetricsRequest struct {
	// The API key ID used to filter the results.
	//
	// example:
	//
	// 1
	ApiKeyId *int64 `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	// The client ID used to filter the results.
	//
	// example:
	//
	// 1
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// The custom end time.
	//
	// example:
	//
	// 2024-01-02T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The field by which to group the results.
	//
	// example:
	//
	// resourceId
	GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
	// The maximum number of results to return.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// Optional. Filters by member IDs. Separate multiple IDs with commas. If not specified, the department and all its members are included. If an empty value is specified, only the department is included without members.
	//
	// example:
	//
	// 30001,30002
	MemberUserIds *string `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty"`
	// The model ID used to filter the results.
	//
	// example:
	//
	// 1
	ModelId *int64 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// Specifies whether to return the total count.
	//
	// example:
	//
	// true
	NeedTotalCount *bool `json:"needTotalCount,omitempty" xml:"needTotalCount,omitempty"`
	// The token for the next query. An empty value indicates the last page.
	//
	// example:
	//
	// 2
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The field by which to sort the results.
	//
	// example:
	//
	// resourceId
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// The sort direction.
	//
	// example:
	//
	// DESC
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The custom start time.
	//
	// example:
	//
	// 2024-01-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The time range for the query. Valid values: 1h, 6h, 24h, 7d, 30d.
	//
	// example:
	//
	// 24h
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s ModelRouterQueryObservationMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryObservationMetricsRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryObservationMetricsRequest) GetApiKeyId() *int64 {
	return s.ApiKeyId
}

func (s *ModelRouterQueryObservationMetricsRequest) GetClientId() *int64 {
	return s.ClientId
}

func (s *ModelRouterQueryObservationMetricsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ModelRouterQueryObservationMetricsRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *ModelRouterQueryObservationMetricsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryObservationMetricsRequest) GetMemberUserIds() *string {
	return s.MemberUserIds
}

func (s *ModelRouterQueryObservationMetricsRequest) GetModelId() *int64 {
	return s.ModelId
}

func (s *ModelRouterQueryObservationMetricsRequest) GetNeedTotalCount() *bool {
	return s.NeedTotalCount
}

func (s *ModelRouterQueryObservationMetricsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryObservationMetricsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ModelRouterQueryObservationMetricsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ModelRouterQueryObservationMetricsRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterQueryObservationMetricsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryObservationMetricsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ModelRouterQueryObservationMetricsRequest) GetTimeRange() *string {
	return s.TimeRange
}

func (s *ModelRouterQueryObservationMetricsRequest) SetApiKeyId(v int64) *ModelRouterQueryObservationMetricsRequest {
	s.ApiKeyId = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsRequest) SetClientId(v int64) *ModelRouterQueryObservationMetricsRequest {
	s.ClientId = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsRequest) SetEndTime(v string) *ModelRouterQueryObservationMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsRequest) SetGroupBy(v string) *ModelRouterQueryObservationMetricsRequest {
	s.GroupBy = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsRequest) SetMaxResults(v int32) *ModelRouterQueryObservationMetricsRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsRequest) SetMemberUserIds(v string) *ModelRouterQueryObservationMetricsRequest {
	s.MemberUserIds = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsRequest) SetModelId(v int64) *ModelRouterQueryObservationMetricsRequest {
	s.ModelId = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsRequest) SetNeedTotalCount(v bool) *ModelRouterQueryObservationMetricsRequest {
	s.NeedTotalCount = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsRequest) SetNextToken(v string) *ModelRouterQueryObservationMetricsRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsRequest) SetOrderBy(v string) *ModelRouterQueryObservationMetricsRequest {
	s.OrderBy = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsRequest) SetOrderDirection(v string) *ModelRouterQueryObservationMetricsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsRequest) SetPageIndex(v int32) *ModelRouterQueryObservationMetricsRequest {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsRequest) SetPageSize(v int32) *ModelRouterQueryObservationMetricsRequest {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsRequest) SetStartTime(v string) *ModelRouterQueryObservationMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsRequest) SetTimeRange(v string) *ModelRouterQueryObservationMetricsRequest {
	s.TimeRange = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsRequest) Validate() error {
	return dara.Validate(s)
}
