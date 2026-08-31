// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkInterfaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *CreateNetworkInterfaceRequest
	GetBusinessType() *string
	SetClientToken(v string) *CreateNetworkInterfaceRequest
	GetClientToken() *string
	SetConnectionTrackingConfiguration(v *CreateNetworkInterfaceRequestConnectionTrackingConfiguration) *CreateNetworkInterfaceRequest
	GetConnectionTrackingConfiguration() *CreateNetworkInterfaceRequestConnectionTrackingConfiguration
	SetDeleteOnRelease(v bool) *CreateNetworkInterfaceRequest
	GetDeleteOnRelease() *bool
	SetDescription(v string) *CreateNetworkInterfaceRequest
	GetDescription() *string
	SetEnablePrimaryIPv6(v bool) *CreateNetworkInterfaceRequest
	GetEnablePrimaryIPv6() *bool
	SetEnhancedNetwork(v *CreateNetworkInterfaceRequestEnhancedNetwork) *CreateNetworkInterfaceRequest
	GetEnhancedNetwork() *CreateNetworkInterfaceRequestEnhancedNetwork
	SetInstanceType(v string) *CreateNetworkInterfaceRequest
	GetInstanceType() *string
	SetIpv4Prefix(v []*string) *CreateNetworkInterfaceRequest
	GetIpv4Prefix() []*string
	SetIpv4PrefixCount(v int32) *CreateNetworkInterfaceRequest
	GetIpv4PrefixCount() *int32
	SetIpv6Address(v []*string) *CreateNetworkInterfaceRequest
	GetIpv6Address() []*string
	SetIpv6AddressCount(v int32) *CreateNetworkInterfaceRequest
	GetIpv6AddressCount() *int32
	SetIpv6Prefix(v []*string) *CreateNetworkInterfaceRequest
	GetIpv6Prefix() []*string
	SetIpv6PrefixCount(v int32) *CreateNetworkInterfaceRequest
	GetIpv6PrefixCount() *int32
	SetNetworkInterfaceName(v string) *CreateNetworkInterfaceRequest
	GetNetworkInterfaceName() *string
	SetNetworkInterfaceTrafficConfig(v *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) *CreateNetworkInterfaceRequest
	GetNetworkInterfaceTrafficConfig() *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig
	SetNetworkInterfaceTrafficMode(v string) *CreateNetworkInterfaceRequest
	GetNetworkInterfaceTrafficMode() *string
	SetOwnerAccount(v string) *CreateNetworkInterfaceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateNetworkInterfaceRequest
	GetOwnerId() *int64
	SetPrimaryIpAddress(v string) *CreateNetworkInterfaceRequest
	GetPrimaryIpAddress() *string
	SetPrivateIpAddress(v []*string) *CreateNetworkInterfaceRequest
	GetPrivateIpAddress() []*string
	SetQueueNumber(v int32) *CreateNetworkInterfaceRequest
	GetQueueNumber() *int32
	SetQueuePairNumber(v int32) *CreateNetworkInterfaceRequest
	GetQueuePairNumber() *int32
	SetRegionId(v string) *CreateNetworkInterfaceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateNetworkInterfaceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateNetworkInterfaceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateNetworkInterfaceRequest
	GetResourceOwnerId() *int64
	SetRxQueueSize(v int32) *CreateNetworkInterfaceRequest
	GetRxQueueSize() *int32
	SetSecondaryPrivateIpAddressCount(v int32) *CreateNetworkInterfaceRequest
	GetSecondaryPrivateIpAddressCount() *int32
	SetSecurityGroupId(v string) *CreateNetworkInterfaceRequest
	GetSecurityGroupId() *string
	SetSecurityGroupIds(v []*string) *CreateNetworkInterfaceRequest
	GetSecurityGroupIds() []*string
	SetSourceDestCheck(v bool) *CreateNetworkInterfaceRequest
	GetSourceDestCheck() *bool
	SetTag(v []*CreateNetworkInterfaceRequestTag) *CreateNetworkInterfaceRequest
	GetTag() []*CreateNetworkInterfaceRequestTag
	SetTxQueueSize(v int32) *CreateNetworkInterfaceRequest
	GetTxQueueSize() *int32
	SetVSwitchId(v string) *CreateNetworkInterfaceRequest
	GetVSwitchId() *string
	SetVisible(v bool) *CreateNetworkInterfaceRequest
	GetVisible() *bool
}

