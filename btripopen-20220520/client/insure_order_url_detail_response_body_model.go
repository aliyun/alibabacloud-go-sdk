// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderUrlDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InsureOrderUrlDetailResponseBody
	GetCode() *string
	SetMessage(v string) *InsureOrderUrlDetailResponseBody
	GetMessage() *string
	SetModule(v string) *InsureOrderUrlDetailResponseBody
	GetModule() *string
	SetRequestId(v string) *InsureOrderUrlDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InsureOrderUrlDetailResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *InsureOrderUrlDetailResponseBody
	GetTraceId() *string
}

type InsureOrderUrlDetailResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// https://siopuat.mingya.com.cn/siop-fliggy/index.html#/list?token=fR4iOOZWgG
	Module *string `json:"module,omitempty" xml:"module,omitempty"`
	// example:
	//
	// 210bc59616861088185764700d7589
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 213e212c16975080875814628effac
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s InsureOrderUrlDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderUrlDetailResponseBody) GoString() string {
	return s.String()
}

func (s *InsureOrderUrlDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *InsureOrderUrlDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InsureOrderUrlDetailResponseBody) GetModule() *string {
	return s.Module
}

func (s *InsureOrderUrlDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsureOrderUrlDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InsureOrderUrlDetailResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *InsureOrderUrlDetailResponseBody) SetCode(v string) *InsureOrderUrlDetailResponseBody {
	s.Code = &v
	return s
}

func (s *InsureOrderUrlDetailResponseBody) SetMessage(v string) *InsureOrderUrlDetailResponseBody {
	s.Message = &v
	return s
}

func (s *InsureOrderUrlDetailResponseBody) SetModule(v string) *InsureOrderUrlDetailResponseBody {
	s.Module = &v
	return s
}

func (s *InsureOrderUrlDetailResponseBody) SetRequestId(v string) *InsureOrderUrlDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsureOrderUrlDetailResponseBody) SetSuccess(v bool) *InsureOrderUrlDetailResponseBody {
	s.Success = &v
	return s
}

func (s *InsureOrderUrlDetailResponseBody) SetTraceId(v string) *InsureOrderUrlDetailResponseBody {
	s.TraceId = &v
	return s
}

func (s *InsureOrderUrlDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
