// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseLindormV2InstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ReleaseLindormV2InstanceResponseBody
	GetAccessDeniedDetail() *string
	SetRequestId(v string) *ReleaseLindormV2InstanceResponseBody
	GetRequestId() *string
}

type ReleaseLindormV2InstanceResponseBody struct {
	// example:
	//
	// {"AuthAction":"xxx","AuthPrincipalDisplayName":"222","AuthPrincipalOwnerId":"111","AuthPrincipalType":"SubUser",,"NoPermissionType":"ImplicitDeny","PolicyType":"AccountLevelIdentityBasedPolicy","EncodedDiagnosticMessage":"xxxxxx"}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 0A7153E4-8354-497E-87E5-5D0EBEF5AEB1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseLindormV2InstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseLindormV2InstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseLindormV2InstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ReleaseLindormV2InstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseLindormV2InstanceResponseBody) SetAccessDeniedDetail(v string) *ReleaseLindormV2InstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ReleaseLindormV2InstanceResponseBody) SetRequestId(v string) *ReleaseLindormV2InstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseLindormV2InstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
