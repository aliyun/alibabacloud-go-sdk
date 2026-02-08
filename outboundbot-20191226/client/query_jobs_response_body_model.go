// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryJobsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *QueryJobsResponseBody
	GetHttpStatusCode() *int32
	SetJobs(v *QueryJobsResponseBodyJobs) *QueryJobsResponseBody
	GetJobs() *QueryJobsResponseBodyJobs
	SetMessage(v string) *QueryJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryJobsResponseBody
	GetSuccess() *bool
}

type QueryJobsResponseBody struct {
	// HTTP status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Job data.
	Jobs *QueryJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Struct"`
	// Response message.
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

func (s QueryJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryJobsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryJobsResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryJobsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryJobsResponseBody) GetJobs() *QueryJobsResponseBodyJobs {
	return s.Jobs
}

func (s *QueryJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryJobsResponseBody) SetCode(v string) *QueryJobsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryJobsResponseBody) SetHttpStatusCode(v int32) *QueryJobsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryJobsResponseBody) SetJobs(v *QueryJobsResponseBodyJobs) *QueryJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *QueryJobsResponseBody) SetMessage(v string) *QueryJobsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryJobsResponseBody) SetRequestId(v string) *QueryJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryJobsResponseBody) SetSuccess(v bool) *QueryJobsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryJobsResponseBody) Validate() error {
	if s.Jobs != nil {
		if err := s.Jobs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryJobsResponseBodyJobs struct {
	// Job array.
	List []*QueryJobsResponseBodyJobsList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Paging size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total number of data entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryJobsResponseBodyJobs) String() string {
	return dara.Prettify(s)
}

func (s QueryJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *QueryJobsResponseBodyJobs) GetList() []*QueryJobsResponseBodyJobsList {
	return s.List
}

func (s *QueryJobsResponseBodyJobs) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryJobsResponseBodyJobs) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryJobsResponseBodyJobs) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryJobsResponseBodyJobs) SetList(v []*QueryJobsResponseBodyJobsList) *QueryJobsResponseBodyJobs {
	s.List = v
	return s
}

func (s *QueryJobsResponseBodyJobs) SetPageNumber(v int32) *QueryJobsResponseBodyJobs {
	s.PageNumber = &v
	return s
}

func (s *QueryJobsResponseBodyJobs) SetPageSize(v int32) *QueryJobsResponseBodyJobs {
	s.PageSize = &v
	return s
}

func (s *QueryJobsResponseBodyJobs) SetTotalCount(v int32) *QueryJobsResponseBodyJobs {
	s.TotalCount = &v
	return s
}

