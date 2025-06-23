// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateMultiOrderRequest struct {
	OrderItems []*CreateMultiOrderRequestOrderItems `json:"OrderItems,omitempty" xml:"OrderItems,omitempty" type:"Repeated"`
	// example:
	//
	// create
	OrderType        *string            `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	Properties       map[string]*string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	ResellerOwnerUid *int64             `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
}

func (s CreateMultiOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMultiOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateMultiOrderRequest) SetOrderItems(v []*CreateMultiOrderRequestOrderItems) *CreateMultiOrderRequest {
	s.OrderItems = v
	return s
}

func (s *CreateMultiOrderRequest) SetOrderType(v string) *CreateMultiOrderRequest {
	s.OrderType = &v
	return s
}

func (s *CreateMultiOrderRequest) SetProperties(v map[string]*string) *CreateMultiOrderRequest {
	s.Properties = v
	return s
}

func (s *CreateMultiOrderRequest) SetResellerOwnerUid(v int64) *CreateMultiOrderRequest {
	s.ResellerOwnerUid = &v
	return s
}

type CreateMultiOrderRequestOrderItems struct {
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// example:
	//
	// false
	AutoRenew  *bool                                          `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	BuyChange  *bool                                          `json:"BuyChange,omitempty" xml:"BuyChange,omitempty"`
	Components []*CreateMultiOrderRequestOrderItemsComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// Year
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	PromotionId *string   `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// DurationPackage
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s CreateMultiOrderRequestOrderItems) String() string {
	return tea.Prettify(s)
}

func (s CreateMultiOrderRequestOrderItems) GoString() string {
	return s.String()
}

func (s *CreateMultiOrderRequestOrderItems) SetAmount(v int32) *CreateMultiOrderRequestOrderItems {
	s.Amount = &v
	return s
}

func (s *CreateMultiOrderRequestOrderItems) SetAutoPay(v bool) *CreateMultiOrderRequestOrderItems {
	s.AutoPay = &v
	return s
}

func (s *CreateMultiOrderRequestOrderItems) SetAutoRenew(v bool) *CreateMultiOrderRequestOrderItems {
	s.AutoRenew = &v
	return s
}

func (s *CreateMultiOrderRequestOrderItems) SetBuyChange(v bool) *CreateMultiOrderRequestOrderItems {
	s.BuyChange = &v
	return s
}

func (s *CreateMultiOrderRequestOrderItems) SetComponents(v []*CreateMultiOrderRequestOrderItemsComponents) *CreateMultiOrderRequestOrderItems {
	s.Components = v
	return s
}

func (s *CreateMultiOrderRequestOrderItems) SetPeriod(v int32) *CreateMultiOrderRequestOrderItems {
	s.Period = &v
	return s
}

func (s *CreateMultiOrderRequestOrderItems) SetPeriodUnit(v string) *CreateMultiOrderRequestOrderItems {
	s.PeriodUnit = &v
	return s
}

func (s *CreateMultiOrderRequestOrderItems) SetPromotionId(v string) *CreateMultiOrderRequestOrderItems {
	s.PromotionId = &v
	return s
}

func (s *CreateMultiOrderRequestOrderItems) SetResourceIds(v []*string) *CreateMultiOrderRequestOrderItems {
	s.ResourceIds = v
	return s
}

func (s *CreateMultiOrderRequestOrderItems) SetResourceType(v string) *CreateMultiOrderRequestOrderItems {
	s.ResourceType = &v
	return s
}

type CreateMultiOrderRequestOrderItemsComponents struct {
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// cn-shanghai
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateMultiOrderRequestOrderItemsComponents) String() string {
	return tea.Prettify(s)
}

func (s CreateMultiOrderRequestOrderItemsComponents) GoString() string {
	return s.String()
}

func (s *CreateMultiOrderRequestOrderItemsComponents) SetKey(v string) *CreateMultiOrderRequestOrderItemsComponents {
	s.Key = &v
	return s
}

func (s *CreateMultiOrderRequestOrderItemsComponents) SetValue(v string) *CreateMultiOrderRequestOrderItemsComponents {
	s.Value = &v
	return s
}