type CreateNetworkInterfaceRequest struct {
	// > This parameter is deprecated.
	//
	// example:
	//
	// null
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The network connectivity tracking configuration.
	//
	// Before you use this parameter, read [Connection timeout management](https://help.aliyun.com/document_detail/2865958.html).
	ConnectionTrackingConfiguration *CreateNetworkInterfaceRequestConnectionTrackingConfiguration `json:"ConnectionTrackingConfiguration,omitempty" xml:"ConnectionTrackingConfiguration,omitempty" type:"Struct"`
	// Specifies whether to retain the ENI when the associated instance is released. Valid values:
	//
	// - true: does not retain the ENI.
	//
	// - false: retains the ENI.
	//
	// example:
	//
	// true
	DeleteOnRelease *bool `json:"DeleteOnRelease,omitempty" xml:"DeleteOnRelease,omitempty"`
	// The description of the network interface controller (NIC). The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// Default value: empty.
	//
	// example:
	//
	// testDescription
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnablePrimaryIPv6 *bool   `json:"EnablePrimaryIPv6,omitempty" xml:"EnablePrimaryIPv6,omitempty"`
	// > This parameter is not publicly available.
	EnhancedNetwork *CreateNetworkInterfaceRequestEnhancedNetwork `json:"EnhancedNetwork,omitempty" xml:"EnhancedNetwork,omitempty" type:"Struct"`
	// The type of the Elastic Network Interface (ENI). Valid values:
	//
	// - Secondary: secondary ENI.
	//
	// - Trunk: trunk network interface controller (NIC) (in invitational preview).
	//
	// Default value: Secondary.
	//
	// example:
	//
	// Secondary
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// One or more IPv4 prefixes to assign to the network interface controller (NIC). Valid values of N: 1 to 10.
	//
	// > If you want to set IPv4 prefixes for the network interface controller (NIC), you must set either the parameter Ipv4Prefix.N or the parameter Ipv4PrefixCount but not both.
	Ipv4Prefix []*string `json:"Ipv4Prefix,omitempty" xml:"Ipv4Prefix,omitempty" type:"Repeated"`
	// The number of IPv4 prefixes to assign to the network interface controller (NIC). Valid values: 1 to 10.
	//
	// > If you want to set IPv4 prefixes for the network interface controller (NIC), you must set either the parameter Ipv4Prefix.N or the parameter Ipv4PrefixCount but not both.
	//
	// example:
	//
	// 1
	Ipv4PrefixCount *int32 `json:"Ipv4PrefixCount,omitempty" xml:"Ipv4PrefixCount,omitempty"`
	// One or more IPv6 addresses to assign to the network interface controller (NIC). You can specify up to 10 IPv6 addresses. Valid values of N: 1 to 10.
	//
	// Example: Ipv6Address.1=2001:db8:1234:1a00::\\*\\*\\*\\*
	//
	// > If you want to set IPv6 addresses for the network interface controller (NIC), you must set either the parameter `Ipv6Addresses.N` or the parameter `Ipv6AddressCount` but not both.
	//
	// example:
	//
	// 2001:db8:1234:1a00::****
	Ipv6Address []*string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty" type:"Repeated"`
	// The number of IPv6 addresses to randomly generate for the network interface controller (NIC). Valid values: 1 to 10.
	//
	// > If you want to set IPv6 addresses for the network interface controller (NIC), you must set either the parameter `Ipv6Addresses.N` or the parameter `Ipv6AddressCount` but not both.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// One or more IPv6 prefixes to assign to the network interface controller (NIC). Valid values of N: 1 to 10.
	//
	// > If you want to set IPv6 prefixes for the network interface controller (NIC), you must set either the parameter Ipv6Prefix.N or the parameter Ipv6PrefixCount but not both.
	Ipv6Prefix []*string `json:"Ipv6Prefix,omitempty" xml:"Ipv6Prefix,omitempty" type:"Repeated"`
	// The number of IPv6 prefixes to assign to the network interface controller (NIC). Valid values: 1 to 10.
	//
	// > If you want to set IPv6 prefixes for the network interface controller (NIC), you must set either the parameter Ipv6Prefix.N or the parameter Ipv6PrefixCount but not both.
	//
	// example:
	//
	// 1
	Ipv6PrefixCount *int32 `json:"Ipv6PrefixCount,omitempty" xml:"Ipv6PrefixCount,omitempty"`
	// The name of the network interface controller (NIC). The name must be 2 to 128 characters in length and can contain characters from the Unicode letter categorization (including English and Chinese characters) and ASCII digits (0-9). The name can contain colons (:), underscores (_), periods (.), or hyphens (-).
	//
	// Default value: empty.
	//
	// example:
	//
	// testNetworkInterfaceName
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	// The traffic configuration parameter set of the network interface controller (NIC).
	NetworkInterfaceTrafficConfig *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig `json:"NetworkInterfaceTrafficConfig,omitempty" xml:"NetworkInterfaceTrafficConfig,omitempty" type:"Struct"`
	// The communication pattern of the network interface controller (NIC). Valid values:
	//
	// - Standard: uses the TCP communication pattern.
	//
	// - HighPerformance: enables the Elastic RDMA Interface (ERI) and uses the RDMA communication pattern.
	//
	// > A network interface controller (NIC) in RDMA communication pattern can be attached only to an instance whose instance type supports ERI. The number of ENIs in RDMA pattern cannot exceed the limit of the instance family. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html) and [Configure eRDMA on enterprise-level instances](https://help.aliyun.com/document_detail/336853.html)<props="china"> and [Configure eRDMA on GPU-accelerated instances](https://help.aliyun.com/document_detail/2248432.html).
	//
	// Default value: Standard.
	//
	// example:
	//
	// Standard
	NetworkInterfaceTrafficMode *string `json:"NetworkInterfaceTrafficMode,omitempty" xml:"NetworkInterfaceTrafficMode,omitempty"`
	OwnerAccount                *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The primary private IP address of the network interface controller (NIC).
	//
	// The specified IP address must be an idle address within the CIDR block of the vSwitch. If you do not specify this parameter, an idle private IP address in the vSwitch CIDR block is randomly allocated by default.
	//
	// example:
	//
	// ``172.17.**.**``
	PrimaryIpAddress *string `json:"PrimaryIpAddress,omitempty" xml:"PrimaryIpAddress,omitempty"`
	// One or more secondary private IP addresses selected from the idle addresses within the CIDR block of the vSwitch to which the network interface controller (NIC) belongs. Valid values of N: 0 to 10.
	//
	// > When you allocate secondary private IP addresses, you cannot specify both the parameter `PrivateIpAddress.N` and the parameter `SecondaryPrivateIpAddressCount` at the same time.
	//
	// example:
	//
	// ``172.17.**.**``
	PrivateIpAddress []*string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty" type:"Repeated"`
	// The number of queues for the network interface controller (NIC). Valid values: 1 to 2048.
	//
	// When you attach the ENI to an instance, the value must be less than the maximum number of queues per network interface controller (NIC) supported by the instance type. You can call [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) to query the `MaximumQueueNumberPerEni` field.
	//
	// Default value: empty. When the ENI is attached, the default queue number for the instance type is used. You can call [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) to query the `SecondaryEniQueueNumber` field.
	//
	// example:
	//
	// 1
	QueueNumber *int32 `json:"QueueNumber,omitempty" xml:"QueueNumber,omitempty"`
	// The number of queues for the RDMA ENI.
	//
	// If you want to attach multiple RDMA ENIs to an instance, we recommend that you manually specify QueuePairNumber for each ENI based on the upper limit of `QueuePairNumber` supported by the instance type and the number of ENIs you plan to use. Make sure that the total QueuePairNumber of all ENIs does not exceed the maximum value allowed by the instance type. Call [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html) to query the upper limit for the instance type.
	//
	// 	Notice: If QueuePairNumber is not specified for an RDMA ENI, the upper limit of QueuePairNumber for all RDMA ENIs supported by the instance type is used by default. Therefore, after an RDMA ENI without a specified QueuePairNumber is attached, no more RDMA ENIs can be added (regular ENIs are not affected by this limit).</notice>
	//
	// example:
	//
	// 22
	QueuePairNumber *int32 `json:"QueuePairNumber,omitempty" xml:"QueuePairNumber,omitempty"`
	// The region ID of the network interface controller (NIC) to create. You can invoke [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. You can call [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) to query resource group information.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The inbound queue depth of the network interface controller (NIC).
	//
	// Take note of the following items:
	//
	// - The inbound queue depth of the network interface controller (NIC) must be equal to the outbound queue depth. Valid values: 8192 to 16384. The value must be a power of 2.
	//
	// - A larger inbound queue depth increases inbound throughput but consumes more memory.
	//
	// > This parameter is not publicly available.
	//
	// example:
	//
	// 8192
	RxQueueSize *int32 `json:"RxQueueSize,omitempty" xml:"RxQueueSize,omitempty"`
	// The number of private IP addresses for automatic creation by ECS. Valid values: 1 to 49.
	//
	// example:
	//
	// 1
	SecondaryPrivateIpAddressCount *int32 `json:"SecondaryPrivateIpAddressCount,omitempty" xml:"SecondaryPrivateIpAddressCount,omitempty"`
	// The ID of the security group to which the network interface controller (NIC) belongs. The security group and the ENI must be in the same VPC.
	//
	// > When you invoke this operation, you must set either `SecurityGroupId` or `SecurityGroupIds.N` but not both.
	//
	// example:
	//
	// sg-bp1fg655nh68xyz9i****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The IDs of one or more security groups to which the network interface controller (NIC) belongs. The security groups and the ENI must be in the same VPC. The valid values of N depend on the quota for the maximum number of security groups to which an ENI can belong. For more information, see [Limits](https://help.aliyun.com/document_detail/25412.html).
	//
	// > When you invoke this operation, you must set either `SecurityGroupId` or `SecurityGroupIds.N` but not both.
	//
	// example:
	//
	// sg-bp1fg655nh68xyz9i****
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// Specifies whether to enable source/destination checking. We recommend that you enable this feature to improve network security. Valid values:
	//
	// - true: enabled.
	//
	// - false: disabled.
	//
	// Default value: false.
	//
	// > This feature is supported only in specific regions. Before you use this feature, read [Source/destination checking](https://help.aliyun.com/document_detail/2863210.html).
	//
	// example:
	//
	// false
	SourceDestCheck *bool `json:"SourceDestCheck,omitempty" xml:"SourceDestCheck,omitempty"`
	// The tags of the network interface controller (NIC).
	Tag []*CreateNetworkInterfaceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The outbound queue depth of the network interface controller (NIC).
	//
	// Take note of the following items:
	//
	// - The outbound queue depth of the network interface controller (NIC) must be equal to the inbound queue depth. Valid values: 8192 to 16384. The value must be a power of 2.
	//
	// - A larger outbound queue depth increases outbound throughput but consumes more memory.
	//
	// > This parameter is not publicly available.
	//
	// example:
	//
	// 8192
	TxQueueSize *int32 `json:"TxQueueSize,omitempty" xml:"TxQueueSize,omitempty"`
	// The vSwitch ID of the network interface controller (NIC). The private IP address of the ENI is allocated from the idle addresses within the CIDR block of the vSwitch.
	//
	// 	Notice: The network interface controller (NIC) and the instance to which you want to attach the ENI must be in the same zone but can belong to different vSwitches.</notice>
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1s5fnvk4gn2tws03****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// > This parameter is deprecated.
	//
	// example:
	//
	// null
	Visible *bool `json:"Visible,omitempty" xml:"Visible,omitempty"`
}

