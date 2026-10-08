// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunRCInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcuType(v string) *RunRCInstancesRequest
	GetAcuType() *string
	SetAmount(v int32) *RunRCInstancesRequest
	GetAmount() *int32
	SetAutoPay(v bool) *RunRCInstancesRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *RunRCInstancesRequest
	GetAutoRenew() *bool
	SetAutoUseCoupon(v bool) *RunRCInstancesRequest
	GetAutoUseCoupon() *bool
	SetBusinessInfo(v string) *RunRCInstancesRequest
	GetBusinessInfo() *string
	SetClientToken(v string) *RunRCInstancesRequest
	GetClientToken() *string
	SetCreateAckEdgeParam(v *RunRCInstancesRequestCreateAckEdgeParam) *RunRCInstancesRequest
	GetCreateAckEdgeParam() *RunRCInstancesRequestCreateAckEdgeParam
	SetCreateExtraParam(v string) *RunRCInstancesRequest
	GetCreateExtraParam() *string
	SetCreateMode(v string) *RunRCInstancesRequest
	GetCreateMode() *string
	SetDataDisk(v []*RunRCInstancesRequestDataDisk) *RunRCInstancesRequest
	GetDataDisk() []*RunRCInstancesRequestDataDisk
	SetDeletionProtection(v bool) *RunRCInstancesRequest
	GetDeletionProtection() *bool
	SetDeploymentSetId(v string) *RunRCInstancesRequest
	GetDeploymentSetId() *string
	SetDescription(v string) *RunRCInstancesRequest
	GetDescription() *string
	SetDryRun(v bool) *RunRCInstancesRequest
	GetDryRun() *bool
	SetHostName(v string) *RunRCInstancesRequest
	GetHostName() *string
	SetImageId(v string) *RunRCInstancesRequest
	GetImageId() *string
	SetInstanceChargeType(v string) *RunRCInstancesRequest
	GetInstanceChargeType() *string
	SetInstanceName(v string) *RunRCInstancesRequest
	GetInstanceName() *string
	SetInstanceType(v string) *RunRCInstancesRequest
	GetInstanceType() *string
	SetInternetChargeType(v string) *RunRCInstancesRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidthOut(v int32) *RunRCInstancesRequest
	GetInternetMaxBandwidthOut() *int32
	SetIoOptimized(v string) *RunRCInstancesRequest
	GetIoOptimized() *string
	SetKeyPairName(v string) *RunRCInstancesRequest
	GetKeyPairName() *string
	SetNetworkOptions(v *RunRCInstancesRequestNetworkOptions) *RunRCInstancesRequest
	GetNetworkOptions() *RunRCInstancesRequestNetworkOptions
	SetPassword(v string) *RunRCInstancesRequest
	GetPassword() *string
	SetPasswordInherit(v bool) *RunRCInstancesRequest
	GetPasswordInherit() *bool
	SetPeriod(v int32) *RunRCInstancesRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RunRCInstancesRequest
	GetPeriodUnit() *string
	SetPrivateIpAddress(v string) *RunRCInstancesRequest
	GetPrivateIpAddress() *string
	SetPromotionCode(v string) *RunRCInstancesRequest
	GetPromotionCode() *string
	SetRegionId(v string) *RunRCInstancesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *RunRCInstancesRequest
	GetResourceGroupId() *string
	SetScheduledRule(v string) *RunRCInstancesRequest
	GetScheduledRule() *string
	SetSecurityEnhancementStrategy(v string) *RunRCInstancesRequest
	GetSecurityEnhancementStrategy() *string
	SetSecurityGroupId(v string) *RunRCInstancesRequest
	GetSecurityGroupId() *string
	SetSecurityGroupIds(v []*string) *RunRCInstancesRequest
	GetSecurityGroupIds() []*string
	SetSpotStrategy(v string) *RunRCInstancesRequest
	GetSpotStrategy() *string
	SetSupportCase(v string) *RunRCInstancesRequest
	GetSupportCase() *string
	SetSystemDisk(v *RunRCInstancesRequestSystemDisk) *RunRCInstancesRequest
	GetSystemDisk() *RunRCInstancesRequestSystemDisk
	SetTag(v []*RunRCInstancesRequestTag) *RunRCInstancesRequest
	GetTag() []*RunRCInstancesRequestTag
	SetUserData(v string) *RunRCInstancesRequest
	GetUserData() *string
	SetUserDataInBase64(v bool) *RunRCInstancesRequest
	GetUserDataInBase64() *bool
	SetVSwitchId(v string) *RunRCInstancesRequest
	GetVSwitchId() *string
	SetZoneId(v string) *RunRCInstancesRequest
	GetZoneId() *string
}

