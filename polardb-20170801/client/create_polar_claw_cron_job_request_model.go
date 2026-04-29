// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolarClawCronJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *CreatePolarClawCronJobRequest
	GetAgentId() *string
	SetApplicationId(v string) *CreatePolarClawCronJobRequest
	GetApplicationId() *string
	SetDeleteAfterRun(v bool) *CreatePolarClawCronJobRequest
	GetDeleteAfterRun() *bool
	SetDelivery(v *CreatePolarClawCronJobRequestDelivery) *CreatePolarClawCronJobRequest
	GetDelivery() *CreatePolarClawCronJobRequestDelivery
	SetDescription(v string) *CreatePolarClawCronJobRequest
	GetDescription() *string
	SetEnabled(v bool) *CreatePolarClawCronJobRequest
	GetEnabled() *bool
	SetFailureAlert(v *CreatePolarClawCronJobRequestFailureAlert) *CreatePolarClawCronJobRequest
	GetFailureAlert() *CreatePolarClawCronJobRequestFailureAlert
	SetName(v string) *CreatePolarClawCronJobRequest
	GetName() *string
	SetPayload(v *CreatePolarClawCronJobRequestPayload) *CreatePolarClawCronJobRequest
	GetPayload() *CreatePolarClawCronJobRequestPayload
	SetRestart(v bool) *CreatePolarClawCronJobRequest
	GetRestart() *bool
	SetRunImmediately(v bool) *CreatePolarClawCronJobRequest
	GetRunImmediately() *bool
	SetSchedule(v *CreatePolarClawCronJobRequestSchedule) *CreatePolarClawCronJobRequest
	GetSchedule() *CreatePolarClawCronJobRequestSchedule
	SetSessionKey(v string) *CreatePolarClawCronJobRequest
	GetSessionKey() *string
	SetSessionTarget(v string) *CreatePolarClawCronJobRequest
	GetSessionTarget() *string
	SetWakeMode(v string) *CreatePolarClawCronJobRequest
	GetWakeMode() *string
}

