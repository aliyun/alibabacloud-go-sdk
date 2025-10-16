// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserAlarmConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmConfig(v []*DescribeUserAlarmConfigResponseBodyAlarmConfig) *DescribeUserAlarmConfigResponseBody
	GetAlarmConfig() []*DescribeUserAlarmConfigResponseBodyAlarmConfig
	SetAlarmLang(v string) *DescribeUserAlarmConfigResponseBody
	GetAlarmLang() *string
	SetContactConfig(v []*DescribeUserAlarmConfigResponseBodyContactConfig) *DescribeUserAlarmConfigResponseBody
	GetContactConfig() []*DescribeUserAlarmConfigResponseBodyContactConfig
	SetDefaultContact(v *DescribeUserAlarmConfigResponseBodyDefaultContact) *DescribeUserAlarmConfigResponseBody
	GetDefaultContact() *DescribeUserAlarmConfigResponseBodyDefaultContact
	SetNotifyConfig(v []*DescribeUserAlarmConfigResponseBodyNotifyConfig) *DescribeUserAlarmConfigResponseBody
	GetNotifyConfig() []*DescribeUserAlarmConfigResponseBodyNotifyConfig
	SetRequestId(v string) *DescribeUserAlarmConfigResponseBody
	GetRequestId() *string
}

type DescribeUserAlarmConfigResponseBody struct {
	AlarmConfig []*DescribeUserAlarmConfigResponseBodyAlarmConfig `json:"AlarmConfig,omitempty" xml:"AlarmConfig,omitempty" type:"Repeated"`
	// example:
	//
	// zh
	AlarmLang      *string                                             `json:"AlarmLang,omitempty" xml:"AlarmLang,omitempty"`
	ContactConfig  []*DescribeUserAlarmConfigResponseBodyContactConfig `json:"ContactConfig,omitempty" xml:"ContactConfig,omitempty" type:"Repeated"`
	DefaultContact *DescribeUserAlarmConfigResponseBodyDefaultContact  `json:"DefaultContact,omitempty" xml:"DefaultContact,omitempty" type:"Struct"`
	NotifyConfig   []*DescribeUserAlarmConfigResponseBodyNotifyConfig  `json:"NotifyConfig,omitempty" xml:"NotifyConfig,omitempty" type:"Repeated"`
	// example:
	//
	// 9D250177-4F11-58B8-9AFE-A4624FF1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserAlarmConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAlarmConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserAlarmConfigResponseBody) GetAlarmConfig() []*DescribeUserAlarmConfigResponseBodyAlarmConfig {
	return s.AlarmConfig
}

func (s *DescribeUserAlarmConfigResponseBody) GetAlarmLang() *string {
	return s.AlarmLang
}

func (s *DescribeUserAlarmConfigResponseBody) GetContactConfig() []*DescribeUserAlarmConfigResponseBodyContactConfig {
	return s.ContactConfig
}

func (s *DescribeUserAlarmConfigResponseBody) GetDefaultContact() *DescribeUserAlarmConfigResponseBodyDefaultContact {
	return s.DefaultContact
}

func (s *DescribeUserAlarmConfigResponseBody) GetNotifyConfig() []*DescribeUserAlarmConfigResponseBodyNotifyConfig {
	return s.NotifyConfig
}

func (s *DescribeUserAlarmConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserAlarmConfigResponseBody) SetAlarmConfig(v []*DescribeUserAlarmConfigResponseBodyAlarmConfig) *DescribeUserAlarmConfigResponseBody {
	s.AlarmConfig = v
	return s
}

func (s *DescribeUserAlarmConfigResponseBody) SetAlarmLang(v string) *DescribeUserAlarmConfigResponseBody {
	s.AlarmLang = &v
	return s
}

func (s *DescribeUserAlarmConfigResponseBody) SetContactConfig(v []*DescribeUserAlarmConfigResponseBodyContactConfig) *DescribeUserAlarmConfigResponseBody {
	s.ContactConfig = v
	return s
}

func (s *DescribeUserAlarmConfigResponseBody) SetDefaultContact(v *DescribeUserAlarmConfigResponseBodyDefaultContact) *DescribeUserAlarmConfigResponseBody {
	s.DefaultContact = v
	return s
}

