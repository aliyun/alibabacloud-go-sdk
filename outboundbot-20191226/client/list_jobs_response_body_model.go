// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListJobsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListJobsResponseBody
	GetHttpStatusCode() *int32
	SetJobs(v []*ListJobsResponseBodyJobs) *ListJobsResponseBody
	GetJobs() []*ListJobsResponseBodyJobs
	SetMessage(v string) *ListJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListJobsResponseBody
	GetSuccess() *bool
}

type ListJobsResponseBody struct {
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Job array
	Jobs []*ListJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListJobsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListJobsResponseBody) GetJobs() []*ListJobsResponseBodyJobs {
	return s.Jobs
}

func (s *ListJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListJobsResponseBody) SetCode(v string) *ListJobsResponseBody {
	s.Code = &v
	return s
}

func (s *ListJobsResponseBody) SetHttpStatusCode(v int32) *ListJobsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListJobsResponseBody) SetJobs(v []*ListJobsResponseBodyJobs) *ListJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListJobsResponseBody) SetMessage(v string) *ListJobsResponseBody {
	s.Message = &v
	return s
}

func (s *ListJobsResponseBody) SetRequestId(v string) *ListJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobsResponseBody) SetSuccess(v bool) *ListJobsResponseBody {
	s.Success = &v
	return s
}

func (s *ListJobsResponseBody) Validate() error {
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

type ListJobsResponseBodyJobs struct {
	// List of calling numbers
	CallingNumbers []*string `json:"CallingNumbers,omitempty" xml:"CallingNumbers,omitempty" type:"Repeated"`
	// List of contacts
	Contacts []*ListJobsResponseBodyJobsContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Repeated"`
	// List of business parameters.
	//
	// > Note: TenantId and ServiceId are system-generated. All other parameters are custom and provided as reference.
	Extras []*ListJobsResponseBodyJobsExtras `json:"Extras,omitempty" xml:"Extras,omitempty" type:"Repeated"`
	// Job failure reason
	//
	// example:
	//
	// NoAnswer
	FailureReason *string `json:"FailureReason,omitempty" xml:"FailureReason,omitempty"`
	// Job group ID
	//
	// example:
	//
	// fce6c599-8ede-40e3-9f78-0928eda7b4e8
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// Job ID
	//
	// example:
	//
	// b72425bd-7871-4050-838e-033d80d754b7
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Job Priority
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
	// Business ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// Job status. Valid values:
	//
	// - **Scheduling**: Scheduling.
	//
	// - **Executing**: Executing.
	//
	// - **Succeeded**: Succeeded.
	//
	// - **Paused**: Suspended.
	//
	// - **Failed**: Failed.
	//
	// - **Cancelled**: Cancelled.
	//
	// - **Drafted**: Draft.
	//
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Policy ID
	//
	// example:
	//
	// c8a2b7f2-ad1a-4865-b872-d0080d9802d9
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// Label data for conversation jobs:
	//
	// In LLM scenarios: tag hit data after conversation completion.
	//
	// In small model scenarios: variable value data after conversation completion.
	Summary []*ListJobsResponseBodyJobsSummary `json:"Summary,omitempty" xml:"Summary,omitempty" type:"Repeated"`
	// System Priority
	//
	// example:
	//
	// 1
	SystemPriority *int32 `json:"SystemPriority,omitempty" xml:"SystemPriority,omitempty"`
	// Task list
	Tasks []*ListJobsResponseBodyJobsTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ListJobsResponseBodyJobs) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobs) GetCallingNumbers() []*string {
	return s.CallingNumbers
}

func (s *ListJobsResponseBodyJobs) GetContacts() []*ListJobsResponseBodyJobsContacts {
	return s.Contacts
}

func (s *ListJobsResponseBodyJobs) GetExtras() []*ListJobsResponseBodyJobsExtras {
	return s.Extras
}

func (s *ListJobsResponseBodyJobs) GetFailureReason() *string {
	return s.FailureReason
}

func (s *ListJobsResponseBodyJobs) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *ListJobsResponseBodyJobs) GetJobId() *string {
	return s.JobId
}

