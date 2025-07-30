// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCuResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyCuResponseBody
	GetAccessDeniedDetail() *string
	SetData(v int64) *ModifyCuResponseBody
	GetData() *int64
	SetErrCode(v string) *ModifyCuResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyCuResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyCuResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifyCuResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyCuResponseBody
	GetSuccess() *bool
}

type ModifyCuResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 24151320976****
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s ModifyCuResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCuResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCuResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyCuResponseBody) GetData() *int64 {
	return s.Data
}

func (s *ModifyCuResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyCuResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyCuResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyCuResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCuResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyCuResponseBody) SetAccessDeniedDetail(v string) *ModifyCuResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyCuResponseBody) SetData(v int64) *ModifyCuResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyCuResponseBody) SetErrCode(v string) *ModifyCuResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyCuResponseBody) SetErrMessage(v string) *ModifyCuResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyCuResponseBody) SetHttpStatusCode(v int32) *ModifyCuResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyCuResponseBody) SetRequestId(v string) *ModifyCuResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCuResponseBody) SetSuccess(v bool) *ModifyCuResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyCuResponseBody) Validate() error {
	return dara.Validate(s)
}
