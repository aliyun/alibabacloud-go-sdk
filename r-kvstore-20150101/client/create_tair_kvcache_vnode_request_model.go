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
	SetVSwitchId(v string) *CreateTairKVCacheVNodeRequest
	GetVSwitchId() *string
	SetVkName(v string) *CreateTairKVCacheVNodeRequest
	GetVkName() *string
	SetZoneId(v string) *CreateTairKVCacheVNodeRequest
	GetZoneId() *string
}

type CreateTairKVCacheVNodeRequest struct {
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// 1
	AutoRenewPeriod *string `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// example:
	//
	// false
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// example:
	//
	// 000000000
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ComputeUnitNum *int32 `json:"ComputeUnitNum,omitempty" xml:"ComputeUnitNum,omitempty"`
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// kvcache.cu.g4b.2
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// example:
	//
	// vnodetest
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string                             `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string                             `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                              `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string                             `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tag                  []*CreateTairKVCacheVNodeRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// tk-2ze4bba3c8fe****
	VkName *string `json:"VkName,omitempty" xml:"VkName,omitempty"`
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
	return dara.Validate(s)
}

type CreateTairKVCacheVNodeRequestTag struct {
	// example:
	//
	// value1_test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
