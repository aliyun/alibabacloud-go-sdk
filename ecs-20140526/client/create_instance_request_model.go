// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHibernationOptions(v *CreateInstanceRequestHibernationOptions) *CreateInstanceRequest
	GetHibernationOptions() *CreateInstanceRequestHibernationOptions
	SetPrivatePoolOptions(v *CreateInstanceRequestPrivatePoolOptions) *CreateInstanceRequest
	GetPrivatePoolOptions() *CreateInstanceRequestPrivatePoolOptions
	SetSystemDisk(v *CreateInstanceRequestSystemDisk) *CreateInstanceRequest
	GetSystemDisk() *CreateInstanceRequestSystemDisk
	SetAffinity(v string) *CreateInstanceRequest
	GetAffinity() *string
	SetArn(v []*CreateInstanceRequestArn) *CreateInstanceRequest
	GetArn() []*CreateInstanceRequestArn
	SetAutoRenew(v bool) *CreateInstanceRequest
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int32) *CreateInstanceRequest
	GetAutoRenewPeriod() *int32
	SetClientToken(v string) *CreateInstanceRequest
	GetClientToken() *string
	SetClusterId(v string) *CreateInstanceRequest
	GetClusterId() *string
	SetCreditSpecification(v string) *CreateInstanceRequest
	GetCreditSpecification() *string
	SetDataDisk(v []*CreateInstanceRequestDataDisk) *CreateInstanceRequest
	GetDataDisk() []*CreateInstanceRequestDataDisk
	SetDedicatedHostId(v string) *CreateInstanceRequest
	GetDedicatedHostId() *string
	SetDeletionProtection(v bool) *CreateInstanceRequest
	GetDeletionProtection() *bool
	SetDeploymentSetGroupNo(v int32) *CreateInstanceRequest
	GetDeploymentSetGroupNo() *int32
	SetDeploymentSetId(v string) *CreateInstanceRequest
	GetDeploymentSetId() *string
	SetDescription(v string) *CreateInstanceRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateInstanceRequest
	GetDryRun() *bool
	SetHostName(v string) *CreateInstanceRequest
	GetHostName() *string
	SetHpcClusterId(v string) *CreateInstanceRequest
	GetHpcClusterId() *string
	SetHttpEndpoint(v string) *CreateInstanceRequest
	GetHttpEndpoint() *string
	SetHttpPutResponseHopLimit(v int32) *CreateInstanceRequest
	GetHttpPutResponseHopLimit() *int32
	SetHttpTokens(v string) *CreateInstanceRequest
	GetHttpTokens() *string
	SetImageFamily(v string) *CreateInstanceRequest
	GetImageFamily() *string
	SetImageId(v string) *CreateInstanceRequest
	GetImageId() *string
	SetInnerIpAddress(v string) *CreateInstanceRequest
	GetInnerIpAddress() *string
	SetInstanceChargeType(v string) *CreateInstanceRequest
	GetInstanceChargeType() *string
	SetInstanceName(v string) *CreateInstanceRequest
	GetInstanceName() *string
	SetInstanceType(v string) *CreateInstanceRequest
	GetInstanceType() *string
	SetInternetChargeType(v string) *CreateInstanceRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidthIn(v int32) *CreateInstanceRequest
	GetInternetMaxBandwidthIn() *int32
	SetInternetMaxBandwidthOut(v int32) *CreateInstanceRequest
	GetInternetMaxBandwidthOut() *int32
	SetIoOptimized(v string) *CreateInstanceRequest
	GetIoOptimized() *string
	SetKeyPairName(v string) *CreateInstanceRequest
	GetKeyPairName() *string
	SetNodeControllerId(v string) *CreateInstanceRequest
	GetNodeControllerId() *string
	SetOwnerAccount(v string) *CreateInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateInstanceRequest
	GetOwnerId() *int64
	SetPassword(v string) *CreateInstanceRequest
	GetPassword() *string
	SetPasswordInherit(v bool) *CreateInstanceRequest
	GetPasswordInherit() *bool
	SetPeriod(v int32) *CreateInstanceRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateInstanceRequest
	GetPeriodUnit() *string
	SetPrivateIpAddress(v string) *CreateInstanceRequest
	GetPrivateIpAddress() *string
	SetRamRoleName(v string) *CreateInstanceRequest
	GetRamRoleName() *string
	SetRegionId(v string) *CreateInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateInstanceRequest
	GetResourceOwnerId() *int64
	SetSecurityEnhancementStrategy(v string) *CreateInstanceRequest
	GetSecurityEnhancementStrategy() *string
	SetSecurityGroupId(v string) *CreateInstanceRequest
	GetSecurityGroupId() *string
	SetSpotDuration(v int32) *CreateInstanceRequest
	GetSpotDuration() *int32
	SetSpotInterruptionBehavior(v string) *CreateInstanceRequest
	GetSpotInterruptionBehavior() *string
	SetSpotPriceLimit(v float32) *CreateInstanceRequest
	GetSpotPriceLimit() *float32
	SetSpotStrategy(v string) *CreateInstanceRequest
	GetSpotStrategy() *string
	SetStorageSetId(v string) *CreateInstanceRequest
	GetStorageSetId() *string
	SetStorageSetPartitionNumber(v int32) *CreateInstanceRequest
	GetStorageSetPartitionNumber() *int32
	SetTag(v []*CreateInstanceRequestTag) *CreateInstanceRequest
	GetTag() []*CreateInstanceRequestTag
	SetTenancy(v string) *CreateInstanceRequest
	GetTenancy() *string
	SetUseAdditionalService(v bool) *CreateInstanceRequest
	GetUseAdditionalService() *bool
	SetUserData(v string) *CreateInstanceRequest
	GetUserData() *string
	SetVSwitchId(v string) *CreateInstanceRequest
	GetVSwitchId() *string
	SetVlanId(v string) *CreateInstanceRequest
	GetVlanId() *string
	SetZoneId(v string) *CreateInstanceRequest
	GetZoneId() *string
}

