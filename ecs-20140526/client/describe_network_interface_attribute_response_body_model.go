// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkInterfaceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssociatedPublicIp(v *DescribeNetworkInterfaceAttributeResponseBodyAssociatedPublicIp) *DescribeNetworkInterfaceAttributeResponseBody
	GetAssociatedPublicIp() *DescribeNetworkInterfaceAttributeResponseBodyAssociatedPublicIp
	SetAttachment(v *DescribeNetworkInterfaceAttributeResponseBodyAttachment) *DescribeNetworkInterfaceAttributeResponseBody
	GetAttachment() *DescribeNetworkInterfaceAttributeResponseBodyAttachment
	SetBondInterfaceSpecification(v *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecification) *DescribeNetworkInterfaceAttributeResponseBody
	GetBondInterfaceSpecification() *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecification
	SetConnectionTrackingConfiguration(v *DescribeNetworkInterfaceAttributeResponseBodyConnectionTrackingConfiguration) *DescribeNetworkInterfaceAttributeResponseBody
	GetConnectionTrackingConfiguration() *DescribeNetworkInterfaceAttributeResponseBodyConnectionTrackingConfiguration
	SetCreationTime(v string) *DescribeNetworkInterfaceAttributeResponseBody
	GetCreationTime() *string
	SetDeleteOnRelease(v bool) *DescribeNetworkInterfaceAttributeResponseBody
	GetDeleteOnRelease() *bool
	SetDescription(v string) *DescribeNetworkInterfaceAttributeResponseBody
	GetDescription() *string
	SetEnhancedNetwork(v *DescribeNetworkInterfaceAttributeResponseBodyEnhancedNetwork) *DescribeNetworkInterfaceAttributeResponseBody
	GetEnhancedNetwork() *DescribeNetworkInterfaceAttributeResponseBodyEnhancedNetwork
	SetInstanceId(v string) *DescribeNetworkInterfaceAttributeResponseBody
	GetInstanceId() *string
	SetIpv4PrefixSets(v *DescribeNetworkInterfaceAttributeResponseBodyIpv4PrefixSets) *DescribeNetworkInterfaceAttributeResponseBody
	GetIpv4PrefixSets() *DescribeNetworkInterfaceAttributeResponseBodyIpv4PrefixSets
	SetIpv6PrefixSets(v *DescribeNetworkInterfaceAttributeResponseBodyIpv6PrefixSets) *DescribeNetworkInterfaceAttributeResponseBody
	GetIpv6PrefixSets() *DescribeNetworkInterfaceAttributeResponseBodyIpv6PrefixSets
	SetIpv6Sets(v *DescribeNetworkInterfaceAttributeResponseBodyIpv6Sets) *DescribeNetworkInterfaceAttributeResponseBody
	GetIpv6Sets() *DescribeNetworkInterfaceAttributeResponseBodyIpv6Sets
	SetMacAddress(v string) *DescribeNetworkInterfaceAttributeResponseBody
	GetMacAddress() *string
	SetNetworkInterfaceId(v string) *DescribeNetworkInterfaceAttributeResponseBody
	GetNetworkInterfaceId() *string
	SetNetworkInterfaceName(v string) *DescribeNetworkInterfaceAttributeResponseBody
	GetNetworkInterfaceName() *string
	SetNetworkInterfaceTrafficConfig(v *DescribeNetworkInterfaceAttributeResponseBodyNetworkInterfaceTrafficConfig) *DescribeNetworkInterfaceAttributeResponseBody
	GetNetworkInterfaceTrafficConfig() *DescribeNetworkInterfaceAttributeResponseBodyNetworkInterfaceTrafficConfig
	SetNetworkInterfaceTrafficMode(v string) *DescribeNetworkInterfaceAttributeResponseBody
	GetNetworkInterfaceTrafficMode() *string
	SetOwnerId(v string) *DescribeNetworkInterfaceAttributeResponseBody
	GetOwnerId() *string
	SetPrivateIpAddress(v string) *DescribeNetworkInterfaceAttributeResponseBody
	GetPrivateIpAddress() *string
	SetPrivateIpSets(v *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSets) *DescribeNetworkInterfaceAttributeResponseBody
	GetPrivateIpSets() *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSets
	SetQoSConfig(v *DescribeNetworkInterfaceAttributeResponseBodyQoSConfig) *DescribeNetworkInterfaceAttributeResponseBody
	GetQoSConfig() *DescribeNetworkInterfaceAttributeResponseBodyQoSConfig
	SetQueueNumber(v int32) *DescribeNetworkInterfaceAttributeResponseBody
	GetQueueNumber() *int32
	SetQueuePairNumber(v int32) *DescribeNetworkInterfaceAttributeResponseBody
	GetQueuePairNumber() *int32
	SetRequestId(v string) *DescribeNetworkInterfaceAttributeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeNetworkInterfaceAttributeResponseBody
	GetResourceGroupId() *string
	SetSecurityGroupIds(v *DescribeNetworkInterfaceAttributeResponseBodySecurityGroupIds) *DescribeNetworkInterfaceAttributeResponseBody
	GetSecurityGroupIds() *DescribeNetworkInterfaceAttributeResponseBodySecurityGroupIds
	SetServiceID(v int64) *DescribeNetworkInterfaceAttributeResponseBody
	GetServiceID() *int64
	SetServiceManaged(v bool) *DescribeNetworkInterfaceAttributeResponseBody
	GetServiceManaged() *bool
	SetSlaveInterfaceSpecification(v *DescribeNetworkInterfaceAttributeResponseBodySlaveInterfaceSpecification) *DescribeNetworkInterfaceAttributeResponseBody
	GetSlaveInterfaceSpecification() *DescribeNetworkInterfaceAttributeResponseBodySlaveInterfaceSpecification
	SetSourceDestCheck(v bool) *DescribeNetworkInterfaceAttributeResponseBody
	GetSourceDestCheck() *bool
	SetStatus(v string) *DescribeNetworkInterfaceAttributeResponseBody
	GetStatus() *string
	SetTags(v *DescribeNetworkInterfaceAttributeResponseBodyTags) *DescribeNetworkInterfaceAttributeResponseBody
	GetTags() *DescribeNetworkInterfaceAttributeResponseBodyTags
	SetTcpOptionAddressEnabled(v string) *DescribeNetworkInterfaceAttributeResponseBody
	GetTcpOptionAddressEnabled() *string
	SetType(v string) *DescribeNetworkInterfaceAttributeResponseBody
	GetType() *string
	SetVSwitchId(v string) *DescribeNetworkInterfaceAttributeResponseBody
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeNetworkInterfaceAttributeResponseBody
	GetVpcId() *string
	SetZoneId(v string) *DescribeNetworkInterfaceAttributeResponseBody
	GetZoneId() *string
}

