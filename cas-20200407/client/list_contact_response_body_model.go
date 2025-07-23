// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContactList(v []*ListContactResponseBodyContactList) *ListContactResponseBody
	GetContactList() []*ListContactResponseBodyContactList
	SetCurrentPage(v int32) *ListContactResponseBody
	GetCurrentPage() *int32
	SetKeyword(v string) *ListContactResponseBody
	GetKeyword() *string
	SetRequestId(v string) *ListContactResponseBody
	GetRequestId() *string
	SetShowSize(v int32) *ListContactResponseBody
	GetShowSize() *int32
	SetTotalCount(v int64) *ListContactResponseBody
	GetTotalCount() *int64
}

type ListContactResponseBody struct {
	// The contacts.
	ContactList []*ListContactResponseBodyContactList `json:"ContactList,omitempty" xml:"ContactList,omitempty" type:"Repeated"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The keyword used in the fuzzy search.
	//
	// example:
	//
	// 186
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 31C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of certificates per page. Default value: **20**.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListContactResponseBody) GoString() string {
	return s.String()
}

func (s *ListContactResponseBody) GetContactList() []*ListContactResponseBodyContactList {
	return s.ContactList
}

func (s *ListContactResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListContactResponseBody) GetKeyword() *string {
	return s.Keyword
}

func (s *ListContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListContactResponseBody) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListContactResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListContactResponseBody) SetContactList(v []*ListContactResponseBodyContactList) *ListContactResponseBody {
	s.ContactList = v
	return s
}

func (s *ListContactResponseBody) SetCurrentPage(v int32) *ListContactResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListContactResponseBody) SetKeyword(v string) *ListContactResponseBody {
	s.Keyword = &v
	return s
}

func (s *ListContactResponseBody) SetRequestId(v string) *ListContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContactResponseBody) SetShowSize(v int32) *ListContactResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListContactResponseBody) SetTotalCount(v int64) *ListContactResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListContactResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListContactResponseBodyContactList struct {
	// The ID of the contact.
	//
	// example:
	//
	// 519580
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The email address of the contact.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Indicates whether the email address passed the verification.
	//
	// example:
	//
	// 1
	EmailStatus *int32 `json:"EmailStatus,omitempty" xml:"EmailStatus,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 139****8888
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// Indicates whether the phone number was verified.
	//
	// example:
	//
	// 1
	MobileStatus *int32 `json:"MobileStatus,omitempty" xml:"MobileStatus,omitempty"`
	// The name of the contact.
	//
	// example:
	//
	// ty-yaoyue.com
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The webhook URL of the chatbot.
	//
	// example:
	//
	// [\\"https://open.feishu.cn/open-apis/bot/v2/hook/XXX\\",\\"https://oapi.dingtalk.com/robot/send?access_token=XXX\\",\\"https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=XXX\\"]
	Webhooks *string `json:"Webhooks,omitempty" xml:"Webhooks,omitempty"`
}

func (s ListContactResponseBodyContactList) String() string {
	return dara.Prettify(s)
}

func (s ListContactResponseBodyContactList) GoString() string {
	return s.String()
}

func (s *ListContactResponseBodyContactList) GetContactId() *int64 {
	return s.ContactId
}

func (s *ListContactResponseBodyContactList) GetEmail() *string {
	return s.Email
}

func (s *ListContactResponseBodyContactList) GetEmailStatus() *int32 {
	return s.EmailStatus
}

func (s *ListContactResponseBodyContactList) GetMobile() *string {
	return s.Mobile
}

func (s *ListContactResponseBodyContactList) GetMobileStatus() *int32 {
	return s.MobileStatus
}

func (s *ListContactResponseBodyContactList) GetName() *string {
	return s.Name
}

func (s *ListContactResponseBodyContactList) GetWebhooks() *string {
	return s.Webhooks
}

func (s *ListContactResponseBodyContactList) SetContactId(v int64) *ListContactResponseBodyContactList {
	s.ContactId = &v
	return s
}

func (s *ListContactResponseBodyContactList) SetEmail(v string) *ListContactResponseBodyContactList {
	s.Email = &v
	return s
}

func (s *ListContactResponseBodyContactList) SetEmailStatus(v int32) *ListContactResponseBodyContactList {
	s.EmailStatus = &v
	return s
}

func (s *ListContactResponseBodyContactList) SetMobile(v string) *ListContactResponseBodyContactList {
	s.Mobile = &v
	return s
}

func (s *ListContactResponseBodyContactList) SetMobileStatus(v int32) *ListContactResponseBodyContactList {
	s.MobileStatus = &v
	return s
}

func (s *ListContactResponseBodyContactList) SetName(v string) *ListContactResponseBodyContactList {
	s.Name = &v
	return s
}

func (s *ListContactResponseBodyContactList) SetWebhooks(v string) *ListContactResponseBodyContactList {
	s.Webhooks = &v
	return s
}

func (s *ListContactResponseBodyContactList) Validate() error {
	return dara.Validate(s)
}
