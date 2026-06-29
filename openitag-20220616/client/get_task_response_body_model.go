// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetTaskResponseBody
	GetCode() *int32
	SetDetails(v string) *GetTaskResponseBody
	GetDetails() *string
	SetErrorCode(v string) *GetTaskResponseBody
	GetErrorCode() *string
	SetMessage(v string) *GetTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTaskResponseBody
	GetSuccess() *bool
	SetTask(v *TaskDetail) *GetTaskResponseBody
	GetTask() *TaskDetail
}

type GetTaskResponseBody struct {
	// The total amount of data under the conditions of this request. This parameter is optional and does not need to be returned by default.
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
	// The response message of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 0F01E603-8A9F-18ED-AD43-D52B5030****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded. Valid values:
	//
	// - true: The request succeeded.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Job.
	Task *TaskDetail `json:"Task,omitempty" xml:"Task,omitempty"`
}

func (s GetTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTaskResponseBody) GetDetails() *string {
	return s.Details
}

func (s *GetTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTaskResponseBody) GetTask() *TaskDetail {
	return s.Task
}

func (s *GetTaskResponseBody) SetCode(v int32) *GetTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskResponseBody) SetDetails(v string) *GetTaskResponseBody {
	s.Details = &v
	return s
}

func (s *GetTaskResponseBody) SetErrorCode(v string) *GetTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskResponseBody) SetMessage(v string) *GetTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskResponseBody) SetRequestId(v string) *GetTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResponseBody) SetSuccess(v bool) *GetTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskResponseBody) SetTask(v *TaskDetail) *GetTaskResponseBody {
	s.Task = v
	return s
}

func (s *GetTaskResponseBody) Validate() error {
	if s.Task != nil {
		if err := s.Task.Validate(); err != nil {
			return err
		}
	}
	return nil
}
