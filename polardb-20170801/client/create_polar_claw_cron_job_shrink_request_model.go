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
	// The ID of the agent that executes the task.
	//
	// example:
	//
	// main
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// Specifies whether to automatically delete the job after its first execution. This is useful for one-time tasks. Default: `false`.
	//
	// example:
	//
	// false
	DeleteAfterRun *bool `json:"DeleteAfterRun,omitempty" xml:"DeleteAfterRun,omitempty"`
	// The configuration for delivering task execution results.
	//
	// example:
	//
	// {"Mode":"announce","Channel":"telegram"}
	DeliveryShrink *string `json:"Delivery,omitempty" xml:"Delivery,omitempty"`
	// A description of the task.
	//
	// example:
	//
	// Daily report generation
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether the cron job is enabled. Default: `true`.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The failure alert configuration.
	//
	// example:
	//
	// {"After":3,"Channel":"telegram"}
	FailureAlertShrink *string `json:"FailureAlert,omitempty" xml:"FailureAlert,omitempty"`
	// The unique name of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// daily-report
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The execution payload configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Kind":"agentTurn","Message":"Generate the daily report."}
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// Specifies whether to restart the gateway upon job creation. Default: `true`.
	//
	// example:
	//
	// true
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
	// Specifies whether to run the job once immediately upon creation. Default: `false`.
	//
	// example:
	//
	// false
	RunImmediately *bool `json:"RunImmediately,omitempty" xml:"RunImmediately,omitempty"`
	// The schedule configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Kind":"cron","Expr":"0 9 	- 	- *","Tz":"Asia/Shanghai"}
	ScheduleShrink *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The session routing key, which determines the conversation session for the task.
	//
	// example:
	//
	// agent:main:feishu:direct:***
	SessionKey *string `json:"SessionKey,omitempty" xml:"SessionKey,omitempty"`
	// The session target. Valid values are `main`, `isolated`, and `current`.
	//
	// This parameter is required.
	//
	// example:
	//
	// main
	SessionTarget *string `json:"SessionTarget,omitempty" xml:"SessionTarget,omitempty"`
	// The wake mode for the agent. Valid values are `now` and `next-heartbeat`.
	//
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
