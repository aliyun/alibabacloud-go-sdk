// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryBillingDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v int64) *ModelRouterQueryBillingDetailsRequest
	GetApiKeyId() *int64
	SetClientId(v int64) *ModelRouterQueryBillingDetailsRequest
	GetClientId() *int64
	SetClientIds(v string) *ModelRouterQueryBillingDetailsRequest
	GetClientIds() *string
	SetEndTime(v int64) *ModelRouterQueryBillingDetailsRequest
	GetEndTime() *int64
	SetModelCodes(v string) *ModelRouterQueryBillingDetailsRequest
	GetModelCodes() *string
	SetModelId(v int64) *ModelRouterQueryBillingDetailsRequest
	GetModelId() *int64
	SetModelTypes(v string) *ModelRouterQueryBillingDetailsRequest
	GetModelTypes() *string
	SetPage(v int32) *ModelRouterQueryBillingDetailsRequest
	GetPage() *int32
	SetPageSize(v int32) *ModelRouterQueryBillingDetailsRequest
	GetPageSize() *int32
	SetRequestId(v string) *ModelRouterQueryBillingDetailsRequest
	GetRequestId() *string
	SetStartTime(v int64) *ModelRouterQueryBillingDetailsRequest
	GetStartTime() *int64
}

type ModelRouterQueryBillingDetailsRequest struct {
	// Optional. Filters results by API Key ID.
	//
	// example:
	//
	// 100
	ApiKeyId *int64 `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	// Optional. Filters results by department ID (single value).
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
	// The query end time, in UNIX timestamp (seconds).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1700086400
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Optional. Filters results by model code. Separate multiple values with commas.
	//
	// example:
	//
	// qwen-plus,qwen-max
	ModelCodes *string `json:"modelCodes,omitempty" xml:"modelCodes,omitempty"`
	// Optional. Filters results by model ID.
	//
	// example:
	//
	// 1
	ModelId *int64 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// Optional. Filters results by model type. Separate multiple values with commas.
	//
	// example:
	//
	// Chat
	ModelTypes *string `json:"modelTypes,omitempty" xml:"modelTypes,omitempty"`
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
	// Optional. Filters results by exact match of the request ID.
	//
	// example:
	//
	// chatcmpl-abc123def456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The query start time, in UNIX timestamp (seconds).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1700000000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ModelRouterQueryBillingDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryBillingDetailsRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryBillingDetailsRequest) GetApiKeyId() *int64 {
	return s.ApiKeyId
}

func (s *ModelRouterQueryBillingDetailsRequest) GetClientId() *int64 {
	return s.ClientId
}

func (s *ModelRouterQueryBillingDetailsRequest) GetClientIds() *string {
	return s.ClientIds
}

func (s *ModelRouterQueryBillingDetailsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ModelRouterQueryBillingDetailsRequest) GetModelCodes() *string {
	return s.ModelCodes
}

func (s *ModelRouterQueryBillingDetailsRequest) GetModelId() *int64 {
	return s.ModelId
}

func (s *ModelRouterQueryBillingDetailsRequest) GetModelTypes() *string {
	return s.ModelTypes
}

func (s *ModelRouterQueryBillingDetailsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterQueryBillingDetailsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryBillingDetailsRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryBillingDetailsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ModelRouterQueryBillingDetailsRequest) SetApiKeyId(v int64) *ModelRouterQueryBillingDetailsRequest {
	s.ApiKeyId = &v
	return s
}

func (s *ModelRouterQueryBillingDetailsRequest) SetClientId(v int64) *ModelRouterQueryBillingDetailsRequest {
	s.ClientId = &v
	return s
}

func (s *ModelRouterQueryBillingDetailsRequest) SetClientIds(v string) *ModelRouterQueryBillingDetailsRequest {
	s.ClientIds = &v
	return s
}

func (s *ModelRouterQueryBillingDetailsRequest) SetEndTime(v int64) *ModelRouterQueryBillingDetailsRequest {
	s.EndTime = &v
	return s
}

func (s *ModelRouterQueryBillingDetailsRequest) SetModelCodes(v string) *ModelRouterQueryBillingDetailsRequest {
	s.ModelCodes = &v
	return s
}

func (s *ModelRouterQueryBillingDetailsRequest) SetModelId(v int64) *ModelRouterQueryBillingDetailsRequest {
	s.ModelId = &v
	return s
}

func (s *ModelRouterQueryBillingDetailsRequest) SetModelTypes(v string) *ModelRouterQueryBillingDetailsRequest {
	s.ModelTypes = &v
	return s
}

func (s *ModelRouterQueryBillingDetailsRequest) SetPage(v int32) *ModelRouterQueryBillingDetailsRequest {
	s.Page = &v
	return s
}

func (s *ModelRouterQueryBillingDetailsRequest) SetPageSize(v int32) *ModelRouterQueryBillingDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryBillingDetailsRequest) SetRequestId(v string) *ModelRouterQueryBillingDetailsRequest {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryBillingDetailsRequest) SetStartTime(v int64) *ModelRouterQueryBillingDetailsRequest {
	s.StartTime = &v
	return s
}

func (s *ModelRouterQueryBillingDetailsRequest) Validate() error {
	return dara.Validate(s)
}
