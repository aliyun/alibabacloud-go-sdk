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
	FilterParams *ListGatewayRouteRequestFilterParams `json:"FilterParams,omitempty" xml:"FilterParams,omitempty" type:"Struct"`
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
	if s.FilterParams != nil {
		if err := s.FilterParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayRouteRequestFilterParams struct {
	// Deprecated
	//
	// Default service ID
	//
	// example:
	//
	// 1
	DefaultServiceId *int64 `json:"DefaultServiceId,omitempty" xml:"DefaultServiceId,omitempty"`
	// Domain ID
	//
	// example:
	//
	// 284
	DomainId *int64 `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// Associated domain
	//
	// example:
	//
	// *.alites.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Gateway ID
	//
	// example:
	//
	// 81
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// Gateway unique identifier ID
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-30a0106924c94bca8712ec4e79fc5acc
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// Route name
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Route path
	//
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Order (unused)
	//
	// example:
	//
	// 1
	RouteOrder *int32 `json:"RouteOrder,omitempty" xml:"RouteOrder,omitempty"`
	// Route status:
	//
	// - Unpublished: 0
	//
	// - Publishing: 2
	//
	// - Published: 3
	//
	// - Editing (updates not published): 4
	//
	// - Unpublishing: 5
	//
	// - Unavailable: 6
	//
	// example:
	//
	// 3
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
