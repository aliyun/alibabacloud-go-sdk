// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFunctionDefaultInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateFunctionDefaultInstanceResponseBody
	GetCode() *string
	SetHttpCode(v int64) *UpdateFunctionDefaultInstanceResponseBody
	GetHttpCode() *int64
	SetLatency(v int64) *UpdateFunctionDefaultInstanceResponseBody
	GetLatency() *int64
	SetMessage(v string) *UpdateFunctionDefaultInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateFunctionDefaultInstanceResponseBody
	GetRequestId() *string
	SetStatus(v string) *UpdateFunctionDefaultInstanceResponseBody
	GetStatus() *string
}

type UpdateFunctionDefaultInstanceResponseBody struct {
	// The error code.
	//
	// example:
	//
	// DefaultInstance.SetFail
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
	// A4D487A9-A456-5AA5-A9C6-B7BF2889CF74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the request.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateFunctionDefaultInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFunctionDefaultInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFunctionDefaultInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateFunctionDefaultInstanceResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *UpdateFunctionDefaultInstanceResponseBody) GetLatency() *int64 {
	return s.Latency
}

func (s *UpdateFunctionDefaultInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateFunctionDefaultInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFunctionDefaultInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateFunctionDefaultInstanceResponseBody) SetCode(v string) *UpdateFunctionDefaultInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFunctionDefaultInstanceResponseBody) SetHttpCode(v int64) *UpdateFunctionDefaultInstanceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateFunctionDefaultInstanceResponseBody) SetLatency(v int64) *UpdateFunctionDefaultInstanceResponseBody {
	s.Latency = &v
	return s
}

func (s *UpdateFunctionDefaultInstanceResponseBody) SetMessage(v string) *UpdateFunctionDefaultInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateFunctionDefaultInstanceResponseBody) SetRequestId(v string) *UpdateFunctionDefaultInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFunctionDefaultInstanceResponseBody) SetStatus(v string) *UpdateFunctionDefaultInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateFunctionDefaultInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
