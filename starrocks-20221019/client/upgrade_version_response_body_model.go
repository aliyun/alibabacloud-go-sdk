// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpgradeVersionResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *UpgradeVersionResponseBody
	GetData() *bool
	SetErrCode(v string) *UpgradeVersionResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *UpgradeVersionResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *UpgradeVersionResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpgradeVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpgradeVersionResponseBody
	GetSuccess() *bool
}

type UpgradeVersionResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The returned data.
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
	// Invalid instance status: [cannot upgrade when instance is not running].
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

func (s UpgradeVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeVersionResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpgradeVersionResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpgradeVersionResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UpgradeVersionResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UpgradeVersionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpgradeVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpgradeVersionResponseBody) SetAccessDeniedDetail(v string) *UpgradeVersionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpgradeVersionResponseBody) SetData(v bool) *UpgradeVersionResponseBody {
	s.Data = &v
	return s
}

func (s *UpgradeVersionResponseBody) SetErrCode(v string) *UpgradeVersionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpgradeVersionResponseBody) SetErrMessage(v string) *UpgradeVersionResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpgradeVersionResponseBody) SetHttpStatusCode(v int32) *UpgradeVersionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpgradeVersionResponseBody) SetRequestId(v string) *UpgradeVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeVersionResponseBody) SetSuccess(v bool) *UpgradeVersionResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
