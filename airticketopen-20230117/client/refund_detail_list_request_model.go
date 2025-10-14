// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundDetailListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderNum(v int64) *RefundDetailListRequest
	GetOrderNum() *int64
	SetPageIndex(v int32) *RefundDetailListRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *RefundDetailListRequest
	GetPageSize() *int32
	SetRefundCreateBeginTime(v int64) *RefundDetailListRequest
	GetRefundCreateBeginTime() *int64
	SetRefundCreateEndTime(v int64) *RefundDetailListRequest
	GetRefundCreateEndTime() *int64
}

type RefundDetailListRequest struct {
	// order number returned by Book
	//
	// example:
	//
	// 49884*****950
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// pagination query parameters, from which page to start querying
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// pagination query parameters, how many orders to return
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// the earliest time(timestamp) of refund order creation
	//
	// This parameter is required.
	//
	// example:
	//
	// 1677229002000
	RefundCreateBeginTime *int64 `json:"refund_create_begin_time,omitempty" xml:"refund_create_begin_time,omitempty"`
	// the latest time(timestamp) of refund order creation
	//
	// This parameter is required.
	//
	// example:
	//
	// 1677229005000
	RefundCreateEndTime *int64 `json:"refund_create_end_time,omitempty" xml:"refund_create_end_time,omitempty"`
}

func (s RefundDetailListRequest) String() string {
	return dara.Prettify(s)
}

func (s RefundDetailListRequest) GoString() string {
	return s.String()
}

func (s *RefundDetailListRequest) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *RefundDetailListRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *RefundDetailListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *RefundDetailListRequest) GetRefundCreateBeginTime() *int64 {
	return s.RefundCreateBeginTime
}

func (s *RefundDetailListRequest) GetRefundCreateEndTime() *int64 {
	return s.RefundCreateEndTime
}

func (s *RefundDetailListRequest) SetOrderNum(v int64) *RefundDetailListRequest {
	s.OrderNum = &v
	return s
}

func (s *RefundDetailListRequest) SetPageIndex(v int32) *RefundDetailListRequest {
	s.PageIndex = &v
	return s
}

func (s *RefundDetailListRequest) SetPageSize(v int32) *RefundDetailListRequest {
	s.PageSize = &v
	return s
}

func (s *RefundDetailListRequest) SetRefundCreateBeginTime(v int64) *RefundDetailListRequest {
	s.RefundCreateBeginTime = &v
	return s
}

func (s *RefundDetailListRequest) SetRefundCreateEndTime(v int64) *RefundDetailListRequest {
	s.RefundCreateEndTime = &v
	return s
}

func (s *RefundDetailListRequest) Validate() error {
	return dara.Validate(s)
}
