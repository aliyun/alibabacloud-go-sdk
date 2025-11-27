// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeRCInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRCInstancesResponseBody
	GetPageSize() *int32
	SetRCInstances(v []*DescribeRCInstancesResponseBodyRCInstances) *DescribeRCInstancesResponseBody
	GetRCInstances() []*DescribeRCInstancesResponseBodyRCInstances
	SetRequestId(v string) *DescribeRCInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeRCInstancesResponseBody
	GetTotalCount() *int32
}

type DescribeRCInstancesResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The details of the instance.
	RCInstances []*DescribeRCInstancesResponseBodyRCInstances `json:"RCInstances,omitempty" xml:"RCInstances,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// E9DD55F4-1A5F-48CA-BA57-DFB3CA8C4C34
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRCInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRCInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRCInstancesResponseBody) GetRCInstances() []*DescribeRCInstancesResponseBodyRCInstances {
	return s.RCInstances
}

func (s *DescribeRCInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRCInstancesResponseBody) SetPageNumber(v int32) *DescribeRCInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRCInstancesResponseBody) SetPageSize(v int32) *DescribeRCInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRCInstancesResponseBody) SetRCInstances(v []*DescribeRCInstancesResponseBodyRCInstances) *DescribeRCInstancesResponseBody {
	s.RCInstances = v
	return s
}

func (s *DescribeRCInstancesResponseBody) SetRequestId(v string) *DescribeRCInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCInstancesResponseBody) SetTotalCount(v int32) *DescribeRCInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRCInstancesResponseBody) Validate() error {
	if s.RCInstances != nil {
		for _, item := range s.RCInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCInstancesResponseBodyRCInstances struct {
	// The cluster name.
	//
	// example:
	//
	// testrdscustom
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	Cpu         *int32  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CreateMode  *string `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	// The database type.
	//
	// example:
	//
	// rds_custom
	DbType          *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The instance description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The time when the task was created. The time is displayed in GMT.
	//
	// example:
	//
	// 2023-03-22 07:56:53.0
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The host IP address.
	//
	// example:
	//
	// 172.30.XXX.XXX
	HostIp *string `json:"HostIp,omitempty" xml:"HostIp,omitempty"`
	// The host name.
	//
	// example:
	//
	// i-2zeaiz4g9u23f40m****
	HostName           *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	ImageId            *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-2ze704f*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// k8s-node
	InstanceName       *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceType       *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	Memory             *int32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NodeType           *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	PublicIp           *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SpotStrategy    *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
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
	// >  If the value returned for the DescribeRCInstances operation is different from the value that is returned for the **DescribeRCInstanceAttribute*	- operation, the value returned for the **DescribeRCInstanceAttribute*	- operation shall prevail.
	//
	// example:
	//
	// Running
	Status       *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	TagResources []*DescribeRCInstancesResponseBodyRCInstancesTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
	Tags         []*DescribeRCInstancesResponseBodyRCInstancesTags         `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// if can be null:
	// true
	VpcAttributes *DescribeRCInstancesResponseBodyRCInstancesVpcAttributes `json:"VpcAttributes,omitempty" xml:"VpcAttributes,omitempty" type:"Struct"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-uf6f7l4fg90****
	VpcId  *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRCInstancesResponseBodyRCInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstancesResponseBodyRCInstances) GoString() string {
	return s.String()
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetCreateMode() *string {
	return s.CreateMode
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetDbType() *string {
	return s.DbType
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetDescription() *string {
	return s.Description
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetHostIp() *string {
	return s.HostIp
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetHostName() *string {
	return s.HostName
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetMemory() *int32 {
	return s.Memory
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetPublicIp() *string {
	return s.PublicIp
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetStatus() *string {
	return s.Status
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetTagResources() []*DescribeRCInstancesResponseBodyRCInstancesTagResources {
	return s.TagResources
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetTags() []*DescribeRCInstancesResponseBodyRCInstancesTags {
	return s.Tags
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetVpcAttributes() *DescribeRCInstancesResponseBodyRCInstancesVpcAttributes {
	return s.VpcAttributes
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeRCInstancesResponseBodyRCInstances) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetClusterName(v string) *DescribeRCInstancesResponseBodyRCInstances {
	s.ClusterName = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetCpu(v int32) *DescribeRCInstancesResponseBodyRCInstances {
	s.Cpu = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetCreateMode(v string) *DescribeRCInstancesResponseBodyRCInstances {
	s.CreateMode = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetDbType(v string) *DescribeRCInstancesResponseBodyRCInstances {
	s.DbType = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetDeploymentSetId(v string) *DescribeRCInstancesResponseBodyRCInstances {
	s.DeploymentSetId = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetDescription(v string) *DescribeRCInstancesResponseBodyRCInstances {
	s.Description = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetExpiredTime(v string) *DescribeRCInstancesResponseBodyRCInstances {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetGmtCreated(v string) *DescribeRCInstancesResponseBodyRCInstances {
	s.GmtCreated = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetHostIp(v string) *DescribeRCInstancesResponseBodyRCInstances {
	s.HostIp = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetHostName(v string) *DescribeRCInstancesResponseBodyRCInstances {
	s.HostName = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetImageId(v string) *DescribeRCInstancesResponseBodyRCInstances {
	s.ImageId = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetInstanceChargeType(v string) *DescribeRCInstancesResponseBodyRCInstances {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetInstanceId(v string) *DescribeRCInstancesResponseBodyRCInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetInstanceName(v string) *DescribeRCInstancesResponseBodyRCInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetInstanceType(v string) *DescribeRCInstancesResponseBodyRCInstances {
	s.InstanceType = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetInstanceTypeFamily(v string) *DescribeRCInstancesResponseBodyRCInstances {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetMemory(v int32) *DescribeRCInstancesResponseBodyRCInstances {
	s.Memory = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetNodeType(v string) *DescribeRCInstancesResponseBodyRCInstances {
	s.NodeType = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetPublicIp(v string) *DescribeRCInstancesResponseBodyRCInstances {
	s.PublicIp = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetRegionId(v string) *DescribeRCInstancesResponseBodyRCInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetSecurityGroupId(v string) *DescribeRCInstancesResponseBodyRCInstances {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetSpotStrategy(v string) *DescribeRCInstancesResponseBodyRCInstances {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetStatus(v string) *DescribeRCInstancesResponseBodyRCInstances {
	s.Status = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetTagResources(v []*DescribeRCInstancesResponseBodyRCInstancesTagResources) *DescribeRCInstancesResponseBodyRCInstances {
	s.TagResources = v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetTags(v []*DescribeRCInstancesResponseBodyRCInstancesTags) *DescribeRCInstancesResponseBodyRCInstances {
	s.Tags = v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetVpcAttributes(v *DescribeRCInstancesResponseBodyRCInstancesVpcAttributes) *DescribeRCInstancesResponseBodyRCInstances {
	s.VpcAttributes = v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetVpcId(v string) *DescribeRCInstancesResponseBodyRCInstances {
	s.VpcId = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) SetZoneId(v string) *DescribeRCInstancesResponseBodyRCInstances {
	s.ZoneId = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstances) Validate() error {
	if s.TagResources != nil {
		for _, item := range s.TagResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VpcAttributes != nil {
		if err := s.VpcAttributes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRCInstancesResponseBodyRCInstancesTagResources struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeRCInstancesResponseBodyRCInstancesTagResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstancesResponseBodyRCInstancesTagResources) GoString() string {
	return s.String()
}

func (s *DescribeRCInstancesResponseBodyRCInstancesTagResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeRCInstancesResponseBodyRCInstancesTagResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeRCInstancesResponseBodyRCInstancesTagResources) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeRCInstancesResponseBodyRCInstancesTagResources) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeRCInstancesResponseBodyRCInstancesTagResources) SetResourceId(v string) *DescribeRCInstancesResponseBodyRCInstancesTagResources {
	s.ResourceId = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstancesTagResources) SetResourceType(v string) *DescribeRCInstancesResponseBodyRCInstancesTagResources {
	s.ResourceType = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstancesTagResources) SetTagKey(v string) *DescribeRCInstancesResponseBodyRCInstancesTagResources {
	s.TagKey = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstancesTagResources) SetTagValue(v string) *DescribeRCInstancesResponseBodyRCInstancesTagResources {
	s.TagValue = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstancesTagResources) Validate() error {
	return dara.Validate(s)
}

type DescribeRCInstancesResponseBodyRCInstancesTags struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeRCInstancesResponseBodyRCInstancesTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstancesResponseBodyRCInstancesTags) GoString() string {
	return s.String()
}

func (s *DescribeRCInstancesResponseBodyRCInstancesTags) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeRCInstancesResponseBodyRCInstancesTags) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeRCInstancesResponseBodyRCInstancesTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeRCInstancesResponseBodyRCInstancesTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeRCInstancesResponseBodyRCInstancesTags) SetResourceId(v string) *DescribeRCInstancesResponseBodyRCInstancesTags {
	s.ResourceId = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstancesTags) SetResourceType(v string) *DescribeRCInstancesResponseBodyRCInstancesTags {
	s.ResourceType = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstancesTags) SetTagKey(v string) *DescribeRCInstancesResponseBodyRCInstancesTags {
	s.TagKey = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstancesTags) SetTagValue(v string) *DescribeRCInstancesResponseBodyRCInstancesTags {
	s.TagValue = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstancesTags) Validate() error {
	return dara.Validate(s)
}

type DescribeRCInstancesResponseBodyRCInstancesVpcAttributes struct {
	NatIpAddress     *string   `json:"NatIpAddress,omitempty" xml:"NatIpAddress,omitempty"`
	PrivateIpAddress []*string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty" type:"Repeated"`
	VSwitchId        *string   `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId            *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeRCInstancesResponseBodyRCInstancesVpcAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstancesResponseBodyRCInstancesVpcAttributes) GoString() string {
	return s.String()
}

func (s *DescribeRCInstancesResponseBodyRCInstancesVpcAttributes) GetNatIpAddress() *string {
	return s.NatIpAddress
}

func (s *DescribeRCInstancesResponseBodyRCInstancesVpcAttributes) GetPrivateIpAddress() []*string {
	return s.PrivateIpAddress
}

func (s *DescribeRCInstancesResponseBodyRCInstancesVpcAttributes) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeRCInstancesResponseBodyRCInstancesVpcAttributes) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeRCInstancesResponseBodyRCInstancesVpcAttributes) SetNatIpAddress(v string) *DescribeRCInstancesResponseBodyRCInstancesVpcAttributes {
	s.NatIpAddress = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstancesVpcAttributes) SetPrivateIpAddress(v []*string) *DescribeRCInstancesResponseBodyRCInstancesVpcAttributes {
	s.PrivateIpAddress = v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstancesVpcAttributes) SetVSwitchId(v string) *DescribeRCInstancesResponseBodyRCInstancesVpcAttributes {
	s.VSwitchId = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstancesVpcAttributes) SetVpcId(v string) *DescribeRCInstancesResponseBodyRCInstancesVpcAttributes {
	s.VpcId = &v
	return s
}

func (s *DescribeRCInstancesResponseBodyRCInstancesVpcAttributes) Validate() error {
	return dara.Validate(s)
}
