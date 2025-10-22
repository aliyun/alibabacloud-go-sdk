// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCuPreCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyCuPreCheckResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *ModifyCuPreCheckResponseBodyData) *ModifyCuPreCheckResponseBody
	GetData() *ModifyCuPreCheckResponseBodyData
	SetErrCode(v string) *ModifyCuPreCheckResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyCuPreCheckResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyCuPreCheckResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifyCuPreCheckResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyCuPreCheckResponseBody
	GetSuccess() *bool
}

type ModifyCuPreCheckResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The returned data.
	Data *ModifyCuPreCheckResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ModifyCuPreCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCuPreCheckResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCuPreCheckResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyCuPreCheckResponseBody) GetData() *ModifyCuPreCheckResponseBodyData {
	return s.Data
}

func (s *ModifyCuPreCheckResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyCuPreCheckResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyCuPreCheckResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyCuPreCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCuPreCheckResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyCuPreCheckResponseBody) SetAccessDeniedDetail(v string) *ModifyCuPreCheckResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyCuPreCheckResponseBody) SetData(v *ModifyCuPreCheckResponseBodyData) *ModifyCuPreCheckResponseBody {
	s.Data = v
	return s
}

func (s *ModifyCuPreCheckResponseBody) SetErrCode(v string) *ModifyCuPreCheckResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyCuPreCheckResponseBody) SetErrMessage(v string) *ModifyCuPreCheckResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyCuPreCheckResponseBody) SetHttpStatusCode(v int32) *ModifyCuPreCheckResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyCuPreCheckResponseBody) SetRequestId(v string) *ModifyCuPreCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCuPreCheckResponseBody) SetSuccess(v bool) *ModifyCuPreCheckResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyCuPreCheckResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyCuPreCheckResponseBodyData struct {
	// Indicates whether the number of CUs can be modified.
	//
	// example:
	//
	// false
	Allow *bool `json:"Allow,omitempty" xml:"Allow,omitempty"`
	// The reason why the number of CUs cannot be modified.
	//
	// example:
	//
	// Failed to find node group[ng-3d5ce6454354****].
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s ModifyCuPreCheckResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyCuPreCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyCuPreCheckResponseBodyData) GetAllow() *bool {
	return s.Allow
}

func (s *ModifyCuPreCheckResponseBodyData) GetReason() *string {
	return s.Reason
}

func (s *ModifyCuPreCheckResponseBodyData) SetAllow(v bool) *ModifyCuPreCheckResponseBodyData {
	s.Allow = &v
	return s
}

func (s *ModifyCuPreCheckResponseBodyData) SetReason(v string) *ModifyCuPreCheckResponseBodyData {
	s.Reason = &v
	return s
}

func (s *ModifyCuPreCheckResponseBodyData) Validate() error {
	return dara.Validate(s)
}
