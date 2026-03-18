// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceMetaTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeInstanceMetaTokenResponseBody
	GetAccessDeniedDetail() *string
	SetData(v string) *DescribeInstanceMetaTokenResponseBody
	GetData() *string
	SetErrCode(v string) *DescribeInstanceMetaTokenResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeInstanceMetaTokenResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeInstanceMetaTokenResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeInstanceMetaTokenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInstanceMetaTokenResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribeInstanceMetaTokenResponseBody
	GetTotal() *int32
}

type DescribeInstanceMetaTokenResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 4ac66c5b-f3ef-4fce-8b52-d7f5f9*******
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 3
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeInstanceMetaTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMetaTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMetaTokenResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeInstanceMetaTokenResponseBody) GetData() *string {
	return s.Data
}

func (s *DescribeInstanceMetaTokenResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeInstanceMetaTokenResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeInstanceMetaTokenResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeInstanceMetaTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceMetaTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInstanceMetaTokenResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeInstanceMetaTokenResponseBody) SetAccessDeniedDetail(v string) *DescribeInstanceMetaTokenResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeInstanceMetaTokenResponseBody) SetData(v string) *DescribeInstanceMetaTokenResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeInstanceMetaTokenResponseBody) SetErrCode(v string) *DescribeInstanceMetaTokenResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeInstanceMetaTokenResponseBody) SetErrMessage(v string) *DescribeInstanceMetaTokenResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeInstanceMetaTokenResponseBody) SetHttpStatusCode(v int32) *DescribeInstanceMetaTokenResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeInstanceMetaTokenResponseBody) SetRequestId(v string) *DescribeInstanceMetaTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceMetaTokenResponseBody) SetSuccess(v bool) *DescribeInstanceMetaTokenResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceMetaTokenResponseBody) SetTotal(v int32) *DescribeInstanceMetaTokenResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeInstanceMetaTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
