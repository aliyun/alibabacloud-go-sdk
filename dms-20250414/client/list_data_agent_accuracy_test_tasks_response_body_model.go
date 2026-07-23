// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentAccuracyTestTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDataAgentAccuracyTestTasksResponseBodyData) *ListDataAgentAccuracyTestTasksResponseBody
	GetData() []*ListDataAgentAccuracyTestTasksResponseBodyData
	SetErrorCode(v string) *ListDataAgentAccuracyTestTasksResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataAgentAccuracyTestTasksResponseBody
	GetErrorMessage() *string
	SetMaxResults(v int32) *ListDataAgentAccuracyTestTasksResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataAgentAccuracyTestTasksResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *ListDataAgentAccuracyTestTasksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataAgentAccuracyTestTasksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDataAgentAccuracyTestTasksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataAgentAccuracyTestTasksResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListDataAgentAccuracyTestTasksResponseBody
	GetTotal() *int32
}

type ListDataAgentAccuracyTestTasksResponseBody struct {
	// The response struct.
	Data []*ListDataAgentAccuracyTestTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the call fails.
	//
	// example:
	//
	// Specified parameter Tid is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// NesLo****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E0D21075-xxx-FD8AD04A63B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 3
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDataAgentAccuracyTestTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentAccuracyTestTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataAgentAccuracyTestTasksResponseBody) GetData() []*ListDataAgentAccuracyTestTasksResponseBodyData {
	return s.Data
}

func (s *ListDataAgentAccuracyTestTasksResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataAgentAccuracyTestTasksResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataAgentAccuracyTestTasksResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataAgentAccuracyTestTasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataAgentAccuracyTestTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataAgentAccuracyTestTasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataAgentAccuracyTestTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataAgentAccuracyTestTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataAgentAccuracyTestTasksResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListDataAgentAccuracyTestTasksResponseBody) SetData(v []*ListDataAgentAccuracyTestTasksResponseBodyData) *ListDataAgentAccuracyTestTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListDataAgentAccuracyTestTasksResponseBody) SetErrorCode(v string) *ListDataAgentAccuracyTestTasksResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksResponseBody) SetErrorMessage(v string) *ListDataAgentAccuracyTestTasksResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksResponseBody) SetMaxResults(v int32) *ListDataAgentAccuracyTestTasksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksResponseBody) SetNextToken(v string) *ListDataAgentAccuracyTestTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksResponseBody) SetPageNumber(v int32) *ListDataAgentAccuracyTestTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksResponseBody) SetPageSize(v int32) *ListDataAgentAccuracyTestTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksResponseBody) SetRequestId(v string) *ListDataAgentAccuracyTestTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksResponseBody) SetSuccess(v bool) *ListDataAgentAccuracyTestTasksResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksResponseBody) SetTotal(v int32) *ListDataAgentAccuracyTestTasksResponseBody {
	s.Total = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataAgentAccuracyTestTasksResponseBodyData struct {
	// The accuracy test instance ID to which the task belongs.
	//
	// example:
	//
	// at-106n4rg17gv9fxxxxxxxxxx
	AccuracyTestInsId *string `json:"AccuracyTestInsId,omitempty" xml:"AccuracyTestInsId,omitempty"`
	// The accuracy test task ID.
	//
	// example:
	//
	// 692abb8f-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	AccuracyTestTaskId *string `json:"AccuracyTestTaskId,omitempty" xml:"AccuracyTestTaskId,omitempty"`
	// The number of completed test subtasks.
	//
	// example:
	//
	// 16
	FinishedTaskNumber *int32 `json:"FinishedTaskNumber,omitempty" xml:"FinishedTaskNumber,omitempty"`
	// The time when the document was created.
	//
	// example:
	//
	// 2026-06-30T08:03:30.000+00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the document was last modified.
	//
	// example:
	//
	// 2026-06-30T08:03:30.000+00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The maximum number of concurrent sessions during the test.
	//
	// example:
	//
	// 5
	MaxConcurrent *int32 `json:"MaxConcurrent,omitempty" xml:"MaxConcurrent,omitempty"`
	// Specifies whether sessions are displayed after analysis. This feature is not supported.
	//
	// example:
	//
	// true
	NeedDelete *bool `json:"NeedDelete,omitempty" xml:"NeedDelete,omitempty"`
	// The status of the custom agent.
	//
	// example:
	//
	// 0 PENDING,1 RUNNING,2 COMPLETED,3 STOPPED,4 FAILED
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of subtasks in the accuracy test task.
	//
	// example:
	//
	// 20
	TotalTaskNumber *int32 `json:"TotalTaskNumber,omitempty" xml:"TotalTaskNumber,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 8wfig6l33n4f4xxxxxxxxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDataAgentAccuracyTestTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentAccuracyTestTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataAgentAccuracyTestTasksResponseBodyData) GetAccuracyTestInsId() *string {
	return s.AccuracyTestInsId
}

func (s *ListDataAgentAccuracyTestTasksResponseBodyData) GetAccuracyTestTaskId() *string {
	return s.AccuracyTestTaskId
}

func (s *ListDataAgentAccuracyTestTasksResponseBodyData) GetFinishedTaskNumber() *int32 {
	return s.FinishedTaskNumber
}

func (s *ListDataAgentAccuracyTestTasksResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListDataAgentAccuracyTestTasksResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListDataAgentAccuracyTestTasksResponseBodyData) GetMaxConcurrent() *int32 {
	return s.MaxConcurrent
}

func (s *ListDataAgentAccuracyTestTasksResponseBodyData) GetNeedDelete() *bool {
	return s.NeedDelete
}

func (s *ListDataAgentAccuracyTestTasksResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *ListDataAgentAccuracyTestTasksResponseBodyData) GetTotalTaskNumber() *int32 {
	return s.TotalTaskNumber
}

func (s *ListDataAgentAccuracyTestTasksResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDataAgentAccuracyTestTasksResponseBodyData) SetAccuracyTestInsId(v string) *ListDataAgentAccuracyTestTasksResponseBodyData {
	s.AccuracyTestInsId = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksResponseBodyData) SetAccuracyTestTaskId(v string) *ListDataAgentAccuracyTestTasksResponseBodyData {
	s.AccuracyTestTaskId = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksResponseBodyData) SetFinishedTaskNumber(v int32) *ListDataAgentAccuracyTestTasksResponseBodyData {
	s.FinishedTaskNumber = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksResponseBodyData) SetGmtCreate(v string) *ListDataAgentAccuracyTestTasksResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksResponseBodyData) SetGmtModified(v string) *ListDataAgentAccuracyTestTasksResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksResponseBodyData) SetMaxConcurrent(v int32) *ListDataAgentAccuracyTestTasksResponseBodyData {
	s.MaxConcurrent = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksResponseBodyData) SetNeedDelete(v bool) *ListDataAgentAccuracyTestTasksResponseBodyData {
	s.NeedDelete = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksResponseBodyData) SetStatus(v int32) *ListDataAgentAccuracyTestTasksResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksResponseBodyData) SetTotalTaskNumber(v int32) *ListDataAgentAccuracyTestTasksResponseBodyData {
	s.TotalTaskNumber = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksResponseBodyData) SetWorkspaceId(v string) *ListDataAgentAccuracyTestTasksResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksResponseBodyData) Validate() error {
	return dara.Validate(s)
}
