// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolarClawCronJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpdatePolarClawCronJobResponseBody
	GetApplicationId() *string
	SetCode(v int32) *UpdatePolarClawCronJobResponseBody
	GetCode() *int32
	SetJob(v *UpdatePolarClawCronJobResponseBodyJob) *UpdatePolarClawCronJobResponseBody
	GetJob() *UpdatePolarClawCronJobResponseBodyJob
	SetMessage(v string) *UpdatePolarClawCronJobResponseBody
	GetMessage() *string
	SetOk(v bool) *UpdatePolarClawCronJobResponseBody
	GetOk() *bool
	SetRequestId(v string) *UpdatePolarClawCronJobResponseBody
	GetRequestId() *string
}

type UpdatePolarClawCronJobResponseBody struct {
	// The application ID.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details of the updated job.
	Job *UpdatePolarClawCronJobResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// The success message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Ok *bool `json:"Ok,omitempty" xml:"Ok,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2281C6C9-CBAB-1AFD-8400-670750CF6025_2212
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePolarClawCronJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawCronJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawCronJobResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdatePolarClawCronJobResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdatePolarClawCronJobResponseBody) GetJob() *UpdatePolarClawCronJobResponseBodyJob {
	return s.Job
}

func (s *UpdatePolarClawCronJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePolarClawCronJobResponseBody) GetOk() *bool {
	return s.Ok
}

func (s *UpdatePolarClawCronJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePolarClawCronJobResponseBody) SetApplicationId(v string) *UpdatePolarClawCronJobResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBody) SetCode(v int32) *UpdatePolarClawCronJobResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBody) SetJob(v *UpdatePolarClawCronJobResponseBodyJob) *UpdatePolarClawCronJobResponseBody {
	s.Job = v
	return s
}

func (s *UpdatePolarClawCronJobResponseBody) SetMessage(v string) *UpdatePolarClawCronJobResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBody) SetOk(v bool) *UpdatePolarClawCronJobResponseBody {
	s.Ok = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBody) SetRequestId(v string) *UpdatePolarClawCronJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBody) Validate() error {
	if s.Job != nil {
		if err := s.Job.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePolarClawCronJobResponseBodyJob struct {
	// The optional ID of the Agent that runs the job.
	//
	// example:
	//
	// work
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The creation timestamp, in milliseconds.
	//
	// example:
	//
	// 1777368967284
	CreatedAtMs *int64 `json:"CreatedAtMs,omitempty" xml:"CreatedAtMs,omitempty"`
	// Indicates whether the job is deleted after its first run.
	//
	// example:
	//
	// false
	DeleteAfterRun *bool `json:"DeleteAfterRun,omitempty" xml:"DeleteAfterRun,omitempty"`
	// The optional result delivery configuration.
	Delivery *UpdatePolarClawCronJobResponseBodyJobDelivery `json:"Delivery,omitempty" xml:"Delivery,omitempty" type:"Struct"`
	// The optional job description.
	//
	// example:
	//
	// Daily report generation
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the job is enabled.
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The job ID (UUID).
	//
	// example:
	//
	// 0ee00f56-f467-4d41-858c-ca4ede2c770e
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The job name.
	//
	// example:
	//
	// daily-report
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The execution payload configuration.
	Payload *UpdatePolarClawCronJobResponseBodyJobPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// **The run history. Returned only if `IncludeRuns` is `true`.**
	Runs []*UpdatePolarClawCronJobResponseBodyJobRuns `json:"Runs,omitempty" xml:"Runs,omitempty" type:"Repeated"`
	// The schedule configuration.
	Schedule *UpdatePolarClawCronJobResponseBodyJobSchedule `json:"Schedule,omitempty" xml:"Schedule,omitempty" type:"Struct"`
	// The optional session routing key.
	//
	// example:
	//
	// abc123
	SessionKey *string `json:"SessionKey,omitempty" xml:"SessionKey,omitempty"`
	// The session target. Valid values: `main`, `isolated`, or `current`.
	//
	// example:
	//
	// main
	SessionTarget *string `json:"SessionTarget,omitempty" xml:"SessionTarget,omitempty"`
	// **The running state of the job.**
	State *UpdatePolarClawCronJobResponseBodyJobState `json:"State,omitempty" xml:"State,omitempty" type:"Struct"`
	// The last update timestamp, in milliseconds.
	//
	// example:
	//
	// 1777370572517
	UpdatedAtMs *int64 `json:"UpdatedAtMs,omitempty" xml:"UpdatedAtMs,omitempty"`
	// The wake mode. Valid values: `now` or `next-heartbeat`.
	//
	// example:
	//
	// now
	WakeMode *string `json:"WakeMode,omitempty" xml:"WakeMode,omitempty"`
}

func (s UpdatePolarClawCronJobResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawCronJobResponseBodyJob) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawCronJobResponseBodyJob) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdatePolarClawCronJobResponseBodyJob) GetCreatedAtMs() *int64 {
	return s.CreatedAtMs
}

