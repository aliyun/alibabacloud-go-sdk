// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryCostModelListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v int64) *ModelRouterQueryCostModelListRequest
	GetApiKeyId() *int64
	SetClientId(v int64) *ModelRouterQueryCostModelListRequest
	GetClientId() *int64
	SetEndTime(v int64) *ModelRouterQueryCostModelListRequest
	GetEndTime() *int64
	SetGranularity(v string) *ModelRouterQueryCostModelListRequest
	GetGranularity() *string
	SetMaxResults(v int32) *ModelRouterQueryCostModelListRequest
	GetMaxResults() *int32
	SetMemberUserIds(v string) *ModelRouterQueryCostModelListRequest
	GetMemberUserIds() *string
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
	// Optional. Filters results by API key ID. This parameter works in conjunction with the department and requires clientId to be specified first.
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
	// The end time, as a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1700086400
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Automatic aggregation. You do not need to pass this parameter. Granularity: hourly/daily. Default value: hourly.
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
	// Optional. Filters results by member IDs, separated by commas. If not specified, the department and all its members are included. If an empty value is passed, only the department is included without members.
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
	// Performs a fuzzy match on the model name or code.
	//
	// example:
	//
	// qwen
	Search *string `json:"search,omitempty" xml:"search,omitempty"`
	// The start time, as a UNIX timestamp in seconds.
	//
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

func (s *ModelRouterQueryCostModelListRequest) GetApiKeyId() *int64 {
	return s.ApiKeyId
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

func (s *ModelRouterQueryCostModelListRequest) GetMemberUserIds() *string {
	return s.MemberUserIds
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

func (s *ModelRouterQueryCostModelListRequest) SetApiKeyId(v int64) *ModelRouterQueryCostModelListRequest {
	s.ApiKeyId = &v
	return s
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

func (s *ModelRouterQueryCostModelListRequest) SetMemberUserIds(v string) *ModelRouterQueryCostModelListRequest {
	s.MemberUserIds = &v
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
