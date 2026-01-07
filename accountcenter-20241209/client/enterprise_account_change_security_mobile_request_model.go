// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountChangeSecurityMobileRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *EnterpriseAccountChangeSecurityMobileRequest
  GetAppName() *string 
  SetEncryptTicket(v string) *EnterpriseAccountChangeSecurityMobileRequest
  GetEncryptTicket() *string 
  SetNewMobile(v string) *EnterpriseAccountChangeSecurityMobileRequest
  GetNewMobile() *string 
  SetOrientedEcId(v string) *EnterpriseAccountChangeSecurityMobileRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseAccountChangeSecurityMobileRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseAccountChangeSecurityMobileRequest
  GetOrientedNbId() *string 
  SetPk(v string) *EnterpriseAccountChangeSecurityMobileRequest
  GetPk() *string 
  SetRequestId(v string) *EnterpriseAccountChangeSecurityMobileRequest
  GetRequestId() *string 
  SetVerificationCode(v string) *EnterpriseAccountChangeSecurityMobileRequest
  GetVerificationCode() *string 
}

type EnterpriseAccountChangeSecurityMobileRequest struct {
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  EncryptTicket *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
  // This parameter is required.
  NewMobile *string `json:"NewMobile,omitempty" xml:"NewMobile,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  // This parameter is required.
  Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
  // This parameter is required.
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // This parameter is required.
  VerificationCode *string `json:"VerificationCode,omitempty" xml:"VerificationCode,omitempty"`
}

func (s EnterpriseAccountChangeSecurityMobileRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountChangeSecurityMobileRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) GetEncryptTicket() *string  {
  return s.EncryptTicket
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) GetNewMobile() *string  {
  return s.NewMobile
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) GetPk() *string  {
  return s.Pk
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) GetVerificationCode() *string  {
  return s.VerificationCode
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) SetAppName(v string) *EnterpriseAccountChangeSecurityMobileRequest {
  s.AppName = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) SetEncryptTicket(v string) *EnterpriseAccountChangeSecurityMobileRequest {
  s.EncryptTicket = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) SetNewMobile(v string) *EnterpriseAccountChangeSecurityMobileRequest {
  s.NewMobile = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) SetOrientedEcId(v string) *EnterpriseAccountChangeSecurityMobileRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) SetOrientedLeId(v string) *EnterpriseAccountChangeSecurityMobileRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) SetOrientedNbId(v string) *EnterpriseAccountChangeSecurityMobileRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) SetPk(v string) *EnterpriseAccountChangeSecurityMobileRequest {
  s.Pk = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) SetRequestId(v string) *EnterpriseAccountChangeSecurityMobileRequest {
  s.RequestId = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) SetVerificationCode(v string) *EnterpriseAccountChangeSecurityMobileRequest {
  s.VerificationCode = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityMobileRequest) Validate() error {
  return dara.Validate(s)
}

