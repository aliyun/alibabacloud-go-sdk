// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEmail(v string) *UpdateContactRequest
	GetEmail() *string
	SetImUserIds(v map[string]*string) *UpdateContactRequest
	GetImUserIds() map[string]*string
	SetLang(v string) *UpdateContactRequest
	GetLang() *string
	SetName(v string) *UpdateContactRequest
	GetName() *string
	SetPhone(v string) *UpdateContactRequest
	GetPhone() *string
	SetPhoneCode(v string) *UpdateContactRequest
	GetPhoneCode() *string
}

type UpdateContactRequest struct {
	// The email address of the contact.
	//
	// example:
	//
	// test@aliyun.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// The instant messaging (IM) user ID mappings. Multiple IM tools are supported.
	//
	// example:
	//
	// {"dingtalk":"abc123","wechat":"def456"}
	ImUserIds map[string]*string `json:"imUserIds,omitempty" xml:"imUserIds,omitempty"`
	// The language.
	//
	// example:
	//
	// zh_CN
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
	// The name of the contact.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The mobile phone number of the contact. Only Chinese mobile phone numbers are supported.
	//
	// example:
	//
	// 150123456789
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// The country code for the mobile phone number, typically +86 for China.
	//
	// example:
	//
	// +86
	PhoneCode *string `json:"phoneCode,omitempty" xml:"phoneCode,omitempty"`
}

func (s UpdateContactRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContactRequest) GoString() string {
	return s.String()
}

func (s *UpdateContactRequest) GetEmail() *string {
	return s.Email
}

func (s *UpdateContactRequest) GetImUserIds() map[string]*string {
	return s.ImUserIds
}

func (s *UpdateContactRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateContactRequest) GetName() *string {
	return s.Name
}

func (s *UpdateContactRequest) GetPhone() *string {
	return s.Phone
}

func (s *UpdateContactRequest) GetPhoneCode() *string {
	return s.PhoneCode
}

func (s *UpdateContactRequest) SetEmail(v string) *UpdateContactRequest {
	s.Email = &v
	return s
}

func (s *UpdateContactRequest) SetImUserIds(v map[string]*string) *UpdateContactRequest {
	s.ImUserIds = v
	return s
}

func (s *UpdateContactRequest) SetLang(v string) *UpdateContactRequest {
	s.Lang = &v
	return s
}

func (s *UpdateContactRequest) SetName(v string) *UpdateContactRequest {
	s.Name = &v
	return s
}

func (s *UpdateContactRequest) SetPhone(v string) *UpdateContactRequest {
	s.Phone = &v
	return s
}

func (s *UpdateContactRequest) SetPhoneCode(v string) *UpdateContactRequest {
	s.PhoneCode = &v
	return s
}

func (s *UpdateContactRequest) Validate() error {
	return dara.Validate(s)
}
