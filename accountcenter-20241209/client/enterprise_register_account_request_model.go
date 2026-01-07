// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseRegisterAccountRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAlias(v string) *EnterpriseRegisterAccountRequest
  GetAlias() *string 
  SetEncryptPassword(v string) *EnterpriseRegisterAccountRequest
  GetEncryptPassword() *string 
  SetEncryptTicket(v string) *EnterpriseRegisterAccountRequest
  GetEncryptTicket() *string 
  SetLoginEmail(v string) *EnterpriseRegisterAccountRequest
  GetLoginEmail() *string 
  SetOrganizationId(v string) *EnterpriseRegisterAccountRequest
  GetOrganizationId() *string 
  SetOrientedEcId(v string) *EnterpriseRegisterAccountRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseRegisterAccountRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseRegisterAccountRequest
  GetOrientedNbId() *string 
  SetRequestId(v string) *EnterpriseRegisterAccountRequest
  GetRequestId() *string 
  SetShowCompleteInfo(v bool) *EnterpriseRegisterAccountRequest
  GetShowCompleteInfo() *bool 
  SetSiteNick(v string) *EnterpriseRegisterAccountRequest
  GetSiteNick() *string 
}

type EnterpriseRegisterAccountRequest struct {
  // example:
  // 
  // 资方支付平台
  Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
  EncryptPassword *string `json:"EncryptPassword,omitempty" xml:"EncryptPassword,omitempty"`
  EncryptTicket *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
  LoginEmail *string `json:"LoginEmail,omitempty" xml:"LoginEmail,omitempty"`
  // example:
  // 
  // 668514d8083f058f78f7199a
  OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // A93073FC-1E86-58BA-AB83-54DA6BDB4F03
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // false
  ShowCompleteInfo *bool `json:"ShowCompleteInfo,omitempty" xml:"ShowCompleteInfo,omitempty"`
  SiteNick *string `json:"SiteNick,omitempty" xml:"SiteNick,omitempty"`
}

func (s EnterpriseRegisterAccountRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRegisterAccountRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseRegisterAccountRequest) GetAlias() *string  {
  return s.Alias
}

func (s *EnterpriseRegisterAccountRequest) GetEncryptPassword() *string  {
  return s.EncryptPassword
}

func (s *EnterpriseRegisterAccountRequest) GetEncryptTicket() *string  {
  return s.EncryptTicket
}

func (s *EnterpriseRegisterAccountRequest) GetLoginEmail() *string  {
  return s.LoginEmail
}

func (s *EnterpriseRegisterAccountRequest) GetOrganizationId() *string  {
  return s.OrganizationId
}

func (s *EnterpriseRegisterAccountRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseRegisterAccountRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseRegisterAccountRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseRegisterAccountRequest) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseRegisterAccountRequest) GetShowCompleteInfo() *bool  {
  return s.ShowCompleteInfo
}

func (s *EnterpriseRegisterAccountRequest) GetSiteNick() *string  {
  return s.SiteNick
}

func (s *EnterpriseRegisterAccountRequest) SetAlias(v string) *EnterpriseRegisterAccountRequest {
  s.Alias = &v
  return s
}

func (s *EnterpriseRegisterAccountRequest) SetEncryptPassword(v string) *EnterpriseRegisterAccountRequest {
  s.EncryptPassword = &v
  return s
}

func (s *EnterpriseRegisterAccountRequest) SetEncryptTicket(v string) *EnterpriseRegisterAccountRequest {
  s.EncryptTicket = &v
  return s
}

func (s *EnterpriseRegisterAccountRequest) SetLoginEmail(v string) *EnterpriseRegisterAccountRequest {
  s.LoginEmail = &v
  return s
}

func (s *EnterpriseRegisterAccountRequest) SetOrganizationId(v string) *EnterpriseRegisterAccountRequest {
  s.OrganizationId = &v
  return s
}

func (s *EnterpriseRegisterAccountRequest) SetOrientedEcId(v string) *EnterpriseRegisterAccountRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseRegisterAccountRequest) SetOrientedLeId(v string) *EnterpriseRegisterAccountRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseRegisterAccountRequest) SetOrientedNbId(v string) *EnterpriseRegisterAccountRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseRegisterAccountRequest) SetRequestId(v string) *EnterpriseRegisterAccountRequest {
  s.RequestId = &v
  return s
}

func (s *EnterpriseRegisterAccountRequest) SetShowCompleteInfo(v bool) *EnterpriseRegisterAccountRequest {
  s.ShowCompleteInfo = &v
  return s
}

func (s *EnterpriseRegisterAccountRequest) SetSiteNick(v string) *EnterpriseRegisterAccountRequest {
  s.SiteNick = &v
  return s
}

func (s *EnterpriseRegisterAccountRequest) Validate() error {
  return dara.Validate(s)
}

