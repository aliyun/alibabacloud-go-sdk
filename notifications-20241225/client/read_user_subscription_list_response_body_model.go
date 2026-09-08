// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadUserSubscriptionListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReadUserSubscriptionListResponseBody
	GetCode() *string
	SetData(v []*ReadUserSubscriptionListResponseBodyData) *ReadUserSubscriptionListResponseBody
	GetData() []*ReadUserSubscriptionListResponseBodyData
	SetMessage(v string) *ReadUserSubscriptionListResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadUserSubscriptionListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadUserSubscriptionListResponseBody
	GetSuccess() *bool
}

type ReadUserSubscriptionListResponseBody struct {
	// The response code of the operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The query result.
	Data []*ReadUserSubscriptionListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The message.
	//
	// example:
	//
	// Succeeded
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// /
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadUserSubscriptionListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadUserSubscriptionListResponseBody) GoString() string {
	return s.String()
}

func (s *ReadUserSubscriptionListResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadUserSubscriptionListResponseBody) GetData() []*ReadUserSubscriptionListResponseBodyData {
	return s.Data
}

func (s *ReadUserSubscriptionListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadUserSubscriptionListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadUserSubscriptionListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadUserSubscriptionListResponseBody) SetCode(v string) *ReadUserSubscriptionListResponseBody {
	s.Code = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBody) SetData(v []*ReadUserSubscriptionListResponseBodyData) *ReadUserSubscriptionListResponseBody {
	s.Data = v
	return s
}

