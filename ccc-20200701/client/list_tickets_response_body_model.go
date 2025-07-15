// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTicketsResponseBody
	GetCode() *string
	SetData(v *ListTicketsResponseBodyData) *ListTicketsResponseBody
	GetData() *ListTicketsResponseBodyData
	SetHttpStatusCode(v int64) *ListTicketsResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *ListTicketsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListTicketsResponseBody
	GetRequestId() *string
}

type ListTicketsResponseBody struct {
	// example:
	//
	// OK
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListTicketsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7BEEA660-A45A-45E3-98CC-AFC65E715C23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTicketsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTicketsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTicketsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTicketsResponseBody) GetData() *ListTicketsResponseBodyData {
	return s.Data
}

func (s *ListTicketsResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ListTicketsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTicketsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTicketsResponseBody) SetCode(v string) *ListTicketsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTicketsResponseBody) SetData(v *ListTicketsResponseBodyData) *ListTicketsResponseBody {
	s.Data = v
	return s
}

func (s *ListTicketsResponseBody) SetHttpStatusCode(v int64) *ListTicketsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTicketsResponseBody) SetMessage(v string) *ListTicketsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTicketsResponseBody) SetRequestId(v string) *ListTicketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTicketsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTicketsResponseBodyData struct {
	List []*ListTicketsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTicketsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTicketsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTicketsResponseBodyData) GetList() []*ListTicketsResponseBodyDataList {
	return s.List
}

func (s *ListTicketsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListTicketsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTicketsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTicketsResponseBodyData) SetList(v []*ListTicketsResponseBodyDataList) *ListTicketsResponseBodyData {
	s.List = v
	return s
}

func (s *ListTicketsResponseBodyData) SetPageNumber(v int64) *ListTicketsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListTicketsResponseBodyData) SetPageSize(v int64) *ListTicketsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListTicketsResponseBodyData) SetTotalCount(v int64) *ListTicketsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListTicketsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListTicketsResponseBodyDataList struct {
	// example:
	//
	// assignee@ccc-test
	Assignee *string `json:"Assignee,omitempty" xml:"Assignee,omitempty"`
	// example:
	//
	// Assignee
	AssigneeName *string `json:"AssigneeName,omitempty" xml:"AssigneeName,omitempty"`
	// example:
	//
	// 43c2671b-*****-4223-86d0-6bd187905cc8
	CategoryId   *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// example:
	//
	// Completed
	CloseCode *string `json:"CloseCode,omitempty" xml:"CloseCode,omitempty"`
	Comment   *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// {"productName":"alynx"}
	Context *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// example:
	//
	// 1631440860000
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// creator@ccc-test
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// Creator
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// example:
	//
	// 0707dab6-34a8-11ef-9823-161e3802b2d4
	CurrentTaskId   *string `json:"CurrentTaskId,omitempty" xml:"CurrentTaskId,omitempty"`
	CurrentTaskName *string `json:"CurrentTaskName,omitempty" xml:"CurrentTaskName,omitempty"`
	// example:
	//
	// 1631440860000
	CurrentTaskStartTime *int64 `json:"CurrentTaskStartTime,omitempty" xml:"CurrentTaskStartTime,omitempty"`
	// example:
	//
	// 51e155ce-*****1-b402-13c69597b920
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// example:
	//
	// 1631440860000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// job-47150***150396416
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// CHAT
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 1631440860000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Processing
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// c844a5f0-496c-4c5b-8a0c-dd27686e8ff6
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// 0
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// example:
	//
	// feb83abd-9f08-49d2-9b56-41d1b66ca0ac
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	Title    *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 1631440860000
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s ListTicketsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListTicketsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListTicketsResponseBodyDataList) GetAssignee() *string {
	return s.Assignee
}

func (s *ListTicketsResponseBodyDataList) GetAssigneeName() *string {
	return s.AssigneeName
}

func (s *ListTicketsResponseBodyDataList) GetCategoryId() *string {
	return s.CategoryId
}

func (s *ListTicketsResponseBodyDataList) GetCategoryName() *string {
	return s.CategoryName
}