func (s *UpdatePolarClawCronJobResponseBodyJob) GetDeleteAfterRun() *bool {
	return s.DeleteAfterRun
}

func (s *UpdatePolarClawCronJobResponseBodyJob) GetDelivery() *UpdatePolarClawCronJobResponseBodyJobDelivery {
	return s.Delivery
}

func (s *UpdatePolarClawCronJobResponseBodyJob) GetDescription() *string {
	return s.Description
}

func (s *UpdatePolarClawCronJobResponseBodyJob) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdatePolarClawCronJobResponseBodyJob) GetId() *string {
	return s.Id
}

func (s *UpdatePolarClawCronJobResponseBodyJob) GetName() *string {
	return s.Name
}

func (s *UpdatePolarClawCronJobResponseBodyJob) GetPayload() *UpdatePolarClawCronJobResponseBodyJobPayload {
	return s.Payload
}

func (s *UpdatePolarClawCronJobResponseBodyJob) GetRuns() []*UpdatePolarClawCronJobResponseBodyJobRuns {
	return s.Runs
}

func (s *UpdatePolarClawCronJobResponseBodyJob) GetSchedule() *UpdatePolarClawCronJobResponseBodyJobSchedule {
	return s.Schedule
}

func (s *UpdatePolarClawCronJobResponseBodyJob) GetSessionKey() *string {
	return s.SessionKey
}

func (s *UpdatePolarClawCronJobResponseBodyJob) GetSessionTarget() *string {
	return s.SessionTarget
}

func (s *UpdatePolarClawCronJobResponseBodyJob) GetState() *UpdatePolarClawCronJobResponseBodyJobState {
	return s.State
}

func (s *UpdatePolarClawCronJobResponseBodyJob) GetUpdatedAtMs() *int64 {
	return s.UpdatedAtMs
}

func (s *UpdatePolarClawCronJobResponseBodyJob) GetWakeMode() *string {
	return s.WakeMode
}

