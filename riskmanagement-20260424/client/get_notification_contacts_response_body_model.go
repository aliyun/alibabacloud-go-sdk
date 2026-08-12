// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNotificationContactsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetNotificationContactsResponseBody
	GetCode() *string
	SetData(v []*GetNotificationContactsResponseBodyData) *GetNotificationContactsResponseBody
	GetData() []*GetNotificationContactsResponseBodyData
	SetMessage(v string) *GetNotificationContactsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetNotificationContactsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetNotificationContactsResponseBody
	GetSuccess() *bool
}

type GetNotificationContactsResponseBody struct {
	// The status code.
	//
	// - **200**: Success.
	//
	// - **Other (400, 500)**: Failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The query result.
	Data []*GetNotificationContactsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The prompt message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2FBDD713-00A5-5C98-B661-3FD31A349B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful.
	//
	// - **true**: Success.
	//
	// - **false**: Failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNotificationContactsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNotificationContactsResponseBody) GoString() string {
	return s.String()
}

func (s *GetNotificationContactsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetNotificationContactsResponseBody) GetData() []*GetNotificationContactsResponseBodyData {
	return s.Data
}

func (s *GetNotificationContactsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetNotificationContactsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNotificationContactsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetNotificationContactsResponseBody) SetCode(v string) *GetNotificationContactsResponseBody {
	s.Code = &v
	return s
}

func (s *GetNotificationContactsResponseBody) SetData(v []*GetNotificationContactsResponseBodyData) *GetNotificationContactsResponseBody {
	s.Data = v
	return s
}

func (s *GetNotificationContactsResponseBody) SetMessage(v string) *GetNotificationContactsResponseBody {
	s.Message = &v
	return s
}

func (s *GetNotificationContactsResponseBody) SetRequestId(v string) *GetNotificationContactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNotificationContactsResponseBody) SetSuccess(v bool) *GetNotificationContactsResponseBody {
	s.Success = &v
	return s
}

func (s *GetNotificationContactsResponseBody) Validate() error {
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

type GetNotificationContactsResponseBodyData struct {
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 1355290655619147
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The message category code.
	//
	// example:
	//
	// prod_edu_content
	CategoryCode *string `json:"CategoryCode,omitempty" xml:"CategoryCode,omitempty"`
	// The message category description.
	//
	// example:
	//
	// Product usage scenarios and technical sharing content.
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
	// Product messages.
	CategoryGroupName *string `json:"CategoryGroupName,omitempty" xml:"CategoryGroupName,omitempty"`
	// The message category name.
	//
	// example:
	//
	// Product educational content.
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The channel list.
	ChannelConfigs []*GetNotificationContactsResponseBodyDataChannelConfigs `json:"ChannelConfigs,omitempty" xml:"ChannelConfigs,omitempty" type:"Repeated"`
	// Indicates whether all notification methods are selected.
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// true
	ChooseAllChannel *bool `json:"ChooseAllChannel,omitempty" xml:"ChooseAllChannel,omitempty"`
	// The general contact list.
	ContactInfoList []*GetNotificationContactsResponseBodyDataContactInfoList `json:"ContactInfoList,omitempty" xml:"ContactInfoList,omitempty" type:"Repeated"`
}

func (s GetNotificationContactsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetNotificationContactsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNotificationContactsResponseBodyData) GetAliUid() *int64 {
	return s.AliUid
}

func (s *GetNotificationContactsResponseBodyData) GetCategoryCode() *string {
	return s.CategoryCode
}

func (s *GetNotificationContactsResponseBodyData) GetCategoryDesc() *string {
	return s.CategoryDesc
}

func (s *GetNotificationContactsResponseBodyData) GetCategoryGroupCode() *string {
	return s.CategoryGroupCode
}

func (s *GetNotificationContactsResponseBodyData) GetCategoryGroupName() *string {
	return s.CategoryGroupName
}

func (s *GetNotificationContactsResponseBodyData) GetCategoryName() *string {
	return s.CategoryName
}

func (s *GetNotificationContactsResponseBodyData) GetChannelConfigs() []*GetNotificationContactsResponseBodyDataChannelConfigs {
	return s.ChannelConfigs
}

func (s *GetNotificationContactsResponseBodyData) GetChooseAllChannel() *bool {
	return s.ChooseAllChannel
}

func (s *GetNotificationContactsResponseBodyData) GetContactInfoList() []*GetNotificationContactsResponseBodyDataContactInfoList {
	return s.ContactInfoList
}

func (s *GetNotificationContactsResponseBodyData) SetAliUid(v int64) *GetNotificationContactsResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *GetNotificationContactsResponseBodyData) SetCategoryCode(v string) *GetNotificationContactsResponseBodyData {
	s.CategoryCode = &v
	return s
}

