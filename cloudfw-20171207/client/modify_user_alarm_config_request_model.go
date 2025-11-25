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
	SetNotifyConfig(v []*ModifyUserAlarmConfigRequestNotifyConfig) *ModifyUserAlarmConfigRequest
	GetNotifyConfig() []*ModifyUserAlarmConfigRequestNotifyConfig
	SetSourceIp(v string) *ModifyUserAlarmConfigRequest
	GetSourceIp() *string
	SetUseDefaultContact(v int32) *ModifyUserAlarmConfigRequest
	GetUseDefaultContact() *int32
}

type ModifyUserAlarmConfigRequest struct {
	// This parameter is required.
	AlarmConfig []*ModifyUserAlarmConfigRequestAlarmConfig `json:"AlarmConfig,omitempty" xml:"AlarmConfig,omitempty" type:"Repeated"`
	// example:
	//
	// zh
	AlarmLang     *string                                      `json:"AlarmLang,omitempty" xml:"AlarmLang,omitempty"`
	ContactConfig []*ModifyUserAlarmConfigRequestContactConfig `json:"ContactConfig,omitempty" xml:"ContactConfig,omitempty" type:"Repeated"`
	// example:
	//
	// zh
	Lang         *string                                     `json:"Lang,omitempty" xml:"Lang,omitempty"`
	NotifyConfig []*ModifyUserAlarmConfigRequestNotifyConfig `json:"NotifyConfig,omitempty" xml:"NotifyConfig,omitempty" type:"Repeated"`
	// example:
	//
	// 117.129.64.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
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

func (s *ModifyUserAlarmConfigRequest) GetNotifyConfig() []*ModifyUserAlarmConfigRequestNotifyConfig {
	return s.NotifyConfig
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

func (s *ModifyUserAlarmConfigRequest) SetNotifyConfig(v []*ModifyUserAlarmConfigRequestNotifyConfig) *ModifyUserAlarmConfigRequest {
	s.NotifyConfig = v
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
	if s.NotifyConfig != nil {
		for _, item := range s.NotifyConfig {
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
	// example:
	//
	// 10
	AlarmHour *string `json:"AlarmHour,omitempty" xml:"AlarmHour,omitempty"`
	// example:
	//
	// 0
	AlarmNotify *string `json:"AlarmNotify,omitempty" xml:"AlarmNotify,omitempty"`
	// example:
	//
	// 0
	AlarmPeriod *string `json:"AlarmPeriod,omitempty" xml:"AlarmPeriod,omitempty"`
	// example:
	//
	// bandwidth
	AlarmType *string `json:"AlarmType,omitempty" xml:"AlarmType,omitempty"`
	// example:
	//
	// on
	AlarmValue *string `json:"AlarmValue,omitempty" xml:"AlarmValue,omitempty"`
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
	// example:
	//
	// 91632****@qq.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 1351234****
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// example:
	//
	// Ben
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *ModifyUserAlarmConfigRequestContactConfig) GetStatus() *string {
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

func (s *ModifyUserAlarmConfigRequestContactConfig) SetStatus(v string) *ModifyUserAlarmConfigRequestContactConfig {
	s.Status = &v
	return s
}

func (s *ModifyUserAlarmConfigRequestContactConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyUserAlarmConfigRequestNotifyConfig struct {
	// example:
	//
	// mail
	NotifyType *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// example:
	//
	// 1351234****
	NotifyValue *string `json:"NotifyValue,omitempty" xml:"NotifyValue,omitempty"`
}

func (s ModifyUserAlarmConfigRequestNotifyConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserAlarmConfigRequestNotifyConfig) GoString() string {
	return s.String()
}

func (s *ModifyUserAlarmConfigRequestNotifyConfig) GetNotifyType() *string {
	return s.NotifyType
}

func (s *ModifyUserAlarmConfigRequestNotifyConfig) GetNotifyValue() *string {
	return s.NotifyValue
}

func (s *ModifyUserAlarmConfigRequestNotifyConfig) SetNotifyType(v string) *ModifyUserAlarmConfigRequestNotifyConfig {
	s.NotifyType = &v
	return s
}

func (s *ModifyUserAlarmConfigRequestNotifyConfig) SetNotifyValue(v string) *ModifyUserAlarmConfigRequestNotifyConfig {
	s.NotifyValue = &v
	return s
}

func (s *ModifyUserAlarmConfigRequestNotifyConfig) Validate() error {
	return dara.Validate(s)
}
