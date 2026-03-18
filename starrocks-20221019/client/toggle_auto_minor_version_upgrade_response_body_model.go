// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iToggleAutoMinorVersionUpgradeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ToggleAutoMinorVersionUpgradeResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *ToggleAutoMinorVersionUpgradeResponseBody
	GetData() *bool
	SetErrCode(v string) *ToggleAutoMinorVersionUpgradeResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ToggleAutoMinorVersionUpgradeResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ToggleAutoMinorVersionUpgradeResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ToggleAutoMinorVersionUpgradeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ToggleAutoMinorVersionUpgradeResponseBody
	GetSuccess() *bool
}

type ToggleAutoMinorVersionUpgradeResponseBody struct {
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

func (s ToggleAutoMinorVersionUpgradeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ToggleAutoMinorVersionUpgradeResponseBody) GoString() string {
	return s.String()
}

func (s *ToggleAutoMinorVersionUpgradeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ToggleAutoMinorVersionUpgradeResponseBody) GetData() *bool {
	return s.Data
}

func (s *ToggleAutoMinorVersionUpgradeResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ToggleAutoMinorVersionUpgradeResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ToggleAutoMinorVersionUpgradeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ToggleAutoMinorVersionUpgradeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ToggleAutoMinorVersionUpgradeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ToggleAutoMinorVersionUpgradeResponseBody) SetAccessDeniedDetail(v string) *ToggleAutoMinorVersionUpgradeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ToggleAutoMinorVersionUpgradeResponseBody) SetData(v bool) *ToggleAutoMinorVersionUpgradeResponseBody {
	s.Data = &v
	return s
}

func (s *ToggleAutoMinorVersionUpgradeResponseBody) SetErrCode(v string) *ToggleAutoMinorVersionUpgradeResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ToggleAutoMinorVersionUpgradeResponseBody) SetErrMessage(v string) *ToggleAutoMinorVersionUpgradeResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ToggleAutoMinorVersionUpgradeResponseBody) SetHttpStatusCode(v int32) *ToggleAutoMinorVersionUpgradeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ToggleAutoMinorVersionUpgradeResponseBody) SetRequestId(v string) *ToggleAutoMinorVersionUpgradeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ToggleAutoMinorVersionUpgradeResponseBody) SetSuccess(v bool) *ToggleAutoMinorVersionUpgradeResponseBody {
	s.Success = &v
	return s
}

func (s *ToggleAutoMinorVersionUpgradeResponseBody) Validate() error {
	return dara.Validate(s)
}
