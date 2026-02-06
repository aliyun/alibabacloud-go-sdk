// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsByGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListJobsByGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListJobsByGroupResponseBody
	GetHttpStatusCode() *int32
	SetJobs(v *ListJobsByGroupResponseBodyJobs) *ListJobsByGroupResponseBody
	GetJobs() *ListJobsByGroupResponseBodyJobs
	SetMessage(v string) *ListJobsByGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListJobsByGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListJobsByGroupResponseBody
	GetSuccess() *bool
}

type ListJobsByGroupResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Jobs           *ListJobsByGroupResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Struct"`
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

func (s ListJobsByGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobsByGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsByGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListJobsByGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListJobsByGroupResponseBody) GetJobs() *ListJobsByGroupResponseBodyJobs {
	return s.Jobs
}

func (s *ListJobsByGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListJobsByGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobsByGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListJobsByGroupResponseBody) SetCode(v string) *ListJobsByGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListJobsByGroupResponseBody) SetHttpStatusCode(v int32) *ListJobsByGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListJobsByGroupResponseBody) SetJobs(v *ListJobsByGroupResponseBodyJobs) *ListJobsByGroupResponseBody {
	s.Jobs = v
	return s
}

func (s *ListJobsByGroupResponseBody) SetMessage(v string) *ListJobsByGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListJobsByGroupResponseBody) SetRequestId(v string) *ListJobsByGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobsByGroupResponseBody) SetSuccess(v bool) *ListJobsByGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ListJobsByGroupResponseBody) Validate() error {
	if s.Jobs != nil {
		if err := s.Jobs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobsByGroupResponseBodyJobs struct {
	List []*ListJobsByGroupResponseBodyJobsList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJobsByGroupResponseBodyJobs) String() string {
	return dara.Prettify(s)
}

func (s ListJobsByGroupResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListJobsByGroupResponseBodyJobs) GetList() []*ListJobsByGroupResponseBodyJobsList {
	return s.List
}

func (s *ListJobsByGroupResponseBodyJobs) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobsByGroupResponseBodyJobs) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobsByGroupResponseBodyJobs) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListJobsByGroupResponseBodyJobs) SetList(v []*ListJobsByGroupResponseBodyJobsList) *ListJobsByGroupResponseBodyJobs {
	s.List = v
	return s
}

func (s *ListJobsByGroupResponseBodyJobs) SetPageNumber(v int32) *ListJobsByGroupResponseBodyJobs {
	s.PageNumber = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobs) SetPageSize(v int32) *ListJobsByGroupResponseBodyJobs {
	s.PageSize = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobs) SetTotalCount(v int32) *ListJobsByGroupResponseBodyJobs {
	s.TotalCount = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobs) Validate() error {
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

type ListJobsByGroupResponseBodyJobsList struct {
	CallingNumbers []*string                                      `json:"CallingNumbers,omitempty" xml:"CallingNumbers,omitempty" type:"Repeated"`
	Contacts       []*ListJobsByGroupResponseBodyJobsListContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Repeated"`
	Extras         []*ListJobsByGroupResponseBodyJobsListExtras   `json:"Extras,omitempty" xml:"Extras,omitempty" type:"Repeated"`
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
	// ade80092-03d9-4f4d-ad4f-ab8a247d3150
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// c8a2b7f2-ad1a-4865-b872-d0080d9802d9
	StrategyId *string                                       `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	Summary    []*ListJobsByGroupResponseBodyJobsListSummary `json:"Summary,omitempty" xml:"Summary,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	SystemPriority *int32 `json:"SystemPriority,omitempty" xml:"SystemPriority,omitempty"`
}

func (s ListJobsByGroupResponseBodyJobsList) String() string {
	return dara.Prettify(s)
}

func (s ListJobsByGroupResponseBodyJobsList) GoString() string {
	return s.String()
}

func (s *ListJobsByGroupResponseBodyJobsList) GetCallingNumbers() []*string {
	return s.CallingNumbers
}

func (s *ListJobsByGroupResponseBodyJobsList) GetContacts() []*ListJobsByGroupResponseBodyJobsListContacts {
	return s.Contacts
}

func (s *ListJobsByGroupResponseBodyJobsList) GetExtras() []*ListJobsByGroupResponseBodyJobsListExtras {
	return s.Extras
}

func (s *ListJobsByGroupResponseBodyJobsList) GetFailureReason() *string {
	return s.FailureReason
}

func (s *ListJobsByGroupResponseBodyJobsList) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *ListJobsByGroupResponseBodyJobsList) GetJobId() *string {
	return s.JobId
}

