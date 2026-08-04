// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterResetMemberAuthorizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ModelRouterResetMemberAuthorizationResponseBody
	GetData() *bool
	SetErrCode(v string) *ModelRouterResetMemberAuthorizationResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterResetMemberAuthorizationResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterResetMemberAuthorizationResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterResetMemberAuthorizationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterResetMemberAuthorizationResponseBody
	GetSuccess() *bool
}

type ModelRouterResetMemberAuthorizationResponseBody struct {
	// The data object. This parameter contains no business data and is always empty.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The fault message code.
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

func (s ModelRouterResetMemberAuthorizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterResetMemberAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterResetMemberAuthorizationResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModelRouterResetMemberAuthorizationResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterResetMemberAuthorizationResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterResetMemberAuthorizationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterResetMemberAuthorizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterResetMemberAuthorizationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterResetMemberAuthorizationResponseBody) SetData(v bool) *ModelRouterResetMemberAuthorizationResponseBody {
	s.Data = &v
	return s
}

func (s *ModelRouterResetMemberAuthorizationResponseBody) SetErrCode(v string) *ModelRouterResetMemberAuthorizationResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterResetMemberAuthorizationResponseBody) SetErrMessage(v string) *ModelRouterResetMemberAuthorizationResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterResetMemberAuthorizationResponseBody) SetHttpStatusCode(v int32) *ModelRouterResetMemberAuthorizationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterResetMemberAuthorizationResponseBody) SetRequestId(v string) *ModelRouterResetMemberAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterResetMemberAuthorizationResponseBody) SetSuccess(v bool) *ModelRouterResetMemberAuthorizationResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterResetMemberAuthorizationResponseBody) Validate() error {
	return dara.Validate(s)
}
