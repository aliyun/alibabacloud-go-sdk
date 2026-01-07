// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountContactEditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *AccountContactEditRequest
	GetAppName() *string
	SetAsyncEmailVerify(v bool) *AccountContactEditRequest
	GetAsyncEmailVerify() *bool
	SetAsyncMobileVerify(v bool) *AccountContactEditRequest
	GetAsyncMobileVerify() *bool
	SetContactEmail(v string) *AccountContactEditRequest
	GetContactEmail() *string
	SetContactId(v int64) *AccountContactEditRequest
	GetContactId() *int64
	SetContactMobile(v string) *AccountContactEditRequest
	GetContactMobile() *string
	SetContactName(v string) *AccountContactEditRequest
	GetContactName() *string
	SetContactPosition(v string) *AccountContactEditRequest
	GetContactPosition() *string
	SetEmailCode(v string) *AccountContactEditRequest
	GetEmailCode() *string
	SetMobileCode(v string) *AccountContactEditRequest
	GetMobileCode() *string
	SetOrientedEcId(v string) *AccountContactEditRequest
	GetOrientedEcId() *string
	SetOrientedLeId(v string) *AccountContactEditRequest
	GetOrientedLeId() *string
	SetOrientedNbId(v string) *AccountContactEditRequest
	GetOrientedNbId() *string
	SetSharedContact(v bool) *AccountContactEditRequest
	GetSharedContact() *bool
}

type AccountContactEditRequest struct {
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
	// xxx
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
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

func (s AccountContactEditRequest) String() string {
	return dara.Prettify(s)
}

func (s AccountContactEditRequest) GoString() string {
	return s.String()
}

func (s *AccountContactEditRequest) GetAppName() *string {
	return s.AppName
}

func (s *AccountContactEditRequest) GetAsyncEmailVerify() *bool {
	return s.AsyncEmailVerify
}

func (s *AccountContactEditRequest) GetAsyncMobileVerify() *bool {
	return s.AsyncMobileVerify
}

func (s *AccountContactEditRequest) GetContactEmail() *string {
	return s.ContactEmail
}

func (s *AccountContactEditRequest) GetContactId() *int64 {
	return s.ContactId
}

func (s *AccountContactEditRequest) GetContactMobile() *string {
	return s.ContactMobile
}

func (s *AccountContactEditRequest) GetContactName() *string {
	return s.ContactName
}

func (s *AccountContactEditRequest) GetContactPosition() *string {
	return s.ContactPosition
}

func (s *AccountContactEditRequest) GetEmailCode() *string {
	return s.EmailCode
}

func (s *AccountContactEditRequest) GetMobileCode() *string {
	return s.MobileCode
}

func (s *AccountContactEditRequest) GetOrientedEcId() *string {
	return s.OrientedEcId
}

func (s *AccountContactEditRequest) GetOrientedLeId() *string {
	return s.OrientedLeId
}

func (s *AccountContactEditRequest) GetOrientedNbId() *string {
	return s.OrientedNbId
}

func (s *AccountContactEditRequest) GetSharedContact() *bool {
	return s.SharedContact
}

func (s *AccountContactEditRequest) SetAppName(v string) *AccountContactEditRequest {
	s.AppName = &v
	return s
}

func (s *AccountContactEditRequest) SetAsyncEmailVerify(v bool) *AccountContactEditRequest {
	s.AsyncEmailVerify = &v
	return s
}

func (s *AccountContactEditRequest) SetAsyncMobileVerify(v bool) *AccountContactEditRequest {
	s.AsyncMobileVerify = &v
	return s
}

func (s *AccountContactEditRequest) SetContactEmail(v string) *AccountContactEditRequest {
	s.ContactEmail = &v
	return s
}

func (s *AccountContactEditRequest) SetContactId(v int64) *AccountContactEditRequest {
	s.ContactId = &v
	return s
}

func (s *AccountContactEditRequest) SetContactMobile(v string) *AccountContactEditRequest {
	s.ContactMobile = &v
	return s
}

func (s *AccountContactEditRequest) SetContactName(v string) *AccountContactEditRequest {
	s.ContactName = &v
	return s
}

func (s *AccountContactEditRequest) SetContactPosition(v string) *AccountContactEditRequest {
	s.ContactPosition = &v
	return s
}

func (s *AccountContactEditRequest) SetEmailCode(v string) *AccountContactEditRequest {
	s.EmailCode = &v
	return s
}

func (s *AccountContactEditRequest) SetMobileCode(v string) *AccountContactEditRequest {
	s.MobileCode = &v
	return s
}

func (s *AccountContactEditRequest) SetOrientedEcId(v string) *AccountContactEditRequest {
	s.OrientedEcId = &v
	return s
}

func (s *AccountContactEditRequest) SetOrientedLeId(v string) *AccountContactEditRequest {
	s.OrientedLeId = &v
	return s
}

func (s *AccountContactEditRequest) SetOrientedNbId(v string) *AccountContactEditRequest {
	s.OrientedNbId = &v
	return s
}

func (s *AccountContactEditRequest) SetSharedContact(v bool) *AccountContactEditRequest {
	s.SharedContact = &v
	return s
}

func (s *AccountContactEditRequest) Validate() error {
	return dara.Validate(s)
}
