// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataTrackJobTableMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *GetDataTrackJobTableMetaRequest
	GetOrderId() *int64
	SetTid(v int64) *GetDataTrackJobTableMetaRequest
	GetTid() *int64
}

type GetDataTrackJobTableMetaRequest struct {
	// The ID of the ticket. You can call the [ListOrders](https://help.aliyun.com/document_detail/144643.html) operation to query the ID of the ticket.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to query the ID of the tenant.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetDataTrackJobTableMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataTrackJobTableMetaRequest) GoString() string {
	return s.String()
}

func (s *GetDataTrackJobTableMetaRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetDataTrackJobTableMetaRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDataTrackJobTableMetaRequest) SetOrderId(v int64) *GetDataTrackJobTableMetaRequest {
	s.OrderId = &v
	return s
}

func (s *GetDataTrackJobTableMetaRequest) SetTid(v int64) *GetDataTrackJobTableMetaRequest {
	s.Tid = &v
	return s
}

func (s *GetDataTrackJobTableMetaRequest) Validate() error {
	return dara.Validate(s)
}
