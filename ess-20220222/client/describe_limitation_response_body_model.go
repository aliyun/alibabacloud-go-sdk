// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLimitationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxNumberOfAlbServerGroup(v int32) *DescribeLimitationResponseBody
	GetMaxNumberOfAlbServerGroup() *int32
	SetMaxNumberOfDBInstances(v int32) *DescribeLimitationResponseBody
	GetMaxNumberOfDBInstances() *int32
	SetMaxNumberOfLifecycleHooks(v int32) *DescribeLimitationResponseBody
	GetMaxNumberOfLifecycleHooks() *int32
	SetMaxNumberOfLoadBalancers(v int32) *DescribeLimitationResponseBody
	GetMaxNumberOfLoadBalancers() *int32
	SetMaxNumberOfMaxSize(v int32) *DescribeLimitationResponseBody
	GetMaxNumberOfMaxSize() *int32
	SetMaxNumberOfMinSize(v int32) *DescribeLimitationResponseBody
	GetMaxNumberOfMinSize() *int32
	SetMaxNumberOfNlbServerGroup(v int32) *DescribeLimitationResponseBody
	GetMaxNumberOfNlbServerGroup() *int32
	SetMaxNumberOfNotificationConfigurations(v int32) *DescribeLimitationResponseBody
	GetMaxNumberOfNotificationConfigurations() *int32
	SetMaxNumberOfScalingConfigurations(v int32) *DescribeLimitationResponseBody
	GetMaxNumberOfScalingConfigurations() *int32
	SetMaxNumberOfScalingGroups(v int32) *DescribeLimitationResponseBody
	GetMaxNumberOfScalingGroups() *int32
	SetMaxNumberOfScalingInstances(v int32) *DescribeLimitationResponseBody
	GetMaxNumberOfScalingInstances() *int32
	SetMaxNumberOfScalingRules(v int32) *DescribeLimitationResponseBody
	GetMaxNumberOfScalingRules() *int32
	SetMaxNumberOfScheduledTasks(v int32) *DescribeLimitationResponseBody
	GetMaxNumberOfScheduledTasks() *int32
	SetMaxNumberOfVServerGroups(v int32) *DescribeLimitationResponseBody
	GetMaxNumberOfVServerGroups() *int32
	SetRequestId(v string) *DescribeLimitationResponseBody
	GetRequestId() *string
}

