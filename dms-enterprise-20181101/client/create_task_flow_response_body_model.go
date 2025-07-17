// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *CreateTaskFlowResponseBody
	GetDagId() *int64
	SetErrorCode(v string) *CreateTaskFlowResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateTaskFlowResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateTaskFlowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTaskFlowResponseBody
	GetSuccess() *bool
}

type CreateTaskFlowResponseBody struct {
	// The ID of the task flow.
	//
	// example:
	//
	// 33***
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
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
	// User can not access to Tenant [1]
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 19DA51A9-AC3E-5C36-8351-07EBCD2B89A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTaskFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskFlowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskFlowResponseBody) GetDagId() *int64 {
	return s.DagId
}

func (s *CreateTaskFlowResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateTaskFlowResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateTaskFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTaskFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTaskFlowResponseBody) SetDagId(v int64) *CreateTaskFlowResponseBody {
	s.DagId = &v
	return s
}

func (s *CreateTaskFlowResponseBody) SetErrorCode(v string) *CreateTaskFlowResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateTaskFlowResponseBody) SetErrorMessage(v string) *CreateTaskFlowResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateTaskFlowResponseBody) SetRequestId(v string) *CreateTaskFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTaskFlowResponseBody) SetSuccess(v bool) *CreateTaskFlowResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTaskFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
