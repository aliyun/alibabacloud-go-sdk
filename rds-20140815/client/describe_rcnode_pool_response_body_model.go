// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCNodePoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodePoolList(v []*DescribeRCNodePoolResponseBodyNodePoolList) *DescribeRCNodePoolResponseBody
	GetNodePoolList() []*DescribeRCNodePoolResponseBodyNodePoolList
	SetRequestId(v string) *DescribeRCNodePoolResponseBody
	GetRequestId() *string
}

type DescribeRCNodePoolResponseBody struct {
	NodePoolList []*DescribeRCNodePoolResponseBodyNodePoolList `json:"NodePoolList,omitempty" xml:"NodePoolList,omitempty" type:"Repeated"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRCNodePoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCNodePoolResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCNodePoolResponseBody) GetNodePoolList() []*DescribeRCNodePoolResponseBodyNodePoolList {
	return s.NodePoolList
}

func (s *DescribeRCNodePoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCNodePoolResponseBody) SetNodePoolList(v []*DescribeRCNodePoolResponseBodyNodePoolList) *DescribeRCNodePoolResponseBody {
	s.NodePoolList = v
	return s
}

func (s *DescribeRCNodePoolResponseBody) SetRequestId(v string) *DescribeRCNodePoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCNodePoolResponseBody) Validate() error {
	if s.NodePoolList != nil {
		for _, item := range s.NodePoolList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCNodePoolResponseBodyNodePoolList struct {
	AutoPay                     *bool                                                 `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew                   *bool                                                 `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ClusterId                   *string                                               `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CreateMode                  *string                                               `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	DataDisk                    []*DescribeRCNodePoolResponseBodyNodePoolListDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	DeploymentSetId             *string                                               `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	Description                 *string                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	HostName                    *string                                               `json:"HostName,omitempty" xml:"HostName,omitempty"`
	ImageId                     *string                                               `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceChargeType          *string                                               `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InstanceName                *string                                               `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceType                *string                                               `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InternetChargeType          *string                                               `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InternetMaxBandwidthOut     *int32                                                `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	IoOptimized                 *string                                               `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	KeyPairName                 *string                                               `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	NodePoolId                  *string                                               `json:"NodePoolId,omitempty" xml:"NodePoolId,omitempty"`
	NodePoolName                *string                                               `json:"NodePoolName,omitempty" xml:"NodePoolName,omitempty"`
	Password                    *string                                               `json:"Password,omitempty" xml:"Password,omitempty"`
	Period                      *int32                                                `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit                  *string                                               `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	RegionId                    *string                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId             *string                                               `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityEnhancementStrategy *string                                               `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	SecurityGroupId             *string                                               `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SpotStrategy                *string                                               `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SystemDisk                  *DescribeRCNodePoolResponseBodyNodePoolListSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	Tag                         []*DescribeRCNodePoolResponseBodyNodePoolListTag      `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	VSwitchId                   *string                                               `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId                      *string                                               `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRCNodePoolResponseBodyNodePoolList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCNodePoolResponseBodyNodePoolList) GoString() string {
	return s.String()
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetCreateMode() *string {
	return s.CreateMode
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetDataDisk() []*DescribeRCNodePoolResponseBodyNodePoolListDataDisk {
	return s.DataDisk
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetDescription() *string {
	return s.Description
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetHostName() *string {
	return s.HostName
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetNodePoolId() *string {
	return s.NodePoolId
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetNodePoolName() *string {
	return s.NodePoolName
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetPassword() *string {
	return s.Password
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetSecurityEnhancementStrategy() *string {
	return s.SecurityEnhancementStrategy
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetSystemDisk() *DescribeRCNodePoolResponseBodyNodePoolListSystemDisk {
	return s.SystemDisk
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetTag() []*DescribeRCNodePoolResponseBodyNodePoolListTag {
	return s.Tag
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetAutoPay(v bool) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.AutoPay = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetAutoRenew(v bool) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.AutoRenew = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetClusterId(v string) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.ClusterId = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetCreateMode(v string) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.CreateMode = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetDataDisk(v []*DescribeRCNodePoolResponseBodyNodePoolListDataDisk) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.DataDisk = v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetDeploymentSetId(v string) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.DeploymentSetId = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetDescription(v string) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.Description = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetHostName(v string) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.HostName = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetImageId(v string) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.ImageId = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetInstanceChargeType(v string) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetInstanceName(v string) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.InstanceName = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetInstanceType(v string) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.InstanceType = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetInternetChargeType(v string) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetInternetMaxBandwidthOut(v int32) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetIoOptimized(v string) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.IoOptimized = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetKeyPairName(v string) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.KeyPairName = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetNodePoolId(v string) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.NodePoolId = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetNodePoolName(v string) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.NodePoolName = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetPassword(v string) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.Password = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetPeriod(v int32) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.Period = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetPeriodUnit(v string) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetRegionId(v string) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.RegionId = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetResourceGroupId(v string) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetSecurityEnhancementStrategy(v string) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetSecurityGroupId(v string) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetSpotStrategy(v string) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetSystemDisk(v *DescribeRCNodePoolResponseBodyNodePoolListSystemDisk) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.SystemDisk = v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetTag(v []*DescribeRCNodePoolResponseBodyNodePoolListTag) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.Tag = v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetVSwitchId(v string) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.VSwitchId = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) SetZoneId(v string) *DescribeRCNodePoolResponseBodyNodePoolList {
	s.ZoneId = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolList) Validate() error {
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

type DescribeRCNodePoolResponseBodyNodePoolListDataDisk struct {
	Category           *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DeleteWithInstance *bool   `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	Encrypted          *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	PerformanceLevel   *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	Size               *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeRCNodePoolResponseBodyNodePoolListDataDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCNodePoolResponseBodyNodePoolListDataDisk) GoString() string {
	return s.String()
}

