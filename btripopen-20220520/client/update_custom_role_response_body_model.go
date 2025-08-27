// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateCustomRoleResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateCustomRoleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateCustomRoleResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateCustomRoleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCustomRoleResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UpdateCustomRoleResponseBody
	GetTraceId() *string
}

type UpdateCustomRoleResponseBody struct {
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
	// 210bcc3a16583004579056128d33d7
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s UpdateCustomRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomRoleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoleResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateCustomRoleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateCustomRoleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateCustomRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomRoleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCustomRoleResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UpdateCustomRoleResponseBody) SetCode(v string) *UpdateCustomRoleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCustomRoleResponseBody) SetHttpStatusCode(v int32) *UpdateCustomRoleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateCustomRoleResponseBody) SetMessage(v string) *UpdateCustomRoleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCustomRoleResponseBody) SetRequestId(v string) *UpdateCustomRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomRoleResponseBody) SetSuccess(v bool) *UpdateCustomRoleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCustomRoleResponseBody) SetTraceId(v string) *UpdateCustomRoleResponseBody {
	s.TraceId = &v
	return s
}

func (s *UpdateCustomRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
