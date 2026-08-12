// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaintainableTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyMaintainableTimeResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *ModifyMaintainableTimeResponseBody
	GetData() *bool
	SetErrCode(v string) *ModifyMaintainableTimeResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyMaintainableTimeResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyMaintainableTimeResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifyMaintainableTimeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyMaintainableTimeResponseBody
	GetSuccess() *bool
}

type ModifyMaintainableTimeResponseBody struct {
	// Details about the access denied error.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether the maintenance window was updated successfully.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// The HTTP status code of the request.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyMaintainableTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaintainableTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMaintainableTimeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyMaintainableTimeResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModifyMaintainableTimeResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyMaintainableTimeResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyMaintainableTimeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyMaintainableTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMaintainableTimeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyMaintainableTimeResponseBody) SetAccessDeniedDetail(v string) *ModifyMaintainableTimeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyMaintainableTimeResponseBody) SetData(v bool) *ModifyMaintainableTimeResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyMaintainableTimeResponseBody) SetErrCode(v string) *ModifyMaintainableTimeResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyMaintainableTimeResponseBody) SetErrMessage(v string) *ModifyMaintainableTimeResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyMaintainableTimeResponseBody) SetHttpStatusCode(v int32) *ModifyMaintainableTimeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyMaintainableTimeResponseBody) SetRequestId(v string) *ModifyMaintainableTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMaintainableTimeResponseBody) SetSuccess(v bool) *ModifyMaintainableTimeResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyMaintainableTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