func (s *ListJobsByGroupResponseBodyJobsList) GetPriority() *int32 {
	return s.Priority
}

func (s *ListJobsByGroupResponseBodyJobsList) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *ListJobsByGroupResponseBodyJobsList) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *ListJobsByGroupResponseBodyJobsList) GetStatus() *string {
	return s.Status
}

func (s *ListJobsByGroupResponseBodyJobsList) GetStrategyId() *string {
	return s.StrategyId
}

func (s *ListJobsByGroupResponseBodyJobsList) GetSummary() []*ListJobsByGroupResponseBodyJobsListSummary {
	return s.Summary
}

func (s *ListJobsByGroupResponseBodyJobsList) GetSystemPriority() *int32 {
	return s.SystemPriority
}

func (s *ListJobsByGroupResponseBodyJobsList) SetCallingNumbers(v []*string) *ListJobsByGroupResponseBodyJobsList {
	s.CallingNumbers = v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsList) SetContacts(v []*ListJobsByGroupResponseBodyJobsListContacts) *ListJobsByGroupResponseBodyJobsList {
	s.Contacts = v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsList) SetExtras(v []*ListJobsByGroupResponseBodyJobsListExtras) *ListJobsByGroupResponseBodyJobsList {
	s.Extras = v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsList) SetFailureReason(v string) *ListJobsByGroupResponseBodyJobsList {
	s.FailureReason = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsList) SetJobGroupId(v string) *ListJobsByGroupResponseBodyJobsList {
	s.JobGroupId = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsList) SetJobId(v string) *ListJobsByGroupResponseBodyJobsList {
	s.JobId = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsList) SetPriority(v int32) *ListJobsByGroupResponseBodyJobsList {
	s.Priority = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsList) SetReferenceId(v string) *ListJobsByGroupResponseBodyJobsList {
	s.ReferenceId = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsList) SetScenarioId(v string) *ListJobsByGroupResponseBodyJobsList {
	s.ScenarioId = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsList) SetStatus(v string) *ListJobsByGroupResponseBodyJobsList {
	s.Status = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsList) SetStrategyId(v string) *ListJobsByGroupResponseBodyJobsList {
	s.StrategyId = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsList) SetSummary(v []*ListJobsByGroupResponseBodyJobsListSummary) *ListJobsByGroupResponseBodyJobsList {
	s.Summary = v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsList) SetSystemPriority(v int32) *ListJobsByGroupResponseBodyJobsList {
	s.SystemPriority = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsList) Validate() error {
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
	return nil
}

type ListJobsByGroupResponseBodyJobsListContacts struct {
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

func (s ListJobsByGroupResponseBodyJobsListContacts) String() string {
	return dara.Prettify(s)
}

func (s ListJobsByGroupResponseBodyJobsListContacts) GoString() string {
	return s.String()
}

func (s *ListJobsByGroupResponseBodyJobsListContacts) GetContactId() *string {
	return s.ContactId
}

func (s *ListJobsByGroupResponseBodyJobsListContacts) GetContactName() *string {
	return s.ContactName
}

func (s *ListJobsByGroupResponseBodyJobsListContacts) GetHonorific() *string {
	return s.Honorific
}

func (s *ListJobsByGroupResponseBodyJobsListContacts) GetJobId() *string {
	return s.JobId
}

func (s *ListJobsByGroupResponseBodyJobsListContacts) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ListJobsByGroupResponseBodyJobsListContacts) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *ListJobsByGroupResponseBodyJobsListContacts) GetRole() *string {
	return s.Role
}

func (s *ListJobsByGroupResponseBodyJobsListContacts) GetState() *string {
	return s.State
}

func (s *ListJobsByGroupResponseBodyJobsListContacts) SetContactId(v string) *ListJobsByGroupResponseBodyJobsListContacts {
	s.ContactId = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsListContacts) SetContactName(v string) *ListJobsByGroupResponseBodyJobsListContacts {
	s.ContactName = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsListContacts) SetHonorific(v string) *ListJobsByGroupResponseBodyJobsListContacts {
	s.Honorific = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsListContacts) SetJobId(v string) *ListJobsByGroupResponseBodyJobsListContacts {
	s.JobId = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsListContacts) SetPhoneNumber(v string) *ListJobsByGroupResponseBodyJobsListContacts {
	s.PhoneNumber = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsListContacts) SetReferenceId(v string) *ListJobsByGroupResponseBodyJobsListContacts {
	s.ReferenceId = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsListContacts) SetRole(v string) *ListJobsByGroupResponseBodyJobsListContacts {
	s.Role = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsListContacts) SetState(v string) *ListJobsByGroupResponseBodyJobsListContacts {
	s.State = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsListContacts) Validate() error {
	return dara.Validate(s)
}

