// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProjectDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ProjectDeleteResponseBody
	GetCode() *string
	SetMessage(v string) *ProjectDeleteResponseBody
	GetMessage() *string
	SetModule(v bool) *ProjectDeleteResponseBody
	GetModule() *bool
	SetRequestId(v string) *ProjectDeleteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ProjectDeleteResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ProjectDeleteResponseBody
	GetTraceId() *string
}

type ProjectDeleteResponseBody struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	Module    *bool   `json:"module,omitempty" xml:"module,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s ProjectDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ProjectDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *ProjectDeleteResponseBody) GetCode() *string {
	return s.Code
}

func (s *ProjectDeleteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ProjectDeleteResponseBody) GetModule() *bool {
	return s.Module
}

func (s *ProjectDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ProjectDeleteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ProjectDeleteResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ProjectDeleteResponseBody) SetCode(v string) *ProjectDeleteResponseBody {
	s.Code = &v
	return s
}

func (s *ProjectDeleteResponseBody) SetMessage(v string) *ProjectDeleteResponseBody {
	s.Message = &v
	return s
}

func (s *ProjectDeleteResponseBody) SetModule(v bool) *ProjectDeleteResponseBody {
	s.Module = &v
	return s
}

func (s *ProjectDeleteResponseBody) SetRequestId(v string) *ProjectDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProjectDeleteResponseBody) SetSuccess(v bool) *ProjectDeleteResponseBody {
	s.Success = &v
	return s
}

func (s *ProjectDeleteResponseBody) SetTraceId(v string) *ProjectDeleteResponseBody {
	s.TraceId = &v
	return s
}

func (s *ProjectDeleteResponseBody) Validate() error {
	return dara.Validate(s)
}
