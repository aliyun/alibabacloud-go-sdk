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
	// A client token to ensure request idempotence. Your client generates this token, which must be unique across requests. The token can contain only ASCII characters and must not exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The connection tracking settings.
	//
	// Before using this parameter, read [Connection timeout management](https://help.aliyun.com/document_detail/2865958.html).
	ConnectionTrackingConfiguration *CreateNetworkInterfaceRequestConnectionTrackingConfiguration `json:"ConnectionTrackingConfiguration,omitempty" xml:"ConnectionTrackingConfiguration,omitempty" type:"Struct"`
	// Specifies whether to release the elastic network interface when its attached instance is released. Valid values:
	//
	// - `true`: The elastic network interface is released.
	//
	// - `false`: The elastic network interface is retained.
	//
	// example:
	//
	// true
	DeleteOnRelease *bool `json:"DeleteOnRelease,omitempty" xml:"DeleteOnRelease,omitempty"`
	// The description of the elastic network interface. The description must be 2 to 256 characters long and cannot start with `http://` or `https://`.
	//
	// Default value: empty.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// > This parameter is not publicly available.
	EnhancedNetwork *CreateNetworkInterfaceRequestEnhancedNetwork `json:"EnhancedNetwork,omitempty" xml:"EnhancedNetwork,omitempty" type:"Struct"`
	// The type of the elastic network interface. Valid values:
	//
	// - `Secondary`: a secondary elastic network interface.
	//
	// - `Trunk`: a trunk network interface. (This feature is available by invitation only.)
	//
	// Default value: `Secondary`.
	//
	// example:
	//
	// Secondary
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// One or more IPv4 prefixes to assign to the elastic network interface. Valid values of N: 1 to 10.
	//
	// > You must specify either `Ipv4Prefix.N` or `Ipv4PrefixCount`, but not both, to assign IPv4 prefixes.
	Ipv4Prefix []*string `json:"Ipv4Prefix,omitempty" xml:"Ipv4Prefix,omitempty" type:"Repeated"`
	// The number of IPv4 prefixes to assign to the elastic network interface. Valid values: 1 to 10.
	//
	// > You must specify either `Ipv4Prefix.N` or `Ipv4PrefixCount`, but not both, to assign IPv4 prefixes.
	//
	// example:
	//
	// 1
	Ipv4PrefixCount *int32 `json:"Ipv4PrefixCount,omitempty" xml:"Ipv4PrefixCount,omitempty"`
	// One or more IPv6 addresses to assign to the elastic network interface. You can specify up to 10 IPv6 addresses. Valid values of N: 1 to 10.
	//
	// Example: `Ipv6Address.1=2001:db8:1234:1a00::****`
	//
	// > You must specify either `Ipv6Address.N` or `Ipv6AddressCount`, but not both, to assign IPv6 addresses.
	//
	// example:
	//
	// 2001:db8:1234:1a00::****
	Ipv6Address []*string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty" type:"Repeated"`
	// The number of random IPv6 addresses to assign to the elastic network interface. Valid values: 1 to 10.
	//
	// > You must specify either `Ipv6Address.N` or `Ipv6AddressCount`, but not both, to assign IPv6 addresses.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// One or more IPv6 prefixes to assign to the elastic network interface. Valid values of N: 1 to 10.
	//
	// > You must specify either `Ipv6Prefix.N` or `Ipv6PrefixCount`, but not both, to assign IPv6 prefixes.
	Ipv6Prefix []*string `json:"Ipv6Prefix,omitempty" xml:"Ipv6Prefix,omitempty" type:"Repeated"`
	// The number of IPv6 prefixes to assign to the elastic network interface. Valid values: 1 to 10.
	//
	// > You must specify either `Ipv6Prefix.N` or `Ipv6PrefixCount`, but not both, to assign IPv6 prefixes.
	//
	// example:
	//
	// 1
	Ipv6PrefixCount *int32 `json:"Ipv6PrefixCount,omitempty" xml:"Ipv6PrefixCount,omitempty"`
	// The name of the elastic network interface. The name must be 2 to 128 characters long and can contain Unicode letters (such as English and Chinese characters), digits (0-9), colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// Default value: empty.
	//
	// example:
	//
	// testNetworkInterfaceName
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	// The communication parameters of the elastic network interface.
	NetworkInterfaceTrafficConfig *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig `json:"NetworkInterfaceTrafficConfig,omitempty" xml:"NetworkInterfaceTrafficConfig,omitempty" type:"Struct"`
	// The traffic mode of the elastic network interface. Valid values:
	//
	// - `Standard`: uses the TCP traffic mode.
	//
	// - `HighPerformance`: enables the Elastic RDMA Interface (ERI) and uses the RDMA traffic mode.
	//
	// > An elastic network interface in RDMA traffic mode can be attached only to an ERI-supported instance type. The number of these elastic network interfaces that can be attached is limited by the instance family. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html), [Configure eRDMA on an enterprise-level instance](https://help.aliyun.com/document_detail/336853.html)<props="china">, and [Configure eRDMA on a GPU instance](https://help.aliyun.com/document_detail/2248432.html).
	//
	// Default value: `Standard`.
	//
	// example:
	//
	// Standard
	NetworkInterfaceTrafficMode *string `json:"NetworkInterfaceTrafficMode,omitempty" xml:"NetworkInterfaceTrafficMode,omitempty"`
	OwnerAccount                *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The primary private IP address of the elastic network interface.
	//
	// The IP address must be an available IP address within the CIDR block of the VSwitch. If this parameter is not specified, the system randomly assigns an available IP address from the VSwitch\\"s CIDR block.
	//
	// example:
	//
	// ``172.17.**.**``
	PrimaryIpAddress *string `json:"PrimaryIpAddress,omitempty" xml:"PrimaryIpAddress,omitempty"`
	// One or more secondary private IP addresses to assign to the elastic network interface. The IP addresses must be available addresses from the CIDR block of the VSwitch to which the elastic network interface belongs. Valid values of N: 0 to 10.
	//
	// > You cannot specify both `PrivateIpAddress.N` and `SecondaryPrivateIpAddressCount` to assign secondary private IP addresses.
	//
	// example:
	//
	// ``172.17.**.**``
	PrivateIpAddress []*string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty" type:"Repeated"`
	// The number of queues for the elastic network interface. Valid values: 1 to 2048.
	//
	// When attached to an instance, this value must be less than the maximum number of queues per elastic network interface that the instance type supports. You can call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation and check the `MaximumQueueNumberPerEni` value in the response to query this limit.
	//
	// If you do not specify this parameter, the default queue number for the instance type is used upon attachment.
	//
	// example:
	//
	// 1
	QueueNumber *int32 `json:"QueueNumber,omitempty" xml:"QueueNumber,omitempty"`
	// The number of queue pairs for the RDMA-enabled elastic network interface.
	//
	// If you want to attach multiple RDMA-enabled elastic network interfaces to an instance, we recommend that you specify a `QueuePairNumber` value for each elastic network interface. The value should be based on the maximum `QueuePairNumber` value supported by the instance type and the number of elastic network interfaces that you plan to use. The total number of queue pairs for all elastic network interfaces cannot exceed the maximum value for the instance type. You can call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html) operation to query the maximum value.
	//
	// 	Notice:
	//
	// If you do not specify `QueuePairNumber` for an RDMA-enabled elastic network interface, the system defaults to the maximum value that the instance type supports. Consequently, you cannot attach any more RDMA-enabled elastic network interfaces to that instance. This does not affect standard elastic network interfaces.
	//
	// example:
	//
	// 22
	QueuePairNumber *int32 `json:"QueuePairNumber,omitempty" xml:"QueuePairNumber,omitempty"`
	// The ID of the region in which to create the elastic network interface. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to view the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query resource groups.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The receive (Rx) queue depth of the elastic network interface.
	//
	// - The receive (Rx) and transmit (Tx) queue depths must be equal. The value must be a power of 2 between 8,192 and 16,384.
	//
	// - A larger Rx queue depth can improve receive throughput but consumes more memory.
	//
	// > This parameter is not publicly available.
	//
	// example:
	//
	// 8192
	RxQueueSize *int32 `json:"RxQueueSize,omitempty" xml:"RxQueueSize,omitempty"`
	// The number of secondary private IP addresses to automatically assign to the elastic network interface. Valid values: 1 to 49.
	//
	// example:
	//
	// 1
	SecondaryPrivateIpAddressCount *int32 `json:"SecondaryPrivateIpAddressCount,omitempty" xml:"SecondaryPrivateIpAddressCount,omitempty"`
	// The ID of the security group for the elastic network interface. The security group and the elastic network interface must be in the same VPC.
	//
	// > You must specify either `SecurityGroupId` or `SecurityGroupIds.N`, but not both.
	//
	// example:
	//
	// sg-bp1fg655nh68xyz9i****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The IDs of one or more security groups to which to add the elastic network interface. The security groups and the elastic network interface must be in the same VPC. The valid values of N depend on the maximum number of security groups to which an elastic network interface can be added. For more information, see [Limits](https://help.aliyun.com/document_detail/25412.html).
	//
	// > You must specify either `SecurityGroupId` or `SecurityGroupIds.N`, but not both.
	//
	// example:
	//
	// sg-bp1fg655nh68xyz9i****
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// Specifies whether to enable source/destination check. Enabling this feature enhances network security. Valid values:
	//
	// - `true`: enabled.
	//
	// - `false`: disabled.
	//
	// Default value: false.
	//
	// > This feature is available only in some regions. Before you use this feature, read [Source/destination check](https://help.aliyun.com/document_detail/2863210.html).
	//
	// example:
	//
	// false
	SourceDestCheck *bool `json:"SourceDestCheck,omitempty" xml:"SourceDestCheck,omitempty"`
	// The tags to add to the elastic network interface.
	Tag []*CreateNetworkInterfaceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The transmit (Tx) queue depth of the elastic network interface.
	//
	// - The transmit (Tx) and receive (Rx) queue depths must be equal. The value must be a power of 2 between 8,192 and 16,384.
	//
	// - A larger Tx queue depth can improve transmit throughput but consumes more memory.
	//
	// > This parameter is not publicly available.
	//
	// example:
	//
	// 8192
	TxQueueSize *int32 `json:"TxQueueSize,omitempty" xml:"TxQueueSize,omitempty"`
	// The ID of the VSwitch for the elastic network interface. The private IP addresses for the elastic network interface are assigned from the available CIDR block of the VSwitch.
	//
	// 	Notice:
	//
	// The elastic network interface and the instance to be attached must be in the same availability zone but can belong to different VSwitches.
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
	// The timeout for a TCP connection in the TIME_WAIT or closing state, in seconds. Valid values: integers from 3 to 15.
	//
	// Default value: 3.
	//
	// > If your ECS instance works with NLB or CLB, the default timeout period for connections in the `TIME_WAIT` state is 15 seconds.
	//
	// example:
	//
	// 3
	TcpClosedAndTimeWaitTimeout *int32 `json:"TcpClosedAndTimeWaitTimeout,omitempty" xml:"TcpClosedAndTimeWaitTimeout,omitempty"`
	// The timeout for an established TCP connection, in seconds. Valid values: 30, 60, 80, 100, 200, 300, 500, 700, and 910.
	//
	// Default value: 910.
	//
	// example:
	//
	// 910
	TcpEstablishedTimeout *int32 `json:"TcpEstablishedTimeout,omitempty" xml:"TcpEstablishedTimeout,omitempty"`
	// The timeout for a UDP stream, in seconds. Valid values: 10, 20, 30, 60, 80, and 100.
	//
	// Default value: 30.
	//
	// > If your ECS instance works with NLB or CLB, the default value is 100 seconds.
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
	// The traffic mode of the elastic network interface.
	//
	// example:
	//
	// HighPerformance
	NetworkInterfaceTrafficMode *string `json:"NetworkInterfaceTrafficMode,omitempty" xml:"NetworkInterfaceTrafficMode,omitempty"`
	// The number of queues for the elastic network interface.
	//
	// example:
	//
	// 8
	QueueNumber *int32 `json:"QueueNumber,omitempty" xml:"QueueNumber,omitempty"`
	// The number of queue pairs for the RDMA-enabled elastic network interface.
	//
	// example:
	//
	// 8
	QueuePairNumber *int32 `json:"QueuePairNumber,omitempty" xml:"QueuePairNumber,omitempty"`
	// The receive (Rx) queue depth of the elastic network interface.
	//
	// <props="china">
	//
	// > This parameter is available by invitation only. To request access, submit a ticket.
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is available by invitation only. To request access, submit a ticket.
	//
	//
	//
	// - This parameter is applicable only to seventh-generation or later ECS instance types.
	//
	// - This parameter is applicable only to Linux images.
	//
	// - A larger Rx queue depth can improve receive throughput and reduce the packet loss rate, but consumes more memory.
	//
	// example:
	//
	// 8192
	RxQueueSize *int32 `json:"RxQueueSize,omitempty" xml:"RxQueueSize,omitempty"`
	// The transmit (Tx) queue depth of the elastic network interface.
	//
	// <props="china">
	//
	// > This parameter is available by invitation only. To request access, submit a ticket.
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is available by invitation only. To request access, submit a ticket.
	//
	//
	//
	// - This parameter is applicable only to seventh-generation or later ECS instance types.
	//
	// - This parameter is applicable only to Linux images.
	//
	// - A larger Tx queue depth can improve transmit throughput and reduce the packet loss rate, but consumes more memory.
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
	// The key of the tag. Valid values for N: 1 to 20. The tag key cannot be an empty string. It can be up to 128 characters long and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. Valid values for N: 1 to 20. The tag value can be an empty string. It can be up to 128 characters long and cannot contain `http://` or `https://`.
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
