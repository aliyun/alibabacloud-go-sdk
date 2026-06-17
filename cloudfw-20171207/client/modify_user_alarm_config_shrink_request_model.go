// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserAlarmConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmConfig(v []*ModifyUserAlarmConfigShrinkRequestAlarmConfig) *ModifyUserAlarmConfigShrinkRequest
	GetAlarmConfig() []*ModifyUserAlarmConfigShrinkRequestAlarmConfig
	SetAlarmLang(v string) *ModifyUserAlarmConfigShrinkRequest
	GetAlarmLang() *string
	SetContactConfigShrink(v string) *ModifyUserAlarmConfigShrinkRequest
	GetContactConfigShrink() *string
	SetLang(v string) *ModifyUserAlarmConfigShrinkRequest
	GetLang() *string
	SetSourceIp(v string) *ModifyUserAlarmConfigShrinkRequest
	GetSourceIp() *string
	SetUseDefaultContact(v int32) *ModifyUserAlarmConfigShrinkRequest
	GetUseDefaultContact() *int32
}

type ModifyUserAlarmConfigShrinkRequest struct {
	// Alert configuration.
	//
	// This parameter is required.
	AlarmConfig []*ModifyUserAlarmConfigShrinkRequestAlarmConfig `json:"AlarmConfig,omitempty" xml:"AlarmConfig,omitempty" type:"Repeated"`
	// Language for message notifications.
	//
	// example:
	//
	// zh
	AlarmLang *string `json:"AlarmLang,omitempty" xml:"AlarmLang,omitempty"`
	// Contact configuration.
	//
	// if can be null:
	// false
	ContactConfigShrink *string `json:"ContactConfig,omitempty" xml:"ContactConfig,omitempty"`
	// Language used for requests and responses.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Source IP address of the requester.
	//
	// example:
	//
	// 117.129.64.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// Use default contact method.
	//
	// example:
	//
	// 1
	UseDefaultContact *int32 `json:"UseDefaultContact,omitempty" xml:"UseDefaultContact,omitempty"`
}

func (s ModifyUserAlarmConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserAlarmConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserAlarmConfigShrinkRequest) GetAlarmConfig() []*ModifyUserAlarmConfigShrinkRequestAlarmConfig {
	return s.AlarmConfig
}

func (s *ModifyUserAlarmConfigShrinkRequest) GetAlarmLang() *string {
	return s.AlarmLang
}

func (s *ModifyUserAlarmConfigShrinkRequest) GetContactConfigShrink() *string {
	return s.ContactConfigShrink
}

func (s *ModifyUserAlarmConfigShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyUserAlarmConfigShrinkRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyUserAlarmConfigShrinkRequest) GetUseDefaultContact() *int32 {
	return s.UseDefaultContact
}

func (s *ModifyUserAlarmConfigShrinkRequest) SetAlarmConfig(v []*ModifyUserAlarmConfigShrinkRequestAlarmConfig) *ModifyUserAlarmConfigShrinkRequest {
	s.AlarmConfig = v
	return s
}

func (s *ModifyUserAlarmConfigShrinkRequest) SetAlarmLang(v string) *ModifyUserAlarmConfigShrinkRequest {
	s.AlarmLang = &v
	return s
}

func (s *ModifyUserAlarmConfigShrinkRequest) SetContactConfigShrink(v string) *ModifyUserAlarmConfigShrinkRequest {
	s.ContactConfigShrink = &v
	return s
}

func (s *ModifyUserAlarmConfigShrinkRequest) SetLang(v string) *ModifyUserAlarmConfigShrinkRequest {
	s.Lang = &v
	return s
}

func (s *ModifyUserAlarmConfigShrinkRequest) SetSourceIp(v string) *ModifyUserAlarmConfigShrinkRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyUserAlarmConfigShrinkRequest) SetUseDefaultContact(v int32) *ModifyUserAlarmConfigShrinkRequest {
	s.UseDefaultContact = &v
	return s
}

