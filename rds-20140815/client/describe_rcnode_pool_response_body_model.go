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
	// The node pool information.
	NodePoolList []*DescribeRCNodePoolResponseBodyNodePoolList `json:"NodePoolList,omitempty" xml:"NodePoolList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C816A4BF-A6EC-4722-95F9-2055859CCFD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// Indicates whether to enable automatic payment. Valid values:
	//
	// 	- **true*	- (default): enables the feature. You must make sure that your account balance is sufficient.
	//
	// 	- **false**: disables the feature. An unpaid order is generated.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Indicates whether to enable auto-renewal for the instance. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The ID of the container cluster in which the RDS Custom instance resides.
	//
	// example:
	//
	// c463aaa89e2b84cacacfbf23c4867****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Indicates whether to add the instance to the ACK cluster.
	//
	// example:
	//
	// 1
	CreateMode *string `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	// The data disks.
	DataDisk []*DescribeRCNodePoolResponseBodyNodePoolListDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// The ID of the deployment set.
	//
	// example:
	//
	// ds-bp18ukv66rlyuffv****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The instance description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance hostname.
	//
	// example:
	//
	// testHost1
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the image used by the instance.
	//
	// example:
	//
	// image-dsvjzw2ii8n4fvr****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The billing method. Valid value:
	//
	// 	- **Prepaid**: subscription
	//
	// 	- **Postpaid**: pay-as-you-go
	//
	// example:
	//
	// Prepaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance name.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type.
	//
	// example:
	//
	// mysql.i8.large.2cm
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The key pair name.
	//
	// example:
	//
	// dell5502
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The node pool ID.
	//
	// example:
	//
	// np31da1b38983f4511b490fc62108a****
	NodePoolId *string `json:"NodePoolId,omitempty" xml:"NodePoolId,omitempty"`
	// The name of the node pool.
	//
	// example:
	//
	// np31da1b38983f4511b490fc62108a****
	NodePoolName *string `json:"NodePoolName,omitempty" xml:"NodePoolName,omitempty"`
	// The password of the root user of the instance.
	//
	// example:
	//
	// testPassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The subscription duration.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription period. Valid values:
	//
	// 	- **Year**
	//
	// 	- **Month*	- (default)
	//
	// example:
	//
	// Year
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	SecurityEnhancementStrategy *string `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-uf6av412xaxixuez****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The specification of the system disk.
	SystemDisk *DescribeRCNodePoolResponseBodyNodePoolListSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// The tags.
	Tag []*DescribeRCNodePoolResponseBodyNodePoolListTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-zm0qvgv3sm3sjzbkr****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-beijing-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	// The type of the data disk. Set the value to **cloud_essd**, which indicates Enterprise SSDs (ESSDs).
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// Indicates whether to encrypt the cloud disk. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The performance level of the ESSD. Valid values:
	//
	// 	- **PL0**: A single ESSD delivers up to 10,000 random read/write IOPS.
	//
	// 	- **PL1**: A single ESSD delivers up to 50,000 random read/write IOPS.
	//
	// 	- **PL2**: A single ESSD delivers up to 100,000 random read/write IOPS.
	//
	// 	- **PL3**: A single ESSD delivers up to 1,000,000 random read/write IOPS.
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The data disk size. Unit: GiB.
	//
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
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
	// The type of the system disk. Set the value to **cloud_essd**, which indicates ESSDs.
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The performance level of the ESSD. Valid values:
	//
	// 	- **PL0**: A single ESSD delivers up to 10,000 random read/write IOPS.
	//
	// 	- **PL1**: A single ESSD delivers up to 50,000 random read/write IOPS.
	//
	// 	- **PL2**: A single ESSD delivers up to 100,000 random read/write IOPS.
	//
	// 	- **PL3**: A single ESSD delivers up to 1,000,000 random read/write IOPS.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of the system disk. Unit: GiB.
	//
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
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
	// The tag keys.
	//
	// example:
	//
	// Testkey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// Testvalue1
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
