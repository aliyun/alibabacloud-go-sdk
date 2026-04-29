// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolarClawCronJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *CreatePolarClawCronJobShrinkRequest
	GetAgentId() *string
	SetApplicationId(v string) *CreatePolarClawCronJobShrinkRequest
	GetApplicationId() *string
	SetDeleteAfterRun(v bool) *CreatePolarClawCronJobShrinkRequest
	GetDeleteAfterRun() *bool
	SetDeliveryShrink(v string) *CreatePolarClawCronJobShrinkRequest
	GetDeliveryShrink() *string
	SetDescription(v string) *CreatePolarClawCronJobShrinkRequest
	GetDescription() *string
	SetEnabled(v bool) *CreatePolarClawCronJobShrinkRequest
	GetEnabled() *bool
	SetFailureAlertShrink(v string) *CreatePolarClawCronJobShrinkRequest
	GetFailureAlertShrink() *string
	SetName(v string) *CreatePolarClawCronJobShrinkRequest
	GetName() *string
	SetPayloadShrink(v string) *CreatePolarClawCronJobShrinkRequest
	GetPayloadShrink() *string
	SetRestart(v bool) *CreatePolarClawCronJobShrinkRequest
	GetRestart() *bool
	SetRunImmediately(v bool) *CreatePolarClawCronJobShrinkRequest
	GetRunImmediately() *bool
	SetScheduleShrink(v string) *CreatePolarClawCronJobShrinkRequest
	GetScheduleShrink() *string
	SetSessionKey(v string) *CreatePolarClawCronJobShrinkRequest
	GetSessionKey() *string
	SetSessionTarget(v string) *CreatePolarClawCronJobShrinkRequest
	GetSessionTarget() *string
	SetWakeMode(v string) *CreatePolarClawCronJobShrinkRequest
	GetWakeMode() *string
}

type CreatePolarClawCronJobShrinkRequest struct {
	// example:
	//
	// main
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// false
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
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// {"After":3,"Channel":"telegram"}
	FailureAlertShrink *string `json:"FailureAlert,omitempty" xml:"FailureAlert,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// daily-report
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"Kind":"agentTurn","Message":"Generate the daily report."}
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// example:
	//
	// true
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
	// example:
	//
	// false
	RunImmediately *bool `json:"RunImmediately,omitempty" xml:"RunImmediately,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"Kind":"cron","Expr":"0 9 	- 	- *","Tz":"Asia/Shanghai"}
	ScheduleShrink *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// example:
	//
	// agent:main:feishu:direct:***
	SessionKey *string `json:"SessionKey,omitempty" xml:"SessionKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// main
	SessionTarget *string `json:"SessionTarget,omitempty" xml:"SessionTarget,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// now
	WakeMode *string `json:"WakeMode,omitempty" xml:"WakeMode,omitempty"`
}

func (s CreatePolarClawCronJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawCronJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePolarClawCronJobShrinkRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *CreatePolarClawCronJobShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreatePolarClawCronJobShrinkRequest) GetDeleteAfterRun() *bool {
	return s.DeleteAfterRun
}

func (s *CreatePolarClawCronJobShrinkRequest) GetDeliveryShrink() *string {
	return s.DeliveryShrink
}

func (s *CreatePolarClawCronJobShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePolarClawCronJobShrinkRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreatePolarClawCronJobShrinkRequest) GetFailureAlertShrink() *string {
	return s.FailureAlertShrink
}

func (s *CreatePolarClawCronJobShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreatePolarClawCronJobShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *CreatePolarClawCronJobShrinkRequest) GetRestart() *bool {
	return s.Restart
}

func (s *CreatePolarClawCronJobShrinkRequest) GetRunImmediately() *bool {
	return s.RunImmediately
}

func (s *CreatePolarClawCronJobShrinkRequest) GetScheduleShrink() *string {
	return s.ScheduleShrink
}

func (s *CreatePolarClawCronJobShrinkRequest) GetSessionKey() *string {
	return s.SessionKey
}

func (s *CreatePolarClawCronJobShrinkRequest) GetSessionTarget() *string {
	return s.SessionTarget
}

func (s *CreatePolarClawCronJobShrinkRequest) GetWakeMode() *string {
	return s.WakeMode
}

func (s *CreatePolarClawCronJobShrinkRequest) SetAgentId(v string) *CreatePolarClawCronJobShrinkRequest {
	s.AgentId = &v
	return s
}

func (s *CreatePolarClawCronJobShrinkRequest) SetApplicationId(v string) *CreatePolarClawCronJobShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *CreatePolarClawCronJobShrinkRequest) SetDeleteAfterRun(v bool) *CreatePolarClawCronJobShrinkRequest {
	s.DeleteAfterRun = &v
	return s
}

func (s *CreatePolarClawCronJobShrinkRequest) SetDeliveryShrink(v string) *CreatePolarClawCronJobShrinkRequest {
	s.DeliveryShrink = &v
	return s
}

func (s *CreatePolarClawCronJobShrinkRequest) SetDescription(v string) *CreatePolarClawCronJobShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreatePolarClawCronJobShrinkRequest) SetEnabled(v bool) *CreatePolarClawCronJobShrinkRequest {
	s.Enabled = &v
	return s
}

func (s *CreatePolarClawCronJobShrinkRequest) SetFailureAlertShrink(v string) *CreatePolarClawCronJobShrinkRequest {
	s.FailureAlertShrink = &v
	return s
}

func (s *CreatePolarClawCronJobShrinkRequest) SetName(v string) *CreatePolarClawCronJobShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreatePolarClawCronJobShrinkRequest) SetPayloadShrink(v string) *CreatePolarClawCronJobShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *CreatePolarClawCronJobShrinkRequest) SetRestart(v bool) *CreatePolarClawCronJobShrinkRequest {
	s.Restart = &v
	return s
}

func (s *CreatePolarClawCronJobShrinkRequest) SetRunImmediately(v bool) *CreatePolarClawCronJobShrinkRequest {
	s.RunImmediately = &v
	return s
}

func (s *CreatePolarClawCronJobShrinkRequest) SetScheduleShrink(v string) *CreatePolarClawCronJobShrinkRequest {
	s.ScheduleShrink = &v
	return s
}

func (s *CreatePolarClawCronJobShrinkRequest) SetSessionKey(v string) *CreatePolarClawCronJobShrinkRequest {
	s.SessionKey = &v
	return s
}

func (s *CreatePolarClawCronJobShrinkRequest) SetSessionTarget(v string) *CreatePolarClawCronJobShrinkRequest {
	s.SessionTarget = &v
	return s
}

func (s *CreatePolarClawCronJobShrinkRequest) SetWakeMode(v string) *CreatePolarClawCronJobShrinkRequest {
	s.WakeMode = &v
	return s
}

func (s *CreatePolarClawCronJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