type CreateInstanceRequest struct {
	HibernationOptions *CreateInstanceRequestHibernationOptions `json:"HibernationOptions,omitempty" xml:"HibernationOptions,omitempty" type:"Struct"`
	PrivatePoolOptions *CreateInstanceRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	SystemDisk         *CreateInstanceRequestSystemDisk         `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// Specifies whether the instance on a dedicated host is associated with the dedicated host. Valid values:
	//
	// example:
	//
	// default
	Affinity *string `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	Arn []*CreateInstanceRequestArn `json:"Arn,omitempty" xml:"Arn,omitempty" type:"Repeated"`
	// Specifies whether to enable auto-renewal. This parameter takes effect only when `InstanceChargeType` is set to `PrePaid`. Valid values:
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period. This parameter is required when AutoRenew is set to True.
	//
	// example:
	//
	// 2
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cluster ID of the instance.
	//
	// example:
	//
	// c-bp67acfmxazb4p****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The running mode of the burstable instance. Valid values:
	//
	// example:
	//
	// Standard
	CreditSpecification *string `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	// The list of data disks.
	DataDisk []*CreateInstanceRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// The ID of the dedicated host.
	//
	// example:
	//
	// dh-bp67acfmxazb4p****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// The release protection attribute of the instance. Specifies whether the instance can be released from the ECS console or by calling the [DeleteInstance](https://help.aliyun.com/document_detail/25507.html) operation.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The group number of the instance in the deployment set. This parameter takes effect only when the deployment set uses the high availability group strategy (AvailabilityGroup). Valid values: 1 to 7.
	//
	// example:
	//
	// 1
	DeploymentSetGroupNo *int32 `json:"DeploymentSetGroupNo,omitempty" xml:"DeploymentSetGroupNo,omitempty"`
	// The ID of the deployment set.
	//
	// example:
	//
	// ds-bp1brhwhoqinyjd6****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The description of the instance. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// InstanceTest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform only a dry run. Valid values:
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The hostname of the Elastic Compute Service server.
	//
	// example:
	//
	// LocalHostName
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the HPC cluster to which the instance belongs.
	//
	// example:
	//
	// hpc-bp67acfmxazb4p****
	HpcClusterId *string `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	// Specifies whether to enable the access channel for instance metadata. Valid values:
	//
	// example:
	//
	// enabled
	HttpEndpoint *string `json:"HttpEndpoint,omitempty" xml:"HttpEndpoint,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// 0
	HttpPutResponseHopLimit *int32 `json:"HttpPutResponseHopLimit,omitempty" xml:"HttpPutResponseHopLimit,omitempty"`
	// Specifies whether to forcefully use the security-hardened mode (IMDSv2) to access instance metadata. Valid values:
	//
	// example:
	//
	// optional
	HttpTokens *string `json:"HttpTokens,omitempty" xml:"HttpTokens,omitempty"`
	// The name of the image family. You can set this parameter to obtain the latest available image from the specified image family to create the instance.
	//
	// example:
	//
	// hangzhou-daily-update
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The ID of the image used to start the instance. To use an Alibaba Cloud Marketplace image, you can view the `ImageId` on the image product page. If you do not specify `ImageFamily` to select the latest available image from an image family, this parameter is required.
	//
	// example:
	//
	// ubuntu_18_04_64_20G_alibase_20190624.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The internal IP address of the instance.
	//
	// example:
	//
	// ``192.168.**.**``
	InnerIpAddress *string `json:"InnerIpAddress,omitempty" xml:"InnerIpAddress,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The name of the instance. The name must be 2 to 128 characters in length and can contain letters, digits, and Unicode characters classified under the letter category (including Chinese characters). The name can also contain colons (:), underscores (_), periods (.), and hyphens (-). If this parameter is not specified, the default value is the instance ID.
	//
	// example:
	//
	// k8s-node-[1,4]-alibabacloud
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The billing method for network usage. Valid values:
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum inbound public bandwidth, in Mbit/s. Valid values:
	//
	// example:
	//
	// 50
	InternetMaxBandwidthIn *int32 `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	// The maximum outbound public bandwidth, in Mbit/s. Valid values: 0 to 100.
	//
	// example:
	//
	// 5
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// Specifies whether the instance is an I/O optimized instance. The I/O optimization improves instance performance. Valid values:
	//
	// The default value for [retired instance types](https://help.aliyun.com/document_detail/55263.html) is none.
	//
	// example:
	//
	// optimized
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// KeyPairTestName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// null
	NodeControllerId *string `json:"NodeControllerId,omitempty" xml:"NodeControllerId,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password of the instance. The password must be 8 to 30 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. The following special characters are supported:
	//
	// example:
	//
	// TestEcs123!
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether to use the password preset in the image. When you use this parameter, the Password parameter must be empty, and you must make sure that the image has a password configured.
	//
	// example:
	//
	// false
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// The subscription period of the instance. The unit is specified by `PeriodUnit`. This parameter takes effect and is required only when `InstanceChargeType` is set to `PrePaid`. If `DedicatedHostId` is specified, the value of this parameter cannot exceed the subscription period of the dedicated host. Valid values:
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription period. Valid values:
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The private IP address of the instance. The IP address must be an available address within the CIDR block of the specified vSwitch (VSwitchId).
	//
	// example:
	//
	// 172.16.236.*
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The name of the instance RAM role. You can call the RAM API [ListRoles](https://help.aliyun.com/document_detail/28713.html) to query the instance RAM roles that you have created.
	//
	// example:
	//
	// RAMTestName
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The region ID of the instance. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent list of regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the enterprise resource group to which the instance belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable security hardening. Valid values:
	//
	// example:
	//
	// Active
	SecurityEnhancementStrategy *string `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	// The ID of the security group to which the new instance belongs.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The protection period of the spot instance, in hours. Default value: 1. Valid values:
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The interruption pattern of the spot instance. Valid values:
	//
	// example:
	//
	// Terminate
	SpotInterruptionBehavior *string `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	// The maximum hourly price of the instance. A maximum of three decimal places is supported. This parameter takes effect only when `SpotStrategy` is set to `SpotWithPriceLimit`.
	//
	// example:
	//
	// 0.98
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The bidding policy for the instance. This parameter takes effect only when `InstanceChargeType` is set to `PostPaid`. Valid values:
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The ID of the storage set.
	//
	// example:
	//
	// ss-bp1j4i2jdf3owlhe****
	StorageSetId *string `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
	// The maximum number of partitions in the storage set. Valid values: greater than or equal to 2.
	//
	// example:
	//
	// 2
	StorageSetPartitionNumber *int32 `json:"StorageSetPartitionNumber,omitempty" xml:"StorageSetPartitionNumber,omitempty"`
	// The tags.
	Tag []*CreateInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies whether to create the instance on a dedicated host. Valid values:
	//
	// example:
	//
	// default
	Tenancy *string `json:"Tenancy,omitempty" xml:"Tenancy,omitempty"`
	// Specifies whether to use the virtual machine system configuration provided by Alibaba Cloud (Windows: NTP and KMS. Linux: NTP and YUM).
	//
	// example:
	//
	// true
	UseAdditionalService *bool `json:"UseAdditionalService,omitempty" xml:"UseAdditionalService,omitempty"`
	// The instance user data. The data must be encoded in Base64. The raw data can be up to 32 KB in size.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The vSwitch ID. This parameter is required when you create a VPC-connected instance. You can call [DescribeVSwitches](https://help.aliyun.com/document_detail/35748.html) to query created vSwitches.
	//
	// example:
	//
	// vsw-bp1s5fnvk4gn2tws0****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The virtual local area network ID.
	//
	// example:
	//
	// 10
	VlanId *string `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
	// The zone ID of the instance. For more information, call [DescribeZones](https://help.aliyun.com/document_detail/25610.html) to query the zone list.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetHibernationOptions() *CreateInstanceRequestHibernationOptions {
	return s.HibernationOptions
}

func (s *CreateInstanceRequest) GetPrivatePoolOptions() *CreateInstanceRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *CreateInstanceRequest) GetSystemDisk() *CreateInstanceRequestSystemDisk {
	return s.SystemDisk
}

