// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAbacPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeletePolicyResult(v bool) *DeleteAbacPolicyResponseBody
	GetDeletePolicyResult() *bool
	SetErrorCode(v string) *DeleteAbacPolicyResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteAbacPolicyResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteAbacPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAbacPolicyResponseBody
	GetSuccess() *bool
}

type DeleteAbacPolicyResponseBody struct {
	// Indicates whether the policy is deleted.
	//
	// example:
	//
	// true
	DeletePolicyResult *bool `json:"DeletePolicyResult,omitempty" xml:"DeletePolicyResult,omitempty"`
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

func (s DeleteAbacPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAbacPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAbacPolicyResponseBody) GetDeletePolicyResult() *bool {
	return s.DeletePolicyResult
}

func (s *DeleteAbacPolicyResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteAbacPolicyResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteAbacPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAbacPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAbacPolicyResponseBody) SetDeletePolicyResult(v bool) *DeleteAbacPolicyResponseBody {
	s.DeletePolicyResult = &v
	return s
}

func (s *DeleteAbacPolicyResponseBody) SetErrorCode(v string) *DeleteAbacPolicyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteAbacPolicyResponseBody) SetErrorMessage(v string) *DeleteAbacPolicyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteAbacPolicyResponseBody) SetRequestId(v string) *DeleteAbacPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAbacPolicyResponseBody) SetSuccess(v bool) *DeleteAbacPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAbacPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
