// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *ListGatewayShrinkRequest
	GetChargeType() *string
	SetGatewayId(v string) *ListGatewayShrinkRequest
	GetGatewayId() *string
	SetGatewayName(v string) *ListGatewayShrinkRequest
	GetGatewayName() *string
	SetGatewayType(v string) *ListGatewayShrinkRequest
	GetGatewayType() *string
	SetInternetEnabled(v bool) *ListGatewayShrinkRequest
	GetInternetEnabled() *bool
	SetLabelShrink(v string) *ListGatewayShrinkRequest
	GetLabelShrink() *string
	SetOrder(v string) *ListGatewayShrinkRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListGatewayShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGatewayShrinkRequest
	GetPageSize() *int32
	SetResourceName(v string) *ListGatewayShrinkRequest
	GetResourceName() *string
	SetSort(v string) *ListGatewayShrinkRequest
	GetSort() *string
	SetStatus(v string) *ListGatewayShrinkRequest
	GetStatus() *string
}

type ListGatewayShrinkRequest struct {
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
	LabelShrink *string `json:"Label,omitempty" xml:"Label,omitempty"`
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

func (s ListGatewayShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayShrinkRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListGatewayShrinkRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListGatewayShrinkRequest) GetGatewayName() *string {
	return s.GatewayName
}

func (s *ListGatewayShrinkRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *ListGatewayShrinkRequest) GetInternetEnabled() *bool {
	return s.InternetEnabled
}

func (s *ListGatewayShrinkRequest) GetLabelShrink() *string {
	return s.LabelShrink
}

func (s *ListGatewayShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListGatewayShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayShrinkRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListGatewayShrinkRequest) GetSort() *string {
	return s.Sort
}

func (s *ListGatewayShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListGatewayShrinkRequest) SetChargeType(v string) *ListGatewayShrinkRequest {
	s.ChargeType = &v
	return s
}

func (s *ListGatewayShrinkRequest) SetGatewayId(v string) *ListGatewayShrinkRequest {
	s.GatewayId = &v
	return s
}

func (s *ListGatewayShrinkRequest) SetGatewayName(v string) *ListGatewayShrinkRequest {
	s.GatewayName = &v
	return s
}

func (s *ListGatewayShrinkRequest) SetGatewayType(v string) *ListGatewayShrinkRequest {
	s.GatewayType = &v
	return s
}

func (s *ListGatewayShrinkRequest) SetInternetEnabled(v bool) *ListGatewayShrinkRequest {
	s.InternetEnabled = &v
	return s
}

func (s *ListGatewayShrinkRequest) SetLabelShrink(v string) *ListGatewayShrinkRequest {
	s.LabelShrink = &v
	return s
}

func (s *ListGatewayShrinkRequest) SetOrder(v string) *ListGatewayShrinkRequest {
	s.Order = &v
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

func (s *ListGatewayShrinkRequest) SetResourceName(v string) *ListGatewayShrinkRequest {
	s.ResourceName = &v
	return s
}

func (s *ListGatewayShrinkRequest) SetSort(v string) *ListGatewayShrinkRequest {
	s.Sort = &v
	return s
}

func (s *ListGatewayShrinkRequest) SetStatus(v string) *ListGatewayShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListGatewayShrinkRequest) Validate() error {
	return dara.Validate(s)
}