func (s *UpdatePolarClawCronJobResponseBodyJob) SetAgentId(v string) *UpdatePolarClawCronJobResponseBodyJob {
	s.AgentId = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJob) SetCreatedAtMs(v int64) *UpdatePolarClawCronJobResponseBodyJob {
	s.CreatedAtMs = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJob) SetDeleteAfterRun(v bool) *UpdatePolarClawCronJobResponseBodyJob {
	s.DeleteAfterRun = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJob) SetDelivery(v *UpdatePolarClawCronJobResponseBodyJobDelivery) *UpdatePolarClawCronJobResponseBodyJob {
	s.Delivery = v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJob) SetDescription(v string) *UpdatePolarClawCronJobResponseBodyJob {
	s.Description = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJob) SetEnabled(v bool) *UpdatePolarClawCronJobResponseBodyJob {
	s.Enabled = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJob) SetId(v string) *UpdatePolarClawCronJobResponseBodyJob {
	s.Id = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJob) SetName(v string) *UpdatePolarClawCronJobResponseBodyJob {
	s.Name = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJob) SetPayload(v *UpdatePolarClawCronJobResponseBodyJobPayload) *UpdatePolarClawCronJobResponseBodyJob {
	s.Payload = v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJob) SetRuns(v []*UpdatePolarClawCronJobResponseBodyJobRuns) *UpdatePolarClawCronJobResponseBodyJob {
	s.Runs = v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJob) SetSchedule(v *UpdatePolarClawCronJobResponseBodyJobSchedule) *UpdatePolarClawCronJobResponseBodyJob {
	s.Schedule = v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJob) SetSessionKey(v string) *UpdatePolarClawCronJobResponseBodyJob {
	s.SessionKey = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJob) SetSessionTarget(v string) *UpdatePolarClawCronJobResponseBodyJob {
	s.SessionTarget = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJob) SetState(v *UpdatePolarClawCronJobResponseBodyJobState) *UpdatePolarClawCronJobResponseBodyJob {
	s.State = v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJob) SetUpdatedAtMs(v int64) *UpdatePolarClawCronJobResponseBodyJob {
	s.UpdatedAtMs = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJob) SetWakeMode(v string) *UpdatePolarClawCronJobResponseBodyJob {
	s.WakeMode = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJob) Validate() error {
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

type UpdatePolarClawCronJobResponseBodyJobDelivery struct {
	// The optional channel account ID.
	//
	// example:
	//
	// default
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Indicates whether to ignore delivery failures.
	BestEffort *bool `json:"BestEffort,omitempty" xml:"BestEffort,omitempty"`
	// The delivery channel.
	//
	// example:
	//
	// telegram
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The delivery mode. Valid values: `none`, `announce`, or `webhook`.
	//
	// example:
	//
	// announce
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The delivery target. Required and must be a URL if `Mode` is `webhook`.
	//
	// example:
	//
	// https://example.com/webhook
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s UpdatePolarClawCronJobResponseBodyJobDelivery) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawCronJobResponseBodyJobDelivery) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawCronJobResponseBodyJobDelivery) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdatePolarClawCronJobResponseBodyJobDelivery) GetBestEffort() *bool {
	return s.BestEffort
}

func (s *UpdatePolarClawCronJobResponseBodyJobDelivery) GetChannel() *string {
	return s.Channel
}

func (s *UpdatePolarClawCronJobResponseBodyJobDelivery) GetMode() *string {
	return s.Mode
}

func (s *UpdatePolarClawCronJobResponseBodyJobDelivery) GetTo() *string {
	return s.To
}

func (s *UpdatePolarClawCronJobResponseBodyJobDelivery) SetAccountId(v string) *UpdatePolarClawCronJobResponseBodyJobDelivery {
	s.AccountId = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobDelivery) SetBestEffort(v bool) *UpdatePolarClawCronJobResponseBodyJobDelivery {
	s.BestEffort = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobDelivery) SetChannel(v string) *UpdatePolarClawCronJobResponseBodyJobDelivery {
	s.Channel = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobDelivery) SetMode(v string) *UpdatePolarClawCronJobResponseBodyJobDelivery {
	s.Mode = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobDelivery) SetTo(v string) *UpdatePolarClawCronJobResponseBodyJobDelivery {
	s.To = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobDelivery) Validate() error {
	return dara.Validate(s)
}

