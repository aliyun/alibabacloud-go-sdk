// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *DescribeRCInstanceAttributeResponseBody
	GetAutoRenew() *bool
	SetClusterId(v string) *DescribeRCInstanceAttributeResponseBody
	GetClusterId() *string
	SetCpu(v int32) *DescribeRCInstanceAttributeResponseBody
	GetCpu() *int32
	SetCreateMode(v int32) *DescribeRCInstanceAttributeResponseBody
	GetCreateMode() *int32
	SetCreationTime(v string) *DescribeRCInstanceAttributeResponseBody
	GetCreationTime() *string
	SetCreditSpecification(v string) *DescribeRCInstanceAttributeResponseBody
	GetCreditSpecification() *string
	SetDataDisks(v *DescribeRCInstanceAttributeResponseBodyDataDisks) *DescribeRCInstanceAttributeResponseBody
	GetDataDisks() *DescribeRCInstanceAttributeResponseBodyDataDisks
	SetDbType(v string) *DescribeRCInstanceAttributeResponseBody
	GetDbType() *string
	SetDedicatedHostAttribute(v *DescribeRCInstanceAttributeResponseBodyDedicatedHostAttribute) *DescribeRCInstanceAttributeResponseBody
	GetDedicatedHostAttribute() *DescribeRCInstanceAttributeResponseBodyDedicatedHostAttribute
	SetDeletionProtection(v bool) *DescribeRCInstanceAttributeResponseBody
	GetDeletionProtection() *bool
	SetDeploymentSetId(v string) *DescribeRCInstanceAttributeResponseBody
	GetDeploymentSetId() *string
	SetDescription(v string) *DescribeRCInstanceAttributeResponseBody
	GetDescription() *string
	SetDiskType(v string) *DescribeRCInstanceAttributeResponseBody
	GetDiskType() *string
	SetEcsInstanceType(v string) *DescribeRCInstanceAttributeResponseBody
	GetEcsInstanceType() *string
	SetEipAddress(v *DescribeRCInstanceAttributeResponseBodyEipAddress) *DescribeRCInstanceAttributeResponseBody
	GetEipAddress() *DescribeRCInstanceAttributeResponseBodyEipAddress
	SetEnableJumboFrame(v bool) *DescribeRCInstanceAttributeResponseBody
	GetEnableJumboFrame() *bool
	SetExpiredTime(v string) *DescribeRCInstanceAttributeResponseBody
	GetExpiredTime() *string
	SetGpu(v int32) *DescribeRCInstanceAttributeResponseBody
	GetGpu() *int32
	SetGpuTypes(v string) *DescribeRCInstanceAttributeResponseBody
	GetGpuTypes() *string
	SetHostName(v string) *DescribeRCInstanceAttributeResponseBody
	GetHostName() *string
	SetHostType(v string) *DescribeRCInstanceAttributeResponseBody
	GetHostType() *string
	SetImageId(v string) *DescribeRCInstanceAttributeResponseBody
	GetImageId() *string
	SetInnerIpAddress(v *DescribeRCInstanceAttributeResponseBodyInnerIpAddress) *DescribeRCInstanceAttributeResponseBody
	GetInnerIpAddress() *DescribeRCInstanceAttributeResponseBodyInnerIpAddress
	SetInstanceChargeType(v string) *DescribeRCInstanceAttributeResponseBody
	GetInstanceChargeType() *string
	SetInstanceId(v string) *DescribeRCInstanceAttributeResponseBody
	GetInstanceId() *string
	SetInstanceName(v string) *DescribeRCInstanceAttributeResponseBody
	GetInstanceName() *string
	SetInstanceNetworkType(v string) *DescribeRCInstanceAttributeResponseBody
	GetInstanceNetworkType() *string
	SetInstanceType(v string) *DescribeRCInstanceAttributeResponseBody
	GetInstanceType() *string
	SetInternetChargeType(v string) *DescribeRCInstanceAttributeResponseBody
	GetInternetChargeType() *string
	SetInternetMaxBandwidthIn(v int32) *DescribeRCInstanceAttributeResponseBody
	GetInternetMaxBandwidthIn() *int32
	SetInternetMaxBandwidthOut(v int32) *DescribeRCInstanceAttributeResponseBody
	GetInternetMaxBandwidthOut() *int32
	SetIoOptimized(v string) *DescribeRCInstanceAttributeResponseBody
	GetIoOptimized() *string
	SetKeyPairName(v string) *DescribeRCInstanceAttributeResponseBody
	GetKeyPairName() *string
	SetMemory(v int32) *DescribeRCInstanceAttributeResponseBody
	GetMemory() *int32
	SetNodeType(v string) *DescribeRCInstanceAttributeResponseBody
	GetNodeType() *string
	SetOperationLocks(v *DescribeRCInstanceAttributeResponseBodyOperationLocks) *DescribeRCInstanceAttributeResponseBody
	GetOperationLocks() *DescribeRCInstanceAttributeResponseBodyOperationLocks
	SetPublicIpAddress(v *DescribeRCInstanceAttributeResponseBodyPublicIpAddress) *DescribeRCInstanceAttributeResponseBody
	GetPublicIpAddress() *DescribeRCInstanceAttributeResponseBodyPublicIpAddress
	SetRegionId(v string) *DescribeRCInstanceAttributeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeRCInstanceAttributeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeRCInstanceAttributeResponseBody
	GetResourceGroupId() *string
	SetSecurityGroupIds(v *DescribeRCInstanceAttributeResponseBodySecurityGroupIds) *DescribeRCInstanceAttributeResponseBody
	GetSecurityGroupIds() *DescribeRCInstanceAttributeResponseBodySecurityGroupIds
	SetSerialNumber(v string) *DescribeRCInstanceAttributeResponseBody
	GetSerialNumber() *string
	SetSpotStrategy(v string) *DescribeRCInstanceAttributeResponseBody
	GetSpotStrategy() *string
	SetStatus(v string) *DescribeRCInstanceAttributeResponseBody
	GetStatus() *string
	SetStoppedMode(v string) *DescribeRCInstanceAttributeResponseBody
	GetStoppedMode() *string
	SetSystemDisk(v *DescribeRCInstanceAttributeResponseBodySystemDisk) *DescribeRCInstanceAttributeResponseBody
	GetSystemDisk() *DescribeRCInstanceAttributeResponseBodySystemDisk
	SetTags(v *DescribeRCInstanceAttributeResponseBodyTags) *DescribeRCInstanceAttributeResponseBody
	GetTags() *DescribeRCInstanceAttributeResponseBodyTags
	SetUserData(v string) *DescribeRCInstanceAttributeResponseBody
	GetUserData() *string
	SetVlanId(v string) *DescribeRCInstanceAttributeResponseBody
	GetVlanId() *string
	SetVpcAttributes(v *DescribeRCInstanceAttributeResponseBodyVpcAttributes) *DescribeRCInstanceAttributeResponseBody
	GetVpcAttributes() *DescribeRCInstanceAttributeResponseBodyVpcAttributes
	SetZoneId(v string) *DescribeRCInstanceAttributeResponseBody
	GetZoneId() *string
}

