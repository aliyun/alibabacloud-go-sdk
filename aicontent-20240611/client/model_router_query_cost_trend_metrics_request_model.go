// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryCostTrendMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v int64) *ModelRouterQueryCostTrendMetricsRequest
	GetApiKeyId() *int64
	SetClientId(v int64) *ModelRouterQueryCostTrendMetricsRequest
	GetClientId() *int64
	SetClientIds(v string) *ModelRouterQueryCostTrendMetricsRequest
	GetClientIds() *string
	SetEndTime(v int64) *ModelRouterQueryCostTrendMetricsRequest
	GetEndTime() *int64
	SetGranularity(v string) *ModelRouterQueryCostTrendMetricsRequest
	GetGranularity() *string
	SetMaxResults(v int32) *ModelRouterQueryCostTrendMetricsRequest
	GetMaxResults() *int32
	SetMemberUserIds(v string) *ModelRouterQueryCostTrendMetricsRequest
	GetMemberUserIds() *string
	SetModelTypes(v string) *ModelRouterQueryCostTrendMetricsRequest
	GetModelTypes() *string
	SetNextToken(v string) *ModelRouterQueryCostTrendMetricsRequest
	GetNextToken() *string
	SetStartTime(v int64) *ModelRouterQueryCostTrendMetricsRequest
	GetStartTime() *int64
}

type ModelRouterQueryCostTrendMetricsRequest struct {
	// Optional. Filters by API Key ID. This parameter is linked to the department and requires clientId to be specified first.
	//
	// example:
	//
	// 100
	ApiKeyId *int64 `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	// The department ID used to filter results.
	//
	// example:
	//
	// 1
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// The list of department IDs, separated by commas. Supports querying data for multiple departments. This parameter is mutually exclusive with clientId.
	//
	// example:
	//
	// 1,2,3
	ClientIds *string `json:"clientIds,omitempty" xml:"clientIds,omitempty"`
	// The end time, as a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1700086400
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The time granularity. Valid values:
	//
	// - hourly
	//
	// - daily
	//
	// Default value: hourly.
	//
	// example:
	//
	// hourly
	Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
	// The maximum number of results to return.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// Optional. Filters by member IDs, separated by commas. If not specified, data for the department and all its members is returned. If an empty value is specified, only department data without members is returned.
	//
	// example:
	//
	// 30001,30002
	MemberUserIds *string `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty"`
	// The model types, separated by commas.
	//
	// example:
	//
	// LLM,VL
	ModelTypes *string `json:"modelTypes,omitempty" xml:"modelTypes,omitempty"`
	// nextToken
	//
	// example:
	//
	// xxxx-xxx-xxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The start time, as a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1700000000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ModelRouterQueryCostTrendMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryCostTrendMetricsRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryCostTrendMetricsRequest) GetApiKeyId() *int64 {
	return s.ApiKeyId
}

func (s *ModelRouterQueryCostTrendMetricsRequest) GetClientId() *int64 {
	return s.ClientId
}

func (s *ModelRouterQueryCostTrendMetricsRequest) GetClientIds() *string {
	return s.ClientIds
}

func (s *ModelRouterQueryCostTrendMetricsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ModelRouterQueryCostTrendMetricsRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *ModelRouterQueryCostTrendMetricsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryCostTrendMetricsRequest) GetMemberUserIds() *string {
	return s.MemberUserIds
}

func (s *ModelRouterQueryCostTrendMetricsRequest) GetModelTypes() *string {
	return s.ModelTypes
}

func (s *ModelRouterQueryCostTrendMetricsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryCostTrendMetricsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ModelRouterQueryCostTrendMetricsRequest) SetApiKeyId(v int64) *ModelRouterQueryCostTrendMetricsRequest {
	s.ApiKeyId = &v
	return s
}

func (s *ModelRouterQueryCostTrendMetricsRequest) SetClientId(v int64) *ModelRouterQueryCostTrendMetricsRequest {
	s.ClientId = &v
	return s
}

func (s *ModelRouterQueryCostTrendMetricsRequest) SetClientIds(v string) *ModelRouterQueryCostTrendMetricsRequest {
	s.ClientIds = &v
	return s
}

func (s *ModelRouterQueryCostTrendMetricsRequest) SetEndTime(v int64) *ModelRouterQueryCostTrendMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *ModelRouterQueryCostTrendMetricsRequest) SetGranularity(v string) *ModelRouterQueryCostTrendMetricsRequest {
	s.Granularity = &v
	return s
}

func (s *ModelRouterQueryCostTrendMetricsRequest) SetMaxResults(v int32) *ModelRouterQueryCostTrendMetricsRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryCostTrendMetricsRequest) SetMemberUserIds(v string) *ModelRouterQueryCostTrendMetricsRequest {
	s.MemberUserIds = &v
	return s
}

func (s *ModelRouterQueryCostTrendMetricsRequest) SetModelTypes(v string) *ModelRouterQueryCostTrendMetricsRequest {
	s.ModelTypes = &v
	return s
}

func (s *ModelRouterQueryCostTrendMetricsRequest) SetNextToken(v string) *ModelRouterQueryCostTrendMetricsRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryCostTrendMetricsRequest) SetStartTime(v int64) *ModelRouterQueryCostTrendMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *ModelRouterQueryCostTrendMetricsRequest) Validate() error {
	return dara.Validate(s)
}
