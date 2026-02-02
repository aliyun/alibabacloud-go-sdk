// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntElementVerifyV2Request interface {
  dara.Model
  String() string
  GoString() string
  SetEntName(v string) *EntElementVerifyV2Request
  GetEntName() *string 
  SetInfoVerifyType(v string) *EntElementVerifyV2Request
  GetInfoVerifyType() *string 
  SetLegalPersonCertNo(v string) *EntElementVerifyV2Request
  GetLegalPersonCertNo() *string 
  SetLegalPersonName(v string) *EntElementVerifyV2Request
  GetLegalPersonName() *string 
  SetLicenseNo(v string) *EntElementVerifyV2Request
  GetLicenseNo() *string 
  SetMerchantBizId(v string) *EntElementVerifyV2Request
  GetMerchantBizId() *string 
  SetMerchantUserId(v string) *EntElementVerifyV2Request
  GetMerchantUserId() *string 
  SetSceneCode(v string) *EntElementVerifyV2Request
  GetSceneCode() *string 
  SetUserAuthorization(v string) *EntElementVerifyV2Request
  GetUserAuthorization() *string 
}

type EntElementVerifyV2Request struct {
  // Enterprise Name.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ****有限公司
  EntName *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
  // Type of enterprise element verification.
  // 
  // - ENT_2META: Two-element verification (enterprise name + unified social credit code)
  // 
  // - ENT_3META: Three-element verification (enterprise name + unified social credit code + legal person\\"s name or organization head)
  // 
  // - ENT_4META: Four-element verification (enterprise name + unified social credit code + legal person\\"s name or organization head + legal person\\"s ID number)
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ENT_2META
  InfoVerifyType *string `json:"InfoVerifyType,omitempty" xml:"InfoVerifyType,omitempty"`
  // Legal Person\\"s ID Number. Required for four-element scenarios.
  // 
  // example:
  // 
  // 1******************9
  LegalPersonCertNo *string `json:"LegalPersonCertNo,omitempty" xml:"LegalPersonCertNo,omitempty"`
  // Legal Person\\"s Name. Required for three-element and four-element scenarios.
  // 
  // example:
  // 
  // 张**
  LegalPersonName *string `json:"LegalPersonName,omitempty" xml:"LegalPersonName,omitempty"`
  // Unified Social Credit Code.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 91330106673959****
  LicenseNo *string `json:"LicenseNo,omitempty" xml:"LicenseNo,omitempty"`
  // A unique business identifier defined by the merchant side, used for subsequent problem localization and troubleshooting. Supports a combination of letters and numbers, with a maximum length of 32 characters. Please ensure uniqueness.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // e0c34a77f5ac40a5aa5e6ed20c35****
  MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
  // Merchant-side user ID. Supports a combination of letters and numbers, with a maximum length of 32 characters.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // mch7x9a2b4c8d3e5f6g1h2i3j4k5****
  MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
  // Custom scene code, defined by the user to distinguish between different business scenarios. Supports a combination of letters, numbers, and underscores, with a maximum length of 32 characters.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1000000006
  SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
  // Whether user authorization has been obtained.
  // 
  // - 1: Authorization obtained
  // 
  // - 0: No authorization obtained
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1
  UserAuthorization *string `json:"UserAuthorization,omitempty" xml:"UserAuthorization,omitempty"`
}

func (s EntElementVerifyV2Request) String() string {
  return dara.Prettify(s)
}

func (s EntElementVerifyV2Request) GoString() string {
  return s.String()
}

func (s *EntElementVerifyV2Request) GetEntName() *string  {
  return s.EntName
}

func (s *EntElementVerifyV2Request) GetInfoVerifyType() *string  {
  return s.InfoVerifyType
}

func (s *EntElementVerifyV2Request) GetLegalPersonCertNo() *string  {
  return s.LegalPersonCertNo
}

func (s *EntElementVerifyV2Request) GetLegalPersonName() *string  {
  return s.LegalPersonName
}

func (s *EntElementVerifyV2Request) GetLicenseNo() *string  {
  return s.LicenseNo
}

func (s *EntElementVerifyV2Request) GetMerchantBizId() *string  {
  return s.MerchantBizId
}

func (s *EntElementVerifyV2Request) GetMerchantUserId() *string  {
  return s.MerchantUserId
}

func (s *EntElementVerifyV2Request) GetSceneCode() *string  {
  return s.SceneCode
}

func (s *EntElementVerifyV2Request) GetUserAuthorization() *string  {
  return s.UserAuthorization
}

func (s *EntElementVerifyV2Request) SetEntName(v string) *EntElementVerifyV2Request {
  s.EntName = &v
  return s
}

func (s *EntElementVerifyV2Request) SetInfoVerifyType(v string) *EntElementVerifyV2Request {
  s.InfoVerifyType = &v
  return s
}

func (s *EntElementVerifyV2Request) SetLegalPersonCertNo(v string) *EntElementVerifyV2Request {
  s.LegalPersonCertNo = &v
  return s
}

func (s *EntElementVerifyV2Request) SetLegalPersonName(v string) *EntElementVerifyV2Request {
  s.LegalPersonName = &v
  return s
}

func (s *EntElementVerifyV2Request) SetLicenseNo(v string) *EntElementVerifyV2Request {
  s.LicenseNo = &v
  return s
}

func (s *EntElementVerifyV2Request) SetMerchantBizId(v string) *EntElementVerifyV2Request {
  s.MerchantBizId = &v
  return s
}

func (s *EntElementVerifyV2Request) SetMerchantUserId(v string) *EntElementVerifyV2Request {
  s.MerchantUserId = &v
  return s
}

func (s *EntElementVerifyV2Request) SetSceneCode(v string) *EntElementVerifyV2Request {
  s.SceneCode = &v
  return s
}

func (s *EntElementVerifyV2Request) SetUserAuthorization(v string) *EntElementVerifyV2Request {
  s.UserAuthorization = &v
  return s
}

func (s *EntElementVerifyV2Request) Validate() error {
  return dara.Validate(s)
}

