// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewLindormInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RenewLindormInstanceResponseBody
	GetAccessDeniedDetail() *string
	SetInstanceId(v string) *RenewLindormInstanceResponseBody
	GetInstanceId() *string
	SetOrderId(v int64) *RenewLindormInstanceResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *RenewLindormInstanceResponseBody
	GetRequestId() *string
}

type RenewLindormInstanceResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// {"AuthAction":"xxx","AuthPrincipalDisplayName":"222","AuthPrincipalOwnerId":"111","AuthPrincipalType":"SubUser",,"NoPermissionType":"ImplicitDeny","PolicyType":"AccountLevelIdentityBasedPolicy","EncodedDiagnosticMessage":"xxxxxx"}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ld-bp1z3506imz2f****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the order. You can obtain the order ID on the Orders page of the Expenses and Costs console.
	//
	// example:
	//
	// 213465921640411
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1556DCB0-043A-4444-8BD9-CF4A68E7EE64
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewLindormInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewLindormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewLindormInstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RenewLindormInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RenewLindormInstanceResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *RenewLindormInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewLindormInstanceResponseBody) SetAccessDeniedDetail(v string) *RenewLindormInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RenewLindormInstanceResponseBody) SetInstanceId(v string) *RenewLindormInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *RenewLindormInstanceResponseBody) SetOrderId(v int64) *RenewLindormInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewLindormInstanceResponseBody) SetRequestId(v string) *RenewLindormInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewLindormInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
