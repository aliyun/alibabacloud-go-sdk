// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayAuthShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListGatewayAuthShrinkRequest
	GetAcceptLanguage() *string
	SetDescSort(v bool) *ListGatewayAuthShrinkRequest
	GetDescSort() *bool
	SetFilterParamsShrink(v string) *ListGatewayAuthShrinkRequest
	GetFilterParamsShrink() *string
	SetOrderItem(v string) *ListGatewayAuthShrinkRequest
	GetOrderItem() *string
	SetPageNumber(v int32) *ListGatewayAuthShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGatewayAuthShrinkRequest
	GetPageSize() *int32
}

type ListGatewayAuthShrinkRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// false
	DescSort           *bool   `json:"DescSort,omitempty" xml:"DescSort,omitempty"`
	FilterParamsShrink *string `json:"FilterParams,omitempty" xml:"FilterParams,omitempty"`
	// example:
	//
	// {}
	OrderItem *string `json:"OrderItem,omitempty" xml:"OrderItem,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListGatewayAuthShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListGatewayAuthShrinkRequest) GetDescSort() *bool {
	return s.DescSort
}

func (s *ListGatewayAuthShrinkRequest) GetFilterParamsShrink() *string {
	return s.FilterParamsShrink
}

func (s *ListGatewayAuthShrinkRequest) GetOrderItem() *string {
	return s.OrderItem
}

func (s *ListGatewayAuthShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayAuthShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayAuthShrinkRequest) SetAcceptLanguage(v string) *ListGatewayAuthShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListGatewayAuthShrinkRequest) SetDescSort(v bool) *ListGatewayAuthShrinkRequest {
	s.DescSort = &v
	return s
}

func (s *ListGatewayAuthShrinkRequest) SetFilterParamsShrink(v string) *ListGatewayAuthShrinkRequest {
	s.FilterParamsShrink = &v
	return s
}

func (s *ListGatewayAuthShrinkRequest) SetOrderItem(v string) *ListGatewayAuthShrinkRequest {
	s.OrderItem = &v
	return s
}

func (s *ListGatewayAuthShrinkRequest) SetPageNumber(v int32) *ListGatewayAuthShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayAuthShrinkRequest) SetPageSize(v int32) *ListGatewayAuthShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListGatewayAuthShrinkRequest) Validate() error {
	return dara.Validate(s)
}
