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
	// not found
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The latency in minutes.
	//
	// example:
	//
	// 145.411
	Latency *int64 `json:"latency,omitempty" xml:"latency,omitempty"`
	// The error message.
	//
	// example:
	//
	// "xx not found"
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 58113A95-1858-5674-87E5-192AEE6FD9DD
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The request status.
	//
	// example:
	//
	// OK
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
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