func (s *CreateInstanceRequest) GetAffinity() *string {
	return s.Affinity
}

func (s *CreateInstanceRequest) GetArn() []*CreateInstanceRequestArn {
	return s.Arn
}

func (s *CreateInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateInstanceRequest) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *CreateInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateInstanceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateInstanceRequest) GetCreditSpecification() *string {
	return s.CreditSpecification
}

func (s *CreateInstanceRequest) GetDataDisk() []*CreateInstanceRequestDataDisk {
	return s.DataDisk
}

func (s *CreateInstanceRequest) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *CreateInstanceRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *CreateInstanceRequest) GetDeploymentSetGroupNo() *int32 {
	return s.DeploymentSetGroupNo
}

func (s *CreateInstanceRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *CreateInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateInstanceRequest) GetHostName() *string {
	return s.HostName
}

func (s *CreateInstanceRequest) GetHpcClusterId() *string {
	return s.HpcClusterId
}

func (s *CreateInstanceRequest) GetHttpEndpoint() *string {
	return s.HttpEndpoint
}

func (s *CreateInstanceRequest) GetHttpPutResponseHopLimit() *int32 {
	return s.HttpPutResponseHopLimit
}

func (s *CreateInstanceRequest) GetHttpTokens() *string {
	return s.HttpTokens
}

