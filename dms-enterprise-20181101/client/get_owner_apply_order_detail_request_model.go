// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOwnerApplyOrderDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *GetOwnerApplyOrderDetailRequest
	GetOrderId() *int64
	SetTid(v int64) *GetOwnerApplyOrderDetailRequest
	GetTid() *int64
}

type GetOwnerApplyOrderDetailRequest struct {
	// The ID of the ticket.
	//
	// This parameter is required.
	//
	// example:
	//
	// 730000
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the tenant.
	//
	// > : To view the ID of the tenant, log on to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// -1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetOwnerApplyOrderDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOwnerApplyOrderDetailRequest) GoString() string {
	return s.String()
}

func (s *GetOwnerApplyOrderDetailRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetOwnerApplyOrderDetailRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetOwnerApplyOrderDetailRequest) SetOrderId(v int64) *GetOwnerApplyOrderDetailRequest {
	s.OrderId = &v
	return s
}

func (s *GetOwnerApplyOrderDetailRequest) SetTid(v int64) *GetOwnerApplyOrderDetailRequest {
	s.Tid = &v
	return s
}

func (s *GetOwnerApplyOrderDetailRequest) Validate() error {
	return dara.Validate(s)
}
