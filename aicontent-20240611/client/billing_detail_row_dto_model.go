// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBillingDetailRowDTO interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v float64) *BillingDetailRowDTO
	GetAmount() *float64
	SetApiKeyId(v int64) *BillingDetailRowDTO
	GetApiKeyId() *int64
	SetApiKeyName(v string) *BillingDetailRowDTO
	GetApiKeyName() *string
	SetCacheCreationTokens(v float64) *BillingDetailRowDTO
	GetCacheCreationTokens() *float64
	SetCachedTokens(v float64) *BillingDetailRowDTO
	GetCachedTokens() *float64
	SetClientId(v int64) *BillingDetailRowDTO
	GetClientId() *int64
	SetClientName(v string) *BillingDetailRowDTO
	GetClientName() *string
	SetDiscount(v float64) *BillingDetailRowDTO
	GetDiscount() *float64
	SetInputTokens(v float64) *BillingDetailRowDTO
	GetInputTokens() *float64
	SetMemberUserId(v int64) *BillingDetailRowDTO
	GetMemberUserId() *int64
	SetMemberUserName(v string) *BillingDetailRowDTO
	GetMemberUserName() *string
	SetMetrics(v string) *BillingDetailRowDTO
	GetMetrics() *string
	SetModelCode(v string) *BillingDetailRowDTO
	GetModelCode() *string
	SetModelId(v int64) *BillingDetailRowDTO
	GetModelId() *int64
	SetModelName(v string) *BillingDetailRowDTO
	GetModelName() *string
	SetModelSymbol(v string) *BillingDetailRowDTO
	GetModelSymbol() *string
	SetModelType(v string) *BillingDetailRowDTO
	GetModelType() *string
	SetModelVersion(v int32) *BillingDetailRowDTO
	GetModelVersion() *int32
	SetOutputTokens(v float64) *BillingDetailRowDTO
	GetOutputTokens() *float64
	SetReasoningTokens(v float64) *BillingDetailRowDTO
	GetReasoningTokens() *float64
	SetRequestId(v string) *BillingDetailRowDTO
	GetRequestId() *string
	SetRequestTime(v int64) *BillingDetailRowDTO
	GetRequestTime() *int64
	SetTotalTokens(v float64) *BillingDetailRowDTO
	GetTotalTokens() *float64
	SetUsageDetail(v string) *BillingDetailRowDTO
	GetUsageDetail() *string
}

