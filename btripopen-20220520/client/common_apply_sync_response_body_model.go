// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommonApplySyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CommonApplySyncResponseBody
	GetCode() *string
	SetMessage(v string) *CommonApplySyncResponseBody
	GetMessage() *string
	SetModule(v bool) *CommonApplySyncResponseBody
	GetModule() *bool
	SetRequestId(v string) *CommonApplySyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CommonApplySyncResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CommonApplySyncResponseBody
	GetTraceId() *string
}

type CommonApplySyncResponseBody struct {
	// example:
	//
	// 0
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module。
	//
	// example:
	//
	// {\\"list\\": [], \\"pageSize\\": 20, \\"pageNo\\": 1}
	Module *bool `json:"module,omitempty" xml:"module,omitempty"`
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CommonApplySyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CommonApplySyncResponseBody) GoString() string {
	return s.String()
}

func (s *CommonApplySyncResponseBody) GetCode() *string {
	return s.Code
}

func (s *CommonApplySyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CommonApplySyncResponseBody) GetModule() *bool {
	return s.Module
}

func (s *CommonApplySyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CommonApplySyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CommonApplySyncResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CommonApplySyncResponseBody) SetCode(v string) *CommonApplySyncResponseBody {
	s.Code = &v
	return s
}

func (s *CommonApplySyncResponseBody) SetMessage(v string) *CommonApplySyncResponseBody {
	s.Message = &v
	return s
}

func (s *CommonApplySyncResponseBody) SetModule(v bool) *CommonApplySyncResponseBody {
	s.Module = &v
	return s
}

func (s *CommonApplySyncResponseBody) SetRequestId(v string) *CommonApplySyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *CommonApplySyncResponseBody) SetSuccess(v bool) *CommonApplySyncResponseBody {
	s.Success = &v
	return s
}

func (s *CommonApplySyncResponseBody) SetTraceId(v string) *CommonApplySyncResponseBody {
	s.TraceId = &v
	return s
}

func (s *CommonApplySyncResponseBody) Validate() error {
	return dara.Validate(s)
}