func (s CreateNetworkInterfaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *CreateNetworkInterfaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateNetworkInterfaceRequest) GetConnectionTrackingConfiguration() *CreateNetworkInterfaceRequestConnectionTrackingConfiguration {
	return s.ConnectionTrackingConfiguration
}

func (s *CreateNetworkInterfaceRequest) GetDeleteOnRelease() *bool {
	return s.DeleteOnRelease
}

func (s *CreateNetworkInterfaceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateNetworkInterfaceRequest) GetEnablePrimaryIPv6() *bool {
	return s.EnablePrimaryIPv6
}

func (s *CreateNetworkInterfaceRequest) GetEnhancedNetwork() *CreateNetworkInterfaceRequestEnhancedNetwork {
	return s.EnhancedNetwork
}

func (s *CreateNetworkInterfaceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateNetworkInterfaceRequest) GetIpv4Prefix() []*string {
	return s.Ipv4Prefix
}

func (s *CreateNetworkInterfaceRequest) GetIpv4PrefixCount() *int32 {
	return s.Ipv4PrefixCount
}

func (s *CreateNetworkInterfaceRequest) GetIpv6Address() []*string {
	return s.Ipv6Address
}

func (s *CreateNetworkInterfaceRequest) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *CreateNetworkInterfaceRequest) GetIpv6Prefix() []*string {
	return s.Ipv6Prefix
}