func (s *ListJobsResponseBodyJobs) GetPriority() *int32 {
	return s.Priority
}

func (s *ListJobsResponseBodyJobs) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *ListJobsResponseBodyJobs) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *ListJobsResponseBodyJobs) GetStatus() *string {
	return s.Status
}

func (s *ListJobsResponseBodyJobs) GetStrategyId() *string {
	return s.StrategyId
}

func (s *ListJobsResponseBodyJobs) GetSummary() []*ListJobsResponseBodyJobsSummary {
	return s.Summary
}

func (s *ListJobsResponseBodyJobs) GetSystemPriority() *int32 {
	return s.SystemPriority
}

func (s *ListJobsResponseBodyJobs) GetTasks() []*ListJobsResponseBodyJobsTasks {
	return s.Tasks
}

func (s *ListJobsResponseBodyJobs) SetCallingNumbers(v []*string) *ListJobsResponseBodyJobs {
	s.CallingNumbers = v
	return s
}

func (s *ListJobsResponseBodyJobs) SetContacts(v []*ListJobsResponseBodyJobsContacts) *ListJobsResponseBodyJobs {
	s.Contacts = v
	return s
}

func (s *ListJobsResponseBodyJobs) SetExtras(v []*ListJobsResponseBodyJobsExtras) *ListJobsResponseBodyJobs {
	s.Extras = v
	return s
}

