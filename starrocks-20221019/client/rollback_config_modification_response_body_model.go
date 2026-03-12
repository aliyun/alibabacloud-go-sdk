// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackConfigModificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RollbackConfigModificationResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *RollbackConfigModificationResponseBody
	GetData() *bool
	SetErrCode(v string) *RollbackConfigModificationResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *RollbackConfigModificationResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *RollbackConfigModificationResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *RollbackConfigModificationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RollbackConfigModificationResponseBody
	GetSuccess() *bool
}

type RollbackConfigModificationResponseBody struct {
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
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RollbackConfigModificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RollbackConfigModificationResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackConfigModificationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RollbackConfigModificationResponseBody) GetData() *bool {
	return s.Data
}

func (s *RollbackConfigModificationResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *RollbackConfigModificationResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *RollbackConfigModificationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RollbackConfigModificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RollbackConfigModificationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RollbackConfigModificationResponseBody) SetAccessDeniedDetail(v string) *RollbackConfigModificationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RollbackConfigModificationResponseBody) SetData(v bool) *RollbackConfigModificationResponseBody {
	s.Data = &v
	return s
}

func (s *RollbackConfigModificationResponseBody) SetErrCode(v string) *RollbackConfigModificationResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RollbackConfigModificationResponseBody) SetErrMessage(v string) *RollbackConfigModificationResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *RollbackConfigModificationResponseBody) SetHttpStatusCode(v int32) *RollbackConfigModificationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RollbackConfigModificationResponseBody) SetRequestId(v string) *RollbackConfigModificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackConfigModificationResponseBody) SetSuccess(v bool) *RollbackConfigModificationResponseBody {
	s.Success = &v
	return s
}

func (s *RollbackConfigModificationResponseBody) Validate() error {
	return dara.Validate(s)
}
