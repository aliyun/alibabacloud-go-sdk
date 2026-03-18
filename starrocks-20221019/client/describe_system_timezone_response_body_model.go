// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemTimezoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeSystemTimezoneResponseBody
	GetAccessDeniedDetail() *string
	SetData(v string) *DescribeSystemTimezoneResponseBody
	GetData() *string
	SetErrCode(v string) *DescribeSystemTimezoneResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeSystemTimezoneResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeSystemTimezoneResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeSystemTimezoneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSystemTimezoneResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribeSystemTimezoneResponseBody
	GetTotal() *int32
}

type DescribeSystemTimezoneResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// Asia/Shanghai
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// 32A44F0D-BFF6-5664-999A-218BBDE74XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeSystemTimezoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemTimezoneResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSystemTimezoneResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeSystemTimezoneResponseBody) GetData() *string {
	return s.Data
}

func (s *DescribeSystemTimezoneResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeSystemTimezoneResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeSystemTimezoneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeSystemTimezoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSystemTimezoneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSystemTimezoneResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeSystemTimezoneResponseBody) SetAccessDeniedDetail(v string) *DescribeSystemTimezoneResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeSystemTimezoneResponseBody) SetData(v string) *DescribeSystemTimezoneResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeSystemTimezoneResponseBody) SetErrCode(v string) *DescribeSystemTimezoneResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSystemTimezoneResponseBody) SetErrMessage(v string) *DescribeSystemTimezoneResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSystemTimezoneResponseBody) SetHttpStatusCode(v int32) *DescribeSystemTimezoneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeSystemTimezoneResponseBody) SetRequestId(v string) *DescribeSystemTimezoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSystemTimezoneResponseBody) SetSuccess(v bool) *DescribeSystemTimezoneResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSystemTimezoneResponseBody) SetTotal(v int32) *DescribeSystemTimezoneResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeSystemTimezoneResponseBody) Validate() error {
	return dara.Validate(s)
}
