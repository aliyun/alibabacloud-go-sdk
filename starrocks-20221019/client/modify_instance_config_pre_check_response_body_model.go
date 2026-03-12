// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceConfigPreCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyInstanceConfigPreCheckResponseBody
	GetAccessDeniedDetail() *string
	SetData(v []*string) *ModifyInstanceConfigPreCheckResponseBody
	GetData() []*string
	SetErrCode(v string) *ModifyInstanceConfigPreCheckResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyInstanceConfigPreCheckResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyInstanceConfigPreCheckResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifyInstanceConfigPreCheckResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyInstanceConfigPreCheckResponseBody
	GetSuccess() *bool
}

type ModifyInstanceConfigPreCheckResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string   `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Success
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
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyInstanceConfigPreCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceConfigPreCheckResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceConfigPreCheckResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyInstanceConfigPreCheckResponseBody) GetData() []*string {
	return s.Data
}

func (s *ModifyInstanceConfigPreCheckResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyInstanceConfigPreCheckResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyInstanceConfigPreCheckResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyInstanceConfigPreCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceConfigPreCheckResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyInstanceConfigPreCheckResponseBody) SetAccessDeniedDetail(v string) *ModifyInstanceConfigPreCheckResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyInstanceConfigPreCheckResponseBody) SetData(v []*string) *ModifyInstanceConfigPreCheckResponseBody {
	s.Data = v
	return s
}

func (s *ModifyInstanceConfigPreCheckResponseBody) SetErrCode(v string) *ModifyInstanceConfigPreCheckResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyInstanceConfigPreCheckResponseBody) SetErrMessage(v string) *ModifyInstanceConfigPreCheckResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyInstanceConfigPreCheckResponseBody) SetHttpStatusCode(v int32) *ModifyInstanceConfigPreCheckResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyInstanceConfigPreCheckResponseBody) SetRequestId(v string) *ModifyInstanceConfigPreCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceConfigPreCheckResponseBody) SetSuccess(v bool) *ModifyInstanceConfigPreCheckResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyInstanceConfigPreCheckResponseBody) Validate() error {
	return dara.Validate(s)
}
