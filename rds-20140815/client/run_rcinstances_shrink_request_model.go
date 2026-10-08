// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunRCInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcuType(v string) *RunRCInstancesShrinkRequest
	GetAcuType() *string
	SetAmount(v int32) *RunRCInstancesShrinkRequest
	GetAmount() *int32
	SetAutoPay(v bool) *RunRCInstancesShrinkRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *RunRCInstancesShrinkRequest
	GetAutoRenew() *bool
	SetAutoUseCoupon(v bool) *RunRCInstancesShrinkRequest
	GetAutoUseCoupon() *bool
	SetBusinessInfo(v string) *RunRCInstancesShrinkRequest
	GetBusinessInfo() *string
	SetClientToken(v string) *RunRCInstancesShrinkRequest
	GetClientToken() *string
	SetCreateAckEdgeParamShrink(v string) *RunRCInstancesShrinkRequest
	GetCreateAckEdgeParamShrink() *string
	SetCreateExtraParam(v string) *RunRCInstancesShrinkRequest
	GetCreateExtraParam() *string
	SetCreateMode(v string) *RunRCInstancesShrinkRequest
	GetCreateMode() *string
	SetDataDiskShrink(v string) *RunRCInstancesShrinkRequest
	GetDataDiskShrink() *string
	SetDeletionProtection(v bool) *RunRCInstancesShrinkRequest
	GetDeletionProtection() *bool
	SetDeploymentSetId(v string) *RunRCInstancesShrinkRequest
	GetDeploymentSetId() *string
	SetDescription(v string) *RunRCInstancesShrinkRequest
	GetDescription() *string
	SetDryRun(v bool) *RunRCInstancesShrinkRequest
	GetDryRun() *bool
	SetHostName(v string) *RunRCInstancesShrinkRequest
	GetHostName() *string
	SetImageId(v string) *RunRCInstancesShrinkRequest
	GetImageId() *string
	SetInstanceChargeType(v string) *RunRCInstancesShrinkRequest
	GetInstanceChargeType() *string
	SetInstanceName(v string) *RunRCInstancesShrinkRequest
	GetInstanceName() *string
	SetInstanceType(v string) *RunRCInstancesShrinkRequest
	GetInstanceType() *string
	SetInternetChargeType(v string) *RunRCInstancesShrinkRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidthOut(v int32) *RunRCInstancesShrinkRequest
	GetInternetMaxBandwidthOut() *int32
	SetIoOptimized(v string) *RunRCInstancesShrinkRequest
	GetIoOptimized() *string
	SetKeyPairName(v string) *RunRCInstancesShrinkRequest
	GetKeyPairName() *string
	SetNetworkOptionsShrink(v string) *RunRCInstancesShrinkRequest
	GetNetworkOptionsShrink() *string
	SetPassword(v string) *RunRCInstancesShrinkRequest
	GetPassword() *string
	SetPasswordInherit(v bool) *RunRCInstancesShrinkRequest
	GetPasswordInherit() *bool
	SetPeriod(v int32) *RunRCInstancesShrinkRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RunRCInstancesShrinkRequest
	GetPeriodUnit() *string
	SetPrivateIpAddress(v string) *RunRCInstancesShrinkRequest
	GetPrivateIpAddress() *string
	SetPromotionCode(v string) *RunRCInstancesShrinkRequest
	GetPromotionCode() *string
	SetRegionId(v string) *RunRCInstancesShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *RunRCInstancesShrinkRequest
	GetResourceGroupId() *string
	SetScheduledRule(v string) *RunRCInstancesShrinkRequest
	GetScheduledRule() *string
	SetSecurityEnhancementStrategy(v string) *RunRCInstancesShrinkRequest
	GetSecurityEnhancementStrategy() *string
	SetSecurityGroupId(v string) *RunRCInstancesShrinkRequest
	GetSecurityGroupId() *string
	SetSecurityGroupIdsShrink(v string) *RunRCInstancesShrinkRequest
	GetSecurityGroupIdsShrink() *string
	SetSpotStrategy(v string) *RunRCInstancesShrinkRequest
	GetSpotStrategy() *string
	SetSupportCase(v string) *RunRCInstancesShrinkRequest
	GetSupportCase() *string
	SetSystemDiskShrink(v string) *RunRCInstancesShrinkRequest
	GetSystemDiskShrink() *string
	SetTag(v []*RunRCInstancesShrinkRequestTag) *RunRCInstancesShrinkRequest
	GetTag() []*RunRCInstancesShrinkRequestTag
	SetUserData(v string) *RunRCInstancesShrinkRequest
	GetUserData() *string
	SetUserDataInBase64(v bool) *RunRCInstancesShrinkRequest
	GetUserDataInBase64() *bool
	SetVSwitchId(v string) *RunRCInstancesShrinkRequest
	GetVSwitchId() *string
	SetZoneId(v string) *RunRCInstancesShrinkRequest
	GetZoneId() *string
}

