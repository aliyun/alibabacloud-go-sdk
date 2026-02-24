// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateDiscoveredResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiscoveredResourceDetail(v *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) *GetAggregateDiscoveredResourceResponseBody
	GetDiscoveredResourceDetail() *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail
	SetRequestId(v string) *GetAggregateDiscoveredResourceResponseBody
	GetRequestId() *string
}

type GetAggregateDiscoveredResourceResponseBody struct {
	// The details of the resource.
	DiscoveredResourceDetail *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail `json:"DiscoveredResourceDetail,omitempty" xml:"DiscoveredResourceDetail,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E4D71ACE-6B0A-46E0-8352-56952378CC7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateDiscoveredResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateDiscoveredResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateDiscoveredResourceResponseBody) GetDiscoveredResourceDetail() *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	return s.DiscoveredResourceDetail
}

func (s *GetAggregateDiscoveredResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAggregateDiscoveredResourceResponseBody) SetDiscoveredResourceDetail(v *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) *GetAggregateDiscoveredResourceResponseBody {
	s.DiscoveredResourceDetail = v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBody) SetRequestId(v string) *GetAggregateDiscoveredResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBody) Validate() error {
	if s.DiscoveredResourceDetail != nil {
		if err := s.DiscoveredResourceDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail struct {
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 100931896542****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the zone where the resource resides.
	//
	// example:
	//
	// cn-hangzhou-h
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	// The compliance evaluation result. Valid values:
	//
	// - COMPLIANT: The resource is compliant.
	//
	// - NON_COMPLIANT: The resource is non-compliant.
	//
	// - NOT_APPLICABLE: The rule did not apply to your resource.
	//
	// - INSUFFICIENT_DATA: No data is available.
	//
	// example:
	//
	// COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The configuration of the resource.
	//
	// example:
	//
	// {\\"ResourceGroupId\\":\\"\\",\\"Memory\\":4096,\\"InstanceChargeType\\":\\"PostPaid\\",\\"Cpu\\":2,\\"OSName\\":\\"Windows Server  2022 数据中心版 64位中文版\\",\\"InstanceNetworkType\\":\\"vpc\\",\\"InnerIpAddress\\":{\\"IpAddress\\":[]},\\"ExpiredTime\\":\\"2099-12-31T15:59Z\\",\\"ImageId\\":\\"win2022_21H2_x64_dtc_zh-cn_40G_alibase_20240110.vhd\\",\\"EipAddress\\":{\\"AllocationId\\":\\"\\",\\"IpAddress\\":\\"\\",\\"InternetChargeType\\":\\"\\"},\\"ImageOptions\\":{},\\"VlanId\\":\\"\\",\\"HostName\\":\\"iZl4i0brknq****\\",\\"Status\\":\\"Stopped\\",\\"HibernationOptions\\":{\\"Configured\\":false},\\"MetadataOptions\\":{\\"HttpTokens\\":\\"\\",\\"HttpEndpoint\\":\\"\\"},\\"InstanceId\\":\\"i-bp12g4xbl4i0brkn****\\",\\"StoppedMode\\":\\"KeepCharging\\",\\"CpuOptions\\":{\\"ThreadsPerCore\\":2,\\"Numa\\":\\"ON\\",\\"CoreCount\\":1},\\"StartTime\\":\\"2024-02-29T07:08Z\\",\\"DeletionProtection\\":false,\\"VpcAttributes\\":{\\"PrivateIpAddress\\":{\\"IpAddress\\":[\\"172.16.XX.XX\\"]},\\"VpcId\\":\\"vpc-bp1wjaw8t272wwmkg****\\",\\"VSwitchId\\":\\"vsw-bp103i8xzww5132ul****\\",\\"NatIpAddress\\":\\"\\"},\\"SecurityGroupIds\\":{\\"SecurityGroupId\\":[\\"sg-bp1h96fz9fagaegp****\\"]},\\"InternetChargeType\\":\\"PayByTraffic\\",\\"InstanceName\\":\\"test123\\",\\"DeploymentSetId\\":\\"\\",\\"InternetMaxBandwidthOut\\":5,\\"SerialNumber\\":\\"6764f567-28fb-4a39-bfc3-48404995****\\",\\"OSType\\":\\"windows\\",\\"CreationTime\\":\\"2024-02-29T07:08Z\\",\\"AutoReleaseTime\\":\\"\\",\\"Description\\":\\"\\",\\"InstanceTypeFamily\\":\\"ecs.c7\\",\\"DedicatedInstanceAttribute\\":{\\"Tenancy\\":\\"\\",\\"Affinity\\":\\"\\"},\\"PublicIpAddress\\":{\\"IpAddress\\":[\\"47.98.XX.XX\\"]},\\"GPUSpec\\":\\"\\",\\"NetworkInterfaces\\":{\\"NetworkInterface\\":[{\\"Type\\":\\"Primary\\",\\"PrimaryIpAddress\\":\\"172.16.XX.XX\\",\\"MacAddress\\":\\"00:16:3e:0c:**:**\\",\\"NetworkInterfaceId\\":\\"eni-bp19uj35v8won3x9****\\",\\"PrivateIpSets\\":{\\"PrivateIpSet\\":[{\\"PrivateIpAddress\\":\\"172.16.XX.XX\\",\\"Primary\\":true}]}}]},\\"SpotPriceLimit\\":0.0,\\"SaleCycle\\":\\"\\",\\"DeviceAvailable\\":true,\\"InstanceType\\":\\"ecs.c7.large\\",\\"OSNameEn\\":\\"Windows Server  2022 DataCenter Edition 64bit Chinese Edition\\",\\"SpotStrategy\\":\\"NoSpot\\",\\"IoOptimized\\":true,\\"ZoneId\\":\\"cn-hangzhou-b\\",\\"ClusterId\\":\\"\\",\\"EcsCapacityReservationAttr\\":{\\"CapacityReservationPreference\\":\\"\\",\\"CapacityReservationId\\":\\"\\"},\\"DedicatedHostAttribute\\":{\\"DedicatedHostId\\":\\"\\",\\"DedicatedHostName\\":\\"\\",\\"DedicatedHostClusterId\\":\\"\\"},\\"GPUAmount\\":0,\\"OperationLocks\\":{\\"LockReason\\":[]},\\"InternetMaxBandwidthIn\\":2000,\\"Recyclable\\":false,\\"RegionId\\":\\"cn-hangzhou\\",\\"CreditSpecification\\":\\"\\"}
	Configuration *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The timestamp when the resource was created.
	//
	// example:
	//
	// 1624961112000
	ResourceCreationTime *int64 `json:"ResourceCreationTime,omitempty" xml:"ResourceCreationTime,omitempty"`
	// Indicates whether the resource was deleted. Valid values:
	//
	// - 1: The resource was not deleted.
	//
	// - 0: The resource was deleted.
	//
	// example:
	//
	// 1
	ResourceDeleted *int32 `json:"ResourceDeleted,omitempty" xml:"ResourceDeleted,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// i-bp12g4xbl4i0brkn****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource name.
	//
	// example:
	//
	// test123
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource status. The value of this parameter varies based on the resource type and may be empty. For example:
	//
	// - If the ResourceType parameter is set to ACS::ECS::Instance, the resource is an ECS instance that has a specific state. In this case, the valid values of this parameter are Running and Stopped.
	//
	// - If the ResourceType parameter is ACS::OSS::Bucket, the resource is an Object Storage Service (OSS) bucket that is not in a specific state. In this case, this parameter is left empty.
	//
	// example:
	//
	// Stopped
	ResourceStatus *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The resource tags.
	//
	// example:
	//
	// {\\"\\"hc\\"\\":[\\"\\"value2\\"\\"]}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The ID of the vSwitch to which the resource belongs, in the format of vsw-t4n7pokxxxxxxxxxxxxxx. If the resource belongs to multiple vSwitches, the IDs are separated by commas, such as vsw-t4n7pokxxxxxxxxxxxxxx, vsw-t4n7pokxxxxxxxxxxxxxx. If the resource does not belong to any vSwitch, an empty string is returned: ""
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-t4n7pokxxxxxxxxxxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC to which the resource belongs, in the format of vpc-t4nhheyvay74fp7n0hxxx. If the resource does not belong to a VPC, an empty string is returned: ""
	//
	// if can be null:
	// true
	//
	// example:
	//
	// vpc-t4nhheyvay74fp7n0hxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GoString() string {
	return s.String()
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetAvailabilityZone() *string {
	return s.AvailabilityZone
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetConfiguration() *string {
	return s.Configuration
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetRegion() *string {
	return s.Region
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetResourceCreationTime() *int64 {
	return s.ResourceCreationTime
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetResourceDeleted() *int32 {
	return s.ResourceDeleted
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetResourceStatus() *string {
	return s.ResourceStatus
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetTags() *string {
	return s.Tags
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetVpcId() *string {
	return s.VpcId
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetAccountId(v int64) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.AccountId = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetAvailabilityZone(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.AvailabilityZone = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetComplianceType(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ComplianceType = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetConfiguration(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.Configuration = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetRegion(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.Region = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceCreationTime(v int64) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceCreationTime = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceDeleted(v int32) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceDeleted = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceId(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceId = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceName(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceName = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceStatus(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceStatus = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceType(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceType = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetTags(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.Tags = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetVSwitchId(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.VSwitchId = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetVpcId(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.VpcId = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) Validate() error {
	return dara.Validate(s)
}
