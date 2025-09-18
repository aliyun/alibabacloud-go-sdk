// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDefaultRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateDefaultRoleResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *CreateDefaultRoleResponseBody
	GetData() *bool
	SetErrCode(v string) *CreateDefaultRoleResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateDefaultRoleResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *CreateDefaultRoleResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateDefaultRoleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDefaultRoleResponseBody
	GetSuccess() *bool
}

type CreateDefaultRoleResponseBody struct {
	// The details about the failed permission verification.
	//
	// example:
	//
	// { "PolicyType": "AccountLevelIdentityBasedPolicy", "AuthPrincipalOwnerId": "xxxx", "EncodedDiagnosticMessage": "xxxx", "AuthPrincipalType": "SubUser", "AuthPrincipalDisplayName": "xxxx", "NoPermissionType": "ImplicitDeny", "AuthAction": "milvus:xxxx" }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The returned result.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code
	//
	// example:
	//
	// Instance.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// Failed to find instance c-123xxx
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
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

func (s CreateDefaultRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDefaultRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDefaultRoleResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateDefaultRoleResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateDefaultRoleResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateDefaultRoleResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateDefaultRoleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDefaultRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDefaultRoleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDefaultRoleResponseBody) SetAccessDeniedDetail(v string) *CreateDefaultRoleResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateDefaultRoleResponseBody) SetData(v bool) *CreateDefaultRoleResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDefaultRoleResponseBody) SetErrCode(v string) *CreateDefaultRoleResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateDefaultRoleResponseBody) SetErrMessage(v string) *CreateDefaultRoleResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateDefaultRoleResponseBody) SetHttpStatusCode(v int32) *CreateDefaultRoleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDefaultRoleResponseBody) SetRequestId(v string) *CreateDefaultRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDefaultRoleResponseBody) SetSuccess(v bool) *CreateDefaultRoleResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDefaultRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
