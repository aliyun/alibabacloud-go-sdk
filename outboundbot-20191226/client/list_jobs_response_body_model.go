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
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Jobs           []*ListJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CallingNumbers []*string                           `json:"CallingNumbers,omitempty" xml:"CallingNumbers,omitempty" type:"Repeated"`
	Contacts       []*ListJobsResponseBodyJobsContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Repeated"`
	Extras         []*ListJobsResponseBodyJobsExtras   `json:"Extras,omitempty" xml:"Extras,omitempty" type:"Repeated"`
	// example:
	//
	// NoAnswer
	FailureReason *string `json:"FailureReason,omitempty" xml:"FailureReason,omitempty"`
	// example:
	//
	// fce6c599-8ede-40e3-9f78-0928eda7b4e8
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// example:
	//
	// b72425bd-7871-4050-838e-033d80d754b7
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// d5971d98-7312-4f0e-a918-a17d67133e28
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// c8a2b7f2-ad1a-4865-b872-d0080d9802d9
	StrategyId *string                            `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	Summary    []*ListJobsResponseBodyJobsSummary `json:"Summary,omitempty" xml:"Summary,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	SystemPriority *int32                           `json:"SystemPriority,omitempty" xml:"SystemPriority,omitempty"`
	Tasks          []*ListJobsResponseBodyJobsTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
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
	// example:
	//
	// db3db762-e421-44c9-9a01-cb423470757c
	ContactId   *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	Honorific   *string `json:"Honorific,omitempty" xml:"Honorific,omitempty"`
	// example:
	//
	// 72dcd26b-f12d-4c27-b3af-18f6aed5b160
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 135****8888
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// example:
	//
	// 2fa6bac3-06da-4315-82ab-72d6fd3a6f34
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// example:
	//
	// *
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
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
	// example:
	//
	// name
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	// example:
	//
	// {}
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// 5
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
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
	// example:
	//
	// 1579068424883
	ActualTime *int64 `json:"ActualTime,omitempty" xml:"ActualTime,omitempty"`
	// example:
	//
	// 1
	Brief *string `json:"Brief,omitempty" xml:"Brief,omitempty"`
	// example:
	//
	// 1528189846043
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// example:
	//
	// 135****8888
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// example:
	//
	// 0571****3106
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// example:
	//
	// 1234
	ChatbotId    *string                                      `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	Contact      *ListJobsResponseBodyJobsTasksContact        `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Struct"`
	Conversation []*ListJobsResponseBodyJobsTasksConversation `json:"Conversation,omitempty" xml:"Conversation,omitempty" type:"Repeated"`
	// example:
	//
	// 120
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// b72425bd-7871-4050-838e-033d80d754b7
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 1579068424883
	PlanedTime *int64 `json:"PlanedTime,omitempty" xml:"PlanedTime,omitempty"`
	// example:
	//
	// ade80092-03d9-4f4d-ad4f-ab8a247d3150
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	// example:
	//
	// db3db762-e421-44c9-9a01-cb423470757c
	ContactId   *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	Honorific   *string `json:"Honorific,omitempty" xml:"Honorific,omitempty"`
	// example:
	//
	// b72425bd-7871-4050-838e-033d80d754b7
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 135****8888
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// example:
	//
	// 2fa6bac3-06da-4315-82ab-72d6fd3a6f34
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// example:
	//
	// *
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
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
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
	// example:
	//
	// Robot
	Speaker *string                                             `json:"Speaker,omitempty" xml:"Speaker,omitempty"`
	Summary []*ListJobsResponseBodyJobsTasksConversationSummary `json:"Summary,omitempty" xml:"Summary,omitempty" type:"Repeated"`
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
	Category    *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
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
