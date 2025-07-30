// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTCInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v string) *CreateTCInstanceRequest
	GetAutoRenew() *string
	SetAutoRenewPeriod(v string) *CreateTCInstanceRequest
	GetAutoRenewPeriod() *string
	SetAutoUseCoupon(v string) *CreateTCInstanceRequest
	GetAutoUseCoupon() *string
	SetBusinessInfo(v string) *CreateTCInstanceRequest
	GetBusinessInfo() *string
	SetClientToken(v string) *CreateTCInstanceRequest
	GetClientToken() *string
	SetCouponNo(v string) *CreateTCInstanceRequest
	GetCouponNo() *string
	SetDataDisk(v []*CreateTCInstanceRequestDataDisk) *CreateTCInstanceRequest
	GetDataDisk() []*CreateTCInstanceRequestDataDisk
	SetDryRun(v bool) *CreateTCInstanceRequest
	GetDryRun() *bool
	SetImageId(v string) *CreateTCInstanceRequest
	GetImageId() *string
	SetInstanceChargeType(v string) *CreateTCInstanceRequest
	GetInstanceChargeType() *string
	SetInstanceClass(v string) *CreateTCInstanceRequest
	GetInstanceClass() *string
	SetInstanceName(v string) *CreateTCInstanceRequest
	GetInstanceName() *string
	SetNeedEni(v bool) *CreateTCInstanceRequest
	GetNeedEni() *bool
	SetNetworkType(v string) *CreateTCInstanceRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *CreateTCInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTCInstanceRequest
	GetOwnerId() *int64
	SetPeriod(v string) *CreateTCInstanceRequest
	GetPeriod() *string
	SetRegionId(v string) *CreateTCInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateTCInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateTCInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTCInstanceRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupId(v string) *CreateTCInstanceRequest
	GetSecurityGroupId() *string
	SetSecurityToken(v string) *CreateTCInstanceRequest
	GetSecurityToken() *string
	SetTag(v []*CreateTCInstanceRequestTag) *CreateTCInstanceRequest
	GetTag() []*CreateTCInstanceRequestTag
	SetVSwitchId(v string) *CreateTCInstanceRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateTCInstanceRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateTCInstanceRequest
	GetZoneId() *string
}

type CreateTCInstanceRequest struct {
	// example:
	//
	// false
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// 1
	AutoRenewPeriod *string `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// example:
	//
	// false
	AutoUseCoupon *string `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// example:
	//
	// 000000000
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// This parameter is required.
	DataDisk []*CreateTCInstanceRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ubuntu_20_04_64_20G_alibase_20210412
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// example:
	//
	// tair.kvcache.guis.8.gu60
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// example:
	//
	// newinstancename
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	NeedEni      *bool   `json:"NeedEni,omitempty" xml:"NeedEni,omitempty"`
	// example:
	//
	// VPC
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 12
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmyiu4e******
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// sg-bpcfmyiu4ekp****
	SecurityGroupId *string                       `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityToken   *string                       `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tag             []*CreateTCInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-bp1nme44gek34slfc****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateTCInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTCInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateTCInstanceRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *CreateTCInstanceRequest) GetAutoRenewPeriod() *string {
	return s.AutoRenewPeriod
}

func (s *CreateTCInstanceRequest) GetAutoUseCoupon() *string {
	return s.AutoUseCoupon
}

func (s *CreateTCInstanceRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *CreateTCInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTCInstanceRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *CreateTCInstanceRequest) GetDataDisk() []*CreateTCInstanceRequestDataDisk {
	return s.DataDisk
}

func (s *CreateTCInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateTCInstanceRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateTCInstanceRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateTCInstanceRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *CreateTCInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateTCInstanceRequest) GetNeedEni() *bool {
	return s.NeedEni
}

func (s *CreateTCInstanceRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateTCInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTCInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTCInstanceRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateTCInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTCInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateTCInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTCInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTCInstanceRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateTCInstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateTCInstanceRequest) GetTag() []*CreateTCInstanceRequestTag {
	return s.Tag
}

