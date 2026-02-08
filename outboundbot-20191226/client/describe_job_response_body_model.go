// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeJobResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DescribeJobResponseBody
	GetHttpStatusCode() *int32
	SetJob(v *DescribeJobResponseBodyJob) *DescribeJobResponseBody
	GetJob() *DescribeJobResponseBodyJob
	SetMessage(v string) *DescribeJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeJobResponseBody
	GetSuccess() *bool
}

type DescribeJobResponseBody struct {
	// API status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Job information.
	//
	// example:
	//
	// {}
	Job *DescribeJobResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// API prompt message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeJobResponseBody) GetJob() *DescribeJobResponseBodyJob {
	return s.Job
}

func (s *DescribeJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeJobResponseBody) SetCode(v string) *DescribeJobResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeJobResponseBody) SetHttpStatusCode(v int32) *DescribeJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeJobResponseBody) SetJob(v *DescribeJobResponseBodyJob) *DescribeJobResponseBody {
	s.Job = v
	return s
}

func (s *DescribeJobResponseBody) SetMessage(v string) *DescribeJobResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeJobResponseBody) SetRequestId(v string) *DescribeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJobResponseBody) SetSuccess(v bool) *DescribeJobResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeJobResponseBody) Validate() error {
	if s.Job != nil {
		if err := s.Job.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeJobResponseBodyJob struct {
	// Actual running time of the job. [Deprecated]
	//
	// example:
	//
	// 1640068026385
	ActualTime *int64 `json:"ActualTime,omitempty" xml:"ActualTime,omitempty"`
	// Called number.
	//
	// example:
	//
	// 137****7777
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// List of calling numbers.
	//
	// example:
	//
	// ["057126883106"]
	CallingNumbers []*string `json:"CallingNumbers,omitempty" xml:"CallingNumbers,omitempty" type:"Repeated"`
	// Contact information.
	//
	// example:
	//
	// []
	Contacts []*DescribeJobResponseBodyJobContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Repeated"`
	// Job hit label status. [Deprecated]
	//
	// example:
	//
	// -
	DsReport *string `json:"DsReport,omitempty" xml:"DsReport,omitempty"`
	// Call termination reason. Valid values:
	//
	// - **1**: FINISHED (Normal completion).
	//
	// - **2**: CHATBOT_HANGUP_AFTER_NOANSWER (Bot hangs up after unrecognized input).
	//
	// - **3**: CHATBOT_HANGUP_AFTER_SILENCE (Bot hangs up due to silence timeout).
	//
	// - **4**: CLIENT_HANGUP_AFTER_NOANSWER (User hangs up after unrecognized input).
	//
	// - **5**: CLIENT_HANGUP (User hangs up without reason).
	//
	// - **6**: TRANSFER_BY_INTENT (Transferred to agent due to intent match).
	//
	// - **7**: TRANSFER_AFTER_NOANSWER (Transferred to agent after unrecognized input).
	//
	// - **8**: NO_INTERACTION (No interaction from user side).
	//
	// - **9**: ERROR (System exception interruption).
	//
	// example:
	//
	// 1
	EndReason *int32 `json:"EndReason,omitempty" xml:"EndReason,omitempty"`
	// Business parameters.
	//
	// > TenantId and ServiceId are system-generated. All other parameters are custom and passed in as references.
	//
	// example:
	//
	// []
	Extras []*DescribeJobResponseBodyJobExtras `json:"Extras,omitempty" xml:"Extras,omitempty" type:"Repeated"`
	// Reason for job failure.
	//
	// example:
	//
	// NoAnswer
	FailureReason *string `json:"FailureReason,omitempty" xml:"FailureReason,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// d5971d98-7312-4f0e-a918-a17d67133e28
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Job ID.
	//
	// example:
	//
	// fce6c599-8ede-40e3-9f78-0928eda7b4e8
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// Job ID.
	//
	// example:
	//
	// b72425bd-7871-4050-838e-033d80d754b7
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The time when the job will be executed next.
	//
	// example:
	//
	// 1640068026385
	NextExecutionTime *int64 `json:"NextExecutionTime,omitempty" xml:"NextExecutionTime,omitempty"`
	// Job priority.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Business ID of the job, defined by the business party.
	//
	// > This is the uploaded ContactId value.
	//
	// example:
	//
	// d5971d98-7312-4f0e-a918-a17d67133e28
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// Scenario ID.
	//
	// example:
	//
	// ade80092-03d9-4f4d-ad4f-ab8a247d3150
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// Script scenario.
	Script *DescribeJobResponseBodyJobScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// Job status.
	//
	// - Scheduling,
	//
	// - Executing,
	//
	// - Succeeded,
	//
	// - Paused,
	//
	// - Failed,
	//
	// - Cancelled,
	//
	// - Drafted (Draft state, a temporary state during file import);
	//
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Policy ID.
	//
	// example:
	//
	// c8a2b7f2-ad1a-4865-b872-d0080d9802d9
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// Label data for the conversation business:
	//
	// - In LLM scenarios: label hit data after the conversation ends.
	//
	// - In small model scenarios: variable value data after the conversation ends.
	//
	// example:
	//
	// []
	Summary []*DescribeJobResponseBodyJobSummary `json:"Summary,omitempty" xml:"Summary,omitempty" type:"Repeated"`
	// System priority of the job.
	//
	// example:
	//
	// 1
	SystemPriority *int32 `json:"SystemPriority,omitempty" xml:"SystemPriority,omitempty"`
	// Call List.
	//
	// example:
	//
	// []
	Tasks []*DescribeJobResponseBodyJobTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s DescribeJobResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResponseBodyJob) GoString() string {
	return s.String()
}

func (s *DescribeJobResponseBodyJob) GetActualTime() *int64 {
	return s.ActualTime
}

func (s *DescribeJobResponseBodyJob) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *DescribeJobResponseBodyJob) GetCallingNumbers() []*string {
	return s.CallingNumbers
}

func (s *DescribeJobResponseBodyJob) GetContacts() []*DescribeJobResponseBodyJobContacts {
	return s.Contacts
}

func (s *DescribeJobResponseBodyJob) GetDsReport() *string {
	return s.DsReport
}

func (s *DescribeJobResponseBodyJob) GetEndReason() *int32 {
	return s.EndReason
}

func (s *DescribeJobResponseBodyJob) GetExtras() []*DescribeJobResponseBodyJobExtras {
	return s.Extras
}

func (s *DescribeJobResponseBodyJob) GetFailureReason() *string {
	return s.FailureReason
}

func (s *DescribeJobResponseBodyJob) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeJobResponseBodyJob) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *DescribeJobResponseBodyJob) GetJobId() *string {
	return s.JobId
}

func (s *DescribeJobResponseBodyJob) GetNextExecutionTime() *int64 {
	return s.NextExecutionTime
}

func (s *DescribeJobResponseBodyJob) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeJobResponseBodyJob) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *DescribeJobResponseBodyJob) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *DescribeJobResponseBodyJob) GetScript() *DescribeJobResponseBodyJobScript {
	return s.Script
}

func (s *DescribeJobResponseBodyJob) GetStatus() *string {
	return s.Status
}

func (s *DescribeJobResponseBodyJob) GetStrategyId() *string {
	return s.StrategyId
}

func (s *DescribeJobResponseBodyJob) GetSummary() []*DescribeJobResponseBodyJobSummary {
	return s.Summary
}

func (s *DescribeJobResponseBodyJob) GetSystemPriority() *int32 {
	return s.SystemPriority
}

func (s *DescribeJobResponseBodyJob) GetTasks() []*DescribeJobResponseBodyJobTasks {
	return s.Tasks
}

func (s *DescribeJobResponseBodyJob) SetActualTime(v int64) *DescribeJobResponseBodyJob {
	s.ActualTime = &v
	return s
}

func (s *DescribeJobResponseBodyJob) SetCalledNumber(v string) *DescribeJobResponseBodyJob {
	s.CalledNumber = &v
	return s
}

func (s *DescribeJobResponseBodyJob) SetCallingNumbers(v []*string) *DescribeJobResponseBodyJob {
	s.CallingNumbers = v
	return s
}

func (s *DescribeJobResponseBodyJob) SetContacts(v []*DescribeJobResponseBodyJobContacts) *DescribeJobResponseBodyJob {
	s.Contacts = v
	return s
}

func (s *DescribeJobResponseBodyJob) SetDsReport(v string) *DescribeJobResponseBodyJob {
	s.DsReport = &v
	return s
}

func (s *DescribeJobResponseBodyJob) SetEndReason(v int32) *DescribeJobResponseBodyJob {
	s.EndReason = &v
	return s
}

func (s *DescribeJobResponseBodyJob) SetExtras(v []*DescribeJobResponseBodyJobExtras) *DescribeJobResponseBodyJob {
	s.Extras = v
	return s
}

func (s *DescribeJobResponseBodyJob) SetFailureReason(v string) *DescribeJobResponseBodyJob {
	s.FailureReason = &v
	return s
}

func (s *DescribeJobResponseBodyJob) SetInstanceId(v string) *DescribeJobResponseBodyJob {
	s.InstanceId = &v
	return s
}

func (s *DescribeJobResponseBodyJob) SetJobGroupId(v string) *DescribeJobResponseBodyJob {
	s.JobGroupId = &v
	return s
}

func (s *DescribeJobResponseBodyJob) SetJobId(v string) *DescribeJobResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *DescribeJobResponseBodyJob) SetNextExecutionTime(v int64) *DescribeJobResponseBodyJob {
	s.NextExecutionTime = &v
	return s
}

func (s *DescribeJobResponseBodyJob) SetPriority(v int32) *DescribeJobResponseBodyJob {
	s.Priority = &v
	return s
}

func (s *DescribeJobResponseBodyJob) SetReferenceId(v string) *DescribeJobResponseBodyJob {
	s.ReferenceId = &v
	return s
}

func (s *DescribeJobResponseBodyJob) SetScenarioId(v string) *DescribeJobResponseBodyJob {
	s.ScenarioId = &v
	return s
}

func (s *DescribeJobResponseBodyJob) SetScript(v *DescribeJobResponseBodyJobScript) *DescribeJobResponseBodyJob {
	s.Script = v
	return s
}

func (s *DescribeJobResponseBodyJob) SetStatus(v string) *DescribeJobResponseBodyJob {
	s.Status = &v
	return s
}

func (s *DescribeJobResponseBodyJob) SetStrategyId(v string) *DescribeJobResponseBodyJob {
	s.StrategyId = &v
	return s
}

func (s *DescribeJobResponseBodyJob) SetSummary(v []*DescribeJobResponseBodyJobSummary) *DescribeJobResponseBodyJob {
	s.Summary = v
	return s
}

func (s *DescribeJobResponseBodyJob) SetSystemPriority(v int32) *DescribeJobResponseBodyJob {
	s.SystemPriority = &v
	return s
}

func (s *DescribeJobResponseBodyJob) SetTasks(v []*DescribeJobResponseBodyJobTasks) *DescribeJobResponseBodyJob {
	s.Tasks = v
	return s
}

func (s *DescribeJobResponseBodyJob) Validate() error {
	if s.Contacts != nil {
		for _, item := range s.Contacts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Extras != nil {
		for _, item := range s.Extras {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Script != nil {
		if err := s.Script.Validate(); err != nil {
			return err
		}
	}
	if s.Summary != nil {
		for _, item := range s.Summary {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeJobResponseBodyJobContacts struct {
	// Contact ID. System-generated.
	//
	// example:
	//
	// db3db762-e421-44c9-9a01-cb423470757c
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// Contact name.
	//
	// example:
	//
	// 张三
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// Contact honorific. Consistent with the contact name.
	//
	// example:
	//
	// 张先生
	Honorific *string `json:"Honorific,omitempty" xml:"Honorific,omitempty"`
	// Job ID.
	//
	// example:
	//
	// 72dcd26b-f12d-4c27-b3af-18f6aed5b160
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Phone number.
	//
	// example:
	//
	// 1358****8888
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The business ID of the contact.
	//
	// example:
	//
	// 2fa6bac3-06da-4315-82ab-72d6fd3a6f34
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// Contact role. [Deprecated]
	//
	// example:
	//
	// *
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// Contact status. [Deprecated]
	//
	// example:
	//
	// Available
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeJobResponseBodyJobContacts) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResponseBodyJobContacts) GoString() string {
	return s.String()
}

func (s *DescribeJobResponseBodyJobContacts) GetContactId() *string {
	return s.ContactId
}

func (s *DescribeJobResponseBodyJobContacts) GetContactName() *string {
	return s.ContactName
}

func (s *DescribeJobResponseBodyJobContacts) GetHonorific() *string {
	return s.Honorific
}

func (s *DescribeJobResponseBodyJobContacts) GetJobId() *string {
	return s.JobId
}

func (s *DescribeJobResponseBodyJobContacts) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *DescribeJobResponseBodyJobContacts) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *DescribeJobResponseBodyJobContacts) GetRole() *string {
	return s.Role
}

func (s *DescribeJobResponseBodyJobContacts) GetState() *string {
	return s.State
}

func (s *DescribeJobResponseBodyJobContacts) SetContactId(v string) *DescribeJobResponseBodyJobContacts {
	s.ContactId = &v
	return s
}

func (s *DescribeJobResponseBodyJobContacts) SetContactName(v string) *DescribeJobResponseBodyJobContacts {
	s.ContactName = &v
	return s
}

func (s *DescribeJobResponseBodyJobContacts) SetHonorific(v string) *DescribeJobResponseBodyJobContacts {
	s.Honorific = &v
	return s
}

func (s *DescribeJobResponseBodyJobContacts) SetJobId(v string) *DescribeJobResponseBodyJobContacts {
	s.JobId = &v
	return s
}

func (s *DescribeJobResponseBodyJobContacts) SetPhoneNumber(v string) *DescribeJobResponseBodyJobContacts {
	s.PhoneNumber = &v
	return s
}

func (s *DescribeJobResponseBodyJobContacts) SetReferenceId(v string) *DescribeJobResponseBodyJobContacts {
	s.ReferenceId = &v
	return s
}

func (s *DescribeJobResponseBodyJobContacts) SetRole(v string) *DescribeJobResponseBodyJobContacts {
	s.Role = &v
	return s
}

func (s *DescribeJobResponseBodyJobContacts) SetState(v string) *DescribeJobResponseBodyJobContacts {
	s.State = &v
	return s
}

func (s *DescribeJobResponseBodyJobContacts) Validate() error {
	return dara.Validate(s)
}

type DescribeJobResponseBodyJobExtras struct {
	// Business parameter name.
	//
	// example:
	//
	// djrq
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Business parameter value.
	//
	// example:
	//
	// 2019-08-21 09:49:59.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeJobResponseBodyJobExtras) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResponseBodyJobExtras) GoString() string {
	return s.String()
}

func (s *DescribeJobResponseBodyJobExtras) GetKey() *string {
	return s.Key
}

func (s *DescribeJobResponseBodyJobExtras) GetValue() *string {
	return s.Value
}

func (s *DescribeJobResponseBodyJobExtras) SetKey(v string) *DescribeJobResponseBodyJobExtras {
	s.Key = &v
	return s
}

func (s *DescribeJobResponseBodyJobExtras) SetValue(v string) *DescribeJobResponseBodyJobExtras {
	s.Value = &v
	return s
}

func (s *DescribeJobResponseBodyJobExtras) Validate() error {
	return dara.Validate(s)
}

type DescribeJobResponseBodyJobScript struct {
	// ASR configuration for the script. [Deprecated]
	//
	// example:
	//
	// {\\"AppKey\\":\\"3GHttnsvir1FeWWb\\"}
	AsrConfig *string `json:"AsrConfig,omitempty" xml:"AsrConfig,omitempty"`
	// Chatbot ID.
	//
	// example:
	//
	// chatbot-cn-EJfqqa***
	ChatbotId *string `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	// Debug status.
	//
	// example:
	//
	// DRAFTED
	DebugStatus *string `json:"DebugStatus,omitempty" xml:"DebugStatus,omitempty"`
	// Industry.
	//
	// example:
	//
	// 金融
	Industry *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	// Indicates whether the debug version is in Draft status.
	//
	// example:
	//
	// true
	IsDebugDrafted *bool `json:"IsDebugDrafted,omitempty" xml:"IsDebugDrafted,omitempty"`
	// Indicates whether the script is in Draft status.
	//
	// example:
	//
	// true
	IsDrafted *bool `json:"IsDrafted,omitempty" xml:"IsDrafted,omitempty"`
	// Toggle for tone continuity feature. No response [Deprecated]
	//
	// example:
	//
	// true
	MiniPlaybackConfigEnabled *bool `json:"MiniPlaybackConfigEnabled,omitempty" xml:"MiniPlaybackConfigEnabled,omitempty"`
	// Script name.
	//
	// example:
	//
	// 催收话术
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Scenario.
	//
	// example:
	//
	// 催收
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// Script description.
	//
	// example:
	//
	// 催收话术
	ScriptDescription *string `json:"ScriptDescription,omitempty" xml:"ScriptDescription,omitempty"`
	// Script ID.
	//
	// example:
	//
	// 810b5872-57f0-4b27-80ab-7b3f4d8a6374
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Status.
	//
	// example:
	//
	// DRAFTED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// TTS configuration for the script. [Deprecated]
	//
	// example:
	//
	// {\\"voice\\":\\"xiaobei\\",\\"volume\\":\\"50\\",\\"speechRate\\":\\"-150\\",\\"pitchRate\\":\\"0\\"}
	TtsConfig *string `json:"TtsConfig,omitempty" xml:"TtsConfig,omitempty"`
	// Update Time.
	//
	// example:
	//
	// 1578881227000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeJobResponseBodyJobScript) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResponseBodyJobScript) GoString() string {
	return s.String()
}

func (s *DescribeJobResponseBodyJobScript) GetAsrConfig() *string {
	return s.AsrConfig
}

func (s *DescribeJobResponseBodyJobScript) GetChatbotId() *string {
	return s.ChatbotId
}

func (s *DescribeJobResponseBodyJobScript) GetDebugStatus() *string {
	return s.DebugStatus
}

func (s *DescribeJobResponseBodyJobScript) GetIndustry() *string {
	return s.Industry
}

func (s *DescribeJobResponseBodyJobScript) GetIsDebugDrafted() *bool {
	return s.IsDebugDrafted
}

func (s *DescribeJobResponseBodyJobScript) GetIsDrafted() *bool {
	return s.IsDrafted
}

func (s *DescribeJobResponseBodyJobScript) GetMiniPlaybackConfigEnabled() *bool {
	return s.MiniPlaybackConfigEnabled
}

func (s *DescribeJobResponseBodyJobScript) GetName() *string {
	return s.Name
}

func (s *DescribeJobResponseBodyJobScript) GetScene() *string {
	return s.Scene
}

func (s *DescribeJobResponseBodyJobScript) GetScriptDescription() *string {
	return s.ScriptDescription
}

func (s *DescribeJobResponseBodyJobScript) GetScriptId() *string {
	return s.ScriptId
}

func (s *DescribeJobResponseBodyJobScript) GetStatus() *string {
	return s.Status
}

func (s *DescribeJobResponseBodyJobScript) GetTtsConfig() *string {
	return s.TtsConfig
}

func (s *DescribeJobResponseBodyJobScript) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeJobResponseBodyJobScript) SetAsrConfig(v string) *DescribeJobResponseBodyJobScript {
	s.AsrConfig = &v
	return s
}

func (s *DescribeJobResponseBodyJobScript) SetChatbotId(v string) *DescribeJobResponseBodyJobScript {
	s.ChatbotId = &v
	return s
}

func (s *DescribeJobResponseBodyJobScript) SetDebugStatus(v string) *DescribeJobResponseBodyJobScript {
	s.DebugStatus = &v
	return s
}

func (s *DescribeJobResponseBodyJobScript) SetIndustry(v string) *DescribeJobResponseBodyJobScript {
	s.Industry = &v
	return s
}

func (s *DescribeJobResponseBodyJobScript) SetIsDebugDrafted(v bool) *DescribeJobResponseBodyJobScript {
	s.IsDebugDrafted = &v
	return s
}

func (s *DescribeJobResponseBodyJobScript) SetIsDrafted(v bool) *DescribeJobResponseBodyJobScript {
	s.IsDrafted = &v
	return s
}

func (s *DescribeJobResponseBodyJobScript) SetMiniPlaybackConfigEnabled(v bool) *DescribeJobResponseBodyJobScript {
	s.MiniPlaybackConfigEnabled = &v
	return s
}

func (s *DescribeJobResponseBodyJobScript) SetName(v string) *DescribeJobResponseBodyJobScript {
	s.Name = &v
	return s
}

func (s *DescribeJobResponseBodyJobScript) SetScene(v string) *DescribeJobResponseBodyJobScript {
	s.Scene = &v
	return s
}

func (s *DescribeJobResponseBodyJobScript) SetScriptDescription(v string) *DescribeJobResponseBodyJobScript {
	s.ScriptDescription = &v
	return s
}

func (s *DescribeJobResponseBodyJobScript) SetScriptId(v string) *DescribeJobResponseBodyJobScript {
	s.ScriptId = &v
	return s
}

func (s *DescribeJobResponseBodyJobScript) SetStatus(v string) *DescribeJobResponseBodyJobScript {
	s.Status = &v
	return s
}

func (s *DescribeJobResponseBodyJobScript) SetTtsConfig(v string) *DescribeJobResponseBodyJobScript {
	s.TtsConfig = &v
	return s
}

func (s *DescribeJobResponseBodyJobScript) SetUpdateTime(v int64) *DescribeJobResponseBodyJobScript {
	s.UpdateTime = &v
	return s
}

func (s *DescribeJobResponseBodyJobScript) Validate() error {
	return dara.Validate(s)
}

type DescribeJobResponseBodyJobSummary struct {
	// Category of the conversation summary (historical field, no longer used) [Deprecated]
	//
	// example:
	//
	// {}
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Tag value
	//
	// example:
	//
	// 5
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Tag name.
	//
	// example:
	//
	// score
	SummaryName *string `json:"SummaryName,omitempty" xml:"SummaryName,omitempty"`
}

func (s DescribeJobResponseBodyJobSummary) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResponseBodyJobSummary) GoString() string {
	return s.String()
}