type RunRCInstancesRequest struct {
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
	CreateAckEdgeParam *RunRCInstancesRequestCreateAckEdgeParam `json:"CreateAckEdgeParam,omitempty" xml:"CreateAckEdgeParam,omitempty" type:"Struct"`
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
	DataDisk []*RunRCInstancesRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
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
	NetworkOptions *RunRCInstancesRequestNetworkOptions `json:"NetworkOptions,omitempty" xml:"NetworkOptions,omitempty" type:"Struct"`
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
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
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
	SystemDisk *RunRCInstancesRequestSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// The list of tags.
	Tag []*RunRCInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s RunRCInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s RunRCInstancesRequest) GoString() string {
	return s.String()
}

func (s *RunRCInstancesRequest) GetAcuType() *string {
	return s.AcuType
}

func (s *RunRCInstancesRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *RunRCInstancesRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *RunRCInstancesRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RunRCInstancesRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *RunRCInstancesRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *RunRCInstancesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RunRCInstancesRequest) GetCreateAckEdgeParam() *RunRCInstancesRequestCreateAckEdgeParam {
	return s.CreateAckEdgeParam
}

func (s *RunRCInstancesRequest) GetCreateExtraParam() *string {
	return s.CreateExtraParam
}

func (s *RunRCInstancesRequest) GetCreateMode() *string {
	return s.CreateMode
}

func (s *RunRCInstancesRequest) GetDataDisk() []*RunRCInstancesRequestDataDisk {
	return s.DataDisk
}

func (s *RunRCInstancesRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *RunRCInstancesRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *RunRCInstancesRequest) GetDescription() *string {
	return s.Description
}

func (s *RunRCInstancesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RunRCInstancesRequest) GetHostName() *string {
	return s.HostName
}

func (s *RunRCInstancesRequest) GetImageId() *string {
	return s.ImageId
}

func (s *RunRCInstancesRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *RunRCInstancesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *RunRCInstancesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *RunRCInstancesRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *RunRCInstancesRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *RunRCInstancesRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *RunRCInstancesRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *RunRCInstancesRequest) GetNetworkOptions() *RunRCInstancesRequestNetworkOptions {
	return s.NetworkOptions
}

func (s *RunRCInstancesRequest) GetPassword() *string {
	return s.Password
}

func (s *RunRCInstancesRequest) GetPasswordInherit() *bool {
	return s.PasswordInherit
}

func (s *RunRCInstancesRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RunRCInstancesRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RunRCInstancesRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *RunRCInstancesRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *RunRCInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RunRCInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RunRCInstancesRequest) GetScheduledRule() *string {
	return s.ScheduledRule
}

func (s *RunRCInstancesRequest) GetSecurityEnhancementStrategy() *string {
	return s.SecurityEnhancementStrategy
}

func (s *RunRCInstancesRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *RunRCInstancesRequest) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *RunRCInstancesRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *RunRCInstancesRequest) GetSupportCase() *string {
	return s.SupportCase
}

func (s *RunRCInstancesRequest) GetSystemDisk() *RunRCInstancesRequestSystemDisk {
	return s.SystemDisk
}

func (s *RunRCInstancesRequest) GetTag() []*RunRCInstancesRequestTag {
	return s.Tag
}

func (s *RunRCInstancesRequest) GetUserData() *string {
	return s.UserData
}

func (s *RunRCInstancesRequest) GetUserDataInBase64() *bool {
	return s.UserDataInBase64
}

func (s *RunRCInstancesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *RunRCInstancesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *RunRCInstancesRequest) SetAcuType(v string) *RunRCInstancesRequest {
	s.AcuType = &v
	return s
}