type DescribeRCInstanceAttributeResponseBody struct {
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The ID of the cluster to which the instance belongs.
	//
	// >  This parameter will be deprecated. We recommend that you use other parameters to ensure compatibility.
	//
	// example:
	//
	// None
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of CPU cores.
	//
	// example:
	//
	// 4
	Cpu        *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CreateMode *int32 `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	// The time when the instance was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-04-22T06:52:23Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The performance mode of the burstable instance.
	//
	// example:
	//
	// None
	CreditSpecification *string `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	// The details of the data disk.
	DataDisks *DescribeRCInstanceAttributeResponseBodyDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Struct"`
	DbType    *string                                           `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The attributes of the dedicated hosts.
	//
	// if can be null:
	// true
	DedicatedHostAttribute *DescribeRCInstanceAttributeResponseBodyDedicatedHostAttribute `json:"DedicatedHostAttribute,omitempty" xml:"DedicatedHostAttribute,omitempty" type:"Struct"`
	DeletionProtection     *bool                                                          `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The ID of the deployment set.
	//
	// example:
	//
	// ds-uf6c8qerk019bj1l****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The instance description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The Elastic Compute Service (ECS) instance family.
	//
	// example:
	//
	// ecs.g6.2xlarge
	EcsInstanceType *string `json:"EcsInstanceType,omitempty" xml:"EcsInstanceType,omitempty"`
	// The elastic IP address (EIP) associated with the instance.
	EipAddress *DescribeRCInstanceAttributeResponseBodyEipAddress `json:"EipAddress,omitempty" xml:"EipAddress,omitempty" type:"Struct"`
	// Indicates whether the Jumbo Frame feature is enabled for the instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	EnableJumboFrame *bool `json:"EnableJumboFrame,omitempty" xml:"EnableJumboFrame,omitempty"`
	// The expiration time. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-08-10T00:00:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Gpu         *int32  `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	GpuTypes    *string `json:"GpuTypes,omitempty" xml:"GpuTypes,omitempty"`
	// The instance hostname.
	//
	// example:
	//
	// iZ2zej1n3cin51rlmby****
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The storage type of the host. Valid values:
	//
	// 	- **dhg_cloud_ssd**: ESSD
	//
	// 	- **dhg_local_ssd**: local SSD
	//
	// example:
	//
	// dhg_cloud_ssd
	HostType *string `json:"HostType,omitempty" xml:"HostType,omitempty"`
	// The image ID of the instance.
	//
	// example:
	//
	// m-2oqiu973jwcxe****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The private IP addresses of the instance in the classic network.
	InnerIpAddress     *DescribeRCInstanceAttributeResponseBodyInnerIpAddress `json:"InnerIpAddress,omitempty" xml:"InnerIpAddress,omitempty" type:"Struct"`
	InstanceChargeType *string                                                `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rc-dh2jf9n6j4s14926****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The network type. Valid values:
	//
	// 	- **classic**
	//
	// 	- **vpc**
	//
	// example:
	//
	// vpc
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The instance type of the instance.
	//
	// example:
	//
	// mysql.x4.xlarge.6cm
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The billing method for network usage. Valid values:
	//
	// 	- **PayByBandwidth**: pay-by-bandwidth
	//
	// 	- **PayByTraffic**: pay-by-data-transfer
	//
	// >  If the **pay-by-traffic*	- billing method is used for network usage, the maximum inbound and outbound bandwidths are used as the upper limits of bandwidths instead of guaranteed performance specifications. In scenarios in which demands exceed resource supplies, the maximum bandwidths may not be reached. If you want guaranteed bandwidths for your instance, use the **pay-by-bandwidth*	- billing method for network usage.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum inbound bandwidth from the Internet. Unit: Mbit/s.
	//
	// example:
	//
	// 1
	InternetMaxBandwidthIn *int32 `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	// The maximum outbound bandwidth to the Internet. Unit: Mbit/s.
	//
	// example:
	//
	// 5
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// Indicates whether the instance is I/O optimized.
	//
	// 	- **optimized**: The instance is I/O optimized.
	//
	// 	- **none**: The instance is not I/O optimized.
	//
	// example:
	//
	// optimized
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// test_01
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The memory capacity of the instance. Unit: MiB.
	//
	// example:
	//
	// 8192
	Memory   *int32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The reasons why the instance is locked.
	OperationLocks *DescribeRCInstanceAttributeResponseBodyOperationLocks `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Struct"`
	// The public IP address of the instance.
	PublicIpAddress *DescribeRCInstanceAttributeResponseBodyPublicIpAddress `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EA2D4F34-01A7-46EB-A339-D80882135206
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The security groups.
	SecurityGroupIds *DescribeRCInstanceAttributeResponseBodySecurityGroupIds `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Struct"`
	// The serial number of the instance.
	//
	// example:
	//
	// b076f6ff-46d1-4234-a608-4e951ed6****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The instance status. Valid values:
	//
	// 	- **Pending**
	//
	// 	- **Running**
	//
	// 	- **Starting**
	//
	// 	- **Stopping**
	//
	// 	- **Stopped**
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the billing of the instance continues after the instance is stopped. Valid values:
	//
	// 	- **KeepCharging**: The billing of the instance continues after the instance is stopped, and resources are retained for the instance.
	//
	// 	- **StopCharging**: The billing of the instance stops after the instance is stopped. After the instance is stopped, resources such as CPU cores, memory resources, and public IP address are released. The instance may be unable to restart if some required resources are out of stock in the current region.
	//
	// 	- **Not-applicable**: The No Fees for Stopped Instances feature is not supported for the instance.
	//
	// example:
	//
	// Not-applicable
	StoppedMode *string                                            `json:"StoppedMode,omitempty" xml:"StoppedMode,omitempty"`
	SystemDisk  *DescribeRCInstanceAttributeResponseBodySystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	Tags        *DescribeRCInstanceAttributeResponseBodyTags       `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	UserData    *string                                            `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The virtual LAN (VLAN) ID of the instance.
	//
	// >  This parameter will be deprecated. We recommend that you use other parameters to ensure compatibility.
	//
	// example:
	//
	// None
	VlanId *string `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
	// The virtual private cloud (VPC) attributes of the instance.
	//
	// if can be null:
	// true
	VpcAttributes *DescribeRCInstanceAttributeResponseBodyVpcAttributes `json:"VpcAttributes,omitempty" xml:"VpcAttributes,omitempty" type:"Struct"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRCInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceAttributeResponseBody) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *DescribeRCInstanceAttributeResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeRCInstanceAttributeResponseBody) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeRCInstanceAttributeResponseBody) GetCreateMode() *int32 {
	return s.CreateMode
}

func (s *DescribeRCInstanceAttributeResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeRCInstanceAttributeResponseBody) GetCreditSpecification() *string {
	return s.CreditSpecification
}

func (s *DescribeRCInstanceAttributeResponseBody) GetDataDisks() *DescribeRCInstanceAttributeResponseBodyDataDisks {
	return s.DataDisks
}

func (s *DescribeRCInstanceAttributeResponseBody) GetDbType() *string {
	return s.DbType
}

func (s *DescribeRCInstanceAttributeResponseBody) GetDedicatedHostAttribute() *DescribeRCInstanceAttributeResponseBodyDedicatedHostAttribute {
	return s.DedicatedHostAttribute
}

func (s *DescribeRCInstanceAttributeResponseBody) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *DescribeRCInstanceAttributeResponseBody) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *DescribeRCInstanceAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeRCInstanceAttributeResponseBody) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeRCInstanceAttributeResponseBody) GetEcsInstanceType() *string {
	return s.EcsInstanceType
}

func (s *DescribeRCInstanceAttributeResponseBody) GetEipAddress() *DescribeRCInstanceAttributeResponseBodyEipAddress {
	return s.EipAddress
}

func (s *DescribeRCInstanceAttributeResponseBody) GetEnableJumboFrame() *bool {
	return s.EnableJumboFrame
}

func (s *DescribeRCInstanceAttributeResponseBody) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeRCInstanceAttributeResponseBody) GetGpu() *int32 {
	return s.Gpu
}

