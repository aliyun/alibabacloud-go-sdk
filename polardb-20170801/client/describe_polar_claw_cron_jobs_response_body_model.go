// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawCronJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribePolarClawCronJobsResponseBody
	GetApplicationId() *string
	SetCode(v int32) *DescribePolarClawCronJobsResponseBody
	GetCode() *int32
	SetJobs(v []*DescribePolarClawCronJobsResponseBodyJobs) *DescribePolarClawCronJobsResponseBody
	GetJobs() []*DescribePolarClawCronJobsResponseBodyJobs
	SetMessage(v string) *DescribePolarClawCronJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribePolarClawCronJobsResponseBody
	GetRequestId() *string
	SetTotal(v int32) *DescribePolarClawCronJobsResponseBody
	GetTotal() *int32
}

type DescribePolarClawCronJobsResponseBody struct {
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 200
	Code *int32                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Jobs []*DescribePolarClawCronJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2281C6C9-CBAB-1AFD-8400-670750CF6025_2212
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribePolarClawCronJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawCronJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarClawCronJobsResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawCronJobsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribePolarClawCronJobsResponseBody) GetJobs() []*DescribePolarClawCronJobsResponseBodyJobs {
	return s.Jobs
}

func (s *DescribePolarClawCronJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePolarClawCronJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarClawCronJobsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribePolarClawCronJobsResponseBody) SetApplicationId(v string) *DescribePolarClawCronJobsResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBody) SetCode(v int32) *DescribePolarClawCronJobsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBody) SetJobs(v []*DescribePolarClawCronJobsResponseBodyJobs) *DescribePolarClawCronJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *DescribePolarClawCronJobsResponseBody) SetMessage(v string) *DescribePolarClawCronJobsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBody) SetRequestId(v string) *DescribePolarClawCronJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBody) SetTotal(v int32) *DescribePolarClawCronJobsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBody) Validate() error {
	if s.Jobs != nil {
		for _, item := range s.Jobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolarClawCronJobsResponseBodyJobs struct {
	// example:
	//
	// main
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 1777370572517
	CreatedAtMs *int64 `json:"CreatedAtMs,omitempty" xml:"CreatedAtMs,omitempty"`
	// example:
	//
	// false
	DeleteAfterRun *bool                                              `json:"DeleteAfterRun,omitempty" xml:"DeleteAfterRun,omitempty"`
	Delivery       *DescribePolarClawCronJobsResponseBodyJobsDelivery `json:"Delivery,omitempty" xml:"Delivery,omitempty" type:"Struct"`
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
	// 0ee00f56-f467-4d41-858c-ca4ede2c770e
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test
	Name     *string                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	Payload  *DescribePolarClawCronJobsResponseBodyJobsPayload  `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	Runs     []*DescribePolarClawCronJobsResponseBodyJobsRuns   `json:"Runs,omitempty" xml:"Runs,omitempty" type:"Repeated"`
	Schedule *DescribePolarClawCronJobsResponseBodyJobsSchedule `json:"Schedule,omitempty" xml:"Schedule,omitempty" type:"Struct"`
	// example:
	//
	// agent:main:feishu:direct:***
	SessionKey *string `json:"SessionKey,omitempty" xml:"SessionKey,omitempty"`
	// example:
	//
	// main
	SessionTarget *string                                         `json:"SessionTarget,omitempty" xml:"SessionTarget,omitempty"`
	State         *DescribePolarClawCronJobsResponseBodyJobsState `json:"State,omitempty" xml:"State,omitempty" type:"Struct"`
	// example:
	//
	// 1777368967284
	UpdatedAtMs *int64 `json:"UpdatedAtMs,omitempty" xml:"UpdatedAtMs,omitempty"`
	// example:
	//
	// now
	WakeMode *string `json:"WakeMode,omitempty" xml:"WakeMode,omitempty"`
}

func (s DescribePolarClawCronJobsResponseBodyJobs) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawCronJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) GetAgentId() *string {
	return s.AgentId
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) GetCreatedAtMs() *int64 {
	return s.CreatedAtMs
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) GetDeleteAfterRun() *bool {
	return s.DeleteAfterRun
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) GetDelivery() *DescribePolarClawCronJobsResponseBodyJobsDelivery {
	return s.Delivery
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) GetDescription() *string {
	return s.Description
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) GetId() *string {
	return s.Id
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) GetName() *string {
	return s.Name
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) GetPayload() *DescribePolarClawCronJobsResponseBodyJobsPayload {
	return s.Payload
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) GetRuns() []*DescribePolarClawCronJobsResponseBodyJobsRuns {
	return s.Runs
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) GetSchedule() *DescribePolarClawCronJobsResponseBodyJobsSchedule {
	return s.Schedule
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) GetSessionKey() *string {
	return s.SessionKey
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) GetSessionTarget() *string {
	return s.SessionTarget
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) GetState() *DescribePolarClawCronJobsResponseBodyJobsState {
	return s.State
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) GetUpdatedAtMs() *int64 {
	return s.UpdatedAtMs
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) GetWakeMode() *string {
	return s.WakeMode
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) SetAgentId(v string) *DescribePolarClawCronJobsResponseBodyJobs {
	s.AgentId = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) SetCreatedAtMs(v int64) *DescribePolarClawCronJobsResponseBodyJobs {
	s.CreatedAtMs = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) SetDeleteAfterRun(v bool) *DescribePolarClawCronJobsResponseBodyJobs {
	s.DeleteAfterRun = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) SetDelivery(v *DescribePolarClawCronJobsResponseBodyJobsDelivery) *DescribePolarClawCronJobsResponseBodyJobs {
	s.Delivery = v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) SetDescription(v string) *DescribePolarClawCronJobsResponseBodyJobs {
	s.Description = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) SetEnabled(v bool) *DescribePolarClawCronJobsResponseBodyJobs {
	s.Enabled = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) SetId(v string) *DescribePolarClawCronJobsResponseBodyJobs {
	s.Id = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) SetName(v string) *DescribePolarClawCronJobsResponseBodyJobs {
	s.Name = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) SetPayload(v *DescribePolarClawCronJobsResponseBodyJobsPayload) *DescribePolarClawCronJobsResponseBodyJobs {
	s.Payload = v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) SetRuns(v []*DescribePolarClawCronJobsResponseBodyJobsRuns) *DescribePolarClawCronJobsResponseBodyJobs {
	s.Runs = v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) SetSchedule(v *DescribePolarClawCronJobsResponseBodyJobsSchedule) *DescribePolarClawCronJobsResponseBodyJobs {
	s.Schedule = v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) SetSessionKey(v string) *DescribePolarClawCronJobsResponseBodyJobs {
	s.SessionKey = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) SetSessionTarget(v string) *DescribePolarClawCronJobsResponseBodyJobs {
	s.SessionTarget = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) SetState(v *DescribePolarClawCronJobsResponseBodyJobsState) *DescribePolarClawCronJobsResponseBodyJobs {
	s.State = v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) SetUpdatedAtMs(v int64) *DescribePolarClawCronJobsResponseBodyJobs {
	s.UpdatedAtMs = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) SetWakeMode(v string) *DescribePolarClawCronJobsResponseBodyJobs {
	s.WakeMode = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobs) Validate() error {
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

type DescribePolarClawCronJobsResponseBodyJobsDelivery struct {
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

func (s DescribePolarClawCronJobsResponseBodyJobsDelivery) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawCronJobsResponseBodyJobsDelivery) GoString() string {
	return s.String()
}

func (s *DescribePolarClawCronJobsResponseBodyJobsDelivery) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribePolarClawCronJobsResponseBodyJobsDelivery) GetBestEffort() *bool {
	return s.BestEffort
}

func (s *DescribePolarClawCronJobsResponseBodyJobsDelivery) GetChannel() *string {
	return s.Channel
}

func (s *DescribePolarClawCronJobsResponseBodyJobsDelivery) GetMode() *string {
	return s.Mode
}

func (s *DescribePolarClawCronJobsResponseBodyJobsDelivery) GetTo() *string {
	return s.To
}

func (s *DescribePolarClawCronJobsResponseBodyJobsDelivery) SetAccountId(v string) *DescribePolarClawCronJobsResponseBodyJobsDelivery {
	s.AccountId = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsDelivery) SetBestEffort(v bool) *DescribePolarClawCronJobsResponseBodyJobsDelivery {
	s.BestEffort = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsDelivery) SetChannel(v string) *DescribePolarClawCronJobsResponseBodyJobsDelivery {
	s.Channel = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsDelivery) SetMode(v string) *DescribePolarClawCronJobsResponseBodyJobsDelivery {
	s.Mode = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsDelivery) SetTo(v string) *DescribePolarClawCronJobsResponseBodyJobsDelivery {
	s.To = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsDelivery) Validate() error {
	return dara.Validate(s)
}

type DescribePolarClawCronJobsResponseBodyJobsPayload struct {
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
	Deliver *bool `json:"Deliver,omitempty" xml:"Deliver,omitempty"`
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
	// 10
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	// example:
	//
	// ou_***
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s DescribePolarClawCronJobsResponseBodyJobsPayload) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawCronJobsResponseBodyJobsPayload) GoString() string {
	return s.String()
}

func (s *DescribePolarClawCronJobsResponseBodyJobsPayload) GetBestEffortDeliver() *bool {
	return s.BestEffortDeliver
}

func (s *DescribePolarClawCronJobsResponseBodyJobsPayload) GetChannel() *string {
	return s.Channel
}

func (s *DescribePolarClawCronJobsResponseBodyJobsPayload) GetDeliver() *bool {
	return s.Deliver
}

func (s *DescribePolarClawCronJobsResponseBodyJobsPayload) GetKind() *string {
	return s.Kind
}

func (s *DescribePolarClawCronJobsResponseBodyJobsPayload) GetLightContext() *bool {
	return s.LightContext
}

func (s *DescribePolarClawCronJobsResponseBodyJobsPayload) GetMessage() *string {
	return s.Message
}

func (s *DescribePolarClawCronJobsResponseBodyJobsPayload) GetModel() *string {
	return s.Model
}

func (s *DescribePolarClawCronJobsResponseBodyJobsPayload) GetText() *string {
	return s.Text
}

func (s *DescribePolarClawCronJobsResponseBodyJobsPayload) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *DescribePolarClawCronJobsResponseBodyJobsPayload) GetTo() *string {
	return s.To
}

func (s *DescribePolarClawCronJobsResponseBodyJobsPayload) SetBestEffortDeliver(v bool) *DescribePolarClawCronJobsResponseBodyJobsPayload {
	s.BestEffortDeliver = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsPayload) SetChannel(v string) *DescribePolarClawCronJobsResponseBodyJobsPayload {
	s.Channel = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsPayload) SetDeliver(v bool) *DescribePolarClawCronJobsResponseBodyJobsPayload {
	s.Deliver = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsPayload) SetKind(v string) *DescribePolarClawCronJobsResponseBodyJobsPayload {
	s.Kind = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsPayload) SetLightContext(v bool) *DescribePolarClawCronJobsResponseBodyJobsPayload {
	s.LightContext = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsPayload) SetMessage(v string) *DescribePolarClawCronJobsResponseBodyJobsPayload {
	s.Message = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsPayload) SetModel(v string) *DescribePolarClawCronJobsResponseBodyJobsPayload {
	s.Model = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsPayload) SetText(v string) *DescribePolarClawCronJobsResponseBodyJobsPayload {
	s.Text = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsPayload) SetTimeoutSeconds(v int32) *DescribePolarClawCronJobsResponseBodyJobsPayload {
	s.TimeoutSeconds = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsPayload) SetTo(v string) *DescribePolarClawCronJobsResponseBodyJobsPayload {
	s.To = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsPayload) Validate() error {
	return dara.Validate(s)
}

type DescribePolarClawCronJobsResponseBodyJobsRuns struct {
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
	Ts    *int64                                              `json:"Ts,omitempty" xml:"Ts,omitempty"`
	Usage *DescribePolarClawCronJobsResponseBodyJobsRunsUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s DescribePolarClawCronJobsResponseBodyJobsRuns) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawCronJobsResponseBodyJobsRuns) GoString() string {
	return s.String()
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) GetAction() *string {
	return s.Action
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) GetDelivered() *bool {
	return s.Delivered
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) GetDeliveryStatus() *string {
	return s.DeliveryStatus
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) GetDurationMs() *int64 {
	return s.DurationMs
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) GetJobId() *string {
	return s.JobId
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) GetJobName() *string {
	return s.JobName
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) GetModel() *string {
	return s.Model
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) GetNextRunAtMs() *int64 {
	return s.NextRunAtMs
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) GetProvider() *string {
	return s.Provider
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) GetRunAtMs() *int64 {
	return s.RunAtMs
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) GetStatus() *string {
	return s.Status
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) GetSummary() *string {
	return s.Summary
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) GetTs() *int64 {
	return s.Ts
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) GetUsage() *DescribePolarClawCronJobsResponseBodyJobsRunsUsage {
	return s.Usage
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) SetAction(v string) *DescribePolarClawCronJobsResponseBodyJobsRuns {
	s.Action = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) SetDelivered(v bool) *DescribePolarClawCronJobsResponseBodyJobsRuns {
	s.Delivered = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) SetDeliveryStatus(v string) *DescribePolarClawCronJobsResponseBodyJobsRuns {
	s.DeliveryStatus = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) SetDurationMs(v int64) *DescribePolarClawCronJobsResponseBodyJobsRuns {
	s.DurationMs = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) SetJobId(v string) *DescribePolarClawCronJobsResponseBodyJobsRuns {
	s.JobId = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) SetJobName(v string) *DescribePolarClawCronJobsResponseBodyJobsRuns {
	s.JobName = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) SetModel(v string) *DescribePolarClawCronJobsResponseBodyJobsRuns {
	s.Model = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) SetNextRunAtMs(v int64) *DescribePolarClawCronJobsResponseBodyJobsRuns {
	s.NextRunAtMs = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) SetProvider(v string) *DescribePolarClawCronJobsResponseBodyJobsRuns {
	s.Provider = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) SetRunAtMs(v int64) *DescribePolarClawCronJobsResponseBodyJobsRuns {
	s.RunAtMs = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) SetSessionId(v string) *DescribePolarClawCronJobsResponseBodyJobsRuns {
	s.SessionId = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) SetStatus(v string) *DescribePolarClawCronJobsResponseBodyJobsRuns {
	s.Status = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) SetSummary(v string) *DescribePolarClawCronJobsResponseBodyJobsRuns {
	s.Summary = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) SetTs(v int64) *DescribePolarClawCronJobsResponseBodyJobsRuns {
	s.Ts = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) SetUsage(v *DescribePolarClawCronJobsResponseBodyJobsRunsUsage) *DescribePolarClawCronJobsResponseBodyJobsRuns {
	s.Usage = v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRuns) Validate() error {
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePolarClawCronJobsResponseBodyJobsRunsUsage struct {
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

func (s DescribePolarClawCronJobsResponseBodyJobsRunsUsage) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawCronJobsResponseBodyJobsRunsUsage) GoString() string {
	return s.String()
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRunsUsage) GetInputTokens() *int32 {
	return s.InputTokens
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRunsUsage) GetOutputTokens() *int32 {
	return s.OutputTokens
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRunsUsage) GetTotalTokens() *int32 {
	return s.TotalTokens
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRunsUsage) SetInputTokens(v int32) *DescribePolarClawCronJobsResponseBodyJobsRunsUsage {
	s.InputTokens = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRunsUsage) SetOutputTokens(v int32) *DescribePolarClawCronJobsResponseBodyJobsRunsUsage {
	s.OutputTokens = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRunsUsage) SetTotalTokens(v int32) *DescribePolarClawCronJobsResponseBodyJobsRunsUsage {
	s.TotalTokens = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsRunsUsage) Validate() error {
	return dara.Validate(s)
}

type DescribePolarClawCronJobsResponseBodyJobsSchedule struct {
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
	// Asia/Shanghai
	Tz *string `json:"Tz,omitempty" xml:"Tz,omitempty"`
}

func (s DescribePolarClawCronJobsResponseBodyJobsSchedule) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawCronJobsResponseBodyJobsSchedule) GoString() string {
	return s.String()
}

func (s *DescribePolarClawCronJobsResponseBodyJobsSchedule) GetAnchorMs() *int64 {
	return s.AnchorMs
}

func (s *DescribePolarClawCronJobsResponseBodyJobsSchedule) GetAt() *string {
	return s.At
}

func (s *DescribePolarClawCronJobsResponseBodyJobsSchedule) GetEveryMs() *int64 {
	return s.EveryMs
}

func (s *DescribePolarClawCronJobsResponseBodyJobsSchedule) GetExpr() *string {
	return s.Expr
}

func (s *DescribePolarClawCronJobsResponseBodyJobsSchedule) GetKind() *string {
	return s.Kind
}

func (s *DescribePolarClawCronJobsResponseBodyJobsSchedule) GetTz() *string {
	return s.Tz
}

func (s *DescribePolarClawCronJobsResponseBodyJobsSchedule) SetAnchorMs(v int64) *DescribePolarClawCronJobsResponseBodyJobsSchedule {
	s.AnchorMs = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsSchedule) SetAt(v string) *DescribePolarClawCronJobsResponseBodyJobsSchedule {
	s.At = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsSchedule) SetEveryMs(v int64) *DescribePolarClawCronJobsResponseBodyJobsSchedule {
	s.EveryMs = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsSchedule) SetExpr(v string) *DescribePolarClawCronJobsResponseBodyJobsSchedule {
	s.Expr = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsSchedule) SetKind(v string) *DescribePolarClawCronJobsResponseBodyJobsSchedule {
	s.Kind = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsSchedule) SetTz(v string) *DescribePolarClawCronJobsResponseBodyJobsSchedule {
	s.Tz = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsSchedule) Validate() error {
	return dara.Validate(s)
}

type DescribePolarClawCronJobsResponseBodyJobsState struct {
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

func (s DescribePolarClawCronJobsResponseBodyJobsState) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawCronJobsResponseBodyJobsState) GoString() string {
	return s.String()
}

func (s *DescribePolarClawCronJobsResponseBodyJobsState) GetConsecutiveErrors() *int32 {
	return s.ConsecutiveErrors
}

func (s *DescribePolarClawCronJobsResponseBodyJobsState) GetLastRunAtMs() *int64 {
	return s.LastRunAtMs
}

func (s *DescribePolarClawCronJobsResponseBodyJobsState) GetLastRunStatus() *string {
	return s.LastRunStatus
}

func (s *DescribePolarClawCronJobsResponseBodyJobsState) GetNextRunAtMs() *int64 {
	return s.NextRunAtMs
}

func (s *DescribePolarClawCronJobsResponseBodyJobsState) SetConsecutiveErrors(v int32) *DescribePolarClawCronJobsResponseBodyJobsState {
	s.ConsecutiveErrors = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsState) SetLastRunAtMs(v int64) *DescribePolarClawCronJobsResponseBodyJobsState {
	s.LastRunAtMs = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsState) SetLastRunStatus(v string) *DescribePolarClawCronJobsResponseBodyJobsState {
	s.LastRunStatus = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsState) SetNextRunAtMs(v int64) *DescribePolarClawCronJobsResponseBodyJobsState {
	s.NextRunAtMs = &v
	return s
}

func (s *DescribePolarClawCronJobsResponseBodyJobsState) Validate() error {
	return dara.Validate(s)
}