type CreateMultiOrderShrinkRequest struct {
	OrderItems []*CreateMultiOrderShrinkRequestOrderItems `json:"OrderItems,omitempty" xml:"OrderItems,omitempty" type:"Repeated"`
	// example:
	//
	// create
	OrderType        *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	PropertiesShrink *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	ResellerOwnerUid *int64  `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
}

func (s CreateMultiOrderShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMultiOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMultiOrderShrinkRequest) SetOrderItems(v []*CreateMultiOrderShrinkRequestOrderItems) *CreateMultiOrderShrinkRequest {
	s.OrderItems = v
	return s
}

func (s *CreateMultiOrderShrinkRequest) SetOrderType(v string) *CreateMultiOrderShrinkRequest {
	s.OrderType = &v
	return s
}

func (s *CreateMultiOrderShrinkRequest) SetPropertiesShrink(v string) *CreateMultiOrderShrinkRequest {
	s.PropertiesShrink = &v
	return s
}

func (s *CreateMultiOrderShrinkRequest) SetResellerOwnerUid(v int64) *CreateMultiOrderShrinkRequest {
	s.ResellerOwnerUid = &v
	return s
}

type CreateMultiOrderShrinkRequestOrderItems struct {
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// example:
	//
	// false
	AutoRenew  *bool                                                `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	BuyChange  *bool                                                `json:"BuyChange,omitempty" xml:"BuyChange,omitempty"`
	Components []*CreateMultiOrderShrinkRequestOrderItemsComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// Year
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	PromotionId *string   `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// DurationPackage
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s CreateMultiOrderShrinkRequestOrderItems) String() string {
	return tea.Prettify(s)
}

func (s CreateMultiOrderShrinkRequestOrderItems) GoString() string {
	return s.String()
}

func (s *CreateMultiOrderShrinkRequestOrderItems) SetAmount(v int32) *CreateMultiOrderShrinkRequestOrderItems {
	s.Amount = &v
	return s
}

func (s *CreateMultiOrderShrinkRequestOrderItems) SetAutoPay(v bool) *CreateMultiOrderShrinkRequestOrderItems {
	s.AutoPay = &v
	return s
}

func (s *CreateMultiOrderShrinkRequestOrderItems) SetAutoRenew(v bool) *CreateMultiOrderShrinkRequestOrderItems {
	s.AutoRenew = &v
	return s
}

func (s *CreateMultiOrderShrinkRequestOrderItems) SetBuyChange(v bool) *CreateMultiOrderShrinkRequestOrderItems {
	s.BuyChange = &v
	return s
}

func (s *CreateMultiOrderShrinkRequestOrderItems) SetComponents(v []*CreateMultiOrderShrinkRequestOrderItemsComponents) *CreateMultiOrderShrinkRequestOrderItems {
	s.Components = v
	return s
}

func (s *CreateMultiOrderShrinkRequestOrderItems) SetPeriod(v int32) *CreateMultiOrderShrinkRequestOrderItems {
	s.Period = &v
	return s
}

func (s *CreateMultiOrderShrinkRequestOrderItems) SetPeriodUnit(v string) *CreateMultiOrderShrinkRequestOrderItems {
	s.PeriodUnit = &v
	return s
}

func (s *CreateMultiOrderShrinkRequestOrderItems) SetPromotionId(v string) *CreateMultiOrderShrinkRequestOrderItems {
	s.PromotionId = &v
	return s
}

func (s *CreateMultiOrderShrinkRequestOrderItems) SetResourceIds(v []*string) *CreateMultiOrderShrinkRequestOrderItems {
	s.ResourceIds = v
	return s
}

func (s *CreateMultiOrderShrinkRequestOrderItems) SetResourceType(v string) *CreateMultiOrderShrinkRequestOrderItems {
	s.ResourceType = &v
	return s
}

type CreateMultiOrderShrinkRequestOrderItemsComponents struct {
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// cn-shanghai
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateMultiOrderShrinkRequestOrderItemsComponents) String() string {
	return tea.Prettify(s)
}

func (s CreateMultiOrderShrinkRequestOrderItemsComponents) GoString() string {
	return s.String()
}

func (s *CreateMultiOrderShrinkRequestOrderItemsComponents) SetKey(v string) *CreateMultiOrderShrinkRequestOrderItemsComponents {
	s.Key = &v
	return s
}

func (s *CreateMultiOrderShrinkRequestOrderItemsComponents) SetValue(v string) *CreateMultiOrderShrinkRequestOrderItemsComponents {
	s.Value = &v
	return s
}

type CreateMultiOrderResponseBody struct {
	OrderIds []*int64 `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Repeated"`
	// example:
	//
	// 833C4D2C-09C7-5CE6-8159-06758B964970
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMultiOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMultiOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMultiOrderResponseBody) SetOrderIds(v []*int64) *CreateMultiOrderResponseBody {
	s.OrderIds = v
	return s
}

func (s *CreateMultiOrderResponseBody) SetRequestId(v string) *CreateMultiOrderResponseBody {
	s.RequestId = &v
	return s
}

type CreateMultiOrderResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMultiOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMultiOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMultiOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateMultiOrderResponse) SetHeaders(v map[string]*string) *CreateMultiOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateMultiOrderResponse) SetStatusCode(v int32) *CreateMultiOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMultiOrderResponse) SetBody(v *CreateMultiOrderResponseBody) *CreateMultiOrderResponse {
	s.Body = v
	return s
}

type DescribeDeliveryAddressResponseBody struct {
	Addresses []*DescribeDeliveryAddressResponseBodyAddresses `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// example:
	//
	// 72481C12-69AB-5CE1-8A35-A8EFA921****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDeliveryAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeliveryAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryAddressResponseBody) SetAddresses(v []*DescribeDeliveryAddressResponseBodyAddresses) *DescribeDeliveryAddressResponseBody {
	s.Addresses = v
	return s
}

func (s *DescribeDeliveryAddressResponseBody) SetRequestId(v string) *DescribeDeliveryAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBody) SetTotalCount(v int32) *DescribeDeliveryAddressResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDeliveryAddressResponseBodyAddresses struct {
	Area *DescribeDeliveryAddressResponseBodyAddressesArea `json:"Area,omitempty" xml:"Area,omitempty" type:"Struct"`
	City *DescribeDeliveryAddressResponseBodyAddressesCity `json:"City,omitempty" xml:"City,omitempty" type:"Struct"`
	// example:
	//
	// Alice
	Contacts *string `json:"Contacts,omitempty" xml:"Contacts,omitempty"`
	// example:
	//
	// true
	DefaultAddress *bool   `json:"DefaultAddress,omitempty" xml:"DefaultAddress,omitempty"`
	Detail         *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// example:
	//
	// 1381111****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// 03****
	PostalCode *string                                               `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	Province   *DescribeDeliveryAddressResponseBodyAddressesProvince `json:"Province,omitempty" xml:"Province,omitempty" type:"Struct"`
	Town       *DescribeDeliveryAddressResponseBodyAddressesTown     `json:"Town,omitempty" xml:"Town,omitempty" type:"Struct"`
}

