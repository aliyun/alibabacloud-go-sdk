// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolarClawCronJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *UpdatePolarClawCronJobRequest
	GetAgentId() *string
	SetApplicationId(v string) *UpdatePolarClawCronJobRequest
	GetApplicationId() *string
	SetDeleteAfterRun(v bool) *UpdatePolarClawCronJobRequest
	GetDeleteAfterRun() *bool
	SetDelivery(v *UpdatePolarClawCronJobRequestDelivery) *UpdatePolarClawCronJobRequest
	GetDelivery() *UpdatePolarClawCronJobRequestDelivery
	SetDescription(v string) *UpdatePolarClawCronJobRequest
	GetDescription() *string
	SetEnabled(v bool) *UpdatePolarClawCronJobRequest
	GetEnabled() *bool
	SetFailureAlert(v *UpdatePolarClawCronJobRequestFailureAlert) *UpdatePolarClawCronJobRequest
	GetFailureAlert() *UpdatePolarClawCronJobRequestFailureAlert
	SetJobId(v string) *UpdatePolarClawCronJobRequest
	GetJobId() *string
	SetName(v string) *UpdatePolarClawCronJobRequest
	GetName() *string
	SetPayload(v *UpdatePolarClawCronJobRequestPayload) *UpdatePolarClawCronJobRequest
	GetPayload() *UpdatePolarClawCronJobRequestPayload
	SetRestart(v bool) *UpdatePolarClawCronJobRequest
	GetRestart() *bool
	SetSchedule(v *UpdatePolarClawCronJobRequestSchedule) *UpdatePolarClawCronJobRequest
	GetSchedule() *UpdatePolarClawCronJobRequestSchedule
	SetSessionKey(v string) *UpdatePolarClawCronJobRequest
	GetSessionKey() *string
	SetSessionTarget(v string) *UpdatePolarClawCronJobRequest
	GetSessionTarget() *string
	SetWakeMode(v string) *UpdatePolarClawCronJobRequest
	GetWakeMode() *string
}

type UpdatePolarClawCronJobRequest struct {
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
	Delivery *UpdatePolarClawCronJobRequestDelivery `json:"Delivery,omitempty" xml:"Delivery,omitempty" type:"Struct"`
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
	FailureAlert *UpdatePolarClawCronJobRequestFailureAlert `json:"FailureAlert,omitempty" xml:"FailureAlert,omitempty" type:"Struct"`
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
	Payload *UpdatePolarClawCronJobRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// true
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
	// example:
	//
	// {"Kind":"cron","Expr":"0 12 	- 	- *","Tz":"America/New_York"}
	Schedule *UpdatePolarClawCronJobRequestSchedule `json:"Schedule,omitempty" xml:"Schedule,omitempty" type:"Struct"`
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

func (s UpdatePolarClawCronJobRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawCronJobRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawCronJobRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdatePolarClawCronJobRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdatePolarClawCronJobRequest) GetDeleteAfterRun() *bool {
	return s.DeleteAfterRun
}

func (s *UpdatePolarClawCronJobRequest) GetDelivery() *UpdatePolarClawCronJobRequestDelivery {
	return s.Delivery
}

func (s *UpdatePolarClawCronJobRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdatePolarClawCronJobRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdatePolarClawCronJobRequest) GetFailureAlert() *UpdatePolarClawCronJobRequestFailureAlert {
	return s.FailureAlert
}

func (s *UpdatePolarClawCronJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *UpdatePolarClawCronJobRequest) GetName() *string {
	return s.Name
}

func (s *UpdatePolarClawCronJobRequest) GetPayload() *UpdatePolarClawCronJobRequestPayload {
	return s.Payload
}

func (s *UpdatePolarClawCronJobRequest) GetRestart() *bool {
	return s.Restart
}

func (s *UpdatePolarClawCronJobRequest) GetSchedule() *UpdatePolarClawCronJobRequestSchedule {
	return s.Schedule
}

func (s *UpdatePolarClawCronJobRequest) GetSessionKey() *string {
	return s.SessionKey
}

func (s *UpdatePolarClawCronJobRequest) GetSessionTarget() *string {
	return s.SessionTarget
}

func (s *UpdatePolarClawCronJobRequest) GetWakeMode() *string {
	return s.WakeMode
}

func (s *UpdatePolarClawCronJobRequest) SetAgentId(v string) *UpdatePolarClawCronJobRequest {
	s.AgentId = &v
	return s
}

func (s *UpdatePolarClawCronJobRequest) SetApplicationId(v string) *UpdatePolarClawCronJobRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdatePolarClawCronJobRequest) SetDeleteAfterRun(v bool) *UpdatePolarClawCronJobRequest {
	s.DeleteAfterRun = &v
	return s
}

