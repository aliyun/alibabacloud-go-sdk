// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListGatewayAuthRequest
	GetAcceptLanguage() *string
	SetDescSort(v bool) *ListGatewayAuthRequest
	GetDescSort() *bool
	SetFilterParams(v *ListGatewayAuthRequestFilterParams) *ListGatewayAuthRequest
	GetFilterParams() *ListGatewayAuthRequestFilterParams
	SetOrderItem(v string) *ListGatewayAuthRequest
	GetOrderItem() *string
	SetPageNumber(v int32) *ListGatewayAuthRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGatewayAuthRequest
	GetPageSize() *int32
}

type ListGatewayAuthRequest struct {
	// The language of the response. Valid values: zh (default): Chinese, en: English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The sort order.
	//
	// example:
	//
	// false
	DescSort *bool `json:"DescSort,omitempty" xml:"DescSort,omitempty"`
	// The filter conditions. Refer to the examples.
	FilterParams *ListGatewayAuthRequestFilterParams `json:"FilterParams,omitempty" xml:"FilterParams,omitempty" type:"Struct"`
	// This parameter is not available.
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
	// The number of items per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListGatewayAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListGatewayAuthRequest) GetDescSort() *bool {
	return s.DescSort
}

func (s *ListGatewayAuthRequest) GetFilterParams() *ListGatewayAuthRequestFilterParams {
	return s.FilterParams
}

func (s *ListGatewayAuthRequest) GetOrderItem() *string {
	return s.OrderItem
}

func (s *ListGatewayAuthRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayAuthRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayAuthRequest) SetAcceptLanguage(v string) *ListGatewayAuthRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListGatewayAuthRequest) SetDescSort(v bool) *ListGatewayAuthRequest {
	s.DescSort = &v
	return s
}

func (s *ListGatewayAuthRequest) SetFilterParams(v *ListGatewayAuthRequestFilterParams) *ListGatewayAuthRequest {
	s.FilterParams = v
	return s
}

func (s *ListGatewayAuthRequest) SetOrderItem(v string) *ListGatewayAuthRequest {
	s.OrderItem = &v
	return s
}

func (s *ListGatewayAuthRequest) SetPageNumber(v int32) *ListGatewayAuthRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayAuthRequest) SetPageSize(v int32) *ListGatewayAuthRequest {
	s.PageSize = &v
	return s
}

func (s *ListGatewayAuthRequest) Validate() error {
	if s.FilterParams != nil {
		if err := s.FilterParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayAuthRequestFilterParams struct {
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-5017305290e14centbrveca****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The whitelist mode.
	//
	// example:
	//
	// true
	IsWhite *bool `json:"IsWhite,omitempty" xml:"IsWhite,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// rutain-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The enabling status.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type.
	//
	// example:
	//
	// JWT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListGatewayAuthRequestFilterParams) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthRequestFilterParams) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthRequestFilterParams) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListGatewayAuthRequestFilterParams) GetIsWhite() *bool {
	return s.IsWhite
}

func (s *ListGatewayAuthRequestFilterParams) GetName() *string {
	return s.Name
}

func (s *ListGatewayAuthRequestFilterParams) GetStatus() *bool {
	return s.Status
}

func (s *ListGatewayAuthRequestFilterParams) GetType() *string {
	return s.Type
}

func (s *ListGatewayAuthRequestFilterParams) SetGatewayUniqueId(v string) *ListGatewayAuthRequestFilterParams {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListGatewayAuthRequestFilterParams) SetIsWhite(v bool) *ListGatewayAuthRequestFilterParams {
	s.IsWhite = &v
	return s
}

func (s *ListGatewayAuthRequestFilterParams) SetName(v string) *ListGatewayAuthRequestFilterParams {
	s.Name = &v
	return s
}

func (s *ListGatewayAuthRequestFilterParams) SetStatus(v bool) *ListGatewayAuthRequestFilterParams {
	s.Status = &v
	return s
}

func (s *ListGatewayAuthRequestFilterParams) SetType(v string) *ListGatewayAuthRequestFilterParams {
	s.Type = &v
	return s
}

func (s *ListGatewayAuthRequestFilterParams) Validate() error {
	return dara.Validate(s)
}