func (s DescribeDeliveryAddressResponseBodyAddresses) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeliveryAddressResponseBodyAddresses) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetArea(v *DescribeDeliveryAddressResponseBodyAddressesArea) *DescribeDeliveryAddressResponseBodyAddresses {
	s.Area = v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetCity(v *DescribeDeliveryAddressResponseBodyAddressesCity) *DescribeDeliveryAddressResponseBodyAddresses {
	s.City = v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetContacts(v string) *DescribeDeliveryAddressResponseBodyAddresses {
	s.Contacts = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetDefaultAddress(v bool) *DescribeDeliveryAddressResponseBodyAddresses {
	s.DefaultAddress = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetDetail(v string) *DescribeDeliveryAddressResponseBodyAddresses {
	s.Detail = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetMobile(v string) *DescribeDeliveryAddressResponseBodyAddresses {
	s.Mobile = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetPostalCode(v string) *DescribeDeliveryAddressResponseBodyAddresses {
	s.PostalCode = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetProvince(v *DescribeDeliveryAddressResponseBodyAddressesProvince) *DescribeDeliveryAddressResponseBodyAddresses {
	s.Province = v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetTown(v *DescribeDeliveryAddressResponseBodyAddressesTown) *DescribeDeliveryAddressResponseBodyAddresses {
	s.Town = v
	return s
}

type DescribeDeliveryAddressResponseBodyAddressesArea struct {
	// example:
	//
	// 33****
	AreaId   *int64  `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	AreaName *string `json:"AreaName,omitempty" xml:"AreaName,omitempty"`
}

func (s DescribeDeliveryAddressResponseBodyAddressesArea) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeliveryAddressResponseBodyAddressesArea) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryAddressResponseBodyAddressesArea) SetAreaId(v int64) *DescribeDeliveryAddressResponseBodyAddressesArea {
	s.AreaId = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddressesArea) SetAreaName(v string) *DescribeDeliveryAddressResponseBodyAddressesArea {
	s.AreaName = &v
	return s
}

type DescribeDeliveryAddressResponseBodyAddressesCity struct {
	// example:
	//
	// 33****
	CityId   *int64  `json:"CityId,omitempty" xml:"CityId,omitempty"`
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
}

func (s DescribeDeliveryAddressResponseBodyAddressesCity) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeliveryAddressResponseBodyAddressesCity) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryAddressResponseBodyAddressesCity) SetCityId(v int64) *DescribeDeliveryAddressResponseBodyAddressesCity {
	s.CityId = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddressesCity) SetCityName(v string) *DescribeDeliveryAddressResponseBodyAddressesCity {
	s.CityName = &v
	return s
}

type DescribeDeliveryAddressResponseBodyAddressesProvince struct {
	// example:
	//
	// 330000
	ProvinceId   *int64  `json:"ProvinceId,omitempty" xml:"ProvinceId,omitempty"`
	ProvinceName *string `json:"ProvinceName,omitempty" xml:"ProvinceName,omitempty"`
}

func (s DescribeDeliveryAddressResponseBodyAddressesProvince) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeliveryAddressResponseBodyAddressesProvince) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryAddressResponseBodyAddressesProvince) SetProvinceId(v int64) *DescribeDeliveryAddressResponseBodyAddressesProvince {
	s.ProvinceId = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddressesProvince) SetProvinceName(v string) *DescribeDeliveryAddressResponseBodyAddressesProvince {
	s.ProvinceName = &v
	return s
}

type DescribeDeliveryAddressResponseBodyAddressesTown struct {
	// example:
	//
	// 3001****
	TownId   *int64  `json:"TownId,omitempty" xml:"TownId,omitempty"`
	TownName *string `json:"TownName,omitempty" xml:"TownName,omitempty"`
}

func (s DescribeDeliveryAddressResponseBodyAddressesTown) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeliveryAddressResponseBodyAddressesTown) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryAddressResponseBodyAddressesTown) SetTownId(v int64) *DescribeDeliveryAddressResponseBodyAddressesTown {
	s.TownId = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddressesTown) SetTownName(v string) *DescribeDeliveryAddressResponseBodyAddressesTown {
	s.TownName = &v
	return s
}

type DescribeDeliveryAddressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeliveryAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeliveryAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeliveryAddressResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryAddressResponse) SetHeaders(v map[string]*string) *DescribeDeliveryAddressResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeliveryAddressResponse) SetStatusCode(v int32) *DescribeDeliveryAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeliveryAddressResponse) SetBody(v *DescribeDeliveryAddressResponseBody) *DescribeDeliveryAddressResponse {
	s.Body = v
	return s
}

type DescribeMultiPriceRequest struct {
	OrderItems []*DescribeMultiPriceRequestOrderItems `json:"OrderItems,omitempty" xml:"OrderItems,omitempty" type:"Repeated"`
	// example:
	//
	// create
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// pacakge
	PackageCode *string `json:"PackageCode,omitempty" xml:"PackageCode,omitempty"`
	// example:
	//
	// 182864463481****
	ResellerOwnerUid *int64 `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
}

func (s DescribeMultiPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceRequest) SetOrderItems(v []*DescribeMultiPriceRequestOrderItems) *DescribeMultiPriceRequest {
	s.OrderItems = v
	return s
}

func (s *DescribeMultiPriceRequest) SetOrderType(v string) *DescribeMultiPriceRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeMultiPriceRequest) SetPackageCode(v string) *DescribeMultiPriceRequest {
	s.PackageCode = &v
	return s
}

func (s *DescribeMultiPriceRequest) SetResellerOwnerUid(v int64) *DescribeMultiPriceRequest {
	s.ResellerOwnerUid = &v
	return s
}

type DescribeMultiPriceRequestOrderItems struct {
	// example:
	//
	// 1
	Amount      *int32                                           `json:"Amount,omitempty" xml:"Amount,omitempty"`
	Components  []*DescribeMultiPriceRequestOrderItemsComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	InstanceIds []*string                                        `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// Year
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	PromotionId *string   `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// example:
	//
	// DurationPackage
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeMultiPriceRequestOrderItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiPriceRequestOrderItems) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceRequestOrderItems) SetAmount(v int32) *DescribeMultiPriceRequestOrderItems {
	s.Amount = &v
	return s
}

func (s *DescribeMultiPriceRequestOrderItems) SetComponents(v []*DescribeMultiPriceRequestOrderItemsComponents) *DescribeMultiPriceRequestOrderItems {
	s.Components = v
	return s
}

func (s *DescribeMultiPriceRequestOrderItems) SetInstanceIds(v []*string) *DescribeMultiPriceRequestOrderItems {
	s.InstanceIds = v
	return s
}

func (s *DescribeMultiPriceRequestOrderItems) SetPeriod(v int32) *DescribeMultiPriceRequestOrderItems {
	s.Period = &v
	return s
}

func (s *DescribeMultiPriceRequestOrderItems) SetPeriodUnit(v string) *DescribeMultiPriceRequestOrderItems {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeMultiPriceRequestOrderItems) SetPromotionId(v string) *DescribeMultiPriceRequestOrderItems {
	s.PromotionId = &v
	return s
}

func (s *DescribeMultiPriceRequestOrderItems) SetResourceIds(v []*string) *DescribeMultiPriceRequestOrderItems {
	s.ResourceIds = v
	return s
}

func (s *DescribeMultiPriceRequestOrderItems) SetResourceType(v string) *DescribeMultiPriceRequestOrderItems {
	s.ResourceType = &v
	return s
}

type DescribeMultiPriceRequestOrderItemsComponents struct {
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// cn-shanghai
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeMultiPriceRequestOrderItemsComponents) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiPriceRequestOrderItemsComponents) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceRequestOrderItemsComponents) SetKey(v string) *DescribeMultiPriceRequestOrderItemsComponents {
	s.Key = &v
	return s
}

func (s *DescribeMultiPriceRequestOrderItemsComponents) SetValue(v string) *DescribeMultiPriceRequestOrderItemsComponents {
	s.Value = &v
	return s
}

type DescribeMultiPriceResponseBody struct {
	PriceInfo *DescribeMultiPriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// example:
	//
	// 833C4D2C-09C7-5CE6-8159-06758B964970
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMultiPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceResponseBody) SetPriceInfo(v *DescribeMultiPriceResponseBodyPriceInfo) *DescribeMultiPriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *DescribeMultiPriceResponseBody) SetRequestId(v string) *DescribeMultiPriceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMultiPriceResponseBodyPriceInfo struct {
	Price *DescribeMultiPriceResponseBodyPriceInfoPrice   `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	Rules []*DescribeMultiPriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeMultiPriceResponseBodyPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiPriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceResponseBodyPriceInfo) SetPrice(v *DescribeMultiPriceResponseBodyPriceInfoPrice) *DescribeMultiPriceResponseBodyPriceInfo {
	s.Price = v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfo) SetRules(v []*DescribeMultiPriceResponseBodyPriceInfoRules) *DescribeMultiPriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

type DescribeMultiPriceResponseBodyPriceInfoPrice struct {
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 534.6
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// example:
	//
	// 6800
	OriginalPrice            *float32                                                    `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	PriceDetails             []*DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails `json:"PriceDetails,omitempty" xml:"PriceDetails,omitempty" type:"Repeated"`
	Promotions               []*DescribeMultiPriceResponseBodyPriceInfoPricePromotions   `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	RefundInstanceIdPriceMap map[string]*float32                                         `json:"RefundInstanceIdPriceMap,omitempty" xml:"RefundInstanceIdPriceMap,omitempty"`
	// example:
	//
	// 60.00
	RefundPrice *float32 `json:"RefundPrice,omitempty" xml:"RefundPrice,omitempty"`
	// example:
	//
	// 82.6
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeMultiPriceResponseBodyPriceInfoPrice) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiPriceResponseBodyPriceInfoPrice) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPrice) SetCurrency(v string) *DescribeMultiPriceResponseBodyPriceInfoPrice {
	s.Currency = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPrice) SetDiscountPrice(v float32) *DescribeMultiPriceResponseBodyPriceInfoPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPrice) SetOriginalPrice(v float32) *DescribeMultiPriceResponseBodyPriceInfoPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPrice) SetPriceDetails(v []*DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails) *DescribeMultiPriceResponseBodyPriceInfoPrice {
	s.PriceDetails = v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPrice) SetPromotions(v []*DescribeMultiPriceResponseBodyPriceInfoPricePromotions) *DescribeMultiPriceResponseBodyPriceInfoPrice {
	s.Promotions = v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPrice) SetRefundInstanceIdPriceMap(v map[string]*float32) *DescribeMultiPriceResponseBodyPriceInfoPrice {
	s.RefundInstanceIdPriceMap = v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPrice) SetRefundPrice(v float32) *DescribeMultiPriceResponseBodyPriceInfoPrice {
	s.RefundPrice = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPrice) SetTradePrice(v float32) *DescribeMultiPriceResponseBodyPriceInfoPrice {
	s.TradePrice = &v
	return s
}

type DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails struct {
	ModuleDetails []*DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails `json:"ModuleDetails,omitempty" xml:"ModuleDetails,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	OrderItem   *int32                                                               `json:"OrderItem,omitempty" xml:"OrderItem,omitempty"`
	PriceDetail *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail `json:"PriceDetail,omitempty" xml:"PriceDetail,omitempty" type:"Struct"`
}

func (s DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails) SetModuleDetails(v []*DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails {
	s.ModuleDetails = v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails) SetOrderItem(v int32) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails {
	s.OrderItem = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails) SetPriceDetail(v *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails {
	s.PriceDetail = v
	return s
}

type DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails struct {
	// example:
	//
	// 734.65
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// example:
	//
	// DesktopType
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// example:
	//
	// eds.enterprise_office.8c32g
	ModuleValue *string `json:"ModuleValue,omitempty" xml:"ModuleValue,omitempty"`
	// example:
	//
	// 10900
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// example:
	//
	// 292.2
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails) SetDiscountPrice(v float32) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails) SetModuleCode(v string) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails {
	s.ModuleCode = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails) SetModuleName(v string) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails {
	s.ModuleName = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails) SetModuleValue(v string) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails {
	s.ModuleValue = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails) SetOriginalPrice(v float32) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails) SetTradePrice(v float32) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails {
	s.TradePrice = &v
	return s
}

type DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail struct {
	// example:
	//
	// 20.00
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// example:
	//
	// 100.00
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// example:
	//
	// DurationPackage
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 80.00
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail) SetDiscountPrice(v float32) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail) SetOriginalPrice(v float32) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail) SetResourceType(v string) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail {
	s.ResourceType = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail) SetTradePrice(v float32) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail {
	s.TradePrice = &v
	return s
}

type DescribeMultiPriceResponseBodyPriceInfoPricePromotions struct {
	// example:
	//
	// new
	OptionCode    *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	PromotionId   *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// example:
	//
	// true
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s DescribeMultiPriceResponseBodyPriceInfoPricePromotions) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiPriceResponseBodyPriceInfoPricePromotions) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePromotions) SetOptionCode(v string) *DescribeMultiPriceResponseBodyPriceInfoPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePromotions) SetPromotionDesc(v string) *DescribeMultiPriceResponseBodyPriceInfoPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePromotions) SetPromotionId(v string) *DescribeMultiPriceResponseBodyPriceInfoPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePromotions) SetPromotionName(v string) *DescribeMultiPriceResponseBodyPriceInfoPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePromotions) SetSelected(v bool) *DescribeMultiPriceResponseBodyPriceInfoPricePromotions {
	s.Selected = &v
	return s
}

