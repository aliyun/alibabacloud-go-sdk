// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAndroidInstanceGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *RenewAndroidInstanceGroupsResponseBody
	GetOrderId() *string
	SetRequestId(v string) *RenewAndroidInstanceGroupsResponseBody
	GetRequestId() *string
}

type RenewAndroidInstanceGroupsResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// 22326560487****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4B886792-2051-5DB4-8AE6-C8E45D3B4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewAndroidInstanceGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewAndroidInstanceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *RenewAndroidInstanceGroupsResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *RenewAndroidInstanceGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewAndroidInstanceGroupsResponseBody) SetOrderId(v string) *RenewAndroidInstanceGroupsResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewAndroidInstanceGroupsResponseBody) SetRequestId(v string) *RenewAndroidInstanceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewAndroidInstanceGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
