// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseContactAddRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *EnterpriseContactAddRequest
  GetAppName() *string 
  SetAsyncEmailVerify(v bool) *EnterpriseContactAddRequest
  GetAsyncEmailVerify() *bool 
  SetAsyncMobileVerify(v bool) *EnterpriseContactAddRequest
  GetAsyncMobileVerify() *bool 
  SetContactEmail(v string) *EnterpriseContactAddRequest
  GetContactEmail() *string 
  SetContactMobile(v string) *EnterpriseContactAddRequest
  GetContactMobile() *string 
  SetContactName(v string) *EnterpriseContactAddRequest
  GetContactName() *string 
  SetContactPosition(v string) *EnterpriseContactAddRequest
  GetContactPosition() *string 
  SetEmailCode(v string) *EnterpriseContactAddRequest
  GetEmailCode() *string 
  SetMobileCode(v string) *EnterpriseContactAddRequest
  GetMobileCode() *string 
  SetOrientedEcId(v string) *EnterpriseContactAddRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseContactAddRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseContactAddRequest
  GetOrientedNbId() *string 
  SetSharedContact(v bool) *EnterpriseContactAddRequest
  GetSharedContact() *bool 
}

type EnterpriseContactAddRequest struct {
  // example:
  // 
  // xxx
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  // example:
  // 
  // true
  AsyncEmailVerify *bool `json:"AsyncEmailVerify,omitempty" xml:"AsyncEmailVerify,omitempty"`
  // example:
  // 
  // true
  AsyncMobileVerify *bool `json:"AsyncMobileVerify,omitempty" xml:"AsyncMobileVerify,omitempty"`
  // example:
  // 
  // xxx@xxx.xx
  ContactEmail *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
  // example:
  // 
  // 1xxxxxxxxxx
  ContactMobile *string `json:"ContactMobile,omitempty" xml:"ContactMobile,omitempty"`
  // example:
  // 
  // xxx
  ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
  // example:
  // 
  // 1
  ContactPosition *string `json:"ContactPosition,omitempty" xml:"ContactPosition,omitempty"`
  // example:
  // 
  // null
  EmailCode *string `json:"EmailCode,omitempty" xml:"EmailCode,omitempty"`
  // example:
  // 
  // null
  MobileCode *string `json:"MobileCode,omitempty" xml:"MobileCode,omitempty"`
  // example:
  // 
  // xxx
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  // example:
  // 
  // xxx
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  // example:
  // 
  // null
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  // example:
  // 
  // true
  SharedContact *bool `json:"SharedContact,omitempty" xml:"SharedContact,omitempty"`
}

func (s EnterpriseContactAddRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseContactAddRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseContactAddRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EnterpriseContactAddRequest) GetAsyncEmailVerify() *bool  {
  return s.AsyncEmailVerify
}

func (s *EnterpriseContactAddRequest) GetAsyncMobileVerify() *bool  {
  return s.AsyncMobileVerify
}

func (s *EnterpriseContactAddRequest) GetContactEmail() *string  {
  return s.ContactEmail
}

func (s *EnterpriseContactAddRequest) GetContactMobile() *string  {
  return s.ContactMobile
}

func (s *EnterpriseContactAddRequest) GetContactName() *string  {
  return s.ContactName
}

func (s *EnterpriseContactAddRequest) GetContactPosition() *string  {
  return s.ContactPosition
}

func (s *EnterpriseContactAddRequest) GetEmailCode() *string  {
  return s.EmailCode
}

func (s *EnterpriseContactAddRequest) GetMobileCode() *string  {
  return s.MobileCode
}

func (s *EnterpriseContactAddRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseContactAddRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseContactAddRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseContactAddRequest) GetSharedContact() *bool  {
  return s.SharedContact
}

func (s *EnterpriseContactAddRequest) SetAppName(v string) *EnterpriseContactAddRequest {
  s.AppName = &v
  return s
}

func (s *EnterpriseContactAddRequest) SetAsyncEmailVerify(v bool) *EnterpriseContactAddRequest {
  s.AsyncEmailVerify = &v
  return s
}

func (s *EnterpriseContactAddRequest) SetAsyncMobileVerify(v bool) *EnterpriseContactAddRequest {
  s.AsyncMobileVerify = &v
  return s
}

func (s *EnterpriseContactAddRequest) SetContactEmail(v string) *EnterpriseContactAddRequest {
  s.ContactEmail = &v
  return s
}

func (s *EnterpriseContactAddRequest) SetContactMobile(v string) *EnterpriseContactAddRequest {
  s.ContactMobile = &v
  return s
}

func (s *EnterpriseContactAddRequest) SetContactName(v string) *EnterpriseContactAddRequest {
  s.ContactName = &v
  return s
}

func (s *EnterpriseContactAddRequest) SetContactPosition(v string) *EnterpriseContactAddRequest {
  s.ContactPosition = &v
  return s
}

func (s *EnterpriseContactAddRequest) SetEmailCode(v string) *EnterpriseContactAddRequest {
  s.EmailCode = &v
  return s
}

func (s *EnterpriseContactAddRequest) SetMobileCode(v string) *EnterpriseContactAddRequest {
  s.MobileCode = &v
  return s
}

func (s *EnterpriseContactAddRequest) SetOrientedEcId(v string) *EnterpriseContactAddRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseContactAddRequest) SetOrientedLeId(v string) *EnterpriseContactAddRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseContactAddRequest) SetOrientedNbId(v string) *EnterpriseContactAddRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseContactAddRequest) SetSharedContact(v bool) *EnterpriseContactAddRequest {
  s.SharedContact = &v
  return s
}

func (s *EnterpriseContactAddRequest) Validate() error {
  return dara.Validate(s)
}

