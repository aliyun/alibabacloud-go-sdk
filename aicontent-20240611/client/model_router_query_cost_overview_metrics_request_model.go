// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryCostOverviewMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v int64) *ModelRouterQueryCostOverviewMetricsRequest
	GetClientId() *int64
	SetEndTime(v int64) *ModelRouterQueryCostOverviewMetricsRequest
	GetEndTime() *int64
	SetGranularity(v string) *ModelRouterQueryCostOverviewMetricsRequest
	GetGranularity() *string
	SetMaxResults(v int32) *ModelRouterQueryCostOverviewMetricsRequest
	GetMaxResults() *int32
	SetModelTypes(v string) *ModelRouterQueryCostOverviewMetricsRequest
	GetModelTypes() *string
	SetNextToken(v string) *ModelRouterQueryCostOverviewMetricsRequest
	GetNextToken() *string
	SetStartTime(v int64) *ModelRouterQueryCostOverviewMetricsRequest
	GetStartTime() *int64
}

type ModelRouterQueryCostOverviewMetricsRequest struct {
	// example:
	//
	// 1
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1700086400
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// hourly
	Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
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

func (s *ModelRouterQueryCostOverviewMetricsRequest) GetClientId() *int64 {
	return s.ClientId
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

func (s *ModelRouterQueryCostOverviewMetricsRequest) GetModelTypes() *string {
	return s.ModelTypes
}

func (s *ModelRouterQueryCostOverviewMetricsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryCostOverviewMetricsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ModelRouterQueryCostOverviewMetricsRequest) SetClientId(v int64) *ModelRouterQueryCostOverviewMetricsRequest {
	s.ClientId = &v
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