type UpdatePolarClawCronJobResponseBodyJobPayload struct {
	// Indicates whether to ignore delivery failures.
	//
	// example:
	//
	// false
	BestEffortDeliver *bool `json:"BestEffortDeliver,omitempty" xml:"BestEffortDeliver,omitempty"`
	// The optional delivery channel ID.
	//
	// example:
	//
	// telegram
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// Indicates whether to deliver the output to a channel.
	//
	// example:
	//
	// false
	Deliver *bool `json:"Deliver,omitempty" xml:"Deliver,omitempty"`
	// The payload type. Valid values: `agentTurn` (Agent conversation) or `systemEvent` (system event).
	//
	// example:
	//
	// agentTurn
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// Indicates whether to use a lightweight context.
	//
	// example:
	//
	// false
	LightContext *bool `json:"LightContext,omitempty" xml:"LightContext,omitempty"`
	// The Agent conversation prompt.
	//
	// example:
	//
	// Generate the daily report.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The model override.
	//
	// example:
	//
	// anthropic/sonnet-4.6
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The system event text.
	//
	// example:
	//
	// Generate the daily report.
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The optional execution timeout in seconds.
	//
	// example:
	//
	// 10
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	// The optional delivery target.
	//
	// example:
	//
	// team
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s UpdatePolarClawCronJobResponseBodyJobPayload) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawCronJobResponseBodyJobPayload) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawCronJobResponseBodyJobPayload) GetBestEffortDeliver() *bool {
	return s.BestEffortDeliver
}

func (s *UpdatePolarClawCronJobResponseBodyJobPayload) GetChannel() *string {
	return s.Channel
}

func (s *UpdatePolarClawCronJobResponseBodyJobPayload) GetDeliver() *bool {
	return s.Deliver
}

func (s *UpdatePolarClawCronJobResponseBodyJobPayload) GetKind() *string {
	return s.Kind
}

func (s *UpdatePolarClawCronJobResponseBodyJobPayload) GetLightContext() *bool {
	return s.LightContext
}

func (s *UpdatePolarClawCronJobResponseBodyJobPayload) GetMessage() *string {
	return s.Message
}

func (s *UpdatePolarClawCronJobResponseBodyJobPayload) GetModel() *string {
	return s.Model
}

func (s *UpdatePolarClawCronJobResponseBodyJobPayload) GetText() *string {
	return s.Text
}

func (s *UpdatePolarClawCronJobResponseBodyJobPayload) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *UpdatePolarClawCronJobResponseBodyJobPayload) GetTo() *string {
	return s.To
}

func (s *UpdatePolarClawCronJobResponseBodyJobPayload) SetBestEffortDeliver(v bool) *UpdatePolarClawCronJobResponseBodyJobPayload {
	s.BestEffortDeliver = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobPayload) SetChannel(v string) *UpdatePolarClawCronJobResponseBodyJobPayload {
	s.Channel = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobPayload) SetDeliver(v bool) *UpdatePolarClawCronJobResponseBodyJobPayload {
	s.Deliver = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobPayload) SetKind(v string) *UpdatePolarClawCronJobResponseBodyJobPayload {
	s.Kind = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobPayload) SetLightContext(v bool) *UpdatePolarClawCronJobResponseBodyJobPayload {
	s.LightContext = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobPayload) SetMessage(v string) *UpdatePolarClawCronJobResponseBodyJobPayload {
	s.Message = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobPayload) SetModel(v string) *UpdatePolarClawCronJobResponseBodyJobPayload {
	s.Model = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobPayload) SetText(v string) *UpdatePolarClawCronJobResponseBodyJobPayload {
	s.Text = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobPayload) SetTimeoutSeconds(v int32) *UpdatePolarClawCronJobResponseBodyJobPayload {
	s.TimeoutSeconds = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobPayload) SetTo(v string) *UpdatePolarClawCronJobResponseBodyJobPayload {
	s.To = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobPayload) Validate() error {
	return dara.Validate(s)
}

