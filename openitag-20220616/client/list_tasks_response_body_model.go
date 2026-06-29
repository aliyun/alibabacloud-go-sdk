// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListTasksResponseBody
	GetCode() *int32
	SetDetails(v string) *ListTasksResponseBody
	GetDetails() *string
	SetErrorCode(v string) *ListTasksResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListTasksResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListTasksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTasksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTasksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTasksResponseBody
	GetSuccess() *bool
	SetTasks(v []*SimpleTask) *ListTasksResponseBody
	GetTasks() []*SimpleTask
	SetTotalCount(v int32) *ListTasksResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListTasksResponseBody
	GetTotalPage() *int32
}

type ListTasksResponseBody struct {
	// Total amount of data under the current request conditions. This parameter is optional and is not returned by default.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details
	//
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// Error code
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Return message of the request
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID
	//
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// List of jobs
	Tasks []*SimpleTask `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// Total count
	//
	// example:
	//
	// 22
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Total number of pages
	//
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListTasksResponseBody) GetDetails() *string {
	return s.Details
}

func (s *ListTasksResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTasksResponseBody) GetTasks() []*SimpleTask {
	return s.Tasks
}

func (s *ListTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTasksResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListTasksResponseBody) SetCode(v int32) *ListTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListTasksResponseBody) SetDetails(v string) *ListTasksResponseBody {
	s.Details = &v
	return s
}

func (s *ListTasksResponseBody) SetErrorCode(v string) *ListTasksResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTasksResponseBody) SetMessage(v string) *ListTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListTasksResponseBody) SetPageNumber(v int32) *ListTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTasksResponseBody) SetPageSize(v int32) *ListTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTasksResponseBody) SetRequestId(v string) *ListTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTasksResponseBody) SetSuccess(v bool) *ListTasksResponseBody {
	s.Success = &v
	return s
}

func (s *ListTasksResponseBody) SetTasks(v []*SimpleTask) *ListTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *ListTasksResponseBody) SetTotalCount(v int32) *ListTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTasksResponseBody) SetTotalPage(v int32) *ListTasksResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListTasksResponseBody) Validate() error {
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
