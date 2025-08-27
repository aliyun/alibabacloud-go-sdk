// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCustomRoleResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateCustomRoleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateCustomRoleResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCustomRoleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCustomRoleResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CreateCustomRoleResponseBody
	GetTraceId() *string
}

type CreateCustomRoleResponseBody struct {
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
	// A5009956-1077-52FB-B520-EA8C7E91D722
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

func (s CreateCustomRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomRoleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCustomRoleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateCustomRoleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCustomRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomRoleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCustomRoleResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CreateCustomRoleResponseBody) SetCode(v string) *CreateCustomRoleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCustomRoleResponseBody) SetHttpStatusCode(v int32) *CreateCustomRoleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateCustomRoleResponseBody) SetMessage(v string) *CreateCustomRoleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCustomRoleResponseBody) SetRequestId(v string) *CreateCustomRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomRoleResponseBody) SetSuccess(v bool) *CreateCustomRoleResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCustomRoleResponseBody) SetTraceId(v string) *CreateCustomRoleResponseBody {
	s.TraceId = &v
	return s
}

func (s *CreateCustomRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
