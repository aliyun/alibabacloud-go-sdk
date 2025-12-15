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
	// Specifies whether to enable automatic payment. Set the value to **true**.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. Default value: false. Valid values:
	//
	// 	- **true**: enabled
	//
	// 	- **false**: disables auto-renewal.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The subscription duration that is supported by auto-renewal. Unit: month. Valid values: **1**, **2**, **3**, **6**, and **12**.
	//
	// >  This parameter is required if the **AutoRenew*	- parameter is set to **true**.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *string `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// Specifies whether to use a coupon. Valid values:
	//
	// 	- **true**: uses a coupon.
	//
	// 	- **false**: does not use a coupon.
	//
	// example:
	//
	// false
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The extended information such as the promotional event ID and business information.
	//
	// example:
	//
	// 000000000
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The new billing method. Valid values:
	//
	// 	- **PrePaid**: subscription. If you set this parameter to PrePaid, you must also specify the **Period*	- parameter.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests and is case-sensitive. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The number of compute units. Valid values: 1.
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
	// 	- **true**: performs a dry run and does not create the instance. The system prechecks the request parameters, request format, service limits, and available resources. If the request fails to pass the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: performs a dry run and performs the actual request. If the request passes the dry run, the instance is created.
	//
	// example:
	//
	// false
	DryRun           *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	ElasticTimeRange *string `json:"ElasticTimeRange,omitempty" xml:"ElasticTimeRange,omitempty"`
	// Instance specification
	//
	// This parameter is required.
	//
	// example:
	//
	// kvcache.cu.g4b.2
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The name of the instance. The name must be 2 to 80 characters in length. The name must start with a letter and cannot contain spaces or the following special characters: `@ / : = " < > { [ ] }`
	//
	// example:
	//
	// vnodetest
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription duration. Valid values: **1*	- to **9**, **12**, **24**, and **36**. Unit: months.
	//
	// >  This parameter is required only if the **ChargeType*	- parameter is set to **PrePaid**.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The ID of the region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group that you want to manage.
	//
	// >
	//
	// 	- You can query resource group IDs in the console or by calling the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation. For more information, see [View the basic information about a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// 	- Before you modify the resource group to which an instance belongs, you can call the [ListResources](https://help.aliyun.com/document_detail/158866.html) operation to view the current resource group of the instance.
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// Details of the tags.
	Tag       []*CreateTairKVCacheVNodeRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	VNodeType *string                             `json:"VNodeType,omitempty" xml:"VNodeType,omitempty"`
	// The ID of the vSwitch to which the instance belongs. The vSwitch must belong to the VPC of the VCluser. You can call the [DescribeVpcs](https://help.aliyun.com/document_detail/35739.html) operation to query the VPC ID.
	//
	// >  The vSwitch and the instance must be deployed in the same zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VCluster that contains the VNode.
	//
	// This parameter is required.
	//
	// example:
	//
	// tk-2ze4bba3c8fe****
	VkName *string `json:"VkName,omitempty" xml:"VkName,omitempty"`
	// The zone ID of the instance.
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
	// >  A maximum of five key-value pairs can be specified at a time.
	//
	// example:
	//
	// value1_test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the instance.
	//
	// >  **N*	- specifies the value of the nth tag. For example, **Tag.1.Value*	- specifies the value of the first tag, and **Tag.2.Value*	- specifies the value of the second tag.
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
