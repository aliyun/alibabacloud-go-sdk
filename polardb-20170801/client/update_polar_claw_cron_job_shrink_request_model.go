// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolarClawCronJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *UpdatePolarClawCronJobShrinkRequest
	GetAgentId() *string
	SetApplicationId(v string) *UpdatePolarClawCronJobShrinkRequest
	GetApplicationId() *string
	SetDeleteAfterRun(v bool) *UpdatePolarClawCronJobShrinkRequest
	GetDeleteAfterRun() *bool
	SetDeliveryShrink(v string) *UpdatePolarClawCronJobShrinkRequest
	GetDeliveryShrink() *string
	SetDescription(v string) *UpdatePolarClawCronJobShrinkRequest
	GetDescription() *string
	SetEnabled(v bool) *UpdatePolarClawCronJobShrinkRequest
	GetEnabled() *bool
	SetFailureAlertShrink(v string) *UpdatePolarClawCronJobShrinkRequest
	GetFailureAlertShrink() *string
	SetJobId(v string) *UpdatePolarClawCronJobShrinkRequest
	GetJobId() *string
	SetName(v string) *UpdatePolarClawCronJobShrinkRequest
	GetName() *string
	SetPayloadShrink(v string) *UpdatePolarClawCronJobShrinkRequest
	GetPayloadShrink() *string
	SetRestart(v bool) *UpdatePolarClawCronJobShrinkRequest
	GetRestart() *bool
	SetScheduleShrink(v string) *UpdatePolarClawCronJobShrinkRequest
	GetScheduleShrink() *string
	SetSessionKey(v string) *UpdatePolarClawCronJobShrinkRequest
	GetSessionKey() *string
	SetSessionTarget(v string) *UpdatePolarClawCronJobShrinkRequest
	GetSessionTarget() *string
	SetWakeMode(v string) *UpdatePolarClawCronJobShrinkRequest
	GetWakeMode() *string
}

type UpdatePolarClawCronJobShrinkRequest struct {
	// example:
	//
	// work
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pa-xxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// true
	DeleteAfterRun *bool `json:"DeleteAfterRun,omitempty" xml:"DeleteAfterRun,omitempty"`
	// example:
	//
	// {"Mode":"announce","Channel":"telegram"}
	DeliveryShrink *string `json:"Delivery,omitempty" xml:"Delivery,omitempty"`
	// example:
	//
	// Daily report generation
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// {"After":3,"Channel":"telegram"}
	FailureAlertShrink *string `json:"FailureAlert,omitempty" xml:"FailureAlert,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0ee00f56-f467-4d41-858c-ca4ede2c770e
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// afternoon-report
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {"Kind":"agentTurn","Message":"Updated: Generate afternoon report."}
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// example:
	//
	// true
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
	// example:
	//
	// {"Kind":"cron","Expr":"0 12 	- 	- *","Tz":"America/New_York"}
	ScheduleShrink *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// example:
	//
	// agent:main:feishu:direct:***
	SessionKey *string `json:"SessionKey,omitempty" xml:"SessionKey,omitempty"`
	// example:
	//
	// isolated
	SessionTarget *string `json:"SessionTarget,omitempty" xml:"SessionTarget,omitempty"`
	// example:
	//
	// now
	WakeMode *string `json:"WakeMode,omitempty" xml:"WakeMode,omitempty"`
}

func (s UpdatePolarClawCronJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawCronJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawCronJobShrinkRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdatePolarClawCronJobShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdatePolarClawCronJobShrinkRequest) GetDeleteAfterRun() *bool {
	return s.DeleteAfterRun
}

func (s *UpdatePolarClawCronJobShrinkRequest) GetDeliveryShrink() *string {
	return s.DeliveryShrink
}

func (s *UpdatePolarClawCronJobShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdatePolarClawCronJobShrinkRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdatePolarClawCronJobShrinkRequest) GetFailureAlertShrink() *string {
	return s.FailureAlertShrink
}

func (s *UpdatePolarClawCronJobShrinkRequest) GetJobId() *string {
	return s.JobId
}

func (s *UpdatePolarClawCronJobShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdatePolarClawCronJobShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *UpdatePolarClawCronJobShrinkRequest) GetRestart() *bool {
	return s.Restart
}

func (s *UpdatePolarClawCronJobShrinkRequest) GetScheduleShrink() *string {
	return s.ScheduleShrink
}

func (s *UpdatePolarClawCronJobShrinkRequest) GetSessionKey() *string {
	return s.SessionKey
}

func (s *UpdatePolarClawCronJobShrinkRequest) GetSessionTarget() *string {
	return s.SessionTarget
}

func (s *UpdatePolarClawCronJobShrinkRequest) GetWakeMode() *string {
	return s.WakeMode
}

func (s *UpdatePolarClawCronJobShrinkRequest) SetAgentId(v string) *UpdatePolarClawCronJobShrinkRequest {
	s.AgentId = &v
	return s
}

func (s *UpdatePolarClawCronJobShrinkRequest) SetApplicationId(v string) *UpdatePolarClawCronJobShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdatePolarClawCronJobShrinkRequest) SetDeleteAfterRun(v bool) *UpdatePolarClawCronJobShrinkRequest {
	s.DeleteAfterRun = &v
	return s
}

func (s *UpdatePolarClawCronJobShrinkRequest) SetDeliveryShrink(v string) *UpdatePolarClawCronJobShrinkRequest {
	s.DeliveryShrink = &v
	return s
}

func (s *UpdatePolarClawCronJobShrinkRequest) SetDescription(v string) *UpdatePolarClawCronJobShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdatePolarClawCronJobShrinkRequest) SetEnabled(v bool) *UpdatePolarClawCronJobShrinkRequest {
	s.Enabled = &v
	return s
}

func (s *UpdatePolarClawCronJobShrinkRequest) SetFailureAlertShrink(v string) *UpdatePolarClawCronJobShrinkRequest {
	s.FailureAlertShrink = &v
	return s
}

func (s *UpdatePolarClawCronJobShrinkRequest) SetJobId(v string) *UpdatePolarClawCronJobShrinkRequest {
	s.JobId = &v
	return s
}

func (s *UpdatePolarClawCronJobShrinkRequest) SetName(v string) *UpdatePolarClawCronJobShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdatePolarClawCronJobShrinkRequest) SetPayloadShrink(v string) *UpdatePolarClawCronJobShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *UpdatePolarClawCronJobShrinkRequest) SetRestart(v bool) *UpdatePolarClawCronJobShrinkRequest {
	s.Restart = &v
	return s
}

func (s *UpdatePolarClawCronJobShrinkRequest) SetScheduleShrink(v string) *UpdatePolarClawCronJobShrinkRequest {
	s.ScheduleShrink = &v
	return s
}

func (s *UpdatePolarClawCronJobShrinkRequest) SetSessionKey(v string) *UpdatePolarClawCronJobShrinkRequest {
	s.SessionKey = &v
	return s
}

func (s *UpdatePolarClawCronJobShrinkRequest) SetSessionTarget(v string) *UpdatePolarClawCronJobShrinkRequest {
	s.SessionTarget = &v
	return s
}

func (s *UpdatePolarClawCronJobShrinkRequest) SetWakeMode(v string) *UpdatePolarClawCronJobShrinkRequest {
	s.WakeMode = &v
	return s
}

func (s *UpdatePolarClawCronJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