func (s *CreateInstanceRequest) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *CreateInstanceRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateInstanceRequest) GetInnerIpAddress() *string {
	return s.InnerIpAddress
}

func (s *CreateInstanceRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateInstanceRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateInstanceRequest) GetInternetMaxBandwidthIn() *int32 {
	return s.InternetMaxBandwidthIn
}

func (s *CreateInstanceRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *CreateInstanceRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *CreateInstanceRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateInstanceRequest) GetNodeControllerId() *string {
	return s.NodeControllerId
}

func (s *CreateInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateInstanceRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateInstanceRequest) GetPasswordInherit() *bool {
	return s.PasswordInherit
}

func (s *CreateInstanceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateInstanceRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateInstanceRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *CreateInstanceRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *CreateInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateInstanceRequest) GetSecurityEnhancementStrategy() *string {
	return s.SecurityEnhancementStrategy
}

func (s *CreateInstanceRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateInstanceRequest) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *CreateInstanceRequest) GetSpotInterruptionBehavior() *string {
	return s.SpotInterruptionBehavior
}

func (s *CreateInstanceRequest) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *CreateInstanceRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *CreateInstanceRequest) GetStorageSetId() *string {
	return s.StorageSetId
}

func (s *CreateInstanceRequest) GetStorageSetPartitionNumber() *int32 {
	return s.StorageSetPartitionNumber
}

