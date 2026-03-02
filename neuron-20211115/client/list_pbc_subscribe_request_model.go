// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPbcSubscribeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderBy(v string) *ListPbcSubscribeRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListPbcSubscribeRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListPbcSubscribeRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPbcSubscribeRequest
	GetPageSize() *int32
}

type ListPbcSubscribeRequest struct {
	OrderBy        *string `json:"order_by,omitempty" xml:"order_by,omitempty"`
	OrderDirection *string `json:"order_direction,omitempty" xml:"order_direction,omitempty"`
	PageNumber     *int32  `json:"page_number,omitempty" xml:"page_number,omitempty"`
	PageSize       *int32  `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

func (s ListPbcSubscribeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPbcSubscribeRequest) GoString() string {
	return s.String()
}

func (s *ListPbcSubscribeRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListPbcSubscribeRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListPbcSubscribeRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPbcSubscribeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPbcSubscribeRequest) SetOrderBy(v string) *ListPbcSubscribeRequest {
	s.OrderBy = &v
	return s
}

func (s *ListPbcSubscribeRequest) SetOrderDirection(v string) *ListPbcSubscribeRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListPbcSubscribeRequest) SetPageNumber(v int32) *ListPbcSubscribeRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPbcSubscribeRequest) SetPageSize(v int32) *ListPbcSubscribeRequest {
	s.PageSize = &v
	return s
}

func (s *ListPbcSubscribeRequest) Validate() error {
	return dara.Validate(s)
}