func (s *UpdatePolarClawCronJobRequest) SetDelivery(v *UpdatePolarClawCronJobRequestDelivery) *UpdatePolarClawCronJobRequest {
	s.Delivery = v
	return s
}

func (s *UpdatePolarClawCronJobRequest) SetDescription(v string) *UpdatePolarClawCronJobRequest {
	s.Description = &v
	return s
}

func (s *UpdatePolarClawCronJobRequest) SetEnabled(v bool) *UpdatePolarClawCronJobRequest {
	s.Enabled = &v
	return s
}

func (s *UpdatePolarClawCronJobRequest) SetFailureAlert(v *UpdatePolarClawCronJobRequestFailureAlert) *UpdatePolarClawCronJobRequest {
	s.FailureAlert = v
	return s
}

func (s *UpdatePolarClawCronJobRequest) SetJobId(v string) *UpdatePolarClawCronJobRequest {
	s.JobId = &v
	return s
}

func (s *UpdatePolarClawCronJobRequest) SetName(v string) *UpdatePolarClawCronJobRequest {
	s.Name = &v
	return s
}

func (s *UpdatePolarClawCronJobRequest) SetPayload(v *UpdatePolarClawCronJobRequestPayload) *UpdatePolarClawCronJobRequest {
	s.Payload = v
	return s
}

func (s *UpdatePolarClawCronJobRequest) SetRestart(v bool) *UpdatePolarClawCronJobRequest {
	s.Restart = &v
	return s
}

func (s *UpdatePolarClawCronJobRequest) SetSchedule(v *UpdatePolarClawCronJobRequestSchedule) *UpdatePolarClawCronJobRequest {
	s.Schedule = v
	return s
}

func (s *UpdatePolarClawCronJobRequest) SetSessionKey(v string) *UpdatePolarClawCronJobRequest {
	s.SessionKey = &v
	return s
}

func (s *UpdatePolarClawCronJobRequest) SetSessionTarget(v string) *UpdatePolarClawCronJobRequest {
	s.SessionTarget = &v
	return s
}

func (s *UpdatePolarClawCronJobRequest) SetWakeMode(v string) *UpdatePolarClawCronJobRequest {
	s.WakeMode = &v
	return s
}

func (s *UpdatePolarClawCronJobRequest) Validate() error {
	if s.Delivery != nil {
		if err := s.Delivery.Validate(); err != nil {
			return err
		}
	}
	if s.FailureAlert != nil {
		if err := s.FailureAlert.Validate(); err != nil {
			return err
		}
	}
	if s.Payload != nil {
		if err := s.Payload.Validate(); err != nil {
			return err
		}
	}
	if s.Schedule != nil {
		if err := s.Schedule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePolarClawCronJobRequestDelivery struct {
	// example:
	//
	// accountId123
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// false
	BestEffort *bool `json:"BestEffort,omitempty" xml:"BestEffort,omitempty"`
	// example:
	//
	// telegram
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// example:
	//
	// announce
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// https://example.com/webhook
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s UpdatePolarClawCronJobRequestDelivery) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawCronJobRequestDelivery) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawCronJobRequestDelivery) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdatePolarClawCronJobRequestDelivery) GetBestEffort() *bool {
	return s.BestEffort
}

func (s *UpdatePolarClawCronJobRequestDelivery) GetChannel() *string {
	return s.Channel
}

func (s *UpdatePolarClawCronJobRequestDelivery) GetMode() *string {
	return s.Mode
}

func (s *UpdatePolarClawCronJobRequestDelivery) GetTo() *string {
	return s.To
}

func (s *UpdatePolarClawCronJobRequestDelivery) SetAccountId(v string) *UpdatePolarClawCronJobRequestDelivery {
	s.AccountId = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestDelivery) SetBestEffort(v bool) *UpdatePolarClawCronJobRequestDelivery {
	s.BestEffort = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestDelivery) SetChannel(v string) *UpdatePolarClawCronJobRequestDelivery {
	s.Channel = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestDelivery) SetMode(v string) *UpdatePolarClawCronJobRequestDelivery {
	s.Mode = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestDelivery) SetTo(v string) *UpdatePolarClawCronJobRequestDelivery {
	s.To = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestDelivery) Validate() error {
	return dara.Validate(s)
}

