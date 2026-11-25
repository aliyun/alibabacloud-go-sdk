// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWarmPoolInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeWarmPoolInstancesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeWarmPoolInstancesResponseBody
	GetRequestId() *string
	SetScalingGroupId(v string) *DescribeWarmPoolInstancesResponseBody
	GetScalingGroupId() *string
	SetWarmPoolInstances(v []*DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) *DescribeWarmPoolInstancesResponseBody
	GetWarmPoolInstances() []*DescribeWarmPoolInstancesResponseBodyWarmPoolInstances
}

type DescribeWarmPoolInstancesResponseBody struct {
	// The token for the next query. An empty NextToken value indicates that no more results are available.
	//
	// example:
	//
	// AAAAAcSz4VTb1Nq****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E38EB733-D714-4658-8A5F-0688AB68****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the scaling group.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The list of ECS instances in the warm pool.
	WarmPoolInstances []*DescribeWarmPoolInstancesResponseBodyWarmPoolInstances `json:"WarmPoolInstances,omitempty" xml:"WarmPoolInstances,omitempty" type:"Repeated"`
}

func (s DescribeWarmPoolInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWarmPoolInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWarmPoolInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeWarmPoolInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWarmPoolInstancesResponseBody) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeWarmPoolInstancesResponseBody) GetWarmPoolInstances() []*DescribeWarmPoolInstancesResponseBodyWarmPoolInstances {
	return s.WarmPoolInstances
}

func (s *DescribeWarmPoolInstancesResponseBody) SetNextToken(v string) *DescribeWarmPoolInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeWarmPoolInstancesResponseBody) SetRequestId(v string) *DescribeWarmPoolInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWarmPoolInstancesResponseBody) SetScalingGroupId(v string) *DescribeWarmPoolInstancesResponseBody {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeWarmPoolInstancesResponseBody) SetWarmPoolInstances(v []*DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) *DescribeWarmPoolInstancesResponseBody {
	s.WarmPoolInstances = v
	return s
}

func (s *DescribeWarmPoolInstancesResponseBody) Validate() error {
	if s.WarmPoolInstances != nil {
		for _, item := range s.WarmPoolInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeWarmPoolInstancesResponseBodyWarmPoolInstances struct {
	// The time when the ECS instance was added to the warm pool, accurate to the second. The time is in UTC and follows the ISO 8601 standard in the YYYY-MM-DDTHH:mm:ssZ format.
	//
	// example:
	//
	// 2020-05-18T03:11:39Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The ID of the ECS instance.
	//
	// example:
	//
	// i-bp1c2hf**********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the instance launch template.
	//
	// example:
	//
	// lt-uf64k10pok76*****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The version of the instance launch template.
	//
	// example:
	//
	// 2
	LaunchTemplateVersion *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The lifecycle state of the ECS instance in the warm pool. Valid values:
	//
	// - Pending: The ECS instance is being added to the warm pool.
	//
	// - Pending:Wait: The addition of the ECS instance is suspended. If a lifecycle hook that applies to scale-out activities for the warm pool is created in the scaling group, the ECS instance is suspended when it is added to the warm pool and waits for the lifecycle hook to time out.
	//
	// - Removing: The ECS instance is being removed from the warm pool.
	//
	// - Removing:Wait: The removal of the ECS instance is suspended. If a lifecycle hook that applies to scale-in activities for the warm pool is created in the scaling group, the ECS instance is suspended when it is removed from the warm pool and waits for the lifecycle hook to time out.
	//
	// - Running: The ECS instance is successfully added to the warm pool and is in the running state.
	//
	// - Stopped: The ECS instance is successfully added to the warm pool and is in the economical mode.
	//
	// example:
	//
	// Running
	LifecycleState *string `json:"LifecycleState,omitempty" xml:"LifecycleState,omitempty"`
	// The private IP address of the instance.
	//
	// example:
	//
	// 1*.1*.1*.2*
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The ID of the scaling activity when the ECS instance was added to the warm pool.
	//
	// example:
	//
	// asa-2ze87i********
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
	// The ID of the associated scaling configuration.
	//
	// example:
	//
	// asc-2ze3j8*********
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	// The ID of the scaling group.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The instance identity in the scaling group, which corresponds one-to-one with the ID of the ECS instance.
	//
	// example:
	//
	// asi-j6cj1gcte640ekhb****
	ScalingInstanceId *string `json:"ScalingInstanceId,omitempty" xml:"ScalingInstanceId,omitempty"`
	// The zone ID of the ECS instance.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) GoString() string {
	return s.String()
}

func (s *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) GetLaunchTemplateVersion() *string {
	return s.LaunchTemplateVersion
}

func (s *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) GetLifecycleState() *string {
	return s.LifecycleState
}

func (s *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) GetScalingActivityId() *string {
	return s.ScalingActivityId
}

func (s *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) GetScalingConfigurationId() *string {
	return s.ScalingConfigurationId
}

func (s *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) GetScalingInstanceId() *string {
	return s.ScalingInstanceId
}

func (s *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) SetCreatedTime(v string) *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances {
	s.CreatedTime = &v
	return s
}

func (s *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) SetInstanceId(v string) *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) SetLaunchTemplateId(v string) *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances {
	s.LaunchTemplateId = &v
	return s
}

func (s *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) SetLaunchTemplateVersion(v string) *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) SetLifecycleState(v string) *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances {
	s.LifecycleState = &v
	return s
}

func (s *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) SetPrivateIpAddress(v string) *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) SetScalingActivityId(v string) *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances {
	s.ScalingActivityId = &v
	return s
}

func (s *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) SetScalingConfigurationId(v string) *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances {
	s.ScalingConfigurationId = &v
	return s
}

func (s *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) SetScalingGroupId(v string) *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) SetScalingInstanceId(v string) *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances {
	s.ScalingInstanceId = &v
	return s
}

func (s *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) SetZoneId(v string) *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances {
	s.ZoneId = &v
	return s
}

func (s *DescribeWarmPoolInstancesResponseBodyWarmPoolInstances) Validate() error {
	return dara.Validate(s)
}
