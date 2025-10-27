// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLindormInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateLindormInstanceAttributeResponseBody
	GetAccessDeniedDetail() *string
	SetRequestId(v string) *UpdateLindormInstanceAttributeResponseBody
	GetRequestId() *string
}

type UpdateLindormInstanceAttributeResponseBody struct {
	// example:
	//
	// {"AuthAction":"xxx","AuthPrincipalDisplayName":"222","AuthPrincipalOwnerId":"111","AuthPrincipalType":"SubUser",,"NoPermissionType":"ImplicitDeny","PolicyType":"AccountLevelIdentityBasedPolicy","EncodedDiagnosticMessage":"xxxxxx"}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 1556DCB0-043A-4444-8BD9-CF4A68E7EE64
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLindormInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLindormInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLindormInstanceAttributeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateLindormInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLindormInstanceAttributeResponseBody) SetAccessDeniedDetail(v string) *UpdateLindormInstanceAttributeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateLindormInstanceAttributeResponseBody) SetRequestId(v string) *UpdateLindormInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLindormInstanceAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
