// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApprovalTasksByUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListApprovalTasksByUserResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListApprovalTasksByUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListApprovalTasksByUserResponseBody
	GetMessage() *string
	SetPageResult(v *ListApprovalTasksByUserResponseBodyPageResult) *ListApprovalTasksByUserResponseBody
	GetPageResult() *ListApprovalTasksByUserResponseBodyPageResult
	SetRequestId(v string) *ListApprovalTasksByUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListApprovalTasksByUserResponseBody
	GetSuccess() *bool
}

type ListApprovalTasksByUserResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend error details.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The paging query result.
	PageResult *ListApprovalTasksByUserResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListApprovalTasksByUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalTasksByUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListApprovalTasksByUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListApprovalTasksByUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListApprovalTasksByUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListApprovalTasksByUserResponseBody) GetPageResult() *ListApprovalTasksByUserResponseBodyPageResult {
	return s.PageResult
}

func (s *ListApprovalTasksByUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApprovalTasksByUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListApprovalTasksByUserResponseBody) SetCode(v string) *ListApprovalTasksByUserResponseBody {
	s.Code = &v
	return s
}

func (s *ListApprovalTasksByUserResponseBody) SetHttpStatusCode(v int32) *ListApprovalTasksByUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListApprovalTasksByUserResponseBody) SetMessage(v string) *ListApprovalTasksByUserResponseBody {
	s.Message = &v
	return s
}

func (s *ListApprovalTasksByUserResponseBody) SetPageResult(v *ListApprovalTasksByUserResponseBodyPageResult) *ListApprovalTasksByUserResponseBody {
	s.PageResult = v
	return s
}

func (s *ListApprovalTasksByUserResponseBody) SetRequestId(v string) *ListApprovalTasksByUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApprovalTasksByUserResponseBody) SetSuccess(v bool) *ListApprovalTasksByUserResponseBody {
	s.Success = &v
	return s
}

func (s *ListApprovalTasksByUserResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApprovalTasksByUserResponseBodyPageResult struct {
	// The list of approval tasks.
	TaskList []*ListApprovalTasksByUserResponseBodyPageResultTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApprovalTasksByUserResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalTasksByUserResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListApprovalTasksByUserResponseBodyPageResult) GetTaskList() []*ListApprovalTasksByUserResponseBodyPageResultTaskList {
	return s.TaskList
}

func (s *ListApprovalTasksByUserResponseBodyPageResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListApprovalTasksByUserResponseBodyPageResult) SetTaskList(v []*ListApprovalTasksByUserResponseBodyPageResultTaskList) *ListApprovalTasksByUserResponseBodyPageResult {
	s.TaskList = v
	return s
}

func (s *ListApprovalTasksByUserResponseBodyPageResult) SetTotalCount(v int64) *ListApprovalTasksByUserResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListApprovalTasksByUserResponseBodyPageResult) Validate() error {
	if s.TaskList != nil {
		for _, item := range s.TaskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApprovalTasksByUserResponseBodyPageResultTaskList struct {
	// The applicant ID.
	//
	// example:
	//
	// user001
	ApplicantId *string `json:"ApplicantId,omitempty" xml:"ApplicantId,omitempty"`
	// The applicant name.
	//
	// example:
	//
	// John
	ApplicantName *string `json:"ApplicantName,omitempty" xml:"ApplicantName,omitempty"`
	// The approval type.
	//
	// example:
	//
	// DATA_SOURCE
	ApprovalType *string `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	// The approval flow ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The relationship between the current user and the approval task.
	//
	// example:
	//
	// SUBMITTED
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// The source system.
	//
	// example:
	//
	// Dataphin
	SourceSystem *string `json:"SourceSystem,omitempty" xml:"SourceSystem,omitempty"`
	// The approval status.
	//
	// example:
	//
	// APPROVED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The submission time.
	//
	// example:
	//
	// 2025-07-30 10:00:00
	SubmittedAt *string `json:"SubmittedAt,omitempty" xml:"SubmittedAt,omitempty"`
	// The task name.
	//
	// example:
	//
	// Datasource application
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListApprovalTasksByUserResponseBodyPageResultTaskList) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalTasksByUserResponseBodyPageResultTaskList) GoString() string {
	return s.String()
}

func (s *ListApprovalTasksByUserResponseBodyPageResultTaskList) GetApplicantId() *string {
	return s.ApplicantId
}

func (s *ListApprovalTasksByUserResponseBodyPageResultTaskList) GetApplicantName() *string {
	return s.ApplicantName
}

func (s *ListApprovalTasksByUserResponseBodyPageResultTaskList) GetApprovalType() *string {
	return s.ApprovalType
}

func (s *ListApprovalTasksByUserResponseBodyPageResultTaskList) GetId() *int64 {
	return s.Id
}

func (s *ListApprovalTasksByUserResponseBodyPageResultTaskList) GetRelationType() *string {
	return s.RelationType
}

func (s *ListApprovalTasksByUserResponseBodyPageResultTaskList) GetSourceSystem() *string {
	return s.SourceSystem
}

func (s *ListApprovalTasksByUserResponseBodyPageResultTaskList) GetStatus() *string {
	return s.Status
}

func (s *ListApprovalTasksByUserResponseBodyPageResultTaskList) GetSubmittedAt() *string {
	return s.SubmittedAt
}

func (s *ListApprovalTasksByUserResponseBodyPageResultTaskList) GetTitle() *string {
	return s.Title
}

func (s *ListApprovalTasksByUserResponseBodyPageResultTaskList) SetApplicantId(v string) *ListApprovalTasksByUserResponseBodyPageResultTaskList {
	s.ApplicantId = &v
	return s
}

func (s *ListApprovalTasksByUserResponseBodyPageResultTaskList) SetApplicantName(v string) *ListApprovalTasksByUserResponseBodyPageResultTaskList {
	s.ApplicantName = &v
	return s
}

func (s *ListApprovalTasksByUserResponseBodyPageResultTaskList) SetApprovalType(v string) *ListApprovalTasksByUserResponseBodyPageResultTaskList {
	s.ApprovalType = &v
	return s
}

func (s *ListApprovalTasksByUserResponseBodyPageResultTaskList) SetId(v int64) *ListApprovalTasksByUserResponseBodyPageResultTaskList {
	s.Id = &v
	return s
}

func (s *ListApprovalTasksByUserResponseBodyPageResultTaskList) SetRelationType(v string) *ListApprovalTasksByUserResponseBodyPageResultTaskList {
	s.RelationType = &v
	return s
}

func (s *ListApprovalTasksByUserResponseBodyPageResultTaskList) SetSourceSystem(v string) *ListApprovalTasksByUserResponseBodyPageResultTaskList {
	s.SourceSystem = &v
	return s
}

func (s *ListApprovalTasksByUserResponseBodyPageResultTaskList) SetStatus(v string) *ListApprovalTasksByUserResponseBodyPageResultTaskList {
	s.Status = &v
	return s
}

func (s *ListApprovalTasksByUserResponseBodyPageResultTaskList) SetSubmittedAt(v string) *ListApprovalTasksByUserResponseBodyPageResultTaskList {
	s.SubmittedAt = &v
	return s
}

func (s *ListApprovalTasksByUserResponseBodyPageResultTaskList) SetTitle(v string) *ListApprovalTasksByUserResponseBodyPageResultTaskList {
	s.Title = &v
	return s
}

func (s *ListApprovalTasksByUserResponseBodyPageResultTaskList) Validate() error {
	return dara.Validate(s)
}
