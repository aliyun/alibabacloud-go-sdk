// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUpgradableVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryUpgradableVersionsResponseBody
	GetAccessDeniedDetail() *string
	SetData(v []*string) *QueryUpgradableVersionsResponseBody
	GetData() []*string
	SetErrCode(v string) *QueryUpgradableVersionsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *QueryUpgradableVersionsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *QueryUpgradableVersionsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryUpgradableVersionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryUpgradableVersionsResponseBody
	GetSuccess() *bool
}

type QueryUpgradableVersionsResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The versions that you can upgrade to.
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s QueryUpgradableVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUpgradableVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUpgradableVersionsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryUpgradableVersionsResponseBody) GetData() []*string {
	return s.Data
}

func (s *QueryUpgradableVersionsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *QueryUpgradableVersionsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *QueryUpgradableVersionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryUpgradableVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUpgradableVersionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryUpgradableVersionsResponseBody) SetAccessDeniedDetail(v string) *QueryUpgradableVersionsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryUpgradableVersionsResponseBody) SetData(v []*string) *QueryUpgradableVersionsResponseBody {
	s.Data = v
	return s
}

func (s *QueryUpgradableVersionsResponseBody) SetErrCode(v string) *QueryUpgradableVersionsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryUpgradableVersionsResponseBody) SetErrMessage(v string) *QueryUpgradableVersionsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryUpgradableVersionsResponseBody) SetHttpStatusCode(v int32) *QueryUpgradableVersionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryUpgradableVersionsResponseBody) SetRequestId(v string) *QueryUpgradableVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUpgradableVersionsResponseBody) SetSuccess(v bool) *QueryUpgradableVersionsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUpgradableVersionsResponseBody) Validate() error {
	return dara.Validate(s)
}