type CreatePolarClawCronJobRequest struct {
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
	Delivery *CreatePolarClawCronJobRequestDelivery `json:"Delivery,omitempty" xml:"Delivery,omitempty" type:"Struct"`
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
	FailureAlert *CreatePolarClawCronJobRequestFailureAlert `json:"FailureAlert,omitempty" xml:"FailureAlert,omitempty" type:"Struct"`
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
	Payload *CreatePolarClawCronJobRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
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
	Schedule *CreatePolarClawCronJobRequestSchedule `json:"Schedule,omitempty" xml:"Schedule,omitempty" type:"Struct"`
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

func (s CreatePolarClawCronJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawCronJobRequest) GoString() string {
	return s.String()
}

func (s *CreatePolarClawCronJobRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *CreatePolarClawCronJobRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreatePolarClawCronJobRequest) GetDeleteAfterRun() *bool {
	return s.DeleteAfterRun
}

func (s *CreatePolarClawCronJobRequest) GetDelivery() *CreatePolarClawCronJobRequestDelivery {
	return s.Delivery
}

func (s *CreatePolarClawCronJobRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePolarClawCronJobRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreatePolarClawCronJobRequest) GetFailureAlert() *CreatePolarClawCronJobRequestFailureAlert {
	return s.FailureAlert
}

func (s *CreatePolarClawCronJobRequest) GetName() *string {
	return s.Name
}

func (s *CreatePolarClawCronJobRequest) GetPayload() *CreatePolarClawCronJobRequestPayload {
	return s.Payload
}

func (s *CreatePolarClawCronJobRequest) GetRestart() *bool {
	return s.Restart
}

func (s *CreatePolarClawCronJobRequest) GetRunImmediately() *bool {
	return s.RunImmediately
}

func (s *CreatePolarClawCronJobRequest) GetSchedule() *CreatePolarClawCronJobRequestSchedule {
	return s.Schedule
}

func (s *CreatePolarClawCronJobRequest) GetSessionKey() *string {
	return s.SessionKey
}

func (s *CreatePolarClawCronJobRequest) GetSessionTarget() *string {
	return s.SessionTarget
}

func (s *CreatePolarClawCronJobRequest) GetWakeMode() *string {
	return s.WakeMode
}

func (s *CreatePolarClawCronJobRequest) SetAgentId(v string) *CreatePolarClawCronJobRequest {
	s.AgentId = &v
	return s
}

func (s *CreatePolarClawCronJobRequest) SetApplicationId(v string) *CreatePolarClawCronJobRequest {
	s.ApplicationId = &v
	return s
}

func (s *CreatePolarClawCronJobRequest) SetDeleteAfterRun(v bool) *CreatePolarClawCronJobRequest {
	s.DeleteAfterRun = &v
	return s
}

func (s *CreatePolarClawCronJobRequest) SetDelivery(v *CreatePolarClawCronJobRequestDelivery) *CreatePolarClawCronJobRequest {
	s.Delivery = v
	return s
}

func (s *CreatePolarClawCronJobRequest) SetDescription(v string) *CreatePolarClawCronJobRequest {
	s.Description = &v
	return s
}

func (s *CreatePolarClawCronJobRequest) SetEnabled(v bool) *CreatePolarClawCronJobRequest {
	s.Enabled = &v
	return s
}

func (s *CreatePolarClawCronJobRequest) SetFailureAlert(v *CreatePolarClawCronJobRequestFailureAlert) *CreatePolarClawCronJobRequest {
	s.FailureAlert = v
	return s
}

func (s *CreatePolarClawCronJobRequest) SetName(v string) *CreatePolarClawCronJobRequest {
	s.Name = &v
	return s
}

func (s *CreatePolarClawCronJobRequest) SetPayload(v *CreatePolarClawCronJobRequestPayload) *CreatePolarClawCronJobRequest {
	s.Payload = v
	return s
}

func (s *CreatePolarClawCronJobRequest) SetRestart(v bool) *CreatePolarClawCronJobRequest {
	s.Restart = &v
	return s
}

func (s *CreatePolarClawCronJobRequest) SetRunImmediately(v bool) *CreatePolarClawCronJobRequest {
	s.RunImmediately = &v
	return s
}

func (s *CreatePolarClawCronJobRequest) SetSchedule(v *CreatePolarClawCronJobRequestSchedule) *CreatePolarClawCronJobRequest {
	s.Schedule = v
	return s
}

func (s *CreatePolarClawCronJobRequest) SetSessionKey(v string) *CreatePolarClawCronJobRequest {
	s.SessionKey = &v
	return s
}

func (s *CreatePolarClawCronJobRequest) SetSessionTarget(v string) *CreatePolarClawCronJobRequest {
	s.SessionTarget = &v
	return s
}

func (s *CreatePolarClawCronJobRequest) SetWakeMode(v string) *CreatePolarClawCronJobRequest {
	s.WakeMode = &v
	return s
}

func (s *CreatePolarClawCronJobRequest) Validate() error {
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

type CreatePolarClawCronJobRequestDelivery struct {
	// example:
	//
	// default
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// false
	BestEffort *bool `json:"BestEffort,omitempty" xml:"BestEffort,omitempty"`
	// example:
	//
	// feishu
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// example:
	//
	// announce
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// ou_***
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s CreatePolarClawCronJobRequestDelivery) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawCronJobRequestDelivery) GoString() string {
	return s.String()
}

func (s *CreatePolarClawCronJobRequestDelivery) GetAccountId() *string {
	return s.AccountId
}

func (s *CreatePolarClawCronJobRequestDelivery) GetBestEffort() *bool {
	return s.BestEffort
}

func (s *CreatePolarClawCronJobRequestDelivery) GetChannel() *string {
	return s.Channel
}

func (s *CreatePolarClawCronJobRequestDelivery) GetMode() *string {
	return s.Mode
}

func (s *CreatePolarClawCronJobRequestDelivery) GetTo() *string {
	return s.To
}

func (s *CreatePolarClawCronJobRequestDelivery) SetAccountId(v string) *CreatePolarClawCronJobRequestDelivery {
	s.AccountId = &v
	return s
}

func (s *CreatePolarClawCronJobRequestDelivery) SetBestEffort(v bool) *CreatePolarClawCronJobRequestDelivery {
	s.BestEffort = &v
	return s
}

func (s *CreatePolarClawCronJobRequestDelivery) SetChannel(v string) *CreatePolarClawCronJobRequestDelivery {
	s.Channel = &v
	return s
}

func (s *CreatePolarClawCronJobRequestDelivery) SetMode(v string) *CreatePolarClawCronJobRequestDelivery {
	s.Mode = &v
	return s
}

func (s *CreatePolarClawCronJobRequestDelivery) SetTo(v string) *CreatePolarClawCronJobRequestDelivery {
	s.To = &v
	return s
}

func (s *CreatePolarClawCronJobRequestDelivery) Validate() error {
	return dara.Validate(s)
}

type CreatePolarClawCronJobRequestFailureAlert struct {
	// example:
	//
	// default
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// 3
	After *int32 `json:"After,omitempty" xml:"After,omitempty"`
	// example:
	//
	// feishu
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
	// ou_***
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s CreatePolarClawCronJobRequestFailureAlert) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawCronJobRequestFailureAlert) GoString() string {
	return s.String()
}

