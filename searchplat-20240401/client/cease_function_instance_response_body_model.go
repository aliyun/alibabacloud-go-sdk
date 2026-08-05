// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCeaseFunctionInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CeaseFunctionInstanceResponseBody
	GetCode() *string
	SetHttpCode(v int64) *CeaseFunctionInstanceResponseBody
	GetHttpCode() *int64
	SetLatency(v int64) *CeaseFunctionInstanceResponseBody
	GetLatency() *int64
	SetMessage(v string) *CeaseFunctionInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CeaseFunctionInstanceResponseBody
	GetRequestId() *string
	SetStatus(v string) *CeaseFunctionInstanceResponseBody
	GetStatus() *string
}

type CeaseFunctionInstanceResponseBody struct {
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
	// 2423C841-91C4-5E51-B296-590D367967FC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The request status.
	//
	// example:
	//
	// OK
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CeaseFunctionInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CeaseFunctionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CeaseFunctionInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CeaseFunctionInstanceResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *CeaseFunctionInstanceResponseBody) GetLatency() *int64 {
	return s.Latency
}

func (s *CeaseFunctionInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CeaseFunctionInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CeaseFunctionInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CeaseFunctionInstanceResponseBody) SetCode(v string) *CeaseFunctionInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CeaseFunctionInstanceResponseBody) SetHttpCode(v int64) *CeaseFunctionInstanceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CeaseFunctionInstanceResponseBody) SetLatency(v int64) *CeaseFunctionInstanceResponseBody {
	s.Latency = &v
	return s
}

func (s *CeaseFunctionInstanceResponseBody) SetMessage(v string) *CeaseFunctionInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CeaseFunctionInstanceResponseBody) SetRequestId(v string) *CeaseFunctionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CeaseFunctionInstanceResponseBody) SetStatus(v string) *CeaseFunctionInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *CeaseFunctionInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
