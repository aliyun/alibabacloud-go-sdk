// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAcceleratorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateAcceleratorRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateAcceleratorRequest
	GetAutoRenew() *bool
	SetAutoRenewDuration(v int32) *CreateAcceleratorRequest
	GetAutoRenewDuration() *int32
	SetAutoUseCoupon(v string) *CreateAcceleratorRequest
	GetAutoUseCoupon() *string
	SetBandwidth(v int32) *CreateAcceleratorRequest
	GetBandwidth() *int32
	SetBandwidthBillingType(v string) *CreateAcceleratorRequest
	GetBandwidthBillingType() *string
	SetClientToken(v string) *CreateAcceleratorRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateAcceleratorRequest
	GetDryRun() *bool
	SetDuration(v int32) *CreateAcceleratorRequest
	GetDuration() *int32
	SetInstanceChargeType(v string) *CreateAcceleratorRequest
	GetInstanceChargeType() *string
	SetIpSetConfig(v *CreateAcceleratorRequestIpSetConfig) *CreateAcceleratorRequest
	GetIpSetConfig() *CreateAcceleratorRequestIpSetConfig
	SetName(v string) *CreateAcceleratorRequest
	GetName() *string
	SetPricingCycle(v string) *CreateAcceleratorRequest
	GetPricingCycle() *string
	SetPromotionOptionNo(v string) *CreateAcceleratorRequest
	GetPromotionOptionNo() *string
	SetRegionId(v string) *CreateAcceleratorRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateAcceleratorRequest
	GetResourceGroupId() *string
	SetSpec(v string) *CreateAcceleratorRequest
	GetSpec() *string
	SetTag(v []*CreateAcceleratorRequestTag) *CreateAcceleratorRequest
	GetTag() []*CreateAcceleratorRequestTag
}

type CreateAcceleratorRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// - **false*	- (default): Disables automatic payment. After an order is generated, go to the Order Hub to complete the payment.
	//
	// - **true**: Enables automatic payment. The system automatically completes the payment.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false*	- (default): No.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal duration. Unit: month.
	//
	// Valid values: **1*	- to **12**. Default value: **1**.
	//
	// > This parameter is valid only when **AutoRenew*	- is set to **true**.
	//
	// example:
	//
	// 1
	AutoRenewDuration *int32 `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	// Specifies whether to automatically apply coupons to your bills. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false*	- (default): No.
	//
	// > This parameter is valid only when **AutoPay*	- is set to **true**.
	//
	// example:
	//
	// false
	AutoUseCoupon *string `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The bandwidth of the standard GA instance. Unit: **Mbps**.
	//
	// Valid values: 200 to 5000.
	//
	// > This parameter is required and valid only when the access mode of the acceleration area is Anycast.
	//
	// example:
	//
	// 200
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The bandwidth billing method.
	//
	// - **BandwidthPackage**: Billed by bandwidth plan.
	//
	// - **CDT**: Billed by data transfer.
	//
	// example:
	//
	// BandwidthPackage
	BandwidthBillingType *string `json:"BandwidthBillingType,omitempty" xml:"BandwidthBillingType,omitempty"`
	// A client token that is used to ensure the idempotence of the request.
	//
	// Generate a parameter value from your client to make sure that the value is unique among different requests. The ClientToken parameter can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the request as the **ClientToken**. The **RequestId*	- may be different for each request.
	//
	// example:
	//
	// 123e4567****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: Performs a dry run. The system checks the required parameters, request format, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): Sends a normal request. If the request passes the check, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The subscription duration.
	//
	// - If **PricingCycle*	- is set to **Month**, the valid values for **Duration*	- are **1*	- to **9**.
	//
	// - If **PricingCycle*	- is set to **Year**, the valid values for **Duration*	- are **1*	- to **3**.
	//
	// > This parameter is required if **InstanceChargeType*	- (the billing method of the GA instance) is set to **PREPAY*	- (subscription).
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The billing method of the GA instance.
	//
	// - PREPAY (default): subscription.
	//
	// - POSTPAY: pay-as-you-go.
	//
	// example:
	//
	// PREPAY
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The configurations of the acceleration area.
	IpSetConfig *CreateAcceleratorRequestIpSetConfig `json:"IpSetConfig,omitempty" xml:"IpSetConfig,omitempty" type:"Struct"`
	// The name of the GA instance.
	//
	// The name must be 1 to 128 characters in length, start with a letter, and can contain digits, underscores (_), and hyphens (-).
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The billing cycle. Valid values:
	//
	// - **Month**: Billed by month.
	//
	// - **Year**: Billed by year.
	//
	// > This parameter is required if **InstanceChargeType*	- (the billing method of the GA instance) is set to **PREPAY*	- (subscription).
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The coupon code.
	//
	// > This parameter is available only on the Alibaba Cloud International Website (www\\.alibabacloud.com).
	//
	// example:
	//
	// 50003298014****
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	// Deprecated
	//
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the standard GA instance belongs.
	//
	// example:
	//
	// rg-aekzrnd67gq****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The instance type of the GA instance. Valid values:
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
	// - **100**: Super Large I
	//
	// - **200**: Super Large II
	//
	// > 	- Currently, the Large III and higher instance types are available only to users on the whitelist. To use these instance types, contact your account manager.
	//
	// >
	//
	// > 	- This parameter is required if **InstanceChargeType*	- (the billing method of the GA instance) is set to **PREPAY*	- (subscription).
	//
	// The definitions of different instance types are different. For more information, see [Instance types](https://help.aliyun.com/document_detail/153127.html).
	//
	// example:
	//
	// 1
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The tags of the GA instance.
	Tag []*CreateAcceleratorRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateAcceleratorRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *CreateAcceleratorRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateAcceleratorRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateAcceleratorRequest) GetAutoRenewDuration() *int32 {
	return s.AutoRenewDuration
}