func (s *CreatePolarClawCronJobRequestFailureAlert) GetAccountId() *string {
	return s.AccountId
}

func (s *CreatePolarClawCronJobRequestFailureAlert) GetAfter() *int32 {
	return s.After
}

func (s *CreatePolarClawCronJobRequestFailureAlert) GetChannel() *string {
	return s.Channel
}

func (s *CreatePolarClawCronJobRequestFailureAlert) GetCooldownMs() *int32 {
	return s.CooldownMs
}

func (s *CreatePolarClawCronJobRequestFailureAlert) GetMode() *string {
	return s.Mode
}

func (s *CreatePolarClawCronJobRequestFailureAlert) GetTo() *string {
	return s.To
}

func (s *CreatePolarClawCronJobRequestFailureAlert) SetAccountId(v string) *CreatePolarClawCronJobRequestFailureAlert {
	s.AccountId = &v
	return s
}

func (s *CreatePolarClawCronJobRequestFailureAlert) SetAfter(v int32) *CreatePolarClawCronJobRequestFailureAlert {
	s.After = &v
	return s
}

func (s *CreatePolarClawCronJobRequestFailureAlert) SetChannel(v string) *CreatePolarClawCronJobRequestFailureAlert {
	s.Channel = &v
	return s
}

func (s *CreatePolarClawCronJobRequestFailureAlert) SetCooldownMs(v int32) *CreatePolarClawCronJobRequestFailureAlert {
	s.CooldownMs = &v
	return s
}

func (s *CreatePolarClawCronJobRequestFailureAlert) SetMode(v string) *CreatePolarClawCronJobRequestFailureAlert {
	s.Mode = &v
	return s
}

func (s *CreatePolarClawCronJobRequestFailureAlert) SetTo(v string) *CreatePolarClawCronJobRequestFailureAlert {
	s.To = &v
	return s
}

func (s *CreatePolarClawCronJobRequestFailureAlert) Validate() error {
	return dara.Validate(s)
}

type CreatePolarClawCronJobRequestPayload struct {
	// example:
	//
	// false
	BestEffortDeliver *bool `json:"BestEffortDeliver,omitempty" xml:"BestEffortDeliver,omitempty"`
	// example:
	//
	// feishu
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// example:
	//
	// true
	Deliver   *bool     `json:"Deliver,omitempty" xml:"Deliver,omitempty"`
	Fallbacks []*string `json:"Fallbacks,omitempty" xml:"Fallbacks,omitempty" type:"Repeated"`
	// example:
	//
	// systemEvent
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// example:
	//
	// false
	LightContext *bool `json:"LightContext,omitempty" xml:"LightContext,omitempty"`
	// example:
	//
	// Generate the daily report.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// bailian/qwen3.5-plus
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// Generate the daily report.
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// xhigh
	Thinking *string `json:"Thinking,omitempty" xml:"Thinking,omitempty"`
	// example:
	//
	// 10
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	// example:
	//
	// ou_***
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s CreatePolarClawCronJobRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawCronJobRequestPayload) GoString() string {
	return s.String()
}

func (s *CreatePolarClawCronJobRequestPayload) GetBestEffortDeliver() *bool {
	return s.BestEffortDeliver
}

func (s *CreatePolarClawCronJobRequestPayload) GetChannel() *string {
	return s.Channel
}

func (s *CreatePolarClawCronJobRequestPayload) GetDeliver() *bool {
	return s.Deliver
}

func (s *CreatePolarClawCronJobRequestPayload) GetFallbacks() []*string {
	return s.Fallbacks
}

func (s *CreatePolarClawCronJobRequestPayload) GetKind() *string {
	return s.Kind
}

func (s *CreatePolarClawCronJobRequestPayload) GetLightContext() *bool {
	return s.LightContext
}

func (s *CreatePolarClawCronJobRequestPayload) GetMessage() *string {
	return s.Message
}

func (s *CreatePolarClawCronJobRequestPayload) GetModel() *string {
	return s.Model
}

func (s *CreatePolarClawCronJobRequestPayload) GetText() *string {
	return s.Text
}

func (s *CreatePolarClawCronJobRequestPayload) GetThinking() *string {
	return s.Thinking
}

func (s *CreatePolarClawCronJobRequestPayload) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *CreatePolarClawCronJobRequestPayload) GetTo() *string {
	return s.To
}