func (s *DescribeRCInstanceAttributeResponseBody) GetGpuTypes() *string {
	return s.GpuTypes
}

func (s *DescribeRCInstanceAttributeResponseBody) GetHostName() *string {
	return s.HostName
}

func (s *DescribeRCInstanceAttributeResponseBody) GetHostType() *string {
	return s.HostType
}

func (s *DescribeRCInstanceAttributeResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeRCInstanceAttributeResponseBody) GetInnerIpAddress() *DescribeRCInstanceAttributeResponseBodyInnerIpAddress {
	return s.InnerIpAddress
}

func (s *DescribeRCInstanceAttributeResponseBody) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeRCInstanceAttributeResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRCInstanceAttributeResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeRCInstanceAttributeResponseBody) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *DescribeRCInstanceAttributeResponseBody) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeRCInstanceAttributeResponseBody) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeRCInstanceAttributeResponseBody) GetInternetMaxBandwidthIn() *int32 {
	return s.InternetMaxBandwidthIn
}

func (s *DescribeRCInstanceAttributeResponseBody) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *DescribeRCInstanceAttributeResponseBody) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *DescribeRCInstanceAttributeResponseBody) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *DescribeRCInstanceAttributeResponseBody) GetMemory() *int32 {
	return s.Memory
}

func (s *DescribeRCInstanceAttributeResponseBody) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeRCInstanceAttributeResponseBody) GetOperationLocks() *DescribeRCInstanceAttributeResponseBodyOperationLocks {
	return s.OperationLocks
}

func (s *DescribeRCInstanceAttributeResponseBody) GetPublicIpAddress() *DescribeRCInstanceAttributeResponseBodyPublicIpAddress {
	return s.PublicIpAddress
}

func (s *DescribeRCInstanceAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCInstanceAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRCInstanceAttributeResponseBody) GetSecurityGroupIds() *DescribeRCInstanceAttributeResponseBodySecurityGroupIds {
	return s.SecurityGroupIds
}

func (s *DescribeRCInstanceAttributeResponseBody) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeRCInstanceAttributeResponseBody) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *DescribeRCInstanceAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeRCInstanceAttributeResponseBody) GetStoppedMode() *string {
	return s.StoppedMode
}

func (s *DescribeRCInstanceAttributeResponseBody) GetSystemDisk() *DescribeRCInstanceAttributeResponseBodySystemDisk {
	return s.SystemDisk
}

func (s *DescribeRCInstanceAttributeResponseBody) GetTags() *DescribeRCInstanceAttributeResponseBodyTags {
	return s.Tags
}

func (s *DescribeRCInstanceAttributeResponseBody) GetUserData() *string {
	return s.UserData
}

func (s *DescribeRCInstanceAttributeResponseBody) GetVlanId() *string {
	return s.VlanId
}

func (s *DescribeRCInstanceAttributeResponseBody) GetVpcAttributes() *DescribeRCInstanceAttributeResponseBodyVpcAttributes {
	return s.VpcAttributes
}

func (s *DescribeRCInstanceAttributeResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRCInstanceAttributeResponseBody) SetAutoRenew(v bool) *DescribeRCInstanceAttributeResponseBody {
	s.AutoRenew = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetClusterId(v string) *DescribeRCInstanceAttributeResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetCpu(v int32) *DescribeRCInstanceAttributeResponseBody {
	s.Cpu = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetCreateMode(v int32) *DescribeRCInstanceAttributeResponseBody {
	s.CreateMode = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetCreationTime(v string) *DescribeRCInstanceAttributeResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetCreditSpecification(v string) *DescribeRCInstanceAttributeResponseBody {
	s.CreditSpecification = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetDataDisks(v *DescribeRCInstanceAttributeResponseBodyDataDisks) *DescribeRCInstanceAttributeResponseBody {
	s.DataDisks = v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetDbType(v string) *DescribeRCInstanceAttributeResponseBody {
	s.DbType = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetDedicatedHostAttribute(v *DescribeRCInstanceAttributeResponseBodyDedicatedHostAttribute) *DescribeRCInstanceAttributeResponseBody {
	s.DedicatedHostAttribute = v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetDeletionProtection(v bool) *DescribeRCInstanceAttributeResponseBody {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetDeploymentSetId(v string) *DescribeRCInstanceAttributeResponseBody {
	s.DeploymentSetId = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetDescription(v string) *DescribeRCInstanceAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetDiskType(v string) *DescribeRCInstanceAttributeResponseBody {
	s.DiskType = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetEcsInstanceType(v string) *DescribeRCInstanceAttributeResponseBody {
	s.EcsInstanceType = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetEipAddress(v *DescribeRCInstanceAttributeResponseBodyEipAddress) *DescribeRCInstanceAttributeResponseBody {
	s.EipAddress = v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetEnableJumboFrame(v bool) *DescribeRCInstanceAttributeResponseBody {
	s.EnableJumboFrame = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetExpiredTime(v string) *DescribeRCInstanceAttributeResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetGpu(v int32) *DescribeRCInstanceAttributeResponseBody {
	s.Gpu = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetGpuTypes(v string) *DescribeRCInstanceAttributeResponseBody {
	s.GpuTypes = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetHostName(v string) *DescribeRCInstanceAttributeResponseBody {
	s.HostName = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetHostType(v string) *DescribeRCInstanceAttributeResponseBody {
	s.HostType = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetImageId(v string) *DescribeRCInstanceAttributeResponseBody {
	s.ImageId = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetInnerIpAddress(v *DescribeRCInstanceAttributeResponseBodyInnerIpAddress) *DescribeRCInstanceAttributeResponseBody {
	s.InnerIpAddress = v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetInstanceChargeType(v string) *DescribeRCInstanceAttributeResponseBody {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetInstanceId(v string) *DescribeRCInstanceAttributeResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetInstanceName(v string) *DescribeRCInstanceAttributeResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetInstanceNetworkType(v string) *DescribeRCInstanceAttributeResponseBody {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetInstanceType(v string) *DescribeRCInstanceAttributeResponseBody {
	s.InstanceType = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetInternetChargeType(v string) *DescribeRCInstanceAttributeResponseBody {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetInternetMaxBandwidthIn(v int32) *DescribeRCInstanceAttributeResponseBody {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetInternetMaxBandwidthOut(v int32) *DescribeRCInstanceAttributeResponseBody {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetIoOptimized(v string) *DescribeRCInstanceAttributeResponseBody {
	s.IoOptimized = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetKeyPairName(v string) *DescribeRCInstanceAttributeResponseBody {
	s.KeyPairName = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetMemory(v int32) *DescribeRCInstanceAttributeResponseBody {
	s.Memory = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetNodeType(v string) *DescribeRCInstanceAttributeResponseBody {
	s.NodeType = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetOperationLocks(v *DescribeRCInstanceAttributeResponseBodyOperationLocks) *DescribeRCInstanceAttributeResponseBody {
	s.OperationLocks = v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetPublicIpAddress(v *DescribeRCInstanceAttributeResponseBodyPublicIpAddress) *DescribeRCInstanceAttributeResponseBody {
	s.PublicIpAddress = v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetRegionId(v string) *DescribeRCInstanceAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetRequestId(v string) *DescribeRCInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetResourceGroupId(v string) *DescribeRCInstanceAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetSecurityGroupIds(v *DescribeRCInstanceAttributeResponseBodySecurityGroupIds) *DescribeRCInstanceAttributeResponseBody {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetSerialNumber(v string) *DescribeRCInstanceAttributeResponseBody {
	s.SerialNumber = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetSpotStrategy(v string) *DescribeRCInstanceAttributeResponseBody {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetStatus(v string) *DescribeRCInstanceAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetStoppedMode(v string) *DescribeRCInstanceAttributeResponseBody {
	s.StoppedMode = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetSystemDisk(v *DescribeRCInstanceAttributeResponseBodySystemDisk) *DescribeRCInstanceAttributeResponseBody {
	s.SystemDisk = v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetTags(v *DescribeRCInstanceAttributeResponseBodyTags) *DescribeRCInstanceAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetUserData(v string) *DescribeRCInstanceAttributeResponseBody {
	s.UserData = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetVlanId(v string) *DescribeRCInstanceAttributeResponseBody {
	s.VlanId = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetVpcAttributes(v *DescribeRCInstanceAttributeResponseBodyVpcAttributes) *DescribeRCInstanceAttributeResponseBody {
	s.VpcAttributes = v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) SetZoneId(v string) *DescribeRCInstanceAttributeResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBody) Validate() error {
	if s.DataDisks != nil {
		if err := s.DataDisks.Validate(); err != nil {
			return err
		}
	}
	if s.DedicatedHostAttribute != nil {
		if err := s.DedicatedHostAttribute.Validate(); err != nil {
			return err
		}
	}
	if s.EipAddress != nil {
		if err := s.EipAddress.Validate(); err != nil {
			return err
		}
	}
	if s.InnerIpAddress != nil {
		if err := s.InnerIpAddress.Validate(); err != nil {
			return err
		}
	}
	if s.OperationLocks != nil {
		if err := s.OperationLocks.Validate(); err != nil {
			return err
		}
	}
	if s.PublicIpAddress != nil {
		if err := s.PublicIpAddress.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityGroupIds != nil {
		if err := s.SecurityGroupIds.Validate(); err != nil {
			return err
		}
	}
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	if s.VpcAttributes != nil {
		if err := s.VpcAttributes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRCInstanceAttributeResponseBodyDataDisks struct {
	DataDisk []*DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
}

func (s DescribeRCInstanceAttributeResponseBodyDataDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceAttributeResponseBodyDataDisks) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceAttributeResponseBodyDataDisks) GetDataDisk() []*DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk {
	return s.DataDisk
}

func (s *DescribeRCInstanceAttributeResponseBodyDataDisks) SetDataDisk(v []*DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk) *DescribeRCInstanceAttributeResponseBodyDataDisks {
	s.DataDisk = v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyDataDisks) Validate() error {
	if s.DataDisk != nil {
		for _, item := range s.DataDisk {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk struct {
	// The category of the data disk.
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Indicates whether the data disk is released when the instance is released. Valid values:
	//
	// 	- **true**: The data disk is released when the instance is released.
	//
	// 	- **false**: The data disk is reserved when the instance is released.
	//
	// example:
	//
	// true
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// example:
	//
	// /dev/xvdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// Indicates whether the data disk is encrypted. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The performance level of data disk. This parameter is available when the data disk is an Enterprise SSD (ESSD).
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of the data disk. Unit: GiB.
	//
	// example:
	//
	// 40
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// rcds-bp18um4r4f2fve24**
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk) GetCategory() *string {
	return s.Category
}

func (s *DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk) GetDevice() *string {
	return s.Device
}

func (s *DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk) GetEncrypted() *string {
	return s.Encrypted
}

func (s *DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk) GetSize() *int64 {
	return s.Size
}

func (s *DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk) SetCategory(v string) *DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk {
	s.Category = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk) SetDeleteWithInstance(v bool) *DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk) SetDevice(v string) *DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk {
	s.Device = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk) SetEncrypted(v string) *DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk {
	s.Encrypted = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk) SetPerformanceLevel(v string) *DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk) SetSize(v int64) *DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk {
	s.Size = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk) SetSnapshotId(v string) *DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk {
	s.SnapshotId = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyDataDisksDataDisk) Validate() error {
	return dara.Validate(s)
}

type DescribeRCInstanceAttributeResponseBodyDedicatedHostAttribute struct {
	// The ID of the dedicated host.
	//
	// example:
	//
	// None
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// The name of the dedicated host.
	//
	// example:
	//
	// None
	DedicatedHostName *string `json:"DedicatedHostName,omitempty" xml:"DedicatedHostName,omitempty"`
}

func (s DescribeRCInstanceAttributeResponseBodyDedicatedHostAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceAttributeResponseBodyDedicatedHostAttribute) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceAttributeResponseBodyDedicatedHostAttribute) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *DescribeRCInstanceAttributeResponseBodyDedicatedHostAttribute) GetDedicatedHostName() *string {
	return s.DedicatedHostName
}

func (s *DescribeRCInstanceAttributeResponseBodyDedicatedHostAttribute) SetDedicatedHostId(v string) *DescribeRCInstanceAttributeResponseBodyDedicatedHostAttribute {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyDedicatedHostAttribute) SetDedicatedHostName(v string) *DescribeRCInstanceAttributeResponseBodyDedicatedHostAttribute {
	s.DedicatedHostName = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyDedicatedHostAttribute) Validate() error {
	return dara.Validate(s)
}

type DescribeRCInstanceAttributeResponseBodyEipAddress struct {
	// The EIP ID.
	//
	// example:
	//
	// eip-bp14k3rz6cbg6zxbe****
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The maximum Internet bandwidth of the EIP. Unit: Mbit/s.
	//
	// example:
	//
	// 5
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The billing method of the Internet-facing instance. Valid values:
	//
	// 	- **paybytraffic:*	- pay-by-data-transfer
	//
	// 	- **paybybandwidth**: pay-by-bandwidth
	//
	// >  If the **pay-by-traffic*	- billing method is used for network usage, the maximum inbound and outbound bandwidths are used as the upper limits of bandwidths instead of guaranteed performance specifications. In scenarios in which demands exceed resource supplies, the maximum bandwidths may not be reached. If you want guaranteed bandwidths for your instance, use the **pay-by-bandwidth*	- billing method for network usage.
	//
	// example:
	//
	// paybytraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The EIP.
	//
	// example:
	//
	// 8.147.XXX.XXX
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
}

func (s DescribeRCInstanceAttributeResponseBodyEipAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceAttributeResponseBodyEipAddress) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceAttributeResponseBodyEipAddress) GetAllocationId() *string {
	return s.AllocationId
}

func (s *DescribeRCInstanceAttributeResponseBodyEipAddress) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeRCInstanceAttributeResponseBodyEipAddress) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeRCInstanceAttributeResponseBodyEipAddress) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeRCInstanceAttributeResponseBodyEipAddress) SetAllocationId(v string) *DescribeRCInstanceAttributeResponseBodyEipAddress {
	s.AllocationId = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyEipAddress) SetBandwidth(v int32) *DescribeRCInstanceAttributeResponseBodyEipAddress {
	s.Bandwidth = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyEipAddress) SetInternetChargeType(v string) *DescribeRCInstanceAttributeResponseBodyEipAddress {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyEipAddress) SetIpAddress(v string) *DescribeRCInstanceAttributeResponseBodyEipAddress {
	s.IpAddress = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyEipAddress) Validate() error {
	return dara.Validate(s)
}

type DescribeRCInstanceAttributeResponseBodyInnerIpAddress struct {
	IpAddress []*string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty" type:"Repeated"`
}

func (s DescribeRCInstanceAttributeResponseBodyInnerIpAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceAttributeResponseBodyInnerIpAddress) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceAttributeResponseBodyInnerIpAddress) GetIpAddress() []*string {
	return s.IpAddress
}

func (s *DescribeRCInstanceAttributeResponseBodyInnerIpAddress) SetIpAddress(v []*string) *DescribeRCInstanceAttributeResponseBodyInnerIpAddress {
	s.IpAddress = v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyInnerIpAddress) Validate() error {
	return dara.Validate(s)
}

type DescribeRCInstanceAttributeResponseBodyOperationLocks struct {
	LockReason []*DescribeRCInstanceAttributeResponseBodyOperationLocksLockReason `json:"LockReason,omitempty" xml:"LockReason,omitempty" type:"Repeated"`
}

func (s DescribeRCInstanceAttributeResponseBodyOperationLocks) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceAttributeResponseBodyOperationLocks) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceAttributeResponseBodyOperationLocks) GetLockReason() []*DescribeRCInstanceAttributeResponseBodyOperationLocksLockReason {
	return s.LockReason
}

func (s *DescribeRCInstanceAttributeResponseBodyOperationLocks) SetLockReason(v []*DescribeRCInstanceAttributeResponseBodyOperationLocksLockReason) *DescribeRCInstanceAttributeResponseBodyOperationLocks {
	s.LockReason = v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyOperationLocks) Validate() error {
	if s.LockReason != nil {
		for _, item := range s.LockReason {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCInstanceAttributeResponseBodyOperationLocksLockReason struct {
	// The reason why the instance is locked. Valid values:
	//
	// 	- **financial**: The instance is locked due to overdue payments.
	//
	// 	- **security**: The instance is locked for security purposes.
	//
	// 	- **recycling**: The instance is locked because the instance is a preemptible instance and pending to be released.
	//
	// 	- **dedicatedhostfinancial**: The instance is locked due to overdue payments for the dedicated host.
	//
	// 	- **refunded**: The instance is locked because a refund was made for the instance.
	//
	// example:
	//
	// None
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
}

func (s DescribeRCInstanceAttributeResponseBodyOperationLocksLockReason) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceAttributeResponseBodyOperationLocksLockReason) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceAttributeResponseBodyOperationLocksLockReason) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeRCInstanceAttributeResponseBodyOperationLocksLockReason) SetLockReason(v string) *DescribeRCInstanceAttributeResponseBodyOperationLocksLockReason {
	s.LockReason = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyOperationLocksLockReason) Validate() error {
	return dara.Validate(s)
}

type DescribeRCInstanceAttributeResponseBodyPublicIpAddress struct {
	IpAddress []*string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty" type:"Repeated"`
}

func (s DescribeRCInstanceAttributeResponseBodyPublicIpAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceAttributeResponseBodyPublicIpAddress) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceAttributeResponseBodyPublicIpAddress) GetIpAddress() []*string {
	return s.IpAddress
}

func (s *DescribeRCInstanceAttributeResponseBodyPublicIpAddress) SetIpAddress(v []*string) *DescribeRCInstanceAttributeResponseBodyPublicIpAddress {
	s.IpAddress = v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyPublicIpAddress) Validate() error {
	return dara.Validate(s)
}

type DescribeRCInstanceAttributeResponseBodySecurityGroupIds struct {
	SecurityGroupId []*string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Repeated"`
}

func (s DescribeRCInstanceAttributeResponseBodySecurityGroupIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceAttributeResponseBodySecurityGroupIds) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceAttributeResponseBodySecurityGroupIds) GetSecurityGroupId() []*string {
	return s.SecurityGroupId
}

func (s *DescribeRCInstanceAttributeResponseBodySecurityGroupIds) SetSecurityGroupId(v []*string) *DescribeRCInstanceAttributeResponseBodySecurityGroupIds {
	s.SecurityGroupId = v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodySecurityGroupIds) Validate() error {
	return dara.Validate(s)
}

type DescribeRCInstanceAttributeResponseBodySystemDisk struct {
	DeleteWithInstance         *bool   `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	Encrypted                  *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	SystemDiskCategory         *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	SystemDiskPerformanceLevel *string `json:"SystemDiskPerformanceLevel,omitempty" xml:"SystemDiskPerformanceLevel,omitempty"`
	SystemDiskSize             *int64  `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribeRCInstanceAttributeResponseBodySystemDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceAttributeResponseBodySystemDisk) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceAttributeResponseBodySystemDisk) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *DescribeRCInstanceAttributeResponseBodySystemDisk) GetEncrypted() *string {
	return s.Encrypted
}

func (s *DescribeRCInstanceAttributeResponseBodySystemDisk) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *DescribeRCInstanceAttributeResponseBodySystemDisk) GetSystemDiskPerformanceLevel() *string {
	return s.SystemDiskPerformanceLevel
}

func (s *DescribeRCInstanceAttributeResponseBodySystemDisk) GetSystemDiskSize() *int64 {
	return s.SystemDiskSize
}

func (s *DescribeRCInstanceAttributeResponseBodySystemDisk) SetDeleteWithInstance(v bool) *DescribeRCInstanceAttributeResponseBodySystemDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodySystemDisk) SetEncrypted(v string) *DescribeRCInstanceAttributeResponseBodySystemDisk {
	s.Encrypted = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodySystemDisk) SetSystemDiskCategory(v string) *DescribeRCInstanceAttributeResponseBodySystemDisk {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodySystemDisk) SetSystemDiskPerformanceLevel(v string) *DescribeRCInstanceAttributeResponseBodySystemDisk {
	s.SystemDiskPerformanceLevel = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodySystemDisk) SetSystemDiskSize(v int64) *DescribeRCInstanceAttributeResponseBodySystemDisk {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodySystemDisk) Validate() error {
	return dara.Validate(s)
}

type DescribeRCInstanceAttributeResponseBodyTags struct {
	Tag []*DescribeRCInstanceAttributeResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeRCInstanceAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceAttributeResponseBodyTags) GetTag() []*DescribeRCInstanceAttributeResponseBodyTagsTag {
	return s.Tag
}

func (s *DescribeRCInstanceAttributeResponseBodyTags) SetTag(v []*DescribeRCInstanceAttributeResponseBodyTagsTag) *DescribeRCInstanceAttributeResponseBodyTags {
	s.Tag = v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyTags) Validate() error {
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

type DescribeRCInstanceAttributeResponseBodyTagsTag struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeRCInstanceAttributeResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceAttributeResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceAttributeResponseBodyTagsTag) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeRCInstanceAttributeResponseBodyTagsTag) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeRCInstanceAttributeResponseBodyTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeRCInstanceAttributeResponseBodyTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeRCInstanceAttributeResponseBodyTagsTag) SetResourceId(v string) *DescribeRCInstanceAttributeResponseBodyTagsTag {
	s.ResourceId = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyTagsTag) SetResourceType(v string) *DescribeRCInstanceAttributeResponseBodyTagsTag {
	s.ResourceType = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyTagsTag) SetTagKey(v string) *DescribeRCInstanceAttributeResponseBodyTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyTagsTag) SetTagValue(v string) *DescribeRCInstanceAttributeResponseBodyTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyTagsTag) Validate() error {
	return dara.Validate(s)
}

type DescribeRCInstanceAttributeResponseBodyVpcAttributes struct {
	// The network address translation (NAT) IP address of the instance. The NAT IP address is used by instances in different VPCs for communication.
	//
	// example:
	//
	// None
	NatIpAddress *string `json:"NatIpAddress,omitempty" xml:"NatIpAddress,omitempty"`
	// The private IP addresses of the instance.
	PrivateIpAddress *DescribeRCInstanceAttributeResponseBodyVpcAttributesPrivateIpAddress `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty" type:"Struct"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1nt15muovrc5qdj****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-2zeu747v4765aw2id****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeRCInstanceAttributeResponseBodyVpcAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceAttributeResponseBodyVpcAttributes) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceAttributeResponseBodyVpcAttributes) GetNatIpAddress() *string {
	return s.NatIpAddress
}

func (s *DescribeRCInstanceAttributeResponseBodyVpcAttributes) GetPrivateIpAddress() *DescribeRCInstanceAttributeResponseBodyVpcAttributesPrivateIpAddress {
	return s.PrivateIpAddress
}

func (s *DescribeRCInstanceAttributeResponseBodyVpcAttributes) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeRCInstanceAttributeResponseBodyVpcAttributes) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeRCInstanceAttributeResponseBodyVpcAttributes) SetNatIpAddress(v string) *DescribeRCInstanceAttributeResponseBodyVpcAttributes {
	s.NatIpAddress = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyVpcAttributes) SetPrivateIpAddress(v *DescribeRCInstanceAttributeResponseBodyVpcAttributesPrivateIpAddress) *DescribeRCInstanceAttributeResponseBodyVpcAttributes {
	s.PrivateIpAddress = v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyVpcAttributes) SetVSwitchId(v string) *DescribeRCInstanceAttributeResponseBodyVpcAttributes {
	s.VSwitchId = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyVpcAttributes) SetVpcId(v string) *DescribeRCInstanceAttributeResponseBodyVpcAttributes {
	s.VpcId = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyVpcAttributes) Validate() error {
	if s.PrivateIpAddress != nil {
		if err := s.PrivateIpAddress.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRCInstanceAttributeResponseBodyVpcAttributesPrivateIpAddress struct {
	IpAddress []*string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty" type:"Repeated"`
}

func (s DescribeRCInstanceAttributeResponseBodyVpcAttributesPrivateIpAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceAttributeResponseBodyVpcAttributesPrivateIpAddress) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceAttributeResponseBodyVpcAttributesPrivateIpAddress) GetIpAddress() []*string {
	return s.IpAddress
}

func (s *DescribeRCInstanceAttributeResponseBodyVpcAttributesPrivateIpAddress) SetIpAddress(v []*string) *DescribeRCInstanceAttributeResponseBodyVpcAttributesPrivateIpAddress {
	s.IpAddress = v
	return s
}

func (s *DescribeRCInstanceAttributeResponseBodyVpcAttributesPrivateIpAddress) Validate() error {
	return dara.Validate(s)
}
