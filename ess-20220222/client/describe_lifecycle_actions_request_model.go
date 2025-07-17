// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLifecycleActionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLifecycleActionStatus(v string) *DescribeLifecycleActionsRequest
	GetLifecycleActionStatus() *string
	SetMaxResults(v int32) *DescribeLifecycleActionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeLifecycleActionsRequest
	GetNextToken() *string
	SetOwnerId(v int64) *DescribeLifecycleActionsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLifecycleActionsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeLifecycleActionsRequest
	GetResourceOwnerAccount() *string
	SetScalingActivityId(v string) *DescribeLifecycleActionsRequest
	GetScalingActivityId() *string
}

type DescribeLifecycleActionsRequest struct {
	// The status of the lifecycle action. Valid values:
	//
	// 	- If a lifecycle action is in the Pending state, Elastic Compute Service (ECS) instances are waiting to be added to a scaling group or waiting to be removed from a scaling group.
	//
	// 	- If a lifecycle action is in the Timeout state, the lifecycle hook that triggers the lifecycle action ends, and ECS instances are added to or removed from the scaling group.
	//
	// 	- If a lifecycle action is in the Completed state, you manually end the lifecycle hook that triggers the lifecycle action ahead of schedule.
	//
	// example:
	//
	// Pending
	LifecycleActionStatus *string `json:"LifecycleActionStatus,omitempty" xml:"LifecycleActionStatus,omitempty"`
	// The maximum number of entries to return on each page. Valid values: 1 to 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to specify the lifecycle action from which the query starts.
	//
	// For example, after the first 10 lifecycle actions are queried, the query starts from the 11th lifecycle action. Set this parameter to the NextToken value that is returned in the previous API call. If you do not specify this parameter, the query starts from the beginning.
	//
	// example:
	//
	// AAAAAcSz4VTb1Nq****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group.
	//
	// example:
	//
	// cn-qingdao
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the scaling activity.
	//
	// This parameter is required.
	//
	// example:
	//
	// asa-bp17mug9t0pegagw****
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s DescribeLifecycleActionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLifecycleActionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLifecycleActionsRequest) GetLifecycleActionStatus() *string {
	return s.LifecycleActionStatus
}

func (s *DescribeLifecycleActionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeLifecycleActionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeLifecycleActionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLifecycleActionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLifecycleActionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLifecycleActionsRequest) GetScalingActivityId() *string {
	return s.ScalingActivityId
}

func (s *DescribeLifecycleActionsRequest) SetLifecycleActionStatus(v string) *DescribeLifecycleActionsRequest {
	s.LifecycleActionStatus = &v
	return s
}

func (s *DescribeLifecycleActionsRequest) SetMaxResults(v int32) *DescribeLifecycleActionsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeLifecycleActionsRequest) SetNextToken(v string) *DescribeLifecycleActionsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeLifecycleActionsRequest) SetOwnerId(v int64) *DescribeLifecycleActionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLifecycleActionsRequest) SetRegionId(v string) *DescribeLifecycleActionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLifecycleActionsRequest) SetResourceOwnerAccount(v string) *DescribeLifecycleActionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLifecycleActionsRequest) SetScalingActivityId(v string) *DescribeLifecycleActionsRequest {
	s.ScalingActivityId = &v
	return s
}

func (s *DescribeLifecycleActionsRequest) Validate() error {
	return dara.Validate(s)
}