func (s *CreateInstanceRequest) GetTag() []*CreateInstanceRequestTag {
	return s.Tag
}

func (s *CreateInstanceRequest) GetTenancy() *string {
	return s.Tenancy
}

func (s *CreateInstanceRequest) GetUseAdditionalService() *bool {
	return s.UseAdditionalService
}

func (s *CreateInstanceRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateInstanceRequest) GetVlanId() *string {
	return s.VlanId
}

func (s *CreateInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateInstanceRequest) SetHibernationOptions(v *CreateInstanceRequestHibernationOptions) *CreateInstanceRequest {
	s.HibernationOptions = v
	return s
}

func (s *CreateInstanceRequest) SetPrivatePoolOptions(v *CreateInstanceRequestPrivatePoolOptions) *CreateInstanceRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *CreateInstanceRequest) SetSystemDisk(v *CreateInstanceRequestSystemDisk) *CreateInstanceRequest {
	s.SystemDisk = v
	return s
}

func (s *CreateInstanceRequest) SetAffinity(v string) *CreateInstanceRequest {
	s.Affinity = &v
	return s
}

func (s *CreateInstanceRequest) SetArn(v []*CreateInstanceRequestArn) *CreateInstanceRequest {
	s.Arn = v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenew(v bool) *CreateInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenewPeriod(v int32) *CreateInstanceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateInstanceRequest) SetClientToken(v string) *CreateInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateInstanceRequest) SetClusterId(v string) *CreateInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateInstanceRequest) SetCreditSpecification(v string) *CreateInstanceRequest {
	s.CreditSpecification = &v
	return s
}

func (s *CreateInstanceRequest) SetDataDisk(v []*CreateInstanceRequestDataDisk) *CreateInstanceRequest {
	s.DataDisk = v
	return s
}

func (s *CreateInstanceRequest) SetDedicatedHostId(v string) *CreateInstanceRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *CreateInstanceRequest) SetDeletionProtection(v bool) *CreateInstanceRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CreateInstanceRequest) SetDeploymentSetGroupNo(v int32) *CreateInstanceRequest {
	s.DeploymentSetGroupNo = &v
	return s
}

func (s *CreateInstanceRequest) SetDeploymentSetId(v string) *CreateInstanceRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *CreateInstanceRequest) SetDescription(v string) *CreateInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateInstanceRequest) SetDryRun(v bool) *CreateInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateInstanceRequest) SetHostName(v string) *CreateInstanceRequest {
	s.HostName = &v
	return s
}

func (s *CreateInstanceRequest) SetHpcClusterId(v string) *CreateInstanceRequest {
	s.HpcClusterId = &v
	return s
}

func (s *CreateInstanceRequest) SetHttpEndpoint(v string) *CreateInstanceRequest {
	s.HttpEndpoint = &v
	return s
}

func (s *CreateInstanceRequest) SetHttpPutResponseHopLimit(v int32) *CreateInstanceRequest {
	s.HttpPutResponseHopLimit = &v
	return s
}

func (s *CreateInstanceRequest) SetHttpTokens(v string) *CreateInstanceRequest {
	s.HttpTokens = &v
	return s
}

func (s *CreateInstanceRequest) SetImageFamily(v string) *CreateInstanceRequest {
	s.ImageFamily = &v
	return s
}

