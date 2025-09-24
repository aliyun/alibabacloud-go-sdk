// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubscriptionPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetSubscriptionPriceRequest
	GetInstanceId() *string
	SetModuleList(v []*GetSubscriptionPriceRequestModuleList) *GetSubscriptionPriceRequest
	GetModuleList() []*GetSubscriptionPriceRequestModuleList
	SetOrderType(v string) *GetSubscriptionPriceRequest
	GetOrderType() *string
	SetOwnerId(v int64) *GetSubscriptionPriceRequest
	GetOwnerId() *int64
	SetProductCode(v string) *GetSubscriptionPriceRequest
	GetProductCode() *string
	SetProductType(v string) *GetSubscriptionPriceRequest
	GetProductType() *string
	SetQuantity(v int32) *GetSubscriptionPriceRequest
	GetQuantity() *int32
	SetRegion(v string) *GetSubscriptionPriceRequest
	GetRegion() *string
	SetServicePeriodQuantity(v int32) *GetSubscriptionPriceRequest
	GetServicePeriodQuantity() *int32
	SetServicePeriodUnit(v string) *GetSubscriptionPriceRequest
	GetServicePeriodUnit() *string
	SetSubscriptionType(v string) *GetSubscriptionPriceRequest
	GetSubscriptionType() *string
}

type GetSubscriptionPriceRequest struct {
	// The ID of the instance for which the price is queried. This parameter is required if you upgrade an instance. You can specify this parameter to obtain the pre-upgrade configurations of the instance.
	//
	// example:
	//
	// i-khkjhxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The information about the pricing module.
	//
	// This parameter is required.
	ModuleList []*GetSubscriptionPriceRequestModuleList `json:"ModuleList,omitempty" xml:"ModuleList,omitempty" type:"Repeated"`
	// The type of the order. Valid values:
	//
	// 	- NewOrder: purchases an instance of an Alibaba Cloud service.
	//
	// 	- Renewal: renews an instance of an Alibaba Cloud service.
	//
	// 	- Upgrade: upgrades an instance of an Alibaba Cloud service.
	//
	// This parameter is required.
	//
	// example:
	//
	// NewOrder
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The code of the service. For more information about the service code, see **Codes of Alibaba Cloud Services**.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The type of the service. Specify the parameter based on the pricing document of the specific service.
	//
	// example:
	//
	// ecs
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The quantity.
	//
	// example:
	//
	// 1
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The service duration.
	//
	// example:
	//
	// 1
	ServicePeriodQuantity *int32 `json:"ServicePeriodQuantity,omitempty" xml:"ServicePeriodQuantity,omitempty"`
	// The unit of the service duration. Valid values:
	//
	// 	- Year
	//
	// 	- Month
	//
	// example:
	//
	// Year
	ServicePeriodUnit *string `json:"ServicePeriodUnit,omitempty" xml:"ServicePeriodUnit,omitempty"`
	// The billing method. Set the value to Subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// Subscription
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s GetSubscriptionPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionPriceRequest) GoString() string {
	return s.String()
}

func (s *GetSubscriptionPriceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSubscriptionPriceRequest) GetModuleList() []*GetSubscriptionPriceRequestModuleList {
	return s.ModuleList
}

func (s *GetSubscriptionPriceRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *GetSubscriptionPriceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetSubscriptionPriceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetSubscriptionPriceRequest) GetProductType() *string {
	return s.ProductType
}

func (s *GetSubscriptionPriceRequest) GetQuantity() *int32 {
	return s.Quantity
}

func (s *GetSubscriptionPriceRequest) GetRegion() *string {
	return s.Region
}

func (s *GetSubscriptionPriceRequest) GetServicePeriodQuantity() *int32 {
	return s.ServicePeriodQuantity
}

func (s *GetSubscriptionPriceRequest) GetServicePeriodUnit() *string {
	return s.ServicePeriodUnit
}

func (s *GetSubscriptionPriceRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *GetSubscriptionPriceRequest) SetInstanceId(v string) *GetSubscriptionPriceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSubscriptionPriceRequest) SetModuleList(v []*GetSubscriptionPriceRequestModuleList) *GetSubscriptionPriceRequest {
	s.ModuleList = v
	return s
}

