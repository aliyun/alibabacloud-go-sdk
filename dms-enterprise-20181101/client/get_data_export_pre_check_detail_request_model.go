// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataExportPreCheckDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *GetDataExportPreCheckDetailRequest
	GetOrderId() *int64
	SetTid(v int64) *GetDataExportPreCheckDetailRequest
	GetTid() *int64
}

type GetDataExportPreCheckDetailRequest struct {
	// The ID of the ticket. You can call the [ListOrders](https://help.aliyun.com/document_detail/144643.html) operation to obtain the ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 34****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetDataExportPreCheckDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataExportPreCheckDetailRequest) GoString() string {
	return s.String()
}

func (s *GetDataExportPreCheckDetailRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetDataExportPreCheckDetailRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDataExportPreCheckDetailRequest) SetOrderId(v int64) *GetDataExportPreCheckDetailRequest {
	s.OrderId = &v
	return s
}

func (s *GetDataExportPreCheckDetailRequest) SetTid(v int64) *GetDataExportPreCheckDetailRequest {
	s.Tid = &v
	return s
}

func (s *GetDataExportPreCheckDetailRequest) Validate() error {
	return dara.Validate(s)
}
