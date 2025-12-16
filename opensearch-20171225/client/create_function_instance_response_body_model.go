// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFunctionInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateFunctionInstanceResponseBody
	GetCode() *string
	SetHttpCode(v int64) *CreateFunctionInstanceResponseBody
	GetHttpCode() *int64
	SetLatency(v int64) *CreateFunctionInstanceResponseBody
	GetLatency() *int64
	SetMessage(v string) *CreateFunctionInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateFunctionInstanceResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateFunctionInstanceResponseBody
	GetStatus() *string
}

type CreateFunctionInstanceResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Version.NotExist
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
	// The error message. If no error occurs, this parameter is left empty.
	//
	// example:
	//
	// version not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 98724351-D6B2-5D8A-B089-7FFD1821A7E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the request.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateFunctionInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFunctionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFunctionInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateFunctionInstanceResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *CreateFunctionInstanceResponseBody) GetLatency() *int64 {
	return s.Latency
}

func (s *CreateFunctionInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateFunctionInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFunctionInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateFunctionInstanceResponseBody) SetCode(v string) *CreateFunctionInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFunctionInstanceResponseBody) SetHttpCode(v int64) *CreateFunctionInstanceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateFunctionInstanceResponseBody) SetLatency(v int64) *CreateFunctionInstanceResponseBody {
	s.Latency = &v
	return s
}

func (s *CreateFunctionInstanceResponseBody) SetMessage(v string) *CreateFunctionInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFunctionInstanceResponseBody) SetRequestId(v string) *CreateFunctionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFunctionInstanceResponseBody) SetStatus(v string) *CreateFunctionInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *CreateFunctionInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
