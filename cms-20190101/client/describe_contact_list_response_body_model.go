// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContactListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeContactListResponseBody
	GetCode() *string
	SetContacts(v *DescribeContactListResponseBodyContacts) *DescribeContactListResponseBody
	GetContacts() *DescribeContactListResponseBodyContacts
	SetMessage(v string) *DescribeContactListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeContactListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeContactListResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribeContactListResponseBody
	GetTotal() *int32
}

type DescribeContactListResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code     *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Contacts *DescribeContactListResponseBodyContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06D5ECC2-B9BE-42A4-8FA3-1A610FB08B83
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 15
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeContactListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContactListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeContactListResponseBody) GetContacts() *DescribeContactListResponseBodyContacts {
	return s.Contacts
}

func (s *DescribeContactListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeContactListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeContactListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeContactListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeContactListResponseBody) SetCode(v string) *DescribeContactListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeContactListResponseBody) SetContacts(v *DescribeContactListResponseBodyContacts) *DescribeContactListResponseBody {
	s.Contacts = v
	return s
}

func (s *DescribeContactListResponseBody) SetMessage(v string) *DescribeContactListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeContactListResponseBody) SetRequestId(v string) *DescribeContactListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContactListResponseBody) SetSuccess(v bool) *DescribeContactListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeContactListResponseBody) SetTotal(v int32) *DescribeContactListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeContactListResponseBody) Validate() error {
	if s.Contacts != nil {
		if err := s.Contacts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeContactListResponseBodyContacts struct {
	Contact []*DescribeContactListResponseBodyContactsContact `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Repeated"`
}

func (s DescribeContactListResponseBodyContacts) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactListResponseBodyContacts) GoString() string {
	return s.String()
}

func (s *DescribeContactListResponseBodyContacts) GetContact() []*DescribeContactListResponseBodyContactsContact {
	return s.Contact
}

func (s *DescribeContactListResponseBodyContacts) SetContact(v []*DescribeContactListResponseBodyContactsContact) *DescribeContactListResponseBodyContacts {
	s.Contact = v
	return s
}

func (s *DescribeContactListResponseBodyContacts) Validate() error {
	if s.Contact != nil {
		for _, item := range s.Contact {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeContactListResponseBodyContactsContact struct {
	Channels      *DescribeContactListResponseBodyContactsContactChannels      `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Struct"`
	ChannelsState *DescribeContactListResponseBodyContactsContactChannelsState `json:"ChannelsState,omitempty" xml:"ChannelsState,omitempty" type:"Struct"`
	ContactGroups *DescribeContactListResponseBodyContactsContactContactGroups `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Struct"`
	CreateTime    *int64                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Desc          *string                                                      `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Lang          *string                                                      `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Name          *string                                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	UpdateTime    *int64                                                       `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeContactListResponseBodyContactsContact) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactListResponseBodyContactsContact) GoString() string {
	return s.String()
}

func (s *DescribeContactListResponseBodyContactsContact) GetChannels() *DescribeContactListResponseBodyContactsContactChannels {
	return s.Channels
}

func (s *DescribeContactListResponseBodyContactsContact) GetChannelsState() *DescribeContactListResponseBodyContactsContactChannelsState {
	return s.ChannelsState
}

func (s *DescribeContactListResponseBodyContactsContact) GetContactGroups() *DescribeContactListResponseBodyContactsContactContactGroups {
	return s.ContactGroups
}

func (s *DescribeContactListResponseBodyContactsContact) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeContactListResponseBodyContactsContact) GetDesc() *string {
	return s.Desc
}

func (s *DescribeContactListResponseBodyContactsContact) GetLang() *string {
	return s.Lang
}

func (s *DescribeContactListResponseBodyContactsContact) GetName() *string {
	return s.Name
}

func (s *DescribeContactListResponseBodyContactsContact) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeContactListResponseBodyContactsContact) SetChannels(v *DescribeContactListResponseBodyContactsContactChannels) *DescribeContactListResponseBodyContactsContact {
	s.Channels = v
	return s
}

func (s *DescribeContactListResponseBodyContactsContact) SetChannelsState(v *DescribeContactListResponseBodyContactsContactChannelsState) *DescribeContactListResponseBodyContactsContact {
	s.ChannelsState = v
	return s
}

func (s *DescribeContactListResponseBodyContactsContact) SetContactGroups(v *DescribeContactListResponseBodyContactsContactContactGroups) *DescribeContactListResponseBodyContactsContact {
	s.ContactGroups = v
	return s
}

func (s *DescribeContactListResponseBodyContactsContact) SetCreateTime(v int64) *DescribeContactListResponseBodyContactsContact {
	s.CreateTime = &v
	return s
}

func (s *DescribeContactListResponseBodyContactsContact) SetDesc(v string) *DescribeContactListResponseBodyContactsContact {
	s.Desc = &v
	return s
}

func (s *DescribeContactListResponseBodyContactsContact) SetLang(v string) *DescribeContactListResponseBodyContactsContact {
	s.Lang = &v
	return s
}

func (s *DescribeContactListResponseBodyContactsContact) SetName(v string) *DescribeContactListResponseBodyContactsContact {
	s.Name = &v
	return s
}

func (s *DescribeContactListResponseBodyContactsContact) SetUpdateTime(v int64) *DescribeContactListResponseBodyContactsContact {
	s.UpdateTime = &v
	return s
}

func (s *DescribeContactListResponseBodyContactsContact) Validate() error {
	if s.Channels != nil {
		if err := s.Channels.Validate(); err != nil {
			return err
		}
	}
	if s.ChannelsState != nil {
		if err := s.ChannelsState.Validate(); err != nil {
			return err
		}
	}
	if s.ContactGroups != nil {
		if err := s.ContactGroups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeContactListResponseBodyContactsContactChannels struct {
	AliIM       *string `json:"AliIM,omitempty" xml:"AliIM,omitempty"`
	DingWebHook *string `json:"DingWebHook,omitempty" xml:"DingWebHook,omitempty"`
	Mail        *string `json:"Mail,omitempty" xml:"Mail,omitempty"`
	SMS         *string `json:"SMS,omitempty" xml:"SMS,omitempty"`
}

func (s DescribeContactListResponseBodyContactsContactChannels) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactListResponseBodyContactsContactChannels) GoString() string {
	return s.String()
}

func (s *DescribeContactListResponseBodyContactsContactChannels) GetAliIM() *string {
	return s.AliIM
}

func (s *DescribeContactListResponseBodyContactsContactChannels) GetDingWebHook() *string {
	return s.DingWebHook
}

func (s *DescribeContactListResponseBodyContactsContactChannels) GetMail() *string {
	return s.Mail
}

func (s *DescribeContactListResponseBodyContactsContactChannels) GetSMS() *string {
	return s.SMS
}

func (s *DescribeContactListResponseBodyContactsContactChannels) SetAliIM(v string) *DescribeContactListResponseBodyContactsContactChannels {
	s.AliIM = &v
	return s
}

func (s *DescribeContactListResponseBodyContactsContactChannels) SetDingWebHook(v string) *DescribeContactListResponseBodyContactsContactChannels {
	s.DingWebHook = &v
	return s
}

func (s *DescribeContactListResponseBodyContactsContactChannels) SetMail(v string) *DescribeContactListResponseBodyContactsContactChannels {
	s.Mail = &v
	return s
}

func (s *DescribeContactListResponseBodyContactsContactChannels) SetSMS(v string) *DescribeContactListResponseBodyContactsContactChannels {
	s.SMS = &v
	return s
}

func (s *DescribeContactListResponseBodyContactsContactChannels) Validate() error {
	return dara.Validate(s)
}

type DescribeContactListResponseBodyContactsContactChannelsState struct {
	AliIM       *string `json:"AliIM,omitempty" xml:"AliIM,omitempty"`
	DingWebHook *string `json:"DingWebHook,omitempty" xml:"DingWebHook,omitempty"`
	Mail        *string `json:"Mail,omitempty" xml:"Mail,omitempty"`
	SMS         *string `json:"SMS,omitempty" xml:"SMS,omitempty"`
}

func (s DescribeContactListResponseBodyContactsContactChannelsState) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactListResponseBodyContactsContactChannelsState) GoString() string {
	return s.String()
}

func (s *DescribeContactListResponseBodyContactsContactChannelsState) GetAliIM() *string {
	return s.AliIM
}

func (s *DescribeContactListResponseBodyContactsContactChannelsState) GetDingWebHook() *string {
	return s.DingWebHook
}

func (s *DescribeContactListResponseBodyContactsContactChannelsState) GetMail() *string {
	return s.Mail
}

func (s *DescribeContactListResponseBodyContactsContactChannelsState) GetSMS() *string {
	return s.SMS
}

func (s *DescribeContactListResponseBodyContactsContactChannelsState) SetAliIM(v string) *DescribeContactListResponseBodyContactsContactChannelsState {
	s.AliIM = &v
	return s
}

func (s *DescribeContactListResponseBodyContactsContactChannelsState) SetDingWebHook(v string) *DescribeContactListResponseBodyContactsContactChannelsState {
	s.DingWebHook = &v
	return s
}

func (s *DescribeContactListResponseBodyContactsContactChannelsState) SetMail(v string) *DescribeContactListResponseBodyContactsContactChannelsState {
	s.Mail = &v
	return s
}

func (s *DescribeContactListResponseBodyContactsContactChannelsState) SetSMS(v string) *DescribeContactListResponseBodyContactsContactChannelsState {
	s.SMS = &v
	return s
}

func (s *DescribeContactListResponseBodyContactsContactChannelsState) Validate() error {
	return dara.Validate(s)
}

type DescribeContactListResponseBodyContactsContactContactGroups struct {
	ContactGroup []*string `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty" type:"Repeated"`
}

func (s DescribeContactListResponseBodyContactsContactContactGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactListResponseBodyContactsContactContactGroups) GoString() string {
	return s.String()
}

func (s *DescribeContactListResponseBodyContactsContactContactGroups) GetContactGroup() []*string {
	return s.ContactGroup
}

func (s *DescribeContactListResponseBodyContactsContactContactGroups) SetContactGroup(v []*string) *DescribeContactListResponseBodyContactsContactContactGroups {
	s.ContactGroup = v
	return s
}

func (s *DescribeContactListResponseBodyContactsContactContactGroups) Validate() error {
	return dara.Validate(s)
}
