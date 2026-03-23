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
	// ACU type
	//
	// example:
	//
	// gn8is
	AcuType *string `json:"AcuType,omitempty" xml:"AcuType,omitempty"`
	// The number of RDS Custom instances that you want to create. The parameter is available if you want to create multiple RDS Custom instances at a time.
	//
	// Valid values: **1*	- to **10**. Default value: **1**.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Specifies whether to enable the automatic payment feature. Valid values:
	//
	// 	- **true*	- (default): enables the feature. Make sure that your account balance is sufficient.
	//
	// 	- **false**: disables the feature. An unpaid order is generated.
	//
	// >  If your account balance is insufficient, you can set the AutoPay parameter to false. In this case, an unpaid order is generated. You can complete the payment in the Expenses and Costs console.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Specifies whether to automatically use a coupon. Valid values:
	//
	// 	- **true*	- (default): Yes.
	//
	// 	- **false**: No.
	//
	// > If you use a coupon and later decrease the quota, the amount offset by the coupon will not be refunded.
	//
	// example:
	//
	// true
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Information about the ACK Edge cluster.
	CreateAckEdgeParam *RunRCInstancesRequestCreateAckEdgeParam `json:"CreateAckEdgeParam,omitempty" xml:"CreateAckEdgeParam,omitempty" type:"Struct"`
	// Reserved parameter. Not supported currently.
	//
	// example:
	//
	// None
	CreateExtraParam *string `json:"CreateExtraParam,omitempty" xml:"CreateExtraParam,omitempty"`
	// Specifies whether the instance can be added to an ACK cluster. When this parameter is set to **1**, the created instance can be added to an ACK cluster by invoking the **AttachRCInstances*	- API operation, enabling efficient management of container applications.
	//
	// - **1**: Yes.
	//
	// - **0*	- (default): No.
	//
	// example:
	//
	// 0
	CreateMode *string `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	// The information about the data disks.
	DataDisk []*RunRCInstancesRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// Specifies whether to enable release protection. Valid values:
	//
	// 	- **true**: Enabled
	//
	// 	- **false*	- (default): Disabled
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
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, service limits, and insufficient inventory errors.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, the instance is directly created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Hostname of the instance (2–64 characters).
	//
	// - Multiple segments separated by periods (.) are supported. Each segment can contain uppercase and lowercase English letters, digits, and hyphens (-).
	//
	// - A period (.) or hyphen (-) cannot appear at the beginning or end of a segment, nor can two periods or hyphens appear consecutively.
	//
	// example:
	//
	// testHost1
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the image used by the instance.
	//
	// example:
	//
	// image-dsvjzw2ii8n4******
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The billing method of the instance. Set the value to **Prepaid**, which indicates the subscription billing method.
	//
	// example:
	//
	// Prepaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance name.
	//
	// example:
	//
	// rc-node-[99,1]-rchost
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type. For more information about the instance types that are supported by RDS Custom instances, see [Instance types for RDS Custom instances](https://help.aliyun.com/document_detail/2844823.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql.i8.large.2cm
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// null
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// 0
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// null
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The name of the AccessKey pair. You can specify only one name.
	//
	// example:
	//
	// dell5502
	KeyPairName    *string                              `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	NetworkOptions *RunRCInstancesRequestNetworkOptions `json:"NetworkOptions,omitempty" xml:"NetworkOptions,omitempty" type:"Struct"`
	// The password of the account that is used to log on to the instance.
	//
	// example:
	//
	// TestRDS123!
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether to use the password preset in the image. When this parameter is used, the Password parameter must be empty, and you must ensure that the selected image has a password already configured. Default value: false.
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// The subscription duration of the instance. Default value: **1**.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// 	- **Year**
	//
	// 	- **Month*	- (default)
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// example:
	//
	// `10.1.**.**`
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The coupon code.
	//
	// example:
	//
	// 72329885****
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
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
	// Scheduled elasticity rule
	//
	// example:
	//
	// {"rule":[{"beginTime":"09:00","endTime":"17:00","acu":4}]}
	ScheduledRule *string `json:"ScheduledRule,omitempty" xml:"ScheduledRule,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// null
	SecurityEnhancementStrategy *string `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	// The ID of the security group to which you want to add the new instance. Instances in the same security group can communicate with each other. The maximum number of instances allowed in a security group varies based on the type of the security group. For more information, see the "Security group limits" section in [Limits](https://help.aliyun.com/document_detail/25412.html).
	//
	// >  The network type of the instance is determined by the security group specified by the SecurityGroupId parameter. For example, if the network type of the specified security group is VPC, the instance is a VPC-type instance. In this case, you must specify the VSwitchId parameter.
	//
	// example:
	//
	// sg-uf6av412xaxixu******
	SecurityGroupId  *string   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The spot strategy for pay-as-you-go instances. This parameter takes effect only when the **InstanceChargeType*	- parameter is set to **PostPaid**. Valid values:
	//
	// - **NoSpot**: Normal pay-as-you-go instance.
	//
	// - **SpotAsPriceGo**: The system automatically bids based on the current market price.
	//
	// Default value: **NoSpot**.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The deployment type of RDS Custom. Valid values:
	//
	// - **eni**: Dual network interface cards.
	//
	// - **edge**: Point of presence (POP) node pool.
	//
	// - **share**: VPC.
	//
	// example:
	//
	// share
	SupportCase *string `json:"SupportCase,omitempty" xml:"SupportCase,omitempty"`
	// The specification of the system disk.
	SystemDisk *RunRCInstancesRequestSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// The list of tags.
	Tag []*RunRCInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The instance user data. The raw data can be up to 32 KB in size.
	//
	// Do not pass sensitive information, such as passwords and private keys, in plaintext. If you must pass such information, encrypt it first, encode it in Base64, and then transmit it. Decrypt and use it inside the instance. The following is an example of converting a script into a Base64-encoded string:
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
	// Specifies whether custom data is Base64-encoded.
	//
	// - **true**: Yes.
	//
	// - **false*	- (default): No.
	//
	// example:
	//
	// true
	UserDataInBase64 *bool `json:"UserDataInBase64,omitempty" xml:"UserDataInBase64,omitempty"`
	// The vSwitch ID of the instance. You must specify this parameter when you create an instance of the virtual private cloud (VPC) type. The specified vSwitch and security group must belong to the same VPC.
	//
	// >  If you specify the VSwitchId parameter, the zone specified by the ZoneId parameter must be the same as the zone in which the specified vSwitch resides. You can leave the ZoneId parameter empty. In this case, the system uses the zone in which the specified vSwitch resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-2vcd61ngm890sk****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the instance. You can call the DescribeZones operation to query the zone IDs.
	//
	// >  If you specify the VSwitchId parameter, the zone specified by the ZoneId parameter must be the same as the zone in which the specified vSwitch resides. You can leave the ZoneId parameter empty. In this case, the system uses the zone in which the specified vSwitch resides.
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
	// null
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The mount target of the data disk.
	//
	// > This parameter is used only in the full image (entire machine image) scenario. You can set this parameter to the mount target of the data disk in the full image and modify the corresponding **DataDisk.Size*	- and **DataDisk.Category*	- parameters to change the disk category and size of the data disk in the full image.
	//
	// example:
	//
	// /dev/xvdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// Specifies whether to encrypt the cloud disk. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of the data disk. Unit: GiB.
	//
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The snapshot used to create the data disk.
	//
	// - If the snapshot size corresponding to **DataDisk.SnapshotId*	- is greater than the value of **DataDisk.Size**, the created disk size matches the snapshot size. If the snapshot size is smaller than the value of **DataDisk.Size**, the created disk size equals the value of **DataDisk.Size**.
	//
	// - Creating an elastic ephemeral disk from a snapshot is not supported.
	//
	// - Snapshots created on or before July 15, 2013 cannot be used to create disks.
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
	// The type of the system disk. Set the value to **cloud_essd**, which indicates ESSDs.
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Performance level when the system disk is an enterprise SSD (ESSD). For details about performance differences among ESSD types, see [enterprise SSD (ESSD)](https://help.aliyun.com/document_detail/2859916.html). Valid values:
	//
	// - **PL0*	-
	//
	// - **PL1*	- (default)
	//
	// - **PL2*	-
	//
	// - **PL3**
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of the system disk. Unit: GiB. Only performance level 1 (PL1) ESSDs are supported. Valid values: 20 to 2048.
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
	// The tag key. You can create up to N tag keys at the same time, where N ranges from **1 to 20**. Empty strings are not allowed.
	//
	// example:
	//
	// Testkey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value corresponding to the tag key. You can create up to N tag values at the same time, where N ranges from **1 to 20**. Empty strings are allowed.
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
