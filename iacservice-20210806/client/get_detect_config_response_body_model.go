// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDetectConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetectConfig(v *GetDetectConfigResponseBodyDetectConfig) *GetDetectConfigResponseBody
	GetDetectConfig() *GetDetectConfigResponseBodyDetectConfig
	SetRequestId(v string) *GetDetectConfigResponseBody
	GetRequestId() *string
}

type GetDetectConfigResponseBody struct {
	DetectConfig *GetDetectConfigResponseBodyDetectConfig `json:"detectConfig,omitempty" xml:"detectConfig,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// BF72A6FB-B071-5F2E-A036-9D62545B962C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetDetectConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDetectConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetDetectConfigResponseBody) GetDetectConfig() *GetDetectConfigResponseBodyDetectConfig {
	return s.DetectConfig
}

func (s *GetDetectConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDetectConfigResponseBody) SetDetectConfig(v *GetDetectConfigResponseBodyDetectConfig) *GetDetectConfigResponseBody {
	s.DetectConfig = v
	return s
}

func (s *GetDetectConfigResponseBody) SetRequestId(v string) *GetDetectConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDetectConfigResponseBody) Validate() error {
	if s.DetectConfig != nil {
		if err := s.DetectConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDetectConfigResponseBodyDetectConfig struct {
	AlarmConfigs []*GetDetectConfigResponseBodyDetectConfigAlarmConfigs `json:"alarmConfigs,omitempty" xml:"alarmConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// 2026-04-10T02:30:04Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 0 0 0 ? 	- 1
	CronExpression *string `json:"cronExpression,omitempty" xml:"cronExpression,omitempty"`
	// example:
	//
	// this is a description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// dc-xxxx
	DetectConfigId *string `json:"detectConfigId,omitempty" xml:"detectConfigId,omitempty"`
	// example:
	//
	// test
	DetectConfigName *string `json:"detectConfigName,omitempty" xml:"detectConfigName,omitempty"`
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// Cron
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s GetDetectConfigResponseBodyDetectConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDetectConfigResponseBodyDetectConfig) GoString() string {
	return s.String()
}

func (s *GetDetectConfigResponseBodyDetectConfig) GetAlarmConfigs() []*GetDetectConfigResponseBodyDetectConfigAlarmConfigs {
	return s.AlarmConfigs
}

func (s *GetDetectConfigResponseBodyDetectConfig) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetDetectConfigResponseBodyDetectConfig) GetCronExpression() *string {
	return s.CronExpression
}

func (s *GetDetectConfigResponseBodyDetectConfig) GetDescription() *string {
	return s.Description
}

func (s *GetDetectConfigResponseBodyDetectConfig) GetDetectConfigId() *string {
	return s.DetectConfigId
}

func (s *GetDetectConfigResponseBodyDetectConfig) GetDetectConfigName() *string {
	return s.DetectConfigName
}

func (s *GetDetectConfigResponseBodyDetectConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetDetectConfigResponseBodyDetectConfig) GetTriggerType() *string {
	return s.TriggerType
}

func (s *GetDetectConfigResponseBodyDetectConfig) SetAlarmConfigs(v []*GetDetectConfigResponseBodyDetectConfigAlarmConfigs) *GetDetectConfigResponseBodyDetectConfig {
	s.AlarmConfigs = v
	return s
}

func (s *GetDetectConfigResponseBodyDetectConfig) SetCreateTime(v string) *GetDetectConfigResponseBodyDetectConfig {
	s.CreateTime = &v
	return s
}

func (s *GetDetectConfigResponseBodyDetectConfig) SetCronExpression(v string) *GetDetectConfigResponseBodyDetectConfig {
	s.CronExpression = &v
	return s
}

func (s *GetDetectConfigResponseBodyDetectConfig) SetDescription(v string) *GetDetectConfigResponseBodyDetectConfig {
	s.Description = &v
	return s
}

func (s *GetDetectConfigResponseBodyDetectConfig) SetDetectConfigId(v string) *GetDetectConfigResponseBodyDetectConfig {
	s.DetectConfigId = &v
	return s
}

func (s *GetDetectConfigResponseBodyDetectConfig) SetDetectConfigName(v string) *GetDetectConfigResponseBodyDetectConfig {
	s.DetectConfigName = &v
	return s
}

func (s *GetDetectConfigResponseBodyDetectConfig) SetEnabled(v bool) *GetDetectConfigResponseBodyDetectConfig {
	s.Enabled = &v
	return s
}

func (s *GetDetectConfigResponseBodyDetectConfig) SetTriggerType(v string) *GetDetectConfigResponseBodyDetectConfig {
	s.TriggerType = &v
	return s
}

func (s *GetDetectConfigResponseBodyDetectConfig) Validate() error {
	if s.AlarmConfigs != nil {
		for _, item := range s.AlarmConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDetectConfigResponseBodyDetectConfigAlarmConfigs struct {
	// example:
	//
	// https://metrichub-cms-cn-hangzhou.aliyuncs.com/event/notify?xxxxx
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// cms
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetDetectConfigResponseBodyDetectConfigAlarmConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetDetectConfigResponseBodyDetectConfigAlarmConfigs) GoString() string {
	return s.String()
}

func (s *GetDetectConfigResponseBodyDetectConfigAlarmConfigs) GetAddress() *string {
	return s.Address
}

func (s *GetDetectConfigResponseBodyDetectConfigAlarmConfigs) GetType() *string {
	return s.Type
}

func (s *GetDetectConfigResponseBodyDetectConfigAlarmConfigs) SetAddress(v string) *GetDetectConfigResponseBodyDetectConfigAlarmConfigs {
	s.Address = &v
	return s
}

func (s *GetDetectConfigResponseBodyDetectConfigAlarmConfigs) SetType(v string) *GetDetectConfigResponseBodyDetectConfigAlarmConfigs {
	s.Type = &v
	return s
}

func (s *GetDetectConfigResponseBodyDetectConfigAlarmConfigs) Validate() error {
	return dara.Validate(s)
}
