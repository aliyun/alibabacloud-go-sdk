// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterUpdateBillingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterUpdateBillingRuleResponseBodyData) *ModelRouterUpdateBillingRuleResponseBody
	GetData() *ModelRouterUpdateBillingRuleResponseBodyData
	SetErrCode(v string) *ModelRouterUpdateBillingRuleResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterUpdateBillingRuleResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterUpdateBillingRuleResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterUpdateBillingRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterUpdateBillingRuleResponseBody
	GetSuccess() *bool
}

type ModelRouterUpdateBillingRuleResponseBody struct {
	Data *ModelRouterUpdateBillingRuleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterUpdateBillingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterUpdateBillingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterUpdateBillingRuleResponseBody) GetData() *ModelRouterUpdateBillingRuleResponseBodyData {
	return s.Data
}

func (s *ModelRouterUpdateBillingRuleResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterUpdateBillingRuleResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterUpdateBillingRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterUpdateBillingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterUpdateBillingRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterUpdateBillingRuleResponseBody) SetData(v *ModelRouterUpdateBillingRuleResponseBodyData) *ModelRouterUpdateBillingRuleResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterUpdateBillingRuleResponseBody) SetErrCode(v string) *ModelRouterUpdateBillingRuleResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterUpdateBillingRuleResponseBody) SetErrMessage(v string) *ModelRouterUpdateBillingRuleResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterUpdateBillingRuleResponseBody) SetHttpStatusCode(v int32) *ModelRouterUpdateBillingRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterUpdateBillingRuleResponseBody) SetRequestId(v string) *ModelRouterUpdateBillingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterUpdateBillingRuleResponseBody) SetSuccess(v bool) *ModelRouterUpdateBillingRuleResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterUpdateBillingRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterUpdateBillingRuleResponseBodyData struct {
	// example:
	//
	// token_tiered
	BillingType *string `json:"billingType,omitempty" xml:"billingType,omitempty"`
	// example:
	//
	// 0
	DeleteTag *int32 `json:"deleteTag,omitempty" xml:"deleteTag,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	EffectiveTime *string `json:"effectiveTime,omitempty" xml:"effectiveTime,omitempty"`
	// example:
	//
	// 2025-01-01T00:00:00Z
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
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
	// 通义千问Plus
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// example:
	//
	// chat
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// example:
	//
	// {}
	PricingConfig interface{} `json:"pricingConfig,omitempty" xml:"pricingConfig,omitempty"`
	// example:
	//
	// alibaba
	Symbol *string `json:"symbol,omitempty" xml:"symbol,omitempty"`
	// example:
	//
	// 1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ModelRouterUpdateBillingRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterUpdateBillingRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) GetBillingType() *string {
	return s.BillingType
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) GetDeleteTag() *int32 {
	return s.DeleteTag
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) GetModelCode() *string {
	return s.ModelCode
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) GetModelId() *int64 {
	return s.ModelId
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) GetModelName() *string {
	return s.ModelName
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) GetModelType() *string {
	return s.ModelType
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) GetPricingConfig() interface{} {
	return s.PricingConfig
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) GetSymbol() *string {
	return s.Symbol
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) GetVersion() *int32 {
	return s.Version
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) SetBillingType(v string) *ModelRouterUpdateBillingRuleResponseBodyData {
	s.BillingType = &v
	return s
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) SetDeleteTag(v int32) *ModelRouterUpdateBillingRuleResponseBodyData {
	s.DeleteTag = &v
	return s
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) SetEffectiveTime(v string) *ModelRouterUpdateBillingRuleResponseBodyData {
	s.EffectiveTime = &v
	return s
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) SetExpireTime(v string) *ModelRouterUpdateBillingRuleResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) SetGmtCreate(v string) *ModelRouterUpdateBillingRuleResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) SetGmtModified(v string) *ModelRouterUpdateBillingRuleResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) SetId(v int64) *ModelRouterUpdateBillingRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) SetModelCode(v string) *ModelRouterUpdateBillingRuleResponseBodyData {
	s.ModelCode = &v
	return s
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) SetModelId(v int64) *ModelRouterUpdateBillingRuleResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) SetModelName(v string) *ModelRouterUpdateBillingRuleResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) SetModelType(v string) *ModelRouterUpdateBillingRuleResponseBodyData {
	s.ModelType = &v
	return s
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) SetPricingConfig(v interface{}) *ModelRouterUpdateBillingRuleResponseBodyData {
	s.PricingConfig = v
	return s
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) SetSymbol(v string) *ModelRouterUpdateBillingRuleResponseBodyData {
	s.Symbol = &v
	return s
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) SetVersion(v int32) *ModelRouterUpdateBillingRuleResponseBodyData {
	s.Version = &v
	return s
}

func (s *ModelRouterUpdateBillingRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
