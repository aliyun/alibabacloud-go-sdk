// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *CreateContactRequest
	GetContactId() *string
	SetEmail(v string) *CreateContactRequest
	GetEmail() *string
	SetImUserIds(v map[string]*string) *CreateContactRequest
	GetImUserIds() map[string]*string
	SetLang(v string) *CreateContactRequest
	GetLang() *string
	SetName(v string) *CreateContactRequest
	GetName() *string
	SetPhone(v string) *CreateContactRequest
	GetPhone() *string
	SetWorkspace(v string) *CreateContactRequest
	GetWorkspace() *string
}

type CreateContactRequest struct {
	// The ID of the contact.
	//
	// example:
	//
	// testId
	ContactId *string `json:"contactId,omitempty" xml:"contactId,omitempty"`
	// The email address.
	//
	// example:
	//
	// test@aliyun.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// The instant messaging (IM) user ID mapping, which supports multi-platform integration. The key is the IM type (such as DING or WEIXIN), and the value is the corresponding user ID.
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
	// test_07
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The mobile phone number. Currently, only mobile phone numbers in the Chinese mainland are supported.
	//
	// example:
	//
	// 13012345678
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// default-cms-5884500827389836-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateContactRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateContactRequest) GoString() string {
	return s.String()
}

func (s *CreateContactRequest) GetContactId() *string {
	return s.ContactId
}

func (s *CreateContactRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateContactRequest) GetImUserIds() map[string]*string {
	return s.ImUserIds
}

func (s *CreateContactRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateContactRequest) GetName() *string {
	return s.Name
}

func (s *CreateContactRequest) GetPhone() *string {
	return s.Phone
}

func (s *CreateContactRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateContactRequest) SetContactId(v string) *CreateContactRequest {
	s.ContactId = &v
	return s
}

func (s *CreateContactRequest) SetEmail(v string) *CreateContactRequest {
	s.Email = &v
	return s
}

func (s *CreateContactRequest) SetImUserIds(v map[string]*string) *CreateContactRequest {
	s.ImUserIds = v
	return s
}

func (s *CreateContactRequest) SetLang(v string) *CreateContactRequest {
	s.Lang = &v
	return s
}

func (s *CreateContactRequest) SetName(v string) *CreateContactRequest {
	s.Name = &v
	return s
}

func (s *CreateContactRequest) SetPhone(v string) *CreateContactRequest {
	s.Phone = &v
	return s
}

func (s *CreateContactRequest) SetWorkspace(v string) *CreateContactRequest {
	s.Workspace = &v
	return s
}

func (s *CreateContactRequest) Validate() error {
	return dara.Validate(s)
}
