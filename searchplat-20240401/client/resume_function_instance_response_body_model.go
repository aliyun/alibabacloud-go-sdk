// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeFunctionInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResumeFunctionInstanceResponseBody
	GetCode() *string
	SetHttpCode(v int64) *ResumeFunctionInstanceResponseBody
	GetHttpCode() *int64
	SetLatency(v int64) *ResumeFunctionInstanceResponseBody
	GetLatency() *int64
	SetMessage(v string) *ResumeFunctionInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResumeFunctionInstanceResponseBody
	GetRequestId() *string
	SetStatus(v string) *ResumeFunctionInstanceResponseBody
	GetStatus() *string
}

type ResumeFunctionInstanceResponseBody struct {
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
	// The time consumed, in milliseconds.
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
	// 5950143C-B8F0-5758-A08A-66F302FD587F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The request status.
	//
	// example:
	//
	// OK
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ResumeFunctionInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeFunctionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeFunctionInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResumeFunctionInstanceResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *ResumeFunctionInstanceResponseBody) GetLatency() *int64 {
	return s.Latency
}

func (s *ResumeFunctionInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResumeFunctionInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeFunctionInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ResumeFunctionInstanceResponseBody) SetCode(v string) *ResumeFunctionInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ResumeFunctionInstanceResponseBody) SetHttpCode(v int64) *ResumeFunctionInstanceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ResumeFunctionInstanceResponseBody) SetLatency(v int64) *ResumeFunctionInstanceResponseBody {
	s.Latency = &v
	return s
}

func (s *ResumeFunctionInstanceResponseBody) SetMessage(v string) *ResumeFunctionInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ResumeFunctionInstanceResponseBody) SetRequestId(v string) *ResumeFunctionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeFunctionInstanceResponseBody) SetStatus(v string) *ResumeFunctionInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *ResumeFunctionInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
