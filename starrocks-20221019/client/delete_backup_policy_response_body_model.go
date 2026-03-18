// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteBackupPolicyResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *DeleteBackupPolicyResponseBody
	GetData() *bool
	SetErrCode(v string) *DeleteBackupPolicyResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DeleteBackupPolicyResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DeleteBackupPolicyResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteBackupPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteBackupPolicyResponseBody
	GetSuccess() *bool
}

type DeleteBackupPolicyResponseBody struct {
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
	// null
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

func (s DeleteBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBackupPolicyResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteBackupPolicyResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteBackupPolicyResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteBackupPolicyResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeleteBackupPolicyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBackupPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteBackupPolicyResponseBody) SetAccessDeniedDetail(v string) *DeleteBackupPolicyResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteBackupPolicyResponseBody) SetData(v bool) *DeleteBackupPolicyResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteBackupPolicyResponseBody) SetErrCode(v string) *DeleteBackupPolicyResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteBackupPolicyResponseBody) SetErrMessage(v string) *DeleteBackupPolicyResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteBackupPolicyResponseBody) SetHttpStatusCode(v int32) *DeleteBackupPolicyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteBackupPolicyResponseBody) SetRequestId(v string) *DeleteBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBackupPolicyResponseBody) SetSuccess(v bool) *DeleteBackupPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
