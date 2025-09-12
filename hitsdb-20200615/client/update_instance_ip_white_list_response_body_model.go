// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceIpWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateInstanceIpWhiteListResponseBody
	GetAccessDeniedDetail() *string
	SetRequestId(v string) *UpdateInstanceIpWhiteListResponseBody
	GetRequestId() *string
}

type UpdateInstanceIpWhiteListResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// {"AuthAction":"xxx","AuthPrincipalDisplayName":"222","AuthPrincipalOwnerId":"111","AuthPrincipalType":"SubUser",,"NoPermissionType":"ImplicitDeny","PolicyType":"AccountLevelIdentityBasedPolicy","EncodedDiagnosticMessage":"xxxxxx"}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4944539D-D27C-458D-95F1-2DCEB5E0EED5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceIpWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceIpWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceIpWhiteListResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateInstanceIpWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceIpWhiteListResponseBody) SetAccessDeniedDetail(v string) *UpdateInstanceIpWhiteListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateInstanceIpWhiteListResponseBody) SetRequestId(v string) *UpdateInstanceIpWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceIpWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}