func (s *QueryJobsResponseBodyJobs) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryJobsResponseBodyJobsList struct {
	// List of calling numbers.
	CallingNumbers []*string `json:"CallingNumbers,omitempty" xml:"CallingNumbers,omitempty" type:"Repeated"`
	// Contact information. This parameter is deprecated.
	//
	// > You can obtain it through the DescribeJob API.
	Contacts []*QueryJobsResponseBodyJobsListContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Repeated"`
	// Business Data, displaying the label collection status in LLM scenarios.
	//
	// > Keys equal to TenantId or ServiceId are system parameters.
	Extras []*QueryJobsResponseBodyJobsListExtras `json:"Extras,omitempty" xml:"Extras,omitempty" type:"Repeated"`
	// Failure reason.
	//
	// - Unknown (unknown fault),
	//
	// - NoAnswer (no answer),
	//
	// - InvalidStrategy (invalid policy, Configuration Incorrect),
	//
	// - TimeUp (timeout detected during schedule),
	//
	// - NoStrategy (policy is empty or not found),
	//
	// - CallFailed (call failed),
	//
	// - PerDayCallCountLimit (daily call count limit for the number reached),
	//
	// - ContactBlockList (contact on do-not-call list),
	//
	// - EmptyNumber (nonexistent number, no further calls),
	//
	// - JobPerDayCallCountLimit (daily call count limit for the job reached),
	//
	// - VerificationCancelled (cancelled because authentication failed before call),
	//
	// - ContactSuspended (suspended),
	//
	// - InArrears (overdue payment),
	//
	// - OutOfService (service suspended);
	//
	// example:
	//
	// NoAnswer
	FailureReason *string `json:"FailureReason,omitempty" xml:"FailureReason,omitempty"`
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
	// fce6c599-8ede-40e3-9f78-0928eda7b4e8
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Associated business ID.
	//
	// example:
	//
	// d5971d98-7312-4f0e-a918-a17d67133e28
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// Scenario ID. Historical parameter. [Deprecated]
	//
	// example:
	//
	// ade80092-03d9-4f4d-ad4f-ab8a247d3150
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// Job status.
	//
	// - Scheduling(0, "Scheduling")
	//
	// - Executing(1, "Executing")
	//
	// - Succeeded(2, "End - Reached")
	//
	// - Paused(3, "Paused")
	//
	// - Failed(4, "End - Not Reached")
	//
	// - Cancelled(5, "Cancelled - Manual Intervention")
	//
	// - Drafted(6, "Draft")
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
	// Conversation summary. (Historical field, no longer used.) [Deprecated]
	Summary []*QueryJobsResponseBodyJobsListSummary `json:"Summary,omitempty" xml:"Summary,omitempty" type:"Repeated"`
	// Tag hit information in the small model scenario.
	TagHits []*QueryJobsResponseBodyJobsListTagHits `json:"TagHits,omitempty" xml:"TagHits,omitempty" type:"Repeated"`
	// Call list. This parameter is deprecated.
	//
	// > You can obtain it through the searchTask API.
	Tasks []*QueryJobsResponseBodyJobsListTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s QueryJobsResponseBodyJobsList) String() string {
	return dara.Prettify(s)
}

func (s QueryJobsResponseBodyJobsList) GoString() string {
	return s.String()
}

func (s *QueryJobsResponseBodyJobsList) GetCallingNumbers() []*string {
	return s.CallingNumbers
}

func (s *QueryJobsResponseBodyJobsList) GetContacts() []*QueryJobsResponseBodyJobsListContacts {
	return s.Contacts
}

func (s *QueryJobsResponseBodyJobsList) GetExtras() []*QueryJobsResponseBodyJobsListExtras {
	return s.Extras
}

func (s *QueryJobsResponseBodyJobsList) GetFailureReason() *string {
	return s.FailureReason
}

func (s *QueryJobsResponseBodyJobsList) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *QueryJobsResponseBodyJobsList) GetJobId() *string {
	return s.JobId
}

func (s *QueryJobsResponseBodyJobsList) GetPriority() *int32 {
	return s.Priority
}

func (s *QueryJobsResponseBodyJobsList) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *QueryJobsResponseBodyJobsList) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *QueryJobsResponseBodyJobsList) GetStatus() *string {
	return s.Status
}

func (s *QueryJobsResponseBodyJobsList) GetStrategyId() *string {
	return s.StrategyId
}

func (s *QueryJobsResponseBodyJobsList) GetSummary() []*QueryJobsResponseBodyJobsListSummary {
	return s.Summary
}

func (s *QueryJobsResponseBodyJobsList) GetTagHits() []*QueryJobsResponseBodyJobsListTagHits {
	return s.TagHits
}

func (s *QueryJobsResponseBodyJobsList) GetTasks() []*QueryJobsResponseBodyJobsListTasks {
	return s.Tasks
}

func (s *QueryJobsResponseBodyJobsList) SetCallingNumbers(v []*string) *QueryJobsResponseBodyJobsList {
	s.CallingNumbers = v
	return s
}

func (s *QueryJobsResponseBodyJobsList) SetContacts(v []*QueryJobsResponseBodyJobsListContacts) *QueryJobsResponseBodyJobsList {
	s.Contacts = v
	return s
}

func (s *QueryJobsResponseBodyJobsList) SetExtras(v []*QueryJobsResponseBodyJobsListExtras) *QueryJobsResponseBodyJobsList {
	s.Extras = v
	return s
}

