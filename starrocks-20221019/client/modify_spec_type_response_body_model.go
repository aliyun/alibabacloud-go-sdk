// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySpecTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifySpecTypeResponseBody
	GetAccessDeniedDetail() *string
	SetData(v int64) *ModifySpecTypeResponseBody
	GetData() *int64
	SetErrCode(v string) *ModifySpecTypeResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifySpecTypeResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifySpecTypeResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifySpecTypeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifySpecTypeResponseBody
	GetSuccess() *bool
}

type ModifySpecTypeResponseBody struct {
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
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifySpecTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySpecTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySpecTypeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifySpecTypeResponseBody) GetData() *int64 {
	return s.Data
}

func (s *ModifySpecTypeResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifySpecTypeResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifySpecTypeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifySpecTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySpecTypeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifySpecTypeResponseBody) SetAccessDeniedDetail(v string) *ModifySpecTypeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifySpecTypeResponseBody) SetData(v int64) *ModifySpecTypeResponseBody {
	s.Data = &v
	return s
}

func (s *ModifySpecTypeResponseBody) SetErrCode(v string) *ModifySpecTypeResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifySpecTypeResponseBody) SetErrMessage(v string) *ModifySpecTypeResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifySpecTypeResponseBody) SetHttpStatusCode(v int32) *ModifySpecTypeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifySpecTypeResponseBody) SetRequestId(v string) *ModifySpecTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySpecTypeResponseBody) SetSuccess(v bool) *ModifySpecTypeResponseBody {
	s.Success = &v
	return s
}

func (s *ModifySpecTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