type UpdatePolarClawCronJobResponseBodyJobRuns struct {
	// The action performed. Valid values: `finished`, `error`, or `skipped`.
	//
	// example:
	//
	// finished
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// Indicates whether the result was delivered.
	//
	// example:
	//
	// true
	Delivered *bool `json:"Delivered,omitempty" xml:"Delivered,omitempty"`
	// The delivery status.
	//
	// example:
	//
	// delivered
	DeliveryStatus *string `json:"DeliveryStatus,omitempty" xml:"DeliveryStatus,omitempty"`
	// The execution duration, in milliseconds.
	//
	// example:
	//
	// 27586
	DurationMs *int64 `json:"DurationMs,omitempty" xml:"DurationMs,omitempty"`
	// The associated job ID.
	//
	// example:
	//
	// 0ee00f56-f467-4d41-858c-ca4ede2c770e
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The job name.
	//
	// example:
	//
	// daily-report
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The model used for the run.
	//
	// example:
	//
	// anthropic/sonnet-4.6
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The timestamp of the next scheduled run, in milliseconds.
	//
	// example:
	//
	// 1777424400000
	NextRunAtMs *int64 `json:"NextRunAtMs,omitempty" xml:"NextRunAtMs,omitempty"`
	// The model provider.
	//
	// example:
	//
	// anthropic
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The actual execution timestamp, in milliseconds.
	//
	// example:
	//
	// 1777370544931
	RunAtMs *int64 `json:"RunAtMs,omitempty" xml:"RunAtMs,omitempty"`
	// The associated session ID.
	//
	// example:
	//
	// abc123
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The status of the run. Valid values: `ok`, `error`, or `skipped`.
	//
	// example:
	//
	// ok
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The run summary text.
	//
	// example:
	//
	// Report generated successfully.
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// The run timestamp, in milliseconds.
	//
	// example:
	//
	// 1777370572518
	Ts *int64 `json:"Ts,omitempty" xml:"Ts,omitempty"`
	// Optional token usage details.
	Usage *UpdatePolarClawCronJobResponseBodyJobRunsUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s UpdatePolarClawCronJobResponseBodyJobRuns) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawCronJobResponseBodyJobRuns) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) GetAction() *string {
	return s.Action
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) GetDelivered() *bool {
	return s.Delivered
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) GetDeliveryStatus() *string {
	return s.DeliveryStatus
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) GetDurationMs() *int64 {
	return s.DurationMs
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) GetJobId() *string {
	return s.JobId
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) GetJobName() *string {
	return s.JobName
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) GetModel() *string {
	return s.Model
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) GetNextRunAtMs() *int64 {
	return s.NextRunAtMs
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) GetProvider() *string {
	return s.Provider
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) GetRunAtMs() *int64 {
	return s.RunAtMs
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) GetSessionId() *string {
	return s.SessionId
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) GetStatus() *string {
	return s.Status
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) GetSummary() *string {
	return s.Summary
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) GetTs() *int64 {
	return s.Ts
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) GetUsage() *UpdatePolarClawCronJobResponseBodyJobRunsUsage {
	return s.Usage
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) SetAction(v string) *UpdatePolarClawCronJobResponseBodyJobRuns {
	s.Action = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) SetDelivered(v bool) *UpdatePolarClawCronJobResponseBodyJobRuns {
	s.Delivered = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) SetDeliveryStatus(v string) *UpdatePolarClawCronJobResponseBodyJobRuns {
	s.DeliveryStatus = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) SetDurationMs(v int64) *UpdatePolarClawCronJobResponseBodyJobRuns {
	s.DurationMs = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) SetJobId(v string) *UpdatePolarClawCronJobResponseBodyJobRuns {
	s.JobId = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) SetJobName(v string) *UpdatePolarClawCronJobResponseBodyJobRuns {
	s.JobName = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) SetModel(v string) *UpdatePolarClawCronJobResponseBodyJobRuns {
	s.Model = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) SetNextRunAtMs(v int64) *UpdatePolarClawCronJobResponseBodyJobRuns {
	s.NextRunAtMs = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) SetProvider(v string) *UpdatePolarClawCronJobResponseBodyJobRuns {
	s.Provider = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) SetRunAtMs(v int64) *UpdatePolarClawCronJobResponseBodyJobRuns {
	s.RunAtMs = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) SetSessionId(v string) *UpdatePolarClawCronJobResponseBodyJobRuns {
	s.SessionId = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) SetStatus(v string) *UpdatePolarClawCronJobResponseBodyJobRuns {
	s.Status = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) SetSummary(v string) *UpdatePolarClawCronJobResponseBodyJobRuns {
	s.Summary = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) SetTs(v int64) *UpdatePolarClawCronJobResponseBodyJobRuns {
	s.Ts = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) SetUsage(v *UpdatePolarClawCronJobResponseBodyJobRunsUsage) *UpdatePolarClawCronJobResponseBodyJobRuns {
	s.Usage = v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobRuns) Validate() error {
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePolarClawCronJobResponseBodyJobRunsUsage struct {
	// The number of input tokens.
	//
	// example:
	//
	// 30250
	InputTokens *int32 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// The number of output tokens.
	//
	// example:
	//
	// 30250
	OutputTokens *int32 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// The total number of tokens.
	//
	// example:
	//
	// 60500
	TotalTokens *int32 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s UpdatePolarClawCronJobResponseBodyJobRunsUsage) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawCronJobResponseBodyJobRunsUsage) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawCronJobResponseBodyJobRunsUsage) GetInputTokens() *int32 {
	return s.InputTokens
}

