// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotifyConfigUnified interface {
	dara.Model
	String() string
	GoString() string
	SetActiveDays(v []*int32) *NotifyConfigUnified
	GetActiveDays() []*int32
	SetActiveEndTime(v string) *NotifyConfigUnified
	GetActiveEndTime() *string
	SetActiveStartTime(v string) *NotifyConfigUnified
	GetActiveStartTime() *string
	SetChannels(v []*DirectNotifyChannel) *NotifyConfigUnified
	GetChannels() []*DirectNotifyChannel
	SetSilenceTimeSecs(v int32) *NotifyConfigUnified
	GetSilenceTimeSecs() *int32
	SetType(v string) *NotifyConfigUnified
	GetType() *string
	SetUtcOffset(v string) *NotifyConfigUnified
	GetUtcOffset() *string
}

type NotifyConfigUnified struct {
	// 一周中发送通知的星期，1-7
	ActiveDays []*int32 `json:"activeDays,omitempty" xml:"activeDays,omitempty" type:"Repeated"`
	// 每天通知生效结束时间
	ActiveEndTime *string `json:"activeEndTime,omitempty" xml:"activeEndTime,omitempty"`
	// 每天通知生效开始时间
	ActiveStartTime *string `json:"activeStartTime,omitempty" xml:"activeStartTime,omitempty"`
	// 通知渠道列表
	//
	// This parameter is required.
	Channels []*DirectNotifyChannel `json:"channels,omitempty" xml:"channels,omitempty" type:"Repeated"`
	// 通道沉默周期（秒）
	SilenceTimeSecs *int32 `json:"silenceTimeSecs,omitempty" xml:"silenceTimeSecs,omitempty"`
	// 通知配置类型
	//
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// UTC 时区偏移量
	UtcOffset *string `json:"utcOffset,omitempty" xml:"utcOffset,omitempty"`
}

func (s NotifyConfigUnified) String() string {
	return dara.Prettify(s)
}

func (s NotifyConfigUnified) GoString() string {
	return s.String()
}

func (s *NotifyConfigUnified) GetActiveDays() []*int32 {
	return s.ActiveDays
}

func (s *NotifyConfigUnified) GetActiveEndTime() *string {
	return s.ActiveEndTime
}

func (s *NotifyConfigUnified) GetActiveStartTime() *string {
	return s.ActiveStartTime
}

func (s *NotifyConfigUnified) GetChannels() []*DirectNotifyChannel {
	return s.Channels
}

func (s *NotifyConfigUnified) GetSilenceTimeSecs() *int32 {
	return s.SilenceTimeSecs
}

func (s *NotifyConfigUnified) GetType() *string {
	return s.Type
}

func (s *NotifyConfigUnified) GetUtcOffset() *string {
	return s.UtcOffset
}

func (s *NotifyConfigUnified) SetActiveDays(v []*int32) *NotifyConfigUnified {
	s.ActiveDays = v
	return s
}

func (s *NotifyConfigUnified) SetActiveEndTime(v string) *NotifyConfigUnified {
	s.ActiveEndTime = &v
	return s
}

func (s *NotifyConfigUnified) SetActiveStartTime(v string) *NotifyConfigUnified {
	s.ActiveStartTime = &v
	return s
}

func (s *NotifyConfigUnified) SetChannels(v []*DirectNotifyChannel) *NotifyConfigUnified {
	s.Channels = v
	return s
}

func (s *NotifyConfigUnified) SetSilenceTimeSecs(v int32) *NotifyConfigUnified {
	s.SilenceTimeSecs = &v
	return s
}

func (s *NotifyConfigUnified) SetType(v string) *NotifyConfigUnified {
	s.Type = &v
	return s
}

func (s *NotifyConfigUnified) SetUtcOffset(v string) *NotifyConfigUnified {
	s.UtcOffset = &v
	return s
}

func (s *NotifyConfigUnified) Validate() error {
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
