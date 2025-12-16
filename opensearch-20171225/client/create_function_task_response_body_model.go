// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFunctionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateFunctionTaskResponseBody
	GetCode() *string
	SetHttpCode(v int64) *CreateFunctionTaskResponseBody
	GetHttpCode() *int64
	SetLatency(v int64) *CreateFunctionTaskResponseBody
	GetLatency() *int64
	SetMessage(v string) *CreateFunctionTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateFunctionTaskResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateFunctionTaskResponseBody
	GetStatus() *string
}

type CreateFunctionTaskResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Task.IsRunning
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	//
	// example:
	//
	// 123
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1638157990724
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the request.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateFunctionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFunctionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFunctionTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateFunctionTaskResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *CreateFunctionTaskResponseBody) GetLatency() *int64 {
	return s.Latency
}

func (s *CreateFunctionTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateFunctionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFunctionTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateFunctionTaskResponseBody) SetCode(v string) *CreateFunctionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFunctionTaskResponseBody) SetHttpCode(v int64) *CreateFunctionTaskResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateFunctionTaskResponseBody) SetLatency(v int64) *CreateFunctionTaskResponseBody {
	s.Latency = &v
	return s
}

func (s *CreateFunctionTaskResponseBody) SetMessage(v string) *CreateFunctionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFunctionTaskResponseBody) SetRequestId(v string) *CreateFunctionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFunctionTaskResponseBody) SetStatus(v string) *CreateFunctionTaskResponseBody {
	s.Status = &v
	return s
}

func (s *CreateFunctionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