func (s *ListJobsResponseBodyJobs) SetFailureReason(v string) *ListJobsResponseBodyJobs {
	s.FailureReason = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetJobGroupId(v string) *ListJobsResponseBodyJobs {
	s.JobGroupId = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetJobId(v string) *ListJobsResponseBodyJobs {
	s.JobId = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetPriority(v int32) *ListJobsResponseBodyJobs {
	s.Priority = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetReferenceId(v string) *ListJobsResponseBodyJobs {
	s.ReferenceId = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetScenarioId(v string) *ListJobsResponseBodyJobs {
	s.ScenarioId = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetStatus(v string) *ListJobsResponseBodyJobs {
	s.Status = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetStrategyId(v string) *ListJobsResponseBodyJobs {
	s.StrategyId = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetSummary(v []*ListJobsResponseBodyJobsSummary) *ListJobsResponseBodyJobs {
	s.Summary = v
	return s
}

func (s *ListJobsResponseBodyJobs) SetSystemPriority(v int32) *ListJobsResponseBodyJobs {
	s.SystemPriority = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetTasks(v []*ListJobsResponseBodyJobsTasks) *ListJobsResponseBodyJobs {
	s.Tasks = v
	return s
}

func (s *ListJobsResponseBodyJobs) Validate() error {
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

type ListJobsResponseBodyJobsContacts struct {
	// Contact ID. Automatically generated by the system
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
	// Honorific. Same as the contact name
	//
	// example:
	//
	// 张先生
	Honorific *string `json:"Honorific,omitempty" xml:"Honorific,omitempty"`
	// Job ID [Deprecated]
	//
	// example:
	//
	// 72dcd26b-f12d-4c27-b3af-18f6aed5b160
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Phone number
	//
	// example:
	//
	// 135****8888
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Third-party system ID of the contact
	//
	// example:
	//
	// 2fa6bac3-06da-4315-82ab-72d6fd3a6f34
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// Contact role [Deprecated]
	//
	// example:
	//
	// *
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// Status [Deprecated]
	//
	// example:
	//
	// Available
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListJobsResponseBodyJobsContacts) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyJobsContacts) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobsContacts) GetContactId() *string {
	return s.ContactId
}

func (s *ListJobsResponseBodyJobsContacts) GetContactName() *string {
	return s.ContactName
}

func (s *ListJobsResponseBodyJobsContacts) GetHonorific() *string {
	return s.Honorific
}

func (s *ListJobsResponseBodyJobsContacts) GetJobId() *string {
	return s.JobId
}

func (s *ListJobsResponseBodyJobsContacts) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ListJobsResponseBodyJobsContacts) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *ListJobsResponseBodyJobsContacts) GetRole() *string {
	return s.Role
}

func (s *ListJobsResponseBodyJobsContacts) GetState() *string {
	return s.State
}

func (s *ListJobsResponseBodyJobsContacts) SetContactId(v string) *ListJobsResponseBodyJobsContacts {
	s.ContactId = &v
	return s
}

func (s *ListJobsResponseBodyJobsContacts) SetContactName(v string) *ListJobsResponseBodyJobsContacts {
	s.ContactName = &v
	return s
}

func (s *ListJobsResponseBodyJobsContacts) SetHonorific(v string) *ListJobsResponseBodyJobsContacts {
	s.Honorific = &v
	return s
}

func (s *ListJobsResponseBodyJobsContacts) SetJobId(v string) *ListJobsResponseBodyJobsContacts {
	s.JobId = &v
	return s
}

func (s *ListJobsResponseBodyJobsContacts) SetPhoneNumber(v string) *ListJobsResponseBodyJobsContacts {
	s.PhoneNumber = &v
	return s
}

func (s *ListJobsResponseBodyJobsContacts) SetReferenceId(v string) *ListJobsResponseBodyJobsContacts {
	s.ReferenceId = &v
	return s
}

func (s *ListJobsResponseBodyJobsContacts) SetRole(v string) *ListJobsResponseBodyJobsContacts {
	s.Role = &v
	return s
}

func (s *ListJobsResponseBodyJobsContacts) SetState(v string) *ListJobsResponseBodyJobsContacts {
	s.State = &v
	return s
}

func (s *ListJobsResponseBodyJobsContacts) Validate() error {
	return dara.Validate(s)
}

type ListJobsResponseBodyJobsExtras struct {
	// Business parameter key
	//
	// example:
	//
	// name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Business parameter value
	//
	// example:
	//
	// 张三
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListJobsResponseBodyJobsExtras) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyJobsExtras) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobsExtras) GetKey() *string {
	return s.Key
}

func (s *ListJobsResponseBodyJobsExtras) GetValue() *string {
	return s.Value
}

func (s *ListJobsResponseBodyJobsExtras) SetKey(v string) *ListJobsResponseBodyJobsExtras {
	s.Key = &v
	return s
}

func (s *ListJobsResponseBodyJobsExtras) SetValue(v string) *ListJobsResponseBodyJobsExtras {
	s.Value = &v
	return s
}

func (s *ListJobsResponseBodyJobsExtras) Validate() error {
	return dara.Validate(s)
}

type ListJobsResponseBodyJobsSummary struct {
	// Conversation summary category
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
	// Label name.
	//
	// example:
	//
	// score
	SummaryName *string `json:"SummaryName,omitempty" xml:"SummaryName,omitempty"`
}

func (s ListJobsResponseBodyJobsSummary) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyJobsSummary) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobsSummary) GetCategory() *string {
	return s.Category
}

func (s *ListJobsResponseBodyJobsSummary) GetContent() *string {
	return s.Content
}

func (s *ListJobsResponseBodyJobsSummary) GetSummaryName() *string {
	return s.SummaryName
}

func (s *ListJobsResponseBodyJobsSummary) SetCategory(v string) *ListJobsResponseBodyJobsSummary {
	s.Category = &v
	return s
}

func (s *ListJobsResponseBodyJobsSummary) SetContent(v string) *ListJobsResponseBodyJobsSummary {
	s.Content = &v
	return s
}

func (s *ListJobsResponseBodyJobsSummary) SetSummaryName(v string) *ListJobsResponseBodyJobsSummary {
	s.SummaryName = &v
	return s
}

func (s *ListJobsResponseBodyJobsSummary) Validate() error {
	return dara.Validate(s)
}

type ListJobsResponseBodyJobsTasks struct {
	// Actual running time
	//
	// example:
	//
	// 1579068424883
	ActualTime *int64 `json:"ActualTime,omitempty" xml:"ActualTime,omitempty"`
	// Business result. Historical parameter [Deprecated]
	//
	// example:
	//
	// 1
	Brief *string `json:"Brief,omitempty" xml:"Brief,omitempty"`
	// Call ID
	//
	// example:
	//
	// 1528189846043
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// Called number
	//
	// example:
	//
	// 135****8888
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// Calling number
	//
	// example:
	//
	// 0571****3106
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// Chatbot ID
	//
	// example:
	//
	// 1234
	ChatbotId *string `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	// Contact for the job
	Contact *ListJobsResponseBodyJobsTasksContact `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Struct"`
	// Conversation List
	Conversation []*ListJobsResponseBodyJobsTasksConversation `json:"Conversation,omitempty" xml:"Conversation,omitempty" type:"Repeated"`
	// Duration.
	//
	// example:
	//
	// 120
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Job ID
	//
	// example:
	//
	// b72425bd-7871-4050-838e-033d80d754b7
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Planned execution time of the task
	//
	// example:
	//
	// 1579068424883
	PlanedTime *int64 `json:"PlanedTime,omitempty" xml:"PlanedTime,omitempty"`
	// Scenario ID. Historical parameter. [Deprecated]
	//
	// example:
	//
	// ade80092-03d9-4f4d-ad4f-ab8a247d3150
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// Task Status, with the following possible values:
	//
	// (Note: The **Succeeded*	- status has been further categorized by specific reasons. The generic **Succeeded**: 1 (Answered) status will no longer be displayed; instead, only the detailed reason types will be returned.)
	//
	// - **Executing**: 0 (Calling in progress).
	//
	// - **Succeeded**: 1 (Answered).
	//
	// - **NoAnswer**: 2 (Not answered – No one picked up).
	//
	// - **NotExist**: 3 (Not answered – Nonexistent number).
	//
	// - **Busy**: 4 (Not answered – Line busy).
	//
	// - **Cancelled**: 5 (Not dialed – Job stopped).
	//
	// - **Failed**: 6 (Failed).
	//
	// - **NotConnected**: 7 (Not answered – Unable to connect).
	//
	// - **PoweredOff**: 8 (Not answered – Powered off).
	//
	// - **OutOfService**: 9 (Not answered – Called party suspended).
	//
	// - **InArrears**: 10 (Not answered – Called party has overdue payment).
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
	// - **SucceededFinish**: 20 (Answered – Normal completion).
	//
	// - **SucceededChatbotHangUpAfterNoAnswer**: 21 (Answered – Bot hung up after unrecognized input).
	//
	// - **SucceededChatbotHangUpAfterSilence**: 22 (Answered – Bot hung up after silence timeout).
	//
	// - **SucceededClientHangUpAfterNoAnswer**: 23 (Answered – User hung up after unrecognized input).
	//
	// - **SucceededClientHangUp**: 24 (Answered – User hung up without reason).
	//
	// - **SucceededTransferByIntent**: 25 (Answered – Transferred to agent due to intent match).
	//
	// - **SucceededTransferAfterNoAnswer**: 26 (Answered – Transferred to agent after unrecognized input).
	//
	// - **SucceededInoInterAction**: 27 (Answered – No interaction from user side).
	//
	// - **SucceededError**: 28 (Answered – Interrupted due to system exception).
	//
	// - **SucceededSpecialInterceptVoiceAssistant**: 29 (Answered – Intercepted under special condition – Voice assistant).
	//
	// - **SucceededSpecialInterceptExtensionNumberTransfer**: 30 (Answered – Intercepted under special condition – Extension transfer).
	//
	// - **SucceededSpecialInterceptCustomSpecialIntercept**: 31 (Answered – Intercepted under special condition – Custom interception).
	//
	// - **HighRiskSipCode**: 32 (Not dialed – High-risk, not called).
	//
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Job ID
	//
	// example:
	//
	// ff44709e-39a6-43ba-959b-20fcabe3e496
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListJobsResponseBodyJobsTasks) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyJobsTasks) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobsTasks) GetActualTime() *int64 {
	return s.ActualTime
}