func (s *GetNotificationContactsResponseBodyData) SetCategoryDesc(v string) *GetNotificationContactsResponseBodyData {
	s.CategoryDesc = &v
	return s
}

func (s *GetNotificationContactsResponseBodyData) SetCategoryGroupCode(v string) *GetNotificationContactsResponseBodyData {
	s.CategoryGroupCode = &v
	return s
}

func (s *GetNotificationContactsResponseBodyData) SetCategoryGroupName(v string) *GetNotificationContactsResponseBodyData {
	s.CategoryGroupName = &v
	return s
}

func (s *GetNotificationContactsResponseBodyData) SetCategoryName(v string) *GetNotificationContactsResponseBodyData {
	s.CategoryName = &v
	return s
}

func (s *GetNotificationContactsResponseBodyData) SetChannelConfigs(v []*GetNotificationContactsResponseBodyDataChannelConfigs) *GetNotificationContactsResponseBodyData {
	s.ChannelConfigs = v
	return s
}

func (s *GetNotificationContactsResponseBodyData) SetChooseAllChannel(v bool) *GetNotificationContactsResponseBodyData {
	s.ChooseAllChannel = &v
	return s
}

func (s *GetNotificationContactsResponseBodyData) SetContactInfoList(v []*GetNotificationContactsResponseBodyDataContactInfoList) *GetNotificationContactsResponseBodyData {
	s.ContactInfoList = v
	return s
}

