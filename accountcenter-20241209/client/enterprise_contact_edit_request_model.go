// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseContactEditRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *EnterpriseContactEditRequest
  GetAppName() *string 
  SetAsyncEmailVerify(v bool) *EnterpriseContactEditRequest
  GetAsyncEmailVerify() *bool 
  SetAsyncMobileVerify(v bool) *EnterpriseContactEditRequest
  GetAsyncMobileVerify() *bool 
  SetContactEmail(v string) *EnterpriseContactEditRequest
  GetContactEmail() *string 
  SetContactId(v int64) *EnterpriseContactEditRequest
  GetContactId() *int64 
  SetContactMobile(v string) *EnterpriseContactEditRequest
  GetContactMobile() *string 
  SetContactName(v string) *EnterpriseContactEditRequest
  GetContactName() *string 
  SetContactPosition(v string) *EnterpriseContactEditRequest
  GetContactPosition() *string 
  SetEmailCode(v string) *EnterpriseContactEditRequest
  GetEmailCode() *string 
  SetMobileCode(v string) *EnterpriseContactEditRequest
  GetMobileCode() *string 
  SetOrientedEcId(v string) *EnterpriseContactEditRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseContactEditRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseContactEditRequest
  GetOrientedNbId() *string 
  SetSharedContact(v bool) *EnterpriseContactEditRequest
  GetSharedContact() *bool 
}

type EnterpriseContactEditRequest struct {
  // The application name.
  // 
  // example:
  // 
  // xxx
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  // Specifies whether to asynchronously verify the email address. Set this parameter to true for API calls. Otherwise, the verification code is synchronously and strictly verified. Call SendAsyncEmailCaptcha to send a verification link.
  // 
  // example:
  // 
  // true
  AsyncEmailVerify *bool `json:"AsyncEmailVerify,omitempty" xml:"AsyncEmailVerify,omitempty"`
  // Specifies whether to asynchronously verify the mobile phone number. Set this parameter to true for API calls. Otherwise, the verification code is synchronously and strictly verified. Call SendAsyncMobileCaptcha to send a verification link.
  // 
  // example:
  // 
  // true
  AsyncMobileVerify *bool `json:"AsyncMobileVerify,omitempty" xml:"AsyncMobileVerify,omitempty"`
  // The email address of the contact.
  // 
  // example:
  // 
  // xxx@xxx.xx
  ContactEmail *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
  // The ID of the contact to modify. You can invoke the EnterpriseContactQueryPageList operation to perform a paged query of account contact information.
  // 
  // example:
  // 
  // xxx
  ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
  // The mobile phone number of the contact.
  // 
  // example:
  // 
  // 1xxxxxxxxxx
  ContactMobile *string `json:"ContactMobile,omitempty" xml:"ContactMobile,omitempty"`
  // The name of the contact.
  // 
  // example:
  // 
  // xxx
  ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
  // The position of the contact.
  // 
  // example:
  // 
  // 2
  ContactPosition *string `json:"ContactPosition,omitempty" xml:"ContactPosition,omitempty"`
  // The email verification code.
  // 
  // example:
  // 
  // null
  EmailCode *string `json:"EmailCode,omitempty" xml:"EmailCode,omitempty"`
  // The SMS verification code.
  // 
  // example:
  // 
  // null
  MobileCode *string `json:"MobileCode,omitempty" xml:"MobileCode,omitempty"`
  // The entity ID of the cross-enterprise management object.
  // 
  // example:
  // 
  // null
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  // The ID of the enterprise to which you have switched.
  // 
  // example:
  // 
  // null
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  // The sales channel ID of the cross-enterprise management object.
  // 
  // example:
  // 
  // null
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  // Specifies whether the contact is a public contact. This operation sets this parameter to true by default.
  // 
  // example:
  // 
  // true
  SharedContact *bool `json:"SharedContact,omitempty" xml:"SharedContact,omitempty"`
}

func (s EnterpriseContactEditRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseContactEditRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseContactEditRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EnterpriseContactEditRequest) GetAsyncEmailVerify() *bool  {
  return s.AsyncEmailVerify
}

func (s *EnterpriseContactEditRequest) GetAsyncMobileVerify() *bool  {
  return s.AsyncMobileVerify
}

func (s *EnterpriseContactEditRequest) GetContactEmail() *string  {
  return s.ContactEmail
}

func (s *EnterpriseContactEditRequest) GetContactId() *int64  {
  return s.ContactId
}

func (s *EnterpriseContactEditRequest) GetContactMobile() *string  {
  return s.ContactMobile
}

func (s *EnterpriseContactEditRequest) GetContactName() *string  {
  return s.ContactName
}

func (s *EnterpriseContactEditRequest) GetContactPosition() *string  {
  return s.ContactPosition
}

func (s *EnterpriseContactEditRequest) GetEmailCode() *string  {
  return s.EmailCode
}

func (s *EnterpriseContactEditRequest) GetMobileCode() *string  {
  return s.MobileCode
}

func (s *EnterpriseContactEditRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseContactEditRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseContactEditRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseContactEditRequest) GetSharedContact() *bool  {
  return s.SharedContact
}

func (s *EnterpriseContactEditRequest) SetAppName(v string) *EnterpriseContactEditRequest {
  s.AppName = &v
  return s
}

func (s *EnterpriseContactEditRequest) SetAsyncEmailVerify(v bool) *EnterpriseContactEditRequest {
  s.AsyncEmailVerify = &v
  return s
}

func (s *EnterpriseContactEditRequest) SetAsyncMobileVerify(v bool) *EnterpriseContactEditRequest {
  s.AsyncMobileVerify = &v
  return s
}

func (s *EnterpriseContactEditRequest) SetContactEmail(v string) *EnterpriseContactEditRequest {
  s.ContactEmail = &v
  return s
}

func (s *EnterpriseContactEditRequest) SetContactId(v int64) *EnterpriseContactEditRequest {
  s.ContactId = &v
  return s
}

func (s *EnterpriseContactEditRequest) SetContactMobile(v string) *EnterpriseContactEditRequest {
  s.ContactMobile = &v
  return s
}

func (s *EnterpriseContactEditRequest) SetContactName(v string) *EnterpriseContactEditRequest {
  s.ContactName = &v
  return s
}

func (s *EnterpriseContactEditRequest) SetContactPosition(v string) *EnterpriseContactEditRequest {
  s.ContactPosition = &v
  return s
}

func (s *EnterpriseContactEditRequest) SetEmailCode(v string) *EnterpriseContactEditRequest {
  s.EmailCode = &v
  return s
}

func (s *EnterpriseContactEditRequest) SetMobileCode(v string) *EnterpriseContactEditRequest {
  s.MobileCode = &v
  return s
}

func (s *EnterpriseContactEditRequest) SetOrientedEcId(v string) *EnterpriseContactEditRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseContactEditRequest) SetOrientedLeId(v string) *EnterpriseContactEditRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseContactEditRequest) SetOrientedNbId(v string) *EnterpriseContactEditRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseContactEditRequest) SetSharedContact(v bool) *EnterpriseContactEditRequest {
  s.SharedContact = &v
  return s
}

func (s *EnterpriseContactEditRequest) Validate() error {
  return dara.Validate(s)
}

