// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessControlListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeAccessControlListResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *DescribeAccessControlListResponseBodyData) *DescribeAccessControlListResponseBody
	GetData() *DescribeAccessControlListResponseBodyData
	SetErrMessage(v string) *DescribeAccessControlListResponseBody
	GetErrMessage() *string
	SetErrorCode(v string) *DescribeAccessControlListResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *DescribeAccessControlListResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeAccessControlListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAccessControlListResponseBody
	GetSuccess() *bool
}

type DescribeAccessControlListResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxxx",     "EncodedDiagnosticMessage": "xxxx",     "AuthPrincipalType": "SubUser",     "AuthPrincipalDisplayName": "xxxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "milvus:xxxx" }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The returned result.
	Data *DescribeAccessControlListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Failed to find instance c-123xxx
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error code returned.
	//
	// example:
	//
	// Instance.NotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
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

func (s DescribeAccessControlListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeAccessControlListResponseBody) GetData() *DescribeAccessControlListResponseBodyData {
	return s.Data
}

func (s *DescribeAccessControlListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeAccessControlListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeAccessControlListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeAccessControlListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccessControlListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAccessControlListResponseBody) SetAccessDeniedDetail(v string) *DescribeAccessControlListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeAccessControlListResponseBody) SetData(v *DescribeAccessControlListResponseBodyData) *DescribeAccessControlListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAccessControlListResponseBody) SetErrMessage(v string) *DescribeAccessControlListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeAccessControlListResponseBody) SetErrorCode(v string) *DescribeAccessControlListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeAccessControlListResponseBody) SetHttpStatusCode(v int32) *DescribeAccessControlListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAccessControlListResponseBody) SetRequestId(v string) *DescribeAccessControlListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessControlListResponseBody) SetSuccess(v bool) *DescribeAccessControlListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAccessControlListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAccessControlListResponseBodyData struct {
	// AclId for public network access control.
	//
	// example:
	//
	// acl-xxxx
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The CIDR blocks.
	Cidr []*string `json:"Cidr,omitempty" xml:"Cidr,omitempty" type:"Repeated"`
}

func (s DescribeAccessControlListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListResponseBodyData) GetAclId() *string {
	return s.AclId
}

func (s *DescribeAccessControlListResponseBodyData) GetCidr() []*string {
	return s.Cidr
}

func (s *DescribeAccessControlListResponseBodyData) SetAclId(v string) *DescribeAccessControlListResponseBodyData {
	s.AclId = &v
	return s
}

func (s *DescribeAccessControlListResponseBodyData) SetCidr(v []*string) *DescribeAccessControlListResponseBodyData {
	s.Cidr = v
	return s
}

func (s *DescribeAccessControlListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
