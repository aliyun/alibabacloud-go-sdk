// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySpecTypePreCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifySpecTypePreCheckResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *ModifySpecTypePreCheckResponseBodyData) *ModifySpecTypePreCheckResponseBody
	GetData() *ModifySpecTypePreCheckResponseBodyData
	SetErrCode(v string) *ModifySpecTypePreCheckResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifySpecTypePreCheckResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifySpecTypePreCheckResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifySpecTypePreCheckResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifySpecTypePreCheckResponseBody
	GetSuccess() *bool
}

type ModifySpecTypePreCheckResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string                                 `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *ModifySpecTypePreCheckResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Success
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [Region id should be select from set [cn-beijing, cn-hangzhou]]
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

func (s ModifySpecTypePreCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySpecTypePreCheckResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySpecTypePreCheckResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifySpecTypePreCheckResponseBody) GetData() *ModifySpecTypePreCheckResponseBodyData {
	return s.Data
}

func (s *ModifySpecTypePreCheckResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifySpecTypePreCheckResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifySpecTypePreCheckResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifySpecTypePreCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySpecTypePreCheckResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifySpecTypePreCheckResponseBody) SetAccessDeniedDetail(v string) *ModifySpecTypePreCheckResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifySpecTypePreCheckResponseBody) SetData(v *ModifySpecTypePreCheckResponseBodyData) *ModifySpecTypePreCheckResponseBody {
	s.Data = v
	return s
}

func (s *ModifySpecTypePreCheckResponseBody) SetErrCode(v string) *ModifySpecTypePreCheckResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifySpecTypePreCheckResponseBody) SetErrMessage(v string) *ModifySpecTypePreCheckResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifySpecTypePreCheckResponseBody) SetHttpStatusCode(v int32) *ModifySpecTypePreCheckResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifySpecTypePreCheckResponseBody) SetRequestId(v string) *ModifySpecTypePreCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySpecTypePreCheckResponseBody) SetSuccess(v bool) *ModifySpecTypePreCheckResponseBody {
	s.Success = &v
	return s
}

func (s *ModifySpecTypePreCheckResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifySpecTypePreCheckResponseBodyData struct {
	// example:
	//
	// false
	Allow *bool `json:"Allow,omitempty" xml:"Allow,omitempty"`
	// example:
	//
	// Failed to find node group[ng-3d5ce6454354****].
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s ModifySpecTypePreCheckResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifySpecTypePreCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifySpecTypePreCheckResponseBodyData) GetAllow() *bool {
	return s.Allow
}

func (s *ModifySpecTypePreCheckResponseBodyData) GetReason() *string {
	return s.Reason
}

func (s *ModifySpecTypePreCheckResponseBodyData) SetAllow(v bool) *ModifySpecTypePreCheckResponseBodyData {
	s.Allow = &v
	return s
}

func (s *ModifySpecTypePreCheckResponseBodyData) SetReason(v string) *ModifySpecTypePreCheckResponseBodyData {
	s.Reason = &v
	return s
}

func (s *ModifySpecTypePreCheckResponseBodyData) Validate() error {
	return dara.Validate(s)
}
