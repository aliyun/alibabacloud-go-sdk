// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIsvRuleSaveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IsvRuleSaveResponseBody
	GetCode() *string
	SetMessage(v string) *IsvRuleSaveResponseBody
	GetMessage() *string
	SetModule(v string) *IsvRuleSaveResponseBody
	GetModule() *string
	SetRequestId(v string) *IsvRuleSaveResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IsvRuleSaveResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IsvRuleSaveResponseBody
	GetTraceId() *string
}

type IsvRuleSaveResponseBody struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	Module    *string `json:"module,omitempty" xml:"module,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s IsvRuleSaveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IsvRuleSaveResponseBody) GoString() string {
	return s.String()
}

func (s *IsvRuleSaveResponseBody) GetCode() *string {
	return s.Code
}

func (s *IsvRuleSaveResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IsvRuleSaveResponseBody) GetModule() *string {
	return s.Module
}

func (s *IsvRuleSaveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IsvRuleSaveResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IsvRuleSaveResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IsvRuleSaveResponseBody) SetCode(v string) *IsvRuleSaveResponseBody {
	s.Code = &v
	return s
}

func (s *IsvRuleSaveResponseBody) SetMessage(v string) *IsvRuleSaveResponseBody {
	s.Message = &v
	return s
}

func (s *IsvRuleSaveResponseBody) SetModule(v string) *IsvRuleSaveResponseBody {
	s.Module = &v
	return s
}

func (s *IsvRuleSaveResponseBody) SetRequestId(v string) *IsvRuleSaveResponseBody {
	s.RequestId = &v
	return s
}

func (s *IsvRuleSaveResponseBody) SetSuccess(v bool) *IsvRuleSaveResponseBody {
	s.Success = &v
	return s
}

func (s *IsvRuleSaveResponseBody) SetTraceId(v string) *IsvRuleSaveResponseBody {
	s.TraceId = &v
	return s
}

func (s *IsvRuleSaveResponseBody) Validate() error {
	return dara.Validate(s)
}