type DescribeMultiPriceResponseBodyPriceInfoRules struct {
	// example:
	//
	// accounts_suspect_users
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// hrzdvc
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeMultiPriceResponseBodyPriceInfoRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiPriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceResponseBodyPriceInfoRules) SetDescription(v string) *DescribeMultiPriceResponseBodyPriceInfoRules {
	s.Description = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoRules) SetRuleId(v int64) *DescribeMultiPriceResponseBodyPriceInfoRules {
	s.RuleId = &v
	return s
}

type DescribeMultiPriceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMultiPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMultiPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiPriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceResponse) SetHeaders(v map[string]*string) *DescribeMultiPriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeMultiPriceResponse) SetStatusCode(v int32) *DescribeMultiPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMultiPriceResponse) SetBody(v *DescribeMultiPriceResponseBody) *DescribeMultiPriceResponse {
	s.Body = v
	return s
}

type DescribePackageDeductionsRequest struct {
	EndTime     *int64    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	PackageIds  []*string `json:"PackageIds,omitempty" xml:"PackageIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// CorePackage
	ResourceType  *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
	StartTime     *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePackageDeductionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePackageDeductionsRequest) GoString() string {
	return s.String()
}

func (s *DescribePackageDeductionsRequest) SetEndTime(v int64) *DescribePackageDeductionsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePackageDeductionsRequest) SetInstanceIds(v []*string) *DescribePackageDeductionsRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribePackageDeductionsRequest) SetPackageIds(v []*string) *DescribePackageDeductionsRequest {
	s.PackageIds = v
	return s
}

func (s *DescribePackageDeductionsRequest) SetPageNum(v int32) *DescribePackageDeductionsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribePackageDeductionsRequest) SetPageSize(v int32) *DescribePackageDeductionsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePackageDeductionsRequest) SetResourceType(v string) *DescribePackageDeductionsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribePackageDeductionsRequest) SetResourceTypes(v []*string) *DescribePackageDeductionsRequest {
	s.ResourceTypes = v
	return s
}

func (s *DescribePackageDeductionsRequest) SetStartTime(v int64) *DescribePackageDeductionsRequest {
	s.StartTime = &v
	return s
}

type DescribePackageDeductionsResponseBody struct {
	Deductions []*DescribePackageDeductionsResponseBodyDeductions `json:"Deductions,omitempty" xml:"Deductions,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 833C4D2C-09C7-5CE6-8159-06758B964970
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount        *int64   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalUsedCoreTime *float32 `json:"TotalUsedCoreTime,omitempty" xml:"TotalUsedCoreTime,omitempty"`
	TotalUsedTime     *int64   `json:"TotalUsedTime,omitempty" xml:"TotalUsedTime,omitempty"`
}

