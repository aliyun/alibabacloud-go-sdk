// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTicketResponseBody
	GetCode() *string
	SetData(v *GetTicketResponseBodyData) *GetTicketResponseBody
	GetData() *GetTicketResponseBodyData
	SetHttpStatusCode(v int32) *GetTicketResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTicketResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetTicketResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetTicketResponseBody
	GetRequestId() *string
}

type GetTicketResponseBody struct {
	// example:
	//
	// OK
	Code *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTicketResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// BF268B34-09C2-43FD-BAC4-5D31EA633111
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTicketResponseBody) GetData() *GetTicketResponseBodyData {
	return s.Data
}

func (s *GetTicketResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTicketResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTicketResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTicketResponseBody) SetCode(v string) *GetTicketResponseBody {
	s.Code = &v
	return s
}

func (s *GetTicketResponseBody) SetData(v *GetTicketResponseBodyData) *GetTicketResponseBody {
	s.Data = v
	return s
}

func (s *GetTicketResponseBody) SetHttpStatusCode(v int32) *GetTicketResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTicketResponseBody) SetMessage(v string) *GetTicketResponseBody {
	s.Message = &v
	return s
}

func (s *GetTicketResponseBody) SetParams(v []*string) *GetTicketResponseBody {
	s.Params = v
	return s
}

func (s *GetTicketResponseBody) SetRequestId(v string) *GetTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTicketResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTicketResponseBodyData struct {
	// example:
	//
	// agent1@ccc-test
	Assignee     *string `json:"Assignee,omitempty" xml:"Assignee,omitempty"`
	AssigneeName *string `json:"AssigneeName,omitempty" xml:"AssigneeName,omitempty"`
	// example:
	//
	// 8939-4223-86d0-6bd187905cc8
	CategoryId   *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// example:
	//
	// Completed
	CloseCode *string `json:"CloseCode,omitempty" xml:"CloseCode,omitempty"`
	Comment   *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	Context   *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// example:
	//
	// 1620259200000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// creator@ccc-test
	Creator     *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// example:
	//
	// 912f0b78-6639-4a93-ae18-0d832885c27e
	CurrentTaskId   *string `json:"CurrentTaskId,omitempty" xml:"CurrentTaskId,omitempty"`
	CurrentTaskName *string `json:"CurrentTaskName,omitempty" xml:"CurrentTaskName,omitempty"`
	// example:
	//
	// 1693793208075
	CurrentTaskStartTime *int64 `json:"CurrentTaskStartTime,omitempty" xml:"CurrentTaskStartTime,omitempty"`
	// example:
	//
	// 4223-86d0-6bd187905-891798749
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// example:
	//
	// 1687846259999
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// job-399383842187575296
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// Audio
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 1620259200000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Processing
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// ccc-test_43c2671b-8939-4223-86d0-6bd187905cc8_*****0666238
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// 0
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// example:
	//
	// b3a6a131-359e-46bd-9bc5-1f5cb0ea093f
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	Title    *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 1693793208075
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s GetTicketResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTicketResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBodyData) GetAssignee() *string {
	return s.Assignee
}

func (s *GetTicketResponseBodyData) GetAssigneeName() *string {
	return s.AssigneeName
}

func (s *GetTicketResponseBodyData) GetCategoryId() *string {
	return s.CategoryId
}

func (s *GetTicketResponseBodyData) GetCategoryName() *string {
	return s.CategoryName
}

func (s *GetTicketResponseBodyData) GetCloseCode() *string {
	return s.CloseCode
}

func (s *GetTicketResponseBodyData) GetComment() *string {
	return s.Comment
}

func (s *GetTicketResponseBodyData) GetContext() *string {
	return s.Context
}

func (s *GetTicketResponseBodyData) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *GetTicketResponseBodyData) GetCreator() *string {
	return s.Creator
}

func (s *GetTicketResponseBodyData) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetTicketResponseBodyData) GetCurrentTaskId() *string {
	return s.CurrentTaskId
}

func (s *GetTicketResponseBodyData) GetCurrentTaskName() *string {
	return s.CurrentTaskName
}