func (s *DescribeJobResponseBodyJobSummary) GetCategory() *string {
	return s.Category
}

func (s *DescribeJobResponseBodyJobSummary) GetContent() *string {
	return s.Content
}

func (s *DescribeJobResponseBodyJobSummary) GetSummaryName() *string {
	return s.SummaryName
}

func (s *DescribeJobResponseBodyJobSummary) SetCategory(v string) *DescribeJobResponseBodyJobSummary {
	s.Category = &v
	return s
}

func (s *DescribeJobResponseBodyJobSummary) SetContent(v string) *DescribeJobResponseBodyJobSummary {
	s.Content = &v
	return s
}

func (s *DescribeJobResponseBodyJobSummary) SetSummaryName(v string) *DescribeJobResponseBodyJobSummary {
	s.SummaryName = &v
	return s
}

func (s *DescribeJobResponseBodyJobSummary) Validate() error {
	return dara.Validate(s)
}

type DescribeJobResponseBodyJobTasks struct {
	// Actual running time.
	//
	// example:
	//
	// 1579068424883
	ActualTime *int64 `json:"ActualTime,omitempty" xml:"ActualTime,omitempty"`
	// Business result (historical field, no longer used).
	//
	// example:
	//
	// 1
	Brief *string `json:"Brief,omitempty" xml:"Brief,omitempty"`
	// SIP call ID.
	//
	// example:
	//
	// 1528189846043
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// Called number.
	//
	// example:
	//
	// 135****8888
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// Calling number.
	//
	// example:
	//
	// 0571****3106
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// ID of the chatbot used for the conversation.
	//
	// example:
	//
	// 1234
	ChatbotId *string `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	// Contact information.
	//
	// example:
	//
	// {}
	Contact *DescribeJobResponseBodyJobTasksContact `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Struct"`
	// List of conversation texts for the job.
	//
	// example:
	//
	// []
	Conversation []*DescribeJobResponseBodyJobTasksConversation `json:"Conversation,omitempty" xml:"Conversation,omitempty" type:"Repeated"`
	// Conversation duration. This field is not returned if unavailable.
	//
	// > This data is returned in the ActionParams parameter of the last object in the Conversation list.
	//
	// example:
	//
	// 120
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Completion reason.
	//
	// example:
	//
	// FINISHED
	EndReason *string `json:"EndReason,omitempty" xml:"EndReason,omitempty"`
	// Actual end time.
	//
	// example:
	//
	// 1579068424883
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Party that hung up the call.
	//
	// - client: Customer
	//
	// - ivr: Robot
	//
	// example:
	//
	// client
	HangUpDirection *string `json:"HangUpDirection,omitempty" xml:"HangUpDirection,omitempty"`
	// Job ID.
	//
	// example:
	//
	// b72425bd-7871-4050-838e-033d80d754b7
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Scheduled calling time.
	//
	// example:
	//
	// 1579068424883
	PlanedTime *int64 `json:"PlanedTime,omitempty" xml:"PlanedTime,omitempty"`
	// Actual ringing duration.
	//
	// example:
	//
	// 25
	RealRingingDuration *int64 `json:"RealRingingDuration,omitempty" xml:"RealRingingDuration,omitempty"`
	// Ringing duration.
	//
	// example:
	//
	// 25
	RingingDuration *int64 `json:"RingingDuration,omitempty" xml:"RingingDuration,omitempty"`
	// Scenario ID.
	//
	// example:
	//
	// ade80092-03d9-4f4d-ad4f-ab8a247d3150
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// SIP status code of the call job.
	//
	// example:
	//
	// 200
	SipCode *string `json:"SipCode,omitempty" xml:"SipCode,omitempty"`
	// SIP signaling duration. [Deprecated]
	//
	// example:
	//
	// 25
	SipDuration *int64 `json:"SipDuration,omitempty" xml:"SipDuration,omitempty"`
	// Task Status, with the following possible values:
	//
	// (Note: The **Succeeded*	- status has been further categorized by specific reasons. The generic **Succeeded**: 1 (Connected) status will no longer be displayed; instead, only the detailed reason types will be returned.)
	//
	// - **Executing**: 0 (Calling in progress).
	//
	// - **Succeeded**: 1 (Connected).
	//
	// - **NoAnswer**: 2 (Not connected – No answer).
	//
	// - **NotExist**: 3 (Not connected – Nonexistent number).
	//
	// - **Busy**: 4 (Not connected – Busy line).
	//
	// - **Cancelled**: 5 (Not dialed – Job stopped).
	//
	// - **Failed**: 6 (Failed).
	//
	// - **NotConnected**: 7 (Not connected – Unable to reach).
	//
	// - **PoweredOff**: 8 (Not connected – Powered off).
	//
	// - **OutOfService**: 9 (Not connected – Called party suspended).
	//
	// - **InArrears**: 10 (Not connected – Called party has overdue payment).
	//
	// - **EmptyNumber**: 11 (Not dialed – Nonexistent number, not called).
	//
	// - **PerDayCallCountLimit**: 12 (Not dialed – Exceeded daily call limit).
	//
	// - **ContactBlockList**: 13 (Not dialed – Blacklist).
	//
	// - **CallerNotRegistered**: 14 (Not dialed – Caller number not registered).
	//
	// - **Terminated**: 15 (Not dialed – Terminated).
	//
	// - **VerificationCancelled**: 16 (Not dialed – Cancelled due to failed pre-call authentication).
	//
	// - **OutOfServiceNoCall**: 17 (Not dialed – Called party suspended, not called).
	//
	// - **InArrearsNoCall**: 18 (Not dialed – Called party has overdue payment, not called).
	//
	// - **CallingNumberNotExist**: 19 (Not dialed – Caller number does not exist).
	//
	// - **SucceededFinish**: 20 (Connected – Normal completion).
	//
	// - **SucceededChatbotHangUpAfterNoAnswer**: 21 (Connected – Bot hung up after unrecognized input).
	//
	// - **SucceededChatbotHangUpAfterSilence**: 22 (Connected – Bot hung up after silence timeout).
	//
	// - **SucceededClientHangUpAfterNoAnswer**: 23 (Connected – User hung up after unrecognized input).
	//
	// - **SucceededClientHangUp**: 24 (Connected – User hung up without reason).
	//
	// - **SucceededTransferByIntent**: 25 (Connected – Transferred to agent due to intent match).
	//
	// - **SucceededTransferAfterNoAnswer**: 26 (Connected – Transferred to agent after unrecognized input).
	//
	// - **SucceededInoInterAction**: 27 (Connected – No interaction from user side).
	//
	// - **SucceededError**: 28 (Connected – Interrupted by system exception).
	//
	// - **SucceededSpecialInterceptVoiceAssistant**: 29 (Connected – Intercepted under special condition – Voice assistant).
	//
	// - **SucceededSpecialInterceptExtensionNumberTransfer**: 30 (Connected – Intercepted under special condition – Extension transfer).
	//
	// - **SucceededSpecialInterceptCustomSpecialIntercept**: 31 (Connected – Intercepted under special condition – Custom intercept).
	//
	// - **HighRiskSipCode**: 32 (Not dialed – High-risk, not called).
	//
	// example:
	//
	// SucceededTransferByIntent
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Call ID.
	//
	// example:
	//
	// ff44709e-39a6-43ba-959b-20fcabe3e496
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeJobResponseBodyJobTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResponseBodyJobTasks) GoString() string {
	return s.String()
}

