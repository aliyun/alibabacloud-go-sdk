// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v int64) *UpdateContactRequest
	GetContactId() *int64
	SetEmail(v string) *UpdateContactRequest
	GetEmail() *string
	SetIdcard(v string) *UpdateContactRequest
	GetIdcard() *string
	SetMobile(v string) *UpdateContactRequest
	GetMobile() *string
	SetName(v string) *UpdateContactRequest
	GetName() *string
	SetWebhooks(v string) *UpdateContactRequest
	GetWebhooks() *string
}

type UpdateContactRequest struct {
	// The contact ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1397591
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The email address of the contact.
	//
	// example:
	//
	// test@136.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The ID card number of the contact. This parameter is required for the CFCA certificate brand and is not required for other brands.
	//
	// example:
	//
	// 142***************
	Idcard *string `json:"Idcard,omitempty" xml:"Idcard,omitempty"`
	// The phone number of the contact.
	//
	// example:
	//
	// 1510108***
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The name of the certificate contact.
	//
	// example:
	//
	// zhang san
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The webhook URLs of DingTalk, WeCom, or Lark chatbots. The value is a string in list format.
	//
	// example:
	//
	// [\\"https://open.feishu.cn/open-apis/bot/v2/hook/dc1aa9b9-47cd-4b34-91ef-73c1034208e5\\"]
	Webhooks *string `json:"Webhooks,omitempty" xml:"Webhooks,omitempty"`
}

func (s UpdateContactRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContactRequest) GoString() string {
	return s.String()
}

func (s *UpdateContactRequest) GetContactId() *int64 {
	return s.ContactId
}

func (s *UpdateContactRequest) GetEmail() *string {
	return s.Email
}

func (s *UpdateContactRequest) GetIdcard() *string {
	return s.Idcard
}

func (s *UpdateContactRequest) GetMobile() *string {
	return s.Mobile
}

func (s *UpdateContactRequest) GetName() *string {
	return s.Name
}

func (s *UpdateContactRequest) GetWebhooks() *string {
	return s.Webhooks
}

func (s *UpdateContactRequest) SetContactId(v int64) *UpdateContactRequest {
	s.ContactId = &v
	return s
}

func (s *UpdateContactRequest) SetEmail(v string) *UpdateContactRequest {
	s.Email = &v
	return s
}

func (s *UpdateContactRequest) SetIdcard(v string) *UpdateContactRequest {
	s.Idcard = &v
	return s
}

func (s *UpdateContactRequest) SetMobile(v string) *UpdateContactRequest {
	s.Mobile = &v
	return s
}

func (s *UpdateContactRequest) SetName(v string) *UpdateContactRequest {
	s.Name = &v
	return s
}

func (s *UpdateContactRequest) SetWebhooks(v string) *UpdateContactRequest {
	s.Webhooks = &v
	return s
}

func (s *UpdateContactRequest) Validate() error {
	return dara.Validate(s)
}
