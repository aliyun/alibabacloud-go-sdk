// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryBillingCostBreakdownRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v int64) *ModelRouterQueryBillingCostBreakdownRequest
	GetApiKeyId() *int64
	SetClientId(v int64) *ModelRouterQueryBillingCostBreakdownRequest
	GetClientId() *int64
	SetEndTime(v int64) *ModelRouterQueryBillingCostBreakdownRequest
	GetEndTime() *int64
	SetGranularity(v string) *ModelRouterQueryBillingCostBreakdownRequest
	GetGranularity() *string
	SetMaxResults(v int32) *ModelRouterQueryBillingCostBreakdownRequest
	GetMaxResults() *int32
	SetMemberUserIds(v string) *ModelRouterQueryBillingCostBreakdownRequest
	GetMemberUserIds() *string
	SetModelId(v int64) *ModelRouterQueryBillingCostBreakdownRequest
	GetModelId() *int64
	SetModelTypes(v string) *ModelRouterQueryBillingCostBreakdownRequest
	GetModelTypes() *string
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
	// Optional. Filters results by API key ID. This parameter is linked with the department. Specify clientId first.
	//
	// example:
	//
	// 100
	ApiKeyId *int64 `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	// Optional. Filters results by department ID.
	//
	// example:
	//
	// 5
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// The query end time, in UNIX timestamp (seconds).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1700086400
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The aggregation granularity. Valid values:
	//
	// - hourly
	//
	// - daily
	//
	// This parameter is required.
	//
	// example:
	//
	// hourly
	Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
	// The maximum number of results to return.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// Optional. Filters results by member ID. Separate multiple values with commas. If not specified, the department and all its members are included. If an empty value is specified, only the department is included without members.
	//
	// example:
	//
	// 30001,30002
	MemberUserIds *string `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty"`
	// Optional. Filters results by model ID.
	//
	// example:
	//
	// 12
	ModelId *int64 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// Optional. Filters results by model type. Separate multiple values with commas.
	//
	// example:
	//
	// Chat
	ModelTypes *string `json:"modelTypes,omitempty" xml:"modelTypes,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// xxxx-xxx-xxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 500.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The query start time, in UNIX timestamp (seconds).
	//
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

func (s *ModelRouterQueryBillingCostBreakdownRequest) GetApiKeyId() *int64 {
	return s.ApiKeyId
}

func (s *ModelRouterQueryBillingCostBreakdownRequest) GetClientId() *int64 {
	return s.ClientId
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

func (s *ModelRouterQueryBillingCostBreakdownRequest) GetMemberUserIds() *string {
	return s.MemberUserIds
}

func (s *ModelRouterQueryBillingCostBreakdownRequest) GetModelId() *int64 {
	return s.ModelId
}

func (s *ModelRouterQueryBillingCostBreakdownRequest) GetModelTypes() *string {
	return s.ModelTypes
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

func (s *ModelRouterQueryBillingCostBreakdownRequest) SetApiKeyId(v int64) *ModelRouterQueryBillingCostBreakdownRequest {
	s.ApiKeyId = &v
	return s
}

func (s *ModelRouterQueryBillingCostBreakdownRequest) SetClientId(v int64) *ModelRouterQueryBillingCostBreakdownRequest {
	s.ClientId = &v
	return s
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

func (s *ModelRouterQueryBillingCostBreakdownRequest) SetMemberUserIds(v string) *ModelRouterQueryBillingCostBreakdownRequest {
	s.MemberUserIds = &v
	return s
}

func (s *ModelRouterQueryBillingCostBreakdownRequest) SetModelId(v int64) *ModelRouterQueryBillingCostBreakdownRequest {
	s.ModelId = &v
	return s
}

func (s *ModelRouterQueryBillingCostBreakdownRequest) SetModelTypes(v string) *ModelRouterQueryBillingCostBreakdownRequest {
	s.ModelTypes = &v
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
