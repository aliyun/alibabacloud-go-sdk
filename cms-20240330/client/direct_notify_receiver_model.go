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
	// The list of notification channels. This parameter is valid only for person types (CONTACT/GROUP/DUTY). Valid values: SMS, CALL, EMAIL.
	Channels []*string `json:"channels,omitempty" xml:"channels,omitempty" type:"Repeated"`
	// The list of notification recipient identifiers. For person types, this is the identifier of the contact, contact group, or on-call schedule. For IM types, this is the webhook identifier.
	Identifiers []*string `json:"identifiers,omitempty" xml:"identifiers,omitempty" type:"Repeated"`
	// The Notification Recipient type. For person types (CONTACT/GROUP/DUTY), you must specify notification channels by using the channels parameter. For IM types (DINGTALK/FEISHU/SLACK/WEIXIN/WEBHOOK), the channels parameter is not required.
	//
	// example:
	//
	// CONTACT
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
