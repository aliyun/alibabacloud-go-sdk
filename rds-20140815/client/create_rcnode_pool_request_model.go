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
	// The number of RDS Custom instances to create. This parameter is applicable only to batch creation of RDS Custom instances.
	//
	// Valid values: **1*	- to **5**. Default value: **1**.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Specifies whether to enable automatic payment.
	//
	// Valid values:
	//
	// - **true**: Automatic payment is enabled. Make sure that your account balance is sufficient.
	//
	// - **false**: Only an order is generated. No payment is made.
	//
	//
	// > The default value is true. If your payment method has an insufficient balance, set AutoPay to false. In this case, an unpaid order is generated. You can log on to the ApsaraDB RDS console to complete the payment.
	//
	// >
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal. This parameter is valid only when you create subscription instances. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// > 	- If you purchase on a monthly basis, the auto-renewal epoch is 1 month.
	//
	// > 	- If you purchase on a yearly basis, the auto-renewal epoch is 1 year.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the RDS Custom container cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// c463aaa89e2b84cacacfbf23c4867****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to allow the instance to join an ACK cluster. If this parameter settings is set to **1**, the created instance can be added to an ACK cluster for efficient container application management.
	//
	// - **1**: Yes.
	//
	// - **0*	- (default): No.
	//
	// example:
	//
	// 1
	CreateMode *string `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	// The list of data cloud disks.
	DataDisk []*CreateRCNodePoolRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// The deployment set ID.
	//
	// example:
	//
	// ds-uf6c8qerk019bj1l****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The instance description. The description must be 2 to 256 characters in length and can contain letters and Chinese characters. The description cannot start with http:// or https://.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run for this request. Valid values:
	//
	// 	- **true**: performs a dry run without creating the instance. The system checks the request parameters, request format, service limits, and available stock.
	//
	// 	- **false*	- (default): sends the request. If the request passes the check, the instance is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The hostname of the instance.
	//
	// example:
	//
	// testHost1
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The image ID used by the instance.
	//
	// example:
	//
	// image-dsvjzw2ii8n4fvr6de
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **Prepaid**: subscription.
	//
	// 	- **Postpaid**: pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance name.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type. For the instance types supported by RDS Custom instances, see [RDS Custom instance types](https://help.aliyun.com/document_detail/2844823.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql.i8.large.2cm
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// A reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// A reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// A reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The name of the key pair. Only a single name is supported.
	//
	// example:
	//
	// dell5502
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The name of the node pool.
	//
	// example:
	//
	// testNodePool
	NodePoolName *string `json:"NodePoolName,omitempty" xml:"NodePoolName,omitempty"`
	// The password of the root account of the instance.
	//
	// example:
	//
	// testPassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The subscription duration of the resource. Default value: **1**.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration for the subscription billable methods. Valid values:
	//
	// - **Year**
	//
	// - **Month*	- (default)
	//
	// example:
	//
	// Year
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// A reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	SecurityEnhancementStrategy *string `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	// The security group ID. You can specify an existing security group ID. If the security group does not exist, automatic creation of a security group is performed.
	//
	// example:
	//
	// sg-m5e9abdu1rtxa12b****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// A reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The supported scenario. This parameter is required when **createMode*	- is set to **1**. Currently, only **edge*	- is supported.
	//
	// example:
	//
	// edge
	SupportCase *string `json:"SupportCase,omitempty" xml:"SupportCase,omitempty"`
	// The system cloud disk specifications.
	SystemDisk *CreateRCNodePoolRequestSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// The list of tags.
	Tag []*CreateRCNodePoolRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// A reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The vSwitch ID.
	//
	// > The vSwitch must be in the same zone as the ApsaraDB RDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-uf6adz52c2p****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the instance.
	//
	// > If you specify the VSwitchId parameter, the ZoneId parameter must match the zone of the specified vSwitch. You can also leave this parameter empty, and the system automatically selects the zone of the specified vSwitch.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	// The type of the data cloud disk. Only **cloud_essd*	- (ESSD cloud disk) is supported. For more information about standard SSDs and other cloud disk types, see the related documentation.
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// A reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// Specifies whether to encrypt the data cloud disk. Valid values:
	//
	// - **true**
	//
	// - **false*	- (default)
	//
	// example:
	//
	// false
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The performance level (PL) of the ESSD cloud disk. For standard SSDs and other cloud disk types, this parameter is not applicable. Valid values:
	//
	// - **PL0**: A single cloud disk can deliver up to 10,000 random read/write IOPS.
	//
	// - **PL1**: A single cloud disk can deliver up to 50,000 random read/write IOPS.
	//
	// - **PL2**: A single cloud disk can deliver up to 100,000 random read/write IOPS.
	//
	// - **PL3**: A single cloud disk can deliver up to 1,000,000 random read/write IOPS.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of the data cloud disk. Unit: GiB. Valid values: 20 to 65536.
	//
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
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
	// The category of the system cloud disk. Only **cloud_essd*	- (Enterprise SSD) is supported.
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The performance level (PL) of the ESSD cloud disk. For standard SSDs and other cloud disk types, this parameter is not applicable. Valid values:
	//
	// - **PL0**: A single cloud disk can deliver up to 10,000 random read/write IOPS.
	//
	// - **PL1**: A single cloud disk can deliver up to 50,000 random read/write IOPS.
	//
	// - **PL2**: A single cloud disk can deliver up to 100,000 random read/write IOPS.
	//
	// - **PL3**: A single cloud disk can deliver up to 1,000,000 random read/write IOPS.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of the system cloud disk. Unit: GiB. Valid values: 20 to 2048.
	//
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
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
	// The tag key. You can create up to N tag keys at a time. Valid values of N: **1 to 20**. The tag key cannot be an empty string.
	//
	// example:
	//
	// testkey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value that corresponds to the tag key. You can create up to N tag values at a time. Valid values of N: **1*	- to **20**. The tag value can be an empty string.
	//
	// example:
	//
	// testvalue1
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
