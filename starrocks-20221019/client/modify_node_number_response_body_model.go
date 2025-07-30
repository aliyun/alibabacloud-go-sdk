// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodeNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyNodeNumberResponseBody
	GetAccessDeniedDetail() *string
	SetData(v int64) *ModifyNodeNumberResponseBody
	GetData() *int64
	SetErrCode(v string) *ModifyNodeNumberResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyNodeNumberResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyNodeNumberResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifyNodeNumberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyNodeNumberResponseBody
	GetSuccess() *bool
}

type ModifyNodeNumberResponseBody struct {
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

func (s ModifyNodeNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodeNumberResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNodeNumberResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyNodeNumberResponseBody) GetData() *int64 {
	return s.Data
}

func (s *ModifyNodeNumberResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyNodeNumberResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyNodeNumberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyNodeNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyNodeNumberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyNodeNumberResponseBody) SetAccessDeniedDetail(v string) *ModifyNodeNumberResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyNodeNumberResponseBody) SetData(v int64) *ModifyNodeNumberResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyNodeNumberResponseBody) SetErrCode(v string) *ModifyNodeNumberResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyNodeNumberResponseBody) SetErrMessage(v string) *ModifyNodeNumberResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyNodeNumberResponseBody) SetHttpStatusCode(v int32) *ModifyNodeNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyNodeNumberResponseBody) SetRequestId(v string) *ModifyNodeNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNodeNumberResponseBody) SetSuccess(v bool) *ModifyNodeNumberResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyNodeNumberResponseBody) Validate() error {
	return dara.Validate(s)
}