func (s *CreateInstanceRequest) SetImageId(v string) *CreateInstanceRequest {
	s.ImageId = &v
	return s
}

func (s *CreateInstanceRequest) SetInnerIpAddress(v string) *CreateInstanceRequest {
	s.InnerIpAddress = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceChargeType(v string) *CreateInstanceRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceType(v string) *CreateInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateInstanceRequest) SetInternetChargeType(v string) *CreateInstanceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetInternetMaxBandwidthIn(v int32) *CreateInstanceRequest {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *CreateInstanceRequest) SetInternetMaxBandwidthOut(v int32) *CreateInstanceRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *CreateInstanceRequest) SetIoOptimized(v string) *CreateInstanceRequest {
	s.IoOptimized = &v
	return s
}

func (s *CreateInstanceRequest) SetKeyPairName(v string) *CreateInstanceRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateInstanceRequest) SetNodeControllerId(v string) *CreateInstanceRequest {
	s.NodeControllerId = &v
	return s
}

func (s *CreateInstanceRequest) SetOwnerAccount(v string) *CreateInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateInstanceRequest) SetOwnerId(v int64) *CreateInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateInstanceRequest) SetPassword(v string) *CreateInstanceRequest {
	s.Password = &v
	return s
}

func (s *CreateInstanceRequest) SetPasswordInherit(v bool) *CreateInstanceRequest {
	s.PasswordInherit = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriod(v int32) *CreateInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriodUnit(v string) *CreateInstanceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateInstanceRequest) SetPrivateIpAddress(v string) *CreateInstanceRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateInstanceRequest) SetRamRoleName(v string) *CreateInstanceRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateInstanceRequest) SetRegionId(v string) *CreateInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceGroupId(v string) *CreateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceOwnerAccount(v string) *CreateInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceOwnerId(v int64) *CreateInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateInstanceRequest) SetSecurityEnhancementStrategy(v string) *CreateInstanceRequest {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *CreateInstanceRequest) SetSecurityGroupId(v string) *CreateInstanceRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetSpotDuration(v int32) *CreateInstanceRequest {
	s.SpotDuration = &v
	return s
}

func (s *CreateInstanceRequest) SetSpotInterruptionBehavior(v string) *CreateInstanceRequest {
	s.SpotInterruptionBehavior = &v
	return s
}

func (s *CreateInstanceRequest) SetSpotPriceLimit(v float32) *CreateInstanceRequest {
	s.SpotPriceLimit = &v
	return s
}

func (s *CreateInstanceRequest) SetSpotStrategy(v string) *CreateInstanceRequest {
	s.SpotStrategy = &v
	return s
}

func (s *CreateInstanceRequest) SetStorageSetId(v string) *CreateInstanceRequest {
	s.StorageSetId = &v
	return s
}

func (s *CreateInstanceRequest) SetStorageSetPartitionNumber(v int32) *CreateInstanceRequest {
	s.StorageSetPartitionNumber = &v
	return s
}

func (s *CreateInstanceRequest) SetTag(v []*CreateInstanceRequestTag) *CreateInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateInstanceRequest) SetTenancy(v string) *CreateInstanceRequest {
	s.Tenancy = &v
	return s
}

func (s *CreateInstanceRequest) SetUseAdditionalService(v bool) *CreateInstanceRequest {
	s.UseAdditionalService = &v
	return s
}

func (s *CreateInstanceRequest) SetUserData(v string) *CreateInstanceRequest {
	s.UserData = &v
	return s
}

func (s *CreateInstanceRequest) SetVSwitchId(v string) *CreateInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateInstanceRequest) SetVlanId(v string) *CreateInstanceRequest {
	s.VlanId = &v
	return s
}

