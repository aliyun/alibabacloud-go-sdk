// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAbacPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateAbacPolicyResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateAbacPolicyResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateAbacPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAbacPolicyResponseBody
	GetSuccess() *bool
	SetUpdatePolicyResult(v int64) *UpdateAbacPolicyResponseBody
	GetUpdatePolicyResult() *int64
}

type UpdateAbacPolicyResponseBody struct {
	// The error code that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// CE43759B-5A72-560A-BF3D-862F38B36B9E
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
	// Indicates whether the policy is updated.
	//
	// example:
	//
	// true
	UpdatePolicyResult *int64 `json:"UpdatePolicyResult,omitempty" xml:"UpdatePolicyResult,omitempty"`
}

func (s UpdateAbacPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAbacPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAbacPolicyResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateAbacPolicyResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateAbacPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAbacPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAbacPolicyResponseBody) GetUpdatePolicyResult() *int64 {
	return s.UpdatePolicyResult
}

func (s *UpdateAbacPolicyResponseBody) SetErrorCode(v string) *UpdateAbacPolicyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateAbacPolicyResponseBody) SetErrorMessage(v string) *UpdateAbacPolicyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateAbacPolicyResponseBody) SetRequestId(v string) *UpdateAbacPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAbacPolicyResponseBody) SetSuccess(v bool) *UpdateAbacPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAbacPolicyResponseBody) SetUpdatePolicyResult(v int64) *UpdateAbacPolicyResponseBody {
	s.UpdatePolicyResult = &v
	return s
}

func (s *UpdateAbacPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
