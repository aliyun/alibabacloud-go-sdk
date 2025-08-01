// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayServiceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListGatewayServiceShrinkRequest
	GetAcceptLanguage() *string
	SetDescSort(v bool) *ListGatewayServiceShrinkRequest
	GetDescSort() *bool
	SetFilterParamsShrink(v string) *ListGatewayServiceShrinkRequest
	GetFilterParamsShrink() *string
	SetOrderItem(v string) *ListGatewayServiceShrinkRequest
	GetOrderItem() *string
	SetPageNumber(v int32) *ListGatewayServiceShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGatewayServiceShrinkRequest
	GetPageSize() *int32
}

type ListGatewayServiceShrinkRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Specifies whether to enable sorting.
	//
	// example:
	//
	// false
	DescSort *bool `json:"DescSort,omitempty" xml:"DescSort,omitempty"`
	// The parameters that are used to specify filter conditions. The values of the parameters are in the format of {"key1":"value1"}.
	FilterParamsShrink *string `json:"FilterParams,omitempty" xml:"FilterParams,omitempty"`
	// The item based on which entries are sorted.
	//
	// example:
	//
	// GmtCreate
	OrderItem *string `json:"OrderItem,omitempty" xml:"OrderItem,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListGatewayServiceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayServiceShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayServiceShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListGatewayServiceShrinkRequest) GetDescSort() *bool {
	return s.DescSort
}

func (s *ListGatewayServiceShrinkRequest) GetFilterParamsShrink() *string {
	return s.FilterParamsShrink
}

func (s *ListGatewayServiceShrinkRequest) GetOrderItem() *string {
	return s.OrderItem
}

func (s *ListGatewayServiceShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayServiceShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayServiceShrinkRequest) SetAcceptLanguage(v string) *ListGatewayServiceShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListGatewayServiceShrinkRequest) SetDescSort(v bool) *ListGatewayServiceShrinkRequest {
	s.DescSort = &v
	return s
}

func (s *ListGatewayServiceShrinkRequest) SetFilterParamsShrink(v string) *ListGatewayServiceShrinkRequest {
	s.FilterParamsShrink = &v
	return s
}

func (s *ListGatewayServiceShrinkRequest) SetOrderItem(v string) *ListGatewayServiceShrinkRequest {
	s.OrderItem = &v
	return s
}

func (s *ListGatewayServiceShrinkRequest) SetPageNumber(v int32) *ListGatewayServiceShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayServiceShrinkRequest) SetPageSize(v int32) *ListGatewayServiceShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListGatewayServiceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
