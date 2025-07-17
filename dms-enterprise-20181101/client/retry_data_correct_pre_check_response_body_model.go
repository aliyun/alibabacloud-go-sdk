// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryDataCorrectPreCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *RetryDataCorrectPreCheckResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *RetryDataCorrectPreCheckResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *RetryDataCorrectPreCheckResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RetryDataCorrectPreCheckResponseBody
	GetSuccess() *bool
}

type RetryDataCorrectPreCheckResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// MissingOrderId
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// OrderId is mandatory for this action.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5BC081C7-5F77-5C92-9758-E1ED17CA****
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

func (s RetryDataCorrectPreCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryDataCorrectPreCheckResponseBody) GoString() string {
	return s.String()
}

func (s *RetryDataCorrectPreCheckResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RetryDataCorrectPreCheckResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RetryDataCorrectPreCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryDataCorrectPreCheckResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RetryDataCorrectPreCheckResponseBody) SetErrorCode(v string) *RetryDataCorrectPreCheckResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RetryDataCorrectPreCheckResponseBody) SetErrorMessage(v string) *RetryDataCorrectPreCheckResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RetryDataCorrectPreCheckResponseBody) SetRequestId(v string) *RetryDataCorrectPreCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryDataCorrectPreCheckResponseBody) SetSuccess(v bool) *RetryDataCorrectPreCheckResponseBody {
	s.Success = &v
	return s
}

func (s *RetryDataCorrectPreCheckResponseBody) Validate() error {
	return dara.Validate(s)
}
