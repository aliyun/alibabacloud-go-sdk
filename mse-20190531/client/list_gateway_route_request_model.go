// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListGatewayRouteRequest
	GetAcceptLanguage() *string
	SetDescSort(v bool) *ListGatewayRouteRequest
	GetDescSort() *bool
	SetFilterParams(v *ListGatewayRouteRequestFilterParams) *ListGatewayRouteRequest
	GetFilterParams() *ListGatewayRouteRequestFilterParams
	SetOrderItem(v string) *ListGatewayRouteRequest
	GetOrderItem() *string
	SetPageNumber(v int32) *ListGatewayRouteRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGatewayRouteRequest
	GetPageSize() *int32
}

type ListGatewayRouteRequest struct {
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
	// Specifies whether to enable sorting. This parameter is unavailable.
	//
	// example:
	//
	// false
	DescSort *bool `json:"DescSort,omitempty" xml:"DescSort,omitempty"`
	// The parameters that specify filter conditions. The parameters are in the format of {"key1":"value1"}.
	FilterParams *ListGatewayRouteRequestFilterParams `json:"FilterParams,omitempty" xml:"FilterParams,omitempty" type:"Struct"`
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

func (s ListGatewayRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListGatewayRouteRequest) GetDescSort() *bool {
	return s.DescSort
}

func (s *ListGatewayRouteRequest) GetFilterParams() *ListGatewayRouteRequestFilterParams {
	return s.FilterParams
}

func (s *ListGatewayRouteRequest) GetOrderItem() *string {
	return s.OrderItem
}

func (s *ListGatewayRouteRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayRouteRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayRouteRequest) SetAcceptLanguage(v string) *ListGatewayRouteRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListGatewayRouteRequest) SetDescSort(v bool) *ListGatewayRouteRequest {
	s.DescSort = &v
	return s
}

func (s *ListGatewayRouteRequest) SetFilterParams(v *ListGatewayRouteRequestFilterParams) *ListGatewayRouteRequest {
	s.FilterParams = v
	return s
}

func (s *ListGatewayRouteRequest) SetOrderItem(v string) *ListGatewayRouteRequest {
	s.OrderItem = &v
	return s
}

func (s *ListGatewayRouteRequest) SetPageNumber(v int32) *ListGatewayRouteRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayRouteRequest) SetPageSize(v int32) *ListGatewayRouteRequest {
	s.PageSize = &v
	return s
}

func (s *ListGatewayRouteRequest) Validate() error {
	return dara.Validate(s)
}

type ListGatewayRouteRequestFilterParams struct {
	// Deprecated
	//
	// The default service ID.
	//
	// example:
	//
	// 1
	DefaultServiceId *int64 `json:"DefaultServiceId,omitempty" xml:"DefaultServiceId,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// 284
	DomainId *int64 `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The associated domain name.
	//
	// example:
	//
	// *.alites.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the gateway.
	//
	// example:
	//
	// 81
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-30a0106924c94bca8712ec4e79fc5acc
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The name of the gateway.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The order.
	//
	// example:
	//
	// 1
	RouteOrder *int32 `json:"RouteOrder,omitempty" xml:"RouteOrder,omitempty"`
	// The status.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListGatewayRouteRequestFilterParams) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteRequestFilterParams) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteRequestFilterParams) GetDefaultServiceId() *int64 {
	return s.DefaultServiceId
}

func (s *ListGatewayRouteRequestFilterParams) GetDomainId() *int64 {
	return s.DomainId
}

func (s *ListGatewayRouteRequestFilterParams) GetDomainName() *string {
	return s.DomainName
}

func (s *ListGatewayRouteRequestFilterParams) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *ListGatewayRouteRequestFilterParams) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListGatewayRouteRequestFilterParams) GetName() *string {
	return s.Name
}

func (s *ListGatewayRouteRequestFilterParams) GetPath() *string {
	return s.Path
}

func (s *ListGatewayRouteRequestFilterParams) GetRouteOrder() *int32 {
	return s.RouteOrder
}

func (s *ListGatewayRouteRequestFilterParams) GetStatus() *int32 {
	return s.Status
}

func (s *ListGatewayRouteRequestFilterParams) SetDefaultServiceId(v int64) *ListGatewayRouteRequestFilterParams {
	s.DefaultServiceId = &v
	return s
}

func (s *ListGatewayRouteRequestFilterParams) SetDomainId(v int64) *ListGatewayRouteRequestFilterParams {
	s.DomainId = &v
	return s
}

func (s *ListGatewayRouteRequestFilterParams) SetDomainName(v string) *ListGatewayRouteRequestFilterParams {
	s.DomainName = &v
	return s
}

func (s *ListGatewayRouteRequestFilterParams) SetGatewayId(v int64) *ListGatewayRouteRequestFilterParams {
	s.GatewayId = &v
	return s
}

func (s *ListGatewayRouteRequestFilterParams) SetGatewayUniqueId(v string) *ListGatewayRouteRequestFilterParams {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListGatewayRouteRequestFilterParams) SetName(v string) *ListGatewayRouteRequestFilterParams {
	s.Name = &v
	return s
}

func (s *ListGatewayRouteRequestFilterParams) SetPath(v string) *ListGatewayRouteRequestFilterParams {
	s.Path = &v
	return s
}

func (s *ListGatewayRouteRequestFilterParams) SetRouteOrder(v int32) *ListGatewayRouteRequestFilterParams {
	s.RouteOrder = &v
	return s
}

func (s *ListGatewayRouteRequestFilterParams) SetStatus(v int32) *ListGatewayRouteRequestFilterParams {
	s.Status = &v
	return s
}

func (s *ListGatewayRouteRequestFilterParams) Validate() error {
	return dara.Validate(s)
}
