// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledTaskPushOptionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChannels(v []*GetScheduledTaskPushOptionsResponseBodyChannels) *GetScheduledTaskPushOptionsResponseBody
	GetChannels() []*GetScheduledTaskPushOptionsResponseBodyChannels
	SetCode(v string) *GetScheduledTaskPushOptionsResponseBody
	GetCode() *string
	SetEmptyHint(v string) *GetScheduledTaskPushOptionsResponseBody
	GetEmptyHint() *string
	SetMessage(v string) *GetScheduledTaskPushOptionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetScheduledTaskPushOptionsResponseBody
	GetRequestId() *string
}

type GetScheduledTaskPushOptionsResponseBody struct {
	// The list of notification channels.
	Channels []*GetScheduledTaskPushOptionsResponseBodyChannels `json:"channels,omitempty" xml:"channels,omitempty" type:"Repeated"`
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The prompt displayed when no third-party accounts are bound.
	//
	// example:
	//
	// No push channels available
	EmptyHint *string `json:"emptyHint,omitempty" xml:"emptyHint,omitempty"`
	// The prompt message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetScheduledTaskPushOptionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskPushOptionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskPushOptionsResponseBody) GetChannels() []*GetScheduledTaskPushOptionsResponseBodyChannels {
	return s.Channels
}

func (s *GetScheduledTaskPushOptionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetScheduledTaskPushOptionsResponseBody) GetEmptyHint() *string {
	return s.EmptyHint
}

func (s *GetScheduledTaskPushOptionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetScheduledTaskPushOptionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScheduledTaskPushOptionsResponseBody) SetChannels(v []*GetScheduledTaskPushOptionsResponseBodyChannels) *GetScheduledTaskPushOptionsResponseBody {
	s.Channels = v
	return s
}

func (s *GetScheduledTaskPushOptionsResponseBody) SetCode(v string) *GetScheduledTaskPushOptionsResponseBody {
	s.Code = &v
	return s
}

func (s *GetScheduledTaskPushOptionsResponseBody) SetEmptyHint(v string) *GetScheduledTaskPushOptionsResponseBody {
	s.EmptyHint = &v
	return s
}

func (s *GetScheduledTaskPushOptionsResponseBody) SetMessage(v string) *GetScheduledTaskPushOptionsResponseBody {
	s.Message = &v
	return s
}

func (s *GetScheduledTaskPushOptionsResponseBody) SetRequestId(v string) *GetScheduledTaskPushOptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScheduledTaskPushOptionsResponseBody) Validate() error {
	if s.Channels != nil {
		for _, item := range s.Channels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetScheduledTaskPushOptionsResponseBodyChannels struct {
	// The channel name.
	//
	// This parameter is required.
	//
	// example:
	//
	// DingTalk
	ChannelName *string `json:"channelName,omitempty" xml:"channelName,omitempty"`
	// The notification method. Valid values:
	//
	// - **hdm_alarm_sms**: SMS.
	//
	// - **dingtalk**: DingTalk chatbot.
	//
	// - **hdm_alarm_sms_and_email**: SMS and email.
	//
	// - **hdm_alarm_sms,dingtalk**: SMS and DingTalk chatbot.
	//
	// This parameter is required.
	//
	// example:
	//
	// DINGTALK
	ChannelType *string `json:"channelType,omitempty" xml:"channelType,omitempty"`
	// The optional IM groups bound to this channel for the collaboration group. This value is empty when querying personal tasks.
	ImGroups []*GetScheduledTaskPushOptionsResponseBodyChannelsImGroups `json:"imGroups,omitempty" xml:"imGroups,omitempty" type:"Repeated"`
	// The supported methods: HEAD, GET, POST, PUT, DELETE, PATCH, OPTIONS.
	Methods []*GetScheduledTaskPushOptionsResponseBodyChannelsMethods `json:"methods,omitempty" xml:"methods,omitempty" type:"Repeated"`
}

func (s GetScheduledTaskPushOptionsResponseBodyChannels) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskPushOptionsResponseBodyChannels) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskPushOptionsResponseBodyChannels) GetChannelName() *string {
	return s.ChannelName
}

func (s *GetScheduledTaskPushOptionsResponseBodyChannels) GetChannelType() *string {
	return s.ChannelType
}

func (s *GetScheduledTaskPushOptionsResponseBodyChannels) GetImGroups() []*GetScheduledTaskPushOptionsResponseBodyChannelsImGroups {
	return s.ImGroups
}

func (s *GetScheduledTaskPushOptionsResponseBodyChannels) GetMethods() []*GetScheduledTaskPushOptionsResponseBodyChannelsMethods {
	return s.Methods
}

func (s *GetScheduledTaskPushOptionsResponseBodyChannels) SetChannelName(v string) *GetScheduledTaskPushOptionsResponseBodyChannels {
	s.ChannelName = &v
	return s
}

