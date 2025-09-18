// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeInstanceConfigsResponseBody
	GetAccessDeniedDetail() *string
	SetData(v []byte) *DescribeInstanceConfigsResponseBody
	GetData() []byte
	SetErrCode(v string) *DescribeInstanceConfigsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeInstanceConfigsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeInstanceConfigsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeInstanceConfigsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInstanceConfigsResponseBody
	GetSuccess() *bool
}

type DescribeInstanceConfigsResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxxx",     "EncodedDiagnosticMessage": "xxxx",     "AuthPrincipalType": "SubUser",     "AuthPrincipalDisplayName": "xxxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "milvus:xxxx" }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The returned result.
	//
	// example:
	//
	// dataCoord:\\n  enableCompaction: true
	Data []byte `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned.
	//
	// example:
	//
	// Instance.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Failed to find instance c-123xxx.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ABCD-1234-5678-EFGH
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeInstanceConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceConfigsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeInstanceConfigsResponseBody) GetData() []byte {
	return s.Data
}

func (s *DescribeInstanceConfigsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeInstanceConfigsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeInstanceConfigsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeInstanceConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceConfigsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInstanceConfigsResponseBody) SetAccessDeniedDetail(v string) *DescribeInstanceConfigsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBody) SetData(v []byte) *DescribeInstanceConfigsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeInstanceConfigsResponseBody) SetErrCode(v string) *DescribeInstanceConfigsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBody) SetErrMessage(v string) *DescribeInstanceConfigsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBody) SetHttpStatusCode(v int32) *DescribeInstanceConfigsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBody) SetRequestId(v string) *DescribeInstanceConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBody) SetSuccess(v bool) *DescribeInstanceConfigsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}