func (s *UpdatePolarClawCronJobResponseBodyJobRunsUsage) GetOutputTokens() *int32 {
	return s.OutputTokens
}

func (s *UpdatePolarClawCronJobResponseBodyJobRunsUsage) GetTotalTokens() *int32 {
	return s.TotalTokens
}

func (s *UpdatePolarClawCronJobResponseBodyJobRunsUsage) SetInputTokens(v int32) *UpdatePolarClawCronJobResponseBodyJobRunsUsage {
	s.InputTokens = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobRunsUsage) SetOutputTokens(v int32) *UpdatePolarClawCronJobResponseBodyJobRunsUsage {
	s.OutputTokens = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobRunsUsage) SetTotalTokens(v int32) *UpdatePolarClawCronJobResponseBodyJobRunsUsage {
	s.TotalTokens = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobRunsUsage) Validate() error {
	return dara.Validate(s)
}

type UpdatePolarClawCronJobResponseBodyJobSchedule struct {
	// The base timestamp for interval alignment, in milliseconds.
	//
	// example:
	//
	// 1777370572518
	AnchorMs *int64 `json:"AnchorMs,omitempty" xml:"AnchorMs,omitempty"`
	// The ISO 8601 timestamp. Required if `Kind` is `at`.
	//
	// example:
	//
	// 2026-04-10T09:00:00+08:00
	At *string `json:"At,omitempty" xml:"At,omitempty"`
	// The interval in milliseconds. Required if `Kind` is `every`.
	//
	// example:
	//
	// 100000
	EveryMs *int64 `json:"EveryMs,omitempty" xml:"EveryMs,omitempty"`
	// The cron expression. Required if `Kind` is `cron`.
	//
	// example:
	//
	// 0 9 	- 	- *
	Expr *string `json:"Expr,omitempty" xml:"Expr,omitempty"`
	// The schedule type. Valid values: `cron` (cron expression), `every` (fixed interval), or `at` (one-time).
	//
	// example:
	//
	// cron
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// The IANA time zone, such as `Asia/Shanghai`.
	//
	// example:
	//
	// Asia/Shanghai
	Tz *string `json:"Tz,omitempty" xml:"Tz,omitempty"`
}

func (s UpdatePolarClawCronJobResponseBodyJobSchedule) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawCronJobResponseBodyJobSchedule) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawCronJobResponseBodyJobSchedule) GetAnchorMs() *int64 {
	return s.AnchorMs
}

