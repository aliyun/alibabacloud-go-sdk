// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetTaskStatusResponseBody
	GetCode() *int32
	SetDetails(v string) *GetTaskStatusResponseBody
	GetDetails() *string
	SetErrorCode(v string) *GetTaskStatusResponseBody
	GetErrorCode() *string
	SetMessage(v string) *GetTaskStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTaskStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTaskStatusResponseBody
	GetSuccess() *bool
	SetTaskStatus(v string) *GetTaskStatusResponseBody
	GetTaskStatus() *string
}

type GetTaskStatusResponseBody struct {
	// Return encoding. The default value is 0, indicating normal execution.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details.
	//
	// example:
	//
	// ""
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
	// 90ABA848-AD74-1F6E-84BC-4182A7F1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Job status. Possible values:
	//
	// - INIT: initialization.
	//
	// - PROCESSING: processing.
	//
	// - SUCC: succeeded.
	//
	// - FAIL: failed.
	//
	// - DELETED: deleted.
	//
	// - OFFLINE: unpublished.
	//
	// - FINISHED: finished.
	//
	// example:
	//
	// SUCC
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GetTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTaskStatusResponseBody) GetDetails() *string {
	return s.Details
}

func (s *GetTaskStatusResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTaskStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTaskStatusResponseBody) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetTaskStatusResponseBody) SetCode(v int32) *GetTaskStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetDetails(v string) *GetTaskStatusResponseBody {
	s.Details = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetErrorCode(v string) *GetTaskStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetMessage(v string) *GetTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetRequestId(v string) *GetTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetSuccess(v bool) *GetTaskStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetTaskStatus(v string) *GetTaskStatusResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetTaskStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
