// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListGatewayRequest
	GetAcceptLanguage() *string
	SetDescSort(v bool) *ListGatewayRequest
	GetDescSort() *bool
	SetFilterParams(v *ListGatewayRequestFilterParams) *ListGatewayRequest
	GetFilterParams() *ListGatewayRequestFilterParams
	SetOrderItem(v string) *ListGatewayRequest
	GetOrderItem() *string
	SetPageNumber(v int32) *ListGatewayRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGatewayRequest
	GetPageSize() *int32
}

type ListGatewayRequest struct {
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
	// Specifies whether to enable the sorting feature. This feature is not available.
	//
	// example:
	//
	// false
	DescSort *bool `json:"DescSort,omitempty" xml:"DescSort,omitempty"`
	// The details of parameters.
	FilterParams *ListGatewayRequestFilterParams `json:"FilterParams,omitempty" xml:"FilterParams,omitempty" type:"Struct"`
	// The order information.
	//
	// example:
	//
	// {}
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

func (s ListGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListGatewayRequest) GetDescSort() *bool {
	return s.DescSort
}

func (s *ListGatewayRequest) GetFilterParams() *ListGatewayRequestFilterParams {
	return s.FilterParams
}

func (s *ListGatewayRequest) GetOrderItem() *string {
	return s.OrderItem
}

func (s *ListGatewayRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayRequest) SetAcceptLanguage(v string) *ListGatewayRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListGatewayRequest) SetDescSort(v bool) *ListGatewayRequest {
	s.DescSort = &v
	return s
}

func (s *ListGatewayRequest) SetFilterParams(v *ListGatewayRequestFilterParams) *ListGatewayRequest {
	s.FilterParams = v
	return s
}

func (s *ListGatewayRequest) SetOrderItem(v string) *ListGatewayRequest {
	s.OrderItem = &v
	return s
}

func (s *ListGatewayRequest) SetPageNumber(v int32) *ListGatewayRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayRequest) SetPageSize(v int32) *ListGatewayRequest {
	s.PageSize = &v
	return s
}

func (s *ListGatewayRequest) Validate() error {
	if s.FilterParams != nil {
		if err := s.FilterParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayRequestFilterParams struct {
	// The type of the gateway.
	//
	// example:
	//
	// Ingress
	GatewayType *string `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-5017305290e14centbrveca****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// mse_ingresspre-cn-****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The tag of the instance.
	//
	// example:
	//
	// [{"key":"tagkey","value":"tagvalue"}]
	MseTag *string `json:"MseTag,omitempty" xml:"MseTag,omitempty"`
	// The name of the gateway.
	//
	// example:
	//
	// rutain-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-7y2uye*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-bp13hhyjntbab7w****
	Vpc *string `json:"Vpc,omitempty" xml:"Vpc,omitempty"`
}

func (s ListGatewayRequestFilterParams) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRequestFilterParams) GoString() string {
	return s.String()
}

func (s *ListGatewayRequestFilterParams) GetGatewayType() *string {
	return s.GatewayType
}

func (s *ListGatewayRequestFilterParams) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListGatewayRequestFilterParams) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListGatewayRequestFilterParams) GetMseTag() *string {
	return s.MseTag
}

func (s *ListGatewayRequestFilterParams) GetName() *string {
	return s.Name
}

func (s *ListGatewayRequestFilterParams) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListGatewayRequestFilterParams) GetVpc() *string {
	return s.Vpc
}

func (s *ListGatewayRequestFilterParams) SetGatewayType(v string) *ListGatewayRequestFilterParams {
	s.GatewayType = &v
	return s
}

func (s *ListGatewayRequestFilterParams) SetGatewayUniqueId(v string) *ListGatewayRequestFilterParams {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListGatewayRequestFilterParams) SetInstanceId(v string) *ListGatewayRequestFilterParams {
	s.InstanceId = &v
	return s
}

func (s *ListGatewayRequestFilterParams) SetMseTag(v string) *ListGatewayRequestFilterParams {
	s.MseTag = &v
	return s
}

func (s *ListGatewayRequestFilterParams) SetName(v string) *ListGatewayRequestFilterParams {
	s.Name = &v
	return s
}

func (s *ListGatewayRequestFilterParams) SetResourceGroupId(v string) *ListGatewayRequestFilterParams {
	s.ResourceGroupId = &v
	return s
}

func (s *ListGatewayRequestFilterParams) SetVpc(v string) *ListGatewayRequestFilterParams {
	s.Vpc = &v
	return s
}

func (s *ListGatewayRequestFilterParams) Validate() error {
	return dara.Validate(s)
}
