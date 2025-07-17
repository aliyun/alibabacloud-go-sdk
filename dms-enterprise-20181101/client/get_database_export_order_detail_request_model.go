// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatabaseExportOrderDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *GetDatabaseExportOrderDetailRequest
	GetOrderId() *int64
	SetTid(v int64) *GetDatabaseExportOrderDetailRequest
	GetTid() *int64
}

type GetDatabaseExportOrderDetailRequest struct {
	// The ticket ID. You can call the [ListOrders](https://help.aliyun.com/document_detail/144643.html) operation to obtain the ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 821****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetDatabaseExportOrderDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseExportOrderDetailRequest) GoString() string {
	return s.String()
}

func (s *GetDatabaseExportOrderDetailRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetDatabaseExportOrderDetailRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDatabaseExportOrderDetailRequest) SetOrderId(v int64) *GetDatabaseExportOrderDetailRequest {
	s.OrderId = &v
	return s
}

func (s *GetDatabaseExportOrderDetailRequest) SetTid(v int64) *GetDatabaseExportOrderDetailRequest {
	s.Tid = &v
	return s
}

func (s *GetDatabaseExportOrderDetailRequest) Validate() error {
	return dara.Validate(s)
}
