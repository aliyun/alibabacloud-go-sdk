// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyInstanceConfigResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *ModifyInstanceConfigResponseBody
	GetData() *bool
	SetErrCode(v string) *ModifyInstanceConfigResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyInstanceConfigResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyInstanceConfigResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifyInstanceConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyInstanceConfigResponseBody
	GetSuccess() *bool
}

type ModifyInstanceConfigResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// True
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
	// B1DBBA4E-6CD4-54F7-81E4-12C3425AB519
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyInstanceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceConfigResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyInstanceConfigResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModifyInstanceConfigResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyInstanceConfigResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyInstanceConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyInstanceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyInstanceConfigResponseBody) SetAccessDeniedDetail(v string) *ModifyInstanceConfigResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyInstanceConfigResponseBody) SetData(v bool) *ModifyInstanceConfigResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyInstanceConfigResponseBody) SetErrCode(v string) *ModifyInstanceConfigResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyInstanceConfigResponseBody) SetErrMessage(v string) *ModifyInstanceConfigResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyInstanceConfigResponseBody) SetHttpStatusCode(v int32) *ModifyInstanceConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyInstanceConfigResponseBody) SetRequestId(v string) *ModifyInstanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceConfigResponseBody) SetSuccess(v bool) *ModifyInstanceConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyInstanceConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
