// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInventoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CheckInventoryResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *CheckInventoryResponseBody
	GetData() *bool
	SetErrCode(v string) *CheckInventoryResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CheckInventoryResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *CheckInventoryResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CheckInventoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckInventoryResponseBody
	GetSuccess() *bool
}

type CheckInventoryResponseBody struct {
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
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckInventoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckInventoryResponseBody) GoString() string {
	return s.String()
}

func (s *CheckInventoryResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CheckInventoryResponseBody) GetData() *bool {
	return s.Data
}

func (s *CheckInventoryResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CheckInventoryResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CheckInventoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CheckInventoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckInventoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckInventoryResponseBody) SetAccessDeniedDetail(v string) *CheckInventoryResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CheckInventoryResponseBody) SetData(v bool) *CheckInventoryResponseBody {
	s.Data = &v
	return s
}

func (s *CheckInventoryResponseBody) SetErrCode(v string) *CheckInventoryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CheckInventoryResponseBody) SetErrMessage(v string) *CheckInventoryResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CheckInventoryResponseBody) SetHttpStatusCode(v int32) *CheckInventoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckInventoryResponseBody) SetRequestId(v string) *CheckInventoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckInventoryResponseBody) SetSuccess(v bool) *CheckInventoryResponseBody {
	s.Success = &v
	return s
}

func (s *CheckInventoryResponseBody) Validate() error {
	return dara.Validate(s)
}
