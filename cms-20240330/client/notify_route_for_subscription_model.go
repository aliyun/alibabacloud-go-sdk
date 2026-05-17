// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotifyRouteForSubscription interface {
	dara.Model
	String() string
	GoString() string
	SetChannels(v []*NotifyRouteForSubscriptionChannels) *NotifyRouteForSubscription
	GetChannels() []*NotifyRouteForSubscriptionChannels
	SetEffectTimeRange(v *NotifyRouteForSubscriptionEffectTimeRange) *NotifyRouteForSubscription
	GetEffectTimeRange() *NotifyRouteForSubscriptionEffectTimeRange
}

type NotifyRouteForSubscription struct {
	Channels        []*NotifyRouteForSubscriptionChannels      `json:"channels,omitempty" xml:"channels,omitempty" type:"Repeated"`
	EffectTimeRange *NotifyRouteForSubscriptionEffectTimeRange `json:"effectTimeRange,omitempty" xml:"effectTimeRange,omitempty" type:"Struct"`
}

func (s NotifyRouteForSubscription) String() string {
	return dara.Prettify(s)
}

func (s NotifyRouteForSubscription) GoString() string {
	return s.String()
}

func (s *NotifyRouteForSubscription) GetChannels() []*NotifyRouteForSubscriptionChannels {
	return s.Channels
}

func (s *NotifyRouteForSubscription) GetEffectTimeRange() *NotifyRouteForSubscriptionEffectTimeRange {
	return s.EffectTimeRange
}

func (s *NotifyRouteForSubscription) SetChannels(v []*NotifyRouteForSubscriptionChannels) *NotifyRouteForSubscription {
	s.Channels = v
	return s
}

func (s *NotifyRouteForSubscription) SetEffectTimeRange(v *NotifyRouteForSubscriptionEffectTimeRange) *NotifyRouteForSubscription {
	s.EffectTimeRange = v
	return s
}

func (s *NotifyRouteForSubscription) Validate() error {
	if s.Channels != nil {
		for _, item := range s.Channels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EffectTimeRange != nil {
		if err := s.EffectTimeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type NotifyRouteForSubscriptionChannels struct {
	ChannelType        *string   `json:"channelType,omitempty" xml:"channelType,omitempty"`
	EnabledSubChannels []*string `json:"enabledSubChannels,omitempty" xml:"enabledSubChannels,omitempty" type:"Repeated"`
	Receivers          []*string `json:"receivers,omitempty" xml:"receivers,omitempty" type:"Repeated"`
}

func (s NotifyRouteForSubscriptionChannels) String() string {
	return dara.Prettify(s)
}

func (s NotifyRouteForSubscriptionChannels) GoString() string {
	return s.String()
}

func (s *NotifyRouteForSubscriptionChannels) GetChannelType() *string {
	return s.ChannelType
}

func (s *NotifyRouteForSubscriptionChannels) GetEnabledSubChannels() []*string {
	return s.EnabledSubChannels
}

func (s *NotifyRouteForSubscriptionChannels) GetReceivers() []*string {
	return s.Receivers
}

func (s *NotifyRouteForSubscriptionChannels) SetChannelType(v string) *NotifyRouteForSubscriptionChannels {
	s.ChannelType = &v
	return s
}

func (s *NotifyRouteForSubscriptionChannels) SetEnabledSubChannels(v []*string) *NotifyRouteForSubscriptionChannels {
	s.EnabledSubChannels = v
	return s
}

func (s *NotifyRouteForSubscriptionChannels) SetReceivers(v []*string) *NotifyRouteForSubscriptionChannels {
	s.Receivers = v
	return s
}

func (s *NotifyRouteForSubscriptionChannels) Validate() error {
	return dara.Validate(s)
}

type NotifyRouteForSubscriptionEffectTimeRange struct {
	DayInWeek         []*int32 `json:"dayInWeek,omitempty" xml:"dayInWeek,omitempty" type:"Repeated"`
	EndTimeInMinute   *int32   `json:"endTimeInMinute,omitempty" xml:"endTimeInMinute,omitempty"`
	StartTimeInMinute *int32   `json:"startTimeInMinute,omitempty" xml:"startTimeInMinute,omitempty"`
	TimeZone          *string  `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s NotifyRouteForSubscriptionEffectTimeRange) String() string {
	return dara.Prettify(s)
}

func (s NotifyRouteForSubscriptionEffectTimeRange) GoString() string {
	return s.String()
}

func (s *NotifyRouteForSubscriptionEffectTimeRange) GetDayInWeek() []*int32 {
	return s.DayInWeek
}

func (s *NotifyRouteForSubscriptionEffectTimeRange) GetEndTimeInMinute() *int32 {
	return s.EndTimeInMinute
}

func (s *NotifyRouteForSubscriptionEffectTimeRange) GetStartTimeInMinute() *int32 {
	return s.StartTimeInMinute
}

func (s *NotifyRouteForSubscriptionEffectTimeRange) GetTimeZone() *string {
	return s.TimeZone
}

func (s *NotifyRouteForSubscriptionEffectTimeRange) SetDayInWeek(v []*int32) *NotifyRouteForSubscriptionEffectTimeRange {
	s.DayInWeek = v
	return s
}

func (s *NotifyRouteForSubscriptionEffectTimeRange) SetEndTimeInMinute(v int32) *NotifyRouteForSubscriptionEffectTimeRange {
	s.EndTimeInMinute = &v
	return s
}

func (s *NotifyRouteForSubscriptionEffectTimeRange) SetStartTimeInMinute(v int32) *NotifyRouteForSubscriptionEffectTimeRange {
	s.StartTimeInMinute = &v
	return s
}

func (s *NotifyRouteForSubscriptionEffectTimeRange) SetTimeZone(v string) *NotifyRouteForSubscriptionEffectTimeRange {
	s.TimeZone = &v
	return s
}

func (s *NotifyRouteForSubscriptionEffectTimeRange) Validate() error {
	return dara.Validate(s)
}
