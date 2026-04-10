// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExchangeEntitlementResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExchangeEntitlementResponseBody
  GetCode() *string 
  SetData(v *ExchangeEntitlementResponseBodyData) *ExchangeEntitlementResponseBody
  GetData() *ExchangeEntitlementResponseBodyData 
  SetMessage(v string) *ExchangeEntitlementResponseBody
  GetMessage() *string 
  SetRetryAble(v bool) *ExchangeEntitlementResponseBody
  GetRetryAble() *bool 
  SetSuccess(v bool) *ExchangeEntitlementResponseBody
  GetSuccess() *bool 
}

type ExchangeEntitlementResponseBody struct {
  // example:
  // 
  // 0
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  Data *ExchangeEntitlementResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
  // example:
  // 
  // success
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  // example:
  // 
  // false
  RetryAble *bool `json:"retryAble,omitempty" xml:"retryAble,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExchangeEntitlementResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExchangeEntitlementResponseBody) GoString() string {
  return s.String()
}

func (s *ExchangeEntitlementResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExchangeEntitlementResponseBody) GetData() *ExchangeEntitlementResponseBodyData  {
  return s.Data
}

func (s *ExchangeEntitlementResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExchangeEntitlementResponseBody) GetRetryAble() *bool  {
  return s.RetryAble
}

func (s *ExchangeEntitlementResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExchangeEntitlementResponseBody) SetCode(v string) *ExchangeEntitlementResponseBody {
  s.Code = &v
  return s
}

func (s *ExchangeEntitlementResponseBody) SetData(v *ExchangeEntitlementResponseBodyData) *ExchangeEntitlementResponseBody {
  s.Data = v
  return s
}

func (s *ExchangeEntitlementResponseBody) SetMessage(v string) *ExchangeEntitlementResponseBody {
  s.Message = &v
  return s
}

func (s *ExchangeEntitlementResponseBody) SetRetryAble(v bool) *ExchangeEntitlementResponseBody {
  s.RetryAble = &v
  return s
}

func (s *ExchangeEntitlementResponseBody) SetSuccess(v bool) *ExchangeEntitlementResponseBody {
  s.Success = &v
  return s
}

func (s *ExchangeEntitlementResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExchangeEntitlementResponseBodyData struct {
  // example:
  // 
  // sk_live_abc123xyz789
  ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
  // example:
  // 
  // 2024-01-01T00:00:00Z
  EffectiveAt *string `json:"effectiveAt,omitempty" xml:"effectiveAt,omitempty"`
  // example:
  // 
  // https://llm-gateway.example.com/acme
  Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
  // example:
  // 
  // 2024-01-31T23:59:59Z
  ExpireAt *string `json:"expireAt,omitempty" xml:"expireAt,omitempty"`
  // example:
  // 
  // a1b2c3d4e5f6...
  KeyHash *string `json:"keyHash,omitempty" xml:"keyHash,omitempty"`
  // example:
  // 
  // ORD20240101000001
  RedemptionOrderNo *string `json:"redemptionOrderNo,omitempty" xml:"redemptionOrderNo,omitempty"`
  // example:
  // 
  // false
  Reused *bool `json:"reused,omitempty" xml:"reused,omitempty"`
  Template *ExchangeEntitlementResponseBodyDataTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
}

func (s ExchangeEntitlementResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExchangeEntitlementResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExchangeEntitlementResponseBodyData) GetApiKey() *string  {
  return s.ApiKey
}

func (s *ExchangeEntitlementResponseBodyData) GetEffectiveAt() *string  {
  return s.EffectiveAt
}

func (s *ExchangeEntitlementResponseBodyData) GetEndpoint() *string  {
  return s.Endpoint
}

func (s *ExchangeEntitlementResponseBodyData) GetExpireAt() *string  {
  return s.ExpireAt
}

func (s *ExchangeEntitlementResponseBodyData) GetKeyHash() *string  {
  return s.KeyHash
}

func (s *ExchangeEntitlementResponseBodyData) GetRedemptionOrderNo() *string  {
  return s.RedemptionOrderNo
}

func (s *ExchangeEntitlementResponseBodyData) GetReused() *bool  {
  return s.Reused
}

func (s *ExchangeEntitlementResponseBodyData) GetTemplate() *ExchangeEntitlementResponseBodyDataTemplate  {
  return s.Template
}

func (s *ExchangeEntitlementResponseBodyData) SetApiKey(v string) *ExchangeEntitlementResponseBodyData {
  s.ApiKey = &v
  return s
}

func (s *ExchangeEntitlementResponseBodyData) SetEffectiveAt(v string) *ExchangeEntitlementResponseBodyData {
  s.EffectiveAt = &v
  return s
}

func (s *ExchangeEntitlementResponseBodyData) SetEndpoint(v string) *ExchangeEntitlementResponseBodyData {
  s.Endpoint = &v
  return s
}

func (s *ExchangeEntitlementResponseBodyData) SetExpireAt(v string) *ExchangeEntitlementResponseBodyData {
  s.ExpireAt = &v
  return s
}

func (s *ExchangeEntitlementResponseBodyData) SetKeyHash(v string) *ExchangeEntitlementResponseBodyData {
  s.KeyHash = &v
  return s
}

func (s *ExchangeEntitlementResponseBodyData) SetRedemptionOrderNo(v string) *ExchangeEntitlementResponseBodyData {
  s.RedemptionOrderNo = &v
  return s
}

func (s *ExchangeEntitlementResponseBodyData) SetReused(v bool) *ExchangeEntitlementResponseBodyData {
  s.Reused = &v
  return s
}

func (s *ExchangeEntitlementResponseBodyData) SetTemplate(v *ExchangeEntitlementResponseBodyDataTemplate) *ExchangeEntitlementResponseBodyData {
  s.Template = v
  return s
}

func (s *ExchangeEntitlementResponseBodyData) Validate() error {
  if s.Template != nil {
    if err := s.Template.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExchangeEntitlementResponseBodyDataTemplate struct {
  AllowedModels []*string `json:"allowedModels,omitempty" xml:"allowedModels,omitempty" type:"Repeated"`
  // example:
  // 
  // 1000
  QuotaLimit *int64 `json:"quotaLimit,omitempty" xml:"quotaLimit,omitempty"`
  // example:
  // 
  // 10001
  TemplateId *int64 `json:"templateId,omitempty" xml:"templateId,omitempty"`
  // example:
  // 
  // 基础版 Token 包
  TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
  // example:
  // 
  // TOKEN_PACK
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ExchangeEntitlementResponseBodyDataTemplate) String() string {
  return dara.Prettify(s)
}

func (s ExchangeEntitlementResponseBodyDataTemplate) GoString() string {
  return s.String()
}

func (s *ExchangeEntitlementResponseBodyDataTemplate) GetAllowedModels() []*string  {
  return s.AllowedModels
}

func (s *ExchangeEntitlementResponseBodyDataTemplate) GetQuotaLimit() *int64  {
  return s.QuotaLimit
}

func (s *ExchangeEntitlementResponseBodyDataTemplate) GetTemplateId() *int64  {
  return s.TemplateId
}

func (s *ExchangeEntitlementResponseBodyDataTemplate) GetTemplateName() *string  {
  return s.TemplateName
}

func (s *ExchangeEntitlementResponseBodyDataTemplate) GetType() *string  {
  return s.Type
}

func (s *ExchangeEntitlementResponseBodyDataTemplate) SetAllowedModels(v []*string) *ExchangeEntitlementResponseBodyDataTemplate {
  s.AllowedModels = v
  return s
}

func (s *ExchangeEntitlementResponseBodyDataTemplate) SetQuotaLimit(v int64) *ExchangeEntitlementResponseBodyDataTemplate {
  s.QuotaLimit = &v
  return s
}

func (s *ExchangeEntitlementResponseBodyDataTemplate) SetTemplateId(v int64) *ExchangeEntitlementResponseBodyDataTemplate {
  s.TemplateId = &v
  return s
}

func (s *ExchangeEntitlementResponseBodyDataTemplate) SetTemplateName(v string) *ExchangeEntitlementResponseBodyDataTemplate {
  s.TemplateName = &v
  return s
}

func (s *ExchangeEntitlementResponseBodyDataTemplate) SetType(v string) *ExchangeEntitlementResponseBodyDataTemplate {
  s.Type = &v
  return s
}

func (s *ExchangeEntitlementResponseBodyDataTemplate) Validate() error {
  return dara.Validate(s)
}

