// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarApplyAddResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CarApplyAddResponseBody
	GetCode() *string
	SetMessage(v string) *CarApplyAddResponseBody
	GetMessage() *string
	SetModule(v int64) *CarApplyAddResponseBody
	GetModule() *int64
	SetRequestId(v string) *CarApplyAddResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CarApplyAddResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CarApplyAddResponseBody
	GetTraceId() *string
}

type CarApplyAddResponseBody struct {
	// example:
	//
	// 0
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 1002923002
	Module *int64 `json:"module,omitempty" xml:"module,omitempty"`
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

func (s CarApplyAddResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CarApplyAddResponseBody) GoString() string {
	return s.String()
}

func (s *CarApplyAddResponseBody) GetCode() *string {
	return s.Code
}

func (s *CarApplyAddResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CarApplyAddResponseBody) GetModule() *int64 {
	return s.Module
}

func (s *CarApplyAddResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CarApplyAddResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CarApplyAddResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CarApplyAddResponseBody) SetCode(v string) *CarApplyAddResponseBody {
	s.Code = &v
	return s
}

func (s *CarApplyAddResponseBody) SetMessage(v string) *CarApplyAddResponseBody {
	s.Message = &v
	return s
}

func (s *CarApplyAddResponseBody) SetModule(v int64) *CarApplyAddResponseBody {
	s.Module = &v
	return s
}

func (s *CarApplyAddResponseBody) SetRequestId(v string) *CarApplyAddResponseBody {
	s.RequestId = &v
	return s
}

func (s *CarApplyAddResponseBody) SetSuccess(v bool) *CarApplyAddResponseBody {
	s.Success = &v
	return s
}

func (s *CarApplyAddResponseBody) SetTraceId(v string) *CarApplyAddResponseBody {
	s.TraceId = &v
	return s
}

func (s *CarApplyAddResponseBody) Validate() error {
	return dara.Validate(s)
}