func (s *DescribeJobResponseBodyJobTasks) GetActualTime() *int64 {
	return s.ActualTime
}

func (s *DescribeJobResponseBodyJobTasks) GetBrief() *string {
	return s.Brief
}

func (s *DescribeJobResponseBodyJobTasks) GetCallId() *string {
	return s.CallId
}

func (s *DescribeJobResponseBodyJobTasks) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *DescribeJobResponseBodyJobTasks) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *DescribeJobResponseBodyJobTasks) GetChatbotId() *string {
	return s.ChatbotId
}

func (s *DescribeJobResponseBodyJobTasks) GetContact() *DescribeJobResponseBodyJobTasksContact {
	return s.Contact
}

func (s *DescribeJobResponseBodyJobTasks) GetConversation() []*DescribeJobResponseBodyJobTasksConversation {
	return s.Conversation
}

func (s *DescribeJobResponseBodyJobTasks) GetDuration() *int32 {
	return s.Duration
}

func (s *DescribeJobResponseBodyJobTasks) GetEndReason() *string {
	return s.EndReason
}

func (s *DescribeJobResponseBodyJobTasks) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeJobResponseBodyJobTasks) GetHangUpDirection() *string {
	return s.HangUpDirection
}

func (s *DescribeJobResponseBodyJobTasks) GetJobId() *string {
	return s.JobId
}

