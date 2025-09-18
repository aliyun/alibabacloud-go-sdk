// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccessControlListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateAccessControlListResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *UpdateAccessControlListResponseBody
	GetData() *bool
	SetErrMessage(v string) *UpdateAccessControlListResponseBody
	GetErrMessage() *string
	SetErrorCode(v string) *UpdateAccessControlListResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *UpdateAccessControlListResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateAccessControlListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAccessControlListResponseBody
	GetSuccess() *bool
}

type UpdateAccessControlListResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxxx",     "EncodedDiagnosticMessage": "xxxx",     "AuthPrincipalType": "SubUser",     "AuthPrincipalDisplayName": "xxxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "milvus:xxxx" }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The returned result.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message.
	//
	// example:
	//
	// The format of cidr is illegal.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error code returned.
	//
	// example:
	//
	// ILLEGAL.MilvusCidrFormat
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ABCD-1234-5678-EFGH
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAccessControlListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccessControlListResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAccessControlListResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateAccessControlListResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateAccessControlListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UpdateAccessControlListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateAccessControlListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateAccessControlListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAccessControlListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAccessControlListResponseBody) SetAccessDeniedDetail(v string) *UpdateAccessControlListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateAccessControlListResponseBody) SetData(v bool) *UpdateAccessControlListResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateAccessControlListResponseBody) SetErrMessage(v string) *UpdateAccessControlListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdateAccessControlListResponseBody) SetErrorCode(v string) *UpdateAccessControlListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateAccessControlListResponseBody) SetHttpStatusCode(v int32) *UpdateAccessControlListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateAccessControlListResponseBody) SetRequestId(v string) *UpdateAccessControlListResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAccessControlListResponseBody) SetSuccess(v bool) *UpdateAccessControlListResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAccessControlListResponseBody) Validate() error {
	return dara.Validate(s)
}