func (s *RunRCInstancesRequest) SetAmount(v int32) *RunRCInstancesRequest {
	s.Amount = &v
	return s
}

func (s *RunRCInstancesRequest) SetAutoPay(v bool) *RunRCInstancesRequest {
	s.AutoPay = &v
	return s
}

func (s *RunRCInstancesRequest) SetAutoRenew(v bool) *RunRCInstancesRequest {
	s.AutoRenew = &v
	return s
}

func (s *RunRCInstancesRequest) SetAutoUseCoupon(v bool) *RunRCInstancesRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *RunRCInstancesRequest) SetBusinessInfo(v string) *RunRCInstancesRequest {
	s.BusinessInfo = &v
	return s
}

func (s *RunRCInstancesRequest) SetClientToken(v string) *RunRCInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *RunRCInstancesRequest) SetCreateAckEdgeParam(v *RunRCInstancesRequestCreateAckEdgeParam) *RunRCInstancesRequest {
	s.CreateAckEdgeParam = v
	return s
}

func (s *RunRCInstancesRequest) SetCreateExtraParam(v string) *RunRCInstancesRequest {
	s.CreateExtraParam = &v
	return s
}

func (s *RunRCInstancesRequest) SetCreateMode(v string) *RunRCInstancesRequest {
	s.CreateMode = &v
	return s
}

func (s *RunRCInstancesRequest) SetDataDisk(v []*RunRCInstancesRequestDataDisk) *RunRCInstancesRequest {
	s.DataDisk = v
	return s
}

func (s *RunRCInstancesRequest) SetDeletionProtection(v bool) *RunRCInstancesRequest {
	s.DeletionProtection = &v
	return s
}

func (s *RunRCInstancesRequest) SetDeploymentSetId(v string) *RunRCInstancesRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *RunRCInstancesRequest) SetDescription(v string) *RunRCInstancesRequest {
	s.Description = &v
	return s
}

func (s *RunRCInstancesRequest) SetDryRun(v bool) *RunRCInstancesRequest {
	s.DryRun = &v
	return s
}

func (s *RunRCInstancesRequest) SetHostName(v string) *RunRCInstancesRequest {
	s.HostName = &v
	return s
}

func (s *RunRCInstancesRequest) SetImageId(v string) *RunRCInstancesRequest {
	s.ImageId = &v
	return s
}

func (s *RunRCInstancesRequest) SetInstanceChargeType(v string) *RunRCInstancesRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *RunRCInstancesRequest) SetInstanceName(v string) *RunRCInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *RunRCInstancesRequest) SetInstanceType(v string) *RunRCInstancesRequest {
	s.InstanceType = &v
	return s
}

func (s *RunRCInstancesRequest) SetInternetChargeType(v string) *RunRCInstancesRequest {
	s.InternetChargeType = &v
	return s
}

func (s *RunRCInstancesRequest) SetInternetMaxBandwidthOut(v int32) *RunRCInstancesRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *RunRCInstancesRequest) SetIoOptimized(v string) *RunRCInstancesRequest {
	s.IoOptimized = &v
	return s
}

func (s *RunRCInstancesRequest) SetKeyPairName(v string) *RunRCInstancesRequest {
	s.KeyPairName = &v
	return s
}

func (s *RunRCInstancesRequest) SetNetworkOptions(v *RunRCInstancesRequestNetworkOptions) *RunRCInstancesRequest {
	s.NetworkOptions = v
	return s
}

func (s *RunRCInstancesRequest) SetPassword(v string) *RunRCInstancesRequest {
	s.Password = &v
	return s
}

func (s *RunRCInstancesRequest) SetPasswordInherit(v bool) *RunRCInstancesRequest {
	s.PasswordInherit = &v
	return s
}

func (s *RunRCInstancesRequest) SetPeriod(v int32) *RunRCInstancesRequest {
	s.Period = &v
	return s
}