func (s *DescribeJobResponseBodyJobTasks) GetPlanedTime() *int64 {
	return s.PlanedTime
}

func (s *DescribeJobResponseBodyJobTasks) GetRealRingingDuration() *int64 {
	return s.RealRingingDuration
}

func (s *DescribeJobResponseBodyJobTasks) GetRingingDuration() *int64 {
	return s.RingingDuration
}

func (s *DescribeJobResponseBodyJobTasks) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *DescribeJobResponseBodyJobTasks) GetSipCode() *string {
	return s.SipCode
}

func (s *DescribeJobResponseBodyJobTasks) GetSipDuration() *int64 {
	return s.SipDuration
}

func (s *DescribeJobResponseBodyJobTasks) GetStatus() *string {
	return s.Status
}

func (s *DescribeJobResponseBodyJobTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeJobResponseBodyJobTasks) SetActualTime(v int64) *DescribeJobResponseBodyJobTasks {
	s.ActualTime = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasks) SetBrief(v string) *DescribeJobResponseBodyJobTasks {
	s.Brief = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasks) SetCallId(v string) *DescribeJobResponseBodyJobTasks {
	s.CallId = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasks) SetCalledNumber(v string) *DescribeJobResponseBodyJobTasks {
	s.CalledNumber = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasks) SetCallingNumber(v string) *DescribeJobResponseBodyJobTasks {
	s.CallingNumber = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasks) SetChatbotId(v string) *DescribeJobResponseBodyJobTasks {
	s.ChatbotId = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasks) SetContact(v *DescribeJobResponseBodyJobTasksContact) *DescribeJobResponseBodyJobTasks {
	s.Contact = v
	return s
}

