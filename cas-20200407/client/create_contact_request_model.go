// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEmail(v string) *CreateContactRequest
	GetEmail() *string
	SetIdcard(v string) *CreateContactRequest
	GetIdcard() *string
	SetMobile(v string) *CreateContactRequest
	GetMobile() *string
	SetName(v string) *CreateContactRequest
	GetName() *string
	SetWebhooks(v string) *CreateContactRequest
	GetWebhooks() *string
}

type CreateContactRequest struct {
	// The email address of the contact.
	//
	// example:
	//
	// test@126.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The ID card number of the contact. This parameter is required for the CFCA certificate brand and is not required for other brands.
	//
	// example:
	//
	// 123
	Idcard *string `json:"Idcard,omitempty" xml:"Idcard,omitempty"`
	// The phone number of the contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1335678****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The name of the certificate contact.
	//
	// This parameter is required.
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

func (s CreateContactRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateContactRequest) GoString() string {
	return s.String()
}

func (s *CreateContactRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateContactRequest) GetIdcard() *string {
	return s.Idcard
}

func (s *CreateContactRequest) GetMobile() *string {
	return s.Mobile
}

func (s *CreateContactRequest) GetName() *string {
	return s.Name
}

func (s *CreateContactRequest) GetWebhooks() *string {
	return s.Webhooks
}

func (s *CreateContactRequest) SetEmail(v string) *CreateContactRequest {
	s.Email = &v
	return s
}

func (s *CreateContactRequest) SetIdcard(v string) *CreateContactRequest {
	s.Idcard = &v
	return s
}

func (s *CreateContactRequest) SetMobile(v string) *CreateContactRequest {
	s.Mobile = &v
	return s
}

func (s *CreateContactRequest) SetName(v string) *CreateContactRequest {
	s.Name = &v
	return s
}

func (s *CreateContactRequest) SetWebhooks(v string) *CreateContactRequest {
	s.Webhooks = &v
	return s
}

func (s *CreateContactRequest) Validate() error {
	return dara.Validate(s)
}
