// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBandwidthPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *UpdateBandwidthPackageRequest
	GetAutoPay() *bool
	SetAutoUseCoupon(v bool) *UpdateBandwidthPackageRequest
	GetAutoUseCoupon() *bool
	SetBandwidth(v int32) *UpdateBandwidthPackageRequest
	GetBandwidth() *int32
	SetBandwidthPackageId(v string) *UpdateBandwidthPackageRequest
	GetBandwidthPackageId() *string
	SetBandwidthType(v string) *UpdateBandwidthPackageRequest
	GetBandwidthType() *string
	SetDescription(v string) *UpdateBandwidthPackageRequest
	GetDescription() *string
	SetName(v string) *UpdateBandwidthPackageRequest
	GetName() *string
	SetRegionId(v string) *UpdateBandwidthPackageRequest
	GetRegionId() *string
}

type UpdateBandwidthPackageRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// - **false*	- (default): disables automatic payment. After an order is generated, you must go to the <props="china">[Order Hub](https://usercenter2.aliyun.com/order/list)<props="intl">[Order Hub](https://usercenter2-intl.aliyun.com/order/list) to complete the payment.
	//
	// - **true**: enables automatic payment. The system automatically pays the bill.
	//
	// > This parameter is required only for upgrade orders.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to use coupons. Valid values:
	//
	// - **true**: yes.
	//
	// - **false*	- (default): no.
	//
	// example:
	//
	// false
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The bandwidth of the bandwidth plan. Unit: Mbit/s.
	//
	// Valid values: **2*	- to **2000**.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The ID of the bandwidth plan that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// gbwp-bp1sgzldyj6b4q7cx****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The bandwidth type. Valid values:
	//
	// - **Basic**: standard bandwidth.
	//
	// - **Enhanced**: enhanced bandwidth.
	//
	// - **Advanced**: premium bandwidth.
	//
	// > You can upgrade a Basic bandwidth plan to Enhanced, or downgrade an **Enhanced*	- bandwidth plan to **Basic**. You cannot change the bandwidth type of an **Advanced*	- bandwidth plan.
	//
	// example:
	//
	// Basic
	BandwidthType *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	// The description of the bandwidth plan.
	//
	// The description can be up to 256 characters in length.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the bandwidth plan.
	//
	// The name must be 1 to 128 characters in length, start with a letter or a Chinese character, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region where the Global Accelerator instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateBandwidthPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *UpdateBandwidthPackageRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *UpdateBandwidthPackageRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *UpdateBandwidthPackageRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *UpdateBandwidthPackageRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *UpdateBandwidthPackageRequest) GetBandwidthType() *string {
	return s.BandwidthType
}

func (s *UpdateBandwidthPackageRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateBandwidthPackageRequest) GetName() *string {
	return s.Name
}

func (s *UpdateBandwidthPackageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateBandwidthPackageRequest) SetAutoPay(v bool) *UpdateBandwidthPackageRequest {
	s.AutoPay = &v
	return s
}

func (s *UpdateBandwidthPackageRequest) SetAutoUseCoupon(v bool) *UpdateBandwidthPackageRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *UpdateBandwidthPackageRequest) SetBandwidth(v int32) *UpdateBandwidthPackageRequest {
	s.Bandwidth = &v
	return s
}

func (s *UpdateBandwidthPackageRequest) SetBandwidthPackageId(v string) *UpdateBandwidthPackageRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *UpdateBandwidthPackageRequest) SetBandwidthType(v string) *UpdateBandwidthPackageRequest {
	s.BandwidthType = &v
	return s
}

func (s *UpdateBandwidthPackageRequest) SetDescription(v string) *UpdateBandwidthPackageRequest {
	s.Description = &v
	return s
}

func (s *UpdateBandwidthPackageRequest) SetName(v string) *UpdateBandwidthPackageRequest {
	s.Name = &v
	return s
}

func (s *UpdateBandwidthPackageRequest) SetRegionId(v string) *UpdateBandwidthPackageRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateBandwidthPackageRequest) Validate() error {
	return dara.Validate(s)
}
