// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRCNodePoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *CreateRCNodePoolRequest
	GetAmount() *int32
	SetAutoPay(v bool) *CreateRCNodePoolRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateRCNodePoolRequest
	GetAutoRenew() *bool
	SetClientToken(v string) *CreateRCNodePoolRequest
	GetClientToken() *string
	SetClusterId(v string) *CreateRCNodePoolRequest
	GetClusterId() *string
	SetCreateMode(v string) *CreateRCNodePoolRequest
	GetCreateMode() *string
	SetDataDisk(v []*CreateRCNodePoolRequestDataDisk) *CreateRCNodePoolRequest
	GetDataDisk() []*CreateRCNodePoolRequestDataDisk
	SetDeploymentSetId(v string) *CreateRCNodePoolRequest
	GetDeploymentSetId() *string
	SetDescription(v string) *CreateRCNodePoolRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateRCNodePoolRequest
	GetDryRun() *bool
	SetHostName(v string) *CreateRCNodePoolRequest
	GetHostName() *string
	SetImageId(v string) *CreateRCNodePoolRequest
	GetImageId() *string
	SetInstanceChargeType(v string) *CreateRCNodePoolRequest
	GetInstanceChargeType() *string
	SetInstanceName(v string) *CreateRCNodePoolRequest
	GetInstanceName() *string
	SetInstanceType(v string) *CreateRCNodePoolRequest
	GetInstanceType() *string
	SetInternetChargeType(v string) *CreateRCNodePoolRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidthOut(v int32) *CreateRCNodePoolRequest
	GetInternetMaxBandwidthOut() *int32
	SetIoOptimized(v string) *CreateRCNodePoolRequest
	GetIoOptimized() *string
	SetKeyPairName(v string) *CreateRCNodePoolRequest
	GetKeyPairName() *string
	SetNodePoolName(v string) *CreateRCNodePoolRequest
	GetNodePoolName() *string
	SetPassword(v string) *CreateRCNodePoolRequest
	GetPassword() *string
	SetPeriod(v int32) *CreateRCNodePoolRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateRCNodePoolRequest
	GetPeriodUnit() *string
	SetRegionId(v string) *CreateRCNodePoolRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateRCNodePoolRequest
	GetResourceGroupId() *string
	SetSecurityEnhancementStrategy(v string) *CreateRCNodePoolRequest
	GetSecurityEnhancementStrategy() *string
	SetSecurityGroupId(v string) *CreateRCNodePoolRequest
	GetSecurityGroupId() *string
	SetSpotStrategy(v string) *CreateRCNodePoolRequest
	GetSpotStrategy() *string
	SetSupportCase(v string) *CreateRCNodePoolRequest
	GetSupportCase() *string
	SetSystemDisk(v *CreateRCNodePoolRequestSystemDisk) *CreateRCNodePoolRequest
	GetSystemDisk() *CreateRCNodePoolRequestSystemDisk
	SetTag(v []*CreateRCNodePoolRequestTag) *CreateRCNodePoolRequest
	GetTag() []*CreateRCNodePoolRequestTag
	SetUserData(v string) *CreateRCNodePoolRequest
	GetUserData() *string
	SetVSwitchId(v string) *CreateRCNodePoolRequest
	GetVSwitchId() *string
	SetZoneId(v string) *CreateRCNodePoolRequest
	GetZoneId() *string
}

type CreateRCNodePoolRequest struct {
	Amount      *int32  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	AutoPay     *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew   *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	ClusterId          *string                            `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CreateMode         *string                            `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	DataDisk           []*CreateRCNodePoolRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	DeploymentSetId    *string                            `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	Description        *string                            `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun             *bool                              `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	HostName           *string                            `json:"HostName,omitempty" xml:"HostName,omitempty"`
	ImageId            *string                            `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceChargeType *string                            `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InstanceName       *string                            `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
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
	RegionId                    *string                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId             *string                            `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityEnhancementStrategy *string                            `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	SecurityGroupId             *string                            `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SpotStrategy                *string                            `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SupportCase                 *string                            `json:"SupportCase,omitempty" xml:"SupportCase,omitempty"`
	SystemDisk                  *CreateRCNodePoolRequestSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	Tag                         []*CreateRCNodePoolRequestTag      `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	UserData                    *string                            `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// This parameter is required.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateRCNodePoolRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRCNodePoolRequest) GoString() string {
	return s.String()
}