func (s *QueryJobsResponseBodyJobsList) SetFailureReason(v string) *QueryJobsResponseBodyJobsList {
	s.FailureReason = &v
	return s
}

func (s *QueryJobsResponseBodyJobsList) SetJobGroupId(v string) *QueryJobsResponseBodyJobsList {
	s.JobGroupId = &v
	return s
}

func (s *QueryJobsResponseBodyJobsList) SetJobId(v string) *QueryJobsResponseBodyJobsList {
	s.JobId = &v
	return s
}

func (s *QueryJobsResponseBodyJobsList) SetPriority(v int32) *QueryJobsResponseBodyJobsList {
	s.Priority = &v
	return s
}

func (s *QueryJobsResponseBodyJobsList) SetReferenceId(v string) *QueryJobsResponseBodyJobsList {
	s.ReferenceId = &v
	return s
}

func (s *QueryJobsResponseBodyJobsList) SetScenarioId(v string) *QueryJobsResponseBodyJobsList {
	s.ScenarioId = &v
	return s
}

func (s *QueryJobsResponseBodyJobsList) SetStatus(v string) *QueryJobsResponseBodyJobsList {
	s.Status = &v
	return s
}

func (s *QueryJobsResponseBodyJobsList) SetStrategyId(v string) *QueryJobsResponseBodyJobsList {
	s.StrategyId = &v
	return s
}

func (s *QueryJobsResponseBodyJobsList) SetSummary(v []*QueryJobsResponseBodyJobsListSummary) *QueryJobsResponseBodyJobsList {
	s.Summary = v
	return s
}

func (s *QueryJobsResponseBodyJobsList) SetTagHits(v []*QueryJobsResponseBodyJobsListTagHits) *QueryJobsResponseBodyJobsList {
	s.TagHits = v
	return s
}

func (s *QueryJobsResponseBodyJobsList) SetTasks(v []*QueryJobsResponseBodyJobsListTasks) *QueryJobsResponseBodyJobsList {
	s.Tasks = v
	return s
}

