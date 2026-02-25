// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyChargeTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyChargeTypeResponseBody
	GetAccessDeniedDetail() *string
	SetData(v int64) *ModifyChargeTypeResponseBody
	GetData() *int64
	SetErrCode(v string) *ModifyChargeTypeResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyChargeTypeResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyChargeTypeResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifyChargeTypeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyChargeTypeResponseBody
	GetSuccess() *bool
}

type ModifyChargeTypeResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 24151320976****
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyChargeTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyChargeTypeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyChargeTypeResponseBody) GetData() *int64 {
	return s.Data
}

func (s *ModifyChargeTypeResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyChargeTypeResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyChargeTypeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyChargeTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyChargeTypeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyChargeTypeResponseBody) SetAccessDeniedDetail(v string) *ModifyChargeTypeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyChargeTypeResponseBody) SetData(v int64) *ModifyChargeTypeResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyChargeTypeResponseBody) SetErrCode(v string) *ModifyChargeTypeResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyChargeTypeResponseBody) SetErrMessage(v string) *ModifyChargeTypeResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyChargeTypeResponseBody) SetHttpStatusCode(v int32) *ModifyChargeTypeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyChargeTypeResponseBody) SetRequestId(v string) *ModifyChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyChargeTypeResponseBody) SetSuccess(v bool) *ModifyChargeTypeResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyChargeTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