func (s *CreateRCNodePoolRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *CreateRCNodePoolRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateRCNodePoolRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateRCNodePoolRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateRCNodePoolRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateRCNodePoolRequest) GetCreateMode() *string {
	return s.CreateMode
}

func (s *CreateRCNodePoolRequest) GetDataDisk() []*CreateRCNodePoolRequestDataDisk {
	return s.DataDisk
}

func (s *CreateRCNodePoolRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *CreateRCNodePoolRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRCNodePoolRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateRCNodePoolRequest) GetHostName() *string {
	return s.HostName
}

func (s *CreateRCNodePoolRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateRCNodePoolRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateRCNodePoolRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateRCNodePoolRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateRCNodePoolRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateRCNodePoolRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *CreateRCNodePoolRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *CreateRCNodePoolRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateRCNodePoolRequest) GetNodePoolName() *string {
	return s.NodePoolName
}

func (s *CreateRCNodePoolRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateRCNodePoolRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateRCNodePoolRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateRCNodePoolRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRCNodePoolRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateRCNodePoolRequest) GetSecurityEnhancementStrategy() *string {
	return s.SecurityEnhancementStrategy
}

func (s *CreateRCNodePoolRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateRCNodePoolRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *CreateRCNodePoolRequest) GetSupportCase() *string {
	return s.SupportCase
}

func (s *CreateRCNodePoolRequest) GetSystemDisk() *CreateRCNodePoolRequestSystemDisk {
	return s.SystemDisk
}

func (s *CreateRCNodePoolRequest) GetTag() []*CreateRCNodePoolRequestTag {
	return s.Tag
}

