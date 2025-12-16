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
	// "Instance.NotExist"
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
	// 10
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	//
	// example:
	//
	// "instance not exist."
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// "3A809095-C554-5CF5-8FCE-BE19D4673790"
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the request. Valid values:
	//
	// 	- OK: The request was successful.
	//
	// 	- FAIL: The request failed.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
