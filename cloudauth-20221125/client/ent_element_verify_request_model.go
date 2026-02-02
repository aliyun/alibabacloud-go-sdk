// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntElementVerifyRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEntName(v string) *EntElementVerifyRequest
  GetEntName() *string 
  SetInfoVerifyType(v string) *EntElementVerifyRequest
  GetInfoVerifyType() *string 
  SetLegalPersonCertNo(v string) *EntElementVerifyRequest
  GetLegalPersonCertNo() *string 
  SetLegalPersonName(v string) *EntElementVerifyRequest
  GetLegalPersonName() *string 
  SetLicenseNo(v string) *EntElementVerifyRequest
  GetLicenseNo() *string 
  SetMerchantBizId(v string) *EntElementVerifyRequest
  GetMerchantBizId() *string 
  SetMerchantUserId(v string) *EntElementVerifyRequest
  GetMerchantUserId() *string 
  SetSceneCode(v string) *EntElementVerifyRequest
  GetSceneCode() *string 
  SetUserAuthorization(v string) *EntElementVerifyRequest
  GetUserAuthorization() *string 
}

type EntElementVerifyRequest struct {
  // Enterprise name.
  // 
  // example:
  // 
  // ***有限公司
  EntName *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
  // Type of enterprise element verification.
  // 
  // - ENT_2META: Two elements
  // 
  // - ENT_3META: Three elements
  // 
  // - ENT_4META: Four elements
  // 
  // example:
  // 
  // ENT_2META
  InfoVerifyType *string `json:"InfoVerifyType,omitempty" xml:"InfoVerifyType,omitempty"`
  // Legal representative\\"s ID number. Required for the four-element scenario.
  // 
  // example:
  // 
  // 370105*****3892
  LegalPersonCertNo *string `json:"LegalPersonCertNo,omitempty" xml:"LegalPersonCertNo,omitempty"`
  // Legal representative\\"s name. Required for three-element and four-element scenarios.
  // 
  // example:
  // 
  // 张**
  LegalPersonName *string `json:"LegalPersonName,omitempty" xml:"LegalPersonName,omitempty"`
  // Business license number.
  // 
  // example:
  // 
  // 32132***328932
  LicenseNo *string `json:"LicenseNo,omitempty" xml:"LicenseNo,omitempty"`
  // A unique business identifier defined by the merchant for subsequent troubleshooting. It supports a combination of letters and numbers, with a maximum length of 32 characters. Please ensure uniqueness.
  // 
  // example:
  // 
  // 32198****193000
  MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
  // Merchant-side user ID. Supports a combination of letters and numbers, with a maximum length of 32 characters.
  // 
  // example:
  // 
  // 432***421
  MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
  // Scene code. Supports a combination of letters, numbers, and underscores, with a maximum length of 32 characters.
  // 
  // example:
  // 
  // withdraw
  SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
  // Whether user authorization is obtained.
  // 
  // - 1: Authorized
  // 
  // - 0: Not authorized
  // 
  // example:
  // 
  // 1
  UserAuthorization *string `json:"UserAuthorization,omitempty" xml:"UserAuthorization,omitempty"`
}

func (s EntElementVerifyRequest) String() string {
  return dara.Prettify(s)
}

func (s EntElementVerifyRequest) GoString() string {
  return s.String()
}

func (s *EntElementVerifyRequest) GetEntName() *string  {
  return s.EntName
}

func (s *EntElementVerifyRequest) GetInfoVerifyType() *string  {
  return s.InfoVerifyType
}

func (s *EntElementVerifyRequest) GetLegalPersonCertNo() *string  {
  return s.LegalPersonCertNo
}

func (s *EntElementVerifyRequest) GetLegalPersonName() *string  {
  return s.LegalPersonName
}

func (s *EntElementVerifyRequest) GetLicenseNo() *string  {
  return s.LicenseNo
}

func (s *EntElementVerifyRequest) GetMerchantBizId() *string  {
  return s.MerchantBizId
}

func (s *EntElementVerifyRequest) GetMerchantUserId() *string  {
  return s.MerchantUserId
}

func (s *EntElementVerifyRequest) GetSceneCode() *string  {
  return s.SceneCode
}

func (s *EntElementVerifyRequest) GetUserAuthorization() *string  {
  return s.UserAuthorization
}

func (s *EntElementVerifyRequest) SetEntName(v string) *EntElementVerifyRequest {
  s.EntName = &v
  return s
}

func (s *EntElementVerifyRequest) SetInfoVerifyType(v string) *EntElementVerifyRequest {
  s.InfoVerifyType = &v
  return s
}

func (s *EntElementVerifyRequest) SetLegalPersonCertNo(v string) *EntElementVerifyRequest {
  s.LegalPersonCertNo = &v
  return s
}

func (s *EntElementVerifyRequest) SetLegalPersonName(v string) *EntElementVerifyRequest {
  s.LegalPersonName = &v
  return s
}

func (s *EntElementVerifyRequest) SetLicenseNo(v string) *EntElementVerifyRequest {
  s.LicenseNo = &v
  return s
}

func (s *EntElementVerifyRequest) SetMerchantBizId(v string) *EntElementVerifyRequest {
  s.MerchantBizId = &v
  return s
}

func (s *EntElementVerifyRequest) SetMerchantUserId(v string) *EntElementVerifyRequest {
  s.MerchantUserId = &v
  return s
}

func (s *EntElementVerifyRequest) SetSceneCode(v string) *EntElementVerifyRequest {
  s.SceneCode = &v
  return s
}

func (s *EntElementVerifyRequest) SetUserAuthorization(v string) *EntElementVerifyRequest {
  s.UserAuthorization = &v
  return s
}

func (s *EntElementVerifyRequest) Validate() error {
  return dara.Validate(s)
}