type BillingDetailRowDTO struct {
	// The actual payment amount (after discount), rounded to 8 decimal places.
	//
	// example:
	//
	// 0.00012800
	Amount *float64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// API Key ID
	//
	// example:
	//
	// 100
	ApiKeyId *int64 `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	// The API key name.
	//
	// example:
	//
	// Default Key
	ApiKeyName *string `json:"apiKeyName,omitempty" xml:"apiKeyName,omitempty"`
	// The number of cache creation tokens (explicit cache writes).
	//
	// example:
	//
	// 0
	CacheCreationTokens *float64 `json:"cacheCreationTokens,omitempty" xml:"cacheCreationTokens,omitempty"`
	// The number of tokens that hit the cache.
	//
	// example:
	//
	// 256
	CachedTokens *float64 `json:"cachedTokens,omitempty" xml:"cachedTokens,omitempty"`
	// The department ID. A value of 0 indicates that no department is associated.
	//
	// example:
	//
	// 1
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// The department name.
	//
	// example:
	//
	// R&D Department
	ClientName *string `json:"clientName,omitempty" xml:"clientName,omitempty"`
	// The discount coefficient. A value of 1.0 indicates no discount.
	//
	// example:
	//
	// 1.0
	Discount *float64 `json:"discount,omitempty" xml:"discount,omitempty"`
	// The number of input tokens, including cached tokens and cache creation tokens.
	//
	// example:
	//
	// 1024
	InputTokens *float64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// The member user ID for a member row. The value is 0 for a department row.
	//
	// example:
	//
	// 30001
	MemberUserId *int64 `json:"memberUserId,omitempty" xml:"memberUserId,omitempty"`
	// The member name for a member row. The value is empty for a department row.
	//
	// example:
	//
	// John
	MemberUserName *string `json:"memberUserName,omitempty" xml:"memberUserName,omitempty"`
	// The JSON of other metering field mapping, such as video duration and image count. Fields with a value of 0 are not included in the output.
	//
	// example:
	//
	// {}
	Metrics *string `json:"metrics,omitempty" xml:"metrics,omitempty"`
	// The model identifier.
	//
	// example:
	//
	// qwen-plus
	ModelCode *string `json:"modelCode,omitempty" xml:"modelCode,omitempty"`
	// The model ID.
	//
	// example:
	//
	// 1
	ModelId *int64 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// The model name.
	//
	// example:
	//
	// Qwen-Plus
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// The model symbol (provider identifier).
	//
	// example:
	//
	// qwen
	ModelSymbol *string `json:"modelSymbol,omitempty" xml:"modelSymbol,omitempty"`
	// The model type.
	//
	// example:
	//
	// Chat
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// The model version number.
	//
	// example:
	//
	// 1
	ModelVersion *int32 `json:"modelVersion,omitempty" xml:"modelVersion,omitempty"`
	// The number of output tokens.
	//
	// example:
	//
	// 512
	OutputTokens *float64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// The number of reasoning tokens.
	//
	// example:
	//
	// 128
	ReasoningTokens *float64 `json:"reasoningTokens,omitempty" xml:"reasoningTokens,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// chatcmpl-abc123def456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The request time as a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1700000000
	RequestTime *int64 `json:"requestTime,omitempty" xml:"requestTime,omitempty"`
	// The total number of tokens.
	//
	// example:
	//
	// 1536
	TotalTokens *float64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
	// The raw JSON of the usage details.
	//
	// example:
	//
	// {"input_tokens": 1024, "output_tokens": 512}
	UsageDetail *string `json:"usageDetail,omitempty" xml:"usageDetail,omitempty"`
}

func (s BillingDetailRowDTO) String() string {
	return dara.Prettify(s)
}

func (s BillingDetailRowDTO) GoString() string {
	return s.String()
}

func (s *BillingDetailRowDTO) GetAmount() *float64 {
	return s.Amount
}

func (s *BillingDetailRowDTO) GetApiKeyId() *int64 {
	return s.ApiKeyId
}

func (s *BillingDetailRowDTO) GetApiKeyName() *string {
	return s.ApiKeyName
}

func (s *BillingDetailRowDTO) GetCacheCreationTokens() *float64 {
	return s.CacheCreationTokens
}

func (s *BillingDetailRowDTO) GetCachedTokens() *float64 {
	return s.CachedTokens
}

func (s *BillingDetailRowDTO) GetClientId() *int64 {
	return s.ClientId
}

func (s *BillingDetailRowDTO) GetClientName() *string {
	return s.ClientName
}

func (s *BillingDetailRowDTO) GetDiscount() *float64 {
	return s.Discount
}

func (s *BillingDetailRowDTO) GetInputTokens() *float64 {
	return s.InputTokens
}

func (s *BillingDetailRowDTO) GetMemberUserId() *int64 {
	return s.MemberUserId
}

func (s *BillingDetailRowDTO) GetMemberUserName() *string {
	return s.MemberUserName
}

func (s *BillingDetailRowDTO) GetMetrics() *string {
	return s.Metrics
}

func (s *BillingDetailRowDTO) GetModelCode() *string {
	return s.ModelCode
}

func (s *BillingDetailRowDTO) GetModelId() *int64 {
	return s.ModelId
}

func (s *BillingDetailRowDTO) GetModelName() *string {
	return s.ModelName
}

func (s *BillingDetailRowDTO) GetModelSymbol() *string {
	return s.ModelSymbol
}

func (s *BillingDetailRowDTO) GetModelType() *string {
	return s.ModelType
}

func (s *BillingDetailRowDTO) GetModelVersion() *int32 {
	return s.ModelVersion
}

func (s *BillingDetailRowDTO) GetOutputTokens() *float64 {
	return s.OutputTokens
}

func (s *BillingDetailRowDTO) GetReasoningTokens() *float64 {
	return s.ReasoningTokens
}