func (s *DescribeRCNodePoolResponseBodyNodePoolListDataDisk) GetCategory() *string {
	return s.Category
}

func (s *DescribeRCNodePoolResponseBodyNodePoolListDataDisk) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *DescribeRCNodePoolResponseBodyNodePoolListDataDisk) GetEncrypted() *string {
	return s.Encrypted
}

func (s *DescribeRCNodePoolResponseBodyNodePoolListDataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribeRCNodePoolResponseBodyNodePoolListDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *DescribeRCNodePoolResponseBodyNodePoolListDataDisk) SetCategory(v string) *DescribeRCNodePoolResponseBodyNodePoolListDataDisk {
	s.Category = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolListDataDisk) SetDeleteWithInstance(v bool) *DescribeRCNodePoolResponseBodyNodePoolListDataDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolListDataDisk) SetEncrypted(v string) *DescribeRCNodePoolResponseBodyNodePoolListDataDisk {
	s.Encrypted = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolListDataDisk) SetPerformanceLevel(v string) *DescribeRCNodePoolResponseBodyNodePoolListDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolListDataDisk) SetSize(v int32) *DescribeRCNodePoolResponseBodyNodePoolListDataDisk {
	s.Size = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolListDataDisk) Validate() error {
	return dara.Validate(s)
}

type DescribeRCNodePoolResponseBodyNodePoolListSystemDisk struct {
	Category         *string `json:"Category,omitempty" xml:"Category,omitempty"`
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	Size             *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeRCNodePoolResponseBodyNodePoolListSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCNodePoolResponseBodyNodePoolListSystemDisk) GoString() string {
	return s.String()
}

func (s *DescribeRCNodePoolResponseBodyNodePoolListSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *DescribeRCNodePoolResponseBodyNodePoolListSystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribeRCNodePoolResponseBodyNodePoolListSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *DescribeRCNodePoolResponseBodyNodePoolListSystemDisk) SetCategory(v string) *DescribeRCNodePoolResponseBodyNodePoolListSystemDisk {
	s.Category = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolListSystemDisk) SetPerformanceLevel(v string) *DescribeRCNodePoolResponseBodyNodePoolListSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolListSystemDisk) SetSize(v int32) *DescribeRCNodePoolResponseBodyNodePoolListSystemDisk {
	s.Size = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolListSystemDisk) Validate() error {
	return dara.Validate(s)
}

type DescribeRCNodePoolResponseBodyNodePoolListTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRCNodePoolResponseBodyNodePoolListTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCNodePoolResponseBodyNodePoolListTag) GoString() string {
	return s.String()
}

func (s *DescribeRCNodePoolResponseBodyNodePoolListTag) GetKey() *string {
	return s.Key
}

func (s *DescribeRCNodePoolResponseBodyNodePoolListTag) GetValue() *string {
	return s.Value
}

func (s *DescribeRCNodePoolResponseBodyNodePoolListTag) SetKey(v string) *DescribeRCNodePoolResponseBodyNodePoolListTag {
	s.Key = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolListTag) SetValue(v string) *DescribeRCNodePoolResponseBodyNodePoolListTag {
	s.Value = &v
	return s
}

func (s *DescribeRCNodePoolResponseBodyNodePoolListTag) Validate() error {
	return dara.Validate(s)
}
