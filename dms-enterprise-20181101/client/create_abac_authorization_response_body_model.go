// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAbacAuthorizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateAbacAuthorizationResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateAbacAuthorizationResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateAbacAuthorizationResponseBody
	GetRequestId() *string
	SetResult(v string) *CreateAbacAuthorizationResponseBody
	GetResult() *string
	SetSuccess(v bool) *CreateAbacAuthorizationResponseBody
	GetSuccess() *bool
}

type CreateAbacAuthorizationResponseBody struct {
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
	// The request ID. You can use the request ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the policy is attached.
	//
	// example:
	//
	// true
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request succeeded.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAbacAuthorizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAbacAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAbacAuthorizationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateAbacAuthorizationResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateAbacAuthorizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAbacAuthorizationResponseBody) GetResult() *string {
	return s.Result
}

func (s *CreateAbacAuthorizationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAbacAuthorizationResponseBody) SetErrorCode(v string) *CreateAbacAuthorizationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateAbacAuthorizationResponseBody) SetErrorMessage(v string) *CreateAbacAuthorizationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateAbacAuthorizationResponseBody) SetRequestId(v string) *CreateAbacAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAbacAuthorizationResponseBody) SetResult(v string) *CreateAbacAuthorizationResponseBody {
	s.Result = &v
	return s
}

func (s *CreateAbacAuthorizationResponseBody) SetSuccess(v bool) *CreateAbacAuthorizationResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAbacAuthorizationResponseBody) Validate() error {
	return dara.Validate(s)
}
