// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterGetBillingBillSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v int64) *ModelRouterGetBillingBillSummaryRequest
	GetApiKeyId() *int64
	SetClientId(v int64) *ModelRouterGetBillingBillSummaryRequest
	GetClientId() *int64
	SetClientIds(v string) *ModelRouterGetBillingBillSummaryRequest
	GetClientIds() *string
	SetEndTime(v int64) *ModelRouterGetBillingBillSummaryRequest
	GetEndTime() *int64
	SetMaxResults(v int32) *ModelRouterGetBillingBillSummaryRequest
	GetMaxResults() *int32
	SetMemberUserIds(v string) *ModelRouterGetBillingBillSummaryRequest
	GetMemberUserIds() *string
	SetModelId(v int64) *ModelRouterGetBillingBillSummaryRequest
	GetModelId() *int64
	SetModelTypes(v string) *ModelRouterGetBillingBillSummaryRequest
	GetModelTypes() *string
	SetNextToken(v string) *ModelRouterGetBillingBillSummaryRequest
	GetNextToken() *string
	SetStartTime(v int64) *ModelRouterGetBillingBillSummaryRequest
	GetStartTime() *int64
}

type ModelRouterGetBillingBillSummaryRequest struct {
	// The API key ID used to filter results. This parameter is optional and linked to the department. You must specify clientId first.
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
	// The end time, in UNIX timestamp format (seconds).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1700086400
	EndTime    *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The member IDs used to filter results, separated by commas. This parameter is optional. If not specified, the query returns data for the department and all its members. If an empty value is specified, the query returns data for the department only, excluding members.
	//
	// example:
	//
	// 30001,30002
	MemberUserIds *string `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty"`
	// The model ID. This parameter is optional and used to filter by model.
	//
	// example:
	//
	// 1
	ModelId *int64 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// The model types, separated by commas.
	//
	// example:
	//
	// Chat,ChatMultimodal
	ModelTypes *string `json:"modelTypes,omitempty" xml:"modelTypes,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The start time, in UNIX timestamp format (seconds).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1700000000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ModelRouterGetBillingBillSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetBillingBillSummaryRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterGetBillingBillSummaryRequest) GetApiKeyId() *int64 {
	return s.ApiKeyId
}

func (s *ModelRouterGetBillingBillSummaryRequest) GetClientId() *int64 {
	return s.ClientId
}

func (s *ModelRouterGetBillingBillSummaryRequest) GetClientIds() *string {
	return s.ClientIds
}

func (s *ModelRouterGetBillingBillSummaryRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ModelRouterGetBillingBillSummaryRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterGetBillingBillSummaryRequest) GetMemberUserIds() *string {
	return s.MemberUserIds
}

func (s *ModelRouterGetBillingBillSummaryRequest) GetModelId() *int64 {
	return s.ModelId
}

func (s *ModelRouterGetBillingBillSummaryRequest) GetModelTypes() *string {
	return s.ModelTypes
}

func (s *ModelRouterGetBillingBillSummaryRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterGetBillingBillSummaryRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ModelRouterGetBillingBillSummaryRequest) SetApiKeyId(v int64) *ModelRouterGetBillingBillSummaryRequest {
	s.ApiKeyId = &v
	return s
}

func (s *ModelRouterGetBillingBillSummaryRequest) SetClientId(v int64) *ModelRouterGetBillingBillSummaryRequest {
	s.ClientId = &v
	return s
}

func (s *ModelRouterGetBillingBillSummaryRequest) SetClientIds(v string) *ModelRouterGetBillingBillSummaryRequest {
	s.ClientIds = &v
	return s
}

func (s *ModelRouterGetBillingBillSummaryRequest) SetEndTime(v int64) *ModelRouterGetBillingBillSummaryRequest {
	s.EndTime = &v
	return s
}

func (s *ModelRouterGetBillingBillSummaryRequest) SetMaxResults(v int32) *ModelRouterGetBillingBillSummaryRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterGetBillingBillSummaryRequest) SetMemberUserIds(v string) *ModelRouterGetBillingBillSummaryRequest {
	s.MemberUserIds = &v
	return s
}

func (s *ModelRouterGetBillingBillSummaryRequest) SetModelId(v int64) *ModelRouterGetBillingBillSummaryRequest {
	s.ModelId = &v
	return s
}

func (s *ModelRouterGetBillingBillSummaryRequest) SetModelTypes(v string) *ModelRouterGetBillingBillSummaryRequest {
	s.ModelTypes = &v
	return s
}

func (s *ModelRouterGetBillingBillSummaryRequest) SetNextToken(v string) *ModelRouterGetBillingBillSummaryRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterGetBillingBillSummaryRequest) SetStartTime(v int64) *ModelRouterGetBillingBillSummaryRequest {
	s.StartTime = &v
	return s
}

func (s *ModelRouterGetBillingBillSummaryRequest) Validate() error {
	return dara.Validate(s)
}
