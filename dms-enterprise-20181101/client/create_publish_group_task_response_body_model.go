// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePublishGroupTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreatePublishGroupTaskResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreatePublishGroupTaskResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreatePublishGroupTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreatePublishGroupTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v int64) *CreatePublishGroupTaskResponseBody
	GetTaskId() *int64
}

type CreatePublishGroupTaskResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 427688B8-ADFB-4C4E-9D45-EF5C1FD6E23D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the job.
	//
	// example:
	//
	// 413452
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreatePublishGroupTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePublishGroupTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePublishGroupTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreatePublishGroupTaskResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreatePublishGroupTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePublishGroupTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreatePublishGroupTaskResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CreatePublishGroupTaskResponseBody) SetErrorCode(v string) *CreatePublishGroupTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreatePublishGroupTaskResponseBody) SetErrorMessage(v string) *CreatePublishGroupTaskResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreatePublishGroupTaskResponseBody) SetRequestId(v string) *CreatePublishGroupTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePublishGroupTaskResponseBody) SetSuccess(v bool) *CreatePublishGroupTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePublishGroupTaskResponseBody) SetTaskId(v int64) *CreatePublishGroupTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreatePublishGroupTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
