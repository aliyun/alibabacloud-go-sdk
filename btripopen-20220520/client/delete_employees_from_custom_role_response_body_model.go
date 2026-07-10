// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEmployeesFromCustomRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteEmployeesFromCustomRoleResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteEmployeesFromCustomRoleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteEmployeesFromCustomRoleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteEmployeesFromCustomRoleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteEmployeesFromCustomRoleResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DeleteEmployeesFromCustomRoleResponseBody
	GetTraceId() *string
}

type DeleteEmployeesFromCustomRoleResponseBody struct {
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId        *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s DeleteEmployeesFromCustomRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEmployeesFromCustomRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEmployeesFromCustomRoleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteEmployeesFromCustomRoleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteEmployeesFromCustomRoleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteEmployeesFromCustomRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEmployeesFromCustomRoleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteEmployeesFromCustomRoleResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeleteEmployeesFromCustomRoleResponseBody) SetCode(v string) *DeleteEmployeesFromCustomRoleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEmployeesFromCustomRoleResponseBody) SetHttpStatusCode(v int32) *DeleteEmployeesFromCustomRoleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteEmployeesFromCustomRoleResponseBody) SetMessage(v string) *DeleteEmployeesFromCustomRoleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEmployeesFromCustomRoleResponseBody) SetRequestId(v string) *DeleteEmployeesFromCustomRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEmployeesFromCustomRoleResponseBody) SetSuccess(v bool) *DeleteEmployeesFromCustomRoleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteEmployeesFromCustomRoleResponseBody) SetTraceId(v string) *DeleteEmployeesFromCustomRoleResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeleteEmployeesFromCustomRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