func (s *CreateRCNodePoolRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateRCNodePoolRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateRCNodePoolRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateRCNodePoolRequest) SetAmount(v int32) *CreateRCNodePoolRequest {
	s.Amount = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetAutoPay(v bool) *CreateRCNodePoolRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetAutoRenew(v bool) *CreateRCNodePoolRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetClientToken(v string) *CreateRCNodePoolRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetClusterId(v string) *CreateRCNodePoolRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetCreateMode(v string) *CreateRCNodePoolRequest {
	s.CreateMode = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetDataDisk(v []*CreateRCNodePoolRequestDataDisk) *CreateRCNodePoolRequest {
	s.DataDisk = v
	return s
}

func (s *CreateRCNodePoolRequest) SetDeploymentSetId(v string) *CreateRCNodePoolRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetDescription(v string) *CreateRCNodePoolRequest {
	s.Description = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetDryRun(v bool) *CreateRCNodePoolRequest {
	s.DryRun = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetHostName(v string) *CreateRCNodePoolRequest {
	s.HostName = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetImageId(v string) *CreateRCNodePoolRequest {
	s.ImageId = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetInstanceChargeType(v string) *CreateRCNodePoolRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetInstanceName(v string) *CreateRCNodePoolRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetInstanceType(v string) *CreateRCNodePoolRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetInternetChargeType(v string) *CreateRCNodePoolRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetInternetMaxBandwidthOut(v int32) *CreateRCNodePoolRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetIoOptimized(v string) *CreateRCNodePoolRequest {
	s.IoOptimized = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetKeyPairName(v string) *CreateRCNodePoolRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetNodePoolName(v string) *CreateRCNodePoolRequest {
	s.NodePoolName = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetPassword(v string) *CreateRCNodePoolRequest {
	s.Password = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetPeriod(v int32) *CreateRCNodePoolRequest {
	s.Period = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetPeriodUnit(v string) *CreateRCNodePoolRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetRegionId(v string) *CreateRCNodePoolRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetResourceGroupId(v string) *CreateRCNodePoolRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetSecurityEnhancementStrategy(v string) *CreateRCNodePoolRequest {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetSecurityGroupId(v string) *CreateRCNodePoolRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetSpotStrategy(v string) *CreateRCNodePoolRequest {
	s.SpotStrategy = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetSupportCase(v string) *CreateRCNodePoolRequest {
	s.SupportCase = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetSystemDisk(v *CreateRCNodePoolRequestSystemDisk) *CreateRCNodePoolRequest {
	s.SystemDisk = v
	return s
}

func (s *CreateRCNodePoolRequest) SetTag(v []*CreateRCNodePoolRequestTag) *CreateRCNodePoolRequest {
	s.Tag = v
	return s
}

func (s *CreateRCNodePoolRequest) SetUserData(v string) *CreateRCNodePoolRequest {
	s.UserData = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetVSwitchId(v string) *CreateRCNodePoolRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateRCNodePoolRequest) SetZoneId(v string) *CreateRCNodePoolRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateRCNodePoolRequest) Validate() error {
	if s.DataDisk != nil {
		for _, item := range s.DataDisk {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
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

type CreateRCNodePoolRequestDataDisk struct {
	Category           *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DeleteWithInstance *bool   `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	Encrypted          *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	PerformanceLevel   *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	Size               *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateRCNodePoolRequestDataDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateRCNodePoolRequestDataDisk) GoString() string {
	return s.String()
}

func (s *CreateRCNodePoolRequestDataDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateRCNodePoolRequestDataDisk) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *CreateRCNodePoolRequestDataDisk) GetEncrypted() *string {
	return s.Encrypted
}

func (s *CreateRCNodePoolRequestDataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateRCNodePoolRequestDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateRCNodePoolRequestDataDisk) SetCategory(v string) *CreateRCNodePoolRequestDataDisk {
	s.Category = &v
	return s
}

func (s *CreateRCNodePoolRequestDataDisk) SetDeleteWithInstance(v bool) *CreateRCNodePoolRequestDataDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *CreateRCNodePoolRequestDataDisk) SetEncrypted(v string) *CreateRCNodePoolRequestDataDisk {
	s.Encrypted = &v
	return s
}

func (s *CreateRCNodePoolRequestDataDisk) SetPerformanceLevel(v string) *CreateRCNodePoolRequestDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateRCNodePoolRequestDataDisk) SetSize(v int32) *CreateRCNodePoolRequestDataDisk {
	s.Size = &v
	return s
}

func (s *CreateRCNodePoolRequestDataDisk) Validate() error {
	return dara.Validate(s)
}

type CreateRCNodePoolRequestSystemDisk struct {
	Category         *string `json:"Category,omitempty" xml:"Category,omitempty"`
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	Size             *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateRCNodePoolRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateRCNodePoolRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateRCNodePoolRequestSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateRCNodePoolRequestSystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateRCNodePoolRequestSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateRCNodePoolRequestSystemDisk) SetCategory(v string) *CreateRCNodePoolRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *CreateRCNodePoolRequestSystemDisk) SetPerformanceLevel(v string) *CreateRCNodePoolRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateRCNodePoolRequestSystemDisk) SetSize(v int32) *CreateRCNodePoolRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *CreateRCNodePoolRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}

type CreateRCNodePoolRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRCNodePoolRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateRCNodePoolRequestTag) GoString() string {
	return s.String()
}

func (s *CreateRCNodePoolRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateRCNodePoolRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateRCNodePoolRequestTag) SetKey(v string) *CreateRCNodePoolRequestTag {
	s.Key = &v
	return s
}

func (s *CreateRCNodePoolRequestTag) SetValue(v string) *CreateRCNodePoolRequestTag {
	s.Value = &v
	return s
}

func (s *CreateRCNodePoolRequestTag) Validate() error {
	return dara.Validate(s)
}
