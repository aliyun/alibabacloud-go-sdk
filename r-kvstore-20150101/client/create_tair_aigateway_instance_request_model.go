// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTairAIGatewayInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateTairAIGatewayInstanceRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateTairAIGatewayInstanceRequest
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v string) *CreateTairAIGatewayInstanceRequest
	GetAutoRenewPeriod() *string
	SetAutoUseCoupon(v bool) *CreateTairAIGatewayInstanceRequest
	GetAutoUseCoupon() *bool
	SetBusinessInfo(v string) *CreateTairAIGatewayInstanceRequest
	GetBusinessInfo() *string
	SetChargeType(v string) *CreateTairAIGatewayInstanceRequest
	GetChargeType() *string
	SetClientToken(v string) *CreateTairAIGatewayInstanceRequest
	GetClientToken() *string
	SetCouponNo(v string) *CreateTairAIGatewayInstanceRequest
	GetCouponNo() *string
	SetDryRun(v bool) *CreateTairAIGatewayInstanceRequest
	GetDryRun() *bool
	SetGatewayName(v string) *CreateTairAIGatewayInstanceRequest
	GetGatewayName() *string
	SetGatewayType(v string) *CreateTairAIGatewayInstanceRequest
	GetGatewayType() *string
	SetPeriod(v int32) *CreateTairAIGatewayInstanceRequest
	GetPeriod() *int32
	SetPluginParams(v string) *CreateTairAIGatewayInstanceRequest
	GetPluginParams() *string
	SetPluginType(v string) *CreateTairAIGatewayInstanceRequest
	GetPluginType() *string
	SetRegionId(v string) *CreateTairAIGatewayInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateTairAIGatewayInstanceRequest
	GetResourceGroupId() *string
	SetSecurityToken(v string) *CreateTairAIGatewayInstanceRequest
	GetSecurityToken() *string
	SetTag(v []*CreateTairAIGatewayInstanceRequestTag) *CreateTairAIGatewayInstanceRequest
	GetTag() []*CreateTairAIGatewayInstanceRequestTag
	SetVSwitchId(v string) *CreateTairAIGatewayInstanceRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateTairAIGatewayInstanceRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateTairAIGatewayInstanceRequest
	GetZoneId() *string
}

