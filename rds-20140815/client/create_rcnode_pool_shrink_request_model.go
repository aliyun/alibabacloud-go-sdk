// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRCNodePoolShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *CreateRCNodePoolShrinkRequest
	GetAmount() *int32
	SetAutoPay(v bool) *CreateRCNodePoolShrinkRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateRCNodePoolShrinkRequest
	GetAutoRenew() *bool
	SetClientToken(v string) *CreateRCNodePoolShrinkRequest
	GetClientToken() *string
	SetClusterId(v string) *CreateRCNodePoolShrinkRequest
	GetClusterId() *string
	SetCreateMode(v string) *CreateRCNodePoolShrinkRequest
	GetCreateMode() *string
	SetDataDiskShrink(v string) *CreateRCNodePoolShrinkRequest
	GetDataDiskShrink() *string
	SetDeploymentSetId(v string) *CreateRCNodePoolShrinkRequest
	GetDeploymentSetId() *string
	SetDescription(v string) *CreateRCNodePoolShrinkRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateRCNodePoolShrinkRequest
	GetDryRun() *bool
	SetHostName(v string) *CreateRCNodePoolShrinkRequest
	GetHostName() *string
	SetImageId(v string) *CreateRCNodePoolShrinkRequest
	GetImageId() *string
	SetInstanceChargeType(v string) *CreateRCNodePoolShrinkRequest
	GetInstanceChargeType() *string
	SetInstanceName(v string) *CreateRCNodePoolShrinkRequest
	GetInstanceName() *string
	SetInstanceType(v string) *CreateRCNodePoolShrinkRequest
	GetInstanceType() *string
	SetInternetChargeType(v string) *CreateRCNodePoolShrinkRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidthOut(v int32) *CreateRCNodePoolShrinkRequest
	GetInternetMaxBandwidthOut() *int32
	SetIoOptimized(v string) *CreateRCNodePoolShrinkRequest
	GetIoOptimized() *string
	SetKeyPairName(v string) *CreateRCNodePoolShrinkRequest
	GetKeyPairName() *string
	SetNodePoolName(v string) *CreateRCNodePoolShrinkRequest
	GetNodePoolName() *string
	SetPassword(v string) *CreateRCNodePoolShrinkRequest
	GetPassword() *string
	SetPeriod(v int32) *CreateRCNodePoolShrinkRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateRCNodePoolShrinkRequest
	GetPeriodUnit() *string
	SetRegionId(v string) *CreateRCNodePoolShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateRCNodePoolShrinkRequest
	GetResourceGroupId() *string
	SetSecurityEnhancementStrategy(v string) *CreateRCNodePoolShrinkRequest
	GetSecurityEnhancementStrategy() *string
	SetSecurityGroupId(v string) *CreateRCNodePoolShrinkRequest
	GetSecurityGroupId() *string
	SetSpotStrategy(v string) *CreateRCNodePoolShrinkRequest
	GetSpotStrategy() *string
	SetSupportCase(v string) *CreateRCNodePoolShrinkRequest
	GetSupportCase() *string
	SetSystemDiskShrink(v string) *CreateRCNodePoolShrinkRequest
	GetSystemDiskShrink() *string
	SetTag(v []*CreateRCNodePoolShrinkRequestTag) *CreateRCNodePoolShrinkRequest
	GetTag() []*CreateRCNodePoolShrinkRequestTag
	SetUserData(v string) *CreateRCNodePoolShrinkRequest
	GetUserData() *string
	SetVSwitchId(v string) *CreateRCNodePoolShrinkRequest
	GetVSwitchId() *string
	SetZoneId(v string) *CreateRCNodePoolShrinkRequest
	GetZoneId() *string
}