func (s *DescribeJobResponseBodyJobTasks) SetConversation(v []*DescribeJobResponseBodyJobTasksConversation) *DescribeJobResponseBodyJobTasks {
	s.Conversation = v
	return s
}

func (s *DescribeJobResponseBodyJobTasks) SetDuration(v int32) *DescribeJobResponseBodyJobTasks {
	s.Duration = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasks) SetEndReason(v string) *DescribeJobResponseBodyJobTasks {
	s.EndReason = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasks) SetEndTime(v int64) *DescribeJobResponseBodyJobTasks {
	s.EndTime = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasks) SetHangUpDirection(v string) *DescribeJobResponseBodyJobTasks {
	s.HangUpDirection = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasks) SetJobId(v string) *DescribeJobResponseBodyJobTasks {
	s.JobId = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasks) SetPlanedTime(v int64) *DescribeJobResponseBodyJobTasks {
	s.PlanedTime = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasks) SetRealRingingDuration(v int64) *DescribeJobResponseBodyJobTasks {
	s.RealRingingDuration = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasks) SetRingingDuration(v int64) *DescribeJobResponseBodyJobTasks {
	s.RingingDuration = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasks) SetScenarioId(v string) *DescribeJobResponseBodyJobTasks {
	s.ScenarioId = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasks) SetSipCode(v string) *DescribeJobResponseBodyJobTasks {
	s.SipCode = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasks) SetSipDuration(v int64) *DescribeJobResponseBodyJobTasks {
	s.SipDuration = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasks) SetStatus(v string) *DescribeJobResponseBodyJobTasks {
	s.Status = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasks) SetTaskId(v string) *DescribeJobResponseBodyJobTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasks) Validate() error {
	if s.Contact != nil {
		if err := s.Contact.Validate(); err != nil {
			return err
		}
	}
	if s.Conversation != nil {
		for _, item := range s.Conversation {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeJobResponseBodyJobTasksContact struct {
	// Contact ID. Generated by the system.
	//
	// example:
	//
	// db3db762-e421-44c9-9a01-cb423470757c
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// Contact name
	//
	// example:
	//
	// 张三
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// Contact honorific. Must be consistent with the contact\\"s name.
	//
	// example:
	//
	// 张先生
	Honorific *string `json:"Honorific,omitempty" xml:"Honorific,omitempty"`
	// Job ID [Deprecated]
	//
	// example:
	//
	// b72425bd-7871-4050-838e-033d80d754b7
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Phone number.
	//
	// example:
	//
	// 1351****8888
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Operational System ID of the contact.
	//
	// > Custom uploaded by the customer via the ContactId parameter.
	//
	// example:
	//
	// 2fa6bac3-06da-4315-82ab-72d6fd3a6f34
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// Contact role [deprecated].
	//
	// example:
	//
	// *
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// Contact status. [Deprecated]
	//
	// example:
	//
	// Available
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeJobResponseBodyJobTasksContact) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResponseBodyJobTasksContact) GoString() string {
	return s.String()
}

func (s *DescribeJobResponseBodyJobTasksContact) GetContactId() *string {
	return s.ContactId
}

func (s *DescribeJobResponseBodyJobTasksContact) GetContactName() *string {
	return s.ContactName
}

func (s *DescribeJobResponseBodyJobTasksContact) GetHonorific() *string {
	return s.Honorific
}

func (s *DescribeJobResponseBodyJobTasksContact) GetJobId() *string {
	return s.JobId
}

func (s *DescribeJobResponseBodyJobTasksContact) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *DescribeJobResponseBodyJobTasksContact) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *DescribeJobResponseBodyJobTasksContact) GetRole() *string {
	return s.Role
}

