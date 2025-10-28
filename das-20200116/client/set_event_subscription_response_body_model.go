// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetEventSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SetEventSubscriptionResponseBody
	GetCode() *string
	SetData(v *SetEventSubscriptionResponseBodyData) *SetEventSubscriptionResponseBody
	GetData() *SetEventSubscriptionResponseBodyData
	SetMessage(v string) *SetEventSubscriptionResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetEventSubscriptionResponseBody
	GetRequestId() *string
	SetSuccess(v string) *SetEventSubscriptionResponseBody
	GetSuccess() *string
}

type SetEventSubscriptionResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information.
	Data *SetEventSubscriptionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 097F0C56-B252-515A-B602-FC56EF93EF8A
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

func (s SetEventSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetEventSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *SetEventSubscriptionResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetEventSubscriptionResponseBody) GetData() *SetEventSubscriptionResponseBodyData {
	return s.Data
}

func (s *SetEventSubscriptionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetEventSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetEventSubscriptionResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *SetEventSubscriptionResponseBody) SetCode(v string) *SetEventSubscriptionResponseBody {
	s.Code = &v
	return s
}

func (s *SetEventSubscriptionResponseBody) SetData(v *SetEventSubscriptionResponseBodyData) *SetEventSubscriptionResponseBody {
	s.Data = v
	return s
}

func (s *SetEventSubscriptionResponseBody) SetMessage(v string) *SetEventSubscriptionResponseBody {
	s.Message = &v
	return s
}

func (s *SetEventSubscriptionResponseBody) SetRequestId(v string) *SetEventSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetEventSubscriptionResponseBody) SetSuccess(v string) *SetEventSubscriptionResponseBody {
	s.Success = &v
	return s
}

func (s *SetEventSubscriptionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetEventSubscriptionResponseBodyData struct {
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
	// The name of the contact who receives alert notifications. Multiple names are separated by commas (,).
	//
	// example:
	//
	// Default contact
	ContactName *string `json:"contactName,omitempty" xml:"contactName,omitempty"`
	// The supported event scenarios. Only **AllContext*	- is returned for this parameter, which indicates that all scenarios are supported.
	//
	// example:
	//
	// AllContext
	EventContext *string `json:"eventContext,omitempty" xml:"eventContext,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The language of event notifications. Only **zh-CN*	- is returned for this parameter, which indicates that event notifications are sent in Chinese.
	//
	// example:
	//
	// zh_CN
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
	// The risk level of the events. Valid values:
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
	// The minimum interval between consecutive event notifications. Unit: seconds.
	//
	// example:
	//
	// 60
	MinInterval *int32 `json:"minInterval,omitempty" xml:"minInterval,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1088760496****
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SetEventSubscriptionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SetEventSubscriptionResponseBodyData) GoString() string {
	return s.String()
}

func (s *SetEventSubscriptionResponseBodyData) GetActive() *int32 {
	return s.Active
}

func (s *SetEventSubscriptionResponseBodyData) GetChannelType() *string {
	return s.ChannelType
}

func (s *SetEventSubscriptionResponseBodyData) GetContactGroupName() *string {
	return s.ContactGroupName
}

func (s *SetEventSubscriptionResponseBodyData) GetContactName() *string {
	return s.ContactName
}

func (s *SetEventSubscriptionResponseBodyData) GetEventContext() *string {
	return s.EventContext
}

func (s *SetEventSubscriptionResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetEventSubscriptionResponseBodyData) GetLang() *string {
	return s.Lang
}

func (s *SetEventSubscriptionResponseBodyData) GetLevel() *string {
	return s.Level
}

func (s *SetEventSubscriptionResponseBodyData) GetMinInterval() *int32 {
	return s.MinInterval
}

func (s *SetEventSubscriptionResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *SetEventSubscriptionResponseBodyData) SetActive(v int32) *SetEventSubscriptionResponseBodyData {
	s.Active = &v
	return s
}

func (s *SetEventSubscriptionResponseBodyData) SetChannelType(v string) *SetEventSubscriptionResponseBodyData {
	s.ChannelType = &v
	return s
}

func (s *SetEventSubscriptionResponseBodyData) SetContactGroupName(v string) *SetEventSubscriptionResponseBodyData {
	s.ContactGroupName = &v
	return s
}

func (s *SetEventSubscriptionResponseBodyData) SetContactName(v string) *SetEventSubscriptionResponseBodyData {
	s.ContactName = &v
	return s
}

func (s *SetEventSubscriptionResponseBodyData) SetEventContext(v string) *SetEventSubscriptionResponseBodyData {
	s.EventContext = &v
	return s
}

func (s *SetEventSubscriptionResponseBodyData) SetInstanceId(v string) *SetEventSubscriptionResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *SetEventSubscriptionResponseBodyData) SetLang(v string) *SetEventSubscriptionResponseBodyData {
	s.Lang = &v
	return s
}

func (s *SetEventSubscriptionResponseBodyData) SetLevel(v string) *SetEventSubscriptionResponseBodyData {
	s.Level = &v
	return s
}

func (s *SetEventSubscriptionResponseBodyData) SetMinInterval(v int32) *SetEventSubscriptionResponseBodyData {
	s.MinInterval = &v
	return s
}

func (s *SetEventSubscriptionResponseBodyData) SetUserId(v string) *SetEventSubscriptionResponseBodyData {
	s.UserId = &v
	return s
}

func (s *SetEventSubscriptionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