func (s *GetScheduledTaskPushOptionsResponseBodyChannels) SetChannelType(v string) *GetScheduledTaskPushOptionsResponseBodyChannels {
	s.ChannelType = &v
	return s
}

func (s *GetScheduledTaskPushOptionsResponseBodyChannels) SetImGroups(v []*GetScheduledTaskPushOptionsResponseBodyChannelsImGroups) *GetScheduledTaskPushOptionsResponseBodyChannels {
	s.ImGroups = v
	return s
}

func (s *GetScheduledTaskPushOptionsResponseBodyChannels) SetMethods(v []*GetScheduledTaskPushOptionsResponseBodyChannelsMethods) *GetScheduledTaskPushOptionsResponseBodyChannels {
	s.Methods = v
	return s
}

func (s *GetScheduledTaskPushOptionsResponseBodyChannels) Validate() error {
	if s.ImGroups != nil {
		for _, item := range s.ImGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Methods != nil {
		for _, item := range s.Methods {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetScheduledTaskPushOptionsResponseBodyChannelsImGroups struct {
	// The external IM group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cidExample
	ImGroupId *string `json:"imGroupId,omitempty" xml:"imGroupId,omitempty"`
	// The external IM group name.
	//
	// example:
	//
	// Project collaboration group
	ImGroupName *string `json:"imGroupName,omitempty" xml:"imGroupName,omitempty"`
	// The binding record ID of the IM group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 101
	MappingId *int64 `json:"mappingId,omitempty" xml:"mappingId,omitempty"`
}

func (s GetScheduledTaskPushOptionsResponseBodyChannelsImGroups) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskPushOptionsResponseBodyChannelsImGroups) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskPushOptionsResponseBodyChannelsImGroups) GetImGroupId() *string {
	return s.ImGroupId
}

func (s *GetScheduledTaskPushOptionsResponseBodyChannelsImGroups) GetImGroupName() *string {
	return s.ImGroupName
}

func (s *GetScheduledTaskPushOptionsResponseBodyChannelsImGroups) GetMappingId() *int64 {
	return s.MappingId
}

func (s *GetScheduledTaskPushOptionsResponseBodyChannelsImGroups) SetImGroupId(v string) *GetScheduledTaskPushOptionsResponseBodyChannelsImGroups {
	s.ImGroupId = &v
	return s
}

func (s *GetScheduledTaskPushOptionsResponseBodyChannelsImGroups) SetImGroupName(v string) *GetScheduledTaskPushOptionsResponseBodyChannelsImGroups {
	s.ImGroupName = &v
	return s
}

func (s *GetScheduledTaskPushOptionsResponseBodyChannelsImGroups) SetMappingId(v int64) *GetScheduledTaskPushOptionsResponseBodyChannelsImGroups {
	s.MappingId = &v
	return s
}

func (s *GetScheduledTaskPushOptionsResponseBodyChannelsImGroups) Validate() error {
	return dara.Validate(s)
}

type GetScheduledTaskPushOptionsResponseBodyChannelsMethods struct {
	// The reason why the option is grayed out.
	//
	// example:
	//
	// No push channel is bound
	DisabledReason *string `json:"disabledReason,omitempty" xml:"disabledReason,omitempty"`
	// The feature switch. This parameter is optional when type is set to web_search.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The method.
	//
	// This parameter is required.
	//
	// example:
	//
	// channel_bot
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// The name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Group chatbot
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetScheduledTaskPushOptionsResponseBodyChannelsMethods) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskPushOptionsResponseBodyChannelsMethods) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskPushOptionsResponseBodyChannelsMethods) GetDisabledReason() *string {
	return s.DisabledReason
}

func (s *GetScheduledTaskPushOptionsResponseBodyChannelsMethods) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetScheduledTaskPushOptionsResponseBodyChannelsMethods) GetMethod() *string {
	return s.Method
}

func (s *GetScheduledTaskPushOptionsResponseBodyChannelsMethods) GetName() *string {
	return s.Name
}

func (s *GetScheduledTaskPushOptionsResponseBodyChannelsMethods) SetDisabledReason(v string) *GetScheduledTaskPushOptionsResponseBodyChannelsMethods {
	s.DisabledReason = &v
	return s
}

func (s *GetScheduledTaskPushOptionsResponseBodyChannelsMethods) SetEnabled(v bool) *GetScheduledTaskPushOptionsResponseBodyChannelsMethods {
	s.Enabled = &v
	return s
}

func (s *GetScheduledTaskPushOptionsResponseBodyChannelsMethods) SetMethod(v string) *GetScheduledTaskPushOptionsResponseBodyChannelsMethods {
	s.Method = &v
	return s
}

func (s *GetScheduledTaskPushOptionsResponseBodyChannelsMethods) SetName(v string) *GetScheduledTaskPushOptionsResponseBodyChannelsMethods {
	s.Name = &v
	return s
}

func (s *GetScheduledTaskPushOptionsResponseBodyChannelsMethods) Validate() error {
	return dara.Validate(s)
}
