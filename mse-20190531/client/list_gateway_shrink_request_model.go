// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListGatewayShrinkRequest
	GetAcceptLanguage() *string
	SetDescSort(v bool) *ListGatewayShrinkRequest
	GetDescSort() *bool
	SetFilterParamsShrink(v string) *ListGatewayShrinkRequest
	GetFilterParamsShrink() *string
	SetOrderItem(v string) *ListGatewayShrinkRequest
	GetOrderItem() *string
	SetPageNumber(v int32) *ListGatewayShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGatewayShrinkRequest
	GetPageSize() *int32
}

type ListGatewayShrinkRequest struct {
	// The language of the returned information. Valid values:
	//
	// - zh: Chinese
	//
	// - en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The sorting order. This parameter is currently not used.
	//
	// example:
	//
	// false
	DescSort *bool `json:"DescSort,omitempty" xml:"DescSort,omitempty"`
	// The filter parameters.
	FilterParamsShrink *string `json:"FilterParams,omitempty" xml:"FilterParams,omitempty"`
	// The order information.
	//
	// example:
	//
	// {}
	OrderItem *string `json:"OrderItem,omitempty" xml:"OrderItem,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of instances per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListGatewayShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListGatewayShrinkRequest) GetDescSort() *bool {
	return s.DescSort
}

func (s *ListGatewayShrinkRequest) GetFilterParamsShrink() *string {
	return s.FilterParamsShrink
}

func (s *ListGatewayShrinkRequest) GetOrderItem() *string {
	return s.OrderItem
}

func (s *ListGatewayShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayShrinkRequest) SetAcceptLanguage(v string) *ListGatewayShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListGatewayShrinkRequest) SetDescSort(v bool) *ListGatewayShrinkRequest {
	s.DescSort = &v
	return s
}

func (s *ListGatewayShrinkRequest) SetFilterParamsShrink(v string) *ListGatewayShrinkRequest {
	s.FilterParamsShrink = &v
	return s
}

func (s *ListGatewayShrinkRequest) SetOrderItem(v string) *ListGatewayShrinkRequest {
	s.OrderItem = &v
	return s
}

func (s *ListGatewayShrinkRequest) SetPageNumber(v int32) *ListGatewayShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayShrinkRequest) SetPageSize(v int32) *ListGatewayShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListGatewayShrinkRequest) Validate() error {
	return dara.Validate(s)
}
