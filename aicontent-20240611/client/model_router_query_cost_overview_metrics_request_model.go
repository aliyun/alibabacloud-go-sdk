// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryCostOverviewMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v int64) *ModelRouterQueryCostOverviewMetricsRequest
	GetApiKeyId() *int64
	SetClientId(v int64) *ModelRouterQueryCostOverviewMetricsRequest
	GetClientId() *int64
	SetClientIds(v string) *ModelRouterQueryCostOverviewMetricsRequest
	GetClientIds() *string
	SetEndTime(v int64) *ModelRouterQueryCostOverviewMetricsRequest
	GetEndTime() *int64
	SetGranularity(v string) *ModelRouterQueryCostOverviewMetricsRequest
	GetGranularity() *string
	SetMaxResults(v int32) *ModelRouterQueryCostOverviewMetricsRequest
	GetMaxResults() *int32
	SetMemberUserIds(v string) *ModelRouterQueryCostOverviewMetricsRequest
	GetMemberUserIds() *string
	SetModelTypes(v string) *ModelRouterQueryCostOverviewMetricsRequest
	GetModelTypes() *string
	SetNextToken(v string) *ModelRouterQueryCostOverviewMetricsRequest
	GetNextToken() *string
	SetStartTime(v int64) *ModelRouterQueryCostOverviewMetricsRequest
	GetStartTime() *int64
}

type ModelRouterQueryCostOverviewMetricsRequest struct {
	// Optional. Filters results by API Key ID. This parameter is linked to the department and requires clientId to be specified first.
	//
	// example:
	//
	// 100
	ApiKeyId *int64 `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	// Filters results by department ID.
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
	// The end time, in UNIX timestamp (seconds).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1700086400
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Automatically aggregated. No input required. The granularity. Valid values: hourly and daily. Default value: hourly.
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
	// Optional. Filters results by member IDs, separated by commas. If not specified, the department and all its members are included. If an empty value is specified, only the department is included without members.
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
	// The start time, in UNIX timestamp (seconds).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1700000000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ModelRouterQueryCostOverviewMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryCostOverviewMetricsRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryCostOverviewMetricsRequest) GetApiKeyId() *int64 {
	return s.ApiKeyId
}

func (s *ModelRouterQueryCostOverviewMetricsRequest) GetClientId() *int64 {
	return s.ClientId
}

func (s *ModelRouterQueryCostOverviewMetricsRequest) GetClientIds() *string {
	return s.ClientIds
}

func (s *ModelRouterQueryCostOverviewMetricsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ModelRouterQueryCostOverviewMetricsRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *ModelRouterQueryCostOverviewMetricsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryCostOverviewMetricsRequest) GetMemberUserIds() *string {
	return s.MemberUserIds
}

func (s *ModelRouterQueryCostOverviewMetricsRequest) GetModelTypes() *string {
	return s.ModelTypes
}

func (s *ModelRouterQueryCostOverviewMetricsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryCostOverviewMetricsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ModelRouterQueryCostOverviewMetricsRequest) SetApiKeyId(v int64) *ModelRouterQueryCostOverviewMetricsRequest {
	s.ApiKeyId = &v
	return s
}

func (s *ModelRouterQueryCostOverviewMetricsRequest) SetClientId(v int64) *ModelRouterQueryCostOverviewMetricsRequest {
	s.ClientId = &v
	return s
}

func (s *ModelRouterQueryCostOverviewMetricsRequest) SetClientIds(v string) *ModelRouterQueryCostOverviewMetricsRequest {
	s.ClientIds = &v
	return s
}

func (s *ModelRouterQueryCostOverviewMetricsRequest) SetEndTime(v int64) *ModelRouterQueryCostOverviewMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *ModelRouterQueryCostOverviewMetricsRequest) SetGranularity(v string) *ModelRouterQueryCostOverviewMetricsRequest {
	s.Granularity = &v
	return s
}

func (s *ModelRouterQueryCostOverviewMetricsRequest) SetMaxResults(v int32) *ModelRouterQueryCostOverviewMetricsRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryCostOverviewMetricsRequest) SetMemberUserIds(v string) *ModelRouterQueryCostOverviewMetricsRequest {
	s.MemberUserIds = &v
	return s
}

func (s *ModelRouterQueryCostOverviewMetricsRequest) SetModelTypes(v string) *ModelRouterQueryCostOverviewMetricsRequest {
	s.ModelTypes = &v
	return s
}

func (s *ModelRouterQueryCostOverviewMetricsRequest) SetNextToken(v string) *ModelRouterQueryCostOverviewMetricsRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryCostOverviewMetricsRequest) SetStartTime(v int64) *ModelRouterQueryCostOverviewMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *ModelRouterQueryCostOverviewMetricsRequest) Validate() error {
	return dara.Validate(s)
}
