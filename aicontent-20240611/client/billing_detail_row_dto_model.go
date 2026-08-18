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
	// example:
	//
	// 0.00012800
	Amount *float64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// example:
	//
	// 100
	ApiKeyId *int64 `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	// example:
	//
	// 默认密钥
	ApiKeyName *string `json:"apiKeyName,omitempty" xml:"apiKeyName,omitempty"`
	// example:
	//
	// 0
	CacheCreationTokens *float64 `json:"cacheCreationTokens,omitempty" xml:"cacheCreationTokens,omitempty"`
	// example:
	//
	// 256
	CachedTokens *float64 `json:"cachedTokens,omitempty" xml:"cachedTokens,omitempty"`
	// example:
	//
	// 1
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// example:
	//
	// 研发部
	ClientName *string `json:"clientName,omitempty" xml:"clientName,omitempty"`
	// example:
	//
	// 1.0
	Discount *float64 `json:"discount,omitempty" xml:"discount,omitempty"`
	// example:
	//
	// 1024
	InputTokens *float64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 30001
	MemberUserId *int64 `json:"memberUserId,omitempty" xml:"memberUserId,omitempty"`
	// example:
	//
	// 张三
	MemberUserName *string `json:"memberUserName,omitempty" xml:"memberUserName,omitempty"`
	// example:
	//
	// {}
	Metrics *string `json:"metrics,omitempty" xml:"metrics,omitempty"`
	// example:
	//
	// qwen-plus
	ModelCode *string `json:"modelCode,omitempty" xml:"modelCode,omitempty"`
	// example:
	//
	// 1
	ModelId *int64 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 通义千问-Plus
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// example:
	//
	// qwen
	ModelSymbol *string `json:"modelSymbol,omitempty" xml:"modelSymbol,omitempty"`
	// example:
	//
	// Chat
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// example:
	//
	// 1
	ModelVersion *int32 `json:"modelVersion,omitempty" xml:"modelVersion,omitempty"`
	// example:
	//
	// 512
	OutputTokens *float64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 128
	ReasoningTokens *float64 `json:"reasoningTokens,omitempty" xml:"reasoningTokens,omitempty"`
	// example:
	//
	// chatcmpl-abc123def456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 1700000000
	RequestTime *int64 `json:"requestTime,omitempty" xml:"requestTime,omitempty"`
	// example:
	//
	// 1536
	TotalTokens *float64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
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