func (s *QueryJobsResponseBodyJobsList) Validate() error {
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
	if s.TagHits != nil {
		for _, item := range s.TagHits {
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

type QueryJobsResponseBodyJobsListContacts struct {
	// Contact ID.
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
	// Honorific.
	//
	// example:
	//
	// 张先生
	Honorific *string `json:"Honorific,omitempty" xml:"Honorific,omitempty"`
	// Job ID.
	//
	// example:
	//
	// fce6c599-8ede-40e3-9f78-0928eda7b4e8
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Phone number.
	//
	// example:
	//
	// 135****8888
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Associated business ID.
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
	// - Available (Normal),
	//
	// - WrongNumber (fault),
	//
	// - DoesNotExist (nonexistent number),
	//
	// - Suspended (suspended);
	//
	// example:
	//
	// Available
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s QueryJobsResponseBodyJobsListContacts) String() string {
	return dara.Prettify(s)
}

func (s QueryJobsResponseBodyJobsListContacts) GoString() string {
	return s.String()
}

func (s *QueryJobsResponseBodyJobsListContacts) GetContactId() *string {
	return s.ContactId
}

func (s *QueryJobsResponseBodyJobsListContacts) GetContactName() *string {
	return s.ContactName
}

func (s *QueryJobsResponseBodyJobsListContacts) GetHonorific() *string {
	return s.Honorific
}

func (s *QueryJobsResponseBodyJobsListContacts) GetJobId() *string {
	return s.JobId
}

func (s *QueryJobsResponseBodyJobsListContacts) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *QueryJobsResponseBodyJobsListContacts) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *QueryJobsResponseBodyJobsListContacts) GetRole() *string {
	return s.Role
}

func (s *QueryJobsResponseBodyJobsListContacts) GetState() *string {
	return s.State
}

func (s *QueryJobsResponseBodyJobsListContacts) SetContactId(v string) *QueryJobsResponseBodyJobsListContacts {
	s.ContactId = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListContacts) SetContactName(v string) *QueryJobsResponseBodyJobsListContacts {
	s.ContactName = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListContacts) SetHonorific(v string) *QueryJobsResponseBodyJobsListContacts {
	s.Honorific = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListContacts) SetJobId(v string) *QueryJobsResponseBodyJobsListContacts {
	s.JobId = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListContacts) SetPhoneNumber(v string) *QueryJobsResponseBodyJobsListContacts {
	s.PhoneNumber = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListContacts) SetReferenceId(v string) *QueryJobsResponseBodyJobsListContacts {
	s.ReferenceId = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListContacts) SetRole(v string) *QueryJobsResponseBodyJobsListContacts {
	s.Role = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListContacts) SetState(v string) *QueryJobsResponseBodyJobsListContacts {
	s.State = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListContacts) Validate() error {
	return dara.Validate(s)
}

type QueryJobsResponseBodyJobsListExtras struct {
	// Business Data key
	//
	// example:
	//
	// djrq
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Business Data value
	//
	// example:
	//
	// 2019-08-21 09:49:59.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryJobsResponseBodyJobsListExtras) String() string {
	return dara.Prettify(s)
}

func (s QueryJobsResponseBodyJobsListExtras) GoString() string {
	return s.String()
}

func (s *QueryJobsResponseBodyJobsListExtras) GetKey() *string {
	return s.Key
}

func (s *QueryJobsResponseBodyJobsListExtras) GetValue() *string {
	return s.Value
}

func (s *QueryJobsResponseBodyJobsListExtras) SetKey(v string) *QueryJobsResponseBodyJobsListExtras {
	s.Key = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListExtras) SetValue(v string) *QueryJobsResponseBodyJobsListExtras {
	s.Value = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListExtras) Validate() error {
	return dara.Validate(s)
}

type QueryJobsResponseBodyJobsListSummary struct {
	// Category
	//
	// example:
	//
	// {}
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Content
	//
	// example:
	//
	// 5
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Call record ID
	//
	// example:
	//
	// 098b9b09-9223-4a8b-a422-99726f0457f3
	ConversationDetailId *string `json:"ConversationDetailId,omitempty" xml:"ConversationDetailId,omitempty"`
	// Job ID.
	//
	// example:
	//
	// ba1ba502-d044-48c0-b710-0f1f840a7c53
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Job ID.
	//
	// example:
	//
	// b72425bd-7871-4050-838e-033d80d754b7
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Summary ID
	//
	// example:
	//
	// dc67d544-df06-4625-ae48-13e3c9f72d8a
	SummaryId *string `json:"SummaryId,omitempty" xml:"SummaryId,omitempty"`
	// Summary name.
	//
	// example:
	//
	// score
	SummaryName *string `json:"SummaryName,omitempty" xml:"SummaryName,omitempty"`
	// Call ID.
	//
	// example:
	//
	// 9fdf7a81-6781-4ab8-92fb-1d4231ef365e
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryJobsResponseBodyJobsListSummary) String() string {
	return dara.Prettify(s)
}

func (s QueryJobsResponseBodyJobsListSummary) GoString() string {
	return s.String()
}

func (s *QueryJobsResponseBodyJobsListSummary) GetCategory() *string {
	return s.Category
}

func (s *QueryJobsResponseBodyJobsListSummary) GetContent() *string {
	return s.Content
}

func (s *QueryJobsResponseBodyJobsListSummary) GetConversationDetailId() *string {
	return s.ConversationDetailId
}

func (s *QueryJobsResponseBodyJobsListSummary) GetGroupId() *string {
	return s.GroupId
}

func (s *QueryJobsResponseBodyJobsListSummary) GetJobId() *string {
	return s.JobId
}

func (s *QueryJobsResponseBodyJobsListSummary) GetSummaryId() *string {
	return s.SummaryId
}

func (s *QueryJobsResponseBodyJobsListSummary) GetSummaryName() *string {
	return s.SummaryName
}

func (s *QueryJobsResponseBodyJobsListSummary) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryJobsResponseBodyJobsListSummary) SetCategory(v string) *QueryJobsResponseBodyJobsListSummary {
	s.Category = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListSummary) SetContent(v string) *QueryJobsResponseBodyJobsListSummary {
	s.Content = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListSummary) SetConversationDetailId(v string) *QueryJobsResponseBodyJobsListSummary {
	s.ConversationDetailId = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListSummary) SetGroupId(v string) *QueryJobsResponseBodyJobsListSummary {
	s.GroupId = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListSummary) SetJobId(v string) *QueryJobsResponseBodyJobsListSummary {
	s.JobId = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListSummary) SetSummaryId(v string) *QueryJobsResponseBodyJobsListSummary {
	s.SummaryId = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListSummary) SetSummaryName(v string) *QueryJobsResponseBodyJobsListSummary {
	s.SummaryName = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListSummary) SetTaskId(v string) *QueryJobsResponseBodyJobsListSummary {
	s.TaskId = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListSummary) Validate() error {
	return dara.Validate(s)
}

