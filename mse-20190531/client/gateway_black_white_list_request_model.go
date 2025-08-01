// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGatewayBlackWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GatewayBlackWhiteListRequest
	GetAcceptLanguage() *string
	SetDescSort(v bool) *GatewayBlackWhiteListRequest
	GetDescSort() *bool
	SetFilterParams(v *GatewayBlackWhiteListRequestFilterParams) *GatewayBlackWhiteListRequest
	GetFilterParams() *GatewayBlackWhiteListRequestFilterParams
	SetOrderItem(v string) *GatewayBlackWhiteListRequest
	GetOrderItem() *string
	SetPageNumber(v int32) *GatewayBlackWhiteListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GatewayBlackWhiteListRequest
	GetPageSize() *int32
}

type GatewayBlackWhiteListRequest struct {
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
	FilterParams *GatewayBlackWhiteListRequestFilterParams `json:"FilterParams,omitempty" xml:"FilterParams,omitempty" type:"Struct"`
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

func (s GatewayBlackWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s GatewayBlackWhiteListRequest) GoString() string {
	return s.String()
}

func (s *GatewayBlackWhiteListRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GatewayBlackWhiteListRequest) GetDescSort() *bool {
	return s.DescSort
}

func (s *GatewayBlackWhiteListRequest) GetFilterParams() *GatewayBlackWhiteListRequestFilterParams {
	return s.FilterParams
}

func (s *GatewayBlackWhiteListRequest) GetOrderItem() *string {
	return s.OrderItem
}

func (s *GatewayBlackWhiteListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GatewayBlackWhiteListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GatewayBlackWhiteListRequest) SetAcceptLanguage(v string) *GatewayBlackWhiteListRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GatewayBlackWhiteListRequest) SetDescSort(v bool) *GatewayBlackWhiteListRequest {
	s.DescSort = &v
	return s
}

func (s *GatewayBlackWhiteListRequest) SetFilterParams(v *GatewayBlackWhiteListRequestFilterParams) *GatewayBlackWhiteListRequest {
	s.FilterParams = v
	return s
}

func (s *GatewayBlackWhiteListRequest) SetOrderItem(v string) *GatewayBlackWhiteListRequest {
	s.OrderItem = &v
	return s
}

func (s *GatewayBlackWhiteListRequest) SetPageNumber(v int32) *GatewayBlackWhiteListRequest {
	s.PageNumber = &v
	return s
}

func (s *GatewayBlackWhiteListRequest) SetPageSize(v int32) *GatewayBlackWhiteListRequest {
	s.PageSize = &v
	return s
}

func (s *GatewayBlackWhiteListRequest) Validate() error {
	return dara.Validate(s)
}

type GatewayBlackWhiteListRequestFilterParams struct {
	// The gateway ID.
	//
	// example:
	//
	// 81
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway. If this parameter is used together with the GatewayId parameter, the value of the GatewayId parameter is used.
	//
	// example:
	//
	// gw-5017305290e14centbrveca****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// This parameter is unavailable for public use.
	//
	// example:
	//
	// ""
	IsWhite *bool `json:"IsWhite,omitempty" xml:"IsWhite,omitempty"`
	// This parameter is unavailable for public use.
	//
	// example:
	//
	// ""
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The content that you want to query.
	//
	// example:
	//
	// 1.1.1.1
	SearchContent *string `json:"SearchContent,omitempty" xml:"SearchContent,omitempty"`
	// The query type. Valid values:
	//
	// 	- ROUTE: The list is queried by route. If the value of this parameter is ROUTE, set the SearchContent parameter to the route name.
	//
	// 	- DOMAIN: The list is queried by domain name. If the value of this parameter is DOMAIN, set the SearchContent parameter to the domain name.
	//
	// 	- IP: The list is queried by specified IP address. If the value of this parameter is IP, set the SearchContent parameter to the IP address.
	//
	// example:
	//
	// IP
	SearchType *string `json:"SearchType,omitempty" xml:"SearchType,omitempty"`
	// This parameter is unavailable for public use.
	//
	// example:
	//
	// ""
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GatewayBlackWhiteListRequestFilterParams) String() string {
	return dara.Prettify(s)
}

func (s GatewayBlackWhiteListRequestFilterParams) GoString() string {
	return s.String()
}

func (s *GatewayBlackWhiteListRequestFilterParams) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *GatewayBlackWhiteListRequestFilterParams) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GatewayBlackWhiteListRequestFilterParams) GetIsWhite() *bool {
	return s.IsWhite
}

func (s *GatewayBlackWhiteListRequestFilterParams) GetResourceType() *string {
	return s.ResourceType
}

func (s *GatewayBlackWhiteListRequestFilterParams) GetSearchContent() *string {
	return s.SearchContent
}

func (s *GatewayBlackWhiteListRequestFilterParams) GetSearchType() *string {
	return s.SearchType
}

func (s *GatewayBlackWhiteListRequestFilterParams) GetType() *string {
	return s.Type
}

func (s *GatewayBlackWhiteListRequestFilterParams) SetGatewayId(v int64) *GatewayBlackWhiteListRequestFilterParams {
	s.GatewayId = &v
	return s
}

func (s *GatewayBlackWhiteListRequestFilterParams) SetGatewayUniqueId(v string) *GatewayBlackWhiteListRequestFilterParams {
	s.GatewayUniqueId = &v
	return s
}

func (s *GatewayBlackWhiteListRequestFilterParams) SetIsWhite(v bool) *GatewayBlackWhiteListRequestFilterParams {
	s.IsWhite = &v
	return s
}

func (s *GatewayBlackWhiteListRequestFilterParams) SetResourceType(v string) *GatewayBlackWhiteListRequestFilterParams {
	s.ResourceType = &v
	return s
}

func (s *GatewayBlackWhiteListRequestFilterParams) SetSearchContent(v string) *GatewayBlackWhiteListRequestFilterParams {
	s.SearchContent = &v
	return s
}

func (s *GatewayBlackWhiteListRequestFilterParams) SetSearchType(v string) *GatewayBlackWhiteListRequestFilterParams {
	s.SearchType = &v
	return s
}

func (s *GatewayBlackWhiteListRequestFilterParams) SetType(v string) *GatewayBlackWhiteListRequestFilterParams {
	s.Type = &v
	return s
}

func (s *GatewayBlackWhiteListRequestFilterParams) Validate() error {
	return dara.Validate(s)
}
