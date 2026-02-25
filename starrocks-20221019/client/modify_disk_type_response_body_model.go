// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyDiskTypeResponseBody
	GetAccessDeniedDetail() *string
	SetData(v int64) *ModifyDiskTypeResponseBody
	GetData() *int64
	SetErrCode(v string) *ModifyDiskTypeResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyDiskTypeResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyDiskTypeResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifyDiskTypeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyDiskTypeResponseBody
	GetSuccess() *bool
}

type ModifyDiskTypeResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// true
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
	// 32A44F0D-BFF6-5664-999A-218BBDE74XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDiskTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDiskTypeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyDiskTypeResponseBody) GetData() *int64 {
	return s.Data
}

func (s *ModifyDiskTypeResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyDiskTypeResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyDiskTypeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyDiskTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDiskTypeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyDiskTypeResponseBody) SetAccessDeniedDetail(v string) *ModifyDiskTypeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyDiskTypeResponseBody) SetData(v int64) *ModifyDiskTypeResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyDiskTypeResponseBody) SetErrCode(v string) *ModifyDiskTypeResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyDiskTypeResponseBody) SetErrMessage(v string) *ModifyDiskTypeResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyDiskTypeResponseBody) SetHttpStatusCode(v int32) *ModifyDiskTypeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyDiskTypeResponseBody) SetRequestId(v string) *ModifyDiskTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDiskTypeResponseBody) SetSuccess(v bool) *ModifyDiskTypeResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDiskTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
