// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntElementVerifyPRORequest interface {
  dara.Model
  String() string
  GoString() string
  SetEntName(v string) *EntElementVerifyPRORequest
  GetEntName() *string 
  SetInfoVerifyType(v string) *EntElementVerifyPRORequest
  GetInfoVerifyType() *string 
  SetLegalPersonCertNo(v string) *EntElementVerifyPRORequest
  GetLegalPersonCertNo() *string 
  SetLegalPersonName(v string) *EntElementVerifyPRORequest
  GetLegalPersonName() *string 
  SetLicenseNo(v string) *EntElementVerifyPRORequest
  GetLicenseNo() *string 
  SetMerchantBizId(v string) *EntElementVerifyPRORequest
  GetMerchantBizId() *string 
  SetMerchantUserId(v string) *EntElementVerifyPRORequest
  GetMerchantUserId() *string 
  SetSceneCode(v string) *EntElementVerifyPRORequest
  GetSceneCode() *string 
  SetUserAuthorization(v string) *EntElementVerifyPRORequest
  GetUserAuthorization() *string 
}

type EntElementVerifyPRORequest struct {
  EntName *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
  // example:
  // 
  // ENT_4META
  InfoVerifyType *string `json:"InfoVerifyType,omitempty" xml:"InfoVerifyType,omitempty"`
  // example:
  // 
  // 370105*****3892
  LegalPersonCertNo *string `json:"LegalPersonCertNo,omitempty" xml:"LegalPersonCertNo,omitempty"`
  LegalPersonName *string `json:"LegalPersonName,omitempty" xml:"LegalPersonName,omitempty"`
  // example:
  // 
  // 91330106673959****
  LicenseNo *string `json:"LicenseNo,omitempty" xml:"LicenseNo,omitempty"`
  // example:
  // 
  // e0c34a77f5ac40a5aa5e6ed20c35****
  MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
  // example:
  // 
  // mch7x9a2b4c8d3e5f6g1h2i3j4k5****
  MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
  // example:
  // 
  // 1000000006
  SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
  // example:
  // 
  // 1
  UserAuthorization *string `json:"UserAuthorization,omitempty" xml:"UserAuthorization,omitempty"`
}

func (s EntElementVerifyPRORequest) String() string {
  return dara.Prettify(s)
}

func (s EntElementVerifyPRORequest) GoString() string {
  return s.String()
}

func (s *EntElementVerifyPRORequest) GetEntName() *string  {
  return s.EntName
}

func (s *EntElementVerifyPRORequest) GetInfoVerifyType() *string  {
  return s.InfoVerifyType
}

func (s *EntElementVerifyPRORequest) GetLegalPersonCertNo() *string  {
  return s.LegalPersonCertNo
}

func (s *EntElementVerifyPRORequest) GetLegalPersonName() *string  {
  return s.LegalPersonName
}

func (s *EntElementVerifyPRORequest) GetLicenseNo() *string  {
  return s.LicenseNo
}

func (s *EntElementVerifyPRORequest) GetMerchantBizId() *string  {
  return s.MerchantBizId
}

func (s *EntElementVerifyPRORequest) GetMerchantUserId() *string  {
  return s.MerchantUserId
}

func (s *EntElementVerifyPRORequest) GetSceneCode() *string  {
  return s.SceneCode
}

func (s *EntElementVerifyPRORequest) GetUserAuthorization() *string  {
  return s.UserAuthorization
}

func (s *EntElementVerifyPRORequest) SetEntName(v string) *EntElementVerifyPRORequest {
  s.EntName = &v
  return s
}

func (s *EntElementVerifyPRORequest) SetInfoVerifyType(v string) *EntElementVerifyPRORequest {
  s.InfoVerifyType = &v
  return s
}

func (s *EntElementVerifyPRORequest) SetLegalPersonCertNo(v string) *EntElementVerifyPRORequest {
  s.LegalPersonCertNo = &v
  return s
}

func (s *EntElementVerifyPRORequest) SetLegalPersonName(v string) *EntElementVerifyPRORequest {
  s.LegalPersonName = &v
  return s
}

func (s *EntElementVerifyPRORequest) SetLicenseNo(v string) *EntElementVerifyPRORequest {
  s.LicenseNo = &v
  return s
}

func (s *EntElementVerifyPRORequest) SetMerchantBizId(v string) *EntElementVerifyPRORequest {
  s.MerchantBizId = &v
  return s
}

func (s *EntElementVerifyPRORequest) SetMerchantUserId(v string) *EntElementVerifyPRORequest {
  s.MerchantUserId = &v
  return s
}

func (s *EntElementVerifyPRORequest) SetSceneCode(v string) *EntElementVerifyPRORequest {
  s.SceneCode = &v
  return s
}

func (s *EntElementVerifyPRORequest) SetUserAuthorization(v string) *EntElementVerifyPRORequest {
  s.UserAuthorization = &v
  return s
}

func (s *EntElementVerifyPRORequest) Validate() error {
  return dara.Validate(s)
}

