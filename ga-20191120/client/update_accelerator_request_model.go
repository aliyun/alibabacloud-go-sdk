// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAcceleratorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *UpdateAcceleratorRequest
	GetAcceleratorId() *string
	SetAutoPay(v bool) *UpdateAcceleratorRequest
	GetAutoPay() *bool
	SetAutoUseCoupon(v bool) *UpdateAcceleratorRequest
	GetAutoUseCoupon() *bool
	SetBandwidth(v int32) *UpdateAcceleratorRequest
	GetBandwidth() *int32
	SetClientToken(v string) *UpdateAcceleratorRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateAcceleratorRequest
	GetDescription() *string
	SetName(v string) *UpdateAcceleratorRequest
	GetName() *string
	SetRegionId(v string) *UpdateAcceleratorRequest
	GetRegionId() *string
	SetSpec(v string) *UpdateAcceleratorRequest
	GetSpec() *string
}

type UpdateAcceleratorRequest struct {
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// Specifies whether to enable automatic payment. Valid values:
	//
	// - **false*	- (default): Disables automatic payment. After an order is generated, you must go to the [Order Hub](https://usercenter2-intl.aliyun.com/order/list) to complete the payment.
	//
	// - **true**: Enables automatic payment. The system automatically pays the bill.
	//
	// > This parameter is valid only for upgrade orders.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to automatically use a coupon to pay for the bill. Valid values:
	//
	// - **true**: Use a coupon.
	//
	// - **false*	- (default): Do not use a coupon.
	//
	// > This parameter is valid only if **AutoPay*	- is set to **true**.
	//
	// example:
	//
	// false
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The bandwidth of the standard GA instance. Unit: **Mbps**.
	//
	// Valid values: 200 to 5000.
	//
	// > This parameter is valid only when the access mode of the acceleration area is Anycast.
	//
	// example:
	//
	// 200
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// Generate a parameter value from your client to make sure that the value is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 123e4567****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the GA instance. The description can be up to 200 characters in length.
	//
	// example:
	//
	// Accelerator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the GA instance.
	//
	// The name must be 1 to 128 characters in length, start with a letter or a Chinese character, and can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// Accelerator
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Deprecated
	//
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The specification of the GA instance. Valid values:
	//
	// - **1**: Small I
	//
	// - **2**: Small II
	//
	// - **3**: Small III
	//
	// - **5**: Medium I
	//
	// - **8**: Medium II
	//
	// - **10**: Medium III
	//
	// - **20**: Large I
	//
	// - **30**: Large II
	//
	// - **40**: Large III
	//
	// - **50**: Large IV
	//
	// - **60**: Large V
	//
	// - **70**: Large VI
	//
	// - **80**: Large VII
	//
	// - **90**: Large VIII
	//
	// - **100**: Ultra-large I
	//
	// - **200**: Ultra-large II
	//
	// > Large III and higher specifications are available only to whitelisted users. To use these specifications, contact your account manager.
	//
	// The definitions of instance types vary. For more information, see [Instance types](https://help.aliyun.com/document_detail/153127.html).
	//
	// example:
	//
	// 1
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s UpdateAcceleratorRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *UpdateAcceleratorRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *UpdateAcceleratorRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *UpdateAcceleratorRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *UpdateAcceleratorRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *UpdateAcceleratorRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAcceleratorRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAcceleratorRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAcceleratorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAcceleratorRequest) GetSpec() *string {
	return s.Spec
}

func (s *UpdateAcceleratorRequest) SetAcceleratorId(v string) *UpdateAcceleratorRequest {
	s.AcceleratorId = &v
	return s
}

func (s *UpdateAcceleratorRequest) SetAutoPay(v bool) *UpdateAcceleratorRequest {
	s.AutoPay = &v
	return s
}

func (s *UpdateAcceleratorRequest) SetAutoUseCoupon(v bool) *UpdateAcceleratorRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *UpdateAcceleratorRequest) SetBandwidth(v int32) *UpdateAcceleratorRequest {
	s.Bandwidth = &v
	return s
}

func (s *UpdateAcceleratorRequest) SetClientToken(v string) *UpdateAcceleratorRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAcceleratorRequest) SetDescription(v string) *UpdateAcceleratorRequest {
	s.Description = &v
	return s
}

func (s *UpdateAcceleratorRequest) SetName(v string) *UpdateAcceleratorRequest {
	s.Name = &v
	return s
}

func (s *UpdateAcceleratorRequest) SetRegionId(v string) *UpdateAcceleratorRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAcceleratorRequest) SetSpec(v string) *UpdateAcceleratorRequest {
	s.Spec = &v
	return s
}

func (s *UpdateAcceleratorRequest) Validate() error {
	return dara.Validate(s)
}