type CreateRCNodePoolShrinkRequest struct {
	Amount      *int32  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	AutoPay     *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew   *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	ClusterId          *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CreateMode         *string `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	DataDiskShrink     *string `json:"DataDisk,omitempty" xml:"DataDisk,omitempty"`
	DeploymentSetId    *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun             *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	HostName           *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	ImageId            *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InstanceName       *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	InstanceType            *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InternetChargeType      *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InternetMaxBandwidthOut *int32  `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	IoOptimized             *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	KeyPairName             *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	NodePoolName            *string `json:"NodePoolName,omitempty" xml:"NodePoolName,omitempty"`
	Password                *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Period                  *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit              *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// This parameter is required.
	RegionId                    *string                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId             *string                             `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityEnhancementStrategy *string                             `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	SecurityGroupId             *string                             `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SpotStrategy                *string                             `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SupportCase                 *string                             `json:"SupportCase,omitempty" xml:"SupportCase,omitempty"`
	SystemDiskShrink            *string                             `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty"`
	Tag                         []*CreateRCNodePoolShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	UserData                    *string                             `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// This parameter is required.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateRCNodePoolShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRCNodePoolShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRCNodePoolShrinkRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *CreateRCNodePoolShrinkRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateRCNodePoolShrinkRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateRCNodePoolShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateRCNodePoolShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateRCNodePoolShrinkRequest) GetCreateMode() *string {
	return s.CreateMode
}

func (s *CreateRCNodePoolShrinkRequest) GetDataDiskShrink() *string {
	return s.DataDiskShrink
}

func (s *CreateRCNodePoolShrinkRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *CreateRCNodePoolShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRCNodePoolShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateRCNodePoolShrinkRequest) GetHostName() *string {
	return s.HostName
}

func (s *CreateRCNodePoolShrinkRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateRCNodePoolShrinkRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateRCNodePoolShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateRCNodePoolShrinkRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateRCNodePoolShrinkRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateRCNodePoolShrinkRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *CreateRCNodePoolShrinkRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *CreateRCNodePoolShrinkRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateRCNodePoolShrinkRequest) GetNodePoolName() *string {
	return s.NodePoolName
}

func (s *CreateRCNodePoolShrinkRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateRCNodePoolShrinkRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateRCNodePoolShrinkRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateRCNodePoolShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRCNodePoolShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateRCNodePoolShrinkRequest) GetSecurityEnhancementStrategy() *string {
	return s.SecurityEnhancementStrategy
}

func (s *CreateRCNodePoolShrinkRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateRCNodePoolShrinkRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *CreateRCNodePoolShrinkRequest) GetSupportCase() *string {
	return s.SupportCase
}

func (s *CreateRCNodePoolShrinkRequest) GetSystemDiskShrink() *string {
	return s.SystemDiskShrink
}

func (s *CreateRCNodePoolShrinkRequest) GetTag() []*CreateRCNodePoolShrinkRequestTag {
	return s.Tag
}

func (s *CreateRCNodePoolShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateRCNodePoolShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateRCNodePoolShrinkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateRCNodePoolShrinkRequest) SetAmount(v int32) *CreateRCNodePoolShrinkRequest {
	s.Amount = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetAutoPay(v bool) *CreateRCNodePoolShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetAutoRenew(v bool) *CreateRCNodePoolShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetClientToken(v string) *CreateRCNodePoolShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetClusterId(v string) *CreateRCNodePoolShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetCreateMode(v string) *CreateRCNodePoolShrinkRequest {
	s.CreateMode = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetDataDiskShrink(v string) *CreateRCNodePoolShrinkRequest {
	s.DataDiskShrink = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetDeploymentSetId(v string) *CreateRCNodePoolShrinkRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetDescription(v string) *CreateRCNodePoolShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetDryRun(v bool) *CreateRCNodePoolShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetHostName(v string) *CreateRCNodePoolShrinkRequest {
	s.HostName = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetImageId(v string) *CreateRCNodePoolShrinkRequest {
	s.ImageId = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetInstanceChargeType(v string) *CreateRCNodePoolShrinkRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetInstanceName(v string) *CreateRCNodePoolShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetInstanceType(v string) *CreateRCNodePoolShrinkRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetInternetChargeType(v string) *CreateRCNodePoolShrinkRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetInternetMaxBandwidthOut(v int32) *CreateRCNodePoolShrinkRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetIoOptimized(v string) *CreateRCNodePoolShrinkRequest {
	s.IoOptimized = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetKeyPairName(v string) *CreateRCNodePoolShrinkRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetNodePoolName(v string) *CreateRCNodePoolShrinkRequest {
	s.NodePoolName = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetPassword(v string) *CreateRCNodePoolShrinkRequest {
	s.Password = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetPeriod(v int32) *CreateRCNodePoolShrinkRequest {
	s.Period = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetPeriodUnit(v string) *CreateRCNodePoolShrinkRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetRegionId(v string) *CreateRCNodePoolShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetResourceGroupId(v string) *CreateRCNodePoolShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetSecurityEnhancementStrategy(v string) *CreateRCNodePoolShrinkRequest {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetSecurityGroupId(v string) *CreateRCNodePoolShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetSpotStrategy(v string) *CreateRCNodePoolShrinkRequest {
	s.SpotStrategy = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetSupportCase(v string) *CreateRCNodePoolShrinkRequest {
	s.SupportCase = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetSystemDiskShrink(v string) *CreateRCNodePoolShrinkRequest {
	s.SystemDiskShrink = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetTag(v []*CreateRCNodePoolShrinkRequestTag) *CreateRCNodePoolShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetUserData(v string) *CreateRCNodePoolShrinkRequest {
	s.UserData = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetVSwitchId(v string) *CreateRCNodePoolShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) SetZoneId(v string) *CreateRCNodePoolShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequest) Validate() error {
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

type CreateRCNodePoolShrinkRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRCNodePoolShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateRCNodePoolShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateRCNodePoolShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateRCNodePoolShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateRCNodePoolShrinkRequestTag) SetKey(v string) *CreateRCNodePoolShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequestTag) SetValue(v string) *CreateRCNodePoolShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateRCNodePoolShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