func (s DescribePackageDeductionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePackageDeductionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePackageDeductionsResponseBody) SetDeductions(v []*DescribePackageDeductionsResponseBodyDeductions) *DescribePackageDeductionsResponseBody {
	s.Deductions = v
	return s
}

func (s *DescribePackageDeductionsResponseBody) SetPageNum(v int32) *DescribePackageDeductionsResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribePackageDeductionsResponseBody) SetPageSize(v int32) *DescribePackageDeductionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePackageDeductionsResponseBody) SetRequestId(v string) *DescribePackageDeductionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePackageDeductionsResponseBody) SetTotalCount(v int64) *DescribePackageDeductionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePackageDeductionsResponseBody) SetTotalUsedCoreTime(v float32) *DescribePackageDeductionsResponseBody {
	s.TotalUsedCoreTime = &v
	return s
}

func (s *DescribePackageDeductionsResponseBody) SetTotalUsedTime(v int64) *DescribePackageDeductionsResponseBody {
	s.TotalUsedTime = &v
	return s
}

type DescribePackageDeductionsResponseBodyDeductions struct {
	// example:
	//
	// 4
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// ecd-6wye9optu0kag****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// example:
	//
	// DemoComputer
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// example:
	//
	// eds.enterprise_office.4c8g
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// example:
	//
	// 2024-07-31T03:00Z
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Deleted
	InstanceState *string `json:"InstanceState,omitempty" xml:"InstanceState,omitempty"`
	InstanceType  *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 8192
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SessionId    *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 2024-07-31T02:00Z
	StaTime *string `json:"StaTime,omitempty" xml:"StaTime,omitempty"`
	// example:
	//
	// 4.0
	UsedCoreTime *float32 `json:"UsedCoreTime,omitempty" xml:"UsedCoreTime,omitempty"`
	// example:
	//
	// 3600
	UsedTime          *int64 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	UsedTimeWithScale *int64 `json:"UsedTimeWithScale,omitempty" xml:"UsedTimeWithScale,omitempty"`
}

