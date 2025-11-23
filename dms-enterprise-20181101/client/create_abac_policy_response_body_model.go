// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAbacPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatePolicyResult(v int64) *CreateAbacPolicyResponseBody
	GetCreatePolicyResult() *int64
	SetErrorCode(v string) *CreateAbacPolicyResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateAbacPolicyResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateAbacPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAbacPolicyResponseBody
	GetSuccess() *bool
}

type CreateAbacPolicyResponseBody struct {
	// The ID of the policy.
	//
	// example:
	//
	// 12****
	CreatePolicyResult *int64 `json:"CreatePolicyResult,omitempty" xml:"CreatePolicyResult,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
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

func (s CreateAbacPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAbacPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAbacPolicyResponseBody) GetCreatePolicyResult() *int64 {
	return s.CreatePolicyResult
}

func (s *CreateAbacPolicyResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateAbacPolicyResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateAbacPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAbacPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAbacPolicyResponseBody) SetCreatePolicyResult(v int64) *CreateAbacPolicyResponseBody {
	s.CreatePolicyResult = &v
	return s
}

func (s *CreateAbacPolicyResponseBody) SetErrorCode(v string) *CreateAbacPolicyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateAbacPolicyResponseBody) SetErrorMessage(v string) *CreateAbacPolicyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateAbacPolicyResponseBody) SetRequestId(v string) *CreateAbacPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAbacPolicyResponseBody) SetSuccess(v bool) *CreateAbacPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAbacPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