func (s *CreateInstanceRequest) SetZoneId(v string) *CreateInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
	if s.HibernationOptions != nil {
		if err := s.HibernationOptions.Validate(); err != nil {
			return err
		}
	}
	if s.PrivatePoolOptions != nil {
		if err := s.PrivatePoolOptions.Validate(); err != nil {
			return err
		}
	}
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
			return err
		}
	}
	if s.Arn != nil {
		for _, item := range s.Arn {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type CreateInstanceRequestHibernationOptions struct {
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// false
	Configured *bool `json:"Configured,omitempty" xml:"Configured,omitempty"`
}

func (s CreateInstanceRequestHibernationOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestHibernationOptions) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestHibernationOptions) GetConfigured() *bool {
	return s.Configured
}

func (s *CreateInstanceRequestHibernationOptions) SetConfigured(v bool) *CreateInstanceRequestHibernationOptions {
	s.Configured = &v
	return s
}

func (s *CreateInstanceRequestHibernationOptions) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestPrivatePoolOptions struct {
	// The ID of the private pool. The ID of an elasticity assurance or capacity reservation.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The private pool option for the instance launch. A private pool is generated when an elasticity assurance or capacity reservation takes effect. You can select a private pool when you start an instance. Valid values:
	//
	// example:
	//
	// Open
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
}

func (s CreateInstanceRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestPrivatePoolOptions) GetId() *string {
	return s.Id
}

func (s *CreateInstanceRequestPrivatePoolOptions) GetMatchCriteria() *string {
	return s.MatchCriteria
}

func (s *CreateInstanceRequestPrivatePoolOptions) SetId(v string) *CreateInstanceRequestPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *CreateInstanceRequestPrivatePoolOptions) SetMatchCriteria(v string) *CreateInstanceRequestPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *CreateInstanceRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestSystemDisk struct {
	// The category of the system disk. Valid values:
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The description of the system disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// TestDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the system disk. The name must be 2 to 128 characters in length and can contain letters, digits, and Unicode characters classified under the letter category (including Chinese characters). The name can also contain colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// SystemDiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The performance level (PL) of the enterprise SSD (ESSD) used as the system disk. Settings for the performance level. If the system disk is a standard SSD, this parameter is ignored. Valid values:
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of the system disk, in GiB. Valid values:
	//
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the dedicated block storage cluster. To use a disk in a dedicated block storage cluster as the system disk when you create an ECS instance, specify this parameter.
	//
	// example:
	//
	// dbsc-j5e1sf2vaf5he8m2****
	StorageClusterId *string `json:"StorageClusterId,omitempty" xml:"StorageClusterId,omitempty"`
}

func (s CreateInstanceRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateInstanceRequestSystemDisk) GetDescription() *string {
	return s.Description
}

func (s *CreateInstanceRequestSystemDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *CreateInstanceRequestSystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateInstanceRequestSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateInstanceRequestSystemDisk) GetStorageClusterId() *string {
	return s.StorageClusterId
}

func (s *CreateInstanceRequestSystemDisk) SetCategory(v string) *CreateInstanceRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *CreateInstanceRequestSystemDisk) SetDescription(v string) *CreateInstanceRequestSystemDisk {
	s.Description = &v
	return s
}

func (s *CreateInstanceRequestSystemDisk) SetDiskName(v string) *CreateInstanceRequestSystemDisk {
	s.DiskName = &v
	return s
}

func (s *CreateInstanceRequestSystemDisk) SetPerformanceLevel(v string) *CreateInstanceRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateInstanceRequestSystemDisk) SetSize(v int32) *CreateInstanceRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *CreateInstanceRequestSystemDisk) SetStorageClusterId(v string) *CreateInstanceRequestSystemDisk {
	s.StorageClusterId = &v
	return s
}

