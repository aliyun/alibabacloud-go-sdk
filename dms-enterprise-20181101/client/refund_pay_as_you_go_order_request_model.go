// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundPayAsYouGoOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RefundPayAsYouGoOrderRequest
	GetInstanceId() *string
	SetOrderId(v string) *RefundPayAsYouGoOrderRequest
	GetOrderId() *string
	SetTid(v int64) *RefundPayAsYouGoOrderRequest
	GetTid() *int64
}

type RefundPayAsYouGoOrderRequest struct {
	// The instance ID in the sales order.
	//
	// This parameter is required.
	//
	// example:
	//
	// dms_pre_public_intl-sg-vf****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The order ID of the order for the pay-as-you-go resource. You can call the ListEffectiveOrders operation to query the order ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2190037****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s RefundPayAsYouGoOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s RefundPayAsYouGoOrderRequest) GoString() string {
	return s.String()
}

func (s *RefundPayAsYouGoOrderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RefundPayAsYouGoOrderRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *RefundPayAsYouGoOrderRequest) GetTid() *int64 {
	return s.Tid
}

func (s *RefundPayAsYouGoOrderRequest) SetInstanceId(v string) *RefundPayAsYouGoOrderRequest {
	s.InstanceId = &v
	return s
}

func (s *RefundPayAsYouGoOrderRequest) SetOrderId(v string) *RefundPayAsYouGoOrderRequest {
	s.OrderId = &v
	return s
}

func (s *RefundPayAsYouGoOrderRequest) SetTid(v int64) *RefundPayAsYouGoOrderRequest {
	s.Tid = &v
	return s
}

func (s *RefundPayAsYouGoOrderRequest) Validate() error {
	return dara.Validate(s)
}