func (s *UpdatePolarClawCronJobResponseBodyJobSchedule) GetAt() *string {
	return s.At
}

func (s *UpdatePolarClawCronJobResponseBodyJobSchedule) GetEveryMs() *int64 {
	return s.EveryMs
}

func (s *UpdatePolarClawCronJobResponseBodyJobSchedule) GetExpr() *string {
	return s.Expr
}

func (s *UpdatePolarClawCronJobResponseBodyJobSchedule) GetKind() *string {
	return s.Kind
}

func (s *UpdatePolarClawCronJobResponseBodyJobSchedule) GetTz() *string {
	return s.Tz
}

func (s *UpdatePolarClawCronJobResponseBodyJobSchedule) SetAnchorMs(v int64) *UpdatePolarClawCronJobResponseBodyJobSchedule {
	s.AnchorMs = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobSchedule) SetAt(v string) *UpdatePolarClawCronJobResponseBodyJobSchedule {
	s.At = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobSchedule) SetEveryMs(v int64) *UpdatePolarClawCronJobResponseBodyJobSchedule {
	s.EveryMs = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobSchedule) SetExpr(v string) *UpdatePolarClawCronJobResponseBodyJobSchedule {
	s.Expr = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobSchedule) SetKind(v string) *UpdatePolarClawCronJobResponseBodyJobSchedule {
	s.Kind = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobSchedule) SetTz(v string) *UpdatePolarClawCronJobResponseBodyJobSchedule {
	s.Tz = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobSchedule) Validate() error {
	return dara.Validate(s)
}

type UpdatePolarClawCronJobResponseBodyJobState struct {
	// **The number of consecutive failures.**
	//
	// example:
	//
	// 0
	ConsecutiveErrors *int32 `json:"ConsecutiveErrors,omitempty" xml:"ConsecutiveErrors,omitempty"`
	// **The optional timestamp of the last run, in milliseconds.**
	//
	// example:
	//
	// 1777370544931
	LastRunAtMs *int64 `json:"LastRunAtMs,omitempty" xml:"LastRunAtMs,omitempty"`
	// **The optional status of the last run.**
	//
	// example:
	//
	// ok
	LastRunStatus *string `json:"LastRunStatus,omitempty" xml:"LastRunStatus,omitempty"`
	// **The timestamp for the next run, in milliseconds.**
	//
	// example:
	//
	// 1777424400000
	NextRunAtMs *int64 `json:"NextRunAtMs,omitempty" xml:"NextRunAtMs,omitempty"`
}

func (s UpdatePolarClawCronJobResponseBodyJobState) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawCronJobResponseBodyJobState) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawCronJobResponseBodyJobState) GetConsecutiveErrors() *int32 {
	return s.ConsecutiveErrors
}

func (s *UpdatePolarClawCronJobResponseBodyJobState) GetLastRunAtMs() *int64 {
	return s.LastRunAtMs
}

func (s *UpdatePolarClawCronJobResponseBodyJobState) GetLastRunStatus() *string {
	return s.LastRunStatus
}

func (s *UpdatePolarClawCronJobResponseBodyJobState) GetNextRunAtMs() *int64 {
	return s.NextRunAtMs
}

func (s *UpdatePolarClawCronJobResponseBodyJobState) SetConsecutiveErrors(v int32) *UpdatePolarClawCronJobResponseBodyJobState {
	s.ConsecutiveErrors = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobState) SetLastRunAtMs(v int64) *UpdatePolarClawCronJobResponseBodyJobState {
	s.LastRunAtMs = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobState) SetLastRunStatus(v string) *UpdatePolarClawCronJobResponseBodyJobState {
	s.LastRunStatus = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobState) SetNextRunAtMs(v int64) *UpdatePolarClawCronJobResponseBodyJobState {
	s.NextRunAtMs = &v
	return s
}

func (s *UpdatePolarClawCronJobResponseBodyJobState) Validate() error {
	return dara.Validate(s)
}
