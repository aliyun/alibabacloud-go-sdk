// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFunctionResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateFunctionResourceResponseBody
	GetCode() *string
	SetHttpCode(v int64) *CreateFunctionResourceResponseBody
	GetHttpCode() *int64
	SetLatency(v float64) *CreateFunctionResourceResponseBody
	GetLatency() *float64
	SetMessage(v string) *CreateFunctionResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateFunctionResourceResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateFunctionResourceResponseBody
	GetStatus() *string
}

type CreateFunctionResourceResponseBody struct {
	// The error code. If no error occurs, this parameter is left empty.
	//
	// example:
	//
	// ""
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
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A4D487A9-A456-5AA5-A9C6-B7BF2889CF74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code. Valid values:
	//
	// 	- OK
	//
	// 	- FAIL
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateFunctionResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFunctionResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFunctionResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateFunctionResourceResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *CreateFunctionResourceResponseBody) GetLatency() *float64 {
	return s.Latency
}

func (s *CreateFunctionResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateFunctionResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFunctionResourceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateFunctionResourceResponseBody) SetCode(v string) *CreateFunctionResourceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFunctionResourceResponseBody) SetHttpCode(v int64) *CreateFunctionResourceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateFunctionResourceResponseBody) SetLatency(v float64) *CreateFunctionResourceResponseBody {
	s.Latency = &v
	return s
}

func (s *CreateFunctionResourceResponseBody) SetMessage(v string) *CreateFunctionResourceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFunctionResourceResponseBody) SetRequestId(v string) *CreateFunctionResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFunctionResourceResponseBody) SetStatus(v string) *CreateFunctionResourceResponseBody {
	s.Status = &v
	return s
}

func (s *CreateFunctionResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