type DescribeNetworkInterfaceAttributeResponseBody struct {
	// The elastic IP address that is associated with the primary private IP address of the elastic network interface.
	AssociatedPublicIp *DescribeNetworkInterfaceAttributeResponseBodyAssociatedPublicIp `json:"AssociatedPublicIp,omitempty" xml:"AssociatedPublicIp,omitempty" type:"Struct"`
	// > This parameter is in invitational preview and is not publicly available.
	Attachment *DescribeNetworkInterfaceAttributeResponseBodyAttachment `json:"Attachment,omitempty" xml:"Attachment,omitempty" type:"Struct"`
	// > This parameter is in invitational preview and is not publicly available.
	BondInterfaceSpecification *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecification `json:"BondInterfaceSpecification,omitempty" xml:"BondInterfaceSpecification,omitempty" type:"Struct"`
	// The connection tracking configuration.
	//
	// For more information, see [Connection timeout management](https://help.aliyun.com/document_detail/2865958.html).
	//
	// > This parameter is returned only if the `Attribute` parameter is set to `connectionTrackingConfiguration` in the request.
	ConnectionTrackingConfiguration *DescribeNetworkInterfaceAttributeResponseBodyConnectionTrackingConfiguration `json:"ConnectionTrackingConfiguration,omitempty" xml:"ConnectionTrackingConfiguration,omitempty" type:"Struct"`
	// The time when the elastic network interface was created.
	//
	// example:
	//
	// 2019-12-25T12:31:31Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Indicates whether to release the elastic network interface when the associated instance is released.
	//
	// - `true`: The interface is released.
	//
	// - `false`: The interface is retained.
	//
	// example:
	//
	// true
	DeleteOnRelease *bool `json:"DeleteOnRelease,omitempty" xml:"DeleteOnRelease,omitempty"`
	// The description of the elastic network interface.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is not publicly available.
	EnhancedNetwork *DescribeNetworkInterfaceAttributeResponseBodyEnhancedNetwork `json:"EnhancedNetwork,omitempty" xml:"EnhancedNetwork,omitempty" type:"Struct"`
	// The ID of the instance to which the elastic network interface is attached.
	//
	// > This parameter is not returned if the elastic network interface is managed by another Alibaba Cloud service.
	//
	// example:
	//
	// i-bp1e2l6djkndyuli****
	InstanceId     *string                                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ipv4PrefixSets *DescribeNetworkInterfaceAttributeResponseBodyIpv4PrefixSets `json:"Ipv4PrefixSets,omitempty" xml:"Ipv4PrefixSets,omitempty" type:"Struct"`
	Ipv6PrefixSets *DescribeNetworkInterfaceAttributeResponseBodyIpv6PrefixSets `json:"Ipv6PrefixSets,omitempty" xml:"Ipv6PrefixSets,omitempty" type:"Struct"`
	Ipv6Sets       *DescribeNetworkInterfaceAttributeResponseBodyIpv6Sets       `json:"Ipv6Sets,omitempty" xml:"Ipv6Sets,omitempty" type:"Struct"`
	// The MAC address of the elastic network interface.
	//
	// example:
	//
	// 00:16:3e:12:**:**
	MacAddress *string `json:"MacAddress,omitempty" xml:"MacAddress,omitempty"`
	// The ID of the elastic network interface.
	//
	// example:
	//
	// eni-bp125p95hhdhn3ot****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The name of the elastic network interface.
	//
	// example:
	//
	// my-eni-name
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	// The communication parameters of the elastic network interface.
	NetworkInterfaceTrafficConfig *DescribeNetworkInterfaceAttributeResponseBodyNetworkInterfaceTrafficConfig `json:"NetworkInterfaceTrafficConfig,omitempty" xml:"NetworkInterfaceTrafficConfig,omitempty" type:"Struct"`
	// The communication mode of the elastic network interface. Valid values:
	//
	// - `Standard`: Uses TCP communication.
	//
	// - `HighPerformance`: Uses the Elastic RDMA Interface (ERI) for RDMA communication.
	//
	// > The `HighPerformance` value is supported only by RDMA-enhanced instances, such as the c7re family.
	//
	// example:
	//
	// Standard
	NetworkInterfaceTrafficMode *string `json:"NetworkInterfaceTrafficMode,omitempty" xml:"NetworkInterfaceTrafficMode,omitempty"`
	// The ID of the account to which the elastic network interface belongs.
	//
	// example:
	//
	// 123456****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The primary private IP address of the elastic network interface.
	//
	// example:
	//
	// ``10.1.**.**``
	PrivateIpAddress *string                                                     `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	PrivateIpSets    *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSets `json:"PrivateIpSets,omitempty" xml:"PrivateIpSets,omitempty" type:"Struct"`
	// The QoS settings.
	QoSConfig *DescribeNetworkInterfaceAttributeResponseBodyQoSConfig `json:"QoSConfig,omitempty" xml:"QoSConfig,omitempty" type:"Struct"`
	// The number of queues supported by the elastic network interface.
	//
	// - For a primary network interface, this parameter returns the default number of queues for the instance type.
	//
	// - For a secondary network interface:
	//
	//   - If the interface is in the `InUse` state:
	//
	//     - If the queue number was not modified, the default value for the instance type is returned.
	//
	//     - If the queue number was modified, the new value is returned.
	//
	//   - If the secondary network interface is in the `Available` state:
	//
	//     - If the queue number was not modified, this parameter is not returned.
	//
	//     - If the queue number was modified, the new value is returned.
	//
	// example:
	//
	// 8
	QueueNumber *int32 `json:"QueueNumber,omitempty" xml:"QueueNumber,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// 22
	QueuePairNumber *int32 `json:"QueuePairNumber,omitempty" xml:"QueuePairNumber,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the enterprise resource group to which the elastic network interface belongs. If you use this parameter to filter resources, the number of resources cannot exceed 1,000.
	//
	// > Resources in the default resource group cannot be filtered.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId  *string                                                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityGroupIds *DescribeNetworkInterfaceAttributeResponseBodySecurityGroupIds `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Struct"`
	// The ID of the virtual service provider (VSP) for the elastic network interface.
	//
	// example:
	//
	// 12345678910
	ServiceID *int64 `json:"ServiceID,omitempty" xml:"ServiceID,omitempty"`
	// Indicates whether the elastic network interface is managed by an Alibaba Cloud service or a VSP.
	//
	// example:
	//
	// true
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	SlaveInterfaceSpecification *DescribeNetworkInterfaceAttributeResponseBodySlaveInterfaceSpecification `json:"SlaveInterfaceSpecification,omitempty" xml:"SlaveInterfaceSpecification,omitempty" type:"Struct"`
	// This parameter is not publicly available.
	//
	// example:
	//
	// false
	SourceDestCheck *bool `json:"SourceDestCheck,omitempty" xml:"SourceDestCheck,omitempty"`
	// The status of the elastic network interface. Valid values:
	//
	// - `Available`: The elastic network interface is available.
	//
	// - `Attaching`: The elastic network interface is being attached.
	//
	// - `InUse`: The elastic network interface is attached.
	//
	// - `Detaching`: The elastic network interface is being detached.
	//
	// - `Deleting`: The elastic network interface is being deleted.
	//
	// example:
	//
	// Available
	Status *string                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   *DescribeNetworkInterfaceAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// null
	TcpOptionAddressEnabled *string `json:"TcpOptionAddressEnabled,omitempty" xml:"TcpOptionAddressEnabled,omitempty"`
	// The type of the elastic network interface. Valid values:
	//
	// - `Primary`: The primary network interface.
	//
	// - `Secondary`: The secondary network interface.
	//
	// example:
	//
	// Secondary
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the vSwitch to which the elastic network interface is connected.
	//
	// example:
	//
	// vsw-bp1s5fnvk4gn2tws0****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC to which the elastic network interface belongs.
	//
	// example:
	//
	// vpc-bp67acfmxazb4p****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeNetworkInterfaceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetAssociatedPublicIp() *DescribeNetworkInterfaceAttributeResponseBodyAssociatedPublicIp {
	return s.AssociatedPublicIp
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetAttachment() *DescribeNetworkInterfaceAttributeResponseBodyAttachment {
	return s.Attachment
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetBondInterfaceSpecification() *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecification {
	return s.BondInterfaceSpecification
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetConnectionTrackingConfiguration() *DescribeNetworkInterfaceAttributeResponseBodyConnectionTrackingConfiguration {
	return s.ConnectionTrackingConfiguration
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetDeleteOnRelease() *bool {
	return s.DeleteOnRelease
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetEnhancedNetwork() *DescribeNetworkInterfaceAttributeResponseBodyEnhancedNetwork {
	return s.EnhancedNetwork
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetIpv4PrefixSets() *DescribeNetworkInterfaceAttributeResponseBodyIpv4PrefixSets {
	return s.Ipv4PrefixSets
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetIpv6PrefixSets() *DescribeNetworkInterfaceAttributeResponseBodyIpv6PrefixSets {
	return s.Ipv6PrefixSets
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetIpv6Sets() *DescribeNetworkInterfaceAttributeResponseBodyIpv6Sets {
	return s.Ipv6Sets
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetMacAddress() *string {
	return s.MacAddress
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetNetworkInterfaceName() *string {
	return s.NetworkInterfaceName
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetNetworkInterfaceTrafficConfig() *DescribeNetworkInterfaceAttributeResponseBodyNetworkInterfaceTrafficConfig {
	return s.NetworkInterfaceTrafficConfig
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetNetworkInterfaceTrafficMode() *string {
	return s.NetworkInterfaceTrafficMode
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetPrivateIpSets() *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSets {
	return s.PrivateIpSets
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetQoSConfig() *DescribeNetworkInterfaceAttributeResponseBodyQoSConfig {
	return s.QoSConfig
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetQueueNumber() *int32 {
	return s.QueueNumber
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetQueuePairNumber() *int32 {
	return s.QueuePairNumber
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetSecurityGroupIds() *DescribeNetworkInterfaceAttributeResponseBodySecurityGroupIds {
	return s.SecurityGroupIds
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetServiceID() *int64 {
	return s.ServiceID
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetSlaveInterfaceSpecification() *DescribeNetworkInterfaceAttributeResponseBodySlaveInterfaceSpecification {
	return s.SlaveInterfaceSpecification
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetSourceDestCheck() *bool {
	return s.SourceDestCheck
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetTags() *DescribeNetworkInterfaceAttributeResponseBodyTags {
	return s.Tags
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetTcpOptionAddressEnabled() *string {
	return s.TcpOptionAddressEnabled
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetAssociatedPublicIp(v *DescribeNetworkInterfaceAttributeResponseBodyAssociatedPublicIp) *DescribeNetworkInterfaceAttributeResponseBody {
	s.AssociatedPublicIp = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetAttachment(v *DescribeNetworkInterfaceAttributeResponseBodyAttachment) *DescribeNetworkInterfaceAttributeResponseBody {
	s.Attachment = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetBondInterfaceSpecification(v *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecification) *DescribeNetworkInterfaceAttributeResponseBody {
	s.BondInterfaceSpecification = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetConnectionTrackingConfiguration(v *DescribeNetworkInterfaceAttributeResponseBodyConnectionTrackingConfiguration) *DescribeNetworkInterfaceAttributeResponseBody {
	s.ConnectionTrackingConfiguration = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetCreationTime(v string) *DescribeNetworkInterfaceAttributeResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetDeleteOnRelease(v bool) *DescribeNetworkInterfaceAttributeResponseBody {
	s.DeleteOnRelease = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetDescription(v string) *DescribeNetworkInterfaceAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetEnhancedNetwork(v *DescribeNetworkInterfaceAttributeResponseBodyEnhancedNetwork) *DescribeNetworkInterfaceAttributeResponseBody {
	s.EnhancedNetwork = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetInstanceId(v string) *DescribeNetworkInterfaceAttributeResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetIpv4PrefixSets(v *DescribeNetworkInterfaceAttributeResponseBodyIpv4PrefixSets) *DescribeNetworkInterfaceAttributeResponseBody {
	s.Ipv4PrefixSets = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetIpv6PrefixSets(v *DescribeNetworkInterfaceAttributeResponseBodyIpv6PrefixSets) *DescribeNetworkInterfaceAttributeResponseBody {
	s.Ipv6PrefixSets = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetIpv6Sets(v *DescribeNetworkInterfaceAttributeResponseBodyIpv6Sets) *DescribeNetworkInterfaceAttributeResponseBody {
	s.Ipv6Sets = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetMacAddress(v string) *DescribeNetworkInterfaceAttributeResponseBody {
	s.MacAddress = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetNetworkInterfaceId(v string) *DescribeNetworkInterfaceAttributeResponseBody {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetNetworkInterfaceName(v string) *DescribeNetworkInterfaceAttributeResponseBody {
	s.NetworkInterfaceName = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetNetworkInterfaceTrafficConfig(v *DescribeNetworkInterfaceAttributeResponseBodyNetworkInterfaceTrafficConfig) *DescribeNetworkInterfaceAttributeResponseBody {
	s.NetworkInterfaceTrafficConfig = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetNetworkInterfaceTrafficMode(v string) *DescribeNetworkInterfaceAttributeResponseBody {
	s.NetworkInterfaceTrafficMode = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetOwnerId(v string) *DescribeNetworkInterfaceAttributeResponseBody {
	s.OwnerId = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetPrivateIpAddress(v string) *DescribeNetworkInterfaceAttributeResponseBody {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetPrivateIpSets(v *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSets) *DescribeNetworkInterfaceAttributeResponseBody {
	s.PrivateIpSets = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetQoSConfig(v *DescribeNetworkInterfaceAttributeResponseBodyQoSConfig) *DescribeNetworkInterfaceAttributeResponseBody {
	s.QoSConfig = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetQueueNumber(v int32) *DescribeNetworkInterfaceAttributeResponseBody {
	s.QueueNumber = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetQueuePairNumber(v int32) *DescribeNetworkInterfaceAttributeResponseBody {
	s.QueuePairNumber = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetRequestId(v string) *DescribeNetworkInterfaceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetResourceGroupId(v string) *DescribeNetworkInterfaceAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetSecurityGroupIds(v *DescribeNetworkInterfaceAttributeResponseBodySecurityGroupIds) *DescribeNetworkInterfaceAttributeResponseBody {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetServiceID(v int64) *DescribeNetworkInterfaceAttributeResponseBody {
	s.ServiceID = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetServiceManaged(v bool) *DescribeNetworkInterfaceAttributeResponseBody {
	s.ServiceManaged = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetSlaveInterfaceSpecification(v *DescribeNetworkInterfaceAttributeResponseBodySlaveInterfaceSpecification) *DescribeNetworkInterfaceAttributeResponseBody {
	s.SlaveInterfaceSpecification = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetSourceDestCheck(v bool) *DescribeNetworkInterfaceAttributeResponseBody {
	s.SourceDestCheck = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetStatus(v string) *DescribeNetworkInterfaceAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetTags(v *DescribeNetworkInterfaceAttributeResponseBodyTags) *DescribeNetworkInterfaceAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetTcpOptionAddressEnabled(v string) *DescribeNetworkInterfaceAttributeResponseBody {
	s.TcpOptionAddressEnabled = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetType(v string) *DescribeNetworkInterfaceAttributeResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetVSwitchId(v string) *DescribeNetworkInterfaceAttributeResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetVpcId(v string) *DescribeNetworkInterfaceAttributeResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) SetZoneId(v string) *DescribeNetworkInterfaceAttributeResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBody) Validate() error {
	if s.AssociatedPublicIp != nil {
		if err := s.AssociatedPublicIp.Validate(); err != nil {
			return err
		}
	}
	if s.Attachment != nil {
		if err := s.Attachment.Validate(); err != nil {
			return err
		}
	}
	if s.BondInterfaceSpecification != nil {
		if err := s.BondInterfaceSpecification.Validate(); err != nil {
			return err
		}
	}
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
	if s.Ipv4PrefixSets != nil {
		if err := s.Ipv4PrefixSets.Validate(); err != nil {
			return err
		}
	}
	if s.Ipv6PrefixSets != nil {
		if err := s.Ipv6PrefixSets.Validate(); err != nil {
			return err
		}
	}
	if s.Ipv6Sets != nil {
		if err := s.Ipv6Sets.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkInterfaceTrafficConfig != nil {
		if err := s.NetworkInterfaceTrafficConfig.Validate(); err != nil {
			return err
		}
	}
	if s.PrivateIpSets != nil {
		if err := s.PrivateIpSets.Validate(); err != nil {
			return err
		}
	}
	if s.QoSConfig != nil {
		if err := s.QoSConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityGroupIds != nil {
		if err := s.SecurityGroupIds.Validate(); err != nil {
			return err
		}
	}
	if s.SlaveInterfaceSpecification != nil {
		if err := s.SlaveInterfaceSpecification.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNetworkInterfaceAttributeResponseBodyAssociatedPublicIp struct {
	// The ID of the elastic IP address.
	//
	// example:
	//
	// null
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The public IP address.
	//
	// example:
	//
	// ``116.62.**.**``
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
}

func (s DescribeNetworkInterfaceAttributeResponseBodyAssociatedPublicIp) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponseBodyAssociatedPublicIp) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyAssociatedPublicIp) GetAllocationId() *string {
	return s.AllocationId
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyAssociatedPublicIp) GetPublicIpAddress() *string {
	return s.PublicIpAddress
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyAssociatedPublicIp) SetAllocationId(v string) *DescribeNetworkInterfaceAttributeResponseBodyAssociatedPublicIp {
	s.AllocationId = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyAssociatedPublicIp) SetPublicIpAddress(v string) *DescribeNetworkInterfaceAttributeResponseBodyAssociatedPublicIp {
	s.PublicIpAddress = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyAssociatedPublicIp) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfaceAttributeResponseBodyAttachment struct {
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// hide
	DeviceIndex *int32 `json:"DeviceIndex,omitempty" xml:"DeviceIndex,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// hide
	InstanceId                *string                                                                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MemberNetworkInterfaceIds *DescribeNetworkInterfaceAttributeResponseBodyAttachmentMemberNetworkInterfaceIds `json:"MemberNetworkInterfaceIds,omitempty" xml:"MemberNetworkInterfaceIds,omitempty" type:"Struct"`
	// The index of the physical network card to which the elastic network interface is attached.
	//
	// - This parameter is not returned if the elastic network interface is `Available`, or if no index was specified during attachment.
	//
	// - If the elastic network interface is `InUse` and an index was specified during attachment, this parameter returns the index of the physical network card.
	//
	// example:
	//
	// 0
	NetworkCardIndex *int32 `json:"NetworkCardIndex,omitempty" xml:"NetworkCardIndex,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// hide
	TrunkNetworkInterfaceId *string `json:"TrunkNetworkInterfaceId,omitempty" xml:"TrunkNetworkInterfaceId,omitempty"`
}

func (s DescribeNetworkInterfaceAttributeResponseBodyAttachment) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponseBodyAttachment) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyAttachment) GetDeviceIndex() *int32 {
	return s.DeviceIndex
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyAttachment) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyAttachment) GetMemberNetworkInterfaceIds() *DescribeNetworkInterfaceAttributeResponseBodyAttachmentMemberNetworkInterfaceIds {
	return s.MemberNetworkInterfaceIds
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyAttachment) GetNetworkCardIndex() *int32 {
	return s.NetworkCardIndex
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyAttachment) GetTrunkNetworkInterfaceId() *string {
	return s.TrunkNetworkInterfaceId
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyAttachment) SetDeviceIndex(v int32) *DescribeNetworkInterfaceAttributeResponseBodyAttachment {
	s.DeviceIndex = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyAttachment) SetInstanceId(v string) *DescribeNetworkInterfaceAttributeResponseBodyAttachment {
	s.InstanceId = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyAttachment) SetMemberNetworkInterfaceIds(v *DescribeNetworkInterfaceAttributeResponseBodyAttachmentMemberNetworkInterfaceIds) *DescribeNetworkInterfaceAttributeResponseBodyAttachment {
	s.MemberNetworkInterfaceIds = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyAttachment) SetNetworkCardIndex(v int32) *DescribeNetworkInterfaceAttributeResponseBodyAttachment {
	s.NetworkCardIndex = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyAttachment) SetTrunkNetworkInterfaceId(v string) *DescribeNetworkInterfaceAttributeResponseBodyAttachment {
	s.TrunkNetworkInterfaceId = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyAttachment) Validate() error {
	if s.MemberNetworkInterfaceIds != nil {
		if err := s.MemberNetworkInterfaceIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNetworkInterfaceAttributeResponseBodyAttachmentMemberNetworkInterfaceIds struct {
	MemberNetworkInterfaceId []*string `json:"MemberNetworkInterfaceId,omitempty" xml:"MemberNetworkInterfaceId,omitempty" type:"Repeated"`
}

func (s DescribeNetworkInterfaceAttributeResponseBodyAttachmentMemberNetworkInterfaceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponseBodyAttachmentMemberNetworkInterfaceIds) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyAttachmentMemberNetworkInterfaceIds) GetMemberNetworkInterfaceId() []*string {
	return s.MemberNetworkInterfaceId
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyAttachmentMemberNetworkInterfaceIds) SetMemberNetworkInterfaceId(v []*string) *DescribeNetworkInterfaceAttributeResponseBodyAttachmentMemberNetworkInterfaceIds {
	s.MemberNetworkInterfaceId = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyAttachmentMemberNetworkInterfaceIds) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecification struct {
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// null
	BondMode                    *string                                                                                             `json:"BondMode,omitempty" xml:"BondMode,omitempty"`
	SlaveInterfaceSpecification *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecification `json:"SlaveInterfaceSpecification,omitempty" xml:"SlaveInterfaceSpecification,omitempty" type:"Struct"`
}

func (s DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecification) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecification) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecification) GetBondMode() *string {
	return s.BondMode
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecification) GetSlaveInterfaceSpecification() *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecification {
	return s.SlaveInterfaceSpecification
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecification) SetBondMode(v string) *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecification {
	s.BondMode = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecification) SetSlaveInterfaceSpecification(v *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecification) *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecification {
	s.SlaveInterfaceSpecification = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecification) Validate() error {
	if s.SlaveInterfaceSpecification != nil {
		if err := s.SlaveInterfaceSpecification.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecification struct {
	SlaveInterfaceSpecificationSet []*DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecificationSlaveInterfaceSpecificationSet `json:"SlaveInterfaceSpecificationSet,omitempty" xml:"SlaveInterfaceSpecificationSet,omitempty" type:"Repeated"`
}

func (s DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecification) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecification) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecification) GetSlaveInterfaceSpecificationSet() []*DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecificationSlaveInterfaceSpecificationSet {
	return s.SlaveInterfaceSpecificationSet
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecification) SetSlaveInterfaceSpecificationSet(v []*DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecificationSlaveInterfaceSpecificationSet) *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecification {
	s.SlaveInterfaceSpecificationSet = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecification) Validate() error {
	if s.SlaveInterfaceSpecificationSet != nil {
		for _, item := range s.SlaveInterfaceSpecificationSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecificationSlaveInterfaceSpecificationSet struct {
	BondNetworkInterfaceId  *string `json:"BondNetworkInterfaceId,omitempty" xml:"BondNetworkInterfaceId,omitempty"`
	SlaveNetworkInterfaceId *string `json:"SlaveNetworkInterfaceId,omitempty" xml:"SlaveNetworkInterfaceId,omitempty"`
	WorkState               *string `json:"WorkState,omitempty" xml:"WorkState,omitempty"`
}

func (s DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecificationSlaveInterfaceSpecificationSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecificationSlaveInterfaceSpecificationSet) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecificationSlaveInterfaceSpecificationSet) GetBondNetworkInterfaceId() *string {
	return s.BondNetworkInterfaceId
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecificationSlaveInterfaceSpecificationSet) GetSlaveNetworkInterfaceId() *string {
	return s.SlaveNetworkInterfaceId
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecificationSlaveInterfaceSpecificationSet) GetWorkState() *string {
	return s.WorkState
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecificationSlaveInterfaceSpecificationSet) SetBondNetworkInterfaceId(v string) *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecificationSlaveInterfaceSpecificationSet {
	s.BondNetworkInterfaceId = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecificationSlaveInterfaceSpecificationSet) SetSlaveNetworkInterfaceId(v string) *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecificationSlaveInterfaceSpecificationSet {
	s.SlaveNetworkInterfaceId = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecificationSlaveInterfaceSpecificationSet) SetWorkState(v string) *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecificationSlaveInterfaceSpecificationSet {
	s.WorkState = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyBondInterfaceSpecificationSlaveInterfaceSpecificationSlaveInterfaceSpecificationSet) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfaceAttributeResponseBodyConnectionTrackingConfiguration struct {
	// The timeout period for TCP connections in the `TIME_WAIT` and `FIN-WAIT-2` states. Unit: seconds. Valid values: an integer from 3 to 15.
	//
	// > For ECS instances used with a Network Load Balancer (NLB) or Classic Load Balancer (CLB), the default timeout for connections in the `TIME_WAIT` state is 15 seconds.
	//
	// example:
	//
	// 3
	TcpClosedAndTimeWaitTimeout *int32 `json:"TcpClosedAndTimeWaitTimeout,omitempty" xml:"TcpClosedAndTimeWaitTimeout,omitempty"`
	// The timeout period for established TCP connections. Unit: seconds. Valid values: 30, 60, 80, 100, 200, 300, 500, 700, and 910.
	//
	// example:
	//
	// 910
	TcpEstablishedTimeout *int32 `json:"TcpEstablishedTimeout,omitempty" xml:"TcpEstablishedTimeout,omitempty"`
	// The timeout period for UDP streams. Unit: seconds. Valid values: 10, 20, 30, 60, 80, and 100.
	//
	// > For ECS instances used with a Network Load Balancer (NLB) or Classic Load Balancer (CLB), the default UDP timeout is 100 seconds.
	//
	// example:
	//
	// 30
	UdpTimeout *int32 `json:"UdpTimeout,omitempty" xml:"UdpTimeout,omitempty"`
}

func (s DescribeNetworkInterfaceAttributeResponseBodyConnectionTrackingConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponseBodyConnectionTrackingConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyConnectionTrackingConfiguration) GetTcpClosedAndTimeWaitTimeout() *int32 {
	return s.TcpClosedAndTimeWaitTimeout
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyConnectionTrackingConfiguration) GetTcpEstablishedTimeout() *int32 {
	return s.TcpEstablishedTimeout
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyConnectionTrackingConfiguration) GetUdpTimeout() *int32 {
	return s.UdpTimeout
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyConnectionTrackingConfiguration) SetTcpClosedAndTimeWaitTimeout(v int32) *DescribeNetworkInterfaceAttributeResponseBodyConnectionTrackingConfiguration {
	s.TcpClosedAndTimeWaitTimeout = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyConnectionTrackingConfiguration) SetTcpEstablishedTimeout(v int32) *DescribeNetworkInterfaceAttributeResponseBodyConnectionTrackingConfiguration {
	s.TcpEstablishedTimeout = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyConnectionTrackingConfiguration) SetUdpTimeout(v int32) *DescribeNetworkInterfaceAttributeResponseBodyConnectionTrackingConfiguration {
	s.UdpTimeout = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyConnectionTrackingConfiguration) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfaceAttributeResponseBodyEnhancedNetwork struct {
	// > This parameter is not publicly available.
	//
	// example:
	//
	// true
	EnableRss *bool `json:"EnableRss,omitempty" xml:"EnableRss,omitempty"`
	// This parameter is not publicly available.
	//
	// example:
	//
	// false
	EnableSriov                     *bool  `json:"EnableSriov,omitempty" xml:"EnableSriov,omitempty"`
	VirtualFunctionQuantity         *int32 `json:"VirtualFunctionQuantity,omitempty" xml:"VirtualFunctionQuantity,omitempty"`
	VirtualFunctionTotalQueueNumber *int32 `json:"VirtualFunctionTotalQueueNumber,omitempty" xml:"VirtualFunctionTotalQueueNumber,omitempty"`
}

func (s DescribeNetworkInterfaceAttributeResponseBodyEnhancedNetwork) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponseBodyEnhancedNetwork) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyEnhancedNetwork) GetEnableRss() *bool {
	return s.EnableRss
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyEnhancedNetwork) GetEnableSriov() *bool {
	return s.EnableSriov
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyEnhancedNetwork) GetVirtualFunctionQuantity() *int32 {
	return s.VirtualFunctionQuantity
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyEnhancedNetwork) GetVirtualFunctionTotalQueueNumber() *int32 {
	return s.VirtualFunctionTotalQueueNumber
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyEnhancedNetwork) SetEnableRss(v bool) *DescribeNetworkInterfaceAttributeResponseBodyEnhancedNetwork {
	s.EnableRss = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyEnhancedNetwork) SetEnableSriov(v bool) *DescribeNetworkInterfaceAttributeResponseBodyEnhancedNetwork {
	s.EnableSriov = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyEnhancedNetwork) SetVirtualFunctionQuantity(v int32) *DescribeNetworkInterfaceAttributeResponseBodyEnhancedNetwork {
	s.VirtualFunctionQuantity = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyEnhancedNetwork) SetVirtualFunctionTotalQueueNumber(v int32) *DescribeNetworkInterfaceAttributeResponseBodyEnhancedNetwork {
	s.VirtualFunctionTotalQueueNumber = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyEnhancedNetwork) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfaceAttributeResponseBodyIpv4PrefixSets struct {
	Ipv4PrefixSet []*DescribeNetworkInterfaceAttributeResponseBodyIpv4PrefixSetsIpv4PrefixSet `json:"Ipv4PrefixSet,omitempty" xml:"Ipv4PrefixSet,omitempty" type:"Repeated"`
}

func (s DescribeNetworkInterfaceAttributeResponseBodyIpv4PrefixSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponseBodyIpv4PrefixSets) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyIpv4PrefixSets) GetIpv4PrefixSet() []*DescribeNetworkInterfaceAttributeResponseBodyIpv4PrefixSetsIpv4PrefixSet {
	return s.Ipv4PrefixSet
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyIpv4PrefixSets) SetIpv4PrefixSet(v []*DescribeNetworkInterfaceAttributeResponseBodyIpv4PrefixSetsIpv4PrefixSet) *DescribeNetworkInterfaceAttributeResponseBodyIpv4PrefixSets {
	s.Ipv4PrefixSet = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyIpv4PrefixSets) Validate() error {
	if s.Ipv4PrefixSet != nil {
		for _, item := range s.Ipv4PrefixSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkInterfaceAttributeResponseBodyIpv4PrefixSetsIpv4PrefixSet struct {
	Ipv4Prefix *string `json:"Ipv4Prefix,omitempty" xml:"Ipv4Prefix,omitempty"`
}

func (s DescribeNetworkInterfaceAttributeResponseBodyIpv4PrefixSetsIpv4PrefixSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponseBodyIpv4PrefixSetsIpv4PrefixSet) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyIpv4PrefixSetsIpv4PrefixSet) GetIpv4Prefix() *string {
	return s.Ipv4Prefix
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyIpv4PrefixSetsIpv4PrefixSet) SetIpv4Prefix(v string) *DescribeNetworkInterfaceAttributeResponseBodyIpv4PrefixSetsIpv4PrefixSet {
	s.Ipv4Prefix = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyIpv4PrefixSetsIpv4PrefixSet) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfaceAttributeResponseBodyIpv6PrefixSets struct {
	Ipv6PrefixSet []*DescribeNetworkInterfaceAttributeResponseBodyIpv6PrefixSetsIpv6PrefixSet `json:"Ipv6PrefixSet,omitempty" xml:"Ipv6PrefixSet,omitempty" type:"Repeated"`
}

func (s DescribeNetworkInterfaceAttributeResponseBodyIpv6PrefixSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponseBodyIpv6PrefixSets) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyIpv6PrefixSets) GetIpv6PrefixSet() []*DescribeNetworkInterfaceAttributeResponseBodyIpv6PrefixSetsIpv6PrefixSet {
	return s.Ipv6PrefixSet
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyIpv6PrefixSets) SetIpv6PrefixSet(v []*DescribeNetworkInterfaceAttributeResponseBodyIpv6PrefixSetsIpv6PrefixSet) *DescribeNetworkInterfaceAttributeResponseBodyIpv6PrefixSets {
	s.Ipv6PrefixSet = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyIpv6PrefixSets) Validate() error {
	if s.Ipv6PrefixSet != nil {
		for _, item := range s.Ipv6PrefixSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkInterfaceAttributeResponseBodyIpv6PrefixSetsIpv6PrefixSet struct {
	Ipv6Prefix *string `json:"Ipv6Prefix,omitempty" xml:"Ipv6Prefix,omitempty"`
}

func (s DescribeNetworkInterfaceAttributeResponseBodyIpv6PrefixSetsIpv6PrefixSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponseBodyIpv6PrefixSetsIpv6PrefixSet) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyIpv6PrefixSetsIpv6PrefixSet) GetIpv6Prefix() *string {
	return s.Ipv6Prefix
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyIpv6PrefixSetsIpv6PrefixSet) SetIpv6Prefix(v string) *DescribeNetworkInterfaceAttributeResponseBodyIpv6PrefixSetsIpv6PrefixSet {
	s.Ipv6Prefix = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyIpv6PrefixSetsIpv6PrefixSet) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfaceAttributeResponseBodyIpv6Sets struct {
	Ipv6Set []*DescribeNetworkInterfaceAttributeResponseBodyIpv6SetsIpv6Set `json:"Ipv6Set,omitempty" xml:"Ipv6Set,omitempty" type:"Repeated"`
}

func (s DescribeNetworkInterfaceAttributeResponseBodyIpv6Sets) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponseBodyIpv6Sets) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyIpv6Sets) GetIpv6Set() []*DescribeNetworkInterfaceAttributeResponseBodyIpv6SetsIpv6Set {
	return s.Ipv6Set
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyIpv6Sets) SetIpv6Set(v []*DescribeNetworkInterfaceAttributeResponseBodyIpv6SetsIpv6Set) *DescribeNetworkInterfaceAttributeResponseBodyIpv6Sets {
	s.Ipv6Set = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyIpv6Sets) Validate() error {
	if s.Ipv6Set != nil {
		for _, item := range s.Ipv6Set {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkInterfaceAttributeResponseBodyIpv6SetsIpv6Set struct {
	Ipv6Address *string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
}

func (s DescribeNetworkInterfaceAttributeResponseBodyIpv6SetsIpv6Set) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponseBodyIpv6SetsIpv6Set) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyIpv6SetsIpv6Set) GetIpv6Address() *string {
	return s.Ipv6Address
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyIpv6SetsIpv6Set) SetIpv6Address(v string) *DescribeNetworkInterfaceAttributeResponseBodyIpv6SetsIpv6Set {
	s.Ipv6Address = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyIpv6SetsIpv6Set) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfaceAttributeResponseBodyNetworkInterfaceTrafficConfig struct {
	// The communication mode of the elastic network interface.
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
}

func (s DescribeNetworkInterfaceAttributeResponseBodyNetworkInterfaceTrafficConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponseBodyNetworkInterfaceTrafficConfig) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyNetworkInterfaceTrafficConfig) GetNetworkInterfaceTrafficMode() *string {
	return s.NetworkInterfaceTrafficMode
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyNetworkInterfaceTrafficConfig) GetQueueNumber() *int32 {
	return s.QueueNumber
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyNetworkInterfaceTrafficConfig) GetQueuePairNumber() *int32 {
	return s.QueuePairNumber
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyNetworkInterfaceTrafficConfig) SetNetworkInterfaceTrafficMode(v string) *DescribeNetworkInterfaceAttributeResponseBodyNetworkInterfaceTrafficConfig {
	s.NetworkInterfaceTrafficMode = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyNetworkInterfaceTrafficConfig) SetQueueNumber(v int32) *DescribeNetworkInterfaceAttributeResponseBodyNetworkInterfaceTrafficConfig {
	s.QueueNumber = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyNetworkInterfaceTrafficConfig) SetQueuePairNumber(v int32) *DescribeNetworkInterfaceAttributeResponseBodyNetworkInterfaceTrafficConfig {
	s.QueuePairNumber = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyNetworkInterfaceTrafficConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSets struct {
	PrivateIpSet []*DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSet `json:"PrivateIpSet,omitempty" xml:"PrivateIpSet,omitempty" type:"Repeated"`
}

func (s DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSets) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSets) GetPrivateIpSet() []*DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSet {
	return s.PrivateIpSet
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSets) SetPrivateIpSet(v []*DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSet) *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSets {
	s.PrivateIpSet = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSets) Validate() error {
	if s.PrivateIpSet != nil {
		for _, item := range s.PrivateIpSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSet struct {
	AssociatedPublicIp *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSetAssociatedPublicIp `json:"AssociatedPublicIp,omitempty" xml:"AssociatedPublicIp,omitempty" type:"Struct"`
	Primary            *bool                                                                                     `json:"Primary,omitempty" xml:"Primary,omitempty"`
	PrivateIpAddress   *string                                                                                   `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
}

func (s DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSet) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSet) GetAssociatedPublicIp() *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSetAssociatedPublicIp {
	return s.AssociatedPublicIp
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSet) GetPrimary() *bool {
	return s.Primary
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSet) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSet) SetAssociatedPublicIp(v *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSetAssociatedPublicIp) *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSet {
	s.AssociatedPublicIp = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSet) SetPrimary(v bool) *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSet {
	s.Primary = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSet) SetPrivateIpAddress(v string) *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSet {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSet) Validate() error {
	if s.AssociatedPublicIp != nil {
		if err := s.AssociatedPublicIp.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSetAssociatedPublicIp struct {
	AllocationId    *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
}

func (s DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSetAssociatedPublicIp) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSetAssociatedPublicIp) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSetAssociatedPublicIp) GetAllocationId() *string {
	return s.AllocationId
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSetAssociatedPublicIp) GetPublicIpAddress() *string {
	return s.PublicIpAddress
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSetAssociatedPublicIp) SetAllocationId(v string) *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSetAssociatedPublicIp {
	s.AllocationId = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSetAssociatedPublicIp) SetPublicIpAddress(v string) *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSetAssociatedPublicIp {
	s.PublicIpAddress = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyPrivateIpSetsPrivateIpSetAssociatedPublicIp) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfaceAttributeResponseBodyQoSConfig struct {
	// Indicates whether QoS is enabled.
	EnableQoS *bool `json:"EnableQoS,omitempty" xml:"EnableQoS,omitempty"`
	// The QoS settings.
	QoS *DescribeNetworkInterfaceAttributeResponseBodyQoSConfigQoS `json:"QoS,omitempty" xml:"QoS,omitempty" type:"Struct"`
}

func (s DescribeNetworkInterfaceAttributeResponseBodyQoSConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponseBodyQoSConfig) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyQoSConfig) GetEnableQoS() *bool {
	return s.EnableQoS
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyQoSConfig) GetQoS() *DescribeNetworkInterfaceAttributeResponseBodyQoSConfigQoS {
	return s.QoS
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyQoSConfig) SetEnableQoS(v bool) *DescribeNetworkInterfaceAttributeResponseBodyQoSConfig {
	s.EnableQoS = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyQoSConfig) SetQoS(v *DescribeNetworkInterfaceAttributeResponseBodyQoSConfigQoS) *DescribeNetworkInterfaceAttributeResponseBodyQoSConfig {
	s.QoS = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyQoSConfig) Validate() error {
	if s.QoS != nil {
		if err := s.QoS.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNetworkInterfaceAttributeResponseBodyQoSConfigQoS struct {
	// The maximum inbound bandwidth on the internal network.
	//
	// example:
	//
	// 50000
	BandwidthRx *int64 `json:"BandwidthRx,omitempty" xml:"BandwidthRx,omitempty"`
	// The maximum outbound bandwidth on the internal network.
	//
	// example:
	//
	// 50000
	BandwidthTx *int64 `json:"BandwidthTx,omitempty" xml:"BandwidthTx,omitempty"`
	// The maximum number of connections.
	//
	// example:
	//
	// 50000
	ConcurrentConnections *int64 `json:"ConcurrentConnections,omitempty" xml:"ConcurrentConnections,omitempty"`
	// The inbound packet transmission rate on the internal network. Unit: packets per second (pps).
	//
	// example:
	//
	// 50000
	PpsRx *int64 `json:"PpsRx,omitempty" xml:"PpsRx,omitempty"`
	// The outbound packet transmission rate on the internal network. Unit: packets per second (pps).
	//
	// example:
	//
	// 50000
	PpsTx *int64 `json:"PpsTx,omitempty" xml:"PpsTx,omitempty"`
}

func (s DescribeNetworkInterfaceAttributeResponseBodyQoSConfigQoS) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponseBodyQoSConfigQoS) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyQoSConfigQoS) GetBandwidthRx() *int64 {
	return s.BandwidthRx
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyQoSConfigQoS) GetBandwidthTx() *int64 {
	return s.BandwidthTx
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyQoSConfigQoS) GetConcurrentConnections() *int64 {
	return s.ConcurrentConnections
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyQoSConfigQoS) GetPpsRx() *int64 {
	return s.PpsRx
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyQoSConfigQoS) GetPpsTx() *int64 {
	return s.PpsTx
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyQoSConfigQoS) SetBandwidthRx(v int64) *DescribeNetworkInterfaceAttributeResponseBodyQoSConfigQoS {
	s.BandwidthRx = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyQoSConfigQoS) SetBandwidthTx(v int64) *DescribeNetworkInterfaceAttributeResponseBodyQoSConfigQoS {
	s.BandwidthTx = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyQoSConfigQoS) SetConcurrentConnections(v int64) *DescribeNetworkInterfaceAttributeResponseBodyQoSConfigQoS {
	s.ConcurrentConnections = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyQoSConfigQoS) SetPpsRx(v int64) *DescribeNetworkInterfaceAttributeResponseBodyQoSConfigQoS {
	s.PpsRx = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyQoSConfigQoS) SetPpsTx(v int64) *DescribeNetworkInterfaceAttributeResponseBodyQoSConfigQoS {
	s.PpsTx = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyQoSConfigQoS) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfaceAttributeResponseBodySecurityGroupIds struct {
	SecurityGroupId []*string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Repeated"`
}

func (s DescribeNetworkInterfaceAttributeResponseBodySecurityGroupIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponseBodySecurityGroupIds) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponseBodySecurityGroupIds) GetSecurityGroupId() []*string {
	return s.SecurityGroupId
}

func (s *DescribeNetworkInterfaceAttributeResponseBodySecurityGroupIds) SetSecurityGroupId(v []*string) *DescribeNetworkInterfaceAttributeResponseBodySecurityGroupIds {
	s.SecurityGroupId = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodySecurityGroupIds) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfaceAttributeResponseBodySlaveInterfaceSpecification struct {
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// null
	BondNetworkInterfaceId *string `json:"BondNetworkInterfaceId,omitempty" xml:"BondNetworkInterfaceId,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// null
	SlaveNetworkInterfaceId *string `json:"SlaveNetworkInterfaceId,omitempty" xml:"SlaveNetworkInterfaceId,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// null
	WorkState *string `json:"WorkState,omitempty" xml:"WorkState,omitempty"`
}

func (s DescribeNetworkInterfaceAttributeResponseBodySlaveInterfaceSpecification) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponseBodySlaveInterfaceSpecification) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponseBodySlaveInterfaceSpecification) GetBondNetworkInterfaceId() *string {
	return s.BondNetworkInterfaceId
}

func (s *DescribeNetworkInterfaceAttributeResponseBodySlaveInterfaceSpecification) GetSlaveNetworkInterfaceId() *string {
	return s.SlaveNetworkInterfaceId
}

func (s *DescribeNetworkInterfaceAttributeResponseBodySlaveInterfaceSpecification) GetWorkState() *string {
	return s.WorkState
}

func (s *DescribeNetworkInterfaceAttributeResponseBodySlaveInterfaceSpecification) SetBondNetworkInterfaceId(v string) *DescribeNetworkInterfaceAttributeResponseBodySlaveInterfaceSpecification {
	s.BondNetworkInterfaceId = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodySlaveInterfaceSpecification) SetSlaveNetworkInterfaceId(v string) *DescribeNetworkInterfaceAttributeResponseBodySlaveInterfaceSpecification {
	s.SlaveNetworkInterfaceId = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodySlaveInterfaceSpecification) SetWorkState(v string) *DescribeNetworkInterfaceAttributeResponseBodySlaveInterfaceSpecification {
	s.WorkState = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodySlaveInterfaceSpecification) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfaceAttributeResponseBodyTags struct {
	Tag []*DescribeNetworkInterfaceAttributeResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeNetworkInterfaceAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyTags) GetTag() []*DescribeNetworkInterfaceAttributeResponseBodyTagsTag {
	return s.Tag
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyTags) SetTag(v []*DescribeNetworkInterfaceAttributeResponseBodyTagsTag) *DescribeNetworkInterfaceAttributeResponseBodyTags {
	s.Tag = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyTags) Validate() error {
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

type DescribeNetworkInterfaceAttributeResponseBodyTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeNetworkInterfaceAttributeResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyTagsTag) SetTagKey(v string) *DescribeNetworkInterfaceAttributeResponseBodyTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyTagsTag) SetTagValue(v string) *DescribeNetworkInterfaceAttributeResponseBodyTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponseBodyTagsTag) Validate() error {
	return dara.Validate(s)
}
