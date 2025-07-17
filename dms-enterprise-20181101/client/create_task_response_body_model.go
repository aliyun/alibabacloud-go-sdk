// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateTaskResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateTaskResponseBody
	GetErrorMessage() *string
	SetNodeId(v int64) *CreateTaskResponseBody
	GetNodeId() *int64
	SetRequestId(v string) *CreateTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTaskResponseBody
	GetSuccess() *bool
}

type CreateTaskResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// InvalidParameterValid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// InvalidParameterValid
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the task node returned when the task was created.
	//
	// example:
	//
	// 3***
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The ID of the request. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// CFD8FE00-36D9-4C1B-940D-65A7B73D9066
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateTaskResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateTaskResponseBody) GetNodeId() *int64 {
	return s.NodeId
}

func (s *CreateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTaskResponseBody) SetErrorCode(v string) *CreateTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateTaskResponseBody) SetErrorMessage(v string) *CreateTaskResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateTaskResponseBody) SetNodeId(v int64) *CreateTaskResponseBody {
	s.NodeId = &v
	return s
}

func (s *CreateTaskResponseBody) SetRequestId(v string) *CreateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTaskResponseBody) SetSuccess(v bool) *CreateTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
