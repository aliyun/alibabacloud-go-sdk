// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountChangeSecurityEmailRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *EnterpriseAccountChangeSecurityEmailRequest
  GetAppName() *string 
  SetOrientedEcId(v string) *EnterpriseAccountChangeSecurityEmailRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseAccountChangeSecurityEmailRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseAccountChangeSecurityEmailRequest
  GetOrientedNbId() *string 
  SetPk(v string) *EnterpriseAccountChangeSecurityEmailRequest
  GetPk() *string 
  SetRequestId(v string) *EnterpriseAccountChangeSecurityEmailRequest
  GetRequestId() *string 
  SetSecurityEmail(v string) *EnterpriseAccountChangeSecurityEmailRequest
  GetSecurityEmail() *string 
  SetVerifyCode(v string) *EnterpriseAccountChangeSecurityEmailRequest
  GetVerifyCode() *string 
}

type EnterpriseAccountChangeSecurityEmailRequest struct {
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  // This parameter is required.
  Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
  // This parameter is required.
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // This parameter is required.
  SecurityEmail *string `json:"SecurityEmail,omitempty" xml:"SecurityEmail,omitempty"`
  // This parameter is required.
  VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s EnterpriseAccountChangeSecurityEmailRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountChangeSecurityEmailRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountChangeSecurityEmailRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EnterpriseAccountChangeSecurityEmailRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseAccountChangeSecurityEmailRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseAccountChangeSecurityEmailRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseAccountChangeSecurityEmailRequest) GetPk() *string  {
  return s.Pk
}

func (s *EnterpriseAccountChangeSecurityEmailRequest) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseAccountChangeSecurityEmailRequest) GetSecurityEmail() *string  {
  return s.SecurityEmail
}

func (s *EnterpriseAccountChangeSecurityEmailRequest) GetVerifyCode() *string  {
  return s.VerifyCode
}

func (s *EnterpriseAccountChangeSecurityEmailRequest) SetAppName(v string) *EnterpriseAccountChangeSecurityEmailRequest {
  s.AppName = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityEmailRequest) SetOrientedEcId(v string) *EnterpriseAccountChangeSecurityEmailRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityEmailRequest) SetOrientedLeId(v string) *EnterpriseAccountChangeSecurityEmailRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityEmailRequest) SetOrientedNbId(v string) *EnterpriseAccountChangeSecurityEmailRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityEmailRequest) SetPk(v string) *EnterpriseAccountChangeSecurityEmailRequest {
  s.Pk = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityEmailRequest) SetRequestId(v string) *EnterpriseAccountChangeSecurityEmailRequest {
  s.RequestId = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityEmailRequest) SetSecurityEmail(v string) *EnterpriseAccountChangeSecurityEmailRequest {
  s.SecurityEmail = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityEmailRequest) SetVerifyCode(v string) *EnterpriseAccountChangeSecurityEmailRequest {
  s.VerifyCode = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityEmailRequest) Validate() error {
  return dara.Validate(s)
}

