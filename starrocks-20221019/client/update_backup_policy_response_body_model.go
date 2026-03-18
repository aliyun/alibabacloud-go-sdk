// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateBackupPolicyResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *UpdateBackupPolicyResponseBody
	GetData() *bool
	SetErrCode(v string) *UpdateBackupPolicyResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *UpdateBackupPolicyResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *UpdateBackupPolicyResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateBackupPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateBackupPolicyResponseBody
	GetSuccess() *bool
}

type UpdateBackupPolicyResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 24151320976****
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// null
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

func (s UpdateBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBackupPolicyResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateBackupPolicyResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateBackupPolicyResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UpdateBackupPolicyResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UpdateBackupPolicyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBackupPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateBackupPolicyResponseBody) SetAccessDeniedDetail(v string) *UpdateBackupPolicyResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateBackupPolicyResponseBody) SetData(v bool) *UpdateBackupPolicyResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateBackupPolicyResponseBody) SetErrCode(v string) *UpdateBackupPolicyResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateBackupPolicyResponseBody) SetErrMessage(v string) *UpdateBackupPolicyResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdateBackupPolicyResponseBody) SetHttpStatusCode(v int32) *UpdateBackupPolicyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateBackupPolicyResponseBody) SetRequestId(v string) *UpdateBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBackupPolicyResponseBody) SetSuccess(v bool) *UpdateBackupPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