func (s *ModifyUserAlarmConfigShrinkRequest) Validate() error {
	if s.AlarmConfig != nil {
		for _, item := range s.AlarmConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyUserAlarmConfigShrinkRequestAlarmConfig struct {
	// Hour for alert notifications.
	//
	// example:
	//
	// 10
	AlarmHour *string `json:"AlarmHour,omitempty" xml:"AlarmHour,omitempty"`
	// Notification method.
	//
	// example:
	//
	// 0
	AlarmNotify *string `json:"AlarmNotify,omitempty" xml:"AlarmNotify,omitempty"`
	// Alert period.
	//
	// example:
	//
	// 0
	AlarmPeriod *string `json:"AlarmPeriod,omitempty" xml:"AlarmPeriod,omitempty"`
	// Alarm metric.
	//
	// example:
	//
	// bandwidth
	AlarmType *string `json:"AlarmType,omitempty" xml:"AlarmType,omitempty"`
	// Alert notification message.
	//
	// example:
	//
	// on
	AlarmValue *string `json:"AlarmValue,omitempty" xml:"AlarmValue,omitempty"`
	// Day of the week for alert notifications.
	//
	// example:
	//
	// 2
	AlarmWeekDay *string `json:"AlarmWeekDay,omitempty" xml:"AlarmWeekDay,omitempty"`
}

func (s ModifyUserAlarmConfigShrinkRequestAlarmConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserAlarmConfigShrinkRequestAlarmConfig) GoString() string {
	return s.String()
}

func (s *ModifyUserAlarmConfigShrinkRequestAlarmConfig) GetAlarmHour() *string {
	return s.AlarmHour
}

func (s *ModifyUserAlarmConfigShrinkRequestAlarmConfig) GetAlarmNotify() *string {
	return s.AlarmNotify
}

func (s *ModifyUserAlarmConfigShrinkRequestAlarmConfig) GetAlarmPeriod() *string {
	return s.AlarmPeriod
}

func (s *ModifyUserAlarmConfigShrinkRequestAlarmConfig) GetAlarmType() *string {
	return s.AlarmType
}

func (s *ModifyUserAlarmConfigShrinkRequestAlarmConfig) GetAlarmValue() *string {
	return s.AlarmValue
}

func (s *ModifyUserAlarmConfigShrinkRequestAlarmConfig) GetAlarmWeekDay() *string {
	return s.AlarmWeekDay
}

func (s *ModifyUserAlarmConfigShrinkRequestAlarmConfig) SetAlarmHour(v string) *ModifyUserAlarmConfigShrinkRequestAlarmConfig {
	s.AlarmHour = &v
	return s
}

func (s *ModifyUserAlarmConfigShrinkRequestAlarmConfig) SetAlarmNotify(v string) *ModifyUserAlarmConfigShrinkRequestAlarmConfig {
	s.AlarmNotify = &v
	return s
}

func (s *ModifyUserAlarmConfigShrinkRequestAlarmConfig) SetAlarmPeriod(v string) *ModifyUserAlarmConfigShrinkRequestAlarmConfig {
	s.AlarmPeriod = &v
	return s
}

func (s *ModifyUserAlarmConfigShrinkRequestAlarmConfig) SetAlarmType(v string) *ModifyUserAlarmConfigShrinkRequestAlarmConfig {
	s.AlarmType = &v
	return s
}

func (s *ModifyUserAlarmConfigShrinkRequestAlarmConfig) SetAlarmValue(v string) *ModifyUserAlarmConfigShrinkRequestAlarmConfig {
	s.AlarmValue = &v
	return s
}

func (s *ModifyUserAlarmConfigShrinkRequestAlarmConfig) SetAlarmWeekDay(v string) *ModifyUserAlarmConfigShrinkRequestAlarmConfig {
	s.AlarmWeekDay = &v
	return s
}

func (s *ModifyUserAlarmConfigShrinkRequestAlarmConfig) Validate() error {
	return dara.Validate(s)
}
