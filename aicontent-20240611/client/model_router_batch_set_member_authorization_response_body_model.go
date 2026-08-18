// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterBatchSetMemberAuthorizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ModelRouterBatchSetMemberAuthorizationResponseBody
	GetData() *bool
	SetErrCode(v string) *ModelRouterBatchSetMemberAuthorizationResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterBatchSetMemberAuthorizationResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterBatchSetMemberAuthorizationResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterBatchSetMemberAuthorizationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterBatchSetMemberAuthorizationResponseBody
	GetSuccess() *bool
}

type ModelRouterBatchSetMemberAuthorizationResponseBody struct {
	// The data object. This object contains no business data. This operation uses all-or-nothing semantics. If the operation succeeds, all changes take effect.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The fault information code.
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
	// Indicates whether the operation is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterBatchSetMemberAuthorizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterBatchSetMemberAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterBatchSetMemberAuthorizationResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModelRouterBatchSetMemberAuthorizationResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterBatchSetMemberAuthorizationResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterBatchSetMemberAuthorizationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterBatchSetMemberAuthorizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterBatchSetMemberAuthorizationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterBatchSetMemberAuthorizationResponseBody) SetData(v bool) *ModelRouterBatchSetMemberAuthorizationResponseBody {
	s.Data = &v
	return s
}

func (s *ModelRouterBatchSetMemberAuthorizationResponseBody) SetErrCode(v string) *ModelRouterBatchSetMemberAuthorizationResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterBatchSetMemberAuthorizationResponseBody) SetErrMessage(v string) *ModelRouterBatchSetMemberAuthorizationResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterBatchSetMemberAuthorizationResponseBody) SetHttpStatusCode(v int32) *ModelRouterBatchSetMemberAuthorizationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterBatchSetMemberAuthorizationResponseBody) SetRequestId(v string) *ModelRouterBatchSetMemberAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterBatchSetMemberAuthorizationResponseBody) SetSuccess(v bool) *ModelRouterBatchSetMemberAuthorizationResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterBatchSetMemberAuthorizationResponseBody) Validate() error {
	return dara.Validate(s)
}
