// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ReleaseInstanceResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *ReleaseInstanceResponseBody
	GetData() *bool
	SetErrCode(v string) *ReleaseInstanceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ReleaseInstanceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ReleaseInstanceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ReleaseInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReleaseInstanceResponseBody
	GetSuccess() *bool
}

type ReleaseInstanceResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The returned data.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReleaseInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ReleaseInstanceResponseBody) GetData() *bool {
	return s.Data
}

func (s *ReleaseInstanceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ReleaseInstanceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ReleaseInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ReleaseInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReleaseInstanceResponseBody) SetAccessDeniedDetail(v string) *ReleaseInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetData(v bool) *ReleaseInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetErrCode(v string) *ReleaseInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetErrMessage(v string) *ReleaseInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetHttpStatusCode(v int32) *ReleaseInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetRequestId(v string) *ReleaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetSuccess(v bool) *ReleaseInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ReleaseInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
