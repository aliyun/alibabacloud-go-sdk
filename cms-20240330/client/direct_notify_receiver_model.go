// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDirectNotifyReceiver interface {
	dara.Model
	String() string
	GoString() string
	SetChannels(v []*string) *DirectNotifyReceiver
	GetChannels() []*string
	SetIdentifiers(v []*string) *DirectNotifyReceiver
	GetIdentifiers() []*string
	SetTargetType(v string) *DirectNotifyReceiver
	GetTargetType() *string
}

type DirectNotifyReceiver struct {
	// The list of notification channels. This parameter is valid only for person-based types (CONTACT/GROUP/DUTY). Valid values: SMS, CALL, EMAIL.
	Channels []*string `json:"channels,omitempty" xml:"channels,omitempty" type:"Repeated"`
	// The list of Notification Recipient identifiers. For person-based types, the identifiers are contacts, contact groups, or on-call schedule identifiers. For IM-based types, the identifiers are webhook identifiers.
	Identifiers []*string `json:"identifiers,omitempty" xml:"identifiers,omitempty" type:"Repeated"`
	// The Notification Recipient type. Person-object types (CONTACT/GROUP/DUTY) require channels to specify notification methods. IM-object types (DINGTALK/FEISHU/SLACK/WEIXIN/WEBHOOK) do not require channels.
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s DirectNotifyReceiver) String() string {
	return dara.Prettify(s)
}

func (s DirectNotifyReceiver) GoString() string {
	return s.String()
}

func (s *DirectNotifyReceiver) GetChannels() []*string {
	return s.Channels
}

func (s *DirectNotifyReceiver) GetIdentifiers() []*string {
	return s.Identifiers
}

func (s *DirectNotifyReceiver) GetTargetType() *string {
	return s.TargetType
}

func (s *DirectNotifyReceiver) SetChannels(v []*string) *DirectNotifyReceiver {
	s.Channels = v
	return s
}

func (s *DirectNotifyReceiver) SetIdentifiers(v []*string) *DirectNotifyReceiver {
	s.Identifiers = v
	return s
}

func (s *DirectNotifyReceiver) SetTargetType(v string) *DirectNotifyReceiver {
	s.TargetType = &v
	return s
}

func (s *DirectNotifyReceiver) Validate() error {
	return dara.Validate(s)
}
