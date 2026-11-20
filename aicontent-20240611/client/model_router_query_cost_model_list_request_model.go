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
	SetClientIds(v string) *ModelRouterQueryCostModelListRequest
	GetClientIds() *string
	SetEndTime(v int64) *ModelRouterQueryCostModelListRequest
	GetEndTime() *int64
	SetGranularity(v string) *ModelRouterQueryCostModelListRequest
	GetGranularity() *string
	SetMaxResults(v int32) *ModelRouterQueryCostModelListRequest
	GetMaxResults() *int32
	SetMemberUserIds(v string) *ModelRouterQueryCostModelListRequest
	GetMemberUserIds() *string
	SetModelIds(v string) *ModelRouterQueryCostModelListRequest
	GetModelIds() *string
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
	// The API key ID used for filtering. This parameter is optional and associated with the department. You must specify clientId first.
	//
	// example:
	//
	// 100
	ApiKeyId *int64 `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	// The department ID used for filtering.
	//
	// example:
	//
	// 1
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// The department IDs. Separate multiple IDs with commas to query data for multiple departments. This parameter is mutually exclusive with clientId.
	//
	// example:
	//
	// 1,2,3
	ClientIds *string `json:"clientIds,omitempty" xml:"clientIds,omitempty"`
	// The end time, specified as a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1700086400
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The time granularity for automatic aggregation. You do not need to pass this parameter. Valid values: hourly and daily. Default value: hourly.
	//
	// example:
	//
	// hourly
	Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
	// The maximum number of results.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The member IDs used for filtering. Separate multiple IDs with commas. This parameter is optional. If you do not specify this parameter, the department and all its members are included. If you pass an empty value, only the department is included, excluding its members.
	//
	// example:
	//
	// 30001,30002
	MemberUserIds *string `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty"`
	// The model IDs used for filtering. Separate multiple positive integer IDs with commas.
	//
	// example:
	//
	// 123,456
	ModelIds *string `json:"modelIds,omitempty" xml:"modelIds,omitempty"`
	// The model types. Separate multiple values with commas.
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
	// The keyword for fuzzy search by model name or code.
	//
	// example:
	//
	// qwen
	Search *string `json:"search,omitempty" xml:"search,omitempty"`
	// The start time, specified as a UNIX timestamp in seconds.
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

func (s *ModelRouterQueryCostModelListRequest) GetClientIds() *string {
	return s.ClientIds
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

func (s *ModelRouterQueryCostModelListRequest) GetModelIds() *string {
	return s.ModelIds
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

func (s *ModelRouterQueryCostModelListRequest) SetClientIds(v string) *ModelRouterQueryCostModelListRequest {
	s.ClientIds = &v
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

func (s *ModelRouterQueryCostModelListRequest) SetModelIds(v string) *ModelRouterQueryCostModelListRequest {
	s.ModelIds = &v
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