func (s *DescribeJobResponseBodyJobTasksContact) GetState() *string {
	return s.State
}

func (s *DescribeJobResponseBodyJobTasksContact) SetContactId(v string) *DescribeJobResponseBodyJobTasksContact {
	s.ContactId = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasksContact) SetContactName(v string) *DescribeJobResponseBodyJobTasksContact {
	s.ContactName = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasksContact) SetHonorific(v string) *DescribeJobResponseBodyJobTasksContact {
	s.Honorific = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasksContact) SetJobId(v string) *DescribeJobResponseBodyJobTasksContact {
	s.JobId = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasksContact) SetPhoneNumber(v string) *DescribeJobResponseBodyJobTasksContact {
	s.PhoneNumber = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasksContact) SetReferenceId(v string) *DescribeJobResponseBodyJobTasksContact {
	s.ReferenceId = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasksContact) SetRole(v string) *DescribeJobResponseBodyJobTasksContact {
	s.Role = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasksContact) SetState(v string) *DescribeJobResponseBodyJobTasksContact {
	s.State = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasksContact) Validate() error {
	return dara.Validate(s)
}

type DescribeJobResponseBodyJobTasksConversation struct {
	// Instruction.
	//
	// example:
	//
	// Broadcast
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// Instruction parameters.
	//
	// - When the interruption capability is Enabled, you can obtain the text content already spoken by the robot in the robot\\"s dialogue text by using the answer parameter value, for example:
	//
	//   - {"abortDialogueInfo":{"abortDialogueType":"Interrupted","answer":"This is the interrupted text information","sequenceId":"xxxxx"},"streamEnd":true}
	//
	// example:
	//
	// {}
	ActionParams *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	// Conversation text.
	//
	// example:
	//
	// 你好，我是**客服
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
	// Session ID.
	//
	// example:
	//
	// fd279983-93b9-b13b-9a34-64e5df473225
	SequenceId *string `json:"SequenceId,omitempty" xml:"SequenceId,omitempty"`
	// Indicates who spoke in the conversation: "Robot" for the robot or "Contact" for the contact.
	//
	// example:
	//
	// Robot
	Speaker *string `json:"Speaker,omitempty" xml:"Speaker,omitempty"`
	// Dialogue summary data (Historical Data, no longer used). [Deprecated]
	//
	// example:
	//
	// []
	Summary []*DescribeJobResponseBodyJobTasksConversationSummary `json:"Summary,omitempty" xml:"Summary,omitempty" type:"Repeated"`
	// Creation Time of the summary.
	//
	// example:
	//
	// 1579068424883
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeJobResponseBodyJobTasksConversation) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResponseBodyJobTasksConversation) GoString() string {
	return s.String()
}

