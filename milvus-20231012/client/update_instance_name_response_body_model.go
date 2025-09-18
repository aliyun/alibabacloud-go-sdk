// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateInstanceNameResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *UpdateInstanceNameResponseBody
	GetData() *bool
	SetErrCode(v string) *UpdateInstanceNameResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *UpdateInstanceNameResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *UpdateInstanceNameResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateInstanceNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateInstanceNameResponseBody
	GetSuccess() *bool
}

type UpdateInstanceNameResponseBody struct {
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
	// The error code returned.
	//
	// example:
	//
	// Instance.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
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

func (s UpdateInstanceNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNameResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateInstanceNameResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateInstanceNameResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UpdateInstanceNameResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UpdateInstanceNameResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateInstanceNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateInstanceNameResponseBody) SetAccessDeniedDetail(v string) *UpdateInstanceNameResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetData(v bool) *UpdateInstanceNameResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetErrCode(v string) *UpdateInstanceNameResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetErrMessage(v string) *UpdateInstanceNameResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetHttpStatusCode(v int32) *UpdateInstanceNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetRequestId(v string) *UpdateInstanceNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetSuccess(v bool) *UpdateInstanceNameResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) Validate() error {
	return dara.Validate(s)
}
