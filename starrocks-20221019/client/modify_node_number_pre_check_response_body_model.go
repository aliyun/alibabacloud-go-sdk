// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodeNumberPreCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyNodeNumberPreCheckResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *ModifyNodeNumberPreCheckResponseBodyData) *ModifyNodeNumberPreCheckResponseBody
	GetData() *ModifyNodeNumberPreCheckResponseBodyData
	SetErrCode(v string) *ModifyNodeNumberPreCheckResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyNodeNumberPreCheckResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyNodeNumberPreCheckResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifyNodeNumberPreCheckResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyNodeNumberPreCheckResponseBody
	GetSuccess() *bool
}

type ModifyNodeNumberPreCheckResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The returned data.
	Data *ModifyNodeNumberPreCheckResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ModifyNodeNumberPreCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodeNumberPreCheckResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNodeNumberPreCheckResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyNodeNumberPreCheckResponseBody) GetData() *ModifyNodeNumberPreCheckResponseBodyData {
	return s.Data
}

func (s *ModifyNodeNumberPreCheckResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyNodeNumberPreCheckResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyNodeNumberPreCheckResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyNodeNumberPreCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyNodeNumberPreCheckResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyNodeNumberPreCheckResponseBody) SetAccessDeniedDetail(v string) *ModifyNodeNumberPreCheckResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyNodeNumberPreCheckResponseBody) SetData(v *ModifyNodeNumberPreCheckResponseBodyData) *ModifyNodeNumberPreCheckResponseBody {
	s.Data = v
	return s
}

func (s *ModifyNodeNumberPreCheckResponseBody) SetErrCode(v string) *ModifyNodeNumberPreCheckResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyNodeNumberPreCheckResponseBody) SetErrMessage(v string) *ModifyNodeNumberPreCheckResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyNodeNumberPreCheckResponseBody) SetHttpStatusCode(v int32) *ModifyNodeNumberPreCheckResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyNodeNumberPreCheckResponseBody) SetRequestId(v string) *ModifyNodeNumberPreCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNodeNumberPreCheckResponseBody) SetSuccess(v bool) *ModifyNodeNumberPreCheckResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyNodeNumberPreCheckResponseBody) Validate() error {
	return dara.Validate(s)
}

type ModifyNodeNumberPreCheckResponseBodyData struct {
	// Indicates whether the number of nodes can be modified.
	//
	// example:
	//
	// true
	Allow *bool `json:"Allow,omitempty" xml:"Allow,omitempty"`
	// The reason why the number of nodes cannot be modified.
	//
	// example:
	//
	// Failed to find node group[ng-3d5ce6454354****].
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s ModifyNodeNumberPreCheckResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodeNumberPreCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyNodeNumberPreCheckResponseBodyData) GetAllow() *bool {
	return s.Allow
}

func (s *ModifyNodeNumberPreCheckResponseBodyData) GetReason() *string {
	return s.Reason
}

func (s *ModifyNodeNumberPreCheckResponseBodyData) SetAllow(v bool) *ModifyNodeNumberPreCheckResponseBodyData {
	s.Allow = &v
	return s
}

func (s *ModifyNodeNumberPreCheckResponseBodyData) SetReason(v string) *ModifyNodeNumberPreCheckResponseBodyData {
	s.Reason = &v
	return s
}

func (s *ModifyNodeNumberPreCheckResponseBodyData) Validate() error {
	return dara.Validate(s)
}