func (s *CreateTCInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateTCInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateTCInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateTCInstanceRequest) SetAutoRenew(v string) *CreateTCInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateTCInstanceRequest) SetAutoRenewPeriod(v string) *CreateTCInstanceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateTCInstanceRequest) SetAutoUseCoupon(v string) *CreateTCInstanceRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateTCInstanceRequest) SetBusinessInfo(v string) *CreateTCInstanceRequest {
	s.BusinessInfo = &v
	return s
}

func (s *CreateTCInstanceRequest) SetClientToken(v string) *CreateTCInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTCInstanceRequest) SetCouponNo(v string) *CreateTCInstanceRequest {
	s.CouponNo = &v
	return s
}

func (s *CreateTCInstanceRequest) SetDataDisk(v []*CreateTCInstanceRequestDataDisk) *CreateTCInstanceRequest {
	s.DataDisk = v
	return s
}

func (s *CreateTCInstanceRequest) SetDryRun(v bool) *CreateTCInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTCInstanceRequest) SetImageId(v string) *CreateTCInstanceRequest {
	s.ImageId = &v
	return s
}

func (s *CreateTCInstanceRequest) SetInstanceChargeType(v string) *CreateTCInstanceRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateTCInstanceRequest) SetInstanceClass(v string) *CreateTCInstanceRequest {
	s.InstanceClass = &v
	return s
}

func (s *CreateTCInstanceRequest) SetInstanceName(v string) *CreateTCInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateTCInstanceRequest) SetNeedEni(v bool) *CreateTCInstanceRequest {
	s.NeedEni = &v
	return s
}

func (s *CreateTCInstanceRequest) SetNetworkType(v string) *CreateTCInstanceRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateTCInstanceRequest) SetOwnerAccount(v string) *CreateTCInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTCInstanceRequest) SetOwnerId(v int64) *CreateTCInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTCInstanceRequest) SetPeriod(v string) *CreateTCInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateTCInstanceRequest) SetRegionId(v string) *CreateTCInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTCInstanceRequest) SetResourceGroupId(v string) *CreateTCInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTCInstanceRequest) SetResourceOwnerAccount(v string) *CreateTCInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTCInstanceRequest) SetResourceOwnerId(v int64) *CreateTCInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTCInstanceRequest) SetSecurityGroupId(v string) *CreateTCInstanceRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateTCInstanceRequest) SetSecurityToken(v string) *CreateTCInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateTCInstanceRequest) SetTag(v []*CreateTCInstanceRequestTag) *CreateTCInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateTCInstanceRequest) SetVSwitchId(v string) *CreateTCInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateTCInstanceRequest) SetVpcId(v string) *CreateTCInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateTCInstanceRequest) SetZoneId(v string) *CreateTCInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateTCInstanceRequest) Validate() error {
	return dara.Validate(s)
}

type CreateTCInstanceRequestDataDisk struct {
	// This parameter is required.
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateTCInstanceRequestDataDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateTCInstanceRequestDataDisk) GoString() string {
	return s.String()
}

func (s *CreateTCInstanceRequestDataDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateTCInstanceRequestDataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateTCInstanceRequestDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateTCInstanceRequestDataDisk) SetCategory(v string) *CreateTCInstanceRequestDataDisk {
	s.Category = &v
	return s
}

func (s *CreateTCInstanceRequestDataDisk) SetPerformanceLevel(v string) *CreateTCInstanceRequestDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateTCInstanceRequestDataDisk) SetSize(v int32) *CreateTCInstanceRequestDataDisk {
	s.Size = &v
	return s
}

func (s *CreateTCInstanceRequestDataDisk) Validate() error {
	return dara.Validate(s)
}

type CreateTCInstanceRequestTag struct {
	// example:
	//
	// key1_test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// testvalue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTCInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateTCInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateTCInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateTCInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateTCInstanceRequestTag) SetKey(v string) *CreateTCInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateTCInstanceRequestTag) SetValue(v string) *CreateTCInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateTCInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}
