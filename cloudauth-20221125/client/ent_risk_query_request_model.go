// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntRiskQueryRequest interface {
  dara.Model
  String() string
  GoString() string
  SetMerchantBizId(v string) *EntRiskQueryRequest
  GetMerchantBizId() *string 
  SetMerchantUserId(v string) *EntRiskQueryRequest
  GetMerchantUserId() *string 
  SetParamType(v string) *EntRiskQueryRequest
  GetParamType() *string 
  SetParamValue(v string) *EntRiskQueryRequest
  GetParamValue() *string 
  SetSceneCode(v string) *EntRiskQueryRequest
  GetSceneCode() *string 
  SetUserAuthorization(v string) *EntRiskQueryRequest
  GetUserAuthorization() *string 
}

type EntRiskQueryRequest struct {
  // A unique business identifier defined by the merchant side, used for subsequent problem localization and troubleshooting. Supports a combination of 32 alphanumeric characters, please ensure uniqueness.
  // 
  // example:
  // 
  // 32198****193000
  MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
  // A custom user ID in your business, used for subsequent problem localization and troubleshooting.
  // 
  // example:
  // 
  // 无
  MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
  // Parameter type.
  // 
  // 00: Company name;
  // 
  // 01: Business registration number;
  // 
  // 02: Unified Social Credit Code;
  // 
  // 03: Organization code;
  // 
  // example:
  // 
  // 00
  ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
  // Enter one of the following based on the ParamType: company name, business registration number, unified social credit code, or organization code.
  // 
  // example:
  // 
  // 91330106673959****
  ParamValue *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
  // Custom scene code, used to distinguish business scenarios, a 10-digit number.
  // 
  // example:
  // 
  // 1000000086
  SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
  // Whether user authorization is obtained.
  // 
  // 1: Authorized
  // 
  // 0: Not authorized
  // 
  // example:
  // 
  // 1
  UserAuthorization *string `json:"UserAuthorization,omitempty" xml:"UserAuthorization,omitempty"`
}

func (s EntRiskQueryRequest) String() string {
  return dara.Prettify(s)
}

func (s EntRiskQueryRequest) GoString() string {
  return s.String()
}

func (s *EntRiskQueryRequest) GetMerchantBizId() *string  {
  return s.MerchantBizId
}

func (s *EntRiskQueryRequest) GetMerchantUserId() *string  {
  return s.MerchantUserId
}

func (s *EntRiskQueryRequest) GetParamType() *string  {
  return s.ParamType
}

func (s *EntRiskQueryRequest) GetParamValue() *string  {
  return s.ParamValue
}

func (s *EntRiskQueryRequest) GetSceneCode() *string  {
  return s.SceneCode
}

func (s *EntRiskQueryRequest) GetUserAuthorization() *string  {
  return s.UserAuthorization
}

func (s *EntRiskQueryRequest) SetMerchantBizId(v string) *EntRiskQueryRequest {
  s.MerchantBizId = &v
  return s
}

func (s *EntRiskQueryRequest) SetMerchantUserId(v string) *EntRiskQueryRequest {
  s.MerchantUserId = &v
  return s
}

func (s *EntRiskQueryRequest) SetParamType(v string) *EntRiskQueryRequest {
  s.ParamType = &v
  return s
}

func (s *EntRiskQueryRequest) SetParamValue(v string) *EntRiskQueryRequest {
  s.ParamValue = &v
  return s
}

func (s *EntRiskQueryRequest) SetSceneCode(v string) *EntRiskQueryRequest {
  s.SceneCode = &v
  return s
}

func (s *EntRiskQueryRequest) SetUserAuthorization(v string) *EntRiskQueryRequest {
  s.UserAuthorization = &v
  return s
}

func (s *EntRiskQueryRequest) Validate() error {
  return dara.Validate(s)
}

