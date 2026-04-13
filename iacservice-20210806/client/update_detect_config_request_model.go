// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDetectConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmConfigs(v []*UpdateDetectConfigRequestAlarmConfigs) *UpdateDetectConfigRequest
	GetAlarmConfigs() []*UpdateDetectConfigRequestAlarmConfigs
	SetClientToken(v string) *UpdateDetectConfigRequest
	GetClientToken() *string
	SetCronExpression(v string) *UpdateDetectConfigRequest
	GetCronExpression() *string
	SetDescription(v string) *UpdateDetectConfigRequest
	GetDescription() *string
	SetDetectConfigName(v string) *UpdateDetectConfigRequest
	GetDetectConfigName() *string
	SetEnabled(v bool) *UpdateDetectConfigRequest
	GetEnabled() *bool
	SetTriggerType(v string) *UpdateDetectConfigRequest
	GetTriggerType() *string
}

type UpdateDetectConfigRequest struct {
	AlarmConfigs []*UpdateDetectConfigRequestAlarmConfigs `json:"alarmConfigs,omitempty" xml:"alarmConfigs,omitempty" type:"Repeated"`
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
	// this is description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
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
	// Manual
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s UpdateDetectConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDetectConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateDetectConfigRequest) GetAlarmConfigs() []*UpdateDetectConfigRequestAlarmConfigs {
	return s.AlarmConfigs
}

func (s *UpdateDetectConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateDetectConfigRequest) GetCronExpression() *string {
	return s.CronExpression
}

func (s *UpdateDetectConfigRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDetectConfigRequest) GetDetectConfigName() *string {
	return s.DetectConfigName
}

func (s *UpdateDetectConfigRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateDetectConfigRequest) GetTriggerType() *string {
	return s.TriggerType
}

func (s *UpdateDetectConfigRequest) SetAlarmConfigs(v []*UpdateDetectConfigRequestAlarmConfigs) *UpdateDetectConfigRequest {
	s.AlarmConfigs = v
	return s
}

func (s *UpdateDetectConfigRequest) SetClientToken(v string) *UpdateDetectConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateDetectConfigRequest) SetCronExpression(v string) *UpdateDetectConfigRequest {
	s.CronExpression = &v
	return s
}

func (s *UpdateDetectConfigRequest) SetDescription(v string) *UpdateDetectConfigRequest {
	s.Description = &v
	return s
}

func (s *UpdateDetectConfigRequest) SetDetectConfigName(v string) *UpdateDetectConfigRequest {
	s.DetectConfigName = &v
	return s
}

func (s *UpdateDetectConfigRequest) SetEnabled(v bool) *UpdateDetectConfigRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateDetectConfigRequest) SetTriggerType(v string) *UpdateDetectConfigRequest {
	s.TriggerType = &v
	return s
}

func (s *UpdateDetectConfigRequest) Validate() error {
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

type UpdateDetectConfigRequestAlarmConfigs struct {
	// example:
	//
	// example@example.com
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// cms
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateDetectConfigRequestAlarmConfigs) String() string {
	return dara.Prettify(s)
}

func (s UpdateDetectConfigRequestAlarmConfigs) GoString() string {
	return s.String()
}

func (s *UpdateDetectConfigRequestAlarmConfigs) GetAddress() *string {
	return s.Address
}

func (s *UpdateDetectConfigRequestAlarmConfigs) GetType() *string {
	return s.Type
}

func (s *UpdateDetectConfigRequestAlarmConfigs) SetAddress(v string) *UpdateDetectConfigRequestAlarmConfigs {
	s.Address = &v
	return s
}

func (s *UpdateDetectConfigRequestAlarmConfigs) SetType(v string) *UpdateDetectConfigRequestAlarmConfigs {
	s.Type = &v
	return s
}

func (s *UpdateDetectConfigRequestAlarmConfigs) Validate() error {
	return dara.Validate(s)
}
