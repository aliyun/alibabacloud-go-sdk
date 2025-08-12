// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContactListByContactGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeContactListByContactGroupResponseBody
	GetCode() *string
	SetContacts(v *DescribeContactListByContactGroupResponseBodyContacts) *DescribeContactListByContactGroupResponseBody
	GetContacts() *DescribeContactListByContactGroupResponseBodyContacts
	SetMessage(v string) *DescribeContactListByContactGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeContactListByContactGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeContactListByContactGroupResponseBody
	GetSuccess() *bool
}

type DescribeContactListByContactGroupResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The alert contacts that receive alert notifications.
	Contacts *DescribeContactListByContactGroupResponseBodyContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// The group is not exists.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06D5ECC2-B9BE-42A4-8FA3-1A610FB08B83
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeContactListByContactGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactListByContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContactListByContactGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeContactListByContactGroupResponseBody) GetContacts() *DescribeContactListByContactGroupResponseBodyContacts {
	return s.Contacts
}

func (s *DescribeContactListByContactGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeContactListByContactGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeContactListByContactGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeContactListByContactGroupResponseBody) SetCode(v string) *DescribeContactListByContactGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeContactListByContactGroupResponseBody) SetContacts(v *DescribeContactListByContactGroupResponseBodyContacts) *DescribeContactListByContactGroupResponseBody {
	s.Contacts = v
	return s
}

func (s *DescribeContactListByContactGroupResponseBody) SetMessage(v string) *DescribeContactListByContactGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeContactListByContactGroupResponseBody) SetRequestId(v string) *DescribeContactListByContactGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContactListByContactGroupResponseBody) SetSuccess(v bool) *DescribeContactListByContactGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeContactListByContactGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeContactListByContactGroupResponseBodyContacts struct {
	Contact []*DescribeContactListByContactGroupResponseBodyContactsContact `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Repeated"`
}

func (s DescribeContactListByContactGroupResponseBodyContacts) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactListByContactGroupResponseBodyContacts) GoString() string {
	return s.String()
}

func (s *DescribeContactListByContactGroupResponseBodyContacts) GetContact() []*DescribeContactListByContactGroupResponseBodyContactsContact {
	return s.Contact
}

func (s *DescribeContactListByContactGroupResponseBodyContacts) SetContact(v []*DescribeContactListByContactGroupResponseBodyContactsContact) *DescribeContactListByContactGroupResponseBodyContacts {
	s.Contact = v
	return s
}

func (s *DescribeContactListByContactGroupResponseBodyContacts) Validate() error {
	return dara.Validate(s)
}

type DescribeContactListByContactGroupResponseBodyContactsContact struct {
	// The alert notification methods.
	Channels *DescribeContactListByContactGroupResponseBodyContactsContactChannels `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Struct"`
	// The time when the alert contact was created.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1552314252000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the alert contact.
	//
	// example:
	//
	// ECS
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The name of the alert contact.
	//
	// example:
	//
	// Alice
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The time when the alert contact was modified.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1552314252000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeContactListByContactGroupResponseBodyContactsContact) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactListByContactGroupResponseBodyContactsContact) GoString() string {
	return s.String()
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContact) GetChannels() *DescribeContactListByContactGroupResponseBodyContactsContactChannels {
	return s.Channels
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContact) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContact) GetDesc() *string {
	return s.Desc
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContact) GetName() *string {
	return s.Name
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContact) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContact) SetChannels(v *DescribeContactListByContactGroupResponseBodyContactsContactChannels) *DescribeContactListByContactGroupResponseBodyContactsContact {
	s.Channels = v
	return s
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContact) SetCreateTime(v int64) *DescribeContactListByContactGroupResponseBodyContactsContact {
	s.CreateTime = &v
	return s
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContact) SetDesc(v string) *DescribeContactListByContactGroupResponseBodyContactsContact {
	s.Desc = &v
	return s
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContact) SetName(v string) *DescribeContactListByContactGroupResponseBodyContactsContact {
	s.Name = &v
	return s
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContact) SetUpdateTime(v int64) *DescribeContactListByContactGroupResponseBodyContactsContact {
	s.UpdateTime = &v
	return s
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContact) Validate() error {
	return dara.Validate(s)
}

type DescribeContactListByContactGroupResponseBodyContactsContactChannels struct {
	// The TradeManager ID of the alert contact.
	//
	// >  This parameter can be returned only on the China site (aliyun.com).
	//
	// example:
	//
	// Alice
	AliIM *string `json:"AliIM,omitempty" xml:"AliIM,omitempty"`
	// The webhook URL of the DingTalk chatbot.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=9bf44f8189597d07dfdd7a123455ffc112****
	DingWebHook *string `json:"DingWebHook,omitempty" xml:"DingWebHook,omitempty"`
	// The email address of the alert contact.
	//
	// example:
	//
	// alice@example.com
	Mail *string `json:"Mail,omitempty" xml:"Mail,omitempty"`
	// The mobile number of the alert contact.
	//
	// >  This parameter can be returned only on the China site (aliyun.com).
	//
	// example:
	//
	// 1333333****
	SMS *string `json:"SMS,omitempty" xml:"SMS,omitempty"`
}

func (s DescribeContactListByContactGroupResponseBodyContactsContactChannels) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactListByContactGroupResponseBodyContactsContactChannels) GoString() string {
	return s.String()
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContactChannels) GetAliIM() *string {
	return s.AliIM
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContactChannels) GetDingWebHook() *string {
	return s.DingWebHook
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContactChannels) GetMail() *string {
	return s.Mail
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContactChannels) GetSMS() *string {
	return s.SMS
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContactChannels) SetAliIM(v string) *DescribeContactListByContactGroupResponseBodyContactsContactChannels {
	s.AliIM = &v
	return s
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContactChannels) SetDingWebHook(v string) *DescribeContactListByContactGroupResponseBodyContactsContactChannels {
	s.DingWebHook = &v
	return s
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContactChannels) SetMail(v string) *DescribeContactListByContactGroupResponseBodyContactsContactChannels {
	s.Mail = &v
	return s
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContactChannels) SetSMS(v string) *DescribeContactListByContactGroupResponseBodyContactsContactChannels {
	s.SMS = &v
	return s
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContactChannels) Validate() error {
	return dara.Validate(s)
}