func (s *ListJobsResponseBodyJobsTasks) GetBrief() *string {
	return s.Brief
}

func (s *ListJobsResponseBodyJobsTasks) GetCallId() *string {
	return s.CallId
}

func (s *ListJobsResponseBodyJobsTasks) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *ListJobsResponseBodyJobsTasks) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *ListJobsResponseBodyJobsTasks) GetChatbotId() *string {
	return s.ChatbotId
}

func (s *ListJobsResponseBodyJobsTasks) GetContact() *ListJobsResponseBodyJobsTasksContact {
	return s.Contact
}

func (s *ListJobsResponseBodyJobsTasks) GetConversation() []*ListJobsResponseBodyJobsTasksConversation {
	return s.Conversation
}

func (s *ListJobsResponseBodyJobsTasks) GetDuration() *int32 {
	return s.Duration
}

func (s *ListJobsResponseBodyJobsTasks) GetJobId() *string {
	return s.JobId
}

func (s *ListJobsResponseBodyJobsTasks) GetPlanedTime() *int64 {
	return s.PlanedTime
}

func (s *ListJobsResponseBodyJobsTasks) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *ListJobsResponseBodyJobsTasks) GetStatus() *string {
	return s.Status
}

func (s *ListJobsResponseBodyJobsTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *ListJobsResponseBodyJobsTasks) SetActualTime(v int64) *ListJobsResponseBodyJobsTasks {
	s.ActualTime = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasks) SetBrief(v string) *ListJobsResponseBodyJobsTasks {
	s.Brief = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasks) SetCallId(v string) *ListJobsResponseBodyJobsTasks {
	s.CallId = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasks) SetCalledNumber(v string) *ListJobsResponseBodyJobsTasks {
	s.CalledNumber = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasks) SetCallingNumber(v string) *ListJobsResponseBodyJobsTasks {
	s.CallingNumber = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasks) SetChatbotId(v string) *ListJobsResponseBodyJobsTasks {
	s.ChatbotId = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasks) SetContact(v *ListJobsResponseBodyJobsTasksContact) *ListJobsResponseBodyJobsTasks {
	s.Contact = v
	return s
}

