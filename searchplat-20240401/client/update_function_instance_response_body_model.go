// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFunctionInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateFunctionInstanceResponseBody
	GetCode() *string
	SetHttpCode(v int64) *UpdateFunctionInstanceResponseBody
	GetHttpCode() *int64
	SetLatency(v int64) *UpdateFunctionInstanceResponseBody
	GetLatency() *int64
	SetMessage(v string) *UpdateFunctionInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateFunctionInstanceResponseBody
	GetRequestId() *string
	SetStatus(v string) *UpdateFunctionInstanceResponseBody
	GetStatus() *string
}

type UpdateFunctionInstanceResponseBody struct {
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
	// The latency.
	//
	// example:
	//
	// 34.946
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

func (s UpdateFunctionInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFunctionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFunctionInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateFunctionInstanceResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *UpdateFunctionInstanceResponseBody) GetLatency() *int64 {
	return s.Latency
}

func (s *UpdateFunctionInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateFunctionInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFunctionInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateFunctionInstanceResponseBody) SetCode(v string) *UpdateFunctionInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFunctionInstanceResponseBody) SetHttpCode(v int64) *UpdateFunctionInstanceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateFunctionInstanceResponseBody) SetLatency(v int64) *UpdateFunctionInstanceResponseBody {
	s.Latency = &v
	return s
}

func (s *UpdateFunctionInstanceResponseBody) SetMessage(v string) *UpdateFunctionInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateFunctionInstanceResponseBody) SetRequestId(v string) *UpdateFunctionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFunctionInstanceResponseBody) SetStatus(v string) *UpdateFunctionInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateFunctionInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
