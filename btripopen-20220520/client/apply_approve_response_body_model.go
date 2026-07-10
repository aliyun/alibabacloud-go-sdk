// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyApproveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ApplyApproveResponseBody
	GetCode() *string
	SetMessage(v string) *ApplyApproveResponseBody
	GetMessage() *string
	SetModule(v string) *ApplyApproveResponseBody
	GetModule() *string
	SetRequestId(v string) *ApplyApproveResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ApplyApproveResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ApplyApproveResponseBody
	GetTraceId() *string
}

type ApplyApproveResponseBody struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	Module    *string `json:"module,omitempty" xml:"module,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s ApplyApproveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyApproveResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyApproveResponseBody) GetCode() *string {
	return s.Code
}

func (s *ApplyApproveResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ApplyApproveResponseBody) GetModule() *string {
	return s.Module
}

func (s *ApplyApproveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyApproveResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyApproveResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ApplyApproveResponseBody) SetCode(v string) *ApplyApproveResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyApproveResponseBody) SetMessage(v string) *ApplyApproveResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyApproveResponseBody) SetModule(v string) *ApplyApproveResponseBody {
	s.Module = &v
	return s
}

func (s *ApplyApproveResponseBody) SetRequestId(v string) *ApplyApproveResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyApproveResponseBody) SetSuccess(v bool) *ApplyApproveResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyApproveResponseBody) SetTraceId(v string) *ApplyApproveResponseBody {
	s.TraceId = &v
	return s
}

func (s *ApplyApproveResponseBody) Validate() error {
	return dara.Validate(s)
}