type ListJobsByGroupResponseBodyJobsListExtras struct {
	// example:
	//
	// djrq
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 2019-08-21 09:49:59.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListJobsByGroupResponseBodyJobsListExtras) String() string {
	return dara.Prettify(s)
}

func (s ListJobsByGroupResponseBodyJobsListExtras) GoString() string {
	return s.String()
}

func (s *ListJobsByGroupResponseBodyJobsListExtras) GetKey() *string {
	return s.Key
}

func (s *ListJobsByGroupResponseBodyJobsListExtras) GetValue() *string {
	return s.Value
}

func (s *ListJobsByGroupResponseBodyJobsListExtras) SetKey(v string) *ListJobsByGroupResponseBodyJobsListExtras {
	s.Key = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsListExtras) SetValue(v string) *ListJobsByGroupResponseBodyJobsListExtras {
	s.Value = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsListExtras) Validate() error {
	return dara.Validate(s)
}

type ListJobsByGroupResponseBodyJobsListSummary struct {
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
	// 62a860f5-a8b3-4b75-9512-c7e04bb7c8bc
	ConversationDetailId *string `json:"ConversationDetailId,omitempty" xml:"ConversationDetailId,omitempty"`
	// example:
	//
	// 88e56cfb-33f8-477a-907c-0fe83292d924
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// example:
	//
	// f102a853-5f5a-47fb-8869-b31ea74a9620
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 680f1905-81ae-4aab-99dd-2964dfb767fa
	SummaryId *string `json:"SummaryId,omitempty" xml:"SummaryId,omitempty"`
	// example:
	//
	// score
	SummaryName *string `json:"SummaryName,omitempty" xml:"SummaryName,omitempty"`
	// example:
	//
	// b0f35dd1-0337-402e-9c4f-3a6c2426950a
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListJobsByGroupResponseBodyJobsListSummary) String() string {
	return dara.Prettify(s)
}

func (s ListJobsByGroupResponseBodyJobsListSummary) GoString() string {
	return s.String()
}

func (s *ListJobsByGroupResponseBodyJobsListSummary) GetCategory() *string {
	return s.Category
}

func (s *ListJobsByGroupResponseBodyJobsListSummary) GetContent() *string {
	return s.Content
}

func (s *ListJobsByGroupResponseBodyJobsListSummary) GetConversationDetailId() *string {
	return s.ConversationDetailId
}

func (s *ListJobsByGroupResponseBodyJobsListSummary) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *ListJobsByGroupResponseBodyJobsListSummary) GetJobId() *string {
	return s.JobId
}

func (s *ListJobsByGroupResponseBodyJobsListSummary) GetSummaryId() *string {
	return s.SummaryId
}

func (s *ListJobsByGroupResponseBodyJobsListSummary) GetSummaryName() *string {
	return s.SummaryName
}

func (s *ListJobsByGroupResponseBodyJobsListSummary) GetTaskId() *string {
	return s.TaskId
}

func (s *ListJobsByGroupResponseBodyJobsListSummary) SetCategory(v string) *ListJobsByGroupResponseBodyJobsListSummary {
	s.Category = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsListSummary) SetContent(v string) *ListJobsByGroupResponseBodyJobsListSummary {
	s.Content = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsListSummary) SetConversationDetailId(v string) *ListJobsByGroupResponseBodyJobsListSummary {
	s.ConversationDetailId = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsListSummary) SetJobGroupId(v string) *ListJobsByGroupResponseBodyJobsListSummary {
	s.JobGroupId = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsListSummary) SetJobId(v string) *ListJobsByGroupResponseBodyJobsListSummary {
	s.JobId = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsListSummary) SetSummaryId(v string) *ListJobsByGroupResponseBodyJobsListSummary {
	s.SummaryId = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsListSummary) SetSummaryName(v string) *ListJobsByGroupResponseBodyJobsListSummary {
	s.SummaryName = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsListSummary) SetTaskId(v string) *ListJobsByGroupResponseBodyJobsListSummary {
	s.TaskId = &v
	return s
}

func (s *ListJobsByGroupResponseBodyJobsListSummary) Validate() error {
	return dara.Validate(s)
}
