// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTimerTriggerConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCronExpression(v string) *TimerTriggerConfig
	GetCronExpression() *string
	SetEnable(v bool) *TimerTriggerConfig
	GetEnable() *bool
	SetPayload(v string) *TimerTriggerConfig
	GetPayload() *string
}

type TimerTriggerConfig struct {
	// The trigger period expression. You can specify to trigger based on a time interval. For example, the expression @every 4m indicates that the triggering is performed every four minutes. You can also specify to trigger based on a cron expression, for example, 0 0 4 \\	- \\	- \\*.
	//
	// example:
	//
	// 0 0 4 	- 	- *
	CronExpression *string `json:"cronExpression,omitempty" xml:"cronExpression,omitempty"`
	// Specify whether to enable the trigger.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Enter custom parameters. The trigger message is used as the value of the payload in the event.
	//
	// example:
	//
	// {"workflowInstanceId":"39639"}
	Payload *string `json:"payload,omitempty" xml:"payload,omitempty"`
}

func (s TimerTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s TimerTriggerConfig) GoString() string {
	return s.String()
}

func (s *TimerTriggerConfig) GetCronExpression() *string {
	return s.CronExpression
}

func (s *TimerTriggerConfig) GetEnable() *bool {
	return s.Enable
}

func (s *TimerTriggerConfig) GetPayload() *string {
	return s.Payload
}

func (s *TimerTriggerConfig) SetCronExpression(v string) *TimerTriggerConfig {
	s.CronExpression = &v
	return s
}

func (s *TimerTriggerConfig) SetEnable(v bool) *TimerTriggerConfig {
	s.Enable = &v
	return s
}

func (s *TimerTriggerConfig) SetPayload(v string) *TimerTriggerConfig {
	s.Payload = &v
	return s
}

func (s *TimerTriggerConfig) Validate() error {
	return dara.Validate(s)
}
