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
	// 	- **false**: disables automatic payment. This is the default value. After an order is generated, you must go to the [Order Center](https://usercenter2-intl.aliyun.com/order/list) to complete the payment.
	//
	// 	- **true**: enables automatic payment. Payments are automatically completed.
	//
	// >  This parameter takes effect only if you call the operation to upgrade a GA instance.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to automatically pay bills by using coupons. Default value: false. Valid values:
	//
	// 	- **true**: automatically pays bills by using coupons.
	//
	// 	- **false**: does not automatically pay bills by using coupons.
	//
	// >  This parameter takes effect only if the **AutoPay*	- parameter is set to **true**.
	//
	// example:
	//
	// false
	AutoUseCoupon *bool  `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	Bandwidth     *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
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
	// The name must be 1 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// Accelerator
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Deprecated
	//
	// The region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The specification of the GA instance. Valid values:
	//
	// 	- **1**: Small Ⅰ
	//
	// 	- **2**: Small Ⅱ
	//
	// 	- **3**: Small Ⅲ
	//
	// 	- **5**: Medium Ⅰ
	//
	// 	- **8**: Medium Ⅱ
	//
	// 	- **10**: Medium Ⅲ
	//
	// 	- **20**: Large Ⅰ
	//
	// 	- **30**: Large Ⅱ
	//
	// 	- **40**: Large Ⅲ
	//
	// 	- **50**: Large Ⅳ
	//
	// 	- **60**: Large Ⅴ
	//
	// 	- **70**: Large Ⅵ
	//
	// 	- **80**: Large VⅡ
	//
	// 	- **90**: Large VⅢ
	//
	// 	- **100**: Super Large Ⅰ
	//
	// 	- **200**: Super Large Ⅱ
	//
	// >  The Large Ⅲ specification and higher specifications are available only for accounts that are added to the whitelist. To use these specifications, contact your Alibaba Cloud account manager.
	//
	// Different specifications provide different capabilities. For more information, see [Instance specifications](https://help.aliyun.com/document_detail/153127.html).
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