type RunRCInstancesShrinkRequest struct {
	// The ACU type.
	//
	// example:
	//
	// gn8is
	AcuType *string `json:"AcuType,omitempty" xml:"AcuType,omitempty"`
	// The number of RDS Custom instances to create. This parameter is applicable only to batch creation of RDS Custom instances.
	//
	// Valid values: **1*	- to **30**. Default value: **1**.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Specifies whether to enable automatic payment. Valid values:
	//
	// - **true*	- (default): Automatic payment is enabled. Ensure that your account balance is sufficient.
	//
	// - **false**: Only an order is generated. No payment is made.
	//
	// > If your payment method has an insufficient balance, set the AutoPay parameter to false. An unpaid order is generated, and you can log on to the ApsaraDB RDS console to complete the payment.
	//
	// >
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// 	- **true*	- (default): Auto-renewal is enabled.
	//
	// 	- **false**: Auto-renewal is disabled.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Specifies whether to automatically use coupons. Valid values:
	//
	// 	- **true*	- (default): Coupons are automatically used.
	//
	// 	- **false**: Coupons are not automatically used.
	//
	// > After a coupon is used, the amount deducted by the coupon is not refunded if you perform a downgrade operation.
	//
	// example:
	//
	// true
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The business information.
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The client token that is used to ensure the idempotence of the request and prevent repeated submissions. The value is generated by the client and must be unique across different requests. The value can be up to 64 ASCII characters in length and cannot contain non-ASCII characters.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ACK Edge cluster information.
	CreateAckEdgeParamShrink *string `json:"CreateAckEdgeParam,omitempty" xml:"CreateAckEdgeParam,omitempty"`
	// Reserved parameter. Not supported.
	//
	// example:
	//
	// None
	CreateExtraParam *string `json:"CreateExtraParam,omitempty" xml:"CreateExtraParam,omitempty"`
	// Specifies whether the instance can be added to an ACK cluster. If this parameter is set to **1**, the created instance can be added to an ACK cluster by calling the **AttachRCInstances*	- API operation, which enables efficient management of containerized applications.
	//
	// - **1**: The instance can be added to an ACK cluster.
	//
	// - **0*	- (default): The instance cannot be added to an ACK cluster.
	//
	// example:
	//
	// 0
	CreateMode *string `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	// The list of data cloud disks.
	DataDiskShrink *string `json:"DataDisk,omitempty" xml:"DataDisk,omitempty"`
	// Specifies whether to enable deletion protection. Valid values:
	//
	// 	- **true**: Deletion protection is enabled.
	//
	// 	- **false*	- (default): Deletion protection is disabled.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The deployment set ID.
	//
	// example:
	//
	// ds-uf6670sipmph********
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The instance description. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// Instance_Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run for the instance creation. Valid values:
	//
	// 	- **true**: A dry run is performed without creating the instance. The check items include request parameters, request format, business limits, and inventory.
	//
	// 	- **false*	- (default): A normal request is sent. After the check is passed, the instance is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The hostname of the instance (2 to 64 characters).
	//
	// - Periods (.) can be used to separate the hostname into multiple segments. Each segment can contain uppercase and lowercase letters, digits, and hyphens (-).
	//
	// - Periods (.) and hyphens (-) cannot be used as the first or last character, and cannot be used consecutively.
	//
	// example:
	//
	// testHost1
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The image ID used by the instance.
	//
	// example:
	//
	// image-dsvjzw2ii8n4******
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **Prepaid**: subscription.
	//
	// 	- **Postpaid**: pay-as-you-go.
	//
	// example:
	//
	// Prepaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance name. The name must be 2 to 128 characters in length and must start with an uppercase or lowercase letter or a Chinese character. The name can contain uppercase and lowercase letters, Chinese characters, digits, periods (.), underscores (_), colons (:), or hyphens (-). The default value is the InstanceId of the instance. When creating multiple RDS Custom instances, you can set sequential instance names that can contain brackets ([]) and commas (,). For more information, see [Create an RDS Custom instance](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/create-an-rds-custom-instance#00481f9ba381u).
	//
	// example:
	//
	// rc-node-[99,1]-rchost
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type. For the instance types supported by RDS Custom instances, see [RDS Custom instance types](https://help.aliyun.com/document_detail/2844823.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql.i8.large.2cm
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Reserved parameter. Not supported.
	//
	// example:
	//
	// null
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum outbound public bandwidth for Custom for SQL Server. Unit: Mbit/s.
	//
	// Valid values: 0 to 1024. Default value: 0.
	//
	// example:
	//
	// 0
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// Reserved parameter. Not supported.
	//
	// example:
	//
	// null
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The name of the key pair. Only a single name is supported.
	//
	// example:
	//
	// dell5502
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The network-related attribute parameters.
	NetworkOptionsShrink *string `json:"NetworkOptions,omitempty" xml:"NetworkOptions,omitempty"`
	// The password of the instance account. The password must be 8 to 30 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. The following special characters are supported: `()~!@#$%^&*-_+=|{}[]:;\\"<>,.?/`.
	//
	// example:
	//
	// TestRDS123!
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether to use the preset password of the image. When this parameter is used, the Password parameter must be empty, and the image must have a password configured. Default value: false.
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// The subscription duration of the resource. Default value: **1**.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription billable methods duration. Valid values:
	//
	// - **Year**
	//
	// - **Month*	- (default)
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The private IP address of the instance. When setting the private IP address for a VPC-type ECS instance, you must select an address from the idle CIDR block of the vSwitch (VSwitchId).
	//
	// example:
	//
	// ``10.1.**.**``
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The coupon code.
	//
	// example:
	//
	// 72329885****
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// The region ID. You can call DescribeRegions to obtain the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The time-based elastic scaling rule.
	//
	// example:
	//
	// {"rule":[{"beginTime":"09:00","endTime":"17:00","acu":4}]}
	ScheduledRule *string `json:"ScheduledRule,omitempty" xml:"ScheduledRule,omitempty"`
	// Reserved parameter. Not supported.
	//
	// example:
	//
	// null
	SecurityEnhancementStrategy *string `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	// The ID of the security group to which the instance belongs. Instances in the same security group can communicate with each other. The maximum number of instances that a security group can contain depends on the security group type. For more information, see the security group section in [Limits](https://help.aliyun.com/document_detail/25412.html).
	//
	// > The SecurityGroupId parameter determines the network type of the instance. For example, if the specified security group is of the VPC type, the instance is a VPC-type instance, and you must also specify the VSwitchId parameter.
	//
	// example:
	//
	// sg-uf6av412xaxixu******
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// Adds the instance to multiple security groups. The maximum number of associated security groups is 10. You cannot set both SecurityGroupId and SecurityGroupIds.N at the same time.
	SecurityGroupIdsShrink *string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty"`
	// The bidding strategy for pay-as-you-go instances. This parameter takes effect only when the **InstanceChargeType*	- parameter is set to **PostPaid**. Valid values:
	//
	// - **NoSpot**: a regular pay-as-you-go instance.
	//
	// - **SpotAsPriceGo**: the system automatically bids at the current market price.
	//
	// Default value: **NoSpot**.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The form factor of RDS Custom. Valid values:
	//
	// - **eni**: dual network interface.
	//
	// - **edge**: edge node pool.
	//
	// - **share**: VPC.
	//
	// example:
	//
	// share
	SupportCase *string `json:"SupportCase,omitempty" xml:"SupportCase,omitempty"`
	// The system cloud disk specifications.
	SystemDiskShrink *string `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty"`
	// The list of tags.
	Tag []*RunRCInstancesShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The instance user data. The raw data can be up to 32 KB in size.
	//
	// Do not pass confidential information such as passwords and private keys in plaintext. If you must pass such information, encrypt it first and then use Base64 encoding before transmission. Perform decryption inside the instance. The following example shows how to transform a script to a Base64 character string:
	//
	// ```
	//
	// echo -n \\"#!/bin/sh
	//
	// echo "Hello World"\\" | base64 -w 0
	//
	// ```
	//
	// example:
	//
	// IyEvYmluL3NoCmVjaG8gIkhlbGxvIFdvcmxkLiBUaGUgdGltZSBpcyBub3cgJChkYXRlIC1SKSIhIHwgdGVlIC9yb290L3VzZXJkYXRhX3Rlc3QudHh0
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// Specifies whether the custom data is Base64-encoded.
	//
	// - **true**: The custom data is Base64-encoded.
	//
	// - **false*	- (default): The custom data is not Base64-encoded.
	//
	// example:
	//
	// true
	UserDataInBase64 *bool `json:"UserDataInBase64,omitempty" xml:"UserDataInBase64,omitempty"`
	// The vSwitch ID of the target instance. If you are creating a VPC-type RDS Custom instance, you must specify the vSwitch ID. The security group and the vSwitch must belong to the same VPC.
	//
	// > If you configure the VSwitchId parameter, the ZoneId parameter must match the zone of the vSwitch. You can also leave ZoneId empty, and the system automatically selects the zone of the specified vSwitch.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-2vcd61ngm890sk****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the instance. You can call DescribeZones to obtain the list of zones.
	//
	// > If you specify the VSwitchId parameter, the ZoneId parameter must match the zone of the vSwitch. You can also leave ZoneId empty, and the system automatically selects the zone of the specified vSwitch.
	//
	// example:
	//
	// cn-beijing-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s RunRCInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunRCInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunRCInstancesShrinkRequest) GetAcuType() *string {
	return s.AcuType
}

func (s *RunRCInstancesShrinkRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *RunRCInstancesShrinkRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *RunRCInstancesShrinkRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RunRCInstancesShrinkRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *RunRCInstancesShrinkRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *RunRCInstancesShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RunRCInstancesShrinkRequest) GetCreateAckEdgeParamShrink() *string {
	return s.CreateAckEdgeParamShrink
}

func (s *RunRCInstancesShrinkRequest) GetCreateExtraParam() *string {
	return s.CreateExtraParam
}

func (s *RunRCInstancesShrinkRequest) GetCreateMode() *string {
	return s.CreateMode
}

func (s *RunRCInstancesShrinkRequest) GetDataDiskShrink() *string {
	return s.DataDiskShrink
}

func (s *RunRCInstancesShrinkRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *RunRCInstancesShrinkRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *RunRCInstancesShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *RunRCInstancesShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RunRCInstancesShrinkRequest) GetHostName() *string {
	return s.HostName
}

func (s *RunRCInstancesShrinkRequest) GetImageId() *string {
	return s.ImageId
}

func (s *RunRCInstancesShrinkRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *RunRCInstancesShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *RunRCInstancesShrinkRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *RunRCInstancesShrinkRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *RunRCInstancesShrinkRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *RunRCInstancesShrinkRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *RunRCInstancesShrinkRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *RunRCInstancesShrinkRequest) GetNetworkOptionsShrink() *string {
	return s.NetworkOptionsShrink
}

