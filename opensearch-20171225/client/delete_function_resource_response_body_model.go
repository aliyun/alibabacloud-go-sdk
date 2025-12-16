// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFunctionResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteFunctionResourceResponseBody
	GetCode() *string
	SetHttpCode(v int64) *DeleteFunctionResourceResponseBody
	GetHttpCode() *int64
	SetLatency(v float64) *DeleteFunctionResourceResponseBody
	GetLatency() *float64
	SetMessage(v string) *DeleteFunctionResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteFunctionResourceResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteFunctionResourceResponseBody
	GetStatus() *string
}

type DeleteFunctionResourceResponseBody struct {
	// The error code returned. If no error occurs, this value is empty.
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
	// The error message.
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

func (s DeleteFunctionResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFunctionResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFunctionResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteFunctionResourceResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *DeleteFunctionResourceResponseBody) GetLatency() *float64 {
	return s.Latency
}

func (s *DeleteFunctionResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteFunctionResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFunctionResourceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteFunctionResourceResponseBody) SetCode(v string) *DeleteFunctionResourceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFunctionResourceResponseBody) SetHttpCode(v int64) *DeleteFunctionResourceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteFunctionResourceResponseBody) SetLatency(v float64) *DeleteFunctionResourceResponseBody {
	s.Latency = &v
	return s
}

func (s *DeleteFunctionResourceResponseBody) SetMessage(v string) *DeleteFunctionResourceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFunctionResourceResponseBody) SetRequestId(v string) *DeleteFunctionResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFunctionResourceResponseBody) SetStatus(v string) *DeleteFunctionResourceResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteFunctionResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
