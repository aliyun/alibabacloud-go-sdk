// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v int64) *GetContactResponseBody
	GetContactId() *int64
	SetEmail(v string) *GetContactResponseBody
	GetEmail() *string
	SetEmailStatus(v int32) *GetContactResponseBody
	GetEmailStatus() *int32
	SetIdCard(v string) *GetContactResponseBody
	GetIdCard() *string
	SetMobile(v string) *GetContactResponseBody
	GetMobile() *string
	SetMobileStatus(v int32) *GetContactResponseBody
	GetMobileStatus() *int32
	SetName(v string) *GetContactResponseBody
	GetName() *string
	SetRequestId(v string) *GetContactResponseBody
	GetRequestId() *string
	SetWebhooks(v string) *GetContactResponseBody
	GetWebhooks() *string
}

type GetContactResponseBody struct {
	// The contact ID.
	//
	// example:
	//
	// 1352570
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The email address of the contact.
	//
	// example:
	//
	// test@163.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Indicates whether the email address is verified.
	//
	// example:
	//
	// 1
	EmailStatus *int32 `json:"EmailStatus,omitempty" xml:"EmailStatus,omitempty"`
	// The ID card number of the contact. This parameter is required for the CFCA certificate brand and is not required for other brands.
	//
	// example:
	//
	// 142***************
	IdCard *string `json:"IdCard,omitempty" xml:"IdCard,omitempty"`
	// The phone number of the contact.
	//
	// example:
	//
	// 1510108****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// Indicates whether the phone number is verified.
	//
	// example:
	//
	// 1
	MobileStatus *int32 `json:"MobileStatus,omitempty" xml:"MobileStatus,omitempty"`
	// The name of the certificate contact.
	//
	// example:
	//
	// zhang san
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EECA10D5-BD0F-4EF1-B3EA-B4578E5C6F8E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The webhook URLs of DingTalk, WeCom, or Lark chatbots. The value is a string in list format.
	//
	// example:
	//
	// [\\"https://open.feishu.cn/open-apis/bot/v2/hook/dc1aa9b9-47cd-4b34-91ef-73c1034208e5\\"]
	Webhooks *string `json:"Webhooks,omitempty" xml:"Webhooks,omitempty"`
}

func (s GetContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetContactResponseBody) GoString() string {
	return s.String()
}

func (s *GetContactResponseBody) GetContactId() *int64 {
	return s.ContactId
}

func (s *GetContactResponseBody) GetEmail() *string {
	return s.Email
}

func (s *GetContactResponseBody) GetEmailStatus() *int32 {
	return s.EmailStatus
}

func (s *GetContactResponseBody) GetIdCard() *string {
	return s.IdCard
}

func (s *GetContactResponseBody) GetMobile() *string {
	return s.Mobile
}

func (s *GetContactResponseBody) GetMobileStatus() *int32 {
	return s.MobileStatus
}

func (s *GetContactResponseBody) GetName() *string {
	return s.Name
}

func (s *GetContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetContactResponseBody) GetWebhooks() *string {
	return s.Webhooks
}

func (s *GetContactResponseBody) SetContactId(v int64) *GetContactResponseBody {
	s.ContactId = &v
	return s
}

func (s *GetContactResponseBody) SetEmail(v string) *GetContactResponseBody {
	s.Email = &v
	return s
}

func (s *GetContactResponseBody) SetEmailStatus(v int32) *GetContactResponseBody {
	s.EmailStatus = &v
	return s
}

func (s *GetContactResponseBody) SetIdCard(v string) *GetContactResponseBody {
	s.IdCard = &v
	return s
}

func (s *GetContactResponseBody) SetMobile(v string) *GetContactResponseBody {
	s.Mobile = &v
	return s
}

func (s *GetContactResponseBody) SetMobileStatus(v int32) *GetContactResponseBody {
	s.MobileStatus = &v
	return s
}

func (s *GetContactResponseBody) SetName(v string) *GetContactResponseBody {
	s.Name = &v
	return s
}

func (s *GetContactResponseBody) SetRequestId(v string) *GetContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContactResponseBody) SetWebhooks(v string) *GetContactResponseBody {
	s.Webhooks = &v
	return s
}

func (s *GetContactResponseBody) Validate() error {
	return dara.Validate(s)
}
