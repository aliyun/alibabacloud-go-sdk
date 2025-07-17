// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataTrackOrderDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *GetDataTrackOrderDetailRequest
	GetOrderId() *int64
	SetTid(v int64) *GetDataTrackOrderDetailRequest
	GetTid() *int64
}

type GetDataTrackOrderDetailRequest struct {
	// The ID of the ticket. You can call the [ListOrders](https://help.aliyun.com/document_detail/144643.html) operation to query the ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4328****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetDataTrackOrderDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataTrackOrderDetailRequest) GoString() string {
	return s.String()
}

func (s *GetDataTrackOrderDetailRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetDataTrackOrderDetailRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDataTrackOrderDetailRequest) SetOrderId(v int64) *GetDataTrackOrderDetailRequest {
	s.OrderId = &v
	return s
}

func (s *GetDataTrackOrderDetailRequest) SetTid(v int64) *GetDataTrackOrderDetailRequest {
	s.Tid = &v
	return s
}

func (s *GetDataTrackOrderDetailRequest) Validate() error {
	return dara.Validate(s)
}