func (s *ListTicketsResponseBodyDataList) GetCloseCode() *string {
	return s.CloseCode
}

func (s *ListTicketsResponseBodyDataList) GetComment() *string {
	return s.Comment
}

func (s *ListTicketsResponseBodyDataList) GetContext() *string {
	return s.Context
}

func (s *ListTicketsResponseBodyDataList) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListTicketsResponseBodyDataList) GetCreator() *string {
	return s.Creator
}

func (s *ListTicketsResponseBodyDataList) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListTicketsResponseBodyDataList) GetCurrentTaskId() *string {
	return s.CurrentTaskId
}

func (s *ListTicketsResponseBodyDataList) GetCurrentTaskName() *string {
	return s.CurrentTaskName
}

func (s *ListTicketsResponseBodyDataList) GetCurrentTaskStartTime() *int64 {
	return s.CurrentTaskStartTime
}

func (s *ListTicketsResponseBodyDataList) GetCustomerId() *string {
	return s.CustomerId
}

func (s *ListTicketsResponseBodyDataList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListTicketsResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTicketsResponseBodyDataList) GetJobId() *string {
	return s.JobId
}

func (s *ListTicketsResponseBodyDataList) GetSource() *string {
	return s.Source
}

func (s *ListTicketsResponseBodyDataList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListTicketsResponseBodyDataList) GetState() *string {
	return s.State
}

func (s *ListTicketsResponseBodyDataList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListTicketsResponseBodyDataList) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *ListTicketsResponseBodyDataList) GetTicketId() *string {
	return s.TicketId
}

func (s *ListTicketsResponseBodyDataList) GetTitle() *string {
	return s.Title
}

func (s *ListTicketsResponseBodyDataList) GetUpdatedTime() *string {
	return s.UpdatedTime
}

func (s *ListTicketsResponseBodyDataList) SetAssignee(v string) *ListTicketsResponseBodyDataList {
	s.Assignee = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetAssigneeName(v string) *ListTicketsResponseBodyDataList {
	s.AssigneeName = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetCategoryId(v string) *ListTicketsResponseBodyDataList {
	s.CategoryId = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetCategoryName(v string) *ListTicketsResponseBodyDataList {
	s.CategoryName = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetCloseCode(v string) *ListTicketsResponseBodyDataList {
	s.CloseCode = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetComment(v string) *ListTicketsResponseBodyDataList {
	s.Comment = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetContext(v string) *ListTicketsResponseBodyDataList {
	s.Context = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetCreatedTime(v string) *ListTicketsResponseBodyDataList {
	s.CreatedTime = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetCreator(v string) *ListTicketsResponseBodyDataList {
	s.Creator = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetCreatorName(v string) *ListTicketsResponseBodyDataList {
	s.CreatorName = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetCurrentTaskId(v string) *ListTicketsResponseBodyDataList {
	s.CurrentTaskId = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetCurrentTaskName(v string) *ListTicketsResponseBodyDataList {
	s.CurrentTaskName = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetCurrentTaskStartTime(v int64) *ListTicketsResponseBodyDataList {
	s.CurrentTaskStartTime = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetCustomerId(v string) *ListTicketsResponseBodyDataList {
	s.CustomerId = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetEndTime(v int64) *ListTicketsResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetInstanceId(v string) *ListTicketsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetJobId(v string) *ListTicketsResponseBodyDataList {
	s.JobId = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetSource(v string) *ListTicketsResponseBodyDataList {
	s.Source = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetStartTime(v int64) *ListTicketsResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetState(v string) *ListTicketsResponseBodyDataList {
	s.State = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetTemplateId(v string) *ListTicketsResponseBodyDataList {
	s.TemplateId = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetTemplateVersion(v string) *ListTicketsResponseBodyDataList {
	s.TemplateVersion = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetTicketId(v string) *ListTicketsResponseBodyDataList {
	s.TicketId = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetTitle(v string) *ListTicketsResponseBodyDataList {
	s.Title = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetUpdatedTime(v string) *ListTicketsResponseBodyDataList {
	s.UpdatedTime = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
