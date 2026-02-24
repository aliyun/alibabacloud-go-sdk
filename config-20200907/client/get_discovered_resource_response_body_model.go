// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDiscoveredResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiscoveredResourceDetail(v *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) *GetDiscoveredResourceResponseBody
	GetDiscoveredResourceDetail() *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail
	SetRequestId(v string) *GetDiscoveredResourceResponseBody
	GetRequestId() *string
}

type GetDiscoveredResourceResponseBody struct {
	// The details of the resource.
	DiscoveredResourceDetail *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail `json:"DiscoveredResourceDetail,omitempty" xml:"DiscoveredResourceDetail,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E4D71ACE-6B0A-46E0-8352-56952378CC7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDiscoveredResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDiscoveredResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceResponseBody) GetDiscoveredResourceDetail() *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	return s.DiscoveredResourceDetail
}

func (s *GetDiscoveredResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDiscoveredResourceResponseBody) SetDiscoveredResourceDetail(v *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) *GetDiscoveredResourceResponseBody {
	s.DiscoveredResourceDetail = v
	return s
}

func (s *GetDiscoveredResourceResponseBody) SetRequestId(v string) *GetDiscoveredResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDiscoveredResourceResponseBody) Validate() error {
	if s.DiscoveredResourceDetail != nil {
		if err := s.DiscoveredResourceDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDiscoveredResourceResponseBodyDiscoveredResourceDetail struct {
	// The ID of the Alibaba Cloud account that owns the resource.
	//
	// example:
	//
	// 120886317861****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The zone where the resource resides.
	//
	// example:
	//
	// cn-hangzhou-h
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	// The compliance type.
	//
	// example:
	//
	// COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The complete configuration information of the resource.
	//
	// example:
	//
	// {"ResourceGroupId":"","Memory":4096,"InstanceChargeType":"PostPaid","Cpu":2,"OSName":"Windows Server 2022 Datacenter 64-bit (Simplified Chinese)","InstanceNetworkType":"vpc","InnerIpAddress":{"IpAddress":[]},"ExpiredTime":"2099-12-31T15:59Z","ImageId":"win2022_21H2_x64_dtc_zh-cn_40G_alibase_20240110.vhd","EipAddress":{"AllocationId":"","IpAddress":"","InternetChargeType":""},"ImageOptions":{},"VlanId":"","HostName":"iZl4i0brknq****","Status":"Stopped","HibernationOptions":{"Configured":false},"MetadataOptions":{"HttpTokens":"","HttpEndpoint":""},"InstanceId":"i-bp12g4xbl4i0brkn****","StoppedMode":"KeepCharging","CpuOptions":{"ThreadsPerCore":2,"Numa":"ON","CoreCount":1},"StartTime":"2024-02-29T07:08Z","DeletionProtection":false,"VpcAttributes":{"PrivateIpAddress":{"IpAddress":["172.16.XX.XX"]},"VpcId":"vpc-bp1wjaw8t272wwmkg****","VSwitchId":"vsw-bp103i8xzww5132ul****","NatIpAddress":""},"SecurityGroupIds":{"SecurityGroupId":["sg-bp1h96fz9fagaegp****"]},"InternetChargeType":"PayByTraffic","InstanceName":"test123","DeploymentSetId":"","InternetMaxBandwidthOut":5,"SerialNumber":"6764f567-28fb-4a39-bfc3-48404995****","OSType":"windows","CreationTime":"2024-02-29T07:08Z","AutoReleaseTime":"","Description":"","InstanceTypeFamily":"ecs.c7","DedicatedInstanceAttribute":{"Tenancy":"","Affinity":""},"PublicIpAddress":{"IpAddress":["47.98.XX.XX"]},"GPUSpec":"","NetworkInterfaces":{"NetworkInterface":[{"Type":"Primary","PrimaryIpAddress":"172.16.XX.XX","MacAddress":"00:16:3e:0c:**:**","NetworkInterfaceId":"eni-bp19uj35v8won3x9****","PrivateIpSets":{"PrivateIpSet":[{"PrivateIpAddress":"172.16.XX.XX","Primary":true}]}}]},"SpotPriceLimit":0.0,"SaleCycle":"","DeviceAvailable":true,"InstanceType":"ecs.c7.large","OSNameEn":"Windows Server  2022 DataCenter Edition 64bit Chinese Edition","SpotStrategy":"NoSpot","IoOptimized":true,"ZoneId":"cn-hangzhou-b","ClusterId":"","EcsCapacityReservationAttr":{"CapacityReservationPreference":"","CapacityReservationId":""},"DedicatedHostAttribute":{"DedicatedHostId":"","DedicatedHostName":"","DedicatedHostClusterId":""},"GPUAmount":0,"OperationLocks":{"LockReason":[]},"InternetMaxBandwidthIn":2000,"Recyclable":false,"RegionId":"cn-hangzhou","CreditSpecification":""}
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
	// 1709190480000
	ResourceCreationTime *int64 `json:"ResourceCreationTime,omitempty" xml:"ResourceCreationTime,omitempty"`
	// The deletion status of the resource. Valid values:
	//
	// - 1: The resource is not deleted.
	//
	// - 0: The resource is deleted.
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
	// The status of the resource. The status of a resource is defined by the corresponding Alibaba Cloud service. This parameter can be empty. Examples:
	//
	// - If the resource type is ACS::ECS::Instance, the resource is stateful. In this case, the value of this parameter is Running or Stopped.
	//
	// - If the resource type is ACS::OSS::Bucket, the resource is stateless. In this case, this parameter is empty.
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
	// The tags of the resource.
	//
	// example:
	//
	// {"key":"value"}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The ID of the vSwitch to which the resource belongs. The ID is in the vsw-t4n7pokxxxxxxxxxxxxxx format. If multiple vSwitch IDs are returned, they are separated by commas (,). Example: vsw-t4n7pokxxxxxxxxxxxxxx,vsw-t4n7pokxxxxxxxxxxxxxx. If the resource does not belong to a vSwitch, an empty string "" is returned.
	//
	// This parameter is required.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// vsw-t4n7pokxxxxxxxxxxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC to which the resource belongs. The ID is in the vpc-t4nhheyvay74fp7n0hxxx format. If the resource does not belong to a VPC, an empty string "" is returned.
	//
	// This parameter is required.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// vpc-t4nhheyvay74fp7n0hxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) String() string {
	return dara.Prettify(s)
}

func (s GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetAvailabilityZone() *string {
	return s.AvailabilityZone
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetConfiguration() *string {
	return s.Configuration
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetRegion() *string {
	return s.Region
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetResourceCreationTime() *int64 {
	return s.ResourceCreationTime
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetResourceDeleted() *int32 {
	return s.ResourceDeleted
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetResourceStatus() *string {
	return s.ResourceStatus
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetTags() *string {
	return s.Tags
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetVpcId() *string {
	return s.VpcId
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetAccountId(v int64) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.AccountId = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetAvailabilityZone(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.AvailabilityZone = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetComplianceType(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ComplianceType = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetConfiguration(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.Configuration = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetRegion(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.Region = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceCreationTime(v int64) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceCreationTime = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceDeleted(v int32) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceDeleted = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceId(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceId = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceName(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceName = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceStatus(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceStatus = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceType(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceType = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetTags(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.Tags = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetVSwitchId(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.VSwitchId = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetVpcId(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.VpcId = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) Validate() error {
	return dara.Validate(s)
}
