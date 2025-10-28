// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetEventSubscriptionResponseBody
	GetCode() *string
	SetData(v *GetEventSubscriptionResponseBodyData) *GetEventSubscriptionResponseBody
	GetData() *GetEventSubscriptionResponseBodyData
	SetMessage(v string) *GetEventSubscriptionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetEventSubscriptionResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetEventSubscriptionResponseBody
	GetSuccess() *string
}

type GetEventSubscriptionResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetEventSubscriptionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEventSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEventSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *GetEventSubscriptionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetEventSubscriptionResponseBody) GetData() *GetEventSubscriptionResponseBodyData {
	return s.Data
}

func (s *GetEventSubscriptionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetEventSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEventSubscriptionResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetEventSubscriptionResponseBody) SetCode(v string) *GetEventSubscriptionResponseBody {
	s.Code = &v
	return s
}

func (s *GetEventSubscriptionResponseBody) SetData(v *GetEventSubscriptionResponseBodyData) *GetEventSubscriptionResponseBody {
	s.Data = v
	return s
}

func (s *GetEventSubscriptionResponseBody) SetMessage(v string) *GetEventSubscriptionResponseBody {
	s.Message = &v
	return s
}

func (s *GetEventSubscriptionResponseBody) SetRequestId(v string) *GetEventSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEventSubscriptionResponseBody) SetSuccess(v string) *GetEventSubscriptionResponseBody {
	s.Success = &v
	return s
}

func (s *GetEventSubscriptionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEventSubscriptionResponseBodyData struct {
	// Indicates whether the event subscription feature is enabled. Valid values:
	//
	// 	- **0**: The event subscription feature is disabled.
	//
	// 	- **1**: The event subscription feature is enabled.
	//
	// example:
	//
	// 1
	Active *int32 `json:"active,omitempty" xml:"active,omitempty"`
	// The notification method. Valid values:
	//
	// 	- **hdm_alarm_sms**: text message.
	//
	// 	- **dingtalk**: DingTalk chatbot.
	//
	// 	- **hdm_alarm_sms_and_email**: text message and email.
	//
	// 	- **hdm_alarm_sms,dingtalk**: text message and DingTalk chatbot.
	//
	// example:
	//
	// hdm_alarm_sms,dingtalk
	ChannelType *string `json:"channelType,omitempty" xml:"channelType,omitempty"`
	// The name of the contact group that receives alert notifications. Multiple names are separated by commas (,).
	//
	// example:
	//
	// Default contact group
	ContactGroupName *string `json:"contactGroupName,omitempty" xml:"contactGroupName,omitempty"`
	// The alert contact groups.
	ContactGroups []*GetEventSubscriptionResponseBodyDataContactGroups `json:"contactGroups,omitempty" xml:"contactGroups,omitempty" type:"Repeated"`
	// The name of the subscriber who receives alert notifications. Multiple names are separated by commas (,).
	//
	// example:
	//
	// Default contact
	ContactName *string `json:"contactName,omitempty" xml:"contactName,omitempty"`
	// The user ID.
	Contacts []*GetEventSubscriptionResponseBodyDataContacts `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	// The supported event scenarios. Only **AllContext*	- may be returned, which indicates that all scenarios are supported.
	//
	// example:
	//
	// AllContext
	EventContext *string `json:"eventContext,omitempty" xml:"eventContext,omitempty"`
	// The supported event scenarios in which event subscription can be sent.
	EventSendGroup []*string `json:"eventSendGroup,omitempty" xml:"eventSendGroup,omitempty" type:"Repeated"`
	// The time when event subscription was enabled. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1633071840000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The time when the event subscription settings were most recently modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1633071850000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The primary key ID of the database.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The language of event notifications. Only **zh-CN*	- may be returned, which indicates that event notifications are sent in Chinese.
	//
	// example:
	//
	// zh_CN
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
	// The risk level of the events that trigger notifications. Valid values:
	//
	// 	- **Notice**
	//
	// 	- **Optimization**
	//
	// 	- **Warn**
	//
	// 	- **Critical**
	//
	// example:
	//
	// Optimization
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// The minimum interval between event notifications. Unit: seconds.
	//
	// example:
	//
	// 60
	MinInterval *string `json:"minInterval,omitempty" xml:"minInterval,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1088760496****
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetEventSubscriptionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetEventSubscriptionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEventSubscriptionResponseBodyData) GetActive() *int32 {
	return s.Active
}

func (s *GetEventSubscriptionResponseBodyData) GetChannelType() *string {
	return s.ChannelType
}

func (s *GetEventSubscriptionResponseBodyData) GetContactGroupName() *string {
	return s.ContactGroupName
}

func (s *GetEventSubscriptionResponseBodyData) GetContactGroups() []*GetEventSubscriptionResponseBodyDataContactGroups {
	return s.ContactGroups
}

func (s *GetEventSubscriptionResponseBodyData) GetContactName() *string {
	return s.ContactName
}

func (s *GetEventSubscriptionResponseBodyData) GetContacts() []*GetEventSubscriptionResponseBodyDataContacts {
	return s.Contacts
}

func (s *GetEventSubscriptionResponseBodyData) GetEventContext() *string {
	return s.EventContext
}

func (s *GetEventSubscriptionResponseBodyData) GetEventSendGroup() []*string {
	return s.EventSendGroup
}

func (s *GetEventSubscriptionResponseBodyData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GetEventSubscriptionResponseBodyData) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GetEventSubscriptionResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetEventSubscriptionResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetEventSubscriptionResponseBodyData) GetLang() *string {
	return s.Lang
}

