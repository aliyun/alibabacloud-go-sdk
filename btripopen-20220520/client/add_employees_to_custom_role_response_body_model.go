// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEmployeesToCustomRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddEmployeesToCustomRoleResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AddEmployeesToCustomRoleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddEmployeesToCustomRoleResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddEmployeesToCustomRoleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddEmployeesToCustomRoleResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *AddEmployeesToCustomRoleResponseBody
	GetTraceId() *string
}

type AddEmployeesToCustomRoleResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
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
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s AddEmployeesToCustomRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddEmployeesToCustomRoleResponseBody) GoString() string {
	return s.String()
}

func (s *AddEmployeesToCustomRoleResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddEmployeesToCustomRoleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddEmployeesToCustomRoleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddEmployeesToCustomRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddEmployeesToCustomRoleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddEmployeesToCustomRoleResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *AddEmployeesToCustomRoleResponseBody) SetCode(v string) *AddEmployeesToCustomRoleResponseBody {
	s.Code = &v
	return s
}

func (s *AddEmployeesToCustomRoleResponseBody) SetHttpStatusCode(v int32) *AddEmployeesToCustomRoleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddEmployeesToCustomRoleResponseBody) SetMessage(v string) *AddEmployeesToCustomRoleResponseBody {
	s.Message = &v
	return s
}

func (s *AddEmployeesToCustomRoleResponseBody) SetRequestId(v string) *AddEmployeesToCustomRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddEmployeesToCustomRoleResponseBody) SetSuccess(v bool) *AddEmployeesToCustomRoleResponseBody {
	s.Success = &v
	return s
}

func (s *AddEmployeesToCustomRoleResponseBody) SetTraceId(v string) *AddEmployeesToCustomRoleResponseBody {
	s.TraceId = &v
	return s
}

func (s *AddEmployeesToCustomRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