func (s DescribePackageDeductionsResponseBodyDeductions) String() string {
	return tea.Prettify(s)
}

func (s DescribePackageDeductionsResponseBodyDeductions) GoString() string {
	return s.String()
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetCpu(v int32) *DescribePackageDeductionsResponseBodyDeductions {
	s.Cpu = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetDesktopId(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.DesktopId = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetDesktopName(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.DesktopName = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetDesktopType(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.DesktopType = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetEndTime(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.EndTime = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetInstanceId(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.InstanceId = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetInstanceState(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.InstanceState = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetInstanceType(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.InstanceType = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetMemory(v int64) *DescribePackageDeductionsResponseBodyDeductions {
	s.Memory = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetOsType(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.OsType = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetRegionId(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.RegionId = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetResourceType(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.ResourceType = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetSessionId(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.SessionId = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetStaTime(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.StaTime = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetUsedCoreTime(v float32) *DescribePackageDeductionsResponseBodyDeductions {
	s.UsedCoreTime = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetUsedTime(v int64) *DescribePackageDeductionsResponseBodyDeductions {
	s.UsedTime = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetUsedTimeWithScale(v int64) *DescribePackageDeductionsResponseBodyDeductions {
	s.UsedTimeWithScale = &v
	return s
}

type DescribePackageDeductionsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePackageDeductionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePackageDeductionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePackageDeductionsResponse) GoString() string {
	return s.String()
}

func (s *DescribePackageDeductionsResponse) SetHeaders(v map[string]*string) *DescribePackageDeductionsResponse {
	s.Headers = v
	return s
}

func (s *DescribePackageDeductionsResponse) SetStatusCode(v int32) *DescribePackageDeductionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePackageDeductionsResponse) SetBody(v *DescribePackageDeductionsResponseBody) *DescribePackageDeductionsResponse {
	s.Body = v
	return s
}

type ModifyInstancePropertiesRequest struct {
	// example:
	//
	// mdp-0c62ayep0nk4v****
	InstanceId  *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// PackageUsedUpStrategy
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DurationPackage
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// Postpaid
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyInstancePropertiesRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstancePropertiesRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstancePropertiesRequest) SetInstanceId(v string) *ModifyInstancePropertiesRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstancePropertiesRequest) SetInstanceIds(v []*string) *ModifyInstancePropertiesRequest {
	s.InstanceIds = v
	return s
}

func (s *ModifyInstancePropertiesRequest) SetKey(v string) *ModifyInstancePropertiesRequest {
	s.Key = &v
	return s
}

func (s *ModifyInstancePropertiesRequest) SetResourceType(v string) *ModifyInstancePropertiesRequest {
	s.ResourceType = &v
	return s
}

func (s *ModifyInstancePropertiesRequest) SetValue(v string) *ModifyInstancePropertiesRequest {
	s.Value = &v
	return s
}

type ModifyInstancePropertiesResponseBody struct {
	// example:
	//
	// 833C4D2C-09C7-5CE6-8159-06758B964970
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstancePropertiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstancePropertiesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstancePropertiesResponseBody) SetRequestId(v string) *ModifyInstancePropertiesResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstancePropertiesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstancePropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstancePropertiesResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstancePropertiesResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstancePropertiesResponse) SetHeaders(v map[string]*string) *ModifyInstancePropertiesResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstancePropertiesResponse) SetStatusCode(v int32) *ModifyInstancePropertiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstancePropertiesResponse) SetBody(v *ModifyInstancePropertiesResponseBody) *ModifyInstancePropertiesResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("wss"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 多商品组合下单
//
// @param tmpReq - CreateMultiOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMultiOrderResponse
func (client *Client) CreateMultiOrderWithOptions(tmpReq *CreateMultiOrderRequest, runtime *util.RuntimeOptions) (_result *CreateMultiOrderResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateMultiOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Properties)) {
		request.PropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Properties, tea.String("Properties"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderItems)) {
		query["OrderItems"] = request.OrderItems
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["OrderType"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.PropertiesShrink)) {
		query["Properties"] = request.PropertiesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResellerOwnerUid)) {
		query["ResellerOwnerUid"] = request.ResellerOwnerUid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMultiOrder"),
		Version:     tea.String("2021-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMultiOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 多商品组合下单
//
// @param request - CreateMultiOrderRequest
//
// @return CreateMultiOrderResponse
func (client *Client) CreateMultiOrder(request *CreateMultiOrderRequest) (_result *CreateMultiOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMultiOrderResponse{}
	_body, _err := client.CreateMultiOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询物流地址
//
// @param request - DescribeDeliveryAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeliveryAddressResponse
func (client *Client) DescribeDeliveryAddressWithOptions(runtime *util.RuntimeOptions) (_result *DescribeDeliveryAddressResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeliveryAddress"),
		Version:     tea.String("2021-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDeliveryAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询物流地址
//
// @return DescribeDeliveryAddressResponse
func (client *Client) DescribeDeliveryAddress() (_result *DescribeDeliveryAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeliveryAddressResponse{}
	_body, _err := client.DescribeDeliveryAddressWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量询价
//
// @param request - DescribeMultiPriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMultiPriceResponse
func (client *Client) DescribeMultiPriceWithOptions(request *DescribeMultiPriceRequest, runtime *util.RuntimeOptions) (_result *DescribeMultiPriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderItems)) {
		query["OrderItems"] = request.OrderItems
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["OrderType"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.PackageCode)) {
		query["PackageCode"] = request.PackageCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResellerOwnerUid)) {
		query["ResellerOwnerUid"] = request.ResellerOwnerUid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMultiPrice"),
		Version:     tea.String("2021-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMultiPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量询价
//
// @param request - DescribeMultiPriceRequest
//
// @return DescribeMultiPriceResponse
func (client *Client) DescribeMultiPrice(request *DescribeMultiPriceRequest) (_result *DescribeMultiPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMultiPriceResponse{}
	_body, _err := client.DescribeMultiPriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询核时包抵扣明细
//
// @param request - DescribePackageDeductionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePackageDeductionsResponse
func (client *Client) DescribePackageDeductionsWithOptions(request *DescribePackageDeductionsRequest, runtime *util.RuntimeOptions) (_result *DescribePackageDeductionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.PackageIds)) {
		query["PackageIds"] = request.PackageIds
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceTypes)) {
		query["ResourceTypes"] = request.ResourceTypes
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePackageDeductions"),
		Version:     tea.String("2021-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePackageDeductionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询核时包抵扣明细
//
// @param request - DescribePackageDeductionsRequest
//
// @return DescribePackageDeductionsResponse
func (client *Client) DescribePackageDeductions(request *DescribePackageDeductionsRequest) (_result *DescribePackageDeductionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePackageDeductionsResponse{}
	_body, _err := client.DescribePackageDeductionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新实例属性
//
// @param request - ModifyInstancePropertiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstancePropertiesResponse
func (client *Client) ModifyInstancePropertiesWithOptions(request *ModifyInstancePropertiesRequest, runtime *util.RuntimeOptions) (_result *ModifyInstancePropertiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyInstanceProperties"),
		Version:     tea.String("2021-12-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyInstancePropertiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新实例属性
//
// @param request - ModifyInstancePropertiesRequest
//
// @return ModifyInstancePropertiesResponse
func (client *Client) ModifyInstanceProperties(request *ModifyInstancePropertiesRequest) (_result *ModifyInstancePropertiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstancePropertiesResponse{}
	_body, _err := client.ModifyInstancePropertiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
