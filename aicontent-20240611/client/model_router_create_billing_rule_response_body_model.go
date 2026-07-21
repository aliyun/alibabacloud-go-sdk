// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateBillingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterCreateBillingRuleResponseBodyData) *ModelRouterCreateBillingRuleResponseBody
	GetData() *ModelRouterCreateBillingRuleResponseBodyData
	SetErrCode(v string) *ModelRouterCreateBillingRuleResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterCreateBillingRuleResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterCreateBillingRuleResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterCreateBillingRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterCreateBillingRuleResponseBody
	GetSuccess() *bool
}

type ModelRouterCreateBillingRuleResponseBody struct {
	// Contains the details of the billing rule.
	Data *ModelRouterCreateBillingRuleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterCreateBillingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateBillingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateBillingRuleResponseBody) GetData() *ModelRouterCreateBillingRuleResponseBodyData {
	return s.Data
}

func (s *ModelRouterCreateBillingRuleResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterCreateBillingRuleResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterCreateBillingRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterCreateBillingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterCreateBillingRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterCreateBillingRuleResponseBody) SetData(v *ModelRouterCreateBillingRuleResponseBodyData) *ModelRouterCreateBillingRuleResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterCreateBillingRuleResponseBody) SetErrCode(v string) *ModelRouterCreateBillingRuleResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterCreateBillingRuleResponseBody) SetErrMessage(v string) *ModelRouterCreateBillingRuleResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterCreateBillingRuleResponseBody) SetHttpStatusCode(v int32) *ModelRouterCreateBillingRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterCreateBillingRuleResponseBody) SetRequestId(v string) *ModelRouterCreateBillingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterCreateBillingRuleResponseBody) SetSuccess(v bool) *ModelRouterCreateBillingRuleResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterCreateBillingRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterCreateBillingRuleResponseBodyData struct {
	// The billing type.
	//
	// example:
	//
	// token_tiered
	BillingType *string `json:"billingType,omitempty" xml:"billingType,omitempty"`
	// The delete marker.
	//
	// example:
	//
	// 0
	DeleteTag *int32 `json:"deleteTag,omitempty" xml:"deleteTag,omitempty"`
	// The effective time of the rule.
	//
	// example:
	//
	// 2024-01-01T00:00:00Z
	EffectiveTime *string `json:"effectiveTime,omitempty" xml:"effectiveTime,omitempty"`
	// The expiration time of the rule.
	//
	// example:
	//
	// 2025-01-01T00:00:00Z
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// The creation time of the rule.
	//
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The modification time of the rule.
	//
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The model code.
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
	// 通义千问Plus
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// The model type.
	//
	// example:
	//
	// chat
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// The pricing configuration.
	//
	// example:
	//
	// {}
	PricingConfig interface{} `json:"pricingConfig,omitempty" xml:"pricingConfig,omitempty"`
	// The provider identifier.
	//
	// example:
	//
	// alibaba
	Symbol *string `json:"symbol,omitempty" xml:"symbol,omitempty"`
	// The version of the rule.
	//
	// example:
	//
	// 1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ModelRouterCreateBillingRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateBillingRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) GetBillingType() *string {
	return s.BillingType
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) GetDeleteTag() *int32 {
	return s.DeleteTag
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) GetModelCode() *string {
	return s.ModelCode
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) GetModelId() *int64 {
	return s.ModelId
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) GetModelName() *string {
	return s.ModelName
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) GetModelType() *string {
	return s.ModelType
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) GetPricingConfig() interface{} {
	return s.PricingConfig
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) GetSymbol() *string {
	return s.Symbol
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) GetVersion() *int32 {
	return s.Version
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) SetBillingType(v string) *ModelRouterCreateBillingRuleResponseBodyData {
	s.BillingType = &v
	return s
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) SetDeleteTag(v int32) *ModelRouterCreateBillingRuleResponseBodyData {
	s.DeleteTag = &v
	return s
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) SetEffectiveTime(v string) *ModelRouterCreateBillingRuleResponseBodyData {
	s.EffectiveTime = &v
	return s
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) SetExpireTime(v string) *ModelRouterCreateBillingRuleResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) SetGmtCreate(v string) *ModelRouterCreateBillingRuleResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) SetGmtModified(v string) *ModelRouterCreateBillingRuleResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) SetId(v int64) *ModelRouterCreateBillingRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) SetModelCode(v string) *ModelRouterCreateBillingRuleResponseBodyData {
	s.ModelCode = &v
	return s
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) SetModelId(v int64) *ModelRouterCreateBillingRuleResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) SetModelName(v string) *ModelRouterCreateBillingRuleResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) SetModelType(v string) *ModelRouterCreateBillingRuleResponseBodyData {
	s.ModelType = &v
	return s
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) SetPricingConfig(v interface{}) *ModelRouterCreateBillingRuleResponseBodyData {
	s.PricingConfig = v
	return s
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) SetSymbol(v string) *ModelRouterCreateBillingRuleResponseBodyData {
	s.Symbol = &v
	return s
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) SetVersion(v int32) *ModelRouterCreateBillingRuleResponseBodyData {
	s.Version = &v
	return s
}

func (s *ModelRouterCreateBillingRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
