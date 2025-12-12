// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLindormV2WhiteIpListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateLindormV2WhiteIpListResponseBody
	GetAccessDeniedDetail() *string
	SetRequestId(v string) *UpdateLindormV2WhiteIpListResponseBody
	GetRequestId() *string
}

type UpdateLindormV2WhiteIpListResponseBody struct {
	// example:
	//
	// {"AuthAction":"xxx","AuthPrincipalDisplayName":"222","AuthPrincipalOwnerId":"111","AuthPrincipalType":"SubUser",,"NoPermissionType":"ImplicitDeny","PolicyType":"AccountLevelIdentityBasedPolicy","EncodedDiagnosticMessage":"xxxxxx"}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 0A7153E4-8354-497E-87E5-5D0EBEF5AEB1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLindormV2WhiteIpListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLindormV2WhiteIpListResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLindormV2WhiteIpListResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateLindormV2WhiteIpListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLindormV2WhiteIpListResponseBody) SetAccessDeniedDetail(v string) *UpdateLindormV2WhiteIpListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateLindormV2WhiteIpListResponseBody) SetRequestId(v string) *UpdateLindormV2WhiteIpListResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLindormV2WhiteIpListResponseBody) Validate() error {
	return dara.Validate(s)
}