func (s *ReadUserSubscriptionListResponseBody) SetMessage(v string) *ReadUserSubscriptionListResponseBody {
	s.Message = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBody) SetRequestId(v string) *ReadUserSubscriptionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBody) SetSuccess(v bool) *ReadUserSubscriptionListResponseBody {
	s.Success = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReadUserSubscriptionListResponseBodyData struct {
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// /
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The message category code.
	//
	// example:
	//
	// prod_edu_content
	CategoryCode *string `json:"CategoryCode,omitempty" xml:"CategoryCode,omitempty"`
	// The description of the message category.
	//
	// example:
	//
	// Content related to product usage scenarios and technical sharing
	CategoryDesc *string `json:"CategoryDesc,omitempty" xml:"CategoryDesc,omitempty"`
	// The category group code.
	//
	// example:
	//
	// prod_msg
	CategoryGroupCode *string `json:"CategoryGroupCode,omitempty" xml:"CategoryGroupCode,omitempty"`
	// The category group name.
	//
	// example:
	//
	// Product Messages
	CategoryGroupName *string `json:"CategoryGroupName,omitempty" xml:"CategoryGroupName,omitempty"`
	// The message category name.
	//
	// example:
	//
	// Product Educational Content
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The channel list.
	ChannelConfigs []*ReadUserSubscriptionListResponseBodyDataChannelConfigs `json:"ChannelConfigs,omitempty" xml:"ChannelConfigs,omitempty" type:"Repeated"`
	// The contact.
	Contact *ReadUserSubscriptionListResponseBodyDataContact `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Struct"`
	// The receiving time list.
	ReceiveTimeList []*int32 `json:"ReceiveTimeList,omitempty" xml:"ReceiveTimeList,omitempty" type:"Repeated"`
}

func (s ReadUserSubscriptionListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReadUserSubscriptionListResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadUserSubscriptionListResponseBodyData) GetAliUid() *int64 {
	return s.AliUid
}

func (s *ReadUserSubscriptionListResponseBodyData) GetCategoryCode() *string {
	return s.CategoryCode
}

func (s *ReadUserSubscriptionListResponseBodyData) GetCategoryDesc() *string {
	return s.CategoryDesc
}

func (s *ReadUserSubscriptionListResponseBodyData) GetCategoryGroupCode() *string {
	return s.CategoryGroupCode
}

func (s *ReadUserSubscriptionListResponseBodyData) GetCategoryGroupName() *string {
	return s.CategoryGroupName
}

func (s *ReadUserSubscriptionListResponseBodyData) GetCategoryName() *string {
	return s.CategoryName
}

func (s *ReadUserSubscriptionListResponseBodyData) GetChannelConfigs() []*ReadUserSubscriptionListResponseBodyDataChannelConfigs {
	return s.ChannelConfigs
}

func (s *ReadUserSubscriptionListResponseBodyData) GetContact() *ReadUserSubscriptionListResponseBodyDataContact {
	return s.Contact
}

func (s *ReadUserSubscriptionListResponseBodyData) GetReceiveTimeList() []*int32 {
	return s.ReceiveTimeList
}

func (s *ReadUserSubscriptionListResponseBodyData) SetAliUid(v int64) *ReadUserSubscriptionListResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyData) SetCategoryCode(v string) *ReadUserSubscriptionListResponseBodyData {
	s.CategoryCode = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyData) SetCategoryDesc(v string) *ReadUserSubscriptionListResponseBodyData {
	s.CategoryDesc = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyData) SetCategoryGroupCode(v string) *ReadUserSubscriptionListResponseBodyData {
	s.CategoryGroupCode = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyData) SetCategoryGroupName(v string) *ReadUserSubscriptionListResponseBodyData {
	s.CategoryGroupName = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyData) SetCategoryName(v string) *ReadUserSubscriptionListResponseBodyData {
	s.CategoryName = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyData) SetChannelConfigs(v []*ReadUserSubscriptionListResponseBodyDataChannelConfigs) *ReadUserSubscriptionListResponseBodyData {
	s.ChannelConfigs = v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyData) SetContact(v *ReadUserSubscriptionListResponseBodyDataContact) *ReadUserSubscriptionListResponseBodyData {
	s.Contact = v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyData) SetReceiveTimeList(v []*int32) *ReadUserSubscriptionListResponseBodyData {
	s.ReceiveTimeList = v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyData) Validate() error {
	if s.ChannelConfigs != nil {
		for _, item := range s.ChannelConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Contact != nil {
		if err := s.Contact.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReadUserSubscriptionListResponseBodyDataChannelConfigs struct {
	// The channel type.
	//
	// example:
	//
	// email
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// Indicates whether the subscription is configured.
	//
	// example:
	//
	// YES
	CheckedState *string `json:"CheckedState,omitempty" xml:"CheckedState,omitempty"`
	// Indicates whether the option is selected by default.
	//
	// example:
	//
	// YES
	DefaultChecked *string `json:"DefaultChecked,omitempty" xml:"DefaultChecked,omitempty"`
	// The fatigue limit.
	//
	// example:
	//
	// 7
	FatigueDayLimit *int32 `json:"FatigueDayLimit,omitempty" xml:"FatigueDayLimit,omitempty"`
	// Indicates whether the option can be modified.
	//
	// example:
	//
	// YES
	Optional *string `json:"Optional,omitempty" xml:"Optional,omitempty"`
}

func (s ReadUserSubscriptionListResponseBodyDataChannelConfigs) String() string {
	return dara.Prettify(s)
}

func (s ReadUserSubscriptionListResponseBodyDataChannelConfigs) GoString() string {
	return s.String()
}

func (s *ReadUserSubscriptionListResponseBodyDataChannelConfigs) GetChannelType() *string {
	return s.ChannelType
}

func (s *ReadUserSubscriptionListResponseBodyDataChannelConfigs) GetCheckedState() *string {
	return s.CheckedState
}

func (s *ReadUserSubscriptionListResponseBodyDataChannelConfigs) GetDefaultChecked() *string {
	return s.DefaultChecked
}

func (s *ReadUserSubscriptionListResponseBodyDataChannelConfigs) GetFatigueDayLimit() *int32 {
	return s.FatigueDayLimit
}

func (s *ReadUserSubscriptionListResponseBodyDataChannelConfigs) GetOptional() *string {
	return s.Optional
}

func (s *ReadUserSubscriptionListResponseBodyDataChannelConfigs) SetChannelType(v string) *ReadUserSubscriptionListResponseBodyDataChannelConfigs {
	s.ChannelType = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataChannelConfigs) SetCheckedState(v string) *ReadUserSubscriptionListResponseBodyDataChannelConfigs {
	s.CheckedState = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataChannelConfigs) SetDefaultChecked(v string) *ReadUserSubscriptionListResponseBodyDataChannelConfigs {
	s.DefaultChecked = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataChannelConfigs) SetFatigueDayLimit(v int32) *ReadUserSubscriptionListResponseBodyDataChannelConfigs {
	s.FatigueDayLimit = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataChannelConfigs) SetOptional(v string) *ReadUserSubscriptionListResponseBodyDataChannelConfigs {
	s.Optional = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataChannelConfigs) Validate() error {
	return dara.Validate(s)
}

type ReadUserSubscriptionListResponseBodyDataContact struct {
	// The Account Center contact list.
	CommonContacts []*ReadUserSubscriptionListResponseBodyDataContactCommonContacts `json:"CommonContacts,omitempty" xml:"CommonContacts,omitempty" type:"Repeated"`
	// The webhook contact list.
	WebhookContacts []*ReadUserSubscriptionListResponseBodyDataContactWebhookContacts `json:"WebhookContacts,omitempty" xml:"WebhookContacts,omitempty" type:"Repeated"`
}

func (s ReadUserSubscriptionListResponseBodyDataContact) String() string {
	return dara.Prettify(s)
}

func (s ReadUserSubscriptionListResponseBodyDataContact) GoString() string {
	return s.String()
}

func (s *ReadUserSubscriptionListResponseBodyDataContact) GetCommonContacts() []*ReadUserSubscriptionListResponseBodyDataContactCommonContacts {
	return s.CommonContacts
}

func (s *ReadUserSubscriptionListResponseBodyDataContact) GetWebhookContacts() []*ReadUserSubscriptionListResponseBodyDataContactWebhookContacts {
	return s.WebhookContacts
}

func (s *ReadUserSubscriptionListResponseBodyDataContact) SetCommonContacts(v []*ReadUserSubscriptionListResponseBodyDataContactCommonContacts) *ReadUserSubscriptionListResponseBodyDataContact {
	s.CommonContacts = v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataContact) SetWebhookContacts(v []*ReadUserSubscriptionListResponseBodyDataContactWebhookContacts) *ReadUserSubscriptionListResponseBodyDataContact {
	s.WebhookContacts = v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataContact) Validate() error {
	if s.CommonContacts != nil {
		for _, item := range s.CommonContacts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WebhookContacts != nil {
		for _, item := range s.WebhookContacts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReadUserSubscriptionListResponseBodyDataContactCommonContacts struct {
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// /
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The email address of the contact.
	//
	// example:
	//
	// t*@qq.*
	ContactEmail *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	// The contact ID.
	//
	// example:
	//
	// 0
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The masked mobile phone number of the Account Center contact.
	//
	// example:
	//
	// 130*90
	ContactMobile *string `json:"ContactMobile,omitempty" xml:"ContactMobile,omitempty"`
	// The name of the Account Center contact.
	//
	// example:
	//
	// test
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// Indicates whether the email address is verified.
	//
	// example:
	//
	// true
	EmailConfirmed *bool `json:"EmailConfirmed,omitempty" xml:"EmailConfirmed,omitempty"`
	// The message source.
	MessageSource *ReadUserSubscriptionListResponseBodyDataContactCommonContactsMessageSource `json:"MessageSource,omitempty" xml:"MessageSource,omitempty" type:"Struct"`
	// Indicates whether the mobile phone number of the Account Center contact is verified.
	//
	// example:
	//
	// true
	MobileConfirmed *bool `json:"MobileConfirmed,omitempty" xml:"MobileConfirmed,omitempty"`
	// The position of the Account Center contact.
	//
	// example:
	//
	// CEO
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s ReadUserSubscriptionListResponseBodyDataContactCommonContacts) String() string {
	return dara.Prettify(s)
}

func (s ReadUserSubscriptionListResponseBodyDataContactCommonContacts) GoString() string {
	return s.String()
}

func (s *ReadUserSubscriptionListResponseBodyDataContactCommonContacts) GetAliUid() *int64 {
	return s.AliUid
}

func (s *ReadUserSubscriptionListResponseBodyDataContactCommonContacts) GetContactEmail() *string {
	return s.ContactEmail
}

func (s *ReadUserSubscriptionListResponseBodyDataContactCommonContacts) GetContactId() *int64 {
	return s.ContactId
}

func (s *ReadUserSubscriptionListResponseBodyDataContactCommonContacts) GetContactMobile() *string {
	return s.ContactMobile
}

func (s *ReadUserSubscriptionListResponseBodyDataContactCommonContacts) GetContactName() *string {
	return s.ContactName
}

func (s *ReadUserSubscriptionListResponseBodyDataContactCommonContacts) GetEmailConfirmed() *bool {
	return s.EmailConfirmed
}

func (s *ReadUserSubscriptionListResponseBodyDataContactCommonContacts) GetMessageSource() *ReadUserSubscriptionListResponseBodyDataContactCommonContactsMessageSource {
	return s.MessageSource
}

func (s *ReadUserSubscriptionListResponseBodyDataContactCommonContacts) GetMobileConfirmed() *bool {
	return s.MobileConfirmed
}

func (s *ReadUserSubscriptionListResponseBodyDataContactCommonContacts) GetPosition() *string {
	return s.Position
}

func (s *ReadUserSubscriptionListResponseBodyDataContactCommonContacts) SetAliUid(v int64) *ReadUserSubscriptionListResponseBodyDataContactCommonContacts {
	s.AliUid = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataContactCommonContacts) SetContactEmail(v string) *ReadUserSubscriptionListResponseBodyDataContactCommonContacts {
	s.ContactEmail = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataContactCommonContacts) SetContactId(v int64) *ReadUserSubscriptionListResponseBodyDataContactCommonContacts {
	s.ContactId = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataContactCommonContacts) SetContactMobile(v string) *ReadUserSubscriptionListResponseBodyDataContactCommonContacts {
	s.ContactMobile = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataContactCommonContacts) SetContactName(v string) *ReadUserSubscriptionListResponseBodyDataContactCommonContacts {
	s.ContactName = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataContactCommonContacts) SetEmailConfirmed(v bool) *ReadUserSubscriptionListResponseBodyDataContactCommonContacts {
	s.EmailConfirmed = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataContactCommonContacts) SetMessageSource(v *ReadUserSubscriptionListResponseBodyDataContactCommonContactsMessageSource) *ReadUserSubscriptionListResponseBodyDataContactCommonContacts {
	s.MessageSource = v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataContactCommonContacts) SetMobileConfirmed(v bool) *ReadUserSubscriptionListResponseBodyDataContactCommonContacts {
	s.MobileConfirmed = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataContactCommonContacts) SetPosition(v string) *ReadUserSubscriptionListResponseBodyDataContactCommonContacts {
	s.Position = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataContactCommonContacts) Validate() error {
	if s.MessageSource != nil {
		if err := s.MessageSource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReadUserSubscriptionListResponseBodyDataContactCommonContactsMessageSource struct {
	// The blacklist.
	KeywordBlacklist []*string `json:"KeywordBlacklist,omitempty" xml:"KeywordBlacklist,omitempty" type:"Repeated"`
	// The whitelist.
	KeywordWhitelist []*string `json:"KeywordWhitelist,omitempty" xml:"KeywordWhitelist,omitempty" type:"Repeated"`
}

func (s ReadUserSubscriptionListResponseBodyDataContactCommonContactsMessageSource) String() string {
	return dara.Prettify(s)
}

func (s ReadUserSubscriptionListResponseBodyDataContactCommonContactsMessageSource) GoString() string {
	return s.String()
}

func (s *ReadUserSubscriptionListResponseBodyDataContactCommonContactsMessageSource) GetKeywordBlacklist() []*string {
	return s.KeywordBlacklist
}

func (s *ReadUserSubscriptionListResponseBodyDataContactCommonContactsMessageSource) GetKeywordWhitelist() []*string {
	return s.KeywordWhitelist
}

func (s *ReadUserSubscriptionListResponseBodyDataContactCommonContactsMessageSource) SetKeywordBlacklist(v []*string) *ReadUserSubscriptionListResponseBodyDataContactCommonContactsMessageSource {
	s.KeywordBlacklist = v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataContactCommonContactsMessageSource) SetKeywordWhitelist(v []*string) *ReadUserSubscriptionListResponseBodyDataContactCommonContactsMessageSource {
	s.KeywordWhitelist = v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataContactCommonContactsMessageSource) Validate() error {
	return dara.Validate(s)
}

type ReadUserSubscriptionListResponseBodyDataContactWebhookContacts struct {
	// The contact ID.
	//
	// example:
	//
	// 0
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The name of the Account Center contact.
	//
	// example:
	//
	// test
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// The message source.
	MessageSource *ReadUserSubscriptionListResponseBodyDataContactWebhookContactsMessageSource `json:"MessageSource,omitempty" xml:"MessageSource,omitempty" type:"Struct"`
	// The security token.
	//
	// example:
	//
	// /
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The webhook URL.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=xxxxxxxx
	ServerUrl *string `json:"ServerUrl,omitempty" xml:"ServerUrl,omitempty"`
	// The webhook type.
	//
	// example:
	//
	// dingtalk
	WebhookType *string `json:"WebhookType,omitempty" xml:"WebhookType,omitempty"`
}

func (s ReadUserSubscriptionListResponseBodyDataContactWebhookContacts) String() string {
	return dara.Prettify(s)
}

func (s ReadUserSubscriptionListResponseBodyDataContactWebhookContacts) GoString() string {
	return s.String()
}

func (s *ReadUserSubscriptionListResponseBodyDataContactWebhookContacts) GetContactId() *int64 {
	return s.ContactId
}

func (s *ReadUserSubscriptionListResponseBodyDataContactWebhookContacts) GetContactName() *string {
	return s.ContactName
}

func (s *ReadUserSubscriptionListResponseBodyDataContactWebhookContacts) GetMessageSource() *ReadUserSubscriptionListResponseBodyDataContactWebhookContactsMessageSource {
	return s.MessageSource
}

func (s *ReadUserSubscriptionListResponseBodyDataContactWebhookContacts) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ReadUserSubscriptionListResponseBodyDataContactWebhookContacts) GetServerUrl() *string {
	return s.ServerUrl
}

func (s *ReadUserSubscriptionListResponseBodyDataContactWebhookContacts) GetWebhookType() *string {
	return s.WebhookType
}

func (s *ReadUserSubscriptionListResponseBodyDataContactWebhookContacts) SetContactId(v int64) *ReadUserSubscriptionListResponseBodyDataContactWebhookContacts {
	s.ContactId = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataContactWebhookContacts) SetContactName(v string) *ReadUserSubscriptionListResponseBodyDataContactWebhookContacts {
	s.ContactName = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataContactWebhookContacts) SetMessageSource(v *ReadUserSubscriptionListResponseBodyDataContactWebhookContactsMessageSource) *ReadUserSubscriptionListResponseBodyDataContactWebhookContacts {
	s.MessageSource = v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataContactWebhookContacts) SetSecurityToken(v string) *ReadUserSubscriptionListResponseBodyDataContactWebhookContacts {
	s.SecurityToken = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataContactWebhookContacts) SetServerUrl(v string) *ReadUserSubscriptionListResponseBodyDataContactWebhookContacts {
	s.ServerUrl = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataContactWebhookContacts) SetWebhookType(v string) *ReadUserSubscriptionListResponseBodyDataContactWebhookContacts {
	s.WebhookType = &v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataContactWebhookContacts) Validate() error {
	if s.MessageSource != nil {
		if err := s.MessageSource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReadUserSubscriptionListResponseBodyDataContactWebhookContactsMessageSource struct {
	// The blacklist.
	KeywordBlacklist []*string `json:"KeywordBlacklist,omitempty" xml:"KeywordBlacklist,omitempty" type:"Repeated"`
	// The whitelist.
	KeywordWhitelist []*string `json:"KeywordWhitelist,omitempty" xml:"KeywordWhitelist,omitempty" type:"Repeated"`
}

func (s ReadUserSubscriptionListResponseBodyDataContactWebhookContactsMessageSource) String() string {
	return dara.Prettify(s)
}

func (s ReadUserSubscriptionListResponseBodyDataContactWebhookContactsMessageSource) GoString() string {
	return s.String()
}

func (s *ReadUserSubscriptionListResponseBodyDataContactWebhookContactsMessageSource) GetKeywordBlacklist() []*string {
	return s.KeywordBlacklist
}

func (s *ReadUserSubscriptionListResponseBodyDataContactWebhookContactsMessageSource) GetKeywordWhitelist() []*string {
	return s.KeywordWhitelist
}

func (s *ReadUserSubscriptionListResponseBodyDataContactWebhookContactsMessageSource) SetKeywordBlacklist(v []*string) *ReadUserSubscriptionListResponseBodyDataContactWebhookContactsMessageSource {
	s.KeywordBlacklist = v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataContactWebhookContactsMessageSource) SetKeywordWhitelist(v []*string) *ReadUserSubscriptionListResponseBodyDataContactWebhookContactsMessageSource {
	s.KeywordWhitelist = v
	return s
}

func (s *ReadUserSubscriptionListResponseBodyDataContactWebhookContactsMessageSource) Validate() error {
	return dara.Validate(s)
}
