// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryJobInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *RetryJobInstanceResponseBody
	GetCode() *int32
	SetMessage(v string) *RetryJobInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *RetryJobInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RetryJobInstanceResponseBody
	GetSuccess() *bool
}

type RetryJobInstanceResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned error message.
	//
	// example:
	//
	// jobId=xxx is not existed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
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

func (s RetryJobInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryJobInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RetryJobInstanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *RetryJobInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RetryJobInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryJobInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RetryJobInstanceResponseBody) SetCode(v int32) *RetryJobInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *RetryJobInstanceResponseBody) SetMessage(v string) *RetryJobInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *RetryJobInstanceResponseBody) SetRequestId(v string) *RetryJobInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryJobInstanceResponseBody) SetSuccess(v bool) *RetryJobInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *RetryJobInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