func (s *CreateNetworkInterfaceRequest) GetIpv6PrefixCount() *int32 {
	return s.Ipv6PrefixCount
}

func (s *CreateNetworkInterfaceRequest) GetNetworkInterfaceName() *string {
	return s.NetworkInterfaceName
}

func (s *CreateNetworkInterfaceRequest) GetNetworkInterfaceTrafficConfig() *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig {
	return s.NetworkInterfaceTrafficConfig
}

func (s *CreateNetworkInterfaceRequest) GetNetworkInterfaceTrafficMode() *string {
	return s.NetworkInterfaceTrafficMode
}

func (s *CreateNetworkInterfaceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateNetworkInterfaceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateNetworkInterfaceRequest) GetPrimaryIpAddress() *string {
	return s.PrimaryIpAddress
}

func (s *CreateNetworkInterfaceRequest) GetPrivateIpAddress() []*string {
	return s.PrivateIpAddress
}

func (s *CreateNetworkInterfaceRequest) GetQueueNumber() *int32 {
	return s.QueueNumber
}

func (s *CreateNetworkInterfaceRequest) GetQueuePairNumber() *int32 {
	return s.QueuePairNumber
}

func (s *CreateNetworkInterfaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNetworkInterfaceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateNetworkInterfaceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateNetworkInterfaceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateNetworkInterfaceRequest) GetRxQueueSize() *int32 {
	return s.RxQueueSize
}