func (s *CreateAcceleratorRequest) GetAutoUseCoupon() *string {
	return s.AutoUseCoupon
}

func (s *CreateAcceleratorRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateAcceleratorRequest) GetBandwidthBillingType() *string {
	return s.BandwidthBillingType
}

func (s *CreateAcceleratorRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAcceleratorRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateAcceleratorRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateAcceleratorRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateAcceleratorRequest) GetIpSetConfig() *CreateAcceleratorRequestIpSetConfig {
	return s.IpSetConfig
}

func (s *CreateAcceleratorRequest) GetName() *string {
	return s.Name
}

func (s *CreateAcceleratorRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateAcceleratorRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *CreateAcceleratorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAcceleratorRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateAcceleratorRequest) GetSpec() *string {
	return s.Spec
}

func (s *CreateAcceleratorRequest) GetTag() []*CreateAcceleratorRequestTag {
	return s.Tag
}

func (s *CreateAcceleratorRequest) SetAutoPay(v bool) *CreateAcceleratorRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateAcceleratorRequest) SetAutoRenew(v bool) *CreateAcceleratorRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateAcceleratorRequest) SetAutoRenewDuration(v int32) *CreateAcceleratorRequest {
	s.AutoRenewDuration = &v
	return s
}

func (s *CreateAcceleratorRequest) SetAutoUseCoupon(v string) *CreateAcceleratorRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateAcceleratorRequest) SetBandwidth(v int32) *CreateAcceleratorRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateAcceleratorRequest) SetBandwidthBillingType(v string) *CreateAcceleratorRequest {
	s.BandwidthBillingType = &v
	return s
}

func (s *CreateAcceleratorRequest) SetClientToken(v string) *CreateAcceleratorRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAcceleratorRequest) SetDryRun(v bool) *CreateAcceleratorRequest {
	s.DryRun = &v
	return s
}

func (s *CreateAcceleratorRequest) SetDuration(v int32) *CreateAcceleratorRequest {
	s.Duration = &v
	return s
}

func (s *CreateAcceleratorRequest) SetInstanceChargeType(v string) *CreateAcceleratorRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateAcceleratorRequest) SetIpSetConfig(v *CreateAcceleratorRequestIpSetConfig) *CreateAcceleratorRequest {
	s.IpSetConfig = v
	return s
}

func (s *CreateAcceleratorRequest) SetName(v string) *CreateAcceleratorRequest {
	s.Name = &v
	return s
}

func (s *CreateAcceleratorRequest) SetPricingCycle(v string) *CreateAcceleratorRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateAcceleratorRequest) SetPromotionOptionNo(v string) *CreateAcceleratorRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *CreateAcceleratorRequest) SetRegionId(v string) *CreateAcceleratorRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAcceleratorRequest) SetResourceGroupId(v string) *CreateAcceleratorRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAcceleratorRequest) SetSpec(v string) *CreateAcceleratorRequest {
	s.Spec = &v
	return s
}

func (s *CreateAcceleratorRequest) SetTag(v []*CreateAcceleratorRequestTag) *CreateAcceleratorRequest {
	s.Tag = v
	return s
}

func (s *CreateAcceleratorRequest) Validate() error {
	if s.IpSetConfig != nil {
		if err := s.IpSetConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAcceleratorRequestIpSetConfig struct {
	// The access mode of the acceleration area. Valid values:
	//
	// - **UserDefine**: Custom nearby access mode. Select acceleration areas and regions as needed. Global Accelerator provides a separate elastic IP address (EIP) for each acceleration region.
	//
	// - **Anycast**: Automatic nearby access mode. You do not need to configure an acceleration area. Global Accelerator provides an Anycast EIP for multiple regions. Users connect to the nearest access point of the Alibaba Cloud network using the Anycast EIP.
	//
	// example:
	//
	// UserDefine
	AccessMode *string `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
}

func (s CreateAcceleratorRequestIpSetConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAcceleratorRequestIpSetConfig) GoString() string {
	return s.String()
}

func (s *CreateAcceleratorRequestIpSetConfig) GetAccessMode() *string {
	return s.AccessMode
}

func (s *CreateAcceleratorRequestIpSetConfig) SetAccessMode(v string) *CreateAcceleratorRequestIpSetConfig {
	s.AccessMode = &v
	return s
}

func (s *CreateAcceleratorRequestIpSetConfig) Validate() error {
	return dara.Validate(s)
}

type CreateAcceleratorRequestTag struct {
	// The tag key of the GA instance. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// You can specify up to 20 tag keys.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the GA instance. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// You can specify up to 20 tag values.
	//
	// example:
	//
	// tag-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAcceleratorRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateAcceleratorRequestTag) GoString() string {
	return s.String()
}

func (s *CreateAcceleratorRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateAcceleratorRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateAcceleratorRequestTag) SetKey(v string) *CreateAcceleratorRequestTag {
	s.Key = &v
	return s
}

func (s *CreateAcceleratorRequestTag) SetValue(v string) *CreateAcceleratorRequestTag {
	s.Value = &v
	return s
}

func (s *CreateAcceleratorRequestTag) Validate() error {
	return dara.Validate(s)
}
