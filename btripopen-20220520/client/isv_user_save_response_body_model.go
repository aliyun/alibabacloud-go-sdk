// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIsvUserSaveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IsvUserSaveResponseBody
	GetCode() *string
	SetMessage(v string) *IsvUserSaveResponseBody
	GetMessage() *string
	SetModule(v string) *IsvUserSaveResponseBody
	GetModule() *string
	SetRequestId(v string) *IsvUserSaveResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IsvUserSaveResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IsvUserSaveResponseBody
	GetTraceId() *string
}

type IsvUserSaveResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	Module  *string `json:"module,omitempty" xml:"module,omitempty"`
	// example:
	//
	// A5009956-1077-52FB-B520-EA8C7E91D722
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 707c9fd116393792883244141e4e7d
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s IsvUserSaveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IsvUserSaveResponseBody) GoString() string {
	return s.String()
}

func (s *IsvUserSaveResponseBody) GetCode() *string {
	return s.Code
}

func (s *IsvUserSaveResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IsvUserSaveResponseBody) GetModule() *string {
	return s.Module
}

func (s *IsvUserSaveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IsvUserSaveResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IsvUserSaveResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IsvUserSaveResponseBody) SetCode(v string) *IsvUserSaveResponseBody {
	s.Code = &v
	return s
}

func (s *IsvUserSaveResponseBody) SetMessage(v string) *IsvUserSaveResponseBody {
	s.Message = &v
	return s
}

func (s *IsvUserSaveResponseBody) SetModule(v string) *IsvUserSaveResponseBody {
	s.Module = &v
	return s
}

func (s *IsvUserSaveResponseBody) SetRequestId(v string) *IsvUserSaveResponseBody {
	s.RequestId = &v
	return s
}

func (s *IsvUserSaveResponseBody) SetSuccess(v bool) *IsvUserSaveResponseBody {
	s.Success = &v
	return s
}

func (s *IsvUserSaveResponseBody) SetTraceId(v string) *IsvUserSaveResponseBody {
	s.TraceId = &v
	return s
}

func (s *IsvUserSaveResponseBody) Validate() error {
	return dara.Validate(s)
}
