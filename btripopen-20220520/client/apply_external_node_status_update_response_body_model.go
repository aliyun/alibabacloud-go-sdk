// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyExternalNodeStatusUpdateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ApplyExternalNodeStatusUpdateResponseBody
	GetCode() *string
	SetMessage(v string) *ApplyExternalNodeStatusUpdateResponseBody
	GetMessage() *string
	SetModule(v bool) *ApplyExternalNodeStatusUpdateResponseBody
	GetModule() *bool
	SetRequestId(v string) *ApplyExternalNodeStatusUpdateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ApplyExternalNodeStatusUpdateResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ApplyExternalNodeStatusUpdateResponseBody
	GetTraceId() *string
}

type ApplyExternalNodeStatusUpdateResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// true
	Module *bool `json:"module,omitempty" xml:"module,omitempty"`
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 210f079416784321627628333de4ab
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s ApplyExternalNodeStatusUpdateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyExternalNodeStatusUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyExternalNodeStatusUpdateResponseBody) GetCode() *string {
	return s.Code
}

func (s *ApplyExternalNodeStatusUpdateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ApplyExternalNodeStatusUpdateResponseBody) GetModule() *bool {
	return s.Module
}

func (s *ApplyExternalNodeStatusUpdateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyExternalNodeStatusUpdateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyExternalNodeStatusUpdateResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ApplyExternalNodeStatusUpdateResponseBody) SetCode(v string) *ApplyExternalNodeStatusUpdateResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyExternalNodeStatusUpdateResponseBody) SetMessage(v string) *ApplyExternalNodeStatusUpdateResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyExternalNodeStatusUpdateResponseBody) SetModule(v bool) *ApplyExternalNodeStatusUpdateResponseBody {
	s.Module = &v
	return s
}

func (s *ApplyExternalNodeStatusUpdateResponseBody) SetRequestId(v string) *ApplyExternalNodeStatusUpdateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyExternalNodeStatusUpdateResponseBody) SetSuccess(v bool) *ApplyExternalNodeStatusUpdateResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyExternalNodeStatusUpdateResponseBody) SetTraceId(v string) *ApplyExternalNodeStatusUpdateResponseBody {
	s.TraceId = &v
	return s
}

func (s *ApplyExternalNodeStatusUpdateResponseBody) Validate() error {
	return dara.Validate(s)
}