func (s *GetTicketResponseBodyData) GetCurrentTaskStartTime() *int64 {
	return s.CurrentTaskStartTime
}

func (s *GetTicketResponseBodyData) GetCustomerId() *string {
	return s.CustomerId
}

func (s *GetTicketResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetTicketResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTicketResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *GetTicketResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *GetTicketResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetTicketResponseBodyData) GetState() *string {
	return s.State
}

func (s *GetTicketResponseBodyData) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTicketResponseBodyData) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *GetTicketResponseBodyData) GetTicketId() *string {
	return s.TicketId
}

func (s *GetTicketResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *GetTicketResponseBodyData) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *GetTicketResponseBodyData) SetAssignee(v string) *GetTicketResponseBodyData {
	s.Assignee = &v
	return s
}

func (s *GetTicketResponseBodyData) SetAssigneeName(v string) *GetTicketResponseBodyData {
	s.AssigneeName = &v
	return s
}

func (s *GetTicketResponseBodyData) SetCategoryId(v string) *GetTicketResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *GetTicketResponseBodyData) SetCategoryName(v string) *GetTicketResponseBodyData {
	s.CategoryName = &v
	return s
}

func (s *GetTicketResponseBodyData) SetCloseCode(v string) *GetTicketResponseBodyData {
	s.CloseCode = &v
	return s
}

func (s *GetTicketResponseBodyData) SetComment(v string) *GetTicketResponseBodyData {
	s.Comment = &v
	return s
}

func (s *GetTicketResponseBodyData) SetContext(v string) *GetTicketResponseBodyData {
	s.Context = &v
	return s
}

func (s *GetTicketResponseBodyData) SetCreatedTime(v int64) *GetTicketResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetTicketResponseBodyData) SetCreator(v string) *GetTicketResponseBodyData {
	s.Creator = &v
	return s
}

func (s *GetTicketResponseBodyData) SetCreatorName(v string) *GetTicketResponseBodyData {
	s.CreatorName = &v
	return s
}

func (s *GetTicketResponseBodyData) SetCurrentTaskId(v string) *GetTicketResponseBodyData {
	s.CurrentTaskId = &v
	return s
}

func (s *GetTicketResponseBodyData) SetCurrentTaskName(v string) *GetTicketResponseBodyData {
	s.CurrentTaskName = &v
	return s
}

func (s *GetTicketResponseBodyData) SetCurrentTaskStartTime(v int64) *GetTicketResponseBodyData {
	s.CurrentTaskStartTime = &v
	return s
}

func (s *GetTicketResponseBodyData) SetCustomerId(v string) *GetTicketResponseBodyData {
	s.CustomerId = &v
	return s
}

func (s *GetTicketResponseBodyData) SetEndTime(v int64) *GetTicketResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetTicketResponseBodyData) SetInstanceId(v string) *GetTicketResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetTicketResponseBodyData) SetJobId(v string) *GetTicketResponseBodyData {
	s.JobId = &v
	return s
}

func (s *GetTicketResponseBodyData) SetSource(v string) *GetTicketResponseBodyData {
	s.Source = &v
	return s
}

func (s *GetTicketResponseBodyData) SetStartTime(v int64) *GetTicketResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetTicketResponseBodyData) SetState(v string) *GetTicketResponseBodyData {
	s.State = &v
	return s
}

func (s *GetTicketResponseBodyData) SetTemplateId(v string) *GetTicketResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *GetTicketResponseBodyData) SetTemplateVersion(v string) *GetTicketResponseBodyData {
	s.TemplateVersion = &v
	return s
}

func (s *GetTicketResponseBodyData) SetTicketId(v string) *GetTicketResponseBodyData {
	s.TicketId = &v
	return s
}

func (s *GetTicketResponseBodyData) SetTitle(v string) *GetTicketResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetTicketResponseBodyData) SetUpdatedTime(v int64) *GetTicketResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *GetTicketResponseBodyData) Validate() error {
	return dara.Validate(s)
}