func (s *CreateInstanceRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestArn struct {
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// 1234567890
	AssumeRoleFor *int64 `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// Primary
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// acs:ram::123456789012****:role/adminrole
	Rolearn *string `json:"Rolearn,omitempty" xml:"Rolearn,omitempty"`
}

func (s CreateInstanceRequestArn) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestArn) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestArn) GetAssumeRoleFor() *int64 {
	return s.AssumeRoleFor
}

func (s *CreateInstanceRequestArn) GetRoleType() *string {
	return s.RoleType
}

func (s *CreateInstanceRequestArn) GetRolearn() *string {
	return s.Rolearn
}

func (s *CreateInstanceRequestArn) SetAssumeRoleFor(v int64) *CreateInstanceRequestArn {
	s.AssumeRoleFor = &v
	return s
}

func (s *CreateInstanceRequestArn) SetRoleType(v string) *CreateInstanceRequestArn {
	s.RoleType = &v
	return s
}

func (s *CreateInstanceRequestArn) SetRolearn(v string) *CreateInstanceRequestArn {
	s.Rolearn = &v
	return s
}

func (s *CreateInstanceRequestArn) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestDataDisk struct {
	// The category of data disk N. Valid values:
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to release the data disk when the instance is released.
	//
	// example:
	//
	// true
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The description of the data disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// TestDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mount point of the data disk.
	//
	// example:
	//
	// /dev/xvdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The name of the data disk. The name must be 2 to 128 characters in length and can contain letters, digits, and Unicode characters classified under the letter category (including Chinese characters). The name can also contain colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// DataDiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	// Specifies whether to encrypt data disk N.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The Key Management Service (KMS) key ID for the disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The performance level (PL) of the enterprise SSD used as a data disk. The value of N must be the same as that in `DataDisk.N.Category=cloud_essd`. Settings for the performance level. If the data disk is a standard SSD, this parameter is ignored. Valid values:
	//
	// example:
	//
	// PL2
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of data disk N, in GiB. Valid values of N: 1 to 16. Valid values:
	//
	// example:
	//
	// 2000
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The snapshot ID used to create data disk N. Valid values of N: 1 to 16.
	//
	// example:
	//
	// s-bp17441ohwka0yuh****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The ID of the dedicated block storage cluster. If you want to use disks in a dedicated block storage cluster as data disks when you create an ECS instance, set this parameter.
	//
	// example:
	//
	// dbsc-j5e1sf2vaf5he8m2****
	StorageClusterId *string `json:"StorageClusterId,omitempty" xml:"StorageClusterId,omitempty"`
}

func (s CreateInstanceRequestDataDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestDataDisk) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestDataDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateInstanceRequestDataDisk) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *CreateInstanceRequestDataDisk) GetDescription() *string {
	return s.Description
}

func (s *CreateInstanceRequestDataDisk) GetDevice() *string {
	return s.Device
}

func (s *CreateInstanceRequestDataDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *CreateInstanceRequestDataDisk) GetEncryptAlgorithm() *string {
	return s.EncryptAlgorithm
}

func (s *CreateInstanceRequestDataDisk) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *CreateInstanceRequestDataDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *CreateInstanceRequestDataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateInstanceRequestDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateInstanceRequestDataDisk) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateInstanceRequestDataDisk) GetStorageClusterId() *string {
	return s.StorageClusterId
}

func (s *CreateInstanceRequestDataDisk) SetCategory(v string) *CreateInstanceRequestDataDisk {
	s.Category = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetDeleteWithInstance(v bool) *CreateInstanceRequestDataDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetDescription(v string) *CreateInstanceRequestDataDisk {
	s.Description = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetDevice(v string) *CreateInstanceRequestDataDisk {
	s.Device = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetDiskName(v string) *CreateInstanceRequestDataDisk {
	s.DiskName = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetEncryptAlgorithm(v string) *CreateInstanceRequestDataDisk {
	s.EncryptAlgorithm = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetEncrypted(v bool) *CreateInstanceRequestDataDisk {
	s.Encrypted = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetKMSKeyId(v string) *CreateInstanceRequestDataDisk {
	s.KMSKeyId = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetPerformanceLevel(v string) *CreateInstanceRequestDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetSize(v int32) *CreateInstanceRequestDataDisk {
	s.Size = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetSnapshotId(v string) *CreateInstanceRequestDataDisk {
	s.SnapshotId = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetStorageClusterId(v string) *CreateInstanceRequestDataDisk {
	s.StorageClusterId = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestTag struct {
	// The tag key of the instance, disk, and primary ENI. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the instance, disk, and primary ENI. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateInstanceRequestTag) SetKey(v string) *CreateInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateInstanceRequestTag) SetValue(v string) *CreateInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}
