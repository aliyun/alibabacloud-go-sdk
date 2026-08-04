// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterSetUserRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ModelRouterSetUserRolesResponseBody
	GetData() *bool
	SetErrCode(v string) *ModelRouterSetUserRolesResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterSetUserRolesResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterSetUserRolesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterSetUserRolesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterSetUserRolesResponseBody
	GetSuccess() *bool
}

type ModelRouterSetUserRolesResponseBody struct {
	// The data object (no business data, always empty).
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The error message code.
	//
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unknown error
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterSetUserRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterSetUserRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterSetUserRolesResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModelRouterSetUserRolesResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterSetUserRolesResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterSetUserRolesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterSetUserRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterSetUserRolesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterSetUserRolesResponseBody) SetData(v bool) *ModelRouterSetUserRolesResponseBody {
	s.Data = &v
	return s
}

func (s *ModelRouterSetUserRolesResponseBody) SetErrCode(v string) *ModelRouterSetUserRolesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterSetUserRolesResponseBody) SetErrMessage(v string) *ModelRouterSetUserRolesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterSetUserRolesResponseBody) SetHttpStatusCode(v int32) *ModelRouterSetUserRolesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterSetUserRolesResponseBody) SetRequestId(v string) *ModelRouterSetUserRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterSetUserRolesResponseBody) SetSuccess(v bool) *ModelRouterSetUserRolesResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterSetUserRolesResponseBody) Validate() error {
	return dara.Validate(s)
}
