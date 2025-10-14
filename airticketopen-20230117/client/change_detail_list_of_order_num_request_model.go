// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDetailListOfOrderNumRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderNum(v int64) *ChangeDetailListOfOrderNumRequest
	GetOrderNum() *int64
	SetPageIndex(v int32) *ChangeDetailListOfOrderNumRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ChangeDetailListOfOrderNumRequest
	GetPageSize() *int32
}

type ChangeDetailListOfOrderNumRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4988430***700
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

func (s ChangeDetailListOfOrderNumRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailListOfOrderNumRequest) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumRequest) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *ChangeDetailListOfOrderNumRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ChangeDetailListOfOrderNumRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ChangeDetailListOfOrderNumRequest) SetOrderNum(v int64) *ChangeDetailListOfOrderNumRequest {
	s.OrderNum = &v
	return s
}

func (s *ChangeDetailListOfOrderNumRequest) SetPageIndex(v int32) *ChangeDetailListOfOrderNumRequest {
	s.PageIndex = &v
	return s
}

func (s *ChangeDetailListOfOrderNumRequest) SetPageSize(v int32) *ChangeDetailListOfOrderNumRequest {
	s.PageSize = &v
	return s
}

func (s *ChangeDetailListOfOrderNumRequest) Validate() error {
	return dara.Validate(s)
}
