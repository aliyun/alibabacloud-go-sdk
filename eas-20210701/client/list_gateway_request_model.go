// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *ListGatewayRequest
	GetChargeType() *string
	SetGatewayId(v string) *ListGatewayRequest
	GetGatewayId() *string
	SetGatewayName(v string) *ListGatewayRequest
	GetGatewayName() *string
	SetGatewayType(v string) *ListGatewayRequest
	GetGatewayType() *string
	SetInternetEnabled(v bool) *ListGatewayRequest
	GetInternetEnabled() *bool
	SetLabel(v map[string]*string) *ListGatewayRequest
	GetLabel() map[string]*string
	SetOrder(v string) *ListGatewayRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListGatewayRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGatewayRequest
	GetPageSize() *int32
	SetResourceName(v string) *ListGatewayRequest
	GetResourceName() *string
	SetSort(v string) *ListGatewayRequest
	GetSort() *string
	SetStatus(v string) *ListGatewayRequest
	GetStatus() *string
}

type ListGatewayRequest struct {
	// The billing method.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The ID of the private gateway. You can obtain the ID from the private_gateway_id field in the response of the ListResources operation.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The alias of the private gateway.
	//
	// example:
	//
	// mygateway1
	GatewayName *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
	// The type of the gateway.
	//
	// example:
	//
	// Application
	GatewayType *string `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	// Specifies whether to enable access over the public network.
	//
	// example:
	//
	// true
	InternetEnabled *bool `json:"InternetEnabled,omitempty" xml:"InternetEnabled,omitempty"`
	// Filter by tag.
	Label map[string]*string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The sort order.
	//
	// example:
	//
	// asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number of the gateway list to return. The default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of gateways to return on each page. The default value is 100.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group. You can obtain the ID from the ResourceId field in the response of the [ListResources](https://help.aliyun.com/document_detail/412133.html) operation.
	//
	// example:
	//
	// eas-r-4gt8twzwllfo******
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The field to sort by.
	//
	// example:
	//
	// CreateTime
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The status of the gateway.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListGatewayRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListGatewayRequest) GetGatewayName() *string {
	return s.GatewayName
}

func (s *ListGatewayRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *ListGatewayRequest) GetInternetEnabled() *bool {
	return s.InternetEnabled
}

func (s *ListGatewayRequest) GetLabel() map[string]*string {
	return s.Label
}

func (s *ListGatewayRequest) GetOrder() *string {
	return s.Order
}

func (s *ListGatewayRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListGatewayRequest) GetSort() *string {
	return s.Sort
}

func (s *ListGatewayRequest) GetStatus() *string {
	return s.Status
}

func (s *ListGatewayRequest) SetChargeType(v string) *ListGatewayRequest {
	s.ChargeType = &v
	return s
}

func (s *ListGatewayRequest) SetGatewayId(v string) *ListGatewayRequest {
	s.GatewayId = &v
	return s
}

func (s *ListGatewayRequest) SetGatewayName(v string) *ListGatewayRequest {
	s.GatewayName = &v
	return s
}

func (s *ListGatewayRequest) SetGatewayType(v string) *ListGatewayRequest {
	s.GatewayType = &v
	return s
}

func (s *ListGatewayRequest) SetInternetEnabled(v bool) *ListGatewayRequest {
	s.InternetEnabled = &v
	return s
}

func (s *ListGatewayRequest) SetLabel(v map[string]*string) *ListGatewayRequest {
	s.Label = v
	return s
}

func (s *ListGatewayRequest) SetOrder(v string) *ListGatewayRequest {
	s.Order = &v
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

func (s *ListGatewayRequest) SetResourceName(v string) *ListGatewayRequest {
	s.ResourceName = &v
	return s
}

func (s *ListGatewayRequest) SetSort(v string) *ListGatewayRequest {
	s.Sort = &v
	return s
}

func (s *ListGatewayRequest) SetStatus(v string) *ListGatewayRequest {
	s.Status = &v
	return s
}

func (s *ListGatewayRequest) Validate() error {
	return dara.Validate(s)
}
