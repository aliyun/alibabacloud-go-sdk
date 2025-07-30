// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyDiskNumberResponseBody
	GetAccessDeniedDetail() *string
	SetData(v int64) *ModifyDiskNumberResponseBody
	GetData() *int64
	SetErrCode(v string) *ModifyDiskNumberResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyDiskNumberResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyDiskNumberResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifyDiskNumberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyDiskNumberResponseBody
	GetSuccess() *bool
}

type ModifyDiskNumberResponseBody struct {
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

func (s ModifyDiskNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskNumberResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDiskNumberResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyDiskNumberResponseBody) GetData() *int64 {
	return s.Data
}

func (s *ModifyDiskNumberResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyDiskNumberResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyDiskNumberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyDiskNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDiskNumberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyDiskNumberResponseBody) SetAccessDeniedDetail(v string) *ModifyDiskNumberResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyDiskNumberResponseBody) SetData(v int64) *ModifyDiskNumberResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyDiskNumberResponseBody) SetErrCode(v string) *ModifyDiskNumberResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyDiskNumberResponseBody) SetErrMessage(v string) *ModifyDiskNumberResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyDiskNumberResponseBody) SetHttpStatusCode(v int32) *ModifyDiskNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyDiskNumberResponseBody) SetRequestId(v string) *ModifyDiskNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDiskNumberResponseBody) SetSuccess(v bool) *ModifyDiskNumberResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDiskNumberResponseBody) Validate() error {
	return dara.Validate(s)
}
