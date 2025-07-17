// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitStructSyncOrderApprovalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *SubmitStructSyncOrderApprovalResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SubmitStructSyncOrderApprovalResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *SubmitStructSyncOrderApprovalResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitStructSyncOrderApprovalResponseBody
	GetSuccess() *bool
	SetWorkflowInstanceId(v int64) *SubmitStructSyncOrderApprovalResponseBody
	GetWorkflowInstanceId() *int64
}

type SubmitStructSyncOrderApprovalResponseBody struct {
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4E1D2B4D-3E53-4ABC-999D-1D2520B3471D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the approval process.
	//
	// example:
	//
	// 432523
	WorkflowInstanceId *int64 `json:"WorkflowInstanceId,omitempty" xml:"WorkflowInstanceId,omitempty"`
}

func (s SubmitStructSyncOrderApprovalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitStructSyncOrderApprovalResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitStructSyncOrderApprovalResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SubmitStructSyncOrderApprovalResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SubmitStructSyncOrderApprovalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitStructSyncOrderApprovalResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitStructSyncOrderApprovalResponseBody) GetWorkflowInstanceId() *int64 {
	return s.WorkflowInstanceId
}

func (s *SubmitStructSyncOrderApprovalResponseBody) SetErrorCode(v string) *SubmitStructSyncOrderApprovalResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitStructSyncOrderApprovalResponseBody) SetErrorMessage(v string) *SubmitStructSyncOrderApprovalResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SubmitStructSyncOrderApprovalResponseBody) SetRequestId(v string) *SubmitStructSyncOrderApprovalResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitStructSyncOrderApprovalResponseBody) SetSuccess(v bool) *SubmitStructSyncOrderApprovalResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitStructSyncOrderApprovalResponseBody) SetWorkflowInstanceId(v int64) *SubmitStructSyncOrderApprovalResponseBody {
	s.WorkflowInstanceId = &v
	return s
}

func (s *SubmitStructSyncOrderApprovalResponseBody) Validate() error {
	return dara.Validate(s)
}
