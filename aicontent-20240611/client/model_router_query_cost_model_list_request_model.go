// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryCostModelListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v int64) *ModelRouterQueryCostModelListRequest
	GetClientId() *int64
	SetEndTime(v int64) *ModelRouterQueryCostModelListRequest
	GetEndTime() *int64
	SetGranularity(v string) *ModelRouterQueryCostModelListRequest
	GetGranularity() *string
	SetMaxResults(v int32) *ModelRouterQueryCostModelListRequest
	GetMaxResults() *int32
	SetModelTypes(v string) *ModelRouterQueryCostModelListRequest
	GetModelTypes() *string
	SetNextToken(v string) *ModelRouterQueryCostModelListRequest
	GetNextToken() *string
	SetSearch(v string) *ModelRouterQueryCostModelListRequest
	GetSearch() *string
	SetStartTime(v int64) *ModelRouterQueryCostModelListRequest
	GetStartTime() *int64
}

type ModelRouterQueryCostModelListRequest struct {
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
	// example:
	//
	// qwen
	Search *string `json:"search,omitempty" xml:"search,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1700000000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ModelRouterQueryCostModelListRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryCostModelListRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryCostModelListRequest) GetClientId() *int64 {
	return s.ClientId
}

func (s *ModelRouterQueryCostModelListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ModelRouterQueryCostModelListRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *ModelRouterQueryCostModelListRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryCostModelListRequest) GetModelTypes() *string {
	return s.ModelTypes
}

func (s *ModelRouterQueryCostModelListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryCostModelListRequest) GetSearch() *string {
	return s.Search
}

func (s *ModelRouterQueryCostModelListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ModelRouterQueryCostModelListRequest) SetClientId(v int64) *ModelRouterQueryCostModelListRequest {
	s.ClientId = &v
	return s
}

func (s *ModelRouterQueryCostModelListRequest) SetEndTime(v int64) *ModelRouterQueryCostModelListRequest {
	s.EndTime = &v
	return s
}

func (s *ModelRouterQueryCostModelListRequest) SetGranularity(v string) *ModelRouterQueryCostModelListRequest {
	s.Granularity = &v
	return s
}

func (s *ModelRouterQueryCostModelListRequest) SetMaxResults(v int32) *ModelRouterQueryCostModelListRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryCostModelListRequest) SetModelTypes(v string) *ModelRouterQueryCostModelListRequest {
	s.ModelTypes = &v
	return s
}

func (s *ModelRouterQueryCostModelListRequest) SetNextToken(v string) *ModelRouterQueryCostModelListRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryCostModelListRequest) SetSearch(v string) *ModelRouterQueryCostModelListRequest {
	s.Search = &v
	return s
}

func (s *ModelRouterQueryCostModelListRequest) SetStartTime(v int64) *ModelRouterQueryCostModelListRequest {
	s.StartTime = &v
	return s
}

func (s *ModelRouterQueryCostModelListRequest) Validate() error {
	return dara.Validate(s)
}
