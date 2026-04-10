// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExchangeEntitlementRequest interface {
  dara.Model
  String() string
  GoString() string
  SetExternalUserId(v string) *ExchangeEntitlementRequest
  GetExternalUserId() *string 
  SetKeyHash(v string) *ExchangeEntitlementRequest
  GetKeyHash() *string 
  SetRequestId(v string) *ExchangeEntitlementRequest
  GetRequestId() *string 
  SetTemplateId(v int64) *ExchangeEntitlementRequest
  GetTemplateId() *int64 
  SetUserName(v string) *ExchangeEntitlementRequest
  GetUserName() *string 
}

type ExchangeEntitlementRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // your_user_id_001
  ExternalUserId *string `json:"externalUserId,omitempty" xml:"externalUserId,omitempty"`
  // example:
  // 
  // a1b2c3d4e5f6...
  KeyHash *string `json:"keyHash,omitempty" xml:"keyHash,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // req_20240101_001
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 10001
  TemplateId *int64 `json:"templateId,omitempty" xml:"templateId,omitempty"`
  // example:
  // 
  // zhangsan
  UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s ExchangeEntitlementRequest) String() string {
  return dara.Prettify(s)
}

func (s ExchangeEntitlementRequest) GoString() string {
  return s.String()
}

func (s *ExchangeEntitlementRequest) GetExternalUserId() *string  {
  return s.ExternalUserId
}

func (s *ExchangeEntitlementRequest) GetKeyHash() *string  {
  return s.KeyHash
}

func (s *ExchangeEntitlementRequest) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExchangeEntitlementRequest) GetTemplateId() *int64  {
  return s.TemplateId
}

func (s *ExchangeEntitlementRequest) GetUserName() *string  {
  return s.UserName
}

func (s *ExchangeEntitlementRequest) SetExternalUserId(v string) *ExchangeEntitlementRequest {
  s.ExternalUserId = &v
  return s
}

func (s *ExchangeEntitlementRequest) SetKeyHash(v string) *ExchangeEntitlementRequest {
  s.KeyHash = &v
  return s
}

func (s *ExchangeEntitlementRequest) SetRequestId(v string) *ExchangeEntitlementRequest {
  s.RequestId = &v
  return s
}

func (s *ExchangeEntitlementRequest) SetTemplateId(v int64) *ExchangeEntitlementRequest {
  s.TemplateId = &v
  return s
}

func (s *ExchangeEntitlementRequest) SetUserName(v string) *ExchangeEntitlementRequest {
  s.UserName = &v
  return s
}

func (s *ExchangeEntitlementRequest) Validate() error {
  return dara.Validate(s)
}

