// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataArchiveOrderDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *GetDataArchiveOrderDetailRequest
	GetOrderId() *int64
	SetTid(v int64) *GetDataArchiveOrderDetailRequest
	GetTid() *int64
}

type GetDataArchiveOrderDetailRequest struct {
	// The IDs of data archiving tickets.
	//
	// This parameter is required.
	//
	// example:
	//
	// 868****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The tenant ID.
	//
	// >  To view the tenant ID, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 5***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetDataArchiveOrderDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataArchiveOrderDetailRequest) GoString() string {
	return s.String()
}

func (s *GetDataArchiveOrderDetailRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetDataArchiveOrderDetailRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDataArchiveOrderDetailRequest) SetOrderId(v int64) *GetDataArchiveOrderDetailRequest {
	s.OrderId = &v
	return s
}

func (s *GetDataArchiveOrderDetailRequest) SetTid(v int64) *GetDataArchiveOrderDetailRequest {
	s.Tid = &v
	return s
}

func (s *GetDataArchiveOrderDetailRequest) Validate() error {
	return dara.Validate(s)
}
