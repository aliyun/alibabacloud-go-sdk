// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotifyChannel interface {
	dara.Model
	String() string
	GoString() string
	SetChannelType(v string) *NotifyChannel
	GetChannelType() *string
	SetEnabledSubChannels(v []*string) *NotifyChannel
	GetEnabledSubChannels() []*string
	SetReceivers(v []*string) *NotifyChannel
	GetReceivers() []*string
}

type NotifyChannel struct {
	ChannelType        *string   `json:"channelType,omitempty" xml:"channelType,omitempty"`
	EnabledSubChannels []*string `json:"enabledSubChannels,omitempty" xml:"enabledSubChannels,omitempty" type:"Repeated"`
	Receivers          []*string `json:"receivers,omitempty" xml:"receivers,omitempty" type:"Repeated"`
}

func (s NotifyChannel) String() string {
	return dara.Prettify(s)
}

func (s NotifyChannel) GoString() string {
	return s.String()
}

func (s *NotifyChannel) GetChannelType() *string {
	return s.ChannelType
}

func (s *NotifyChannel) GetEnabledSubChannels() []*string {
	return s.EnabledSubChannels
}

func (s *NotifyChannel) GetReceivers() []*string {
	return s.Receivers
}

func (s *NotifyChannel) SetChannelType(v string) *NotifyChannel {
	s.ChannelType = &v
	return s
}

func (s *NotifyChannel) SetEnabledSubChannels(v []*string) *NotifyChannel {
	s.EnabledSubChannels = v
	return s
}

func (s *NotifyChannel) SetReceivers(v []*string) *NotifyChannel {
	s.Receivers = v
	return s
}

func (s *NotifyChannel) Validate() error {
	return dara.Validate(s)
}