func (s *DescribeUserAlarmConfigResponseBody) SetNotifyConfig(v []*DescribeUserAlarmConfigResponseBodyNotifyConfig) *DescribeUserAlarmConfigResponseBody {
	s.NotifyConfig = v
	return s
}

func (s *DescribeUserAlarmConfigResponseBody) SetRequestId(v string) *DescribeUserAlarmConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserAlarmConfigResponseBody) Validate() error {
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
	if s.DefaultContact != nil {
		if err := s.DefaultContact.Validate(); err != nil {
			return err
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

type DescribeUserAlarmConfigResponseBodyAlarmConfig struct {
	// example:
	//
	// 0
	AlarmHour *int32 `json:"AlarmHour,omitempty" xml:"AlarmHour,omitempty"`
	// example:
	//
	// 1
	AlarmNotify *int32 `json:"AlarmNotify,omitempty" xml:"AlarmNotify,omitempty"`
	// example:
	//
	// 30
	AlarmPeriod *int32 `json:"AlarmPeriod,omitempty" xml:"AlarmPeriod,omitempty"`
	// example:
	//
	// bandwidth
	AlarmType *string `json:"AlarmType,omitempty" xml:"AlarmType,omitempty"`
	// example:
	//
	// 80
	AlarmValue *string `json:"AlarmValue,omitempty" xml:"AlarmValue,omitempty"`
	// example:
	//
	// 0
	AlarmWeekDay *int32 `json:"AlarmWeekDay,omitempty" xml:"AlarmWeekDay,omitempty"`
}

func (s DescribeUserAlarmConfigResponseBodyAlarmConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAlarmConfigResponseBodyAlarmConfig) GoString() string {
	return s.String()
}

func (s *DescribeUserAlarmConfigResponseBodyAlarmConfig) GetAlarmHour() *int32 {
	return s.AlarmHour
}

func (s *DescribeUserAlarmConfigResponseBodyAlarmConfig) GetAlarmNotify() *int32 {
	return s.AlarmNotify
}

func (s *DescribeUserAlarmConfigResponseBodyAlarmConfig) GetAlarmPeriod() *int32 {
	return s.AlarmPeriod
}

func (s *DescribeUserAlarmConfigResponseBodyAlarmConfig) GetAlarmType() *string {
	return s.AlarmType
}

func (s *DescribeUserAlarmConfigResponseBodyAlarmConfig) GetAlarmValue() *string {
	return s.AlarmValue
}

func (s *DescribeUserAlarmConfigResponseBodyAlarmConfig) GetAlarmWeekDay() *int32 {
	return s.AlarmWeekDay
}

func (s *DescribeUserAlarmConfigResponseBodyAlarmConfig) SetAlarmHour(v int32) *DescribeUserAlarmConfigResponseBodyAlarmConfig {
	s.AlarmHour = &v
	return s
}

func (s *DescribeUserAlarmConfigResponseBodyAlarmConfig) SetAlarmNotify(v int32) *DescribeUserAlarmConfigResponseBodyAlarmConfig {
	s.AlarmNotify = &v
	return s
}

func (s *DescribeUserAlarmConfigResponseBodyAlarmConfig) SetAlarmPeriod(v int32) *DescribeUserAlarmConfigResponseBodyAlarmConfig {
	s.AlarmPeriod = &v
	return s
}

func (s *DescribeUserAlarmConfigResponseBodyAlarmConfig) SetAlarmType(v string) *DescribeUserAlarmConfigResponseBodyAlarmConfig {
	s.AlarmType = &v
	return s
}

func (s *DescribeUserAlarmConfigResponseBodyAlarmConfig) SetAlarmValue(v string) *DescribeUserAlarmConfigResponseBodyAlarmConfig {
	s.AlarmValue = &v
	return s
}

func (s *DescribeUserAlarmConfigResponseBodyAlarmConfig) SetAlarmWeekDay(v int32) *DescribeUserAlarmConfigResponseBodyAlarmConfig {
	s.AlarmWeekDay = &v
	return s
}

func (s *DescribeUserAlarmConfigResponseBodyAlarmConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeUserAlarmConfigResponseBodyContactConfig struct {
	// example:
	//
	// 1530811****@qq.com
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// example:
	//
	// 1531123****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeUserAlarmConfigResponseBodyContactConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAlarmConfigResponseBodyContactConfig) GoString() string {
	return s.String()
}

func (s *DescribeUserAlarmConfigResponseBodyContactConfig) GetEmail() *string {
	return s.Email
}

func (s *DescribeUserAlarmConfigResponseBodyContactConfig) GetMobilePhone() *string {
	return s.MobilePhone
}

func (s *DescribeUserAlarmConfigResponseBodyContactConfig) GetName() *string {
	return s.Name
}

func (s *DescribeUserAlarmConfigResponseBodyContactConfig) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeUserAlarmConfigResponseBodyContactConfig) SetEmail(v string) *DescribeUserAlarmConfigResponseBodyContactConfig {
	s.Email = &v
	return s
}

func (s *DescribeUserAlarmConfigResponseBodyContactConfig) SetMobilePhone(v string) *DescribeUserAlarmConfigResponseBodyContactConfig {
	s.MobilePhone = &v
	return s
}

func (s *DescribeUserAlarmConfigResponseBodyContactConfig) SetName(v string) *DescribeUserAlarmConfigResponseBodyContactConfig {
	s.Name = &v
	return s
}

func (s *DescribeUserAlarmConfigResponseBodyContactConfig) SetStatus(v int32) *DescribeUserAlarmConfigResponseBodyContactConfig {
	s.Status = &v
	return s
}

func (s *DescribeUserAlarmConfigResponseBodyContactConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeUserAlarmConfigResponseBodyDefaultContact struct {
	// example:
	//
	// 1530811****@qq.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 1531123****
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeUserAlarmConfigResponseBodyDefaultContact) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAlarmConfigResponseBodyDefaultContact) GoString() string {
	return s.String()
}

func (s *DescribeUserAlarmConfigResponseBodyDefaultContact) GetEmail() *string {
	return s.Email
}

func (s *DescribeUserAlarmConfigResponseBodyDefaultContact) GetMobilePhone() *string {
	return s.MobilePhone
}

func (s *DescribeUserAlarmConfigResponseBodyDefaultContact) GetName() *string {
	return s.Name
}

func (s *DescribeUserAlarmConfigResponseBodyDefaultContact) GetStatus() *string {
	return s.Status
}

func (s *DescribeUserAlarmConfigResponseBodyDefaultContact) SetEmail(v string) *DescribeUserAlarmConfigResponseBodyDefaultContact {
	s.Email = &v
	return s
}

func (s *DescribeUserAlarmConfigResponseBodyDefaultContact) SetMobilePhone(v string) *DescribeUserAlarmConfigResponseBodyDefaultContact {
	s.MobilePhone = &v
	return s
}

func (s *DescribeUserAlarmConfigResponseBodyDefaultContact) SetName(v string) *DescribeUserAlarmConfigResponseBodyDefaultContact {
	s.Name = &v
	return s
}

func (s *DescribeUserAlarmConfigResponseBodyDefaultContact) SetStatus(v string) *DescribeUserAlarmConfigResponseBodyDefaultContact {
	s.Status = &v
	return s
}

func (s *DescribeUserAlarmConfigResponseBodyDefaultContact) Validate() error {
	return dara.Validate(s)
}

type DescribeUserAlarmConfigResponseBodyNotifyConfig struct {
	// example:
	//
	// mail
	NotifyType *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// example:
	//
	// 0
	NotifyValue *string `json:"NotifyValue,omitempty" xml:"NotifyValue,omitempty"`
}

func (s DescribeUserAlarmConfigResponseBodyNotifyConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAlarmConfigResponseBodyNotifyConfig) GoString() string {
	return s.String()
}

func (s *DescribeUserAlarmConfigResponseBodyNotifyConfig) GetNotifyType() *string {
	return s.NotifyType
}

func (s *DescribeUserAlarmConfigResponseBodyNotifyConfig) GetNotifyValue() *string {
	return s.NotifyValue
}

func (s *DescribeUserAlarmConfigResponseBodyNotifyConfig) SetNotifyType(v string) *DescribeUserAlarmConfigResponseBodyNotifyConfig {
	s.NotifyType = &v
	return s
}

func (s *DescribeUserAlarmConfigResponseBodyNotifyConfig) SetNotifyValue(v string) *DescribeUserAlarmConfigResponseBodyNotifyConfig {
	s.NotifyValue = &v
	return s
}

func (s *DescribeUserAlarmConfigResponseBodyNotifyConfig) Validate() error {
	return dara.Validate(s)
}
