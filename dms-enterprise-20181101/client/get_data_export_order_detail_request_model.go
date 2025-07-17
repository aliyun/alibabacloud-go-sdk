// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataExportOrderDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *GetDataExportOrderDetailRequest
	GetOrderId() *int64
	SetTid(v int64) *GetDataExportOrderDetailRequest
	GetTid() *int64
}

type GetDataExportOrderDetailRequest struct {
	// The ticket ID. You can call the [ListOrders](https://help.aliyun.com/document_detail/465867.html) operation to query the ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the tenant.
	//
	// > To view the tenant ID, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// -1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetDataExportOrderDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataExportOrderDetailRequest) GoString() string {
	return s.String()
}

func (s *GetDataExportOrderDetailRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetDataExportOrderDetailRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDataExportOrderDetailRequest) SetOrderId(v int64) *GetDataExportOrderDetailRequest {
	s.OrderId = &v
	return s
}

func (s *GetDataExportOrderDetailRequest) SetTid(v int64) *GetDataExportOrderDetailRequest {
	s.Tid = &v
	return s
}

func (s *GetDataExportOrderDetailRequest) Validate() error {
	return dara.Validate(s)
}