func (s *RunRCInstancesShrinkRequest) GetPassword() *string {
	return s.Password
}

func (s *RunRCInstancesShrinkRequest) GetPasswordInherit() *bool {
	return s.PasswordInherit
}

func (s *RunRCInstancesShrinkRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RunRCInstancesShrinkRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RunRCInstancesShrinkRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *RunRCInstancesShrinkRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *RunRCInstancesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RunRCInstancesShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RunRCInstancesShrinkRequest) GetScheduledRule() *string {
	return s.ScheduledRule
}

func (s *RunRCInstancesShrinkRequest) GetSecurityEnhancementStrategy() *string {
	return s.SecurityEnhancementStrategy
}

func (s *RunRCInstancesShrinkRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *RunRCInstancesShrinkRequest) GetSecurityGroupIdsShrink() *string {
	return s.SecurityGroupIdsShrink
}

func (s *RunRCInstancesShrinkRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *RunRCInstancesShrinkRequest) GetSupportCase() *string {
	return s.SupportCase
}

func (s *RunRCInstancesShrinkRequest) GetSystemDiskShrink() *string {
	return s.SystemDiskShrink
}

func (s *RunRCInstancesShrinkRequest) GetTag() []*RunRCInstancesShrinkRequestTag {
	return s.Tag
}

func (s *RunRCInstancesShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *RunRCInstancesShrinkRequest) GetUserDataInBase64() *bool {
	return s.UserDataInBase64
}