func (s *RunRCInstancesRequest) SetPeriodUnit(v string) *RunRCInstancesRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RunRCInstancesRequest) SetPrivateIpAddress(v string) *RunRCInstancesRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *RunRCInstancesRequest) SetPromotionCode(v string) *RunRCInstancesRequest {
	s.PromotionCode = &v
	return s
}

func (s *RunRCInstancesRequest) SetRegionId(v string) *RunRCInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *RunRCInstancesRequest) SetResourceGroupId(v string) *RunRCInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *RunRCInstancesRequest) SetScheduledRule(v string) *RunRCInstancesRequest {
	s.ScheduledRule = &v
	return s
}

func (s *RunRCInstancesRequest) SetSecurityEnhancementStrategy(v string) *RunRCInstancesRequest {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *RunRCInstancesRequest) SetSecurityGroupId(v string) *RunRCInstancesRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *RunRCInstancesRequest) SetSecurityGroupIds(v []*string) *RunRCInstancesRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *RunRCInstancesRequest) SetSpotStrategy(v string) *RunRCInstancesRequest {
	s.SpotStrategy = &v
	return s
}

func (s *RunRCInstancesRequest) SetSupportCase(v string) *RunRCInstancesRequest {
	s.SupportCase = &v
	return s
}

func (s *RunRCInstancesRequest) SetSystemDisk(v *RunRCInstancesRequestSystemDisk) *RunRCInstancesRequest {
	s.SystemDisk = v
	return s
}

func (s *RunRCInstancesRequest) SetTag(v []*RunRCInstancesRequestTag) *RunRCInstancesRequest {
	s.Tag = v
	return s
}

func (s *RunRCInstancesRequest) SetUserData(v string) *RunRCInstancesRequest {
	s.UserData = &v
	return s
}

func (s *RunRCInstancesRequest) SetUserDataInBase64(v bool) *RunRCInstancesRequest {
	s.UserDataInBase64 = &v
	return s
}

func (s *RunRCInstancesRequest) SetVSwitchId(v string) *RunRCInstancesRequest {
	s.VSwitchId = &v
	return s
}

func (s *RunRCInstancesRequest) SetZoneId(v string) *RunRCInstancesRequest {
	s.ZoneId = &v
	return s
}

