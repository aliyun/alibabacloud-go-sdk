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
	// 	- **false**: disables automatic payment. This is the default value. If you select this option, you must go to [Order Center](https://usercenter2-intl.aliyun.com/order/list) to complete the payment after an order is generated.
	//
	// 	- **true**: enables automatic payment. Payments are automatically completed.
	//
	// >  This parameter takes effect only if you call the UpdateBandwidthPackage operation to upgrade a bandwidth plan.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to use coupons. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The bandwidth value of the bandwidth plan. Unit: Mbit/s.
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
	// The type of bandwidth. Valid values:
	//
	// 	- **Basic**
	//
	// 	- **Enhanced**
	//
	// 	- **Advanced**
	//
	// >  You can upgrade **Basic*	- bandwidth to **Enhanced*	- bandwidth or downgrade Enhanced bandwidth to Basic bandwidth. You cannot change **Advanced*	- bandwidth to another type of bandwidth.
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
	// The name of the bandwidth plan. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
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
