// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateApplicationDescriptionResponseBody
	GetCode() *string
	SetErrorCode(v string) *UpdateApplicationDescriptionResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpdateApplicationDescriptionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateApplicationDescriptionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateApplicationDescriptionResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UpdateApplicationDescriptionResponseBody
	GetTraceId() *string
}

type UpdateApplicationDescriptionResponseBody struct {
	// example:
	//
	// 200
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 01CF26C7-00A3-4AA6-BA76-7E95F2A3***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// ac1a0b2215622246421415014e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UpdateApplicationDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationDescriptionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateApplicationDescriptionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateApplicationDescriptionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateApplicationDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationDescriptionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateApplicationDescriptionResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UpdateApplicationDescriptionResponseBody) SetCode(v string) *UpdateApplicationDescriptionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateApplicationDescriptionResponseBody) SetErrorCode(v string) *UpdateApplicationDescriptionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateApplicationDescriptionResponseBody) SetMessage(v string) *UpdateApplicationDescriptionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateApplicationDescriptionResponseBody) SetRequestId(v string) *UpdateApplicationDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationDescriptionResponseBody) SetSuccess(v bool) *UpdateApplicationDescriptionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateApplicationDescriptionResponseBody) SetTraceId(v string) *UpdateApplicationDescriptionResponseBody {
	s.TraceId = &v
	return s
}

func (s *UpdateApplicationDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