func (s *CreatePolarClawCronJobRequestPayload) SetBestEffortDeliver(v bool) *CreatePolarClawCronJobRequestPayload {
	s.BestEffortDeliver = &v
	return s
}

func (s *CreatePolarClawCronJobRequestPayload) SetChannel(v string) *CreatePolarClawCronJobRequestPayload {
	s.Channel = &v
	return s
}

func (s *CreatePolarClawCronJobRequestPayload) SetDeliver(v bool) *CreatePolarClawCronJobRequestPayload {
	s.Deliver = &v
	return s
}

func (s *CreatePolarClawCronJobRequestPayload) SetFallbacks(v []*string) *CreatePolarClawCronJobRequestPayload {
	s.Fallbacks = v
	return s
}

func (s *CreatePolarClawCronJobRequestPayload) SetKind(v string) *CreatePolarClawCronJobRequestPayload {
	s.Kind = &v
	return s
}

func (s *CreatePolarClawCronJobRequestPayload) SetLightContext(v bool) *CreatePolarClawCronJobRequestPayload {
	s.LightContext = &v
	return s
}

func (s *CreatePolarClawCronJobRequestPayload) SetMessage(v string) *CreatePolarClawCronJobRequestPayload {
	s.Message = &v
	return s
}

func (s *CreatePolarClawCronJobRequestPayload) SetModel(v string) *CreatePolarClawCronJobRequestPayload {
	s.Model = &v
	return s
}

func (s *CreatePolarClawCronJobRequestPayload) SetText(v string) *CreatePolarClawCronJobRequestPayload {
	s.Text = &v
	return s
}

func (s *CreatePolarClawCronJobRequestPayload) SetThinking(v string) *CreatePolarClawCronJobRequestPayload {
	s.Thinking = &v
	return s
}

func (s *CreatePolarClawCronJobRequestPayload) SetTimeoutSeconds(v int32) *CreatePolarClawCronJobRequestPayload {
	s.TimeoutSeconds = &v
	return s
}

func (s *CreatePolarClawCronJobRequestPayload) SetTo(v string) *CreatePolarClawCronJobRequestPayload {
	s.To = &v
	return s
}

func (s *CreatePolarClawCronJobRequestPayload) Validate() error {
	return dara.Validate(s)
}

type CreatePolarClawCronJobRequestSchedule struct {
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

func (s CreatePolarClawCronJobRequestSchedule) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawCronJobRequestSchedule) GoString() string {
	return s.String()
}

func (s *CreatePolarClawCronJobRequestSchedule) GetAnchorMs() *int64 {
	return s.AnchorMs
}

func (s *CreatePolarClawCronJobRequestSchedule) GetAt() *string {
	return s.At
}

func (s *CreatePolarClawCronJobRequestSchedule) GetEveryMs() *int64 {
	return s.EveryMs
}

func (s *CreatePolarClawCronJobRequestSchedule) GetExpr() *string {
	return s.Expr
}

func (s *CreatePolarClawCronJobRequestSchedule) GetKind() *string {
	return s.Kind
}

func (s *CreatePolarClawCronJobRequestSchedule) GetStaggerMs() *int32 {
	return s.StaggerMs
}

func (s *CreatePolarClawCronJobRequestSchedule) GetTz() *string {
	return s.Tz
}

func (s *CreatePolarClawCronJobRequestSchedule) SetAnchorMs(v int64) *CreatePolarClawCronJobRequestSchedule {
	s.AnchorMs = &v
	return s
}

func (s *CreatePolarClawCronJobRequestSchedule) SetAt(v string) *CreatePolarClawCronJobRequestSchedule {
	s.At = &v
	return s
}

func (s *CreatePolarClawCronJobRequestSchedule) SetEveryMs(v int64) *CreatePolarClawCronJobRequestSchedule {
	s.EveryMs = &v
	return s
}

func (s *CreatePolarClawCronJobRequestSchedule) SetExpr(v string) *CreatePolarClawCronJobRequestSchedule {
	s.Expr = &v
	return s
}

func (s *CreatePolarClawCronJobRequestSchedule) SetKind(v string) *CreatePolarClawCronJobRequestSchedule {
	s.Kind = &v
	return s
}

func (s *CreatePolarClawCronJobRequestSchedule) SetStaggerMs(v int32) *CreatePolarClawCronJobRequestSchedule {
	s.StaggerMs = &v
	return s
}

func (s *CreatePolarClawCronJobRequestSchedule) SetTz(v string) *CreatePolarClawCronJobRequestSchedule {
	s.Tz = &v
	return s
}

func (s *CreatePolarClawCronJobRequestSchedule) Validate() error {
	return dara.Validate(s)
}