func (s *DescribeJobResponseBodyJobTasksConversation) GetAction() *string {
	return s.Action
}

func (s *DescribeJobResponseBodyJobTasksConversation) GetActionParams() *string {
	return s.ActionParams
}

func (s *DescribeJobResponseBodyJobTasksConversation) GetScript() *string {
	return s.Script
}

func (s *DescribeJobResponseBodyJobTasksConversation) GetSequenceId() *string {
	return s.SequenceId
}

func (s *DescribeJobResponseBodyJobTasksConversation) GetSpeaker() *string {
	return s.Speaker
}

func (s *DescribeJobResponseBodyJobTasksConversation) GetSummary() []*DescribeJobResponseBodyJobTasksConversationSummary {
	return s.Summary
}

func (s *DescribeJobResponseBodyJobTasksConversation) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeJobResponseBodyJobTasksConversation) SetAction(v string) *DescribeJobResponseBodyJobTasksConversation {
	s.Action = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasksConversation) SetActionParams(v string) *DescribeJobResponseBodyJobTasksConversation {
	s.ActionParams = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasksConversation) SetScript(v string) *DescribeJobResponseBodyJobTasksConversation {
	s.Script = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasksConversation) SetSequenceId(v string) *DescribeJobResponseBodyJobTasksConversation {
	s.SequenceId = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasksConversation) SetSpeaker(v string) *DescribeJobResponseBodyJobTasksConversation {
	s.Speaker = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasksConversation) SetSummary(v []*DescribeJobResponseBodyJobTasksConversationSummary) *DescribeJobResponseBodyJobTasksConversation {
	s.Summary = v
	return s
}