func (s *ListJobsResponseBodyJobsTasks) SetConversation(v []*ListJobsResponseBodyJobsTasksConversation) *ListJobsResponseBodyJobsTasks {
	s.Conversation = v
	return s
}

func (s *ListJobsResponseBodyJobsTasks) SetDuration(v int32) *ListJobsResponseBodyJobsTasks {
	s.Duration = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasks) SetJobId(v string) *ListJobsResponseBodyJobsTasks {
	s.JobId = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasks) SetPlanedTime(v int64) *ListJobsResponseBodyJobsTasks {
	s.PlanedTime = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasks) SetScenarioId(v string) *ListJobsResponseBodyJobsTasks {
	s.ScenarioId = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasks) SetStatus(v string) *ListJobsResponseBodyJobsTasks {
	s.Status = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasks) SetTaskId(v string) *ListJobsResponseBodyJobsTasks {
	s.TaskId = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasks) Validate() error {
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

type ListJobsResponseBodyJobsTasksContact struct {
	// Contact ID
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
	// Honorific title of the contact. If not specified, it is the same as the contact name.
	//
	// example:
	//
	// 张先生
	Honorific *string `json:"Honorific,omitempty" xml:"Honorific,omitempty"`
	// Job ID. Deprecated
	//
	// example:
	//
	// b72425bd-7871-4050-838e-033d80d754b7
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Phone number
	//
	// example:
	//
	// 135****8888
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Third-party system ID of the contact
	//
	// example:
	//
	// 2fa6bac3-06da-4315-82ab-72d6fd3a6f34
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// Role.
	//
	// example:
	//
	// *
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// Status.
	//
	// example:
	//
	// Available
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListJobsResponseBodyJobsTasksContact) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyJobsTasksContact) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobsTasksContact) GetContactId() *string {
	return s.ContactId
}

func (s *ListJobsResponseBodyJobsTasksContact) GetContactName() *string {
	return s.ContactName
}

func (s *ListJobsResponseBodyJobsTasksContact) GetHonorific() *string {
	return s.Honorific
}

func (s *ListJobsResponseBodyJobsTasksContact) GetJobId() *string {
	return s.JobId
}

func (s *ListJobsResponseBodyJobsTasksContact) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ListJobsResponseBodyJobsTasksContact) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *ListJobsResponseBodyJobsTasksContact) GetRole() *string {
	return s.Role
}

func (s *ListJobsResponseBodyJobsTasksContact) GetState() *string {
	return s.State
}