func (s *GetSubscriptionPriceRequest) SetOrderType(v string) *GetSubscriptionPriceRequest {
	s.OrderType = &v
	return s
}

func (s *GetSubscriptionPriceRequest) SetOwnerId(v int64) *GetSubscriptionPriceRequest {
	s.OwnerId = &v
	return s
}

func (s *GetSubscriptionPriceRequest) SetProductCode(v string) *GetSubscriptionPriceRequest {
	s.ProductCode = &v
	return s
}

func (s *GetSubscriptionPriceRequest) SetProductType(v string) *GetSubscriptionPriceRequest {
	s.ProductType = &v
	return s
}

func (s *GetSubscriptionPriceRequest) SetQuantity(v int32) *GetSubscriptionPriceRequest {
	s.Quantity = &v
	return s
}

func (s *GetSubscriptionPriceRequest) SetRegion(v string) *GetSubscriptionPriceRequest {
	s.Region = &v
	return s
}

func (s *GetSubscriptionPriceRequest) SetServicePeriodQuantity(v int32) *GetSubscriptionPriceRequest {
	s.ServicePeriodQuantity = &v
	return s
}

func (s *GetSubscriptionPriceRequest) SetServicePeriodUnit(v string) *GetSubscriptionPriceRequest {
	s.ServicePeriodUnit = &v
	return s
}

func (s *GetSubscriptionPriceRequest) SetSubscriptionType(v string) *GetSubscriptionPriceRequest {
	s.SubscriptionType = &v
	return s
}

func (s *GetSubscriptionPriceRequest) Validate() error {
	return dara.Validate(s)
}

type GetSubscriptionPriceRequestModuleList struct {
	// The configurations of the Nth pricing module. Valid values of N: 1 to 50. Format: AA:aa,BB:bb. The values of AA and BB are the property IDs of the pricing module. The values of aa and bb are the property values of the pricing module.
	//
	// This parameter is required.
	//
	// example:
	//
	// PackageCode:version_1
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The identifier of the Nth pricing module.
	//
	// This parameter is required.
	//
	// example:
	//
	// PackageCode
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	// The status of the pricing module. This parameter is required only if the order type is Upgrade. Valid values:
	//
	// 	- 1: adds one or more instances.
	//
	// 	- 2: modifies the configurations of an instance. In the upgrade scenario, if the configurations of the pricing module change, you must specify this value for the parameter.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	ModuleStatus *int32 `json:"ModuleStatus,omitempty" xml:"ModuleStatus,omitempty"`
	// The tag of the specified resource. This parameter is required only if you upgrade or modify the configurations of an Alibaba Cloud service. For example, if you want to modify the configurations of a disk, you can use a tag to identify the ID of the disk.
	//
	// example:
	//
	// 213213123
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetSubscriptionPriceRequestModuleList) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionPriceRequestModuleList) GoString() string {
	return s.String()
}

func (s *GetSubscriptionPriceRequestModuleList) GetConfig() *string {
	return s.Config
}

func (s *GetSubscriptionPriceRequestModuleList) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *GetSubscriptionPriceRequestModuleList) GetModuleStatus() *int32 {
	return s.ModuleStatus
}

func (s *GetSubscriptionPriceRequestModuleList) GetTag() *string {
	return s.Tag
}

func (s *GetSubscriptionPriceRequestModuleList) SetConfig(v string) *GetSubscriptionPriceRequestModuleList {
	s.Config = &v
	return s
}

func (s *GetSubscriptionPriceRequestModuleList) SetModuleCode(v string) *GetSubscriptionPriceRequestModuleList {
	s.ModuleCode = &v
	return s
}

func (s *GetSubscriptionPriceRequestModuleList) SetModuleStatus(v int32) *GetSubscriptionPriceRequestModuleList {
	s.ModuleStatus = &v
	return s
}

func (s *GetSubscriptionPriceRequestModuleList) SetTag(v string) *GetSubscriptionPriceRequestModuleList {
	s.Tag = &v
	return s
}

func (s *GetSubscriptionPriceRequestModuleList) Validate() error {
	return dara.Validate(s)
}
