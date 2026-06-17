// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserAlarmConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmConfig(v []*ModifyUserAlarmConfigRequestAlarmConfig) *ModifyUserAlarmConfigRequest
	GetAlarmConfig() []*ModifyUserAlarmConfigRequestAlarmConfig
	SetAlarmLang(v string) *ModifyUserAlarmConfigRequest
	GetAlarmLang() *string
	SetContactConfig(v []*ModifyUserAlarmConfigRequestContactConfig) *ModifyUserAlarmConfigRequest
	GetContactConfig() []*ModifyUserAlarmConfigRequestContactConfig
	SetLang(v string) *ModifyUserAlarmConfigRequest
	GetLang() *string
	SetSourceIp(v string) *ModifyUserAlarmConfigRequest
	GetSourceIp() *string
	SetUseDefaultContact(v int32) *ModifyUserAlarmConfigRequest
	GetUseDefaultContact() *int32
}

type ModifyUserAlarmConfigRequest struct {
	// Alert configuration.
	//
	// This parameter is required.
	AlarmConfig []*ModifyUserAlarmConfigRequestAlarmConfig `json:"AlarmConfig,omitempty" xml:"AlarmConfig,omitempty" type:"Repeated"`
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
	ContactConfig []*ModifyUserAlarmConfigRequestContactConfig `json:"ContactConfig,omitempty" xml:"ContactConfig,omitempty" type:"Repeated"`
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

func (s ModifyUserAlarmConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserAlarmConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserAlarmConfigRequest) GetAlarmConfig() []*ModifyUserAlarmConfigRequestAlarmConfig {
	return s.AlarmConfig
}

func (s *ModifyUserAlarmConfigRequest) GetAlarmLang() *string {
	return s.AlarmLang
}

func (s *ModifyUserAlarmConfigRequest) GetContactConfig() []*ModifyUserAlarmConfigRequestContactConfig {
	return s.ContactConfig
}

func (s *ModifyUserAlarmConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyUserAlarmConfigRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyUserAlarmConfigRequest) GetUseDefaultContact() *int32 {
	return s.UseDefaultContact
}

func (s *ModifyUserAlarmConfigRequest) SetAlarmConfig(v []*ModifyUserAlarmConfigRequestAlarmConfig) *ModifyUserAlarmConfigRequest {
	s.AlarmConfig = v
	return s
}

func (s *ModifyUserAlarmConfigRequest) SetAlarmLang(v string) *ModifyUserAlarmConfigRequest {
	s.AlarmLang = &v
	return s
}

func (s *ModifyUserAlarmConfigRequest) SetContactConfig(v []*ModifyUserAlarmConfigRequestContactConfig) *ModifyUserAlarmConfigRequest {
	s.ContactConfig = v
	return s
}

func (s *ModifyUserAlarmConfigRequest) SetLang(v string) *ModifyUserAlarmConfigRequest {
	s.Lang = &v
	return s
}

func (s *ModifyUserAlarmConfigRequest) SetSourceIp(v string) *ModifyUserAlarmConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyUserAlarmConfigRequest) SetUseDefaultContact(v int32) *ModifyUserAlarmConfigRequest {
	s.UseDefaultContact = &v
	return s
}

func (s *ModifyUserAlarmConfigRequest) Validate() error {
	if s.AlarmConfig != nil {
		for _, item := range s.AlarmConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ContactConfig != nil {
		for _, item := range s.ContactConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyUserAlarmConfigRequestAlarmConfig struct {
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

func (s ModifyUserAlarmConfigRequestAlarmConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserAlarmConfigRequestAlarmConfig) GoString() string {
	return s.String()
}

func (s *ModifyUserAlarmConfigRequestAlarmConfig) GetAlarmHour() *string {
	return s.AlarmHour
}

func (s *ModifyUserAlarmConfigRequestAlarmConfig) GetAlarmNotify() *string {
	return s.AlarmNotify
}

func (s *ModifyUserAlarmConfigRequestAlarmConfig) GetAlarmPeriod() *string {
	return s.AlarmPeriod
}

func (s *ModifyUserAlarmConfigRequestAlarmConfig) GetAlarmType() *string {
	return s.AlarmType
}

func (s *ModifyUserAlarmConfigRequestAlarmConfig) GetAlarmValue() *string {
	return s.AlarmValue
}

func (s *ModifyUserAlarmConfigRequestAlarmConfig) GetAlarmWeekDay() *string {
	return s.AlarmWeekDay
}

func (s *ModifyUserAlarmConfigRequestAlarmConfig) SetAlarmHour(v string) *ModifyUserAlarmConfigRequestAlarmConfig {
	s.AlarmHour = &v
	return s
}

func (s *ModifyUserAlarmConfigRequestAlarmConfig) SetAlarmNotify(v string) *ModifyUserAlarmConfigRequestAlarmConfig {
	s.AlarmNotify = &v
	return s
}

func (s *ModifyUserAlarmConfigRequestAlarmConfig) SetAlarmPeriod(v string) *ModifyUserAlarmConfigRequestAlarmConfig {
	s.AlarmPeriod = &v
	return s
}

func (s *ModifyUserAlarmConfigRequestAlarmConfig) SetAlarmType(v string) *ModifyUserAlarmConfigRequestAlarmConfig {
	s.AlarmType = &v
	return s
}

func (s *ModifyUserAlarmConfigRequestAlarmConfig) SetAlarmValue(v string) *ModifyUserAlarmConfigRequestAlarmConfig {
	s.AlarmValue = &v
	return s
}

func (s *ModifyUserAlarmConfigRequestAlarmConfig) SetAlarmWeekDay(v string) *ModifyUserAlarmConfigRequestAlarmConfig {
	s.AlarmWeekDay = &v
	return s
}

func (s *ModifyUserAlarmConfigRequestAlarmConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyUserAlarmConfigRequestContactConfig struct {
	// Mailbox.
	//
	// example:
	//
	// 91632****@qq.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Mobile number.
	//
	// example:
	//
	// 1351234****
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// Alert notification recipient.
	//
	// example:
	//
	// Ben
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Alert status.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyUserAlarmConfigRequestContactConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserAlarmConfigRequestContactConfig) GoString() string {
	return s.String()
}

func (s *ModifyUserAlarmConfigRequestContactConfig) GetEmail() *string {
	return s.Email
}

func (s *ModifyUserAlarmConfigRequestContactConfig) GetMobilePhone() *string {
	return s.MobilePhone
}

func (s *ModifyUserAlarmConfigRequestContactConfig) GetName() *string {
	return s.Name
}

func (s *ModifyUserAlarmConfigRequestContactConfig) GetStatus() *int32 {
	return s.Status
}

func (s *ModifyUserAlarmConfigRequestContactConfig) SetEmail(v string) *ModifyUserAlarmConfigRequestContactConfig {
	s.Email = &v
	return s
}

func (s *ModifyUserAlarmConfigRequestContactConfig) SetMobilePhone(v string) *ModifyUserAlarmConfigRequestContactConfig {
	s.MobilePhone = &v
	return s
}

func (s *ModifyUserAlarmConfigRequestContactConfig) SetName(v string) *ModifyUserAlarmConfigRequestContactConfig {
	s.Name = &v
	return s
}

func (s *ModifyUserAlarmConfigRequestContactConfig) SetStatus(v int32) *ModifyUserAlarmConfigRequestContactConfig {
	s.Status = &v
	return s
}

func (s *ModifyUserAlarmConfigRequestContactConfig) Validate() error {
	return dara.Validate(s)
}
