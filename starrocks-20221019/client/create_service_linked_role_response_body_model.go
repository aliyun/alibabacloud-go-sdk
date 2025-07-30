// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceLinkedRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateServiceLinkedRoleResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *CreateServiceLinkedRoleResponseBody
	GetData() *bool
	SetErrCode(v string) *CreateServiceLinkedRoleResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateServiceLinkedRoleResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *CreateServiceLinkedRoleResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateServiceLinkedRoleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateServiceLinkedRoleResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *CreateServiceLinkedRoleResponseBody
	GetTotal() *int32
}

type CreateServiceLinkedRoleResponseBody struct {
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
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 3
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s CreateServiceLinkedRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateServiceLinkedRoleResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateServiceLinkedRoleResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateServiceLinkedRoleResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateServiceLinkedRoleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateServiceLinkedRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceLinkedRoleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateServiceLinkedRoleResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *CreateServiceLinkedRoleResponseBody) SetAccessDeniedDetail(v string) *CreateServiceLinkedRoleResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBody) SetData(v bool) *CreateServiceLinkedRoleResponseBody {
	s.Data = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBody) SetErrCode(v string) *CreateServiceLinkedRoleResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBody) SetErrMessage(v string) *CreateServiceLinkedRoleResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBody) SetHttpStatusCode(v int32) *CreateServiceLinkedRoleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBody) SetRequestId(v string) *CreateServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBody) SetSuccess(v bool) *CreateServiceLinkedRoleResponseBody {
	s.Success = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBody) SetTotal(v int32) *CreateServiceLinkedRoleResponseBody {
	s.Total = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
