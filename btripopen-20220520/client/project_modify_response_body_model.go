// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProjectModifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ProjectModifyResponseBody
	GetCode() *string
	SetMessage(v string) *ProjectModifyResponseBody
	GetMessage() *string
	SetModule(v bool) *ProjectModifyResponseBody
	GetModule() *bool
	SetRequestId(v string) *ProjectModifyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ProjectModifyResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ProjectModifyResponseBody
	GetTraceId() *string
}

type ProjectModifyResponseBody struct {
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

func (s ProjectModifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ProjectModifyResponseBody) GoString() string {
	return s.String()
}

func (s *ProjectModifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *ProjectModifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ProjectModifyResponseBody) GetModule() *bool {
	return s.Module
}

func (s *ProjectModifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ProjectModifyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ProjectModifyResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ProjectModifyResponseBody) SetCode(v string) *ProjectModifyResponseBody {
	s.Code = &v
	return s
}

func (s *ProjectModifyResponseBody) SetMessage(v string) *ProjectModifyResponseBody {
	s.Message = &v
	return s
}

func (s *ProjectModifyResponseBody) SetModule(v bool) *ProjectModifyResponseBody {
	s.Module = &v
	return s
}

func (s *ProjectModifyResponseBody) SetRequestId(v string) *ProjectModifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProjectModifyResponseBody) SetSuccess(v bool) *ProjectModifyResponseBody {
	s.Success = &v
	return s
}

func (s *ProjectModifyResponseBody) SetTraceId(v string) *ProjectModifyResponseBody {
	s.TraceId = &v
	return s
}

func (s *ProjectModifyResponseBody) Validate() error {
	return dara.Validate(s)
}
