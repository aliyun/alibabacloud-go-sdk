// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCustomRoleResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteCustomRoleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteCustomRoleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCustomRoleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCustomRoleResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DeleteCustomRoleResponseBody
	GetTraceId() *string
}

type DeleteCustomRoleResponseBody struct {
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
	// 210f079e16603757182131635d866a
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s DeleteCustomRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCustomRoleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteCustomRoleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCustomRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomRoleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCustomRoleResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeleteCustomRoleResponseBody) SetCode(v string) *DeleteCustomRoleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCustomRoleResponseBody) SetHttpStatusCode(v int32) *DeleteCustomRoleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteCustomRoleResponseBody) SetMessage(v string) *DeleteCustomRoleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomRoleResponseBody) SetRequestId(v string) *DeleteCustomRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomRoleResponseBody) SetSuccess(v bool) *DeleteCustomRoleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCustomRoleResponseBody) SetTraceId(v string) *DeleteCustomRoleResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeleteCustomRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