type UpdatePolarClawCronJobRequestFailureAlert struct {
	// example:
	//
	// accountId123
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// 3
	After *int32 `json:"After,omitempty" xml:"After,omitempty"`
	// example:
	//
	// email
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// example:
	//
	// 5000
	CooldownMs *int32 `json:"CooldownMs,omitempty" xml:"CooldownMs,omitempty"`
	// example:
	//
	// announce
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// user@example.com
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s UpdatePolarClawCronJobRequestFailureAlert) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawCronJobRequestFailureAlert) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawCronJobRequestFailureAlert) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdatePolarClawCronJobRequestFailureAlert) GetAfter() *int32 {
	return s.After
}

func (s *UpdatePolarClawCronJobRequestFailureAlert) GetChannel() *string {
	return s.Channel
}

func (s *UpdatePolarClawCronJobRequestFailureAlert) GetCooldownMs() *int32 {
	return s.CooldownMs
}

func (s *UpdatePolarClawCronJobRequestFailureAlert) GetMode() *string {
	return s.Mode
}

func (s *UpdatePolarClawCronJobRequestFailureAlert) GetTo() *string {
	return s.To
}

func (s *UpdatePolarClawCronJobRequestFailureAlert) SetAccountId(v string) *UpdatePolarClawCronJobRequestFailureAlert {
	s.AccountId = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestFailureAlert) SetAfter(v int32) *UpdatePolarClawCronJobRequestFailureAlert {
	s.After = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestFailureAlert) SetChannel(v string) *UpdatePolarClawCronJobRequestFailureAlert {
	s.Channel = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestFailureAlert) SetCooldownMs(v int32) *UpdatePolarClawCronJobRequestFailureAlert {
	s.CooldownMs = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestFailureAlert) SetMode(v string) *UpdatePolarClawCronJobRequestFailureAlert {
	s.Mode = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestFailureAlert) SetTo(v string) *UpdatePolarClawCronJobRequestFailureAlert {
	s.To = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestFailureAlert) Validate() error {
	return dara.Validate(s)
}

type UpdatePolarClawCronJobRequestPayload struct {
	BestEffortDeliver *bool `json:"BestEffortDeliver,omitempty" xml:"BestEffortDeliver,omitempty"`
	// example:
	//
	// telegram
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// example:
	//
	// false
	Deliver   *bool     `json:"Deliver,omitempty" xml:"Deliver,omitempty"`
	Fallbacks []*string `json:"Fallbacks,omitempty" xml:"Fallbacks,omitempty" type:"Repeated"`
	// example:
	//
	// agentTurn
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// example:
	//
	// false
	LightContext *bool `json:"LightContext,omitempty" xml:"LightContext,omitempty"`
	// example:
	//
	// Generate the daily report and send it to the team.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// anthropic/sonnet-4.6
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// Send a reminder to the user.
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// medium
	Thinking *string `json:"Thinking,omitempty" xml:"Thinking,omitempty"`
	// example:
	//
	// 10
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	// example:
	//
	// team
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s UpdatePolarClawCronJobRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawCronJobRequestPayload) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawCronJobRequestPayload) GetBestEffortDeliver() *bool {
	return s.BestEffortDeliver
}

func (s *UpdatePolarClawCronJobRequestPayload) GetChannel() *string {
	return s.Channel
}

func (s *UpdatePolarClawCronJobRequestPayload) GetDeliver() *bool {
	return s.Deliver
}

func (s *UpdatePolarClawCronJobRequestPayload) GetFallbacks() []*string {
	return s.Fallbacks
}

func (s *UpdatePolarClawCronJobRequestPayload) GetKind() *string {
	return s.Kind
}

func (s *UpdatePolarClawCronJobRequestPayload) GetLightContext() *bool {
	return s.LightContext
}

func (s *UpdatePolarClawCronJobRequestPayload) GetMessage() *string {
	return s.Message
}

func (s *UpdatePolarClawCronJobRequestPayload) GetModel() *string {
	return s.Model
}

func (s *UpdatePolarClawCronJobRequestPayload) GetText() *string {
	return s.Text
}

func (s *UpdatePolarClawCronJobRequestPayload) GetThinking() *string {
	return s.Thinking
}

func (s *UpdatePolarClawCronJobRequestPayload) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *UpdatePolarClawCronJobRequestPayload) GetTo() *string {
	return s.To
}

