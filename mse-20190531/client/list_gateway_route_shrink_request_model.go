// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayRouteShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListGatewayRouteShrinkRequest
	GetAcceptLanguage() *string
	SetDescSort(v bool) *ListGatewayRouteShrinkRequest
	GetDescSort() *bool
	SetFilterParamsShrink(v string) *ListGatewayRouteShrinkRequest
	GetFilterParamsShrink() *string
	SetOrderItem(v string) *ListGatewayRouteShrinkRequest
	GetOrderItem() *string
	SetPageNumber(v int32) *ListGatewayRouteShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGatewayRouteShrinkRequest
	GetPageSize() *int32
}

type ListGatewayRouteShrinkRequest struct {
	// Language type for returned information:
	//
	// - zh: Chinese
	//
	// - en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Sorting. This field is currently unavailable and being deprecated.
	//
	// example:
	//
	// false
	DescSort *bool `json:"DescSort,omitempty" xml:"DescSort,omitempty"`
	// Filter conditions in the format {"key1":"value1"}
	FilterParamsShrink *string `json:"FilterParams,omitempty" xml:"FilterParams,omitempty"`
	// Sorting property.
	//
	// example:
	//
	// GmtCreate
	OrderItem *string `json:"OrderItem,omitempty" xml:"OrderItem,omitempty"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of instances displayed per page (currently not in effect, queries all instances by default).
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListGatewayRouteShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListGatewayRouteShrinkRequest) GetDescSort() *bool {
	return s.DescSort
}

func (s *ListGatewayRouteShrinkRequest) GetFilterParamsShrink() *string {
	return s.FilterParamsShrink
}

func (s *ListGatewayRouteShrinkRequest) GetOrderItem() *string {
	return s.OrderItem
}

func (s *ListGatewayRouteShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayRouteShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayRouteShrinkRequest) SetAcceptLanguage(v string) *ListGatewayRouteShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListGatewayRouteShrinkRequest) SetDescSort(v bool) *ListGatewayRouteShrinkRequest {
	s.DescSort = &v
	return s
}

func (s *ListGatewayRouteShrinkRequest) SetFilterParamsShrink(v string) *ListGatewayRouteShrinkRequest {
	s.FilterParamsShrink = &v
	return s
}

func (s *ListGatewayRouteShrinkRequest) SetOrderItem(v string) *ListGatewayRouteShrinkRequest {
	s.OrderItem = &v
	return s
}

func (s *ListGatewayRouteShrinkRequest) SetPageNumber(v int32) *ListGatewayRouteShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayRouteShrinkRequest) SetPageSize(v int32) *ListGatewayRouteShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListGatewayRouteShrinkRequest) Validate() error {
	return dara.Validate(s)
}
