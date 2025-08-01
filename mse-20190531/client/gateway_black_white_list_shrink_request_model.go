// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGatewayBlackWhiteListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GatewayBlackWhiteListShrinkRequest
	GetAcceptLanguage() *string
	SetDescSort(v bool) *GatewayBlackWhiteListShrinkRequest
	GetDescSort() *bool
	SetFilterParamsShrink(v string) *GatewayBlackWhiteListShrinkRequest
	GetFilterParamsShrink() *string
	SetOrderItem(v string) *GatewayBlackWhiteListShrinkRequest
	GetOrderItem() *string
	SetPageNumber(v int32) *GatewayBlackWhiteListShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GatewayBlackWhiteListShrinkRequest
	GetPageSize() *int32
}

type GatewayBlackWhiteListShrinkRequest struct {
	// The language in which you want to display the results. Valid values: zh and en. zh indicates Chinese, which is the default value. en indicates English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// This parameter is unavailable for public use.
	//
	// example:
	//
	// ""
	DescSort *bool `json:"DescSort,omitempty" xml:"DescSort,omitempty"`
	// The filter parameters.
	FilterParamsShrink *string `json:"FilterParams,omitempty" xml:"FilterParams,omitempty"`
	// This parameter is unavailable for public use.
	//
	// example:
	//
	// ""
	OrderItem *string `json:"OrderItem,omitempty" xml:"OrderItem,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 1.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GatewayBlackWhiteListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GatewayBlackWhiteListShrinkRequest) GoString() string {
	return s.String()
}

func (s *GatewayBlackWhiteListShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GatewayBlackWhiteListShrinkRequest) GetDescSort() *bool {
	return s.DescSort
}

func (s *GatewayBlackWhiteListShrinkRequest) GetFilterParamsShrink() *string {
	return s.FilterParamsShrink
}

func (s *GatewayBlackWhiteListShrinkRequest) GetOrderItem() *string {
	return s.OrderItem
}

func (s *GatewayBlackWhiteListShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GatewayBlackWhiteListShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GatewayBlackWhiteListShrinkRequest) SetAcceptLanguage(v string) *GatewayBlackWhiteListShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GatewayBlackWhiteListShrinkRequest) SetDescSort(v bool) *GatewayBlackWhiteListShrinkRequest {
	s.DescSort = &v
	return s
}

func (s *GatewayBlackWhiteListShrinkRequest) SetFilterParamsShrink(v string) *GatewayBlackWhiteListShrinkRequest {
	s.FilterParamsShrink = &v
	return s
}

func (s *GatewayBlackWhiteListShrinkRequest) SetOrderItem(v string) *GatewayBlackWhiteListShrinkRequest {
	s.OrderItem = &v
	return s
}

func (s *GatewayBlackWhiteListShrinkRequest) SetPageNumber(v int32) *GatewayBlackWhiteListShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *GatewayBlackWhiteListShrinkRequest) SetPageSize(v int32) *GatewayBlackWhiteListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GatewayBlackWhiteListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
