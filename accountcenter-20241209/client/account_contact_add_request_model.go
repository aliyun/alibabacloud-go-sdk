// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountContactAddRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *AccountContactAddRequest
	GetAppName() *string
	SetAsyncEmailVerify(v bool) *AccountContactAddRequest
	GetAsyncEmailVerify() *bool
	SetAsyncMobileVerify(v bool) *AccountContactAddRequest
	GetAsyncMobileVerify() *bool
	SetContactEmail(v string) *AccountContactAddRequest
	GetContactEmail() *string
	SetContactMobile(v string) *AccountContactAddRequest
	GetContactMobile() *string
	SetContactName(v string) *AccountContactAddRequest
	GetContactName() *string
	SetContactPosition(v string) *AccountContactAddRequest
	GetContactPosition() *string
	SetEmailCode(v string) *AccountContactAddRequest
	GetEmailCode() *string
	SetMobileCode(v string) *AccountContactAddRequest
	GetMobileCode() *string
	SetOrientedEcId(v string) *AccountContactAddRequest
	GetOrientedEcId() *string
	SetOrientedLeId(v string) *AccountContactAddRequest
	GetOrientedLeId() *string
	SetOrientedNbId(v string) *AccountContactAddRequest
	GetOrientedNbId() *string
	SetSharedContact(v bool) *AccountContactAddRequest
	GetSharedContact() *bool
}

type AccountContactAddRequest struct {
	// example:
	//
	// xxx
	AppName           *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AsyncEmailVerify  *bool   `json:"AsyncEmailVerify,omitempty" xml:"AsyncEmailVerify,omitempty"`
	AsyncMobileVerify *bool   `json:"AsyncMobileVerify,omitempty" xml:"AsyncMobileVerify,omitempty"`
	// example:
	//
	// xxx@xxx.xxx
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
	// 0
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
	// null
	OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	// example:
	//
	// null
	OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	// example:
	//
	// null
	OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	// example:
	//
	// false
	SharedContact *bool `json:"SharedContact,omitempty" xml:"SharedContact,omitempty"`
}

func (s AccountContactAddRequest) String() string {
	return dara.Prettify(s)
}

func (s AccountContactAddRequest) GoString() string {
	return s.String()
}

func (s *AccountContactAddRequest) GetAppName() *string {
	return s.AppName
}

func (s *AccountContactAddRequest) GetAsyncEmailVerify() *bool {
	return s.AsyncEmailVerify
}

func (s *AccountContactAddRequest) GetAsyncMobileVerify() *bool {
	return s.AsyncMobileVerify
}

func (s *AccountContactAddRequest) GetContactEmail() *string {
	return s.ContactEmail
}

func (s *AccountContactAddRequest) GetContactMobile() *string {
	return s.ContactMobile
}

func (s *AccountContactAddRequest) GetContactName() *string {
	return s.ContactName
}

func (s *AccountContactAddRequest) GetContactPosition() *string {
	return s.ContactPosition
}

func (s *AccountContactAddRequest) GetEmailCode() *string {
	return s.EmailCode
}

func (s *AccountContactAddRequest) GetMobileCode() *string {
	return s.MobileCode
}

func (s *AccountContactAddRequest) GetOrientedEcId() *string {
	return s.OrientedEcId
}

func (s *AccountContactAddRequest) GetOrientedLeId() *string {
	return s.OrientedLeId
}

func (s *AccountContactAddRequest) GetOrientedNbId() *string {
	return s.OrientedNbId
}

func (s *AccountContactAddRequest) GetSharedContact() *bool {
	return s.SharedContact
}

func (s *AccountContactAddRequest) SetAppName(v string) *AccountContactAddRequest {
	s.AppName = &v
	return s
}

func (s *AccountContactAddRequest) SetAsyncEmailVerify(v bool) *AccountContactAddRequest {
	s.AsyncEmailVerify = &v
	return s
}

func (s *AccountContactAddRequest) SetAsyncMobileVerify(v bool) *AccountContactAddRequest {
	s.AsyncMobileVerify = &v
	return s
}

func (s *AccountContactAddRequest) SetContactEmail(v string) *AccountContactAddRequest {
	s.ContactEmail = &v
	return s
}

func (s *AccountContactAddRequest) SetContactMobile(v string) *AccountContactAddRequest {
	s.ContactMobile = &v
	return s
}

func (s *AccountContactAddRequest) SetContactName(v string) *AccountContactAddRequest {
	s.ContactName = &v
	return s
}

func (s *AccountContactAddRequest) SetContactPosition(v string) *AccountContactAddRequest {
	s.ContactPosition = &v
	return s
}

func (s *AccountContactAddRequest) SetEmailCode(v string) *AccountContactAddRequest {
	s.EmailCode = &v
	return s
}

func (s *AccountContactAddRequest) SetMobileCode(v string) *AccountContactAddRequest {
	s.MobileCode = &v
	return s
}

func (s *AccountContactAddRequest) SetOrientedEcId(v string) *AccountContactAddRequest {
	s.OrientedEcId = &v
	return s
}

func (s *AccountContactAddRequest) SetOrientedLeId(v string) *AccountContactAddRequest {
	s.OrientedLeId = &v
	return s
}

func (s *AccountContactAddRequest) SetOrientedNbId(v string) *AccountContactAddRequest {
	s.OrientedNbId = &v
	return s
}

func (s *AccountContactAddRequest) SetSharedContact(v bool) *AccountContactAddRequest {
	s.SharedContact = &v
	return s
}

func (s *AccountContactAddRequest) Validate() error {
	return dara.Validate(s)
}
