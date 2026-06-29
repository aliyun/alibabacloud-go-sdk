// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubtasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListSubtasksResponseBody
	GetCode() *int32
	SetDetails(v string) *ListSubtasksResponseBody
	GetDetails() *string
	SetErrorCode(v string) *ListSubtasksResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListSubtasksResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListSubtasksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSubtasksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListSubtasksResponseBody
	GetRequestId() *string
	SetSubtasks(v []*SubtaskDetail) *ListSubtasksResponseBody
	GetSubtasks() []*SubtaskDetail
	SetSuccess(v bool) *ListSubtasksResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListSubtasksResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListSubtasksResponseBody
	GetTotalPage() *int32
}

type ListSubtasksResponseBody struct {
	// Return code. The default value is 0, indicating normal execution.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details.
	//
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// Error code.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Response message of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Page number of the returned subtask list.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of subtasks displayed per page in the response.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of subtasks.
	Subtasks []*SubtaskDetail `json:"Subtasks,omitempty" xml:"Subtasks,omitempty" type:"Repeated"`
	// Indicates whether the operation succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Total number of subtasks.
	//
	// example:
	//
	// 22
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListSubtasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSubtasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubtasksResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListSubtasksResponseBody) GetDetails() *string {
	return s.Details
}

func (s *ListSubtasksResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListSubtasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSubtasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSubtasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSubtasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSubtasksResponseBody) GetSubtasks() []*SubtaskDetail {
	return s.Subtasks
}

func (s *ListSubtasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSubtasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSubtasksResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListSubtasksResponseBody) SetCode(v int32) *ListSubtasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListSubtasksResponseBody) SetDetails(v string) *ListSubtasksResponseBody {
	s.Details = &v
	return s
}

func (s *ListSubtasksResponseBody) SetErrorCode(v string) *ListSubtasksResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSubtasksResponseBody) SetMessage(v string) *ListSubtasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListSubtasksResponseBody) SetPageNumber(v int32) *ListSubtasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSubtasksResponseBody) SetPageSize(v int32) *ListSubtasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSubtasksResponseBody) SetRequestId(v string) *ListSubtasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubtasksResponseBody) SetSubtasks(v []*SubtaskDetail) *ListSubtasksResponseBody {
	s.Subtasks = v
	return s
}

func (s *ListSubtasksResponseBody) SetSuccess(v bool) *ListSubtasksResponseBody {
	s.Success = &v
	return s
}

func (s *ListSubtasksResponseBody) SetTotalCount(v int32) *ListSubtasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSubtasksResponseBody) SetTotalPage(v int32) *ListSubtasksResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListSubtasksResponseBody) Validate() error {
	if s.Subtasks != nil {
		for _, item := range s.Subtasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