type CreateTairAIGatewayInstanceRequest struct {
	// Only pay-as-you-go is supported. Ignore this parameter. Specifies whether to enable automatic payment. The value is fixed to **true**.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Only pay-as-you-go is supported. Ignore this parameter. Specifies whether to enable auto-renewal. Valid values:
	//
	// 	- **true**: Enable auto-renewal.
	//
	// 	- **false*	- (default): Disable auto-renewal.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Only pay-as-you-go is supported. Ignore this parameter. The auto-renewal epoch. Unit: months. Valid values: **1**, **2**, **3**, **6**, and **12**.
	//
	// > This parameter is required when **AutoRenew*	- is set to **true**.
	//
	// example:
	//
	// 3
	AutoRenewPeriod *string `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// Only pay-as-you-go is supported. Ignore this parameter.
	//
	// Specifies whether to use a voucher. Valid values:
	//
	// 	- **true**: Use a voucher.
	//
	// 	- **false*	- (default): Do not use a voucher.
	//
	// example:
	//
	// false
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The activity ID and business information.
	//
	// example:
	//
	// 000000000
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The billing method. Only pay-as-you-go is supported. Valid values:
	//
	// - **PostPaid**: pay-as-you-go.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token is case-sensitive and can be up to 64 ASCII characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The coupon code. Default value: `default`.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// Specifies whether to perform a dry run before you create the instance. Valid values:
	//
	// 	- **true**: performs a dry run and does not create the instance. The system checks the request parameters, request format, service limits, and available resources. If the request fails the dry run, the corresponding error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): sends a normal request and creates the instance after the request passes the check.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The name of the gateway.
	//
	// example:
	//
	// TestGateway01
	GatewayName *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
	// The specifications of the gateway.
	//
	// example:
	//
	// aigw.std.2
	GatewayType *string `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	// Only pay-as-you-go is supported. Ignore this parameter.
	//
	// The renewal period. Unit: months. Valid values: **1*	- to **9**, **12**, **24**, and **36**.
	//
	// > For example, if an instance expires in three months and you renew it for three months, the instance will expire in six months after the renewal.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated. Do not specify this parameter.
	//
	// example:
	//
	// -
	PluginParams *string `json:"PluginParams,omitempty" xml:"PluginParams,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated. Do not specify this parameter.
	//
	// example:
	//
	// -
	PluginType *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The tags of the instance.
	Tag []*CreateTairAIGatewayInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch in the VPC. You can call the [DescribeVpcs](https://help.aliyun.com/document_detail/35739.html) operation to query the vSwitch ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp1nme44gek34slfc****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing-l
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateTairAIGatewayInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTairAIGatewayInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateTairAIGatewayInstanceRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateTairAIGatewayInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateTairAIGatewayInstanceRequest) GetAutoRenewPeriod() *string {
	return s.AutoRenewPeriod
}

func (s *CreateTairAIGatewayInstanceRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *CreateTairAIGatewayInstanceRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *CreateTairAIGatewayInstanceRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateTairAIGatewayInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTairAIGatewayInstanceRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *CreateTairAIGatewayInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateTairAIGatewayInstanceRequest) GetGatewayName() *string {
	return s.GatewayName
}

func (s *CreateTairAIGatewayInstanceRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *CreateTairAIGatewayInstanceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateTairAIGatewayInstanceRequest) GetPluginParams() *string {
	return s.PluginParams
}

func (s *CreateTairAIGatewayInstanceRequest) GetPluginType() *string {
	return s.PluginType
}

func (s *CreateTairAIGatewayInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTairAIGatewayInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateTairAIGatewayInstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateTairAIGatewayInstanceRequest) GetTag() []*CreateTairAIGatewayInstanceRequestTag {
	return s.Tag
}

func (s *CreateTairAIGatewayInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateTairAIGatewayInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateTairAIGatewayInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateTairAIGatewayInstanceRequest) SetAutoPay(v bool) *CreateTairAIGatewayInstanceRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateTairAIGatewayInstanceRequest) SetAutoRenew(v bool) *CreateTairAIGatewayInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateTairAIGatewayInstanceRequest) SetAutoRenewPeriod(v string) *CreateTairAIGatewayInstanceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateTairAIGatewayInstanceRequest) SetAutoUseCoupon(v bool) *CreateTairAIGatewayInstanceRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateTairAIGatewayInstanceRequest) SetBusinessInfo(v string) *CreateTairAIGatewayInstanceRequest {
	s.BusinessInfo = &v
	return s
}

func (s *CreateTairAIGatewayInstanceRequest) SetChargeType(v string) *CreateTairAIGatewayInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateTairAIGatewayInstanceRequest) SetClientToken(v string) *CreateTairAIGatewayInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTairAIGatewayInstanceRequest) SetCouponNo(v string) *CreateTairAIGatewayInstanceRequest {
	s.CouponNo = &v
	return s
}

func (s *CreateTairAIGatewayInstanceRequest) SetDryRun(v bool) *CreateTairAIGatewayInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTairAIGatewayInstanceRequest) SetGatewayName(v string) *CreateTairAIGatewayInstanceRequest {
	s.GatewayName = &v
	return s
}

func (s *CreateTairAIGatewayInstanceRequest) SetGatewayType(v string) *CreateTairAIGatewayInstanceRequest {
	s.GatewayType = &v
	return s
}

func (s *CreateTairAIGatewayInstanceRequest) SetPeriod(v int32) *CreateTairAIGatewayInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateTairAIGatewayInstanceRequest) SetPluginParams(v string) *CreateTairAIGatewayInstanceRequest {
	s.PluginParams = &v
	return s
}

func (s *CreateTairAIGatewayInstanceRequest) SetPluginType(v string) *CreateTairAIGatewayInstanceRequest {
	s.PluginType = &v
	return s
}

func (s *CreateTairAIGatewayInstanceRequest) SetRegionId(v string) *CreateTairAIGatewayInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTairAIGatewayInstanceRequest) SetResourceGroupId(v string) *CreateTairAIGatewayInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTairAIGatewayInstanceRequest) SetSecurityToken(v string) *CreateTairAIGatewayInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateTairAIGatewayInstanceRequest) SetTag(v []*CreateTairAIGatewayInstanceRequestTag) *CreateTairAIGatewayInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateTairAIGatewayInstanceRequest) SetVSwitchId(v string) *CreateTairAIGatewayInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateTairAIGatewayInstanceRequest) SetVpcId(v string) *CreateTairAIGatewayInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateTairAIGatewayInstanceRequest) SetZoneId(v string) *CreateTairAIGatewayInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateTairAIGatewayInstanceRequest) Validate() error {
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

type CreateTairAIGatewayInstanceRequestTag struct {
	// The tag key.
	//
	// > 	- **N*	- specifies the sequence number of the tag key. You can attach up to 20 tags to an instance. For example, Tag.1.Key specifies the key of the first tag, and Tag.2.Key specifies the key of the second tag.
	//
	// > 	- If the tag key does not exist, it is automatically created.
	//
	// example:
	//
	// demokey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// > **N*	- specifies the sequence number of the tag value. For example, **Tag.1.Value*	- specifies the value of the first tag, and **Tag.2.Value*	- specifies the value of the second tag.
	//
	// example:
	//
	// testvalue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTairAIGatewayInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateTairAIGatewayInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateTairAIGatewayInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateTairAIGatewayInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateTairAIGatewayInstanceRequestTag) SetKey(v string) *CreateTairAIGatewayInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateTairAIGatewayInstanceRequestTag) SetValue(v string) *CreateTairAIGatewayInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateTairAIGatewayInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}
