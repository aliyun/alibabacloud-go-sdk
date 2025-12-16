// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFunctionResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateFunctionResourceResponseBody
	GetCode() *string
	SetHttpCode(v int64) *UpdateFunctionResourceResponseBody
	GetHttpCode() *int64
	SetLatency(v float64) *UpdateFunctionResourceResponseBody
	GetLatency() *float64
	SetMessage(v string) *UpdateFunctionResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateFunctionResourceResponseBody
	GetRequestId() *string
	SetStatus(v string) *UpdateFunctionResourceResponseBody
	GetStatus() *string
}

type UpdateFunctionResourceResponseBody struct {
	// The error code. If no error occurs, this parameter is left empty.
	//
	// example:
	//
	// InvalidRequest
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request. Unit: milliseconds.
	//
	// example:
	//
	// 123
	Latency *float64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// Invalid request.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7E375703-5B12-5466-8B48-C4D01AE1291A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the request.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateFunctionResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFunctionResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateFunctionResourceResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *UpdateFunctionResourceResponseBody) GetLatency() *float64 {
	return s.Latency
}

func (s *UpdateFunctionResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateFunctionResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFunctionResourceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateFunctionResourceResponseBody) SetCode(v string) *UpdateFunctionResourceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFunctionResourceResponseBody) SetHttpCode(v int64) *UpdateFunctionResourceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateFunctionResourceResponseBody) SetLatency(v float64) *UpdateFunctionResourceResponseBody {
	s.Latency = &v
	return s
}

func (s *UpdateFunctionResourceResponseBody) SetMessage(v string) *UpdateFunctionResourceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateFunctionResourceResponseBody) SetRequestId(v string) *UpdateFunctionResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFunctionResourceResponseBody) SetStatus(v string) *UpdateFunctionResourceResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateFunctionResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
