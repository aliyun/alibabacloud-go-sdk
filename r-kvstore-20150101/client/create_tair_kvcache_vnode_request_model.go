// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTairKVCacheVNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateTairKVCacheVNodeRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateTairKVCacheVNodeRequest
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v string) *CreateTairKVCacheVNodeRequest
	GetAutoRenewPeriod() *string
	SetAutoUseCoupon(v bool) *CreateTairKVCacheVNodeRequest
	GetAutoUseCoupon() *bool
	SetBusinessInfo(v string) *CreateTairKVCacheVNodeRequest
	GetBusinessInfo() *string
	SetChargeType(v string) *CreateTairKVCacheVNodeRequest
	GetChargeType() *string
	SetClientToken(v string) *CreateTairKVCacheVNodeRequest
	GetClientToken() *string
	SetComputeUnitNum(v int32) *CreateTairKVCacheVNodeRequest
	GetComputeUnitNum() *int32
	SetCouponNo(v string) *CreateTairKVCacheVNodeRequest
	GetCouponNo() *string
	SetDryRun(v bool) *CreateTairKVCacheVNodeRequest
	GetDryRun() *bool
	SetElasticTimeRange(v string) *CreateTairKVCacheVNodeRequest
	GetElasticTimeRange() *string
	SetInstanceClass(v string) *CreateTairKVCacheVNodeRequest
	GetInstanceClass() *string
	SetInstanceName(v string) *CreateTairKVCacheVNodeRequest
	GetInstanceName() *string
	SetOwnerAccount(v string) *CreateTairKVCacheVNodeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTairKVCacheVNodeRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *CreateTairKVCacheVNodeRequest
	GetPeriod() *int32
	SetRegionId(v string) *CreateTairKVCacheVNodeRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateTairKVCacheVNodeRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateTairKVCacheVNodeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTairKVCacheVNodeRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *CreateTairKVCacheVNodeRequest
	GetSecurityToken() *string
	SetTag(v []*CreateTairKVCacheVNodeRequestTag) *CreateTairKVCacheVNodeRequest
	GetTag() []*CreateTairKVCacheVNodeRequestTag
	SetVNodeType(v string) *CreateTairKVCacheVNodeRequest
	GetVNodeType() *string
	SetVSwitchId(v string) *CreateTairKVCacheVNodeRequest
	GetVSwitchId() *string
	SetVkName(v string) *CreateTairKVCacheVNodeRequest
	GetVkName() *string
	SetZoneId(v string) *CreateTairKVCacheVNodeRequest
	GetZoneId() *string
}

type CreateTairKVCacheVNodeRequest struct {
	// Specifies whether to automatically complete the payment. The value must be **true**.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// - **true**: Enables auto-renewal.
	//
	// - **false*	- (default): Disables auto-renewal.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period, in months. Valid values: **1**, **2**, **3**, **6**, and **12**.
	//
	// > This parameter is required when the **AutoRenew*	- parameter is set to **true**.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *string `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// Specifies whether to use a coupon. Valid values:
	//
	// - **true**: Use a coupon.
	//
	// - **false*	- (default): Do not use a coupon.
	//
	// example:
	//
	// false
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// Additional business information, such as a promotion ID.
	//
	// example:
	//
	// 000000000
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The billing method for the instance. Valid value:
	//
	// - **PrePaid**: Subscription. If you specify this value, you must also specify the **Period*	- parameter.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// A client-generated token that ensures request idempotence. This token must be unique across requests, is case-sensitive, and cannot exceed 64 ASCII characters.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The number of compute units. Currently, only one compute unit is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ComputeUnitNum *int32 `json:"ComputeUnitNum,omitempty" xml:"ComputeUnitNum,omitempty"`
	// The coupon code.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: Performs a dry run and does not create the instance. The system checks the request parameters, request format, business limits, and available inventory. If the check fails, the system returns the corresponding error. If the check passes, the system returns the `DryRunOperation` error code.
	//
	// - **false*	- (default): Sends a normal request. If the check passes, the system creates the instance.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is no longer used.
	ElasticTimeRange *string `json:"ElasticTimeRange,omitempty" xml:"ElasticTimeRange,omitempty"`
	// The instance specification.
	//
	// This parameter is required.
	//
	// example:
	//
	// kvcache.cu.g4b.2
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The name of the new instance. The name must be 2 to 80 characters long and must start with a letter (case-insensitive) or a Chinese character. Spaces and the following special characters are not supported: `@/:=”<>{[]}`.
	//
	// example:
	//
	// vnodetest
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription period in months. Valid values: **1*	- to **9**, **12**, **24**, and **36**.
	//
	// > This parameter is required when the **ChargeType*	- parameter is set to **PrePaid**.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The ID of the region where you want to create the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance will belong.
	//
	// > - You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation or view resource group IDs in the console. For more information, see [View the basic information about a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// >
	//
	// > - Before changing the resource group of an instance, call the [ListResources](158866) API to view the current resource group of the instance.
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The tags to add to the instance. You can specify a maximum of five tags.
	Tag []*CreateTairKVCacheVNodeRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter is no longer used.
	VNodeType *string `json:"VNodeType,omitempty" xml:"VNodeType,omitempty"`
	// The ID of the vSwitch for the instance. The vSwitch must belong to the VPC that is associated with the specified virtual cluster. You can call the [DescribeVpcs](https://help.aliyun.com/document_detail/35739.html) operation to obtain the vSwitch ID.
	//
	// > The vSwitch must be in the same zone as the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual cluster that hosts the VNode.
	//
	// This parameter is required.
	//
	// example:
	//
	// tk-2ze4bba3c8fe****
	VkName *string `json:"VkName,omitempty" xml:"VkName,omitempty"`
	// The ID of the zone where you want to create the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateTairKVCacheVNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTairKVCacheVNodeRequest) GoString() string {
	return s.String()
}

func (s *CreateTairKVCacheVNodeRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateTairKVCacheVNodeRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateTairKVCacheVNodeRequest) GetAutoRenewPeriod() *string {
	return s.AutoRenewPeriod
}

func (s *CreateTairKVCacheVNodeRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *CreateTairKVCacheVNodeRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *CreateTairKVCacheVNodeRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateTairKVCacheVNodeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTairKVCacheVNodeRequest) GetComputeUnitNum() *int32 {
	return s.ComputeUnitNum
}

func (s *CreateTairKVCacheVNodeRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *CreateTairKVCacheVNodeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateTairKVCacheVNodeRequest) GetElasticTimeRange() *string {
	return s.ElasticTimeRange
}

func (s *CreateTairKVCacheVNodeRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *CreateTairKVCacheVNodeRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateTairKVCacheVNodeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTairKVCacheVNodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTairKVCacheVNodeRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateTairKVCacheVNodeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTairKVCacheVNodeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateTairKVCacheVNodeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTairKVCacheVNodeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTairKVCacheVNodeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateTairKVCacheVNodeRequest) GetTag() []*CreateTairKVCacheVNodeRequestTag {
	return s.Tag
}

func (s *CreateTairKVCacheVNodeRequest) GetVNodeType() *string {
	return s.VNodeType
}

func (s *CreateTairKVCacheVNodeRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateTairKVCacheVNodeRequest) GetVkName() *string {
	return s.VkName
}

func (s *CreateTairKVCacheVNodeRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateTairKVCacheVNodeRequest) SetAutoPay(v bool) *CreateTairKVCacheVNodeRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) SetAutoRenew(v bool) *CreateTairKVCacheVNodeRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) SetAutoRenewPeriod(v string) *CreateTairKVCacheVNodeRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) SetAutoUseCoupon(v bool) *CreateTairKVCacheVNodeRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) SetBusinessInfo(v string) *CreateTairKVCacheVNodeRequest {
	s.BusinessInfo = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) SetChargeType(v string) *CreateTairKVCacheVNodeRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) SetClientToken(v string) *CreateTairKVCacheVNodeRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) SetComputeUnitNum(v int32) *CreateTairKVCacheVNodeRequest {
	s.ComputeUnitNum = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) SetCouponNo(v string) *CreateTairKVCacheVNodeRequest {
	s.CouponNo = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) SetDryRun(v bool) *CreateTairKVCacheVNodeRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) SetElasticTimeRange(v string) *CreateTairKVCacheVNodeRequest {
	s.ElasticTimeRange = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) SetInstanceClass(v string) *CreateTairKVCacheVNodeRequest {
	s.InstanceClass = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) SetInstanceName(v string) *CreateTairKVCacheVNodeRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) SetOwnerAccount(v string) *CreateTairKVCacheVNodeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) SetOwnerId(v int64) *CreateTairKVCacheVNodeRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) SetPeriod(v int32) *CreateTairKVCacheVNodeRequest {
	s.Period = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) SetRegionId(v string) *CreateTairKVCacheVNodeRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) SetResourceGroupId(v string) *CreateTairKVCacheVNodeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) SetResourceOwnerAccount(v string) *CreateTairKVCacheVNodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) SetResourceOwnerId(v int64) *CreateTairKVCacheVNodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) SetSecurityToken(v string) *CreateTairKVCacheVNodeRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) SetTag(v []*CreateTairKVCacheVNodeRequestTag) *CreateTairKVCacheVNodeRequest {
	s.Tag = v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) SetVNodeType(v string) *CreateTairKVCacheVNodeRequest {
	s.VNodeType = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) SetVSwitchId(v string) *CreateTairKVCacheVNodeRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) SetVkName(v string) *CreateTairKVCacheVNodeRequest {
	s.VkName = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) SetZoneId(v string) *CreateTairKVCacheVNodeRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequest) Validate() error {
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

type CreateTairKVCacheVNodeRequestTag struct {
	// The tag key.
	//
	// > You can specify up to 5 tag key-value pairs at a time.
	//
	// example:
	//
	// value1_test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// > **N*	- represents the index of a tag, starting from 1. For example, **Tag.1.Value*	- is the value of the first tag.
	//
	// example:
	//
	// key1_test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTairKVCacheVNodeRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateTairKVCacheVNodeRequestTag) GoString() string {
	return s.String()
}

func (s *CreateTairKVCacheVNodeRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateTairKVCacheVNodeRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateTairKVCacheVNodeRequestTag) SetKey(v string) *CreateTairKVCacheVNodeRequestTag {
	s.Key = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequestTag) SetValue(v string) *CreateTairKVCacheVNodeRequestTag {
	s.Value = &v
	return s
}

func (s *CreateTairKVCacheVNodeRequestTag) Validate() error {
	return dara.Validate(s)
}
