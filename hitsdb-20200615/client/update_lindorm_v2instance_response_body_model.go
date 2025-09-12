// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLindormV2InstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateLindormV2InstanceResponseBody
	GetAccessDeniedDetail() *string
	SetInstanceId(v string) *UpdateLindormV2InstanceResponseBody
	GetInstanceId() *string
	SetOrderId(v int64) *UpdateLindormV2InstanceResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *UpdateLindormV2InstanceResponseBody
	GetRequestId() *string
}

type UpdateLindormV2InstanceResponseBody struct {
	// example:
	//
	// {"AuthAction":"xxx","AuthPrincipalDisplayName":"222","AuthPrincipalOwnerId":"111","AuthPrincipalType":"SubUser",,"NoPermissionType":"ImplicitDeny","PolicyType":"AccountLevelIdentityBasedPolicy","EncodedDiagnosticMessage":"xxxxxx"}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// ld-bp1478w1603****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 240136741090345
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 1556DCB0-043A-4444-8BD9-CF4A68E7EE64
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLindormV2InstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLindormV2InstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLindormV2InstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateLindormV2InstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateLindormV2InstanceResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *UpdateLindormV2InstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLindormV2InstanceResponseBody) SetAccessDeniedDetail(v string) *UpdateLindormV2InstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateLindormV2InstanceResponseBody) SetInstanceId(v string) *UpdateLindormV2InstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *UpdateLindormV2InstanceResponseBody) SetOrderId(v int64) *UpdateLindormV2InstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpdateLindormV2InstanceResponseBody) SetRequestId(v string) *UpdateLindormV2InstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLindormV2InstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
