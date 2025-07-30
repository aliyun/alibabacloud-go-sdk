// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskSizeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyDiskSizeResponseBody
	GetAccessDeniedDetail() *string
	SetData(v int64) *ModifyDiskSizeResponseBody
	GetData() *int64
	SetErrCode(v string) *ModifyDiskSizeResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyDiskSizeResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyDiskSizeResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifyDiskSizeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyDiskSizeResponseBody
	GetSuccess() *bool
}

type ModifyDiskSizeResponseBody struct {
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

func (s ModifyDiskSizeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskSizeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDiskSizeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyDiskSizeResponseBody) GetData() *int64 {
	return s.Data
}

func (s *ModifyDiskSizeResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyDiskSizeResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyDiskSizeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyDiskSizeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDiskSizeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyDiskSizeResponseBody) SetAccessDeniedDetail(v string) *ModifyDiskSizeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyDiskSizeResponseBody) SetData(v int64) *ModifyDiskSizeResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyDiskSizeResponseBody) SetErrCode(v string) *ModifyDiskSizeResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyDiskSizeResponseBody) SetErrMessage(v string) *ModifyDiskSizeResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyDiskSizeResponseBody) SetHttpStatusCode(v int32) *ModifyDiskSizeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyDiskSizeResponseBody) SetRequestId(v string) *ModifyDiskSizeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDiskSizeResponseBody) SetSuccess(v bool) *ModifyDiskSizeResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDiskSizeResponseBody) Validate() error {
	return dara.Validate(s)
}
