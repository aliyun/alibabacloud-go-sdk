// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloseReason(v string) *CloseOrderRequest
	GetCloseReason() *string
	SetOrderId(v int64) *CloseOrderRequest
	GetOrderId() *int64
	SetTid(v int64) *CloseOrderRequest
	GetTid() *int64
}

type CloseOrderRequest struct {
	// The reason why the ticket is closed.
	//
	// This parameter is required.
	//
	// example:
	//
	// close reason
	CloseReason *string `json:"CloseReason,omitempty" xml:"CloseReason,omitempty"`
	// The ID of the ticket.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1343
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

func (s CloseOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseOrderRequest) GoString() string {
	return s.String()
}

func (s *CloseOrderRequest) GetCloseReason() *string {
	return s.CloseReason
}

func (s *CloseOrderRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CloseOrderRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CloseOrderRequest) SetCloseReason(v string) *CloseOrderRequest {
	s.CloseReason = &v
	return s
}

func (s *CloseOrderRequest) SetOrderId(v int64) *CloseOrderRequest {
	s.OrderId = &v
	return s
}

func (s *CloseOrderRequest) SetTid(v int64) *CloseOrderRequest {
	s.Tid = &v
	return s
}

func (s *CloseOrderRequest) Validate() error {
	return dara.Validate(s)
}
