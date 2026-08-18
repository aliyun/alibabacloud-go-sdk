// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryCostModelDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v int64) *ModelRouterQueryCostModelDetailRequest
	GetApiKeyId() *int64
	SetClientId(v int64) *ModelRouterQueryCostModelDetailRequest
	GetClientId() *int64
	SetClientIds(v string) *ModelRouterQueryCostModelDetailRequest
	GetClientIds() *string
	SetEndTime(v int64) *ModelRouterQueryCostModelDetailRequest
	GetEndTime() *int64
	SetMaxResults(v int32) *ModelRouterQueryCostModelDetailRequest
	GetMaxResults() *int32
	SetMemberUserIds(v string) *ModelRouterQueryCostModelDetailRequest
	GetMemberUserIds() *string
	SetModelId(v int64) *ModelRouterQueryCostModelDetailRequest
	GetModelId() *int64
	SetNextToken(v string) *ModelRouterQueryCostModelDetailRequest
	GetNextToken() *string
	SetPage(v int32) *ModelRouterQueryCostModelDetailRequest
	GetPage() *int32
	SetPageIndex(v int32) *ModelRouterQueryCostModelDetailRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterQueryCostModelDetailRequest
	GetPageSize() *int32
	SetStartTime(v int64) *ModelRouterQueryCostModelDetailRequest
	GetStartTime() *int64
}

type ModelRouterQueryCostModelDetailRequest struct {
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
	// The end time, in UNIX timestamp (seconds).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1700086400
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// maxResults
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// Optional. Filters by member IDs, separated by commas. If not specified, the department and all its members are included. If an empty value is specified, only the department is included without members.
	//
	// example:
	//
	// 30001,30002
	MemberUserIds *string `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty"`
	// The model ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ModelId *int64 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// nextToken
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
	// The page number. This parameter takes priority over the page parameter.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The start time, in UNIX timestamp (seconds).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1700000000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ModelRouterQueryCostModelDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryCostModelDetailRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryCostModelDetailRequest) GetApiKeyId() *int64 {
	return s.ApiKeyId
}

func (s *ModelRouterQueryCostModelDetailRequest) GetClientId() *int64 {
	return s.ClientId
}

func (s *ModelRouterQueryCostModelDetailRequest) GetClientIds() *string {
	return s.ClientIds
}

func (s *ModelRouterQueryCostModelDetailRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ModelRouterQueryCostModelDetailRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryCostModelDetailRequest) GetMemberUserIds() *string {
	return s.MemberUserIds
}

func (s *ModelRouterQueryCostModelDetailRequest) GetModelId() *int64 {
	return s.ModelId
}

func (s *ModelRouterQueryCostModelDetailRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryCostModelDetailRequest) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterQueryCostModelDetailRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterQueryCostModelDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryCostModelDetailRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ModelRouterQueryCostModelDetailRequest) SetApiKeyId(v int64) *ModelRouterQueryCostModelDetailRequest {
	s.ApiKeyId = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailRequest) SetClientId(v int64) *ModelRouterQueryCostModelDetailRequest {
	s.ClientId = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailRequest) SetClientIds(v string) *ModelRouterQueryCostModelDetailRequest {
	s.ClientIds = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailRequest) SetEndTime(v int64) *ModelRouterQueryCostModelDetailRequest {
	s.EndTime = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailRequest) SetMaxResults(v int32) *ModelRouterQueryCostModelDetailRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailRequest) SetMemberUserIds(v string) *ModelRouterQueryCostModelDetailRequest {
	s.MemberUserIds = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailRequest) SetModelId(v int64) *ModelRouterQueryCostModelDetailRequest {
	s.ModelId = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailRequest) SetNextToken(v string) *ModelRouterQueryCostModelDetailRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailRequest) SetPage(v int32) *ModelRouterQueryCostModelDetailRequest {
	s.Page = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailRequest) SetPageIndex(v int32) *ModelRouterQueryCostModelDetailRequest {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailRequest) SetPageSize(v int32) *ModelRouterQueryCostModelDetailRequest {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailRequest) SetStartTime(v int64) *ModelRouterQueryCostModelDetailRequest {
	s.StartTime = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailRequest) Validate() error {
	return dara.Validate(s)
}