func (s *BillingDetailRowDTO) GetRequestId() *string {
	return s.RequestId
}

func (s *BillingDetailRowDTO) GetRequestTime() *int64 {
	return s.RequestTime
}

func (s *BillingDetailRowDTO) GetTotalTokens() *float64 {
	return s.TotalTokens
}

func (s *BillingDetailRowDTO) GetUsageDetail() *string {
	return s.UsageDetail
}

func (s *BillingDetailRowDTO) SetAmount(v float64) *BillingDetailRowDTO {
	s.Amount = &v
	return s
}

func (s *BillingDetailRowDTO) SetApiKeyId(v int64) *BillingDetailRowDTO {
	s.ApiKeyId = &v
	return s
}

func (s *BillingDetailRowDTO) SetApiKeyName(v string) *BillingDetailRowDTO {
	s.ApiKeyName = &v
	return s
}

func (s *BillingDetailRowDTO) SetCacheCreationTokens(v float64) *BillingDetailRowDTO {
	s.CacheCreationTokens = &v
	return s
}

func (s *BillingDetailRowDTO) SetCachedTokens(v float64) *BillingDetailRowDTO {
	s.CachedTokens = &v
	return s
}

func (s *BillingDetailRowDTO) SetClientId(v int64) *BillingDetailRowDTO {
	s.ClientId = &v
	return s
}

func (s *BillingDetailRowDTO) SetClientName(v string) *BillingDetailRowDTO {
	s.ClientName = &v
	return s
}

func (s *BillingDetailRowDTO) SetDiscount(v float64) *BillingDetailRowDTO {
	s.Discount = &v
	return s
}

func (s *BillingDetailRowDTO) SetInputTokens(v float64) *BillingDetailRowDTO {
	s.InputTokens = &v
	return s
}

func (s *BillingDetailRowDTO) SetMemberUserId(v int64) *BillingDetailRowDTO {
	s.MemberUserId = &v
	return s
}

func (s *BillingDetailRowDTO) SetMemberUserName(v string) *BillingDetailRowDTO {
	s.MemberUserName = &v
	return s
}

func (s *BillingDetailRowDTO) SetMetrics(v string) *BillingDetailRowDTO {
	s.Metrics = &v
	return s
}

func (s *BillingDetailRowDTO) SetModelCode(v string) *BillingDetailRowDTO {
	s.ModelCode = &v
	return s
}

func (s *BillingDetailRowDTO) SetModelId(v int64) *BillingDetailRowDTO {
	s.ModelId = &v
	return s
}

func (s *BillingDetailRowDTO) SetModelName(v string) *BillingDetailRowDTO {
	s.ModelName = &v
	return s
}

func (s *BillingDetailRowDTO) SetModelSymbol(v string) *BillingDetailRowDTO {
	s.ModelSymbol = &v
	return s
}

func (s *BillingDetailRowDTO) SetModelType(v string) *BillingDetailRowDTO {
	s.ModelType = &v
	return s
}

func (s *BillingDetailRowDTO) SetModelVersion(v int32) *BillingDetailRowDTO {
	s.ModelVersion = &v
	return s
}

func (s *BillingDetailRowDTO) SetOutputTokens(v float64) *BillingDetailRowDTO {
	s.OutputTokens = &v
	return s
}

func (s *BillingDetailRowDTO) SetReasoningTokens(v float64) *BillingDetailRowDTO {
	s.ReasoningTokens = &v
	return s
}

func (s *BillingDetailRowDTO) SetRequestId(v string) *BillingDetailRowDTO {
	s.RequestId = &v
	return s
}

func (s *BillingDetailRowDTO) SetRequestTime(v int64) *BillingDetailRowDTO {
	s.RequestTime = &v
	return s
}

func (s *BillingDetailRowDTO) SetTotalTokens(v float64) *BillingDetailRowDTO {
	s.TotalTokens = &v
	return s
}

func (s *BillingDetailRowDTO) SetUsageDetail(v string) *BillingDetailRowDTO {
	s.UsageDetail = &v
	return s
}

func (s *BillingDetailRowDTO) Validate() error {
	return dara.Validate(s)
}
