// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDepartmentSaveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DepartmentSaveResponseBody
	GetCode() *string
	SetMessage(v string) *DepartmentSaveResponseBody
	GetMessage() *string
	SetModule(v string) *DepartmentSaveResponseBody
	GetModule() *string
	SetRequestId(v string) *DepartmentSaveResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DepartmentSaveResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DepartmentSaveResponseBody
	GetTraceId() *string
}

type DepartmentSaveResponseBody struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	Module    *string `json:"module,omitempty" xml:"module,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s DepartmentSaveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DepartmentSaveResponseBody) GoString() string {
	return s.String()
}

func (s *DepartmentSaveResponseBody) GetCode() *string {
	return s.Code
}

func (s *DepartmentSaveResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DepartmentSaveResponseBody) GetModule() *string {
	return s.Module
}

func (s *DepartmentSaveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DepartmentSaveResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DepartmentSaveResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DepartmentSaveResponseBody) SetCode(v string) *DepartmentSaveResponseBody {
	s.Code = &v
	return s
}

func (s *DepartmentSaveResponseBody) SetMessage(v string) *DepartmentSaveResponseBody {
	s.Message = &v
	return s
}

func (s *DepartmentSaveResponseBody) SetModule(v string) *DepartmentSaveResponseBody {
	s.Module = &v
	return s
}

func (s *DepartmentSaveResponseBody) SetRequestId(v string) *DepartmentSaveResponseBody {
	s.RequestId = &v
	return s
}

func (s *DepartmentSaveResponseBody) SetSuccess(v bool) *DepartmentSaveResponseBody {
	s.Success = &v
	return s
}

func (s *DepartmentSaveResponseBody) SetTraceId(v string) *DepartmentSaveResponseBody {
	s.TraceId = &v
	return s
}

func (s *DepartmentSaveResponseBody) Validate() error {
	return dara.Validate(s)
}
