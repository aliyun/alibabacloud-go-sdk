// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDetectConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmConfigs(v []*CreateDetectConfigRequestAlarmConfigs) *CreateDetectConfigRequest
	GetAlarmConfigs() []*CreateDetectConfigRequestAlarmConfigs
	SetClientToken(v string) *CreateDetectConfigRequest
	GetClientToken() *string
	SetCronExpression(v string) *CreateDetectConfigRequest
	GetCronExpression() *string
	SetDescription(v string) *CreateDetectConfigRequest
	GetDescription() *string
	SetDetectConfigName(v string) *CreateDetectConfigRequest
	GetDetectConfigName() *string
	SetEnabled(v bool) *CreateDetectConfigRequest
	GetEnabled() *bool
	SetTriggerType(v string) *CreateDetectConfigRequest
	GetTriggerType() *string
}

type CreateDetectConfigRequest struct {
	AlarmConfigs []*CreateDetectConfigRequestAlarmConfigs `json:"alarmConfigs,omitempty" xml:"alarmConfigs,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// 0 0 0 ? 	- 1
	CronExpression *string `json:"cronExpression,omitempty" xml:"cronExpression,omitempty"`
	// example:
	//
	// example
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	DetectConfigName *string `json:"detectConfigName,omitempty" xml:"detectConfigName,omitempty"`
	Enabled          *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// Manual
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s CreateDetectConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDetectConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateDetectConfigRequest) GetAlarmConfigs() []*CreateDetectConfigRequestAlarmConfigs {
	return s.AlarmConfigs
}

func (s *CreateDetectConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDetectConfigRequest) GetCronExpression() *string {
	return s.CronExpression
}

func (s *CreateDetectConfigRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDetectConfigRequest) GetDetectConfigName() *string {
	return s.DetectConfigName
}

func (s *CreateDetectConfigRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateDetectConfigRequest) GetTriggerType() *string {
	return s.TriggerType
}

func (s *CreateDetectConfigRequest) SetAlarmConfigs(v []*CreateDetectConfigRequestAlarmConfigs) *CreateDetectConfigRequest {
	s.AlarmConfigs = v
	return s
}

func (s *CreateDetectConfigRequest) SetClientToken(v string) *CreateDetectConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDetectConfigRequest) SetCronExpression(v string) *CreateDetectConfigRequest {
	s.CronExpression = &v
	return s
}

func (s *CreateDetectConfigRequest) SetDescription(v string) *CreateDetectConfigRequest {
	s.Description = &v
	return s
}

func (s *CreateDetectConfigRequest) SetDetectConfigName(v string) *CreateDetectConfigRequest {
	s.DetectConfigName = &v
	return s
}

func (s *CreateDetectConfigRequest) SetEnabled(v bool) *CreateDetectConfigRequest {
	s.Enabled = &v
	return s
}

func (s *CreateDetectConfigRequest) SetTriggerType(v string) *CreateDetectConfigRequest {
	s.TriggerType = &v
	return s
}

func (s *CreateDetectConfigRequest) Validate() error {
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

type CreateDetectConfigRequestAlarmConfigs struct {
	// example:
	//
	// https://metrichub-cms-cn-hangzhou.aliyuncs.com/event/notify?xxxxx
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// cms
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateDetectConfigRequestAlarmConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreateDetectConfigRequestAlarmConfigs) GoString() string {
	return s.String()
}

func (s *CreateDetectConfigRequestAlarmConfigs) GetAddress() *string {
	return s.Address
}

func (s *CreateDetectConfigRequestAlarmConfigs) GetType() *string {
	return s.Type
}

func (s *CreateDetectConfigRequestAlarmConfigs) SetAddress(v string) *CreateDetectConfigRequestAlarmConfigs {
	s.Address = &v
	return s
}

func (s *CreateDetectConfigRequestAlarmConfigs) SetType(v string) *CreateDetectConfigRequestAlarmConfigs {
	s.Type = &v
	return s
}

func (s *CreateDetectConfigRequestAlarmConfigs) Validate() error {
	return dara.Validate(s)
}
