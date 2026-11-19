// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourcePackagePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v int32) *GetResourcePackagePriceRequest
	GetDuration() *int32
	SetEffectiveDate(v string) *GetResourcePackagePriceRequest
	GetEffectiveDate() *string
	SetInstanceId(v string) *GetResourcePackagePriceRequest
	GetInstanceId() *string
	SetOrderType(v string) *GetResourcePackagePriceRequest
	GetOrderType() *string
	SetOwnerId(v int64) *GetResourcePackagePriceRequest
	GetOwnerId() *int64
	SetPackageType(v string) *GetResourcePackagePriceRequest
	GetPackageType() *string
	SetPricingCycle(v string) *GetResourcePackagePriceRequest
	GetPricingCycle() *string
	SetProductCode(v string) *GetResourcePackagePriceRequest
	GetProductCode() *string
	SetSpecification(v string) *GetResourcePackagePriceRequest
	GetSpecification() *string
}

type GetResourcePackagePriceRequest struct {
	// The validity period of the resource plan. The value must match the available validity period of the specification size.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The effective period of the resource plan. If you do not specify this parameter, the default is immediate.
	//
	// When the order type is **BUY**, you cannot query the price of a resource plan whose effective period is more than six months later than the current time.
	//
	// When the order type is **UPGRADE**, the effective period must be earlier than or equal to the actual expiration time of the resource plan to be upgraded.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2020-02-10T12:00:00Z
	EffectiveDate *string `json:"EffectiveDate,omitempty" xml:"EffectiveDate,omitempty"`
	// The instance ID. **This parameter is required when the order type is renewal or upgrade.**
	//
	// example:
	//
	// OSSBAG-cn-0xl0002
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The order type. Valid values:
	//
	// - BUY: new purchase.
	//
	// - UPGRADE: upgrade.
	//
	// - RENEW: renewal.
	//
	// Default value: BUY.
	//
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource plan type. The value must be the same as the Code value of **ResourcePackage*	- returned by the operation that describes resource plans.
	//
	// This parameter is required.
	//
	// example:
	//
	// FPT_ossbag_periodMonthlyAcc_NetworkOut_finance_common
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The unit of the validity period for the resource plan. Valid values:
	//
	// - Month: months.
	//
	// - Year: years.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The product code. You can obtain the product code by calling the **QueryProductList*	- operation or referring to the **Alibaba Cloud product code quick reference guide**.
	//
	// This parameter is required.
	//
	// example:
	//
	// ossbag
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The specification size of the resource plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// 500
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
}

func (s GetResourcePackagePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePackagePriceRequest) GoString() string {
	return s.String()
}

func (s *GetResourcePackagePriceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *GetResourcePackagePriceRequest) GetEffectiveDate() *string {
	return s.EffectiveDate
}

func (s *GetResourcePackagePriceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetResourcePackagePriceRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *GetResourcePackagePriceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetResourcePackagePriceRequest) GetPackageType() *string {
	return s.PackageType
}

func (s *GetResourcePackagePriceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *GetResourcePackagePriceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetResourcePackagePriceRequest) GetSpecification() *string {
	return s.Specification
}

func (s *GetResourcePackagePriceRequest) SetDuration(v int32) *GetResourcePackagePriceRequest {
	s.Duration = &v
	return s
}

func (s *GetResourcePackagePriceRequest) SetEffectiveDate(v string) *GetResourcePackagePriceRequest {
	s.EffectiveDate = &v
	return s
}

func (s *GetResourcePackagePriceRequest) SetInstanceId(v string) *GetResourcePackagePriceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetResourcePackagePriceRequest) SetOrderType(v string) *GetResourcePackagePriceRequest {
	s.OrderType = &v
	return s
}

func (s *GetResourcePackagePriceRequest) SetOwnerId(v int64) *GetResourcePackagePriceRequest {
	s.OwnerId = &v
	return s
}

func (s *GetResourcePackagePriceRequest) SetPackageType(v string) *GetResourcePackagePriceRequest {
	s.PackageType = &v
	return s
}

func (s *GetResourcePackagePriceRequest) SetPricingCycle(v string) *GetResourcePackagePriceRequest {
	s.PricingCycle = &v
	return s
}

func (s *GetResourcePackagePriceRequest) SetProductCode(v string) *GetResourcePackagePriceRequest {
	s.ProductCode = &v
	return s
}

func (s *GetResourcePackagePriceRequest) SetSpecification(v string) *GetResourcePackagePriceRequest {
	s.Specification = &v
	return s
}

func (s *GetResourcePackagePriceRequest) Validate() error {
	return dara.Validate(s)
}