func (s *DescribeJobResponseBodyJobTasksConversation) SetTimestamp(v int64) *DescribeJobResponseBodyJobTasksConversation {
	s.Timestamp = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasksConversation) Validate() error {
	if s.Summary != nil {
		for _, item := range s.Summary {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeJobResponseBodyJobTasksConversationSummary struct {
	// Dialogue summary category (Historical field, no longer used).
	//
	// example:
	//
	// {}
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Summary Content.
	//
	// example:
	//
	// 5
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Summary name.
	//
	// example:
	//
	// score
	SummaryName *string `json:"SummaryName,omitempty" xml:"SummaryName,omitempty"`
}

func (s DescribeJobResponseBodyJobTasksConversationSummary) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResponseBodyJobTasksConversationSummary) GoString() string {
	return s.String()
}

func (s *DescribeJobResponseBodyJobTasksConversationSummary) GetCategory() *string {
	return s.Category
}

func (s *DescribeJobResponseBodyJobTasksConversationSummary) GetContent() *string {
	return s.Content
}

func (s *DescribeJobResponseBodyJobTasksConversationSummary) GetSummaryName() *string {
	return s.SummaryName
}

func (s *DescribeJobResponseBodyJobTasksConversationSummary) SetCategory(v string) *DescribeJobResponseBodyJobTasksConversationSummary {
	s.Category = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasksConversationSummary) SetContent(v string) *DescribeJobResponseBodyJobTasksConversationSummary {
	s.Content = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasksConversationSummary) SetSummaryName(v string) *DescribeJobResponseBodyJobTasksConversationSummary {
	s.SummaryName = &v
	return s
}

func (s *DescribeJobResponseBodyJobTasksConversationSummary) Validate() error {
	return dara.Validate(s)
}
