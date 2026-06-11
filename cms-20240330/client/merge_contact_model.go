// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMergeContact interface {
	dara.Model
	String() string
	GoString() string
	SetEmail(v string) *MergeContact
	GetEmail() *string
	SetEmailVerify(v bool) *MergeContact
	GetEmailVerify() *bool
	SetExtend(v map[string]interface{}) *MergeContact
	GetExtend() map[string]interface{}
	SetGmtCreate(v string) *MergeContact
	GetGmtCreate() *string
	SetGmtModified(v string) *MergeContact
	GetGmtModified() *string
	SetIdentifier(v string) *MergeContact
	GetIdentifier() *string
	SetLang(v string) *MergeContact
	GetLang() *string
	SetName(v string) *MergeContact
	GetName() *string
	SetPhone(v string) *MergeContact
	GetPhone() *string
	SetPhoneCode(v string) *MergeContact
	GetPhoneCode() *string
	SetPhoneVerify(v bool) *MergeContact
	GetPhoneVerify() *bool
	SetSource(v string) *MergeContact
	GetSource() *string
}

type MergeContact struct {
	// The email address.
	//
	// example:
	//
	// zhangsan@company.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// Specifies whether the email address is authenticated.
	//
	// example:
	//
	// true
	EmailVerify *bool `json:"emailVerify,omitempty" xml:"emailVerify,omitempty"`
	// An extension field used to store additional information.
	//
	// example:
	//
	// { "department": "运维部", "role": "工程师" }
	Extend map[string]interface{} `json:"extend,omitempty" xml:"extend,omitempty"`
	// The time when the contact was created.
	//
	// example:
	//
	// 2025-03-11T08:21:58.789Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The time when the contact was last modified.
	//
	// example:
	//
	// 2025-03-11T08:21:58.789Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The unique identifier of the contact.
	//
	// example:
	//
	// user-12345
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The language preference.
	//
	// example:
	//
	// zh-CN
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
	// The name of the contact.
	//
	// example:
	//
	// 张三
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The mobile phone number.
	//
	// example:
	//
	// 13800138000
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// The country code for the mobile phone number.
	//
	// example:
	//
	// 86
	PhoneCode *string `json:"phoneCode,omitempty" xml:"phoneCode,omitempty"`
	// Specifies whether the mobile phone number is authenticated.
	//
	// example:
	//
	// true
	PhoneVerify *bool `json:"phoneVerify,omitempty" xml:"phoneVerify,omitempty"`
	// The source system of the contact.
	//
	// example:
	//
	// dingtalk
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s MergeContact) String() string {
	return dara.Prettify(s)
}

func (s MergeContact) GoString() string {
	return s.String()
}

func (s *MergeContact) GetEmail() *string {
	return s.Email
}

func (s *MergeContact) GetEmailVerify() *bool {
	return s.EmailVerify
}

func (s *MergeContact) GetExtend() map[string]interface{} {
	return s.Extend
}

func (s *MergeContact) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *MergeContact) GetGmtModified() *string {
	return s.GmtModified
}

func (s *MergeContact) GetIdentifier() *string {
	return s.Identifier
}

func (s *MergeContact) GetLang() *string {
	return s.Lang
}

func (s *MergeContact) GetName() *string {
	return s.Name
}

func (s *MergeContact) GetPhone() *string {
	return s.Phone
}

func (s *MergeContact) GetPhoneCode() *string {
	return s.PhoneCode
}

func (s *MergeContact) GetPhoneVerify() *bool {
	return s.PhoneVerify
}

func (s *MergeContact) GetSource() *string {
	return s.Source
}

func (s *MergeContact) SetEmail(v string) *MergeContact {
	s.Email = &v
	return s
}

func (s *MergeContact) SetEmailVerify(v bool) *MergeContact {
	s.EmailVerify = &v
	return s
}

func (s *MergeContact) SetExtend(v map[string]interface{}) *MergeContact {
	s.Extend = v
	return s
}

func (s *MergeContact) SetGmtCreate(v string) *MergeContact {
	s.GmtCreate = &v
	return s
}

func (s *MergeContact) SetGmtModified(v string) *MergeContact {
	s.GmtModified = &v
	return s
}

func (s *MergeContact) SetIdentifier(v string) *MergeContact {
	s.Identifier = &v
	return s
}

func (s *MergeContact) SetLang(v string) *MergeContact {
	s.Lang = &v
	return s
}

func (s *MergeContact) SetName(v string) *MergeContact {
	s.Name = &v
	return s
}

func (s *MergeContact) SetPhone(v string) *MergeContact {
	s.Phone = &v
	return s
}

func (s *MergeContact) SetPhoneCode(v string) *MergeContact {
	s.PhoneCode = &v
	return s
}

func (s *MergeContact) SetPhoneVerify(v bool) *MergeContact {
	s.PhoneVerify = &v
	return s
}

func (s *MergeContact) SetSource(v string) *MergeContact {
	s.Source = &v
	return s
}

func (s *MergeContact) Validate() error {
	return dara.Validate(s)
}
