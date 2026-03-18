// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AddBackupPolicyResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *AddBackupPolicyResponseBody
	GetData() *bool
	SetErrCode(v string) *AddBackupPolicyResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *AddBackupPolicyResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *AddBackupPolicyResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *AddBackupPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddBackupPolicyResponseBody
	GetSuccess() *bool
}

type AddBackupPolicyResponseBody struct {
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
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [Region id should be select from set [cn-beijing, cn-hangzhou]]
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

func (s AddBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *AddBackupPolicyResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AddBackupPolicyResponseBody) GetData() *bool {
	return s.Data
}

func (s *AddBackupPolicyResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *AddBackupPolicyResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *AddBackupPolicyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddBackupPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddBackupPolicyResponseBody) SetAccessDeniedDetail(v string) *AddBackupPolicyResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddBackupPolicyResponseBody) SetData(v bool) *AddBackupPolicyResponseBody {
	s.Data = &v
	return s
}

func (s *AddBackupPolicyResponseBody) SetErrCode(v string) *AddBackupPolicyResponseBody {
	s.ErrCode = &v
	return s
}

func (s *AddBackupPolicyResponseBody) SetErrMessage(v string) *AddBackupPolicyResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *AddBackupPolicyResponseBody) SetHttpStatusCode(v int32) *AddBackupPolicyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddBackupPolicyResponseBody) SetRequestId(v string) *AddBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddBackupPolicyResponseBody) SetSuccess(v bool) *AddBackupPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *AddBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
