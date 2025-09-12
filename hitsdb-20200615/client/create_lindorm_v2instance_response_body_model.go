// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLindormV2InstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateLindormV2InstanceResponseBody
	GetAccessDeniedDetail() *string
	SetInstanceId(v string) *CreateLindormV2InstanceResponseBody
	GetInstanceId() *string
	SetOrderId(v int64) *CreateLindormV2InstanceResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *CreateLindormV2InstanceResponseBody
	GetRequestId() *string
}

type CreateLindormV2InstanceResponseBody struct {
	// example:
	//
	// {"AuthAction":"xxx","AuthPrincipalDisplayName":"222","AuthPrincipalOwnerId":"111","AuthPrincipalType":"SubUser",,"NoPermissionType":"ImplicitDeny","PolicyType":"AccountLevelIdentityBasedPolicy","EncodedDiagnosticMessage":"xxxxxx"}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 211110656240000
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 1556DCB0-043A-4444-8BD9-CF4A68E7EE64
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLindormV2InstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLindormV2InstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLindormV2InstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateLindormV2InstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateLindormV2InstanceResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateLindormV2InstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLindormV2InstanceResponseBody) SetAccessDeniedDetail(v string) *CreateLindormV2InstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateLindormV2InstanceResponseBody) SetInstanceId(v string) *CreateLindormV2InstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateLindormV2InstanceResponseBody) SetOrderId(v int64) *CreateLindormV2InstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateLindormV2InstanceResponseBody) SetRequestId(v string) *CreateLindormV2InstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLindormV2InstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
