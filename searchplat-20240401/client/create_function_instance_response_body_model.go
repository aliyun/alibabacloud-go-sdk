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
	// not found
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Indicates whether the business request is successful. A non-empty value other than 200 indicates a business processing failure.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The latency score.
	//
	// example:
	//
	// 145.411
	Latency *int64 `json:"latency,omitempty" xml:"latency,omitempty"`
	// The error message.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2BA0504F-B179-586D-8210-A7C7C09A9907
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The request status.
	//
	// example:
	//
	// OK
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
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