type QueryJobsResponseBodyJobsListTagHits struct {
	// Tag Group name.
	//
	// example:
	//
	// 意向收集
	TagGroup *string `json:"TagGroup,omitempty" xml:"TagGroup,omitempty"`
	// Tag Name.
	//
	// example:
	//
	// 有意向
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s QueryJobsResponseBodyJobsListTagHits) String() string {
	return dara.Prettify(s)
}

func (s QueryJobsResponseBodyJobsListTagHits) GoString() string {
	return s.String()
}

func (s *QueryJobsResponseBodyJobsListTagHits) GetTagGroup() *string {
	return s.TagGroup
}

func (s *QueryJobsResponseBodyJobsListTagHits) GetTagName() *string {
	return s.TagName
}

func (s *QueryJobsResponseBodyJobsListTagHits) SetTagGroup(v string) *QueryJobsResponseBodyJobsListTagHits {
	s.TagGroup = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListTagHits) SetTagName(v string) *QueryJobsResponseBodyJobsListTagHits {
	s.TagName = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListTagHits) Validate() error {
	return dara.Validate(s)
}

type QueryJobsResponseBodyJobsListTasks struct {
	// Actual call time.
	//
	// example:
	//
	// 1579068424883
	ActualTime *int64 `json:"ActualTime,omitempty" xml:"ActualTime,omitempty"`
	// Summary (historical field, no longer used)
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
	// Chatbot ID.
	//
	// example:
	//
	// 1234
	ChatbotId *string `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	// Contact information.
	Contact *QueryJobsResponseBodyJobsListTasksContact `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Struct"`
	// Call duration.
	//
	// example:
	//
	// 120
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Job ID.
	//
	// example:
	//
	// b72425bd-7871-4050-838e-033d80d754b7
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Planned outbound call time.
	//
	// example:
	//
	// 1579068424883
	PlanedTime *int64 `json:"PlanedTime,omitempty" xml:"PlanedTime,omitempty"`
	// Scenario ID.
	//
	// example:
	//
	// ade80092-03d9-4f4d-ad4f-ab8a247d3150
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// Task Status. Valid values are as follows (note: the Succeeded status has been further categorized by reason; the generic "Succeeded: 1 (Answered)" status will not be displayed. Instead, specific sub-reasons will be returned):
	//
	// - Executing: 0 (Calling in progress).
	//
	// - Succeeded: 1 (Answered).
	//
	// - NoAnswer: 2 (Not answered – no one picked up).
	//
	// - NotExist: 3 (Not answered – nonexistent number).
	//
	// - Busy: 4 (Not answered – line busy).
	//
	// - Cancelled: 5 (Not dialed – job stopped).
	//
	// - Failed: 6 (Failed).
	//
	// - NotConnected: 7 (Not answered – unable to connect).
	//
	// - PoweredOff: 8 (Not answered – powered off).
	//
	// - OutOfService: 9 (Not answered – callee out of service).
	//
	// - InArrears: 10 (Not answered – callee has overdue payment).
	//
	// - EmptyNumber: 11 (Not dialed – nonexistent number, call not placed).
	//
	// - PerDayCallCountLimit: 12 (Not dialed – daily call limit exceeded).
	//
	// - ContactBlockList: 13 (Not dialed – on blacklist).
	//
	// - CallerNotRegistered: 14 (Not dialed – caller number not registered).
	//
	// - Terminated: 15 (Not dialed – terminated).
	//
	// - VerificationCancelled: 16 (Not dialed – canceled due to failed pre-call authentication).
	//
	// - OutOfServiceNoCall: 17 (Not dialed – callee out of service, call not placed).
	//
	// - InArrearsNoCall: 18 (Not dialed – callee has overdue payment, call not placed).
	//
	// - CallingNumberNotExist: 19 (Not dialed – caller number does not exist).
	//
	// - SucceededFinish: 20 (Answered – completed normally).
	//
	// - SucceededChatbotHangUpAfterNoAnswer: 21 (Answered – chatbot hung up after unrecognized input).
	//
	// - SucceededChatbotHangUpAfterSilence: 22 (Answered – chatbot hung up after silence timeout).
	//
	// - SucceededClientHangUpAfterNoAnswer: 23 (Answered – user hung up after unrecognized input).
	//
	// - SucceededClientHangUp: 24 (Answered – user hung up without reason).
	//
	// - SucceededTransferByIntent: 25 (Answered – transferred to agent due to intent hit).
	//
	// - SucceededTransferAfterNoAnswer: 26 (Answered – transferred to agent after unrecognized input).
	//
	// - SucceededInoInterAction: 27 (Answered – no interaction from user side).
	//
	// - SucceededError: 28 (Answered – interrupted due to system exception).
	//
	// - SucceededSpecialInterceptVoiceAssistant: 29 (Answered – intercepted under special condition – voice assistant).
	//
	// - SucceededSpecialInterceptExtensionNumberTransfer: 30 (Answered – intercepted under special condition – extension transfer).
	//
	// - SucceededSpecialInterceptCustomSpecialIntercept: 31 (Answered – intercepted under special condition – custom interception).
	//
	// - HighRiskSipCode: 32 (Not dialed – high-risk, call not placed).
	//
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Call ID.
	//
	// example:
	//
	// ff44709e-39a6-43ba-959b-20fcabe3e496
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryJobsResponseBodyJobsListTasks) String() string {
	return dara.Prettify(s)
}

