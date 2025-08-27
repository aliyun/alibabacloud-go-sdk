// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDepartmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDepartmentResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteDepartmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDepartmentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDepartmentResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DeleteDepartmentResponseBody
	GetTraceId() *string
}

type DeleteDepartmentResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 210e847f16611516748613869de4f6
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s DeleteDepartmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDepartmentResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDepartmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDepartmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDepartmentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDepartmentResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeleteDepartmentResponseBody) SetCode(v string) *DeleteDepartmentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDepartmentResponseBody) SetMessage(v string) *DeleteDepartmentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDepartmentResponseBody) SetRequestId(v string) *DeleteDepartmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDepartmentResponseBody) SetSuccess(v bool) *DeleteDepartmentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDepartmentResponseBody) SetTraceId(v string) *DeleteDepartmentResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeleteDepartmentResponseBody) Validate() error {
	return dara.Validate(s)
}