type DescribeLimitationResponseBody struct {
	// The maximum number of Application Load Balancer (ALB) server groups that can be attached to a scaling group.
	//
	// >  To view the quota or request a quota increase, go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas).
	//
	// example:
	//
	// 30
	MaxNumberOfAlbServerGroup *int32 `json:"MaxNumberOfAlbServerGroup,omitempty" xml:"MaxNumberOfAlbServerGroup,omitempty"`
	// The maximum number of ApsaraDB RDS instances that can be associated with a scaling group.
	//
	// >  To view the quota or request a quota increase, go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas).
	//
	// example:
	//
	// 30
	MaxNumberOfDBInstances *int32 `json:"MaxNumberOfDBInstances,omitempty" xml:"MaxNumberOfDBInstances,omitempty"`
	// The maximum number of lifecycle hooks that can be created in a scaling group.
	//
	// example:
	//
	// 10
	MaxNumberOfLifecycleHooks *int32 `json:"MaxNumberOfLifecycleHooks,omitempty" xml:"MaxNumberOfLifecycleHooks,omitempty"`
	// The maximum number of Classic Load Balancer (CLB, formerly known as SLB) instances that can be associated with a scaling group.
	//
	// >  To view the quota or request a quota increase, go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas).
	//
	// example:
	//
	// 30
	MaxNumberOfLoadBalancers *int32 `json:"MaxNumberOfLoadBalancers,omitempty" xml:"MaxNumberOfLoadBalancers,omitempty"`
	// The maximum number of instances that can be contained in a scaling group.
	//
	// >  To view the quota or request a quota increase, go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas).
	//
	// example:
	//
	// 2000
	MaxNumberOfMaxSize *int32 `json:"MaxNumberOfMaxSize,omitempty" xml:"MaxNumberOfMaxSize,omitempty"`
	// The minimum number of instances that must be contained in a scaling group. The value of `MaxNumberOfMinSize` must be consistent with the value of `MaxNumberOfMaxSize`.
	//
	// example:
	//
	// 2000
	MaxNumberOfMinSize *int32 `json:"MaxNumberOfMinSize,omitempty" xml:"MaxNumberOfMinSize,omitempty"`
	// The maximum number of Network Load Balancer (NLB) server groups that can be attached to a scaling group.
	//
	// >  To view the quota or request a quota increase, go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas).
	//
	// example:
	//
	// 30
	MaxNumberOfNlbServerGroup *int32 `json:"MaxNumberOfNlbServerGroup,omitempty" xml:"MaxNumberOfNlbServerGroup,omitempty"`
	// The maximum number of notification rules that can be created in a scaling group.
	//
	// example:
	//
	// 6
	MaxNumberOfNotificationConfigurations *int32 `json:"MaxNumberOfNotificationConfigurations,omitempty" xml:"MaxNumberOfNotificationConfigurations,omitempty"`
	// The maximum number of scaling configurations that can be created in a scaling group.
	//
	// >  To view the quota or request a quota increase, go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas).
	//
	// example:
	//
	// 70
	MaxNumberOfScalingConfigurations *int32 `json:"MaxNumberOfScalingConfigurations,omitempty" xml:"MaxNumberOfScalingConfigurations,omitempty"`
	// The maximum number of scaling groups that can be created in a region by using an Alibaba Cloud account.
	//
	// >  To view the quota or request a quota increase, go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas).
	//
	// example:
	//
	// 200
	MaxNumberOfScalingGroups *int32 `json:"MaxNumberOfScalingGroups,omitempty" xml:"MaxNumberOfScalingGroups,omitempty"`
	// The maximum number of Elastic Compute Service (ECS) instances or elastic container instances that can be automatically scaled in a scaling group at the same time.
	//
	// example:
	//
	// 500
	MaxNumberOfScalingInstances *int32 `json:"MaxNumberOfScalingInstances,omitempty" xml:"MaxNumberOfScalingInstances,omitempty"`
	// The maximum number of scaling rules that can be created in a scaling group.
	//
	// >  To view the quota or request a quota increase, go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas).
	//
	// example:
	//
	// 70
	MaxNumberOfScalingRules *int32 `json:"MaxNumberOfScalingRules,omitempty" xml:"MaxNumberOfScalingRules,omitempty"`
	// The maximum number of scheduled tasks that can be created in a region by using an Alibaba Cloud account.
	//
	// >  To view the quota or request a quota increase, go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas).
	//
	// example:
	//
	// 70
	MaxNumberOfScheduledTasks *int32 `json:"MaxNumberOfScheduledTasks,omitempty" xml:"MaxNumberOfScheduledTasks,omitempty"`
	// The maximum number of CLB vServer groups that can be attached to a scaling group.
	//
	// >  To view the quota or request a quota increase, go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas).
	//
	// example:
	//
	// 5
	MaxNumberOfVServerGroups *int32 `json:"MaxNumberOfVServerGroups,omitempty" xml:"MaxNumberOfVServerGroups,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BE9BEB41-E7B8-4C7D-A3CF-2DCB1066****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLimitationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLimitationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLimitationResponseBody) GetMaxNumberOfAlbServerGroup() *int32 {
	return s.MaxNumberOfAlbServerGroup
}

func (s *DescribeLimitationResponseBody) GetMaxNumberOfDBInstances() *int32 {
	return s.MaxNumberOfDBInstances
}

func (s *DescribeLimitationResponseBody) GetMaxNumberOfLifecycleHooks() *int32 {
	return s.MaxNumberOfLifecycleHooks
}

func (s *DescribeLimitationResponseBody) GetMaxNumberOfLoadBalancers() *int32 {
	return s.MaxNumberOfLoadBalancers
}

func (s *DescribeLimitationResponseBody) GetMaxNumberOfMaxSize() *int32 {
	return s.MaxNumberOfMaxSize
}

func (s *DescribeLimitationResponseBody) GetMaxNumberOfMinSize() *int32 {
	return s.MaxNumberOfMinSize
}

func (s *DescribeLimitationResponseBody) GetMaxNumberOfNlbServerGroup() *int32 {
	return s.MaxNumberOfNlbServerGroup
}

func (s *DescribeLimitationResponseBody) GetMaxNumberOfNotificationConfigurations() *int32 {
	return s.MaxNumberOfNotificationConfigurations
}

func (s *DescribeLimitationResponseBody) GetMaxNumberOfScalingConfigurations() *int32 {
	return s.MaxNumberOfScalingConfigurations
}

func (s *DescribeLimitationResponseBody) GetMaxNumberOfScalingGroups() *int32 {
	return s.MaxNumberOfScalingGroups
}

func (s *DescribeLimitationResponseBody) GetMaxNumberOfScalingInstances() *int32 {
	return s.MaxNumberOfScalingInstances
}

func (s *DescribeLimitationResponseBody) GetMaxNumberOfScalingRules() *int32 {
	return s.MaxNumberOfScalingRules
}

func (s *DescribeLimitationResponseBody) GetMaxNumberOfScheduledTasks() *int32 {
	return s.MaxNumberOfScheduledTasks
}

func (s *DescribeLimitationResponseBody) GetMaxNumberOfVServerGroups() *int32 {
	return s.MaxNumberOfVServerGroups
}

func (s *DescribeLimitationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfAlbServerGroup(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfAlbServerGroup = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfDBInstances(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfDBInstances = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfLifecycleHooks(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfLifecycleHooks = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfLoadBalancers(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfLoadBalancers = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfMaxSize(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfMaxSize = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfMinSize(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfMinSize = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfNlbServerGroup(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfNlbServerGroup = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfNotificationConfigurations(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfNotificationConfigurations = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfScalingConfigurations(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfScalingConfigurations = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfScalingGroups(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfScalingGroups = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfScalingInstances(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfScalingInstances = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfScalingRules(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfScalingRules = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfScheduledTasks(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfScheduledTasks = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfVServerGroups(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfVServerGroups = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetRequestId(v string) *DescribeLimitationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLimitationResponseBody) Validate() error {
	return dara.Validate(s)
}