func (s *UpdatePolarClawCronJobRequestPayload) SetBestEffortDeliver(v bool) *UpdatePolarClawCronJobRequestPayload {
	s.BestEffortDeliver = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestPayload) SetChannel(v string) *UpdatePolarClawCronJobRequestPayload {
	s.Channel = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestPayload) SetDeliver(v bool) *UpdatePolarClawCronJobRequestPayload {
	s.Deliver = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestPayload) SetFallbacks(v []*string) *UpdatePolarClawCronJobRequestPayload {
	s.Fallbacks = v
	return s
}

func (s *UpdatePolarClawCronJobRequestPayload) SetKind(v string) *UpdatePolarClawCronJobRequestPayload {
	s.Kind = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestPayload) SetLightContext(v bool) *UpdatePolarClawCronJobRequestPayload {
	s.LightContext = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestPayload) SetMessage(v string) *UpdatePolarClawCronJobRequestPayload {
	s.Message = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestPayload) SetModel(v string) *UpdatePolarClawCronJobRequestPayload {
	s.Model = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestPayload) SetText(v string) *UpdatePolarClawCronJobRequestPayload {
	s.Text = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestPayload) SetThinking(v string) *UpdatePolarClawCronJobRequestPayload {
	s.Thinking = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestPayload) SetTimeoutSeconds(v int32) *UpdatePolarClawCronJobRequestPayload {
	s.TimeoutSeconds = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestPayload) SetTo(v string) *UpdatePolarClawCronJobRequestPayload {
	s.To = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestPayload) Validate() error {
	return dara.Validate(s)
}

type UpdatePolarClawCronJobRequestSchedule struct {
	// example:
	//
	// 1777370572518
	AnchorMs *int64 `json:"AnchorMs,omitempty" xml:"AnchorMs,omitempty"`
	// example:
	//
	// 2026-04-10T09:00:00+08:00
	At *string `json:"At,omitempty" xml:"At,omitempty"`
	// example:
	//
	// 100000
	EveryMs *int64 `json:"EveryMs,omitempty" xml:"EveryMs,omitempty"`
	// example:
	//
	// 0 9 	- 	- *
	Expr *string `json:"Expr,omitempty" xml:"Expr,omitempty"`
	// example:
	//
	// cron
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// example:
	//
	// 0
	StaggerMs *int32 `json:"StaggerMs,omitempty" xml:"StaggerMs,omitempty"`
	// example:
	//
	// Asia/Shanghai
	Tz *string `json:"Tz,omitempty" xml:"Tz,omitempty"`
}

func (s UpdatePolarClawCronJobRequestSchedule) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawCronJobRequestSchedule) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawCronJobRequestSchedule) GetAnchorMs() *int64 {
	return s.AnchorMs
}

func (s *UpdatePolarClawCronJobRequestSchedule) GetAt() *string {
	return s.At
}

func (s *UpdatePolarClawCronJobRequestSchedule) GetEveryMs() *int64 {
	return s.EveryMs
}

func (s *UpdatePolarClawCronJobRequestSchedule) GetExpr() *string {
	return s.Expr
}

func (s *UpdatePolarClawCronJobRequestSchedule) GetKind() *string {
	return s.Kind
}

func (s *UpdatePolarClawCronJobRequestSchedule) GetStaggerMs() *int32 {
	return s.StaggerMs
}

func (s *UpdatePolarClawCronJobRequestSchedule) GetTz() *string {
	return s.Tz
}

func (s *UpdatePolarClawCronJobRequestSchedule) SetAnchorMs(v int64) *UpdatePolarClawCronJobRequestSchedule {
	s.AnchorMs = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestSchedule) SetAt(v string) *UpdatePolarClawCronJobRequestSchedule {
	s.At = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestSchedule) SetEveryMs(v int64) *UpdatePolarClawCronJobRequestSchedule {
	s.EveryMs = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestSchedule) SetExpr(v string) *UpdatePolarClawCronJobRequestSchedule {
	s.Expr = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestSchedule) SetKind(v string) *UpdatePolarClawCronJobRequestSchedule {
	s.Kind = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestSchedule) SetStaggerMs(v int32) *UpdatePolarClawCronJobRequestSchedule {
	s.StaggerMs = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestSchedule) SetTz(v string) *UpdatePolarClawCronJobRequestSchedule {
	s.Tz = &v
	return s
}

func (s *UpdatePolarClawCronJobRequestSchedule) Validate() error {
	return dara.Validate(s)
}
