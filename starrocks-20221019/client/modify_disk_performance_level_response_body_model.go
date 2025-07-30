// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskPerformanceLevelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyDiskPerformanceLevelResponseBody
	GetAccessDeniedDetail() *string
	SetData(v int64) *ModifyDiskPerformanceLevelResponseBody
	GetData() *int64
	SetErrCode(v string) *ModifyDiskPerformanceLevelResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyDiskPerformanceLevelResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyDiskPerformanceLevelResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifyDiskPerformanceLevelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyDiskPerformanceLevelResponseBody
	GetSuccess() *bool
}

type ModifyDiskPerformanceLevelResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 24151320976****
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s ModifyDiskPerformanceLevelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskPerformanceLevelResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDiskPerformanceLevelResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyDiskPerformanceLevelResponseBody) GetData() *int64 {
	return s.Data
}

func (s *ModifyDiskPerformanceLevelResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyDiskPerformanceLevelResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyDiskPerformanceLevelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyDiskPerformanceLevelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDiskPerformanceLevelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyDiskPerformanceLevelResponseBody) SetAccessDeniedDetail(v string) *ModifyDiskPerformanceLevelResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyDiskPerformanceLevelResponseBody) SetData(v int64) *ModifyDiskPerformanceLevelResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyDiskPerformanceLevelResponseBody) SetErrCode(v string) *ModifyDiskPerformanceLevelResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyDiskPerformanceLevelResponseBody) SetErrMessage(v string) *ModifyDiskPerformanceLevelResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyDiskPerformanceLevelResponseBody) SetHttpStatusCode(v int32) *ModifyDiskPerformanceLevelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyDiskPerformanceLevelResponseBody) SetRequestId(v string) *ModifyDiskPerformanceLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDiskPerformanceLevelResponseBody) SetSuccess(v bool) *ModifyDiskPerformanceLevelResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDiskPerformanceLevelResponseBody) Validate() error {
	return dara.Validate(s)
}
