// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchActiveStandbyZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SwitchActiveStandbyZonesResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *SwitchActiveStandbyZonesResponseBody
	GetData() *bool
	SetErrCode(v string) *SwitchActiveStandbyZonesResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *SwitchActiveStandbyZonesResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *SwitchActiveStandbyZonesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *SwitchActiveStandbyZonesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SwitchActiveStandbyZonesResponseBody
	GetSuccess() *bool
}

type SwitchActiveStandbyZonesResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 24151320976****
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
}

func (s SwitchActiveStandbyZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchActiveStandbyZonesResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchActiveStandbyZonesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SwitchActiveStandbyZonesResponseBody) GetData() *bool {
	return s.Data
}

func (s *SwitchActiveStandbyZonesResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *SwitchActiveStandbyZonesResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *SwitchActiveStandbyZonesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SwitchActiveStandbyZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchActiveStandbyZonesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SwitchActiveStandbyZonesResponseBody) SetAccessDeniedDetail(v string) *SwitchActiveStandbyZonesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SwitchActiveStandbyZonesResponseBody) SetData(v bool) *SwitchActiveStandbyZonesResponseBody {
	s.Data = &v
	return s
}

func (s *SwitchActiveStandbyZonesResponseBody) SetErrCode(v string) *SwitchActiveStandbyZonesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SwitchActiveStandbyZonesResponseBody) SetErrMessage(v string) *SwitchActiveStandbyZonesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SwitchActiveStandbyZonesResponseBody) SetHttpStatusCode(v int32) *SwitchActiveStandbyZonesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SwitchActiveStandbyZonesResponseBody) SetRequestId(v string) *SwitchActiveStandbyZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchActiveStandbyZonesResponseBody) SetSuccess(v bool) *SwitchActiveStandbyZonesResponseBody {
	s.Success = &v
	return s
}

func (s *SwitchActiveStandbyZonesResponseBody) Validate() error {
	return dara.Validate(s)
}
