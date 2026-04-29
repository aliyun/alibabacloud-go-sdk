// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolarClawCronJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *CreatePolarClawCronJobResponseBody
	GetApplicationId() *string
	SetCode(v int32) *CreatePolarClawCronJobResponseBody
	GetCode() *int32
	SetJob(v *CreatePolarClawCronJobResponseBodyJob) *CreatePolarClawCronJobResponseBody
	GetJob() *CreatePolarClawCronJobResponseBodyJob
	SetMessage(v string) *CreatePolarClawCronJobResponseBody
	GetMessage() *string
	SetOk(v bool) *CreatePolarClawCronJobResponseBody
	GetOk() *bool
	SetRanImmediately(v bool) *CreatePolarClawCronJobResponseBody
	GetRanImmediately() *bool
	SetRequestId(v string) *CreatePolarClawCronJobResponseBody
	GetRequestId() *string
}

type CreatePolarClawCronJobResponseBody struct {
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 200
	Code *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Job  *CreatePolarClawCronJobResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Ok *bool `json:"Ok,omitempty" xml:"Ok,omitempty"`
	// example:
	//
	// false
	RanImmediately *bool `json:"RanImmediately,omitempty" xml:"RanImmediately,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6BD9CDE4-5E7B-4BF3-9BB8-83C73E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePolarClawCronJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawCronJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolarClawCronJobResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreatePolarClawCronJobResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreatePolarClawCronJobResponseBody) GetJob() *CreatePolarClawCronJobResponseBodyJob {
	return s.Job
}

func (s *CreatePolarClawCronJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePolarClawCronJobResponseBody) GetOk() *bool {
	return s.Ok
}

func (s *CreatePolarClawCronJobResponseBody) GetRanImmediately() *bool {
	return s.RanImmediately
}

func (s *CreatePolarClawCronJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePolarClawCronJobResponseBody) SetApplicationId(v string) *CreatePolarClawCronJobResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBody) SetCode(v int32) *CreatePolarClawCronJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBody) SetJob(v *CreatePolarClawCronJobResponseBodyJob) *CreatePolarClawCronJobResponseBody {
	s.Job = v
	return s
}

func (s *CreatePolarClawCronJobResponseBody) SetMessage(v string) *CreatePolarClawCronJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBody) SetOk(v bool) *CreatePolarClawCronJobResponseBody {
	s.Ok = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBody) SetRanImmediately(v bool) *CreatePolarClawCronJobResponseBody {
	s.RanImmediately = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBody) SetRequestId(v string) *CreatePolarClawCronJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBody) Validate() error {
	if s.Job != nil {
		if err := s.Job.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePolarClawCronJobResponseBodyJob struct {
	// example:
	//
	// main
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 1777368967284
	CreatedAtMs *int64 `json:"CreatedAtMs,omitempty" xml:"CreatedAtMs,omitempty"`
	// example:
	//
	// false
	DeleteAfterRun *bool                                          `json:"DeleteAfterRun,omitempty" xml:"DeleteAfterRun,omitempty"`
	Delivery       *CreatePolarClawCronJobResponseBodyJobDelivery `json:"Delivery,omitempty" xml:"Delivery,omitempty" type:"Struct"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// e2c57423-12f0-45cc-a387-6155168b3201
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test
	Name     *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	Payload  *CreatePolarClawCronJobResponseBodyJobPayload  `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	Runs     []*CreatePolarClawCronJobResponseBodyJobRuns   `json:"Runs,omitempty" xml:"Runs,omitempty" type:"Repeated"`
	Schedule *CreatePolarClawCronJobResponseBodyJobSchedule `json:"Schedule,omitempty" xml:"Schedule,omitempty" type:"Struct"`
	// example:
	//
	// agent:main:feishu:direct:***
	SessionKey *string `json:"SessionKey,omitempty" xml:"SessionKey,omitempty"`
	// example:
	//
	// main
	SessionTarget *string                                     `json:"SessionTarget,omitempty" xml:"SessionTarget,omitempty"`
	State         *CreatePolarClawCronJobResponseBodyJobState `json:"State,omitempty" xml:"State,omitempty" type:"Struct"`
	// example:
	//
	// 1777370572517
	UpdatedAtMs *int64 `json:"UpdatedAtMs,omitempty" xml:"UpdatedAtMs,omitempty"`
	// example:
	//
	// now
	WakeMode *string `json:"WakeMode,omitempty" xml:"WakeMode,omitempty"`
}

func (s CreatePolarClawCronJobResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawCronJobResponseBodyJob) GoString() string {
	return s.String()
}

func (s *CreatePolarClawCronJobResponseBodyJob) GetAgentId() *string {
	return s.AgentId
}

func (s *CreatePolarClawCronJobResponseBodyJob) GetCreatedAtMs() *int64 {
	return s.CreatedAtMs
}

func (s *CreatePolarClawCronJobResponseBodyJob) GetDeleteAfterRun() *bool {
	return s.DeleteAfterRun
}

func (s *CreatePolarClawCronJobResponseBodyJob) GetDelivery() *CreatePolarClawCronJobResponseBodyJobDelivery {
	return s.Delivery
}

func (s *CreatePolarClawCronJobResponseBodyJob) GetDescription() *string {
	return s.Description
}

func (s *CreatePolarClawCronJobResponseBodyJob) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreatePolarClawCronJobResponseBodyJob) GetId() *string {
	return s.Id
}

func (s *CreatePolarClawCronJobResponseBodyJob) GetName() *string {
	return s.Name
}

func (s *CreatePolarClawCronJobResponseBodyJob) GetPayload() *CreatePolarClawCronJobResponseBodyJobPayload {
	return s.Payload
}

func (s *CreatePolarClawCronJobResponseBodyJob) GetRuns() []*CreatePolarClawCronJobResponseBodyJobRuns {
	return s.Runs
}

func (s *CreatePolarClawCronJobResponseBodyJob) GetSchedule() *CreatePolarClawCronJobResponseBodyJobSchedule {
	return s.Schedule
}

func (s *CreatePolarClawCronJobResponseBodyJob) GetSessionKey() *string {
	return s.SessionKey
}

func (s *CreatePolarClawCronJobResponseBodyJob) GetSessionTarget() *string {
	return s.SessionTarget
}

func (s *CreatePolarClawCronJobResponseBodyJob) GetState() *CreatePolarClawCronJobResponseBodyJobState {
	return s.State
}

func (s *CreatePolarClawCronJobResponseBodyJob) GetUpdatedAtMs() *int64 {
	return s.UpdatedAtMs
}

func (s *CreatePolarClawCronJobResponseBodyJob) GetWakeMode() *string {
	return s.WakeMode
}

func (s *CreatePolarClawCronJobResponseBodyJob) SetAgentId(v string) *CreatePolarClawCronJobResponseBodyJob {
	s.AgentId = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJob) SetCreatedAtMs(v int64) *CreatePolarClawCronJobResponseBodyJob {
	s.CreatedAtMs = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJob) SetDeleteAfterRun(v bool) *CreatePolarClawCronJobResponseBodyJob {
	s.DeleteAfterRun = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJob) SetDelivery(v *CreatePolarClawCronJobResponseBodyJobDelivery) *CreatePolarClawCronJobResponseBodyJob {
	s.Delivery = v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJob) SetDescription(v string) *CreatePolarClawCronJobResponseBodyJob {
	s.Description = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJob) SetEnabled(v bool) *CreatePolarClawCronJobResponseBodyJob {
	s.Enabled = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJob) SetId(v string) *CreatePolarClawCronJobResponseBodyJob {
	s.Id = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJob) SetName(v string) *CreatePolarClawCronJobResponseBodyJob {
	s.Name = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJob) SetPayload(v *CreatePolarClawCronJobResponseBodyJobPayload) *CreatePolarClawCronJobResponseBodyJob {
	s.Payload = v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJob) SetRuns(v []*CreatePolarClawCronJobResponseBodyJobRuns) *CreatePolarClawCronJobResponseBodyJob {
	s.Runs = v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJob) SetSchedule(v *CreatePolarClawCronJobResponseBodyJobSchedule) *CreatePolarClawCronJobResponseBodyJob {
	s.Schedule = v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJob) SetSessionKey(v string) *CreatePolarClawCronJobResponseBodyJob {
	s.SessionKey = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJob) SetSessionTarget(v string) *CreatePolarClawCronJobResponseBodyJob {
	s.SessionTarget = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJob) SetState(v *CreatePolarClawCronJobResponseBodyJobState) *CreatePolarClawCronJobResponseBodyJob {
	s.State = v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJob) SetUpdatedAtMs(v int64) *CreatePolarClawCronJobResponseBodyJob {
	s.UpdatedAtMs = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJob) SetWakeMode(v string) *CreatePolarClawCronJobResponseBodyJob {
	s.WakeMode = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJob) Validate() error {
	if s.Delivery != nil {
		if err := s.Delivery.Validate(); err != nil {
			return err
		}
	}
	if s.Payload != nil {
		if err := s.Payload.Validate(); err != nil {
			return err
		}
	}
	if s.Runs != nil {
		for _, item := range s.Runs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Schedule != nil {
		if err := s.Schedule.Validate(); err != nil {
			return err
		}
	}
	if s.State != nil {
		if err := s.State.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePolarClawCronJobResponseBodyJobDelivery struct {
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

func (s CreatePolarClawCronJobResponseBodyJobDelivery) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawCronJobResponseBodyJobDelivery) GoString() string {
	return s.String()
}

func (s *CreatePolarClawCronJobResponseBodyJobDelivery) GetAccountId() *string {
	return s.AccountId
}

func (s *CreatePolarClawCronJobResponseBodyJobDelivery) GetBestEffort() *bool {
	return s.BestEffort
}

func (s *CreatePolarClawCronJobResponseBodyJobDelivery) GetChannel() *string {
	return s.Channel
}

func (s *CreatePolarClawCronJobResponseBodyJobDelivery) GetMode() *string {
	return s.Mode
}

func (s *CreatePolarClawCronJobResponseBodyJobDelivery) GetTo() *string {
	return s.To
}

func (s *CreatePolarClawCronJobResponseBodyJobDelivery) SetAccountId(v string) *CreatePolarClawCronJobResponseBodyJobDelivery {
	s.AccountId = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobDelivery) SetBestEffort(v bool) *CreatePolarClawCronJobResponseBodyJobDelivery {
	s.BestEffort = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobDelivery) SetChannel(v string) *CreatePolarClawCronJobResponseBodyJobDelivery {
	s.Channel = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobDelivery) SetMode(v string) *CreatePolarClawCronJobResponseBodyJobDelivery {
	s.Mode = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobDelivery) SetTo(v string) *CreatePolarClawCronJobResponseBodyJobDelivery {
	s.To = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobDelivery) Validate() error {
	return dara.Validate(s)
}

type CreatePolarClawCronJobResponseBodyJobPayload struct {
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
	// false
	Deliver *bool `json:"Deliver,omitempty" xml:"Deliver,omitempty"`
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
	// 10
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	// example:
	//
	// ou_***
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s CreatePolarClawCronJobResponseBodyJobPayload) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawCronJobResponseBodyJobPayload) GoString() string {
	return s.String()
}

func (s *CreatePolarClawCronJobResponseBodyJobPayload) GetBestEffortDeliver() *bool {
	return s.BestEffortDeliver
}

func (s *CreatePolarClawCronJobResponseBodyJobPayload) GetChannel() *string {
	return s.Channel
}

func (s *CreatePolarClawCronJobResponseBodyJobPayload) GetDeliver() *bool {
	return s.Deliver
}

func (s *CreatePolarClawCronJobResponseBodyJobPayload) GetKind() *string {
	return s.Kind
}

func (s *CreatePolarClawCronJobResponseBodyJobPayload) GetLightContext() *bool {
	return s.LightContext
}

func (s *CreatePolarClawCronJobResponseBodyJobPayload) GetMessage() *string {
	return s.Message
}

func (s *CreatePolarClawCronJobResponseBodyJobPayload) GetModel() *string {
	return s.Model
}

func (s *CreatePolarClawCronJobResponseBodyJobPayload) GetText() *string {
	return s.Text
}

func (s *CreatePolarClawCronJobResponseBodyJobPayload) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *CreatePolarClawCronJobResponseBodyJobPayload) GetTo() *string {
	return s.To
}

func (s *CreatePolarClawCronJobResponseBodyJobPayload) SetBestEffortDeliver(v bool) *CreatePolarClawCronJobResponseBodyJobPayload {
	s.BestEffortDeliver = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobPayload) SetChannel(v string) *CreatePolarClawCronJobResponseBodyJobPayload {
	s.Channel = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobPayload) SetDeliver(v bool) *CreatePolarClawCronJobResponseBodyJobPayload {
	s.Deliver = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobPayload) SetKind(v string) *CreatePolarClawCronJobResponseBodyJobPayload {
	s.Kind = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobPayload) SetLightContext(v bool) *CreatePolarClawCronJobResponseBodyJobPayload {
	s.LightContext = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobPayload) SetMessage(v string) *CreatePolarClawCronJobResponseBodyJobPayload {
	s.Message = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobPayload) SetModel(v string) *CreatePolarClawCronJobResponseBodyJobPayload {
	s.Model = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobPayload) SetText(v string) *CreatePolarClawCronJobResponseBodyJobPayload {
	s.Text = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobPayload) SetTimeoutSeconds(v int32) *CreatePolarClawCronJobResponseBodyJobPayload {
	s.TimeoutSeconds = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobPayload) SetTo(v string) *CreatePolarClawCronJobResponseBodyJobPayload {
	s.To = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobPayload) Validate() error {
	return dara.Validate(s)
}

type CreatePolarClawCronJobResponseBodyJobRuns struct {
	// example:
	//
	// finished
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// false
	Delivered *bool `json:"Delivered,omitempty" xml:"Delivered,omitempty"`
	// example:
	//
	// not-requested
	DeliveryStatus *string `json:"DeliveryStatus,omitempty" xml:"DeliveryStatus,omitempty"`
	// example:
	//
	// 27586
	DurationMs *int64 `json:"DurationMs,omitempty" xml:"DurationMs,omitempty"`
	// example:
	//
	// f83f5278-1abe-40a6-b10e-ad3ecdc05de2
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// test
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// bailian/qwen3.5-plus
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 1777424400000
	NextRunAtMs *int64 `json:"NextRunAtMs,omitempty" xml:"NextRunAtMs,omitempty"`
	// example:
	//
	// bailian
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// example:
	//
	// 1777370544931
	RunAtMs *int64 `json:"RunAtMs,omitempty" xml:"RunAtMs,omitempty"`
	// example:
	//
	// ***
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// ok
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Generate the daily report.
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 1777370572518
	Ts    *int64                                          `json:"Ts,omitempty" xml:"Ts,omitempty"`
	Usage *CreatePolarClawCronJobResponseBodyJobRunsUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s CreatePolarClawCronJobResponseBodyJobRuns) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawCronJobResponseBodyJobRuns) GoString() string {
	return s.String()
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) GetAction() *string {
	return s.Action
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) GetDelivered() *bool {
	return s.Delivered
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) GetDeliveryStatus() *string {
	return s.DeliveryStatus
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) GetDurationMs() *int64 {
	return s.DurationMs
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) GetJobId() *string {
	return s.JobId
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) GetJobName() *string {
	return s.JobName
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) GetModel() *string {
	return s.Model
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) GetNextRunAtMs() *int64 {
	return s.NextRunAtMs
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) GetProvider() *string {
	return s.Provider
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) GetRunAtMs() *int64 {
	return s.RunAtMs
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) GetSessionId() *string {
	return s.SessionId
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) GetStatus() *string {
	return s.Status
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) GetSummary() *string {
	return s.Summary
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) GetTs() *int64 {
	return s.Ts
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) GetUsage() *CreatePolarClawCronJobResponseBodyJobRunsUsage {
	return s.Usage
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) SetAction(v string) *CreatePolarClawCronJobResponseBodyJobRuns {
	s.Action = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) SetDelivered(v bool) *CreatePolarClawCronJobResponseBodyJobRuns {
	s.Delivered = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) SetDeliveryStatus(v string) *CreatePolarClawCronJobResponseBodyJobRuns {
	s.DeliveryStatus = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) SetDurationMs(v int64) *CreatePolarClawCronJobResponseBodyJobRuns {
	s.DurationMs = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) SetJobId(v string) *CreatePolarClawCronJobResponseBodyJobRuns {
	s.JobId = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) SetJobName(v string) *CreatePolarClawCronJobResponseBodyJobRuns {
	s.JobName = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) SetModel(v string) *CreatePolarClawCronJobResponseBodyJobRuns {
	s.Model = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) SetNextRunAtMs(v int64) *CreatePolarClawCronJobResponseBodyJobRuns {
	s.NextRunAtMs = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) SetProvider(v string) *CreatePolarClawCronJobResponseBodyJobRuns {
	s.Provider = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) SetRunAtMs(v int64) *CreatePolarClawCronJobResponseBodyJobRuns {
	s.RunAtMs = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) SetSessionId(v string) *CreatePolarClawCronJobResponseBodyJobRuns {
	s.SessionId = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) SetStatus(v string) *CreatePolarClawCronJobResponseBodyJobRuns {
	s.Status = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) SetSummary(v string) *CreatePolarClawCronJobResponseBodyJobRuns {
	s.Summary = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) SetTs(v int64) *CreatePolarClawCronJobResponseBodyJobRuns {
	s.Ts = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) SetUsage(v *CreatePolarClawCronJobResponseBodyJobRunsUsage) *CreatePolarClawCronJobResponseBodyJobRuns {
	s.Usage = v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobRuns) Validate() error {
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePolarClawCronJobResponseBodyJobRunsUsage struct {
	// example:
	//
	// 30250
	InputTokens *int32 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 30250
	OutputTokens *int32 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 60500
	TotalTokens *int32 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s CreatePolarClawCronJobResponseBodyJobRunsUsage) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawCronJobResponseBodyJobRunsUsage) GoString() string {
	return s.String()
}

func (s *CreatePolarClawCronJobResponseBodyJobRunsUsage) GetInputTokens() *int32 {
	return s.InputTokens
}

func (s *CreatePolarClawCronJobResponseBodyJobRunsUsage) GetOutputTokens() *int32 {
	return s.OutputTokens
}

func (s *CreatePolarClawCronJobResponseBodyJobRunsUsage) GetTotalTokens() *int32 {
	return s.TotalTokens
}

func (s *CreatePolarClawCronJobResponseBodyJobRunsUsage) SetInputTokens(v int32) *CreatePolarClawCronJobResponseBodyJobRunsUsage {
	s.InputTokens = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobRunsUsage) SetOutputTokens(v int32) *CreatePolarClawCronJobResponseBodyJobRunsUsage {
	s.OutputTokens = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobRunsUsage) SetTotalTokens(v int32) *CreatePolarClawCronJobResponseBodyJobRunsUsage {
	s.TotalTokens = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobRunsUsage) Validate() error {
	return dara.Validate(s)
}

type CreatePolarClawCronJobResponseBodyJobSchedule struct {
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
	// 1000
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
	// Asia/Shanghai
	Tz *string `json:"Tz,omitempty" xml:"Tz,omitempty"`
}

func (s CreatePolarClawCronJobResponseBodyJobSchedule) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawCronJobResponseBodyJobSchedule) GoString() string {
	return s.String()
}

func (s *CreatePolarClawCronJobResponseBodyJobSchedule) GetAnchorMs() *int64 {
	return s.AnchorMs
}

func (s *CreatePolarClawCronJobResponseBodyJobSchedule) GetAt() *string {
	return s.At
}

func (s *CreatePolarClawCronJobResponseBodyJobSchedule) GetEveryMs() *int64 {
	return s.EveryMs
}

func (s *CreatePolarClawCronJobResponseBodyJobSchedule) GetExpr() *string {
	return s.Expr
}

func (s *CreatePolarClawCronJobResponseBodyJobSchedule) GetKind() *string {
	return s.Kind
}

func (s *CreatePolarClawCronJobResponseBodyJobSchedule) GetTz() *string {
	return s.Tz
}

func (s *CreatePolarClawCronJobResponseBodyJobSchedule) SetAnchorMs(v int64) *CreatePolarClawCronJobResponseBodyJobSchedule {
	s.AnchorMs = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobSchedule) SetAt(v string) *CreatePolarClawCronJobResponseBodyJobSchedule {
	s.At = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobSchedule) SetEveryMs(v int64) *CreatePolarClawCronJobResponseBodyJobSchedule {
	s.EveryMs = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobSchedule) SetExpr(v string) *CreatePolarClawCronJobResponseBodyJobSchedule {
	s.Expr = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobSchedule) SetKind(v string) *CreatePolarClawCronJobResponseBodyJobSchedule {
	s.Kind = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobSchedule) SetTz(v string) *CreatePolarClawCronJobResponseBodyJobSchedule {
	s.Tz = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobSchedule) Validate() error {
	return dara.Validate(s)
}

type CreatePolarClawCronJobResponseBodyJobState struct {
	// example:
	//
	// 0
	ConsecutiveErrors *int32 `json:"ConsecutiveErrors,omitempty" xml:"ConsecutiveErrors,omitempty"`
	// example:
	//
	// 1777370544931
	LastRunAtMs *int64 `json:"LastRunAtMs,omitempty" xml:"LastRunAtMs,omitempty"`
	// example:
	//
	// ok
	LastRunStatus *string `json:"LastRunStatus,omitempty" xml:"LastRunStatus,omitempty"`
	// example:
	//
	// 1777424400000
	NextRunAtMs *int64 `json:"NextRunAtMs,omitempty" xml:"NextRunAtMs,omitempty"`
}

func (s CreatePolarClawCronJobResponseBodyJobState) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawCronJobResponseBodyJobState) GoString() string {
	return s.String()
}

func (s *CreatePolarClawCronJobResponseBodyJobState) GetConsecutiveErrors() *int32 {
	return s.ConsecutiveErrors
}

func (s *CreatePolarClawCronJobResponseBodyJobState) GetLastRunAtMs() *int64 {
	return s.LastRunAtMs
}

func (s *CreatePolarClawCronJobResponseBodyJobState) GetLastRunStatus() *string {
	return s.LastRunStatus
}

func (s *CreatePolarClawCronJobResponseBodyJobState) GetNextRunAtMs() *int64 {
	return s.NextRunAtMs
}

func (s *CreatePolarClawCronJobResponseBodyJobState) SetConsecutiveErrors(v int32) *CreatePolarClawCronJobResponseBodyJobState {
	s.ConsecutiveErrors = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobState) SetLastRunAtMs(v int64) *CreatePolarClawCronJobResponseBodyJobState {
	s.LastRunAtMs = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobState) SetLastRunStatus(v string) *CreatePolarClawCronJobResponseBodyJobState {
	s.LastRunStatus = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobState) SetNextRunAtMs(v int64) *CreatePolarClawCronJobResponseBodyJobState {
	s.NextRunAtMs = &v
	return s
}

func (s *CreatePolarClawCronJobResponseBodyJobState) Validate() error {
	return dara.Validate(s)
}