func (s *GetEventSubscriptionResponseBodyData) GetLevel() *string {
	return s.Level
}

func (s *GetEventSubscriptionResponseBodyData) GetMinInterval() *string {
	return s.MinInterval
}

func (s *GetEventSubscriptionResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *GetEventSubscriptionResponseBodyData) SetActive(v int32) *GetEventSubscriptionResponseBodyData {
	s.Active = &v
	return s
}

func (s *GetEventSubscriptionResponseBodyData) SetChannelType(v string) *GetEventSubscriptionResponseBodyData {
	s.ChannelType = &v
	return s
}

func (s *GetEventSubscriptionResponseBodyData) SetContactGroupName(v string) *GetEventSubscriptionResponseBodyData {
	s.ContactGroupName = &v
	return s
}

func (s *GetEventSubscriptionResponseBodyData) SetContactGroups(v []*GetEventSubscriptionResponseBodyDataContactGroups) *GetEventSubscriptionResponseBodyData {
	s.ContactGroups = v
	return s
}

func (s *GetEventSubscriptionResponseBodyData) SetContactName(v string) *GetEventSubscriptionResponseBodyData {
	s.ContactName = &v
	return s
}

func (s *GetEventSubscriptionResponseBodyData) SetContacts(v []*GetEventSubscriptionResponseBodyDataContacts) *GetEventSubscriptionResponseBodyData {
	s.Contacts = v
	return s
}

func (s *GetEventSubscriptionResponseBodyData) SetEventContext(v string) *GetEventSubscriptionResponseBodyData {
	s.EventContext = &v
	return s
}

func (s *GetEventSubscriptionResponseBodyData) SetEventSendGroup(v []*string) *GetEventSubscriptionResponseBodyData {
	s.EventSendGroup = v
	return s
}

func (s *GetEventSubscriptionResponseBodyData) SetGmtCreate(v int64) *GetEventSubscriptionResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetEventSubscriptionResponseBodyData) SetGmtModified(v int64) *GetEventSubscriptionResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetEventSubscriptionResponseBodyData) SetId(v int64) *GetEventSubscriptionResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetEventSubscriptionResponseBodyData) SetInstanceId(v string) *GetEventSubscriptionResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetEventSubscriptionResponseBodyData) SetLang(v string) *GetEventSubscriptionResponseBodyData {
	s.Lang = &v
	return s
}

func (s *GetEventSubscriptionResponseBodyData) SetLevel(v string) *GetEventSubscriptionResponseBodyData {
	s.Level = &v
	return s
}

func (s *GetEventSubscriptionResponseBodyData) SetMinInterval(v string) *GetEventSubscriptionResponseBodyData {
	s.MinInterval = &v
	return s
}