func (s QueryJobsResponseBodyJobsListTasks) GoString() string {
	return s.String()
}

func (s *QueryJobsResponseBodyJobsListTasks) GetActualTime() *int64 {
	return s.ActualTime
}

func (s *QueryJobsResponseBodyJobsListTasks) GetBrief() *string {
	return s.Brief
}

func (s *QueryJobsResponseBodyJobsListTasks) GetCallId() *string {
	return s.CallId
}

func (s *QueryJobsResponseBodyJobsListTasks) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *QueryJobsResponseBodyJobsListTasks) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *QueryJobsResponseBodyJobsListTasks) GetChatbotId() *string {
	return s.ChatbotId
}

func (s *QueryJobsResponseBodyJobsListTasks) GetContact() *QueryJobsResponseBodyJobsListTasksContact {
	return s.Contact
}

func (s *QueryJobsResponseBodyJobsListTasks) GetDuration() *int32 {
	return s.Duration
}

func (s *QueryJobsResponseBodyJobsListTasks) GetJobId() *string {
	return s.JobId
}

func (s *QueryJobsResponseBodyJobsListTasks) GetPlanedTime() *int64 {
	return s.PlanedTime
}

func (s *QueryJobsResponseBodyJobsListTasks) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *QueryJobsResponseBodyJobsListTasks) GetStatus() *string {
	return s.Status
}

func (s *QueryJobsResponseBodyJobsListTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryJobsResponseBodyJobsListTasks) SetActualTime(v int64) *QueryJobsResponseBodyJobsListTasks {
	s.ActualTime = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListTasks) SetBrief(v string) *QueryJobsResponseBodyJobsListTasks {
	s.Brief = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListTasks) SetCallId(v string) *QueryJobsResponseBodyJobsListTasks {
	s.CallId = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListTasks) SetCalledNumber(v string) *QueryJobsResponseBodyJobsListTasks {
	s.CalledNumber = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListTasks) SetCallingNumber(v string) *QueryJobsResponseBodyJobsListTasks {
	s.CallingNumber = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListTasks) SetChatbotId(v string) *QueryJobsResponseBodyJobsListTasks {
	s.ChatbotId = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListTasks) SetContact(v *QueryJobsResponseBodyJobsListTasksContact) *QueryJobsResponseBodyJobsListTasks {
	s.Contact = v
	return s
}

func (s *QueryJobsResponseBodyJobsListTasks) SetDuration(v int32) *QueryJobsResponseBodyJobsListTasks {
	s.Duration = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListTasks) SetJobId(v string) *QueryJobsResponseBodyJobsListTasks {
	s.JobId = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListTasks) SetPlanedTime(v int64) *QueryJobsResponseBodyJobsListTasks {
	s.PlanedTime = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListTasks) SetScenarioId(v string) *QueryJobsResponseBodyJobsListTasks {
	s.ScenarioId = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListTasks) SetStatus(v string) *QueryJobsResponseBodyJobsListTasks {
	s.Status = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListTasks) SetTaskId(v string) *QueryJobsResponseBodyJobsListTasks {
	s.TaskId = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListTasks) Validate() error {
	if s.Contact != nil {
		if err := s.Contact.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryJobsResponseBodyJobsListTasksContact struct {
	// Contact ID.
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
	// Honorific.
	//
	// example:
	//
	// 张先生
	Honorific *string `json:"Honorific,omitempty" xml:"Honorific,omitempty"`
	// Job ID.
	//
	// example:
	//
	// b72425bd-7871-4050-838e-033d80d754b7
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Contact phone number.
	//
	// example:
	//
	// 135****8888
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Business association ID.
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

func (s QueryJobsResponseBodyJobsListTasksContact) String() string {
	return dara.Prettify(s)
}

func (s QueryJobsResponseBodyJobsListTasksContact) GoString() string {
	return s.String()
}

func (s *QueryJobsResponseBodyJobsListTasksContact) GetContactId() *string {
	return s.ContactId
}

func (s *QueryJobsResponseBodyJobsListTasksContact) GetContactName() *string {
	return s.ContactName
}

func (s *QueryJobsResponseBodyJobsListTasksContact) GetHonorific() *string {
	return s.Honorific
}

func (s *QueryJobsResponseBodyJobsListTasksContact) GetJobId() *string {
	return s.JobId
}

func (s *QueryJobsResponseBodyJobsListTasksContact) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *QueryJobsResponseBodyJobsListTasksContact) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *QueryJobsResponseBodyJobsListTasksContact) GetRole() *string {
	return s.Role
}

func (s *QueryJobsResponseBodyJobsListTasksContact) GetState() *string {
	return s.State
}

func (s *QueryJobsResponseBodyJobsListTasksContact) SetContactId(v string) *QueryJobsResponseBodyJobsListTasksContact {
	s.ContactId = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListTasksContact) SetContactName(v string) *QueryJobsResponseBodyJobsListTasksContact {
	s.ContactName = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListTasksContact) SetHonorific(v string) *QueryJobsResponseBodyJobsListTasksContact {
	s.Honorific = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListTasksContact) SetJobId(v string) *QueryJobsResponseBodyJobsListTasksContact {
	s.JobId = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListTasksContact) SetPhoneNumber(v string) *QueryJobsResponseBodyJobsListTasksContact {
	s.PhoneNumber = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListTasksContact) SetReferenceId(v string) *QueryJobsResponseBodyJobsListTasksContact {
	s.ReferenceId = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListTasksContact) SetRole(v string) *QueryJobsResponseBodyJobsListTasksContact {
	s.Role = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListTasksContact) SetState(v string) *QueryJobsResponseBodyJobsListTasksContact {
	s.State = &v
	return s
}

func (s *QueryJobsResponseBodyJobsListTasksContact) Validate() error {
	return dara.Validate(s)
}
