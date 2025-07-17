// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStructSyncOrderDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *GetStructSyncOrderDetailRequest
	GetOrderId() *int64
	SetTid(v int64) *GetStructSyncOrderDetailRequest
	GetTid() *int64
}

type GetStructSyncOrderDetailRequest struct {
	// The ID of the ticket.
	//
	// This parameter is required.
	//
	// example:
	//
	// 43****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the tenant.
	//
	// > To view the tenant ID, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetStructSyncOrderDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStructSyncOrderDetailRequest) GoString() string {
	return s.String()
}

func (s *GetStructSyncOrderDetailRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetStructSyncOrderDetailRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetStructSyncOrderDetailRequest) SetOrderId(v int64) *GetStructSyncOrderDetailRequest {
	s.OrderId = &v
	return s
}

func (s *GetStructSyncOrderDetailRequest) SetTid(v int64) *GetStructSyncOrderDetailRequest {
	s.Tid = &v
	return s
}

func (s *GetStructSyncOrderDetailRequest) Validate() error {
	return dara.Validate(s)
}