func (s *RunRCInstancesRequest) Validate() error {
	if s.CreateAckEdgeParam != nil {
		if err := s.CreateAckEdgeParam.Validate(); err != nil {
			return err
		}
	}
	if s.DataDisk != nil {
		for _, item := range s.DataDisk {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NetworkOptions != nil {
		if err := s.NetworkOptions.Validate(); err != nil {
			return err
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

type RunRCInstancesRequestCreateAckEdgeParam struct {
	// The ID of the target ACK Edge cluster.
	//
	// example:
	//
	// c463aaa89e2b84cacacfbf23c4867****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the target edge node pool in the ACK Edge cluster.
	//
	// example:
	//
	// np47e018268fb34e2289ff4c4d22b5****
	NodePoolId *string `json:"NodePoolId,omitempty" xml:"NodePoolId,omitempty"`
}

func (s RunRCInstancesRequestCreateAckEdgeParam) String() string {
	return dara.Prettify(s)
}

func (s RunRCInstancesRequestCreateAckEdgeParam) GoString() string {
	return s.String()
}

func (s *RunRCInstancesRequestCreateAckEdgeParam) GetClusterId() *string {
	return s.ClusterId
}

func (s *RunRCInstancesRequestCreateAckEdgeParam) GetNodePoolId() *string {
	return s.NodePoolId
}

func (s *RunRCInstancesRequestCreateAckEdgeParam) SetClusterId(v string) *RunRCInstancesRequestCreateAckEdgeParam {
	s.ClusterId = &v
	return s
}

func (s *RunRCInstancesRequestCreateAckEdgeParam) SetNodePoolId(v string) *RunRCInstancesRequestCreateAckEdgeParam {
	s.NodePoolId = &v
	return s
}

func (s *RunRCInstancesRequestCreateAckEdgeParam) Validate() error {
	return dara.Validate(s)
}

type RunRCInstancesRequestDataDisk struct {
	// The type of the data cloud disk. Valid values:
	//
	// - **cloud_efficiency**: ultra cloud disk.
	//
	// - **cloud_ssd**: standard SSD.
	//
	// - **cloud_essd*	- (default): ESSD.
	//
	// - **cloud_auto**: premium performance disk.
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Reserved parameter. Not supported.
	//
	// example:
	//
	// null
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The mount point of the data cloud disk.
	//
	// >This parameter is applicable only to full image (whole-machine image) scenarios. You can set this parameter to the mount point of the data cloud disk in the full image and modify the corresponding **DataDisk.Size*	- and **DataDisk.Category*	- parameters to change the cloud disk type and size of the data cloud disk in the full image.
	//
	// example:
	//
	// /dev/xvdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// Specifies whether to encrypt the cloud disk. Valid values:
	//
	// - **true**: The cloud disk is encrypted.
	//
	// - **false*	- (default): The cloud disk is not encrypted.
	//
	// example:
	//
	// false
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The performance level (PL) of the data cloud disk when it is an ESSD. For information about the performance differences of ESSDs, see [ESSD](https://help.aliyun.com/document_detail/2859916.html). Valid values:
	//
	// - **PL0**
	//
	// - **PL1*	- (default)
	//
	// - **PL2**
	//
	// - **PL3**
	//
	// > When the data cloud disk type is standard SSD, this parameter is not applicable.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of the data cloud disk. Unit: GiB. Valid values:
	//
	// - cloud_efficiency: 20 to 32,768.
	//
	// - cloud_ssd: 20 to 32,768.
	//
	// - cloud_auto: 1 to 65,536.
	//
	// - cloud_essd: The valid values depend on the value of **DataDisk.PerformanceLevel**.
	//
	//   - PL0: 1 to 65,536.
	//
	//   - PL1: 20 to 65,536.
	//
	//   - PL2: 461 to 65,536.
	//
	//   - PL3: 1,261 to 65,536.
	//
	// If the **DataDisk.SnapshotId*	- parameter is specified and the snapshot size is greater than the value of **DataDisk.Size**, the cloud disk is created with the same size as the snapshot. If the snapshot size is smaller than the value of **DataDisk.Size**, the cloud disk is created with the size specified by **DataDisk.Size**.
	//
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The snapshot used to create the data cloud disk.
	//
	// - If the snapshot size corresponding to **DataDisk.SnapshotId*	- is greater than the value of **DataDisk.Size**, the cloud disk is created with the same size as the snapshot. If the snapshot size is smaller than the value of **DataDisk.Size**, the cloud disk is created with the size specified by **DataDisk.Size**.
	//
	// - Snapshots cannot be used to create elastic ephemeral disks.
	//
	// - Snapshots created on or before July 15, 2013 cannot be used to create cloud disks.
	//
	// example:
	//
	// s-bp17441ohwka0yuh****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s RunRCInstancesRequestDataDisk) String() string {
	return dara.Prettify(s)
}

func (s RunRCInstancesRequestDataDisk) GoString() string {
	return s.String()
}

func (s *RunRCInstancesRequestDataDisk) GetCategory() *string {
	return s.Category
}

func (s *RunRCInstancesRequestDataDisk) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *RunRCInstancesRequestDataDisk) GetDevice() *string {
	return s.Device
}

func (s *RunRCInstancesRequestDataDisk) GetEncrypted() *string {
	return s.Encrypted
}

func (s *RunRCInstancesRequestDataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *RunRCInstancesRequestDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *RunRCInstancesRequestDataDisk) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *RunRCInstancesRequestDataDisk) SetCategory(v string) *RunRCInstancesRequestDataDisk {
	s.Category = &v
	return s
}

func (s *RunRCInstancesRequestDataDisk) SetDeleteWithInstance(v bool) *RunRCInstancesRequestDataDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *RunRCInstancesRequestDataDisk) SetDevice(v string) *RunRCInstancesRequestDataDisk {
	s.Device = &v
	return s
}

func (s *RunRCInstancesRequestDataDisk) SetEncrypted(v string) *RunRCInstancesRequestDataDisk {
	s.Encrypted = &v
	return s
}

func (s *RunRCInstancesRequestDataDisk) SetPerformanceLevel(v string) *RunRCInstancesRequestDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *RunRCInstancesRequestDataDisk) SetSize(v int32) *RunRCInstancesRequestDataDisk {
	s.Size = &v
	return s
}

func (s *RunRCInstancesRequestDataDisk) SetSnapshotId(v string) *RunRCInstancesRequestDataDisk {
	s.SnapshotId = &v
	return s
}

func (s *RunRCInstancesRequestDataDisk) Validate() error {
	return dara.Validate(s)
}

type RunRCInstancesRequestNetworkOptions struct {
	// Specifies whether to enable the Jumbo frame feature for the instance. Valid values:
	//
	// - **false*	- (default): Jumbo frame is disabled. The MTU of all NICs (including the primary NIC and secondary NICs) on the instance is set to 1500.
	//
	// - **true**: Jumbo frame is enabled. The MTU of all NICs (including the primary NIC and secondary NICs) on the instance is set to 8500.
	//
	// > Only specific instance types of the eighth generation or later support the Jumbo frame feature. For more information, see ECS Instance MTU.
	EnableJumboFrame *bool `json:"EnableJumboFrame,omitempty" xml:"EnableJumboFrame,omitempty"`
}

func (s RunRCInstancesRequestNetworkOptions) String() string {
	return dara.Prettify(s)
}

func (s RunRCInstancesRequestNetworkOptions) GoString() string {
	return s.String()
}

func (s *RunRCInstancesRequestNetworkOptions) GetEnableJumboFrame() *bool {
	return s.EnableJumboFrame
}

func (s *RunRCInstancesRequestNetworkOptions) SetEnableJumboFrame(v bool) *RunRCInstancesRequestNetworkOptions {
	s.EnableJumboFrame = &v
	return s
}

func (s *RunRCInstancesRequestNetworkOptions) Validate() error {
	return dara.Validate(s)
}

type RunRCInstancesRequestSystemDisk struct {
	// The type of the system cloud disk. Valid values:
	//
	// - **cloud_efficiency**: ultra cloud disk.
	//
	// - **cloud_ssd**: standard SSD.
	//
	// - **cloud_essd*	- (default): ESSD.
	//
	// - **cloud_auto**: premium performance disk.
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The performance level (PL) of the system cloud disk when it is an ESSD. For information about the performance differences of ESSDs, see [ESSD](https://help.aliyun.com/document_detail/2859916.html). Valid values:
	//
	// - **PL0**
	//
	// - **PL1*	- (default)
	//
	// - **PL2**
	//
	// - **PL3**
	//
	// > When the system cloud disk type is standard SSD, this parameter is not applicable.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of the system cloud disk. Unit: GiB. The value must be greater than or equal to the size of the image specified by the **ImageId*	- parameter. Valid values:
	//
	// - **cloud_efficiency**: 20 to 2048.
	//
	// - **cloud_ssd**: 20 to 2048.
	//
	// - **cloud_auto**: 1 to 2048.
	//
	// - **cloud_essd**: The valid values depend on the value of **SystemDisk.PerformanceLevel**.
	//
	//   - PL0: 1 to 2048.
	//
	//   - PL1: 20 to 2048.
	//
	//   - PL2: 461 to 2048.
	//
	//   - PL3: 1,261 to 2048.
	//
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s RunRCInstancesRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s RunRCInstancesRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *RunRCInstancesRequestSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *RunRCInstancesRequestSystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *RunRCInstancesRequestSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *RunRCInstancesRequestSystemDisk) SetCategory(v string) *RunRCInstancesRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *RunRCInstancesRequestSystemDisk) SetPerformanceLevel(v string) *RunRCInstancesRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *RunRCInstancesRequestSystemDisk) SetSize(v int32) *RunRCInstancesRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *RunRCInstancesRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}

type RunRCInstancesRequestTag struct {
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

func (s RunRCInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s RunRCInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *RunRCInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *RunRCInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *RunRCInstancesRequestTag) SetKey(v string) *RunRCInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *RunRCInstancesRequestTag) SetValue(v string) *RunRCInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *RunRCInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}
