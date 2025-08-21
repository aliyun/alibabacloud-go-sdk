// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInterruptLogstashTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InterruptLogstashTaskResponseBody
	GetCode() *string
	SetMessage(v string) *InterruptLogstashTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *InterruptLogstashTaskResponseBody
	GetRequestId() *string
	SetResult(v bool) *InterruptLogstashTaskResponseBody
	GetResult() *bool
}

type InterruptLogstashTaskResponseBody struct {
	// The error code returned. If the API operation is successfully called, this parameter is not returned.
	//
	// example:
	//
	// .
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned. If the API operation is successfully called, this parameter is not returned.
	//
	// example:
	//
	// .
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0FA05123-745C-42FD-A69B-AFF48EF9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the task is suspended. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s InterruptLogstashTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InterruptLogstashTaskResponseBody) GoString() string {
	return s.String()
}

func (s *InterruptLogstashTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *InterruptLogstashTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InterruptLogstashTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InterruptLogstashTaskResponseBody) GetResult() *bool {
	return s.Result
}

func (s *InterruptLogstashTaskResponseBody) SetCode(v string) *InterruptLogstashTaskResponseBody {
	s.Code = &v
	return s
}

func (s *InterruptLogstashTaskResponseBody) SetMessage(v string) *InterruptLogstashTaskResponseBody {
	s.Message = &v
	return s
}

func (s *InterruptLogstashTaskResponseBody) SetRequestId(v string) *InterruptLogstashTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *InterruptLogstashTaskResponseBody) SetResult(v bool) *InterruptLogstashTaskResponseBody {
	s.Result = &v
	return s
}

func (s *InterruptLogstashTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