func (s *RunRCInstancesShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *RunRCInstancesShrinkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *RunRCInstancesShrinkRequest) SetAcuType(v string) *RunRCInstancesShrinkRequest {
	s.AcuType = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetAmount(v int32) *RunRCInstancesShrinkRequest {
	s.Amount = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetAutoPay(v bool) *RunRCInstancesShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetAutoRenew(v bool) *RunRCInstancesShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetAutoUseCoupon(v bool) *RunRCInstancesShrinkRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetBusinessInfo(v string) *RunRCInstancesShrinkRequest {
	s.BusinessInfo = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetClientToken(v string) *RunRCInstancesShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetCreateAckEdgeParamShrink(v string) *RunRCInstancesShrinkRequest {
	s.CreateAckEdgeParamShrink = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetCreateExtraParam(v string) *RunRCInstancesShrinkRequest {
	s.CreateExtraParam = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetCreateMode(v string) *RunRCInstancesShrinkRequest {
	s.CreateMode = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetDataDiskShrink(v string) *RunRCInstancesShrinkRequest {
	s.DataDiskShrink = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetDeletionProtection(v bool) *RunRCInstancesShrinkRequest {
	s.DeletionProtection = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetDeploymentSetId(v string) *RunRCInstancesShrinkRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetDescription(v string) *RunRCInstancesShrinkRequest {
	s.Description = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetDryRun(v bool) *RunRCInstancesShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetHostName(v string) *RunRCInstancesShrinkRequest {
	s.HostName = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetImageId(v string) *RunRCInstancesShrinkRequest {
	s.ImageId = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetInstanceChargeType(v string) *RunRCInstancesShrinkRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetInstanceName(v string) *RunRCInstancesShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetInstanceType(v string) *RunRCInstancesShrinkRequest {
	s.InstanceType = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetInternetChargeType(v string) *RunRCInstancesShrinkRequest {
	s.InternetChargeType = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetInternetMaxBandwidthOut(v int32) *RunRCInstancesShrinkRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetIoOptimized(v string) *RunRCInstancesShrinkRequest {
	s.IoOptimized = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetKeyPairName(v string) *RunRCInstancesShrinkRequest {
	s.KeyPairName = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetNetworkOptionsShrink(v string) *RunRCInstancesShrinkRequest {
	s.NetworkOptionsShrink = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetPassword(v string) *RunRCInstancesShrinkRequest {
	s.Password = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetPasswordInherit(v bool) *RunRCInstancesShrinkRequest {
	s.PasswordInherit = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetPeriod(v int32) *RunRCInstancesShrinkRequest {
	s.Period = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetPeriodUnit(v string) *RunRCInstancesShrinkRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetPrivateIpAddress(v string) *RunRCInstancesShrinkRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetPromotionCode(v string) *RunRCInstancesShrinkRequest {
	s.PromotionCode = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetRegionId(v string) *RunRCInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetResourceGroupId(v string) *RunRCInstancesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetScheduledRule(v string) *RunRCInstancesShrinkRequest {
	s.ScheduledRule = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetSecurityEnhancementStrategy(v string) *RunRCInstancesShrinkRequest {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetSecurityGroupId(v string) *RunRCInstancesShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetSecurityGroupIdsShrink(v string) *RunRCInstancesShrinkRequest {
	s.SecurityGroupIdsShrink = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetSpotStrategy(v string) *RunRCInstancesShrinkRequest {
	s.SpotStrategy = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetSupportCase(v string) *RunRCInstancesShrinkRequest {
	s.SupportCase = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetSystemDiskShrink(v string) *RunRCInstancesShrinkRequest {
	s.SystemDiskShrink = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetTag(v []*RunRCInstancesShrinkRequestTag) *RunRCInstancesShrinkRequest {
	s.Tag = v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetUserData(v string) *RunRCInstancesShrinkRequest {
	s.UserData = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetUserDataInBase64(v bool) *RunRCInstancesShrinkRequest {
	s.UserDataInBase64 = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetVSwitchId(v string) *RunRCInstancesShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetZoneId(v string) *RunRCInstancesShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) Validate() error {
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

type RunRCInstancesShrinkRequestTag struct {
	// The tag key. You can create up to N tag keys at a time. Valid values of N: **1 to 20**. Empty strings are not allowed.
	//
	// example:
	//
	// Testkey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value corresponding to the tag key. You can create up to N tag values at a time. Valid values of N: **1 to 20**. Empty strings are allowed.
	//
	// example:
	//
	// Testvalue1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RunRCInstancesShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s RunRCInstancesShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *RunRCInstancesShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *RunRCInstancesShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *RunRCInstancesShrinkRequestTag) SetKey(v string) *RunRCInstancesShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *RunRCInstancesShrinkRequestTag) SetValue(v string) *RunRCInstancesShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *RunRCInstancesShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
