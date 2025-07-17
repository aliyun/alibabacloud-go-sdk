// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPermApplyOrderDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *GetPermApplyOrderDetailRequest
	GetOrderId() *int64
	SetTid(v int64) *GetPermApplyOrderDetailRequest
	GetTid() *int64
}

type GetPermApplyOrderDetailRequest struct {
	// The ticket ID. You can call the [ListOrders](https://help.aliyun.com/document_detail/465867.html) operation to query the ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 730000
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The tenant ID.
	//
	// > To view the tenant ID, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// -1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetPermApplyOrderDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPermApplyOrderDetailRequest) GoString() string {
	return s.String()
}

func (s *GetPermApplyOrderDetailRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetPermApplyOrderDetailRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetPermApplyOrderDetailRequest) SetOrderId(v int64) *GetPermApplyOrderDetailRequest {
	s.OrderId = &v
	return s
}

func (s *GetPermApplyOrderDetailRequest) SetTid(v int64) *GetPermApplyOrderDetailRequest {
	s.Tid = &v
	return s
}

func (s *GetPermApplyOrderDetailRequest) Validate() error {
	return dara.Validate(s)
}
