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
	DataDiskShrink *string `json:"DataDisk,omitempty" xml:"DataDisk,omitempty"`
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
	SystemDiskShrink *string `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty"`
	// The list of tags.
	Tag []*CreateRCNodePoolShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