func (s *GetNotificationContactsResponseBodyData) Validate() error {
	if s.ChannelConfigs != nil {
		for _, item := range s.ChannelConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ContactInfoList != nil {
		for _, item := range s.ContactInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetNotificationContactsResponseBodyDataChannelConfigs struct {
	// The channel type.
	//
	// example:
	//
	// email
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// Indicates whether the subscription is configured.
	//
	// - **NO**
	//
	// - **YES**
	//
	// example:
	//
	// NO
	CheckedState *string `json:"CheckedState,omitempty" xml:"CheckedState,omitempty"`
	// Indicates whether the channel is selected by default.
	//
	// - **NO**
	//
	// - **YES**
	//
	// example:
	//
	// NO
	DefaultChecked *string `json:"DefaultChecked,omitempty" xml:"DefaultChecked,omitempty"`
	// The fatigue limit.
	//
	// example:
	//
	// 7
	FatigueDayLimit *int32 `json:"FatigueDayLimit,omitempty" xml:"FatigueDayLimit,omitempty"`
	// Indicates whether the channel is modifiable.
	//
	// - **NO**
	//
	// - **YES**
	//
	// example:
	//
	// NO
	Optional *string `json:"Optional,omitempty" xml:"Optional,omitempty"`
}

func (s GetNotificationContactsResponseBodyDataChannelConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetNotificationContactsResponseBodyDataChannelConfigs) GoString() string {
	return s.String()
}

func (s *GetNotificationContactsResponseBodyDataChannelConfigs) GetChannelType() *string {
	return s.ChannelType
}

func (s *GetNotificationContactsResponseBodyDataChannelConfigs) GetCheckedState() *string {
	return s.CheckedState
}

func (s *GetNotificationContactsResponseBodyDataChannelConfigs) GetDefaultChecked() *string {
	return s.DefaultChecked
}

func (s *GetNotificationContactsResponseBodyDataChannelConfigs) GetFatigueDayLimit() *int32 {
	return s.FatigueDayLimit
}

func (s *GetNotificationContactsResponseBodyDataChannelConfigs) GetOptional() *string {
	return s.Optional
}

func (s *GetNotificationContactsResponseBodyDataChannelConfigs) SetChannelType(v string) *GetNotificationContactsResponseBodyDataChannelConfigs {
	s.ChannelType = &v
	return s
}

func (s *GetNotificationContactsResponseBodyDataChannelConfigs) SetCheckedState(v string) *GetNotificationContactsResponseBodyDataChannelConfigs {
	s.CheckedState = &v
	return s
}

func (s *GetNotificationContactsResponseBodyDataChannelConfigs) SetDefaultChecked(v string) *GetNotificationContactsResponseBodyDataChannelConfigs {
	s.DefaultChecked = &v
	return s
}

func (s *GetNotificationContactsResponseBodyDataChannelConfigs) SetFatigueDayLimit(v int32) *GetNotificationContactsResponseBodyDataChannelConfigs {
	s.FatigueDayLimit = &v
	return s
}

func (s *GetNotificationContactsResponseBodyDataChannelConfigs) SetOptional(v string) *GetNotificationContactsResponseBodyDataChannelConfigs {
	s.Optional = &v
	return s
}

func (s *GetNotificationContactsResponseBodyDataChannelConfigs) Validate() error {
	return dara.Validate(s)
}

type GetNotificationContactsResponseBodyDataContactInfoList struct {
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 1492387044070147
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// Indicates whether the contact is bound.
	//
	// - **true**
	//
	// - **fasle**
	//
	// example:
	//
	// true
	BindContact *bool `json:"BindContact,omitempty" xml:"BindContact,omitempty"`
	// The contact email address.
	//
	// example:
	//
	// t*@qq.*
	ContactEmail *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	// The Account Center contact ID. A value of 0 indicates the account contact.
	//
	// example:
	//
	// 0
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The Account Center contact mobile number (masked).
	//
	// example:
	//
	// 13580xxx136
	ContactMobile *string `json:"ContactMobile,omitempty" xml:"ContactMobile,omitempty"`
	// The Account Center contact name.
	//
	// example:
	//
	// shianyu
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// Indicates whether the email address is verified.
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// true
	EmailConfirmed *bool `json:"EmailConfirmed,omitempty" xml:"EmailConfirmed,omitempty"`
	// Indicates whether the Account Center contact mobile number is verified.
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// true
	MobileConfirmed *bool `json:"MobileConfirmed,omitempty" xml:"MobileConfirmed,omitempty"`
	// The Account Center contact position.
	//
	// example:
	//
	// CEO
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s GetNotificationContactsResponseBodyDataContactInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetNotificationContactsResponseBodyDataContactInfoList) GoString() string {
	return s.String()
}

func (s *GetNotificationContactsResponseBodyDataContactInfoList) GetAliUid() *int64 {
	return s.AliUid
}

func (s *GetNotificationContactsResponseBodyDataContactInfoList) GetBindContact() *bool {
	return s.BindContact
}

func (s *GetNotificationContactsResponseBodyDataContactInfoList) GetContactEmail() *string {
	return s.ContactEmail
}

func (s *GetNotificationContactsResponseBodyDataContactInfoList) GetContactId() *int64 {
	return s.ContactId
}

func (s *GetNotificationContactsResponseBodyDataContactInfoList) GetContactMobile() *string {
	return s.ContactMobile
}

func (s *GetNotificationContactsResponseBodyDataContactInfoList) GetContactName() *string {
	return s.ContactName
}

func (s *GetNotificationContactsResponseBodyDataContactInfoList) GetEmailConfirmed() *bool {
	return s.EmailConfirmed
}

func (s *GetNotificationContactsResponseBodyDataContactInfoList) GetMobileConfirmed() *bool {
	return s.MobileConfirmed
}

func (s *GetNotificationContactsResponseBodyDataContactInfoList) GetPosition() *string {
	return s.Position
}

func (s *GetNotificationContactsResponseBodyDataContactInfoList) SetAliUid(v int64) *GetNotificationContactsResponseBodyDataContactInfoList {
	s.AliUid = &v
	return s
}

func (s *GetNotificationContactsResponseBodyDataContactInfoList) SetBindContact(v bool) *GetNotificationContactsResponseBodyDataContactInfoList {
	s.BindContact = &v
	return s
}

func (s *GetNotificationContactsResponseBodyDataContactInfoList) SetContactEmail(v string) *GetNotificationContactsResponseBodyDataContactInfoList {
	s.ContactEmail = &v
	return s
}

func (s *GetNotificationContactsResponseBodyDataContactInfoList) SetContactId(v int64) *GetNotificationContactsResponseBodyDataContactInfoList {
	s.ContactId = &v
	return s
}

func (s *GetNotificationContactsResponseBodyDataContactInfoList) SetContactMobile(v string) *GetNotificationContactsResponseBodyDataContactInfoList {
	s.ContactMobile = &v
	return s
}

func (s *GetNotificationContactsResponseBodyDataContactInfoList) SetContactName(v string) *GetNotificationContactsResponseBodyDataContactInfoList {
	s.ContactName = &v
	return s
}

func (s *GetNotificationContactsResponseBodyDataContactInfoList) SetEmailConfirmed(v bool) *GetNotificationContactsResponseBodyDataContactInfoList {
	s.EmailConfirmed = &v
	return s
}

func (s *GetNotificationContactsResponseBodyDataContactInfoList) SetMobileConfirmed(v bool) *GetNotificationContactsResponseBodyDataContactInfoList {
	s.MobileConfirmed = &v
	return s
}

func (s *GetNotificationContactsResponseBodyDataContactInfoList) SetPosition(v string) *GetNotificationContactsResponseBodyDataContactInfoList {
	s.Position = &v
	return s
}

func (s *GetNotificationContactsResponseBodyDataContactInfoList) Validate() error {
	return dara.Validate(s)
}
