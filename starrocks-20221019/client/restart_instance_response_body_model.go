// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RestartInstanceResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *RestartInstanceResponseBody
	GetData() *bool
	SetErrCode(v string) *RestartInstanceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *RestartInstanceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *RestartInstanceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *RestartInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RestartInstanceResponseBody
	GetSuccess() *bool
}

type RestartInstanceResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The returned data.
	//
	// example:
	//
	// true
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
	// Unsupported.Non.Whitelist.Operation
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

func (s RestartInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RestartInstanceResponseBody) GetData() *bool {
	return s.Data
}

func (s *RestartInstanceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *RestartInstanceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *RestartInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RestartInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RestartInstanceResponseBody) SetAccessDeniedDetail(v string) *RestartInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RestartInstanceResponseBody) SetData(v bool) *RestartInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *RestartInstanceResponseBody) SetErrCode(v string) *RestartInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RestartInstanceResponseBody) SetErrMessage(v string) *RestartInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *RestartInstanceResponseBody) SetHttpStatusCode(v int32) *RestartInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RestartInstanceResponseBody) SetRequestId(v string) *RestartInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartInstanceResponseBody) SetSuccess(v bool) *RestartInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *RestartInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
