// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterSetMemberAuthorizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ModelRouterSetMemberAuthorizationResponseBody
	GetData() *bool
	SetErrCode(v string) *ModelRouterSetMemberAuthorizationResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterSetMemberAuthorizationResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterSetMemberAuthorizationResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterSetMemberAuthorizationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterSetMemberAuthorizationResponseBody
	GetSuccess() *bool
}

type ModelRouterSetMemberAuthorizationResponseBody struct {
	// The data object (contains no business data and is always empty).
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The fault code.
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

func (s ModelRouterSetMemberAuthorizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterSetMemberAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterSetMemberAuthorizationResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModelRouterSetMemberAuthorizationResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterSetMemberAuthorizationResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterSetMemberAuthorizationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterSetMemberAuthorizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterSetMemberAuthorizationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterSetMemberAuthorizationResponseBody) SetData(v bool) *ModelRouterSetMemberAuthorizationResponseBody {
	s.Data = &v
	return s
}

func (s *ModelRouterSetMemberAuthorizationResponseBody) SetErrCode(v string) *ModelRouterSetMemberAuthorizationResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterSetMemberAuthorizationResponseBody) SetErrMessage(v string) *ModelRouterSetMemberAuthorizationResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterSetMemberAuthorizationResponseBody) SetHttpStatusCode(v int32) *ModelRouterSetMemberAuthorizationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterSetMemberAuthorizationResponseBody) SetRequestId(v string) *ModelRouterSetMemberAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterSetMemberAuthorizationResponseBody) SetSuccess(v bool) *ModelRouterSetMemberAuthorizationResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterSetMemberAuthorizationResponseBody) Validate() error {
	return dara.Validate(s)
}
