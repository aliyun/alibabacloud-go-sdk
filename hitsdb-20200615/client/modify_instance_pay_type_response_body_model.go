// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstancePayTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyInstancePayTypeResponseBody
	GetAccessDeniedDetail() *string
	SetInstanceId(v string) *ModifyInstancePayTypeResponseBody
	GetInstanceId() *string
	SetOrderId(v int64) *ModifyInstancePayTypeResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *ModifyInstancePayTypeResponseBody
	GetRequestId() *string
}

type ModifyInstancePayTypeResponseBody struct {
	// The detailed reason why the access was denied.
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
	// The ID of the order.
	//
	// example:
	//
	// 211662251220224
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 587BCA54-50DA-4885-ADE9-80A848339151
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstancePayTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstancePayTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstancePayTypeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyInstancePayTypeResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstancePayTypeResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ModifyInstancePayTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstancePayTypeResponseBody) SetAccessDeniedDetail(v string) *ModifyInstancePayTypeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyInstancePayTypeResponseBody) SetInstanceId(v string) *ModifyInstancePayTypeResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstancePayTypeResponseBody) SetOrderId(v int64) *ModifyInstancePayTypeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyInstancePayTypeResponseBody) SetRequestId(v string) *ModifyInstancePayTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstancePayTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
