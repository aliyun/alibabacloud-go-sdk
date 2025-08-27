// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarApplyModifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CarApplyModifyResponseBody
	GetCode() *string
	SetMessage(v string) *CarApplyModifyResponseBody
	GetMessage() *string
	SetModule(v bool) *CarApplyModifyResponseBody
	GetModule() *bool
	SetRequestId(v string) *CarApplyModifyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CarApplyModifyResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CarApplyModifyResponseBody
	GetTraceId() *string
}

type CarApplyModifyResponseBody struct {
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
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
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

func (s CarApplyModifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CarApplyModifyResponseBody) GoString() string {
	return s.String()
}

func (s *CarApplyModifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *CarApplyModifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CarApplyModifyResponseBody) GetModule() *bool {
	return s.Module
}

func (s *CarApplyModifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CarApplyModifyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CarApplyModifyResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CarApplyModifyResponseBody) SetCode(v string) *CarApplyModifyResponseBody {
	s.Code = &v
	return s
}

func (s *CarApplyModifyResponseBody) SetMessage(v string) *CarApplyModifyResponseBody {
	s.Message = &v
	return s
}

func (s *CarApplyModifyResponseBody) SetModule(v bool) *CarApplyModifyResponseBody {
	s.Module = &v
	return s
}

func (s *CarApplyModifyResponseBody) SetRequestId(v string) *CarApplyModifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CarApplyModifyResponseBody) SetSuccess(v bool) *CarApplyModifyResponseBody {
	s.Success = &v
	return s
}

func (s *CarApplyModifyResponseBody) SetTraceId(v string) *CarApplyModifyResponseBody {
	s.TraceId = &v
	return s
}

func (s *CarApplyModifyResponseBody) Validate() error {
	return dara.Validate(s)
}
