// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntVerifyRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAccountNo(v string) *EntVerifyRequest
  GetAccountNo() *string 
  SetEntName(v string) *EntVerifyRequest
  GetEntName() *string 
  SetInfoVerifyType(v string) *EntVerifyRequest
  GetInfoVerifyType() *string 
  SetLegalPersonCertNo(v string) *EntVerifyRequest
  GetLegalPersonCertNo() *string 
  SetLegalPersonMobile(v string) *EntVerifyRequest
  GetLegalPersonMobile() *string 
  SetLegalPersonName(v string) *EntVerifyRequest
  GetLegalPersonName() *string 
  SetLicenseNo(v string) *EntVerifyRequest
  GetLicenseNo() *string 
  SetMerchantBizId(v string) *EntVerifyRequest
  GetMerchantBizId() *string 
  SetMerchantUserId(v string) *EntVerifyRequest
  GetMerchantUserId() *string 
  SetRiskModelVersion(v string) *EntVerifyRequest
  GetRiskModelVersion() *string 
  SetRiskVerifyType(v string) *EntVerifyRequest
  GetRiskVerifyType() *string 
  SetSceneCode(v string) *EntVerifyRequest
  GetSceneCode() *string 
  SetUserAuthorization(v string) *EntVerifyRequest
  GetUserAuthorization() *string 
}

type EntVerifyRequest struct {
  // Receiving account, to assist in improving the risk assessment effect.
  // 
  // example:
  // 
  // 321324***38293
  AccountNo *string `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
  // Enterprise name.
  // 
  // example:
  // 
  // ***有限公司
  EntName *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
  // Enterprise element verification type, currently not supported.
  // 
  // example:
  // 
  // 无
  InfoVerifyType *string `json:"InfoVerifyType,omitempty" xml:"InfoVerifyType,omitempty"`
  // Legal person\\"s ID number.
  // 
  // example:
  // 
  // 370105*****3892
  LegalPersonCertNo *string `json:"LegalPersonCertNo,omitempty" xml:"LegalPersonCertNo,omitempty"`
  // Legal person\\"s mobile phone number.
  // 
  // example:
  // 
  // 1300***53
  LegalPersonMobile *string `json:"LegalPersonMobile,omitempty" xml:"LegalPersonMobile,omitempty"`
  // Legal person\\"s name.
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
  // A unique business identifier defined by the merchant, used for subsequent problem localization and troubleshooting. It supports a combination of letters and numbers, with a maximum length of 32 characters. Please ensure its uniqueness.
  // 
  // example:
  // 
  // 32198****193000
  MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
  // Merchant-side user ID. It supports a combination of letters and numbers, with a maximum length of 32 characters.
  // 
  // example:
  // 
  // 432***421
  MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
  // Enterprise risk verification model version, required when RiskVerifyType is not empty. Currently supported:
  // 
  // - BASIC: Basic version
  // 
  // example:
  // 
  // BASIC
  RiskModelVersion *string `json:"RiskModelVersion,omitempty" xml:"RiskModelVersion,omitempty"`
  // Enterprise risk verification type, at least one of InfoVerifyType or this must be selected. Currently supported:
  // 
  // - BUSINESS_ANTIFRAUD
  // 
  // example:
  // 
  // BUSINESS_ANTIFRAUD
  RiskVerifyType *string `json:"RiskVerifyType,omitempty" xml:"RiskVerifyType,omitempty"`
  // Scene code.
  // 
  // example:
  // 
  // withdraw
  SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
  // Whether the user authorization is obtained.
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

func (s EntVerifyRequest) String() string {
  return dara.Prettify(s)
}

func (s EntVerifyRequest) GoString() string {
  return s.String()
}

func (s *EntVerifyRequest) GetAccountNo() *string  {
  return s.AccountNo
}

func (s *EntVerifyRequest) GetEntName() *string  {
  return s.EntName
}

func (s *EntVerifyRequest) GetInfoVerifyType() *string  {
  return s.InfoVerifyType
}

func (s *EntVerifyRequest) GetLegalPersonCertNo() *string  {
  return s.LegalPersonCertNo
}

func (s *EntVerifyRequest) GetLegalPersonMobile() *string  {
  return s.LegalPersonMobile
}

func (s *EntVerifyRequest) GetLegalPersonName() *string  {
  return s.LegalPersonName
}

func (s *EntVerifyRequest) GetLicenseNo() *string  {
  return s.LicenseNo
}

func (s *EntVerifyRequest) GetMerchantBizId() *string  {
  return s.MerchantBizId
}

func (s *EntVerifyRequest) GetMerchantUserId() *string  {
  return s.MerchantUserId
}

func (s *EntVerifyRequest) GetRiskModelVersion() *string  {
  return s.RiskModelVersion
}

func (s *EntVerifyRequest) GetRiskVerifyType() *string  {
  return s.RiskVerifyType
}

func (s *EntVerifyRequest) GetSceneCode() *string  {
  return s.SceneCode
}

func (s *EntVerifyRequest) GetUserAuthorization() *string  {
  return s.UserAuthorization
}

func (s *EntVerifyRequest) SetAccountNo(v string) *EntVerifyRequest {
  s.AccountNo = &v
  return s
}

func (s *EntVerifyRequest) SetEntName(v string) *EntVerifyRequest {
  s.EntName = &v
  return s
}

func (s *EntVerifyRequest) SetInfoVerifyType(v string) *EntVerifyRequest {
  s.InfoVerifyType = &v
  return s
}

func (s *EntVerifyRequest) SetLegalPersonCertNo(v string) *EntVerifyRequest {
  s.LegalPersonCertNo = &v
  return s
}

func (s *EntVerifyRequest) SetLegalPersonMobile(v string) *EntVerifyRequest {
  s.LegalPersonMobile = &v
  return s
}

func (s *EntVerifyRequest) SetLegalPersonName(v string) *EntVerifyRequest {
  s.LegalPersonName = &v
  return s
}

func (s *EntVerifyRequest) SetLicenseNo(v string) *EntVerifyRequest {
  s.LicenseNo = &v
  return s
}

func (s *EntVerifyRequest) SetMerchantBizId(v string) *EntVerifyRequest {
  s.MerchantBizId = &v
  return s
}

func (s *EntVerifyRequest) SetMerchantUserId(v string) *EntVerifyRequest {
  s.MerchantUserId = &v
  return s
}

func (s *EntVerifyRequest) SetRiskModelVersion(v string) *EntVerifyRequest {
  s.RiskModelVersion = &v
  return s
}

func (s *EntVerifyRequest) SetRiskVerifyType(v string) *EntVerifyRequest {
  s.RiskVerifyType = &v
  return s
}

func (s *EntVerifyRequest) SetSceneCode(v string) *EntVerifyRequest {
  s.SceneCode = &v
  return s
}

func (s *EntVerifyRequest) SetUserAuthorization(v string) *EntVerifyRequest {
  s.UserAuthorization = &v
  return s
}

func (s *EntVerifyRequest) Validate() error {
  return dara.Validate(s)
}