func (s *CreateNetworkInterfaceRequest) GetSecondaryPrivateIpAddressCount() *int32 {
	return s.SecondaryPrivateIpAddressCount
}

func (s *CreateNetworkInterfaceRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateNetworkInterfaceRequest) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *CreateNetworkInterfaceRequest) GetSourceDestCheck() *bool {
	return s.SourceDestCheck
}

func (s *CreateNetworkInterfaceRequest) GetTag() []*CreateNetworkInterfaceRequestTag {
	return s.Tag
}

func (s *CreateNetworkInterfaceRequest) GetTxQueueSize() *int32 {
	return s.TxQueueSize
}

func (s *CreateNetworkInterfaceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateNetworkInterfaceRequest) GetVisible() *bool {
	return s.Visible
}

func (s *CreateNetworkInterfaceRequest) SetBusinessType(v string) *CreateNetworkInterfaceRequest {
	s.BusinessType = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetClientToken(v string) *CreateNetworkInterfaceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetConnectionTrackingConfiguration(v *CreateNetworkInterfaceRequestConnectionTrackingConfiguration) *CreateNetworkInterfaceRequest {
	s.ConnectionTrackingConfiguration = v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetDeleteOnRelease(v bool) *CreateNetworkInterfaceRequest {
	s.DeleteOnRelease = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetDescription(v string) *CreateNetworkInterfaceRequest {
	s.Description = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetEnablePrimaryIPv6(v bool) *CreateNetworkInterfaceRequest {
	s.EnablePrimaryIPv6 = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetEnhancedNetwork(v *CreateNetworkInterfaceRequestEnhancedNetwork) *CreateNetworkInterfaceRequest {
	s.EnhancedNetwork = v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetInstanceType(v string) *CreateNetworkInterfaceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetIpv4Prefix(v []*string) *CreateNetworkInterfaceRequest {
	s.Ipv4Prefix = v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetIpv4PrefixCount(v int32) *CreateNetworkInterfaceRequest {
	s.Ipv4PrefixCount = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetIpv6Address(v []*string) *CreateNetworkInterfaceRequest {
	s.Ipv6Address = v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetIpv6AddressCount(v int32) *CreateNetworkInterfaceRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetIpv6Prefix(v []*string) *CreateNetworkInterfaceRequest {
	s.Ipv6Prefix = v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetIpv6PrefixCount(v int32) *CreateNetworkInterfaceRequest {
	s.Ipv6PrefixCount = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetNetworkInterfaceName(v string) *CreateNetworkInterfaceRequest {
	s.NetworkInterfaceName = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetNetworkInterfaceTrafficConfig(v *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) *CreateNetworkInterfaceRequest {
	s.NetworkInterfaceTrafficConfig = v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetNetworkInterfaceTrafficMode(v string) *CreateNetworkInterfaceRequest {
	s.NetworkInterfaceTrafficMode = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetOwnerAccount(v string) *CreateNetworkInterfaceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetOwnerId(v int64) *CreateNetworkInterfaceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetPrimaryIpAddress(v string) *CreateNetworkInterfaceRequest {
	s.PrimaryIpAddress = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetPrivateIpAddress(v []*string) *CreateNetworkInterfaceRequest {
	s.PrivateIpAddress = v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetQueueNumber(v int32) *CreateNetworkInterfaceRequest {
	s.QueueNumber = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetQueuePairNumber(v int32) *CreateNetworkInterfaceRequest {
	s.QueuePairNumber = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetRegionId(v string) *CreateNetworkInterfaceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetResourceGroupId(v string) *CreateNetworkInterfaceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetResourceOwnerAccount(v string) *CreateNetworkInterfaceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetResourceOwnerId(v int64) *CreateNetworkInterfaceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetRxQueueSize(v int32) *CreateNetworkInterfaceRequest {
	s.RxQueueSize = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetSecondaryPrivateIpAddressCount(v int32) *CreateNetworkInterfaceRequest {
	s.SecondaryPrivateIpAddressCount = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetSecurityGroupId(v string) *CreateNetworkInterfaceRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetSecurityGroupIds(v []*string) *CreateNetworkInterfaceRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetSourceDestCheck(v bool) *CreateNetworkInterfaceRequest {
	s.SourceDestCheck = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetTag(v []*CreateNetworkInterfaceRequestTag) *CreateNetworkInterfaceRequest {
	s.Tag = v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetTxQueueSize(v int32) *CreateNetworkInterfaceRequest {
	s.TxQueueSize = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetVSwitchId(v string) *CreateNetworkInterfaceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetVisible(v bool) *CreateNetworkInterfaceRequest {
	s.Visible = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) Validate() error {
	if s.ConnectionTrackingConfiguration != nil {
		if err := s.ConnectionTrackingConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.EnhancedNetwork != nil {
		if err := s.EnhancedNetwork.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkInterfaceTrafficConfig != nil {
		if err := s.NetworkInterfaceTrafficConfig.Validate(); err != nil {
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

type CreateNetworkInterfaceRequestConnectionTrackingConfiguration struct {
	// The timeout period for TCP connections in the TIME_WAIT and CLOSED states. Unit: seconds. Valid values: integers from 3 to 15.
	//
	// Default value: 3.
	//
	// > If your ECS instance is used with NLB/CLB, the default timeout period for connections in the `TIME_WAIT` state is 15 seconds.
	//
	// example:
	//
	// 3
	TcpClosedAndTimeWaitTimeout *int32 `json:"TcpClosedAndTimeWaitTimeout,omitempty" xml:"TcpClosedAndTimeWaitTimeout,omitempty"`
	// The timeout period for established TCP connections. Unit: seconds. Valid values: [30, 60, 80, 100, 200, 300, 500, 700, 910].
	//
	// Default value: 910.
	//
	// example:
	//
	// 910
	TcpEstablishedTimeout *int32 `json:"TcpEstablishedTimeout,omitempty" xml:"TcpEstablishedTimeout,omitempty"`
	// The timeout period for UDP flows. Unit: seconds. Valid values: [10, 20, 30, 60, 80, 100].
	//
	// Default value: 30.
	//
	// > If your ECS instance is used with NLB/CLB, the default value is 100 seconds.
	//
	// example:
	//
	// 30
	UdpTimeout *int32 `json:"UdpTimeout,omitempty" xml:"UdpTimeout,omitempty"`
}

func (s CreateNetworkInterfaceRequestConnectionTrackingConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceRequestConnectionTrackingConfiguration) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceRequestConnectionTrackingConfiguration) GetTcpClosedAndTimeWaitTimeout() *int32 {
	return s.TcpClosedAndTimeWaitTimeout
}

func (s *CreateNetworkInterfaceRequestConnectionTrackingConfiguration) GetTcpEstablishedTimeout() *int32 {
	return s.TcpEstablishedTimeout
}

func (s *CreateNetworkInterfaceRequestConnectionTrackingConfiguration) GetUdpTimeout() *int32 {
	return s.UdpTimeout
}

func (s *CreateNetworkInterfaceRequestConnectionTrackingConfiguration) SetTcpClosedAndTimeWaitTimeout(v int32) *CreateNetworkInterfaceRequestConnectionTrackingConfiguration {
	s.TcpClosedAndTimeWaitTimeout = &v
	return s
}

func (s *CreateNetworkInterfaceRequestConnectionTrackingConfiguration) SetTcpEstablishedTimeout(v int32) *CreateNetworkInterfaceRequestConnectionTrackingConfiguration {
	s.TcpEstablishedTimeout = &v
	return s
}

func (s *CreateNetworkInterfaceRequestConnectionTrackingConfiguration) SetUdpTimeout(v int32) *CreateNetworkInterfaceRequestConnectionTrackingConfiguration {
	s.UdpTimeout = &v
	return s
}

func (s *CreateNetworkInterfaceRequestConnectionTrackingConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateNetworkInterfaceRequestEnhancedNetwork struct {
	EnableExpress *bool `json:"EnableExpress,omitempty" xml:"EnableExpress,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// true
	EnableRss *bool `json:"EnableRss,omitempty" xml:"EnableRss,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// true
	EnableSriov                     *bool  `json:"EnableSriov,omitempty" xml:"EnableSriov,omitempty"`
	VirtualFunctionQuantity         *int32 `json:"VirtualFunctionQuantity,omitempty" xml:"VirtualFunctionQuantity,omitempty"`
	VirtualFunctionTotalQueueNumber *int32 `json:"VirtualFunctionTotalQueueNumber,omitempty" xml:"VirtualFunctionTotalQueueNumber,omitempty"`
}

func (s CreateNetworkInterfaceRequestEnhancedNetwork) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceRequestEnhancedNetwork) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceRequestEnhancedNetwork) GetEnableExpress() *bool {
	return s.EnableExpress
}

func (s *CreateNetworkInterfaceRequestEnhancedNetwork) GetEnableRss() *bool {
	return s.EnableRss
}

func (s *CreateNetworkInterfaceRequestEnhancedNetwork) GetEnableSriov() *bool {
	return s.EnableSriov
}

func (s *CreateNetworkInterfaceRequestEnhancedNetwork) GetVirtualFunctionQuantity() *int32 {
	return s.VirtualFunctionQuantity
}

func (s *CreateNetworkInterfaceRequestEnhancedNetwork) GetVirtualFunctionTotalQueueNumber() *int32 {
	return s.VirtualFunctionTotalQueueNumber
}

func (s *CreateNetworkInterfaceRequestEnhancedNetwork) SetEnableExpress(v bool) *CreateNetworkInterfaceRequestEnhancedNetwork {
	s.EnableExpress = &v
	return s
}

func (s *CreateNetworkInterfaceRequestEnhancedNetwork) SetEnableRss(v bool) *CreateNetworkInterfaceRequestEnhancedNetwork {
	s.EnableRss = &v
	return s
}

func (s *CreateNetworkInterfaceRequestEnhancedNetwork) SetEnableSriov(v bool) *CreateNetworkInterfaceRequestEnhancedNetwork {
	s.EnableSriov = &v
	return s
}

func (s *CreateNetworkInterfaceRequestEnhancedNetwork) SetVirtualFunctionQuantity(v int32) *CreateNetworkInterfaceRequestEnhancedNetwork {
	s.VirtualFunctionQuantity = &v
	return s
}

func (s *CreateNetworkInterfaceRequestEnhancedNetwork) SetVirtualFunctionTotalQueueNumber(v int32) *CreateNetworkInterfaceRequestEnhancedNetwork {
	s.VirtualFunctionTotalQueueNumber = &v
	return s
}

func (s *CreateNetworkInterfaceRequestEnhancedNetwork) Validate() error {
	return dara.Validate(s)
}

type CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig struct {
	// The communication pattern of the network interface controller (NIC).
	//
	// example:
	//
	// HighPerformance
	NetworkInterfaceTrafficMode *string `json:"NetworkInterfaceTrafficMode,omitempty" xml:"NetworkInterfaceTrafficMode,omitempty"`
	// The number of queues for the network interface controller (NIC).
	//
	// example:
	//
	// 8
	QueueNumber *int32 `json:"QueueNumber,omitempty" xml:"QueueNumber,omitempty"`
	// The number of queues for the RDMA ENI.
	//
	// example:
	//
	// 8
	QueuePairNumber *int32 `json:"QueuePairNumber,omitempty" xml:"QueuePairNumber,omitempty"`
	// The inbound queue depth of the network interface controller (NIC).
	//
	//
	// <props="china">
	//
	// >This parameter is in invitational preview and is not publicly available. If you want to use this parameter, [submit a ticket](https://selfservice.console.aliyun.com/ticket/createIndex) to request access.
	//
	//
	//
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is in invitational preview and is not publicly available. If you want to use this parameter, [submit a ticket](https://smartservice.console.aliyun.com/service/create-ticket-intl) to request access.
	//
	//
	//
	// Take note of the following items:
	//
	// - This parameter applies only to seventh-generation and later ECS instance types.
	//
	// - This parameter currently applies only to Linux images.
	//
	// - A larger inbound queue depth of the network interface controller (NIC) increases inbound throughput and reduces packet loss probability but consumes more memory.
	//
	// example:
	//
	// 8192
	RxQueueSize *int32 `json:"RxQueueSize,omitempty" xml:"RxQueueSize,omitempty"`
	// The outbound queue depth of the network interface controller (NIC).
	//
	//
	// <props="china">
	//
	// >This parameter is in invitational preview and is not publicly available. If you want to use this parameter, [submit a ticket](https://selfservice.console.aliyun.com/ticket/createIndex) to request access.
	//
	//
	//
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is in invitational preview and is not publicly available. If you want to use this parameter, [submit a ticket](https://smartservice.console.aliyun.com/service/create-ticket-intl) to request access.
	//
	//
	//
	// Take note of the following items:
	//
	// - This parameter applies only to seventh-generation and later ECS instance types.
	//
	// - This parameter currently applies only to Linux images.
	//
	// - A larger outbound queue depth of the network interface controller (NIC) increases outbound throughput and reduces packet loss probability but consumes more memory.
	//
	// example:
	//
	// 8192
	TxQueueSize *int32 `json:"TxQueueSize,omitempty" xml:"TxQueueSize,omitempty"`
}

func (s CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) GetNetworkInterfaceTrafficMode() *string {
	return s.NetworkInterfaceTrafficMode
}

func (s *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) GetQueueNumber() *int32 {
	return s.QueueNumber
}

func (s *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) GetQueuePairNumber() *int32 {
	return s.QueuePairNumber
}

func (s *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) GetRxQueueSize() *int32 {
	return s.RxQueueSize
}

func (s *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) GetTxQueueSize() *int32 {
	return s.TxQueueSize
}

func (s *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) SetNetworkInterfaceTrafficMode(v string) *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig {
	s.NetworkInterfaceTrafficMode = &v
	return s
}

func (s *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) SetQueueNumber(v int32) *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig {
	s.QueueNumber = &v
	return s
}

func (s *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) SetQueuePairNumber(v int32) *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig {
	s.QueuePairNumber = &v
	return s
}

func (s *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) SetRxQueueSize(v int32) *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig {
	s.RxQueueSize = &v
	return s
}

func (s *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) SetTxQueueSize(v int32) *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig {
	s.TxQueueSize = &v
	return s
}

func (s *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) Validate() error {
	return dara.Validate(s)
}

type CreateNetworkInterfaceRequestTag struct {
	// The tag key of the network interface controller (NIC). Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with aliyun or acs:. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the network interface controller (NIC). Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateNetworkInterfaceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateNetworkInterfaceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateNetworkInterfaceRequestTag) SetKey(v string) *CreateNetworkInterfaceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateNetworkInterfaceRequestTag) SetValue(v string) *CreateNetworkInterfaceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateNetworkInterfaceRequestTag) Validate() error {
	return dara.Validate(s)
}
