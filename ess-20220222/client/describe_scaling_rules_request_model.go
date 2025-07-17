// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeScalingRulesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeScalingRulesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeScalingRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeScalingRulesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeScalingRulesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeScalingRulesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeScalingRulesRequest
	GetResourceOwnerId() *int64
	SetScalingGroupId(v string) *DescribeScalingRulesRequest
	GetScalingGroupId() *string
	SetScalingRuleAris(v []*string) *DescribeScalingRulesRequest
	GetScalingRuleAris() []*string
	SetScalingRuleIds(v []*string) *DescribeScalingRulesRequest
	GetScalingRuleIds() []*string
	SetScalingRuleNames(v []*string) *DescribeScalingRulesRequest
	GetScalingRuleNames() []*string
	SetScalingRuleType(v string) *DescribeScalingRulesRequest
	GetScalingRuleType() *string
	SetShowAlarmRules(v bool) *DescribeScalingRulesRequest
	GetShowAlarmRules() *bool
}

type DescribeScalingRulesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the scaling group to which the scaling rules that you want to query belong.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the scaling group.
	//
	// example:
	//
	// asg-bp1ffogfdauy0jw0****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The unique identifiers of the scaling rules that you want to query.
	ScalingRuleAris []*string `json:"ScalingRuleAris,omitempty" xml:"ScalingRuleAris,omitempty" type:"Repeated"`
	// The IDs of the scaling rules that you want to query.
	ScalingRuleIds []*string `json:"ScalingRuleIds,omitempty" xml:"ScalingRuleIds,omitempty" type:"Repeated"`
	// The names of the scaling rules that you want to query.
	ScalingRuleNames []*string `json:"ScalingRuleNames,omitempty" xml:"ScalingRuleNames,omitempty" type:"Repeated"`
	// The type of the scaling rule. Valid values:
	//
	// 	- SimpleScalingRule: adjusts the number of ECS instances based on the values of the AdjustmentType and AdjustmentValue parameters.
	//
	// 	- TargetTrackingScalingRule: calculates the number of ECS instances that need to be scaled in a dynamic manner and maintains the value of a predefined metric close to the value of the TargetValue parameter.
	//
	// 	- StepScalingRule: scales ECS instances in steps based on the specified thresholds and metric values.
	//
	// 	- PredictiveScalingRule: uses machine learning to analyze historical monitoring data of the scaling group and predicts the future values of metrics. In addition, Auto Scaling automatically creates scheduled tasks to adjust the boundary values for the scaling group.
	//
	// example:
	//
	// SimpleScalingRule
	ScalingRuleType *string `json:"ScalingRuleType,omitempty" xml:"ScalingRuleType,omitempty"`
	// Specifies whether to return the event-triggered tasks that are associated with the scaling rule. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ShowAlarmRules *bool `json:"ShowAlarmRules,omitempty" xml:"ShowAlarmRules,omitempty"`
}

func (s DescribeScalingRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeScalingRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeScalingRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeScalingRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeScalingRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeScalingRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeScalingRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeScalingRulesRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeScalingRulesRequest) GetScalingRuleAris() []*string {
	return s.ScalingRuleAris
}

func (s *DescribeScalingRulesRequest) GetScalingRuleIds() []*string {
	return s.ScalingRuleIds
}

func (s *DescribeScalingRulesRequest) GetScalingRuleNames() []*string {
	return s.ScalingRuleNames
}

func (s *DescribeScalingRulesRequest) GetScalingRuleType() *string {
	return s.ScalingRuleType
}

func (s *DescribeScalingRulesRequest) GetShowAlarmRules() *bool {
	return s.ShowAlarmRules
}

func (s *DescribeScalingRulesRequest) SetOwnerAccount(v string) *DescribeScalingRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetOwnerId(v int64) *DescribeScalingRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetPageNumber(v int32) *DescribeScalingRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetPageSize(v int32) *DescribeScalingRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetRegionId(v string) *DescribeScalingRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetResourceOwnerAccount(v string) *DescribeScalingRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetResourceOwnerId(v int64) *DescribeScalingRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetScalingGroupId(v string) *DescribeScalingRulesRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetScalingRuleAris(v []*string) *DescribeScalingRulesRequest {
	s.ScalingRuleAris = v
	return s
}

func (s *DescribeScalingRulesRequest) SetScalingRuleIds(v []*string) *DescribeScalingRulesRequest {
	s.ScalingRuleIds = v
	return s
}

func (s *DescribeScalingRulesRequest) SetScalingRuleNames(v []*string) *DescribeScalingRulesRequest {
	s.ScalingRuleNames = v
	return s
}

func (s *DescribeScalingRulesRequest) SetScalingRuleType(v string) *DescribeScalingRulesRequest {
	s.ScalingRuleType = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetShowAlarmRules(v bool) *DescribeScalingRulesRequest {
	s.ShowAlarmRules = &v
	return s
}

func (s *DescribeScalingRulesRequest) Validate() error {
	return dara.Validate(s)
}
