// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyUserPasswordResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *ModifyUserPasswordResponseBody
	GetData() *bool
	SetErrCode(v string) *ModifyUserPasswordResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyUserPasswordResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyUserPasswordResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifyUserPasswordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyUserPasswordResponseBody
	GetSuccess() *bool
}

type ModifyUserPasswordResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Success
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
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyUserPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserPasswordResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyUserPasswordResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModifyUserPasswordResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyUserPasswordResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyUserPasswordResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyUserPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyUserPasswordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyUserPasswordResponseBody) SetAccessDeniedDetail(v string) *ModifyUserPasswordResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyUserPasswordResponseBody) SetData(v bool) *ModifyUserPasswordResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyUserPasswordResponseBody) SetErrCode(v string) *ModifyUserPasswordResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyUserPasswordResponseBody) SetErrMessage(v string) *ModifyUserPasswordResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyUserPasswordResponseBody) SetHttpStatusCode(v int32) *ModifyUserPasswordResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyUserPasswordResponseBody) SetRequestId(v string) *ModifyUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyUserPasswordResponseBody) SetSuccess(v bool) *ModifyUserPasswordResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyUserPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