func (s *ListJobsResponseBodyJobsTasksContact) SetContactId(v string) *ListJobsResponseBodyJobsTasksContact {
	s.ContactId = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasksContact) SetContactName(v string) *ListJobsResponseBodyJobsTasksContact {
	s.ContactName = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasksContact) SetHonorific(v string) *ListJobsResponseBodyJobsTasksContact {
	s.Honorific = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasksContact) SetJobId(v string) *ListJobsResponseBodyJobsTasksContact {
	s.JobId = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasksContact) SetPhoneNumber(v string) *ListJobsResponseBodyJobsTasksContact {
	s.PhoneNumber = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasksContact) SetReferenceId(v string) *ListJobsResponseBodyJobsTasksContact {
	s.ReferenceId = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasksContact) SetRole(v string) *ListJobsResponseBodyJobsTasksContact {
	s.Role = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasksContact) SetState(v string) *ListJobsResponseBodyJobsTasksContact {
	s.State = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasksContact) Validate() error {
	return dara.Validate(s)
}

type ListJobsResponseBodyJobsTasksConversation struct {
	// Conversation text
	//
	// example:
	//
	// 你好，我是**客服
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
	// Speaker of the utterance: "Robot" for the bot, "Contact" for the contact
	//
	// example:
	//
	// Robot
	Speaker *string `json:"Speaker,omitempty" xml:"Speaker,omitempty"`
	// Summary information. [Deprecated]
	Summary []*ListJobsResponseBodyJobsTasksConversationSummary `json:"Summary,omitempty" xml:"Summary,omitempty" type:"Repeated"`
	// UNIX timestamp when the conversation text was stored
	//
	// example:
	//
	// 1579068424883
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ListJobsResponseBodyJobsTasksConversation) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyJobsTasksConversation) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobsTasksConversation) GetScript() *string {
	return s.Script
}

func (s *ListJobsResponseBodyJobsTasksConversation) GetSpeaker() *string {
	return s.Speaker
}

func (s *ListJobsResponseBodyJobsTasksConversation) GetSummary() []*ListJobsResponseBodyJobsTasksConversationSummary {
	return s.Summary
}

func (s *ListJobsResponseBodyJobsTasksConversation) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *ListJobsResponseBodyJobsTasksConversation) SetScript(v string) *ListJobsResponseBodyJobsTasksConversation {
	s.Script = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasksConversation) SetSpeaker(v string) *ListJobsResponseBodyJobsTasksConversation {
	s.Speaker = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasksConversation) SetSummary(v []*ListJobsResponseBodyJobsTasksConversationSummary) *ListJobsResponseBodyJobsTasksConversation {
	s.Summary = v
	return s
}

func (s *ListJobsResponseBodyJobsTasksConversation) SetTimestamp(v int64) *ListJobsResponseBodyJobsTasksConversation {
	s.Timestamp = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasksConversation) Validate() error {
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

type ListJobsResponseBodyJobsTasksConversationSummary struct {
	// category
	//
	// example:
	//
	// 标签
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Summary content
	//
	// example:
	//
	// 是
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Summary Name
	//
	// example:
	//
	// 是本人
	SummaryName *string `json:"SummaryName,omitempty" xml:"SummaryName,omitempty"`
}

func (s ListJobsResponseBodyJobsTasksConversationSummary) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyJobsTasksConversationSummary) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobsTasksConversationSummary) GetCategory() *string {
	return s.Category
}

func (s *ListJobsResponseBodyJobsTasksConversationSummary) GetContent() *string {
	return s.Content
}

func (s *ListJobsResponseBodyJobsTasksConversationSummary) GetSummaryName() *string {
	return s.SummaryName
}

func (s *ListJobsResponseBodyJobsTasksConversationSummary) SetCategory(v string) *ListJobsResponseBodyJobsTasksConversationSummary {
	s.Category = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasksConversationSummary) SetContent(v string) *ListJobsResponseBodyJobsTasksConversationSummary {
	s.Content = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasksConversationSummary) SetSummaryName(v string) *ListJobsResponseBodyJobsTasksConversationSummary {
	s.SummaryName = &v
	return s
}

func (s *ListJobsResponseBodyJobsTasksConversationSummary) Validate() error {
	return dara.Validate(s)
}
