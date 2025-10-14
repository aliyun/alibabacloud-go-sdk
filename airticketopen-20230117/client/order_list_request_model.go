// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOrderListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBookTimeEnd(v int64) *OrderListRequest
	GetBookTimeEnd() *int64
	SetBookTimeStart(v int64) *OrderListRequest
	GetBookTimeStart() *int64
	SetPageIndex(v int32) *OrderListRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *OrderListRequest
	GetPageSize() *int32
	SetStatus(v int32) *OrderListRequest
	GetStatus() *int32
}

type OrderListRequest struct {
	// latest booking time (timestamp)
	//
	// This parameter is required.
	//
	// example:
	//
	// 1677229005000
	BookTimeEnd *int64 `json:"book_time_end,omitempty" xml:"book_time_end,omitempty"`
	// earliest book time(timestamp)
	//
	// This parameter is required.
	//
	// example:
	//
	// 1677227005000
	BookTimeStart *int64 `json:"book_time_start,omitempty" xml:"book_time_start,omitempty"`
	// pagination query parameters, from which page to start querying,querying starts with 0
	//
	// example:
	//
	// 0
	PageIndex *int32 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// pagination query parameters, how many orders to return
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// which order status will be query
	//
	// 1: order reservation in process
	//
	// 2: order reservation successful
	//
	// 3: order paid
	//
	// 4: order successful
	//
	// 5: order closed
	//
	// example:
	//
	// 4
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s OrderListRequest) String() string {
	return dara.Prettify(s)
}

func (s OrderListRequest) GoString() string {
	return s.String()
}

func (s *OrderListRequest) GetBookTimeEnd() *int64 {
	return s.BookTimeEnd
}

func (s *OrderListRequest) GetBookTimeStart() *int64 {
	return s.BookTimeStart
}

func (s *OrderListRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *OrderListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *OrderListRequest) GetStatus() *int32 {
	return s.Status
}

func (s *OrderListRequest) SetBookTimeEnd(v int64) *OrderListRequest {
	s.BookTimeEnd = &v
	return s
}

func (s *OrderListRequest) SetBookTimeStart(v int64) *OrderListRequest {
	s.BookTimeStart = &v
	return s
}

func (s *OrderListRequest) SetPageIndex(v int32) *OrderListRequest {
	s.PageIndex = &v
	return s
}

func (s *OrderListRequest) SetPageSize(v int32) *OrderListRequest {
	s.PageSize = &v
	return s
}

func (s *OrderListRequest) SetStatus(v int32) *OrderListRequest {
	s.Status = &v
	return s
}

func (s *OrderListRequest) Validate() error {
	return dara.Validate(s)
}
