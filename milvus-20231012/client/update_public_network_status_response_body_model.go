// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublicNetworkStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdatePublicNetworkStatusResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *UpdatePublicNetworkStatusResponseBody
	GetData() *bool
	SetErrCode(v string) *UpdatePublicNetworkStatusResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *UpdatePublicNetworkStatusResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *UpdatePublicNetworkStatusResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdatePublicNetworkStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdatePublicNetworkStatusResponseBody
	GetSuccess() *bool
}

type UpdatePublicNetworkStatusResponseBody struct {
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
	// Failed to find instance
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

func (s UpdatePublicNetworkStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublicNetworkStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePublicNetworkStatusResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdatePublicNetworkStatusResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdatePublicNetworkStatusResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UpdatePublicNetworkStatusResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UpdatePublicNetworkStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdatePublicNetworkStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePublicNetworkStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdatePublicNetworkStatusResponseBody) SetAccessDeniedDetail(v string) *UpdatePublicNetworkStatusResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdatePublicNetworkStatusResponseBody) SetData(v bool) *UpdatePublicNetworkStatusResponseBody {
	s.Data = &v
	return s
}

func (s *UpdatePublicNetworkStatusResponseBody) SetErrCode(v string) *UpdatePublicNetworkStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdatePublicNetworkStatusResponseBody) SetErrMessage(v string) *UpdatePublicNetworkStatusResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdatePublicNetworkStatusResponseBody) SetHttpStatusCode(v int32) *UpdatePublicNetworkStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdatePublicNetworkStatusResponseBody) SetRequestId(v string) *UpdatePublicNetworkStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePublicNetworkStatusResponseBody) SetSuccess(v bool) *UpdatePublicNetworkStatusResponseBody {
	s.Success = &v
	return s
}

func (s *UpdatePublicNetworkStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
