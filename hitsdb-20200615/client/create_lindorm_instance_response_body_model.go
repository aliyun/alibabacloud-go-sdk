// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLindormInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateLindormInstanceResponseBody
	GetAccessDeniedDetail() *string
	SetInstanceId(v string) *CreateLindormInstanceResponseBody
	GetInstanceId() *string
	SetOrderId(v int64) *CreateLindormInstanceResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *CreateLindormInstanceResponseBody
	GetRequestId() *string
}

type CreateLindormInstanceResponseBody struct {
	// The detailed reason why the access was denied.
	//
	// example:
	//
	// {"AuthAction":"xxx","AuthPrincipalDisplayName":"222","AuthPrincipalOwnerId":"111","AuthPrincipalType":"SubUser",,"NoPermissionType":"ImplicitDeny","PolicyType":"AccountLevelIdentityBasedPolicy","EncodedDiagnosticMessage":"xxxxxx"}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The ID of the Lindorm instance that is created.
	//
	// example:
	//
	// ld-bp1o3y0yme2i2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 111111111111111
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 93BE8227-3406-4D7A-883D-9A421D42****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLindormInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLindormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLindormInstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateLindormInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateLindormInstanceResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateLindormInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLindormInstanceResponseBody) SetAccessDeniedDetail(v string) *CreateLindormInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateLindormInstanceResponseBody) SetInstanceId(v string) *CreateLindormInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateLindormInstanceResponseBody) SetOrderId(v int64) *CreateLindormInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateLindormInstanceResponseBody) SetRequestId(v string) *CreateLindormInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLindormInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
