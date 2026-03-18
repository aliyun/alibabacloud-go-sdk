// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMinorVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryMinorVersionResponseBody
	GetAccessDeniedDetail() *string
	SetData(v string) *QueryMinorVersionResponseBody
	GetData() *string
	SetErrCode(v string) *QueryMinorVersionResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *QueryMinorVersionResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *QueryMinorVersionResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryMinorVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryMinorVersionResponseBody
	GetSuccess() *bool
}

type QueryMinorVersionResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 3.3.13-1.0-1.7.12
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE74XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMinorVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMinorVersionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMinorVersionResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryMinorVersionResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryMinorVersionResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *QueryMinorVersionResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *QueryMinorVersionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryMinorVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMinorVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMinorVersionResponseBody) SetAccessDeniedDetail(v string) *QueryMinorVersionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryMinorVersionResponseBody) SetData(v string) *QueryMinorVersionResponseBody {
	s.Data = &v
	return s
}

func (s *QueryMinorVersionResponseBody) SetErrCode(v string) *QueryMinorVersionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryMinorVersionResponseBody) SetErrMessage(v string) *QueryMinorVersionResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryMinorVersionResponseBody) SetHttpStatusCode(v int32) *QueryMinorVersionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryMinorVersionResponseBody) SetRequestId(v string) *QueryMinorVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMinorVersionResponseBody) SetSuccess(v bool) *QueryMinorVersionResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMinorVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