func (s *GetEventSubscriptionResponseBodyData) SetUserId(v string) *GetEventSubscriptionResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetEventSubscriptionResponseBodyData) Validate() error {
	if s.ContactGroups != nil {
		for _, item := range s.ContactGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Contacts != nil {
		for _, item := range s.Contacts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetEventSubscriptionResponseBodyDataContactGroups struct {
	// The members of the alert contact group.
	//
	// example:
	//
	// "[\\"Mr. Zhang\\",\\"Ms. Wang\\",\\"Mr. Li\\"]"
	Contacts *string `json:"contacts,omitempty" xml:"contacts,omitempty"`
	// The description of the alert contact group.
	//
	// example:
	//
	// Default contact
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the alert contact group.
	//
	// example:
	//
	// Mr. Zhang
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1088760496****
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetEventSubscriptionResponseBodyDataContactGroups) String() string {
	return dara.Prettify(s)
}

func (s GetEventSubscriptionResponseBodyDataContactGroups) GoString() string {
	return s.String()
}

func (s *GetEventSubscriptionResponseBodyDataContactGroups) GetContacts() *string {
	return s.Contacts
}

func (s *GetEventSubscriptionResponseBodyDataContactGroups) GetDescription() *string {
	return s.Description
}

func (s *GetEventSubscriptionResponseBodyDataContactGroups) GetName() *string {
	return s.Name
}

func (s *GetEventSubscriptionResponseBodyDataContactGroups) GetUserId() *string {
	return s.UserId
}

func (s *GetEventSubscriptionResponseBodyDataContactGroups) SetContacts(v string) *GetEventSubscriptionResponseBodyDataContactGroups {
	s.Contacts = &v
	return s
}

func (s *GetEventSubscriptionResponseBodyDataContactGroups) SetDescription(v string) *GetEventSubscriptionResponseBodyDataContactGroups {
	s.Description = &v
	return s
}

func (s *GetEventSubscriptionResponseBodyDataContactGroups) SetName(v string) *GetEventSubscriptionResponseBodyDataContactGroups {
	s.Name = &v
	return s
}

func (s *GetEventSubscriptionResponseBodyDataContactGroups) SetUserId(v string) *GetEventSubscriptionResponseBodyDataContactGroups {
	s.UserId = &v
	return s
}

func (s *GetEventSubscriptionResponseBodyDataContactGroups) Validate() error {
	return dara.Validate(s)
}

type GetEventSubscriptionResponseBodyDataContacts struct {
	// The webhook URL of the DingTalk chatbot.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=68fa29a9eaf3ba9994f54fxxxc1aa9879700308f90e9c23ebfb3663642c9****
	DingtalkHook *string `json:"dingtalkHook,omitempty" xml:"dingtalkHook,omitempty"`
	// The email address of the alert contact.
	//
	// example:
	//
	// a***@example.net
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// The contact groups to which the alert contact belongs.
	Groups []*string `json:"groups,omitempty" xml:"groups,omitempty" type:"Repeated"`
	// Indicates whether the alert contact name is the same as the contact name on CloudMonitor.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsCmsReduplicated *bool `json:"isCmsReduplicated,omitempty" xml:"isCmsReduplicated,omitempty"`
	// The name of the alert contact.
	//
	// example:
	//
	// Mr. Zhang
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The mobile number of the alert contact.
	//
	// example:
	//
	// 1390000****
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1088760496****
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetEventSubscriptionResponseBodyDataContacts) String() string {
	return dara.Prettify(s)
}

func (s GetEventSubscriptionResponseBodyDataContacts) GoString() string {
	return s.String()
}

func (s *GetEventSubscriptionResponseBodyDataContacts) GetDingtalkHook() *string {
	return s.DingtalkHook
}

func (s *GetEventSubscriptionResponseBodyDataContacts) GetEmail() *string {
	return s.Email
}

func (s *GetEventSubscriptionResponseBodyDataContacts) GetGroups() []*string {
	return s.Groups
}

func (s *GetEventSubscriptionResponseBodyDataContacts) GetIsCmsReduplicated() *bool {
	return s.IsCmsReduplicated
}

func (s *GetEventSubscriptionResponseBodyDataContacts) GetName() *string {
	return s.Name
}

func (s *GetEventSubscriptionResponseBodyDataContacts) GetPhone() *string {
	return s.Phone
}

func (s *GetEventSubscriptionResponseBodyDataContacts) GetUserId() *string {
	return s.UserId
}

func (s *GetEventSubscriptionResponseBodyDataContacts) SetDingtalkHook(v string) *GetEventSubscriptionResponseBodyDataContacts {
	s.DingtalkHook = &v
	return s
}

func (s *GetEventSubscriptionResponseBodyDataContacts) SetEmail(v string) *GetEventSubscriptionResponseBodyDataContacts {
	s.Email = &v
	return s
}

func (s *GetEventSubscriptionResponseBodyDataContacts) SetGroups(v []*string) *GetEventSubscriptionResponseBodyDataContacts {
	s.Groups = v
	return s
}

func (s *GetEventSubscriptionResponseBodyDataContacts) SetIsCmsReduplicated(v bool) *GetEventSubscriptionResponseBodyDataContacts {
	s.IsCmsReduplicated = &v
	return s
}

func (s *GetEventSubscriptionResponseBodyDataContacts) SetName(v string) *GetEventSubscriptionResponseBodyDataContacts {
	s.Name = &v
	return s
}

func (s *GetEventSubscriptionResponseBodyDataContacts) SetPhone(v string) *GetEventSubscriptionResponseBodyDataContacts {
	s.Phone = &v
	return s
}

func (s *GetEventSubscriptionResponseBodyDataContacts) SetUserId(v string) *GetEventSubscriptionResponseBodyDataContacts {
	s.UserId = &v
	return s
}

func (s *GetEventSubscriptionResponseBodyDataContacts) Validate() error {
	return dara.Validate(s)
}
